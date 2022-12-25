package render

import (
	"bytes"
	"log"
	"net/http"
	"path/filepath"
	"text/template"

	"github.com/Hell0W0rID/Booking_App/pkg/handlers/config"
	"github.com/Hell0W0rID/Booking_App/pkg/models"
)

// 1 . RenderTemplates renders templates
// func RenderTemplateTest(w http.ResponseWriter, tmpl string) {

// 	parsedTemplate, _ := template.ParseFiles("./templates/" + tmpl)
// 	err := parsedTemplate.Execute(w, nil)
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// }

// 2 . Render templates with simple cache
// var tempCache = make(map[string]*template.Template)
// func RenderTemplate(w http.ResponseWriter , t string)
// {
// 	var tmpl *template.Template
// 	var err error

// 	//check to see if we already have the template in our cache
// 	_ , inMap := tc[t]

// 	if !inMap {
// 		// need to create template
// 		err = createTemplateCache(t)
// 		if err != nil {
// 			log.Println("error in creating template")
// 		}
// 	} else {
// 		//we have template in cache
// 		log.Println("using cached template")
// 	}

// 	tmpl = tc[t]
// 	err = tmpl.Execute(w,nil)
// }
// func createTemplateCache(t string ) error {
// 	templates := []string{
// 		fmt.Sprintf("./templates/%s" + t),
// 		"./templates/base.layout.tmpl",
// 	}

// 	//parse template
// 	tmpl , err := template.ParseFiles(templates...)
// 	if err != nil {
// 		return err
// 	}
// 	tc[t] = tmpl
// }

var app *config.AppConfig

func NewTemplates(ptr *config.AppConfig) {
	app = ptr
}

// 3. render template with complex cache

func RenderTemplate(w http.ResponseWriter, tmpl string, templData *models.TemplateData) {
	var tc map[string]*template.Template
	var err error
	if app.UseCache {
		// get the template cache from app config
		tc = app.TemplateCache
	} else {
		tc, err = CreateTemplateCache()
		if err != nil {
			log.Println(err)
		}
	}
	// // create template cache
	// tc, err := CreateTemplateCache()
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// get requested template from cache
	temp, found := tc[tmpl]
	if !found {
		log.Fatal("could not get template from template cache")
	}

	buf := new(bytes.Buffer)

	err = temp.Execute(buf, templData)
	if err != nil {
		log.Println(err)
	}
	// render the template
	_, err = buf.WriteTo(w)
	if err != nil {
		log.Println(err)
	}
}

func CreateTemplateCache() (map[string]*template.Template, error) {

	myCache := map[string]*template.Template{} // create empty map

	//get all of the files path named *.page.tmpl from ./templates
	pages, err := filepath.Glob("./templates/*.page.tmpl")
	if err != nil {
		return myCache, err
	}
	// range through all the files ending with *.page.tmpl
	for _, page := range pages {
		//provide the last element from path
		name := filepath.Base(page)
		ts, err := template.New(name).ParseFiles(page)
		if err != nil {
			return myCache, err

		}
		layouts, err := filepath.Glob("./templates/*.layout.tmpl")
		if err != nil {
			return myCache, err
		}
		if len(layouts) > 0 {
			ts, err = ts.ParseGlob("./templates/*.layout.tmpl")
			if err != nil {
				return myCache, err
			}
		}
		myCache[name] = ts
	}
	return myCache, nil

}
