package main

import (
	"flag"
	"fmt"

	"service/api/rest/internal/config"
	"service/api/rest/internal/handler"
	"service/api/rest/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "etc/service-api.yaml", "the config file")

type App struct {
	h *handler.Handler
}

func NewApp(h *handler.Handler) *App {
	return &App{
		h: h,
	}
}

func (a *App) Start() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	ctx := svc.NewServiceContext(c)
	a.h.RegisterHandlers(server, ctx)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}

func main() {
	app := initApp()
	app.Start()
}
