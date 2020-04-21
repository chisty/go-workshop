package handlers

import (
	"context"
	"net/http"

	"github.com/chisty/microservice_go/data"
)

// MiddlewareValidateProduct validates the product in the request and calls next if ok
func (p *Products) MiddlewareValidateProduct(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		prod := data.NewProduct()

		err := data.FromJSON(prod, r.Body)
		if err != nil {
			p.l.Println("ERROR: Unable to unmarshal JSON", err)
			rw.WriteHeader(http.StatusBadRequest)
			// http.Error(rw, "Unable to unmarshal JSON", http.StatusBadRequest)
			return
		}

		errs := p.v.Validate(prod)
		if len(errs) != 0 {
			p.l.Println("ERROR: on validating product", err)
			rw.WriteHeader(http.StatusUnprocessableEntity)
			// http.Error(rw, "Product validation error", http.StatusBadRequest)
			return
		}

		ctx := context.WithValue(r.Context(), KeyProduct{}, *prod)
		req := r.WithContext(ctx)

		next.ServeHTTP(rw, req)
	})
}
