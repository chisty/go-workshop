package handlers

import (
	"net/http"

	"github.com/chisty/microservice_go/data"
)

// Create handles post request
func (p *Products) Create(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle Post Request")

	prod := r.Context().Value(KeyProduct{}).(data.Product)
	p.l.Println("Running prod: ", prod)
	data.AddProduct(prod)
}
