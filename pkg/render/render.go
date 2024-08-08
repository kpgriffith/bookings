package render

import (
	"bytes"
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/justinas/nosurf"
	"github.com/kpgriffith/bookings/pkg/config"
	"github.com/kpgriffith/bookings/pkg/models"
)

var app *config.AppConfig

// NewTemplates sets the config for the render package
func NewTemplates(a *config.AppConfig) {
	app = a
}

func AddDefaultData(td *models.TemplateData, r *http.Request) *models.TemplateData {
	td.CSRFToken = nosurf.Token(r)
	return td
}

// RenderTemplate renders a template
func RenderTemplate(w http.ResponseWriter, r *http.Request, tmpl string, data *models.TemplateData) {
	var tc map[string]*template.Template
	if app.UseCache {
		// get the template cache from the application config
		tc = app.TemplateCache
	} else {
		tc, _ = CreateTemplateCache()
	}

	// get requested template from cache
	t, ok := tc[tmpl]
	if !ok {
		log.Fatal("Could not get template from TemplateCache")
	}

	// execute the template found to make double sure
	// we receive the correct template
	buf := new(bytes.Buffer)
	data = AddDefaultData(data, r)
	_ = t.Execute(buf, data)

	// render the template
	_, err := buf.WriteTo(w)
	if err != nil {
		log.Println("Error writing template to browser", err)
	}
}

func CreateTemplateCache() (map[string]*template.Template, error) {
	myCache := map[string]*template.Template{}

	// get all *.page.tmpl templates first
	pages, err := filepath.Glob("./templates/*.page.tmpl")
	if err != nil {
		return myCache, err
	}

	for _, p := range pages {
		name := filepath.Base(p) // just get the file name without the path
		ts, err := template.New(name).ParseFiles(p)
		if err != nil {
			return myCache, err
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
