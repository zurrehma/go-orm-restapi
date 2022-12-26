package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func (app *App) getAllTasks(w http.ResponseWriter, r *http.Request) {
	app.Logger.Println("FunctionName: getAllTasks")
	fmt.Print(&buf)
	vars := mux.Vars(r)
	title := vars["title"]
	project := getProjectOR404(title, app.DB, w)
	if project == nil {
		return
	}
	tasks := []Task{}
	if err := app.DB.Model(&project).Association("Tasks").Find(&tasks); err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		app.Logger.Println(err)
		fmt.Print(&buf)
		return
	}
	respondJson(w, http.StatusOK, tasks)
}

func (app *App) createTask(w http.ResponseWriter, r *http.Request) {
	app.Logger.Println("FunctionName: createTask")
	fmt.Print(&buf)
	vars := mux.Vars(r)
	title := vars["title"]
	project := getProjectOR404(title, app.DB, w)
	if project == nil {
		return
	}
	tasks := Task{ProjectID: project.ID}

	taskRequest := json.NewDecoder(r.Body)
	err := taskRequest.Decode(&tasks)
	if err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		app.Logger.Println(err)
		fmt.Print(&buf)
		return
	}
	defer r.Body.Close()
	if err := app.DB.Save(&tasks).Error; err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		app.Logger.Println(err)
		fmt.Print(&buf)
		return
	}
	respondJson(w, http.StatusOK, tasks)

}

func (app *App) updateTask(w http.ResponseWriter, r *http.Request) {
	app.Logger.Println("FunctionName: updateTask")
	fmt.Print(&buf)
	vars := mux.Vars(r)
	title := vars["title"]
	project := getProjectOR404(title, app.DB, w)
	if project == nil {
		return
	}
	task_id, _ := strconv.Atoi(vars["id"])
	task := getTaskOR404(task_id, app.DB, w)
	if task == nil {
		return
	}
	taskRequest := json.NewDecoder(r.Body)
	err := taskRequest.Decode(&task)
	if err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		app.Logger.Println(err)
		fmt.Print(&buf)
		return
	}
	defer r.Body.Close()
	if err := app.DB.Save(&task).Error; err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		app.Logger.Println(err)
		fmt.Print(&buf)
		return
	}
	respondJson(w, http.StatusOK, task)
}

func (app *App) deleteTask(w http.ResponseWriter, r *http.Request) {
	app.Logger.Println("FunctionName: deleteTask")
	fmt.Print(&buf)
	vars := mux.Vars(r)
	title := vars["title"]
	project := getProjectOR404(title, app.DB, w)
	if project == nil {
		return
	}
	task_id, _ := strconv.Atoi(vars["id"])
	task := getTaskOR404(task_id, app.DB, w)
	if task == nil {
		return
	}
	if err := app.DB.Unscoped().Delete(&task).Error; err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		app.Logger.Println(err)
		fmt.Print(&buf)
		return
	}
	tasks := []Task{}
	if err := app.DB.Model(&project).Association("Tasks").Find(&tasks); err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		app.Logger.Println(err)
		fmt.Print(&buf)
		return
	}
	project.Tasks = tasks
	respondJson(w, http.StatusOK, project)
}

func getTaskOR404(id int, db *gorm.DB, w http.ResponseWriter) *Task {
	task := Task{}
	if err := db.First(&task, id).Error; err != nil {
		respondError(w, http.StatusNotFound, err.Error())
		return nil
	}
	return &task
}
