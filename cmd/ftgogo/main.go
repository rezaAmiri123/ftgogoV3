package main

import (
	"database/sql"
	"fmt"
	"net/http"
	_ "net/http/pprof"
	"os"

	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/rezaAmiri123/ftgogoV3/accounting"
	"github.com/rezaAmiri123/ftgogoV3/consumer"
	"github.com/rezaAmiri123/ftgogoV3/cosec"
	customerweb "github.com/rezaAmiri123/ftgogoV3/customer-web"
	"github.com/rezaAmiri123/ftgogoV3/delivery"
	"github.com/rezaAmiri123/ftgogoV3/internal/config"
	"github.com/rezaAmiri123/ftgogoV3/internal/system"
	"github.com/rezaAmiri123/ftgogoV3/internal/web/swagger"
	"github.com/rezaAmiri123/ftgogoV3/kitchen"
	"github.com/rezaAmiri123/ftgogoV3/migrations"
	"github.com/rezaAmiri123/ftgogoV3/order"
	"github.com/rezaAmiri123/ftgogoV3/restaurant"
	storeweb "github.com/rezaAmiri123/ftgogoV3/store-web"
)

type monolith struct {
	*system.System
	modules []system.Module
}

func main() {
	if err := run(); err != nil {
		fmt.Printf("mallbots exitted abnormally: %s\n", err.Error())
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

	s, err := system.NewSystem(cfg)
	if err != nil {
		return err
	}

	m := monolith{
		System: s,
		modules: []system.Module{
			&accounting.Module{},
			&consumer.Module{},
			&restaurant.Module{},
			&kitchen.Module{},
			&delivery.Module{},
			&order.Module{},
			&cosec.Module{},
			&customerweb.Module{},
			&storeweb.Module{},
		},
	}
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			return
		}
	}(m.DB())
	err = m.MigrateDB(migrations.FS)
	if err != nil {
		return err
	}
	if err = m.startupModules(); err != nil {
		return err
	}

	// Mount general web resources
	m.Mux().Mount("/", http.FileServer(http.FS(swagger.WebUI)))

	fmt.Println("started ftgogo application")
	defer fmt.Println("stopped ftgogo application")

	m.Waiter().Add(
		m.WaitForWeb,
		m.WaitForRPC,
		m.WaitForStream,
	)

	// enable profiler
	go func() {
		http.ListenAndServe(":6060", nil)
	}()

	return m.Waiter().Wait()
}

func (m *monolith) startupModules() error {
	for _, module := range m.modules {
		ctx := m.Waiter().Context()
		if err := module.Startup(ctx, m); err != nil {
			return err
		}
	}
	return nil
}
