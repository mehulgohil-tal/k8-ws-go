package service

import (
    "bytes"
    "encoding/json"
    "fmt"
    "net/http"
    "os"
    "strconv"
    "sync"
    "order-service/model"
	"log"
)

type OrderService struct {
    mu          sync.Mutex
    coffees     map[string]map[string]int
    inventoryURL string
}

func NewOrderService() *OrderService {
    return &OrderService{
        coffees: map[string]map[string]int{
            "cappuccino": {
                "espressoShot": 1,
                "milk":         200,
                "milkFoam":     50,
            },
            "americano": {
                "espressoShot": 1,
                "hotWater":     150,
            },
        },
        inventoryURL: os.Getenv("INVENTORY_URL"),
    }
}

func (s *OrderService) PlaceOrder(w http.ResponseWriter, r *http.Request) {
    coffeeType := r.URL.Query().Get("coffeeType")
    if coffeeType == "" {
        coffeeType = "cappuccino"
    }
    quantityStr := r.URL.Query().Get("quantity")
    if quantityStr == "" {
        quantityStr = "1"
    }
    quantity, err := strconv.Atoi(quantityStr)
    if err != nil {
        http.Error(w, "Invalid quantity", http.StatusBadRequest)
        return
    }

    ingredients := s.getIngredients(coffeeType, quantity)
    available, err := s.checkInventory(ingredients)
    if err != nil {
        http.Error(w, "Failed to check inventory", http.StatusInternalServerError)
        return
    }

    order := model.Order{
        CoffeeType: coffeeType,
        Quantity:   quantity,
        Status:     "Out of Stock",
    }
    if available {
        order.Status = "Confirmed"
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(order)
}

func (s *OrderService) getIngredients(coffeeType string, quantity int) map[string]int {
    s.mu.Lock()
    defer s.mu.Unlock()

    ingredientsOrg := s.coffees[coffeeType]
    ingredients := make(map[string]int)
    for key, value := range ingredientsOrg {
        ingredients[key] = value * quantity
    }
    return ingredients
}

func (s *OrderService) checkInventory(ingredients map[string]int) (bool, error) {
    body, err := json.Marshal(ingredients)
    if err != nil {
        return false, err
    }

	log.Printf("Posting to inventory URL: %s", s.inventoryURL)
    resp, err := http.Post(fmt.Sprintf("%s/inventory/used", s.inventoryURL), "application/json", bytes.NewBuffer(body))
    if err != nil {
        return false, err
    }
    defer resp.Body.Close()

    var available bool
    if err := json.NewDecoder(resp.Body).Decode(&available); err != nil {
        return false, err
    }
    return available, nil
}