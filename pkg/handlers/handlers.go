package handlers

import (
	"net/http"

	"github.com/vikashparashar/bookings/pkg/config"
	"github.com/vikashparashar/bookings/pkg/models"
	"github.com/vikashparashar/bookings/pkg/render"
)

// Repo is the Repository used by the handlers
var Repo *Repository

// Repository is the Repository type
type Repository struct {
	App *config.AppConfig
}

// NewRepo creates the new Repository
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

// NewHandlers sets the repository for handlers
func NewHandlers(r *Repository) {
	Repo = r
}

func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	remoteIP := r.RemoteAddr
	m.App.Session.Put(r.Context(), "remote_ip", remoteIP)

	// perform some logic
	stringMap := map[string]string{}
	stringMap["Name"] = "Vikash"

	// sending data to the template
	render.RenderTemplates(w, "home.page.tmpl", &models.Template_Data{})
}

func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	remoteIP := m.App.Session.GetString(r.Context(), "remote_ip")
	m.App.Session.Put(r.Context(), "remote_ip", remoteIP)
	stringMap := map[string]string{}
	// stringMap["Name"] = "Vikash"
	stringMap["remote_ip"] = remoteIP

	// sending data to the template
	render.RenderTemplates(w, "about.page.tmpl", &models.Template_Data{StringMap: stringMap})
}

func (m *Repository) General(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplates(w, "general.page.tmpl", &models.Template_Data{})
}
func (m *Repository) Major(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplates(w, "major.page.tmpl", &models.Template_Data{})
}
func (m *Repository) Contact(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplates(w, "contact.page.tmpl", &models.Template_Data{})
}
func (m *Repository) CheckAvailability(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplates(w, "check.page.tmpl", &models.Template_Data{})
}
