package controllers

import (
	"bufio"
	"net/http"
	"os"
	"strings"
	"text/template"
)

func Register(templates *template.Template) {
	/*http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		requestedFile := req.URL.Path[1:]
		template := templates.Lookup(requestedFile + ".html")

		var context interface{} = nil
		switch requestedFile {
		case "home":
			context = viewmodels.GetHome()
		case "categories":
			context = viewmodels.GetCategories()
		}

		if template != nil {
			template.Execute(w, context)
		} else {
			w.WriteHeader(404)
		}
	})*/

	hc := new(homeController)
	hc.template = templates.Lookup("home.html")
	http.HandleFunc("/home", hc.get)

	cc := new(categoriesController)
	cc.template = templates.Lookup("categories.html")
	http.HandleFunc("/categories", cc.get)

	http.HandleFunc("/img/", serveResources)
	http.HandleFunc("/css/", serveResources)
}

func serveResources(w http.ResponseWriter, req *http.Request) {
	path := "../../public" + req.URL.Path
	var contentType string
	if strings.HasSuffix(path, ".css") {
		contentType = "text/css"
	} else if strings.HasPrefix(path, ".png") {
		contentType = "image/png"
	} else {
		contentType = "text/plain"
	}

	f, err := os.Open(path)

	if err == nil {
		defer f.Close()

		w.Header().Add("Content-Type", contentType)

		br := bufio.NewReader(f)
		br.WriteTo(w)
	} else {
		w.WriteHeader(404)
	}
}
