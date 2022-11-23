package render

import (
	"bytes"
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/vikashparashar/Hotel_Bookings_2/pkg/config"
	"github.com/vikashparashar/Hotel_Bookings_2/pkg/models"
)

// stored parsed template into it
// tc is for template cache
var tc = make(map[string]*template.Template)

var functions = template.FuncMap{}

var app *config.AppConfig

// NewTemplates sets the config for the template package
func NewTemplates(a *config.AppConfig) {
	app = a
}

func AddDefaultData(td *models.Template_Data) *models.Template_Data {
	return td
}

func RenderTemplates(w http.ResponseWriter, temp string, td *models.Template_Data) {

	var tc = app.TemplateCache
	if !app.UseCache {
		// get the template cache from the app config
		tc = app.TemplateCache
	} else {
		tc, _ = CreateTemplateCache()
	}

	t, ok := tc[temp]
	if !ok {
		log.Fatalln("could not get template from template cache")
	}

	buf := new(bytes.Buffer)
	td = AddDefaultData(td)
	err := t.Execute(buf, td)

	if err != nil {
		log.Fatalln(err)
	}
	// rander the template

	_, err = buf.WriteTo(w)
	if err != nil {
		log.Println("error writing template to browser")
	}

}

func CreateTemplateCache() (map[string]*template.Template, error) {
	// var myCache = make(map[string]*template.Template)
	myCache := map[string]*template.Template{}

	// get all the files name *.page.tmpl from ./templates
	pages, err := filepath.Glob("./templates/*.page.tmpl")
	if err != nil {
		return myCache, err
	}

	// range through all pages ending with *.page.tmpl

	for _, page := range pages {
		name := filepath.Base(page)
		ts, err := template.New(name).ParseFiles(page)
		if err != nil {
			log.Println("failed to parse file : ", ts)
		}

		matches, err := filepath.Glob("./templates/*.layout.tmpl")
		if err != nil {
			return myCache, err
		}
		if len(matches) > 0 {
			ts, err = ts.ParseGlob("./templates/*.layout.tmpl")
			if err != nil {
				return myCache, err
			}
		}

		myCache[name] = ts
	}

	return myCache, nil
}
