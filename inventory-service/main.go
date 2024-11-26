package main

import (
    "net/http"
    "inventory-service/controller"
	"os"
)

func main() {
    http.HandleFunc("/inventory/stock", controller.GetStock)
    http.HandleFunc("/inventory/used", controller.UseIngredient)
	port := getEnv("PORT", "3000")

    http.ListenAndServe(":"+port, nil)
}

func getEnv(key string, defaultValue string) string {
    if value, exists := os.LookupEnv(key); exists {
        return value
    }
    return defaultValue
}