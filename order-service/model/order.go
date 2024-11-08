package model

type Order struct {
    CoffeeType string `json:"coffeeType"`
    Quantity   int    `json:"quantity"`
    Status     string `json:"status"`
}