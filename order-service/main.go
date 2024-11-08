package main

import (
    "net/http"
    "order-service/controller"
    "os"
)

func main() {
    port := getEnv("PORT", "3000")
    http.HandleFunc("/order/place", controller.PlaceOrder)

    http.ListenAndServe(":"+port, nil)
}


func getEnv(key string, defaultValue string) string {
    if value, exists := os.LookupEnv(key); exists {
        return value
    }
    return defaultValue
}