package handler

import (
	"Order/services/common/genproto/orders"
	"Order/services/orders/types"
	"encoding/json"
	"net/http"
)

type OrdersHttpHandler struct {
	ordersService types.OrderService
}

func NewHttpOrdersService(ordersService types.OrderService) *OrdersHttpHandler {
	return &OrdersHttpHandler{
		ordersService: ordersService,
	}
}

func (h *OrdersHttpHandler) RegisterRouter(router *http.ServeMux) {
	router.HandleFunc("POST /orders", h.CreateOrder)
}

func (h *OrdersHttpHandler) CreateOrder(w http.ResponseWriter, r *http.Request) {

	var req orders.CreateOrderRequest

	err := json.NewDecoder(r.Body).Decode(&req)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	order := &orders.Order{
		OrderID:    42,
		CustomerID: req.CustomerID,
		ProductID:  req.ProductID,
		Quantity:   req.Quantity,
	}

	err = h.ordersService.CreateOrder(r.Context(), order)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	res := &orders.CreateOrderResponse{
		Status: "success",
	}

	data, err := json.Marshal(res)

	if err != nil {
		http.Error(w, err.Error(), http.StatusUnprocessableEntity)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(data)

}
