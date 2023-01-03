package main

import (
	"github.com/zurrehma/go-orm-restapi/app"
	"github.com/zurrehma/go-orm-restapi/app/handler"
	"github.com/zurrehma/go-orm-restapi/config"
)

func main() {

	app := &app.App{
		ProjectHandler: &handler.ProjectHandler{},
		TaskHandler:    &handler.TaskHandler{},
	}
	app.Initialize(config.GetConfig())
	app.Run(":3000")
}
