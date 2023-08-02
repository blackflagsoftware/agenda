package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/blackflagsoftware/agenda/config"
	ae "github.com/blackflagsoftware/agenda/internal/api_error"
	m "github.com/blackflagsoftware/agenda/internal/middleware"
	"github.com/blackflagsoftware/agenda/internal/v1/agenda"
	"github.com/blackflagsoftware/agenda/internal/v1/announcement"
	"github.com/blackflagsoftware/agenda/internal/v1/bishopbusiness"
	"github.com/blackflagsoftware/agenda/internal/v1/defaultcalling"
	"github.com/blackflagsoftware/agenda/internal/v1/hymn"
	"github.com/blackflagsoftware/agenda/internal/v1/newmember"
	"github.com/blackflagsoftware/agenda/internal/v1/ordinance"
	"github.com/blackflagsoftware/agenda/internal/v1/roles"
	"github.com/blackflagsoftware/agenda/internal/v1/roleuser"
	"github.com/blackflagsoftware/agenda/internal/v1/speaker"
	"github.com/blackflagsoftware/agenda/internal/v1/visitor"
	"github.com/blackflagsoftware/agenda/internal/v1/wardbusinessrel"
	"github.com/blackflagsoftware/agenda/internal/v1/wardbusinesssus"
	mig "github.com/blackflagsoftware/agenda/tools/migration/src"
	p "github.com/labstack/echo-contrib/prometheus"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	// --- replace main header text - do not remove ---
)

func main() {
	setPidFile()

	// argument flag
	var restPort string
	flag.StringVar(&restPort, "restPort", "", "the port number used for the REST listener")

	flag.Parse()

	if restPort == "" {
		restPort = config.RestPort
	}
	if config.UseMigration {
		err := os.MkdirAll(config.MigrationPath, 0744)
		if err != nil {
			m.Default.Printf("Unable to make scripts/migrations directory structure: %s\n", err)
		}
		c := mig.Connection{
			Host:           config.SqlitePath,
			MigrationPath:  config.MigrationPath,
			SkipInitialize: config.MigrationSkipInit,
			Engine:         "sqlite",
		}
		if err := mig.StartMigration(c); err != nil {
			m.Default.Panicf("Migration failed due to: %s", err)
		}
	}

	e := echo.New()
	e.HTTPErrorHandler = ae.ErrorHandler // set echo's error handler
	if !strings.Contains(config.Env, "prod") {
		m.Default.Infoln("Logging set to debug...")
		e.Debug = true
		e.Use(m.DebugHandler)
	}
	e.Use(
		middleware.Recover(),
		middleware.CORSWithConfig(middleware.CORSConfig{
			AllowOrigins: []string{"*"},
		}),
		m.Handler,
	)
	if config.EnableMetrics {
		prom := p.NewPrometheus("echo", nil)
		prom.Use(e)
	}

	// set all non-endpoints here
	e.GET("/", Index)
	e.HEAD("/status", ServerStatus) // for traditional server check
	e.GET("/liveness", Liveness)    // for k8s liveness
	e.Static("/documents", config.DocumentDir)

	InitializeRoutes(e)

	e.Start(fmt.Sprintf(":%s", restPort))
}

func setPidFile() {
	// purpose: to set the starting applications pid number to file
	if pidFile, err := os.Create(config.PidPath); err != nil {
		m.Default.Panicln("Unable to create pid file...")
	} else if _, err := pidFile.Write([]byte(fmt.Sprintf("%d", os.Getpid()))); err != nil {
		m.Default.Panicln("Unable to write pid to file...")
	}
}

func Index(c echo.Context) error {
	return c.String(http.StatusOK, "Welcome to the AGENDA API")
}

func ServerStatus(c echo.Context) error {
	c.Response().Header().Add("AGENDA", config.AppVersion)
	c.Response().WriteHeader(http.StatusOK)
	return nil
}

func Liveness(c echo.Context) error {
	return c.String(http.StatusOK, "live")
}

func InitializeRoutes(e *echo.Echo) {
	// initialize all routes here
	routeGroup := e.Group("v1") // change to match your uri prefix
	agenda.InitializeRest(routeGroup)
	visitor.InitializeRest(routeGroup)
	wardbusinessrel.InitializeRest(routeGroup)
	wardbusinesssus.InitializeRest(routeGroup)
	bishopbusiness.InitializeRest(routeGroup)
	newmember.InitializeRest(routeGroup)
	ordinance.InitializeRest(routeGroup)
	speaker.InitializeRest(routeGroup)
	announcement.InitializeRest(routeGroup)
	hymn.InitializeRest(routeGroup)
	defaultcalling.InitializeRest(routeGroup)
	roles.InitializeRest(routeGroup)
	roleuser.InitializeRest(routeGroup)
	// --- replace server text - do not remove ---

}