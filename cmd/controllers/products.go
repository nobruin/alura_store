package controllers

import (
	"alura_store/cmd/models"
	"log"
	"net/http"
	"strconv"
	"text/template"
)

const (
	TEMPLATES_PATH = "./templates/*html"
)

var templates = template.Must(template.ParseGlob(TEMPLATES_PATH))

func Index(w http.ResponseWriter, r *http.Request) {
	products := models.FindAll()
	templates.ExecuteTemplate(w, "Index", products)
}

func New(w http.ResponseWriter, r *http.Request) {
	templates.ExecuteTemplate(w, "New", nil)
}

func Insert(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		product := bindRequestToProduct(r)
		models.CreateProduct(product)
	}

	http.Redirect(w, r, "/", http.StatusMovedPermanently)
}

func Delete(w http.ResponseWriter, r *http.Request) {
	id := bindIdByQueryParam(r)
	models.Delete(id)
	http.Redirect(w, r, "/", http.StatusMovedPermanently)
}

func Update(w http.ResponseWriter, r *http.Request) {

	if r.Method == "POST" {
		id := bindIdByQueryParam(r)
		product := bindRequestToProduct(r)
		product.ID = id
		models.Update(product)
	}

	http.Redirect(w, r, "/", http.StatusMovedPermanently)
}

func Edit(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	product := models.FindById(id)
	templates.ExecuteTemplate(w, "Edit", product)
}

func bindRequestToProduct(r *http.Request) models.Product {
	price, err := strconv.ParseFloat(r.FormValue("price"), 64)
	if err != nil {
		log.Println("Bind float parse error", err)
	}
	quantity, err := strconv.Atoi(r.FormValue("quantity"))
	if err != nil {
		log.Println("Bind Int parse error", err)

	}

	return models.NewModelProduct(r.FormValue("name"), price, quantity, r.FormValue("description"))
}

func bindIdByQueryParam(r *http.Request) int {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))

	if err != nil {
		log.Println("Bind Int parse error", err)
	}

	return id
}
