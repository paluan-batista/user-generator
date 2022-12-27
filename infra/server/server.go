package server

import (
	"fmt"
	"github.com/joomcode/errorx"
	"github.com/labstack/echo/v4"
	"user-generator/infra/config"
	"user-generator/infra/log"
	"user-generator/internal/api"
)

const (
	failToStartServer     = "Failed to start the server"
	failToStartMetaServer = "Failed to start the meta server"
	startingServer        = "Started a server at http(s)://%s"
	startingMetaServer    = "Started a meta server at http(s)://%s"
)

var logger = log.NewLogger()

func newApp() *echo.Echo {
	app := echo.New()
	app.HideBanner = true
	app.HidePort = true
	api.NewUserController().Register(app)

	return app
}

func newMetaApp() *echo.Echo {
	app := echo.New()

	app.HideBanner = true
	app.HidePort = true
	NewHealth().Register(app)
	return app
}

func StartServer() {
	host := config.GetServerConfig().Host
	logger.Info(fmt.Sprintf(startingServer, host))
	err := newApp().Start(host)
	logger.Fatal(errorx.InternalError.Wrap(err, failToStartServer).Error())
}

func StartMetaServer() {
	host := config.GetServerConfig().MetaHost
	logger.Info(fmt.Sprintf(startingMetaServer, host))
	err := newMetaApp().Start(host)
	logger.Fatal(errorx.InternalError.Wrap(err, failToStartMetaServer).Error())
}
