package handlers

import (
	"myapp/pkg/config"
	"myapp/pkg/render"
	"net/http"
)

//Repo - the repository used by the handlers
var Repo *Repository

//Repository - is the repository type (struct)
type Repository struct {
	App *config.AppConfig
}

//NewRepo - creates a new repository for the handlers
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

//NewHandlers - sets the repository for the handlers
func NewHandlers(r *Repository) {
	Repo = r
}

func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "home.page.tmpl")
}

func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "about.page.tmpl")

}

