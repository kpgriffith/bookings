package handlers

import (
	"net/http"

	"github.com/kpgriffith/bookings/pkg/config"
	"github.com/kpgriffith/bookings/pkg/models"
	"github.com/kpgriffith/bookings/pkg/render"
)

// Repo the repository used by the handlers
var Repo *Repository

// Repository is the repository type
type Repository struct {
	App *config.AppConfig
}

// NewRepo creates the new Repository
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

// NewHandlers sets the Repository for the handlers
func NewHandlers(r *Repository) {
	Repo = r
}

// Home is the handler for the home page
func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	remoteIp := r.RemoteAddr
	m.App.Session.Put(r.Context(), "remote_ip", remoteIp)
	render.RenderTemplate(w, "home.page.tmpl", &models.TemplateData{})
}

// About is the handler for the about page
func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	sm := make(map[string]string)
	sm["test"] = "Hello, again."

	remoteIp := m.App.Session.GetString(r.Context(), "remote_ip")
	sm["remoteip"] = remoteIp

	render.RenderTemplate(w, "about.page.tmpl", &models.TemplateData{
		StringMap: sm,
	})
}
