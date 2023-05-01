package render

import (
	"bedandbreakfast/pkg/config"
	"bedandbreakfast/pkg/models"
	"bytes"
	"log"
	"net/http"
	"path/filepath"
	"text/template"
)



var app *config.AppConfig

// NewTemplates sets the config for new templatge cache
func NewTemplates(a *config.AppConfig) {
	app = a
}

func AddDefaultData(td *models.TemplateData) *models.TemplateData {
	return td
}

func RenderTemplate(w http.ResponseWriter, t string, td *models.TemplateData) {
	var tc map[string]*template.Template
	if app.UseCache {
		tc = app.TemplateCache
	} else {
		tc, _ = CreateTemplateCache()
	}
	
	tmpl, ok := tc[t]

	if !ok {
		log.Fatal("error parsing template, could not get template from template cache")
	}

	buf := new(bytes.Buffer)

	td = AddDefaultData(td)
	err := tmpl.Execute(buf, td)

	if err != nil {
		log.Println("error parsing template " + err.Error())
	}

	_, err = buf.WriteTo(w)
	if err != nil {
		log.Println(err)
	}
}

func CreateTemplateCache() (map[string] *template.Template, error) {
	myCache := map[string]*template.Template{}

	pages, err := filepath.Glob("./templates/*.page.tmpl")
	if err != nil {
		return myCache, err
	}

	for _, page := range pages {
		name := filepath.Base(page)
		ts, err := template.New(name).ParseFiles(page)
		if err != nil {
			return myCache, err
		}
		matches, err := filepath.Glob("./templates/*.layout.tmpl")
		if err != nil {
			return myCache, err
		}

		if len(matches) > 0 {
			ts, err := ts.ParseGlob("./templates/*.layout.tmpl")
			if err != nil {
				return myCache, err
			}	
			myCache[name] = ts
		}
	}

	return myCache, nil
}
