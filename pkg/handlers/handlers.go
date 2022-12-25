package handlers

import (
	"net/http"

	"github.com/Hell0W0rID/Booking_Web_App/pkg/handlers/config"
	render "github.com/Hell0W0rID/Booking_Web_App/pkg/handlers/renders"
	"github.com/Hell0W0rID/Booking_Web_App/pkg/models"
)

// repo the repository used by the handlers
var Repo *Repository

// repository type
type Repository struct {
	App *config.AppConfig
}

// NewRepo creates a new repository
func NewRepo(ptr *config.AppConfig) *Repository {
	return &Repository{
		App: ptr,
	}
}

// NewHandlers sets repository for the handlers
func NewHandlers(r *Repository) {
	Repo = r
}

// Home is the home page handler
func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "home.page.tmpl", &models.TemplateData{})
}

// About is the about page Handler
func (m *Repository) About(w http.ResponseWriter, r *http.Request) {

	stringMap := make(map[string]string)
	stringMap["test"] = "Hello World"
	render.RenderTemplate(w, "about.page.tmpl", &models.TemplateData{
		StringMap: stringMap,
	})
}
