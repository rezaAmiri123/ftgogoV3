package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/rezaAmiri123/ftgogoV3/accounting"
	"github.com/rezaAmiri123/ftgogoV3/consumer"
	customerweb "github.com/rezaAmiri123/ftgogoV3/customer-web"
	"github.com/rezaAmiri123/ftgogoV3/delivery"
	"github.com/rezaAmiri123/ftgogoV3/internal/config"
	"github.com/rezaAmiri123/ftgogoV3/internal/logger"
	"github.com/rezaAmiri123/ftgogoV3/internal/monolith"
	"github.com/rezaAmiri123/ftgogoV3/internal/rpc"
	"github.com/rezaAmiri123/ftgogoV3/internal/waiter"
	"github.com/rezaAmiri123/ftgogoV3/internal/web"
	"github.com/rezaAmiri123/ftgogoV3/internal/web/swagger"
	"github.com/rezaAmiri123/ftgogoV3/kitchen"
	"github.com/rezaAmiri123/ftgogoV3/order"
	"github.com/rezaAmiri123/ftgogoV3/restaurant"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	if err := run(); err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}
func run() (err error) {
	var cfg config.AppConfig
	// parse config/env/...
	cfg, err = config.InitConfig()
	if err != nil {
		return err
	}

	m := app{cfg: cfg}

	// init infrastructure...
	m.db, err = sql.Open("pgx", cfg.PG.Conn)
	if err != nil {
		return err
	}
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			return
		}
	}(m.db)
	m.logger = logger.New(logger.LogConfig{
		Environment: cfg.Environment,
		LogLevel:    logger.Level(cfg.LogLevel),
	})
	m.rpc = initRpc(cfg.Rpc)
	m.mux = initMux(cfg.Web)
	m.waiter = waiter.New(waiter.CatchSignals())

	// init modules
	m.modules = []monolith.Module{
		&accounting.Module{},
		&consumer.Module{},
		&restaurant.Module{},
		&kitchen.Module{},
		&delivery.Module{},
		&order.Module{},
		&customerweb.Module{},
	}

	if err = m.startupModules(); err != nil {
		return err
	}

	// Mount general web resources
	m.mux.Mount("/", http.FileServer(http.FS(swagger.WebUI)))

	fmt.Println("started ftgogo application")
	defer fmt.Println("stopped ftgogo application")

	m.waiter.Add(
		m.waitForWeb,
		m.waitForRPC,
	)

	return m.waiter.Wait()
}

func initRpc(_ rpc.RpcConfig) *grpc.Server {
	server := grpc.NewServer()
	reflection.Register(server)

	return server
}

func initMux(_ web.WebConfig) *chi.Mux {
	mux := chi.NewMux()

	mux.Use(
		middleware.Recoverer,
		middleware.Compress(5),
		middleware.Timeout(time.Second*60),
		middleware.Heartbeat("/liveness"),
	)

	// secure
	mux.Use(
		middleware.SetHeader("X-Content-Type-Options", "nosniff"),
		middleware.SetHeader("X-Frame-Options", "SAMEORIGIN"),
		middleware.SetHeader("X-XSS-Protection", "1; mode-block"),
		middleware.NoCache,
	)

	// cors
	mux.Use(cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	}).Handler)

	return mux
}
