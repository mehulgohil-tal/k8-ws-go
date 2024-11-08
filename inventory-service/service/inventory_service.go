package service

import (
    "encoding/json"
    "net/http"
    "sync"
    "inventory-service/model"
    "os"
    "strconv"
)

type InventoryService struct {
    mu    sync.Mutex
    stock map[string]*model.Ingredient
}

func NewInventoryService() *InventoryService {
	espressoShotQuantity := getEnvAsInt("ESPRESSO_SHOT_QUANTITY", 10)
    milkQuantity := getEnvAsInt("MILK_QUANTITY", 1000)
    milkFoamQuantity := getEnvAsInt("MILK_FOAM_QUANTITY", 500)
    hotWaterQuantity := getEnvAsInt("HOT_WATER_QUANTITY", 99999999)

    return &InventoryService{
        stock: map[string]*model.Ingredient{
            "espressoShot": {Name: "Espresso Shot", Quantity: espressoShotQuantity},
            "milk":         {Name: "Milk", Quantity: milkQuantity},
            "milkFoam":     {Name: "Milk Foam", Quantity: milkFoamQuantity},
            "hotWater":     {Name: "Hot Water", Quantity: hotWaterQuantity},
        },
    }
}

func getEnv(key string, defaultValue string) string {
    if value, exists := os.LookupEnv(key); exists {
        return value
    }
    return defaultValue
}

func getEnvAsInt(name string, defaultValue int) int {
    valueStr := getEnv(name, "")
    if value, err := strconv.Atoi(valueStr); err == nil {
        return value
    }
    return defaultValue
}

func (s *InventoryService) GetStock(w http.ResponseWriter, r *http.Request) {
    s.mu.Lock()
    defer s.mu.Unlock()
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(s.stock)
}

func (s *InventoryService) UseIngredient(w http.ResponseWriter, r *http.Request) {
    var ingredients map[string]int
    if err := json.NewDecoder(r.Body).Decode(&ingredients); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    s.mu.Lock()
    defer s.mu.Unlock()

    available := true
    for ingredient, quantity := range ingredients {
        if item, exists := s.stock[ingredient]; !exists || item.Quantity < quantity {
            available = false
            break
        }
    }

    if available {
        for ingredient, quantity := range ingredients {
            s.stock[ingredient].Quantity -= quantity
        }
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(available)
}