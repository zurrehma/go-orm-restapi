package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/zurrehma/go-orm-restapi/app"
	"github.com/zurrehma/go-orm-restapi/app/model"
	"gorm.io/gorm"
)

type ProjectHandler struct {
}

func (*ProjectHandler) GetAllObjects(app *app.App, w http.ResponseWriter, r *http.Request) {
	app.Logger.Println("FunctionName: getAllProjects")
	fmt.Print(&app.Buf)
	projects := []model.Project{}
	if err := app.DB.Find(&projects).Error; err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		app.Logger.Println(err)
		fmt.Print(&app.Buf)
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

func (*ProjectHandler) CreateObject(app *app.App, w http.ResponseWriter, r *http.Request) {
	app.Logger.Println("FunctionName: createProject")
	fmt.Print(&app.Buf)
	project := model.Project{}
	projectRequest := json.NewDecoder(r.Body)
	err := projectRequest.Decode(&project)
	if err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		app.Logger.Println(err)
		fmt.Print(&app.Buf)
		return
	}
	defer r.Body.Close()
	project.CreatedAt = time.Now()
	if err := app.DB.Save(&project).Error; err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		app.Logger.Println(err)
		fmt.Print(&app.Buf)
		return
	}
	respondJson(w, http.StatusCreated, project)
}

func (*ProjectHandler) GetObject(app *app.App, w http.ResponseWriter, r *http.Request) {
	app.Logger.Println("FunctionName: getProject")
	fmt.Print(&app.Buf)
	vars := mux.Vars(r)
	title := vars["title"]
	project := getProjectOR404(title, app.DB, w)
	if project == nil {
		return
	}
	respondJson(w, http.StatusOK, project)
}

func (*ProjectHandler) UpdateObject(app *app.App, w http.ResponseWriter, r *http.Request) {
	app.Logger.Println("FunctionName: updateProject")
	fmt.Print(&app.Buf)
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
		fmt.Print(&app.Buf)
		return
	}
	defer r.Body.Close()
	project.UpdatedAt = time.Now()
	if err := app.DB.Save(&project).Error; err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		app.Logger.Println(err)
		fmt.Print(&app.Buf)
		return
	}
	respondJson(w, http.StatusOK, project)
}

func (*ProjectHandler) DeleteObject(app *app.App, w http.ResponseWriter, r *http.Request) {
	app.Logger.Println("FunctionName: deleteProject")
	fmt.Print(&app.Buf)
	vars := mux.Vars(r)
	title := vars["title"]
	project := getProjectOR404(title, app.DB, w)
	if project == nil {
		return
	}
	if err := app.DB.Unscoped().Delete(project).Error; err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		app.Logger.Println(err)
		fmt.Print(&app.Buf)
		return
	}
	respondJson(w, http.StatusNoContent, project)
}

func getProjectOR404(title string, db *gorm.DB, w http.ResponseWriter) *model.Project {
	project := model.Project{}
	if err := db.First(&project, model.Project{Title: title}).Error; err != nil {
		respondError(w, http.StatusNotFound, err.Error())
		return nil
	}
	return &project
}
