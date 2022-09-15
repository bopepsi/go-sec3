package render

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"path/filepath"
	"text/template"
)

// serve html files
func RenderTemplateBasic(w http.ResponseWriter, tmpl string) {
	parsed, _ := template.ParseFiles("templates/"+tmpl, "templates/base.layout.html")
	err := parsed.Execute(w, nil)
	if err != nil {
		fmt.Println("Error:", err)
	}
}

// Use a cache to store templates used on server
var cache map[string]*template.Template = map[string]*template.Template{}

func RenderTemplateBasicTwo(w http.ResponseWriter, t string) {

	var tmpl *template.Template
	var err error

	_, ok := cache[t]

	if !ok {
		fmt.Println("Adding to cache")
		err = createAndAnddToCache(t)
	} else {
		fmt.Println("Reading from cache")
	}

	tmpl = cache[t]

	err = tmpl.Execute(w, nil)

	if err != nil {
		fmt.Println(err)
	}

}

func createAndAnddToCache(t string) error {

	templateFiles := []string{fmt.Sprintf("templates/%v", t), "templates/base.layout.html"}

	//ParseFiles takes arr of paths
	tmpl, err := template.ParseFiles(templateFiles...)

	if err != nil {
		return err
	}

	cache[t] = tmpl
	return nil
}

// Best way to implement render
func RenderTemplate(w http.ResponseWriter, page string) {
	// create a tmplate cache
	cache, err := createTemplateCache()
	if err != nil {
		log.Fatal(err)
	}

	// for k,v := range cache {
	// 	fmt.Printf("key %s, value %v", k,v)
	// }

	// get requested template from cache
	tmpl, ok := cache[page]
	if !ok {
		okErr := errors.New("no views found")
		log.Fatal(okErr)
	}

	// render template
	tmpl.Execute(w, nil)
}

func createTemplateCache() (map[string]*template.Template, error) {
	var templatesCache map[string]*template.Template = map[string]*template.Template{}

	pages, err := filepath.Glob("templates/*.page.html")
	if err != nil {
		return nil, err
	}

	// range through all files ending with *.page.html
	for _, page := range pages {
		name := filepath.Base(page)
		parsedTmplate, err := template.ParseFiles(page)
		
		if err != nil {
			return nil, err
		}

		layoutPages, err := filepath.Glob("templates/*.layout.html")
		if err != nil {
			return nil, err
		}
		if len(layoutPages) > 0 {
			_, err := parsedTmplate.ParseGlob("templates/*.layout.html")
			if err != nil {
				return nil, err
			}
		}
		fmt.Println(parsedTmplate)
		templatesCache[name] = parsedTmplate

	}

	return templatesCache, nil

}
