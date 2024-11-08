package controller

import (
    "net/http"
    "inventory-service/service"
)

var inventoryService = service.NewInventoryService()

func GetStock(w http.ResponseWriter, r *http.Request) {
    inventoryService.GetStock(w, r)
}

func UseIngredient(w http.ResponseWriter, r *http.Request) {
    inventoryService.UseIngredient(w, r)
}