package handlers

import (
	"bedandbreakfast/pkg/config"
	"bedandbreakfast/pkg/models"
	"bedandbreakfast/pkg/render"
	"net/http"
)

// Repository
type Repository struct {
	App *config.AppConfig
}

// Repo is the repository type
var Repo *Repository


//NewRepo creates a new Repository
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

//NewHandlers sets the repository for the handlers
func NewHandler(r *Repository) {
	Repo = r
}

// Home is the root page handler
func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	remoteIP := r.RemoteAddr
	m.App.Session.Put(r.Context(), "remote_ip", remoteIP)

	render.RenderTemplate(w, "home.page.tmpl", &models.TemplateData{})
}

// About is the about page handler
func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	stringMap := make(map[string]string)
	stringMap["test"] = "Hello , again!"

	stringMap["remoteIP"] = m.App.Session.GetString(r.Context(), "remote_ip")

	
	render.RenderTemplate(w, "about.page.tmpl", &models.TemplateData{StringMap: stringMap})
}