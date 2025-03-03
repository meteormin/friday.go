package main

import (
	"flag"
	"fmt"
	"github.com/meteormin/friday.go/api"
	"github.com/meteormin/friday.go/internal/adapter/rest"
	"github.com/meteormin/friday.go/internal/bootstrap"
	"github.com/meteormin/friday.go/internal/core"
	"github.com/meteormin/friday.go/internal/core/http"
	"os"
	"os/signal"
	"strconv"
	"strings"
	"syscall"
)

// Friday.go API
// @title Friday.go API
// @version {{.Version}}
// @description Friday.go API
// @contact.name meteormin
// @contact.url https://github.com/meteormin/friday.go
// @contact.email miniyu97@gmail.com
// @schemes http https
// @host {{.Host}}
// @BasePath /
// @tag.name auth
// @tag.name users
// @tag.name upload-file
// @tag.name posts
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func main() {
	cfgPath := flag.String("config", "config.yml", "path to config file")
	flag.Parse()

	// initialize app
	cfg := bootstrap.Initialize(*cfgPath)
	if cfg.Server.Port <= 0 {
		cfg.Server.Port = 8080
	}

	app := http.NewApp(cfg.App)

	// set swagger info
	apiInfo(cfg.Server.Host, cfg.App.Version, strconv.Itoa(cfg.Server.Port))

	// set routes
	rest.RegisterRoutes(app)

	l := core.Logger()
	taskScheduler := core.Scheduler()
	storage := core.Storage()

	// Listen from a different goroutine
	go func() {
		taskScheduler.Start()
		if err := app.Listen(":" + strconv.Itoa(cfg.Server.Port)); err != nil {
			l.Panic(err)
		}
	}()

	c := make(chan os.Signal, 1)                    // Create channel to signify a signal being sent
	signal.Notify(c, os.Interrupt, syscall.SIGTERM) // When an interrupt or termination signal is sent, notify the channel

	<-c // This blocks the main thread until an interrupt is received
	fmt.Println("Gracefully shutting down...")
	err := app.Shutdown()
	if err != nil {
		l.Error("Graceful shutdown failed", err)
	} else {
		l.Info("Fiber was successful shutdown.")
	}

	fmt.Println("Running cleanup tasks...")

	err = taskScheduler.Shutdown()
	if err != nil {
		l.Error("Scheduler shutdown failed", err)
	} else {
		l.Info("Scheduler was successful shutdown.")
	}

	err = storage.Close()
	if err != nil {
		l.Error("Storage shutdown failed", err)
	} else {
		l.Error("Storage was successful shutdown.")
	}
}

// apiInfo sets the API information based on the given port.
//
// It takes a string parameter 'port' and does not return anything.
func apiInfo(host string, version string, port string) {

	if host == "" {
		host = fmt.Sprintf("%s:%s", "localhost", port)
	}

	schemaHost := strings.Split(host, "://")
	if len(schemaHost) > 1 {
		host = schemaHost[1]
	}

	api.SwaggerInfo.Title = "Friday.go API"
	api.SwaggerInfo.Version = version
	api.SwaggerInfo.Schemes = []string{"http", "https"}
	api.SwaggerInfo.Host = host
	api.SwaggerInfo.BasePath = "/"
	api.SwaggerInfo.SwaggerTemplate = api.OpenAPITemplate
}
