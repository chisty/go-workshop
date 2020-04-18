package handlers

import (
	"net/http"

	"github.com/chisty/microservice_go/data"
)

// Update handles PUT requests to update products
func (p *Products) Update(rw http.ResponseWriter, r *http.Request) {
	id := getProductId(r)
	prod := r.Context().Value(KeyProduct{}).(data.Product)

	err := data.UpdateProduct(id, &prod)
	if err == data.ErrProductNotFound {
		p.l.Println("[ERROR] product not found", err)
		rw.WriteHeader(http.StatusNotFound)
		data.ToJSON(data.NewGenericError(err.Error()), rw)
		return
	}

	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		data.ToJSON(data.NewGenericError(err.Error()), rw)
		return
	}

	rw.WriteHeader(http.StatusNoContent)
}
