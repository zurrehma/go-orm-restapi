package app

import (
	"bytes"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/zurrehma/go-orm-restapi/app/model"
	"github.com/zurrehma/go-orm-restapi/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type App struct {
	Router         *mux.Router
	DB             *gorm.DB
	Logger         *log.Logger
	Buf            bytes.Buffer
	ProjectHandler Routes
	TaskHandler    Routes
}

func (app *App) Initialize(config *config.Config) {

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
	app.DB = model.DBMigrate(db)
	app.Logger = log.New(&app.Buf, "todoapp: ", log.Lshortfile)
	app.SetRoutes()
}

func (app *App) SetRoutes() {
	app.SetHandler("/projects", app.GetAllProjects, "GET")
	app.SetHandler("/projects", app.CreateProject, "POST")
	app.SetHandler("/projects/{title}", app.GetProject, "GET")
	app.SetHandler("/projects/{title}", app.UpdateProject, "PUT")
	app.SetHandler("/projects/{title}", app.DeleteProject, "DELETE")

	app.SetHandler("/projects/{title}/tasks", app.GetAllTasks, "GET")
	app.SetHandler("/projects/{title}/tasks", app.CreateTask, "POST")
	app.SetHandler("/projects/{title}/tasks/{id}", app.GetTask, "GET")
	app.SetHandler("/projects/{title}/tasks/{id}", app.UpdateTask, "PUT")
	app.SetHandler("/projects/{title}/tasks/{id}", app.DeleteTask, "DELETE")
}

func (app *App) SetHandler(path string, f func(w http.ResponseWriter, r *http.Request), method string) {
	app.Router.HandleFunc(path, f).Methods(method)
}

/*
** Project Handlers
 */
func (a *App) GetAllProjects(w http.ResponseWriter, r *http.Request) {
	a.ProjectHandler.GetAllObjects(a, w, r)
}

func (a *App) CreateProject(w http.ResponseWriter, r *http.Request) {
	a.ProjectHandler.CreateObject(a, w, r)
}

func (a *App) GetProject(w http.ResponseWriter, r *http.Request) {
	a.ProjectHandler.GetObject(a, w, r)
}

func (a *App) UpdateProject(w http.ResponseWriter, r *http.Request) {
	a.ProjectHandler.UpdateObject(a, w, r)
}

func (a *App) DeleteProject(w http.ResponseWriter, r *http.Request) {
	a.ProjectHandler.DeleteObject(a, w, r)
}

/*
** Tasks Handlers
 */
func (a *App) GetAllTasks(w http.ResponseWriter, r *http.Request) {
	a.TaskHandler.GetAllObjects(a, w, r)
}

func (a *App) CreateTask(w http.ResponseWriter, r *http.Request) {
	a.TaskHandler.CreateObject(a, w, r)
}

func (a *App) GetTask(w http.ResponseWriter, r *http.Request) {
	a.TaskHandler.GetObject(a, w, r)
}

func (a *App) UpdateTask(w http.ResponseWriter, r *http.Request) {
	a.TaskHandler.UpdateObject(a, w, r)
}

func (a *App) DeleteTask(w http.ResponseWriter, r *http.Request) {
	a.TaskHandler.DeleteObject(a, w, r)
}

func (app *App) Run(host string) {
	log.Fatal(http.ListenAndServe(host, app.Router))
}
