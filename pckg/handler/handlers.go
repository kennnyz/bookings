package Handler

import (
	"github.com/kennnyz/bookings/pckg/config"
	"github.com/kennnyz/bookings/pckg/models"
	"github.com/kennnyz/bookings/pckg/render"
	"net/http"
)

//TemplateDate holds data sent from handlers to templates

//Repo the repository used by handlers
var Repo *Repository

//Repository is the repository type
type Repository struct {
	App *config.AppConfig
}

//NewRepo create new repository
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

//NewHandlers sets the repository for the handlers
func NewHandlers(r *Repository) {
	Repo = r
}

func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	remoteIP := r.RemoteAddr
	m.App.Session.Put(r.Context(), "remote_ip", remoteIP)
	render.RenderTemplate(w, "home.page.html", &models.TemplateDate{})
}

//perform some logic

func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	// perform some logic
	stringMap := make(map[string]string)
	stringMap["test"] = "Hello again"

	remoteIP := m.App.Session.GetString(r.Context(), "remote_ip")
	stringMap["remote_ip"] = remoteIP
	// send the data to the template
	render.RenderTemplate(w, "about.page.html", &models.TemplateDate{
		StringMap: stringMap,
	})
	// perform some logic
}
