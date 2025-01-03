package handlers

import (
	"net/http"

	"github.com/chisty/microservice_go/data"
)

// Delete handles DELETE requests and removes items from the database
func (p *Products) Delete(rw http.ResponseWriter, r *http.Request) {
	id := getProductId(r)
	p.l.Println("[DEBUG] deleting record id", id)

	err := data.DeleteProduct(id)
	if err == data.ErrProductNotFound {
		p.l.Println("[ERROR] deleting record id does not exist")
		rw.WriteHeader(http.StatusNotFound)
		data.ToJSON(data.NewGenericError(err.Error()), rw)
		return
	}

	if err != nil {
		p.l.Println("[ERROR] deleting record", err)
		rw.WriteHeader(http.StatusInternalServerError)
		data.ToJSON(data.NewGenericError(err.Error()), rw)
		return
	}

	rw.WriteHeader(http.StatusNoContent)
}
