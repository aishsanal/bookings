package handlers

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/aishsanal/bookings/pkg/config"
	"github.com/aishsanal/bookings/pkg/models"
	"github.com/justinas/nosurf"
)

func Home(w http.ResponseWriter, r *http.Request) {
	ipAddress := r.RemoteAddr
	appConfig.Session.Put(r.Context(), "ipAddress", ipAddress)
	renderTemplate(w, r, "home.page.tmpl", &models.TemplateData{})
}

func About(w http.ResponseWriter, r *http.Request) {
	stringMap := make(map[string]string)
	stringMap["test"] = "hii"
	stringMap["ipAddress"] = appConfig.Session.GetString(r.Context(), "ipAddress")
	tempData := models.TemplateData{
		StringMap: stringMap,
	}
	renderTemplate(w, r, "about.page.tmpl", &tempData)
}

func Thumpa(w http.ResponseWriter, r *http.Request) {
	ipAddress := r.RemoteAddr
	appConfig.Session.Put(r.Context(), "ipAddress", ipAddress)
	renderTemplate(w, r, "thumpa.tmpl", &models.TemplateData{})
}

func Mulla(w http.ResponseWriter, r *http.Request) {
	ipAddress := r.RemoteAddr
	appConfig.Session.Put(r.Context(), "ipAddress", ipAddress)
	renderTemplate(w, r, "mulla.tmpl", &models.TemplateData{})
}

func Reservation(w http.ResponseWriter, r *http.Request) {
	ipAddress := r.RemoteAddr
	appConfig.Session.Put(r.Context(), "ipAddress", ipAddress)
	renderTemplate(w, r, "make.reservation.tmpl", &models.TemplateData{})
}

func Availability(w http.ResponseWriter, r *http.Request) {
	ipAddress := r.RemoteAddr
	appConfig.Session.Put(r.Context(), "ipAddress", ipAddress)
	renderTemplate(w, r, "availability.tmpl", &models.TemplateData{})
}

func PostAvailability(w http.ResponseWriter, r *http.Request) {
	ipAddress := r.RemoteAddr
	appConfig.Session.Put(r.Context(), "ipAddress", ipAddress)

	start := r.Form.Get("start")
	end := r.Form.Get("end")

	w.Write([]byte(fmt.Sprintf("start date is %s and end date is %s", start, end)))
}

type jsonResponse struct {
	OK      bool   `json:"ok"`
	Message string `json:"message"`
}

func JSONPostAvailability(w http.ResponseWriter, r *http.Request) {
	responseMsg := jsonResponse{
		OK:      true,
		Message: "Rooms are available",
	}

	jsonResp, err := json.MarshalIndent(responseMsg, "", "    ")
	if err != nil {
		log.Println(err)
	}
	w.Header().Set("content-type", "application/json")
	w.Write(jsonResp)
}

func Contact(w http.ResponseWriter, r *http.Request) {
	ipAddress := r.RemoteAddr
	appConfig.Session.Put(r.Context(), "ipAddress", ipAddress)
	renderTemplate(w, r, "contact.tmpl", &models.TemplateData{})
}

func getCSRFToken(r *http.Request) string {
	return nosurf.Token(r)
}

var appConfig config.AppConfig

func renderTemplate(w http.ResponseWriter, r *http.Request, t string, templateData *models.TemplateData) {
	var template *template.Template
	if appConfig.UseCache {
		template = appConfig.TemplateCache[t]
	} else {
		cache, _ := CreateTemplateCache()
		template = cache[t]
	}

	csrfToken := getCSRFToken(r)
	templateData.CSRFToken = csrfToken
	err := template.Execute(w, templateData)
	if err != nil {
		log.Fatal(err)
	}
}

func SetConfig(a config.AppConfig) {
	appConfig = a
}

type Repository struct {
	app config.AppConfig
}

var repository *Repository

func SetRepository(repo *Repository) {
	repository = repo
}

func CreateRepository(cnf config.AppConfig) *Repository {
	return &Repository{
		app: cnf,
	}
}

func CreateTemplateCache() (map[string]*template.Template, error) {
	cache := map[string]*template.Template{}

	pagesPath, err := filepath.Glob("./templates/*.tmpl")
	if err != nil {
		return cache, err
	}

	for _, path := range pagesPath {
		name := filepath.Base(path)
		templ, err := template.New(name).ParseFiles(path)
		if err != nil {
			return cache, err
		}

		layoutPath, err := filepath.Glob("./templates/*.layout.tmpl")
		if err != nil {
			return cache, err
		}

		if len(layoutPath) > 0 {
			templ, err = templ.ParseGlob(layoutPath[0])
			if err != nil {
				return cache, err
			}
		}
		cache[name] = templ
	}
	return cache, nil
}

// The below functions were written first and the above ones are the refined versions
var tempCache = make(map[string]*template.Template)

func renderTemplateOld(w http.ResponseWriter, t string) {
	var tmpl *template.Template
	var err error
	_, inMap := tempCache[t]
	if !inMap {
		log.Println("Loading template from disk")
		tmpl, err = loadTemplateOld(t)
		if err != nil {
			log.Println("Error loading template from disk")
			return
		}
		tempCache[t] = tmpl
	} else {
		log.Println("Loading template from template cache")
		tmpl = tempCache[t]
	}
	err = tmpl.Execute(w, nil)
	if err != nil {
		log.Println("Error rendering template ", t)
	}
}

func loadTemplateOld(t string) (*template.Template, error) {
	templatesToLoad := []string{
		fmt.Sprintf("../../templates/%s", t),
		"../../templates/base.layout.tmpl",
	}
	tmpl, err := template.ParseFiles(templatesToLoad...)
	return tmpl, err
}
