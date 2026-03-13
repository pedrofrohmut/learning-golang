package products

import (
	"ecommerce/internal/json"
	"log"
	"net/http"
)

type Handler struct {
	service Service
}

func NewHandler(service Service) *Handler {
	return &Handler { service: service }
}

func (this *Handler) ListProducts(writer http.ResponseWriter, request *http.Request) {
	var products, err = this.service.ListProducts(request.Context())
	if err != nil {
		log.Println(err)
		http.Error(writer, err.Error(), 500)
	}

	json.WriteResponse(writer, 200, products)
}
