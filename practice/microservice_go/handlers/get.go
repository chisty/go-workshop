package handlers

import (
	"net/http"

	"github.com/chisty/microservice_go/data"
)

// ListAll handles GET requests and returns all current products
func (p *Products) ListAll(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("[DEBUG] get all products")

	prods := data.GetProducts()
	err := data.ToJSON(prods, rw)
	if err != nil {
		//http.Error(rw, "Unable to marshal json.", http.StatusInternalServerError)
		p.l.Println("[ERROR] serializing product", err)

	}
}

// ListSingle handles GET requests
func (p *Products) ListSingle(rw http.ResponseWriter, r *http.Request) {
	id := getProductId(r)
	p.l.Println("[DEBUG] get record id", id)

	prod, err := data.GetProductByID(id)

	switch err {
	case nil:
	case data.ErrProductNotFound:
		p.l.Println("[ERROR] fetching product", err)
		rw.WriteHeader(http.StatusNotFound)
		data.ToJSON(data.NewGenericError(err.Error()), rw)
		return

	default:
		p.l.Println("[ERROR] fetching product", err)

		rw.WriteHeader(http.StatusInternalServerError)
		data.ToJSON(data.NewGenericError(err.Error()), rw)
		return
	}

	err = data.ToJSON(prod, rw)
	if err != nil {
		p.l.Println("[ERROR] serializing product", err)
	}

}
