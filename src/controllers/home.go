package controllers

import (
	"controllers/util"
	"net/http"
	"text/template"
	"viewmodels"
)

type homeController struct {
	template *template.Template
}

func (this *homeController) get(w http.ResponseWriter, req *http.Request) {
	vm := viewmodels.GetHome()

	w.Header().Add("Content-Type", "text/html")
	responseWriter := util.GetReponseWriter(w, req)
	defer responseWriter.Close()

	this.template.Execute(responseWriter, vm)
}
