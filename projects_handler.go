package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func (app *App) getAllProjects(w http.ResponseWriter, r *http.Request) {
	app.Logger.Println("FunctionName: getAllProjects")
	fmt.Print(&buf)
	projects := []Project{}
	if err := app.DB.Find(&projects).Error; err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		app.Logger.Println(err)
		fmt.Print(&buf)
		return
	}
	// tasks := []Task{}
	// if err := app.DB.Model(&projects).Association("Tasks").Find(&tasks); err != nil {
	// 	respondError(w, http.StatusInternalServerError, err.Error())
	// 	app.Logger.Println(err)
	// 	fmt.Print(&buf)
	// 	return
	// }
	respondJson(w, http.StatusOK, projects)
}

func (app *App) createProject(w http.ResponseWriter, r *http.Request) {
	app.Logger.Println("FunctionName: createProject")
	fmt.Print(&buf)
	project := Project{}
	projectRequest := json.NewDecoder(r.Body)
	err := projectRequest.Decode(&project)
	if err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		app.Logger.Println(err)
		fmt.Print(&buf)
		return
	}
	defer r.Body.Close()
	project.CreatedAt = time.Now()
	if err := app.DB.Save(&project).Error; err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		app.Logger.Println(err)
		fmt.Print(&buf)
		return
	}
	respondJson(w, http.StatusCreated, project)
}

func (app *App) getProject(w http.ResponseWriter, r *http.Request) {
	app.Logger.Println("FunctionName: getProject")
	fmt.Print(&buf)
	vars := mux.Vars(r)
	title := vars["title"]
	project := getProjectOR404(title, app.DB, w)
	if project == nil {
		return
	}
	respondJson(w, http.StatusOK, project)
}

func (app *App) updateProject(w http.ResponseWriter, r *http.Request) {
	app.Logger.Println("FunctionName: updateProject")
	fmt.Print(&buf)
	vars := mux.Vars(r)
	title := vars["title"]
	project := getProjectOR404(title, app.DB, w)
	if project == nil {
		return
	}

	projectRequest := json.NewDecoder(r.Body)
	err := projectRequest.Decode(&project)
	if err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		app.Logger.Println(err)
		fmt.Print(&buf)
		return
	}
	defer r.Body.Close()
	project.UpdatedAt = time.Now()
	if err := app.DB.Save(&project).Error; err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		app.Logger.Println(err)
		fmt.Print(&buf)
		return
	}
	respondJson(w, http.StatusOK, project)
}

func (app *App) deleteProject(w http.ResponseWriter, r *http.Request) {
	app.Logger.Println("FunctionName: deleteProject")
	fmt.Print(&buf)
	vars := mux.Vars(r)
	title := vars["title"]
	project := getProjectOR404(title, app.DB, w)
	if project == nil {
		return
	}
	if err := app.DB.Unscoped().Delete(project).Error; err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		app.Logger.Println(err)
		fmt.Print(&buf)
		return
	}
	respondJson(w, http.StatusNoContent, project)
}

func getProjectOR404(title string, db *gorm.DB, w http.ResponseWriter) *Project {
	project := Project{}
	if err := db.First(&project, Project{Title: title}).Error; err != nil {
		respondError(w, http.StatusNotFound, err.Error())
		return nil
	}
	return &project
}
