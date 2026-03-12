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
	var err = this.service.ListProducts(request.Context())
	if err != nil {
		log.Println(err)
		http.Error(writer, err.Error(), 500)
	}

	var products = struct {
		Product []string `json:"products"`
	} {
	}

	json.WriteResponse(writer, 200, products)
}
