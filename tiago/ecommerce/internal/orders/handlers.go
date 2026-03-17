package orders

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

func (this *Handler) PlaceOrder(writer http.ResponseWriter, request *http.Request) {
	var tempOrder CreateOrderParams
	var err = json.ReadRequest(request, &tempOrder)
	if err != nil {
		log.Println(err)
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}

	var createdOrder, errCreate = this.service.PlaceOrder(request.Context(), tempOrder)
	if errCreate != nil {
		log.Println(err)
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}

	json.WriteResponse(writer, http.StatusCreated, createdOrder)
}
