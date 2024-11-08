package controller

import (
    "net/http"
    "order-service/service"
)

var orderService = service.NewOrderService()

func PlaceOrder(w http.ResponseWriter, r *http.Request) {
    orderService.PlaceOrder(w, r)
}