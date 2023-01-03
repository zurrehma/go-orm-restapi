package app

import "net/http"

type Routes interface {
	GetAllObjects(a *App, w http.ResponseWriter, r *http.Request)
	CreateObject(a *App, w http.ResponseWriter, r *http.Request)
	GetObject(a *App, w http.ResponseWriter, r *http.Request)
	UpdateObject(a *App, w http.ResponseWriter, r *http.Request)
	DeleteObject(a *App, w http.ResponseWriter, r *http.Request)
}
