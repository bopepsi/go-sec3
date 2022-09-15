package handlers

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/bopepsi/go-app/pkg/config"
	"github.com/bopepsi/go-app/pkg/models"
	"github.com/bopepsi/go-app/pkg/render"
)

var Repo *Repository

type Repository struct {
	App *config.AppConfig
}

func SetupRepo(a *config.AppConfig) {
	Repo = &Repository{
		App: a,
	}
}

func (this *Repository) Home(w http.ResponseWriter, r *http.Request) {
	strMap := map[string]string{
		"test": "Hello from home page",
	}
	render.RenderTemplate(w, "home.page.html", &models.TemplateData{
		StringMap: strMap,
	})
}

func (this *Repository) About(w http.ResponseWriter, r *http.Request) {
	parsed, _ := template.ParseFiles("templates/about.page.html", "templates/base.layout.html")
	err := parsed.Execute(w, nil)
	if err != nil {
		fmt.Println("Error:", err)
	}
}
