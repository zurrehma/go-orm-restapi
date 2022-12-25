package main

import (
	"bytes"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var buf bytes.Buffer

type App struct {
	Router *mux.Router
	DB     *gorm.DB
	Logger *log.Logger
}

func (app *App) Initialize(config *Config) {

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.DB.Username,
		config.DB.Password,
		config.DB.IPAddr,
		config.DB.Port,
		config.DB.AppName)

	app.Router = mux.NewRouter()
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Database connection Issue: ", err)
	}
	app.DB = dbMigrate(db)
	app.Logger = log.New(&buf, "todoapp: ", log.Lshortfile)
	app.SetRoutes()
}

func (app *App) SetRoutes() {
	app.SetHandler("/projects", app.getAllProjects, "GET")
	app.SetHandler("/projects", app.createProject, "POST")
	app.SetHandler("/projects/{title}", app.getProject, "GET")
	app.SetHandler("/projects/{title}", app.updateProject, "PUT")
	app.SetHandler("/projects/{title}", app.deleteProject, "DELETE")
}

func (app *App) SetHandler(path string, f func(w http.ResponseWriter, r *http.Request), method string) {
	app.Router.HandleFunc(path, f).Methods(method)
}

func (app *App) Run(host string) {
	log.Fatal(http.ListenAndServe(host, app.Router))
}
