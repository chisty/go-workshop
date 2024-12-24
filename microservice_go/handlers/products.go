// Package classification of Product API
//
// Documentation for Product API
//
//	Schemes: http
//	BasePath: /
//	Version: 1.0.0
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
// swagger:meta
package handlers

import (
	"log"
	"net/http"
	"strconv"

	"github.com/chisty/microservice_go/data"
	"github.com/gorilla/mux"
)

// A list of products returns in the response
//swagger:response productsResponse
type productsResponse struct {
	// All products in system
	// in:Body
	Body []data.Product
}

// Products is a http.handler
type Products struct {
	l *log.Logger
	v *data.Validation
}

// NewProducts creates a products handler with the given logger
func NewProducts(l *log.Logger, v *data.Validation) *Products {
	return &Products{l, v}
}

func getProductId(r *http.Request) int {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		panic(err)
	}

	return id
}

// func (p *Products) AddProduct(rw http.ResponseWriter, r *http.Request) {
// 	p.l.Println("Handle Post Request")

// 	prod := r.Context().Value(keyProduct{}).(data.Product)
// 	p.l.Println("Running prod: ", prod)
// 	data.AddProduct(&prod)
// }

// func (p *Products) UpdateProduct(rw http.ResponseWriter, r *http.Request) {
// 	vars := mux.Vars(r)
// 	id, err := strconv.Atoi(vars["id"])
// 	if err != nil {
// 		http.Error(rw, "Invalid id parameter", http.StatusBadRequest)
// 	}

// 	p.l.Println("Handle PUT Request with id= ", id)

// 	prod := r.Context().Value(KeyProduct{}).(data.Product)

// 	err = data.UpdateProduct(id, &prod)
// 	if err == data.ErrProductNotFound {
// 		http.Error(rw, "Product not found", http.StatusNotFound)
// 		return
// 	}

// 	if err != nil {
// 		http.Error(rw, "Internal server error.", http.StatusInternalServerError)
// 		return
// 	}
// }

type KeyProduct struct{}
