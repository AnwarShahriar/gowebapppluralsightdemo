package controllers

import (
	"controllers/util"
	"net/http"
	"strconv"
	"text/template"
	"viewmodels"

	"converters"
	"models"

	"github.com/gorilla/mux"
)

type productController struct {
	template *template.Template
}

func (this *productController) get(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)

	idRaw := vars["id"]

	id, err := strconv.Atoi(idRaw)
	if err == nil {
		product, err := models.GetProductById(id)

		if err == nil {

			vm := viewmodels.GetProduct(product.Name())
			vm.Product = converters.ConvertProductToViewModel(product)

			w.Header().Add("Content-Type", "text/html")
			responseWriter := util.GetReponseWriter(w, req)
			defer responseWriter.Close()

			this.template.Execute(responseWriter, vm)
		}
	} else {
		w.WriteHeader(404)
	}
}
