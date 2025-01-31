package system

import (
	"context"
	"database/sql"
	"fmt"
	"io/fs"
	"net"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/nats-io/nats.go"
	"github.com/pressly/goose/v3"
	"github.com/rezaAmiri123/ftgogoV3/internal/config"
	"github.com/rezaAmiri123/ftgogoV3/internal/logger"
	"github.com/rezaAmiri123/ftgogoV3/internal/waiter"
	"github.com/rs/zerolog"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var _ Service = (*System)(nil)

type System struct {
	cfg    config.AppConfig
	db     *sql.DB
	nc     *nats.Conn
	js     nats.JetStreamContext
	mux    *chi.Mux
	rpc    *grpc.Server
	waiter waiter.Waiter
	logger zerolog.Logger
}

func NewSystem(cfg config.AppConfig)(*System,error){
	s := &System{ cfg: cfg}

	if err := s.initDB();err!= nil{
		return nil,err
	}

	if err:= s.initJS();err!= nil{
		return nil,err
	}

	s.initMux()
	s.initRpc()
	s.initWaiter()
	s.initLogger()

	return s,nil
}

func (s *System) Config() config.AppConfig  { return s.cfg }
func (s *System) DB() *sql.DB               { return s.db }
func (s *System) JS() nats.JetStreamContext { return s.js }
func (s *System) Mux() *chi.Mux             { return s.mux }
func (s *System) RPC() *grpc.Server         { return s.rpc }
func (s *System) Waiter() waiter.Waiter     { return s.waiter }
func (s *System) Logger() zerolog.Logger    { return s.logger }

func (s *System) initDB() (err error) {
	s.db, err = sql.Open("pgx", s.cfg.PG.Conn)
	return err
}

func (s *System) MigrateDB(fs fs.FS) error {
	goose.SetBaseFS(fs)
	if err := goose.SetDialect("postgres"); err != nil {
		return err
	}
	// if err := goose.Up(s.db,".");err!= nil{
	// 	return err
	// }
	return nil
}

func (s *System) initJS() (err error) {
	s.nc, err = nats.Connect(s.cfg.Nats.URL)
	if err != nil {
		return err
	}
	s.js, err = s.nc.JetStream()
	if err != nil {
		return err
	}

	_, err = s.js.AddStream(&nats.StreamConfig{
		Name:     s.cfg.Nats.Stream,
		Subjects: []string{fmt.Sprintf("%s.>", s.cfg.Nats.Stream)},
	})

	return err
}

func (s *System) initLogger() {
	s.logger = logger.New(logger.LogConfig{
		Environment: s.cfg.Environment,
		LogLevel: logger.Level(s.cfg.LogLevel),
	})
}

func (s *System) initMux() {
	s.mux = chi.NewMux()
	s.mux.Use(
		middleware.Recoverer,
		middleware.Compress(5),
		middleware.Timeout(time.Second*60),
		middleware.Heartbeat("/liveness"),
	)

	// secure
	s.mux.Use(
		middleware.SetHeader("X-Content-Type-Options", "nosniff"),
		middleware.SetHeader("X-Frame-Options", "SAMEORIGIN"),
		middleware.SetHeader("X-XSS-Protection", "1; mode-block"),
		middleware.NoCache,
	)

	// cors
	s.mux.Use(cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	}).Handler)

}

func (s *System) initRpc() {
	s.rpc = grpc.NewServer()
	reflection.Register(s.rpc)
}

func (s *System) initWaiter() {
	s.waiter = waiter.New(waiter.CatchSignals())
}

func (s *System)WaitForWeb(ctx context.Context)error{
	webServer := &http.Server{
		Addr: s.cfg.Web.Address(),
		Handler: s.mux,
	}

	group,gCtx := errgroup.WithContext(ctx)
	group.Go(func() error {
		fmt.Printf("web server started; listening at http://localhost%s\n", s.cfg.Web.Port)
		defer fmt.Println("web server shutdown")
		if err := webServer.ListenAndServe();err!= nil&&err!= http.ErrServerClosed{
			return err
		}
		return nil
	})
	group.Go(func() error {
		<-gCtx.Done()
		fmt.Println("web server to be shutdown")
		ctx,cancel := context.WithTimeout(context.Background(),s.cfg.ShutdownTimeout)
		defer cancel()
		if err := webServer.Shutdown(ctx);err!=nil{
			return err
		}
		return nil
	})

	return group.Wait()
}

func(s *System)WaitForRPC(ctx context.Context)error{
	listener,err := net.Listen("tcp", s.cfg.Rpc.Address())
	if err!= nil{
		return err
	}

	group,gCtx := errgroup.WithContext(ctx)
	group.Go(func() error {
		fmt.Println("rpc server started")
		defer fmt.Println("rpc server shutdown")
		if err:= s.RPC().Serve(listener);err!= nil&&err != grpc.ErrServerStopped{
			return err
		}
		return nil
	})
	group.Go(func() error {
		<-gCtx.Done()
		fmt.Println("rpc server to be shutdown")
		stopped := make(chan struct{})
		go func(){
			s.RPC().GracefulStop()
			close(stopped)
		}()
		timeout := time.NewTimer(s.cfg.ShutdownTimeout)
		select{
		case <-timeout.C:
			// Force it to stop
			s.RPC().Stop()
			return fmt.Errorf("rpc server failed to stop gracefully")
		case <-stopped:
			return nil
		}
	})

	return group.Wait()
}

func(s *System)WaitForStream(ctx context.Context)error{
	closed := make(chan struct{})
	s.nc.SetClosedHandler(func(c *nats.Conn) {
		close(closed)
	})
	group,gCtx := errgroup.WithContext(ctx)
	group.Go(func() error {
		fmt.Println("message stream started")
		defer fmt.Println("message stream stopped")
		<-closed
		return nil
	})
	group.Go(func() error {
		<-gCtx.Done()
		return s.nc.Drain()
	})
	return group.Wait()
}
