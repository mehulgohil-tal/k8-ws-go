**K8 Coffee Shop** :

---

## **K8 Coffee Shop Overview**  
**Goal:**  
This application simulates the process of ordering coffee. The **Order Service** handles customer orders, while the **Inventory Service** manages stock availability. When a customer places an order, the Order Service checks with the Inventory Service to confirm if the requested coffee ingredients are in stock.


---

### **Architecture**  
1. **Service 1: Order Service**  
   - **Purpose:** Manages customer orders.  
   - **Responsibilities:**
     - Accepts new coffee orders via API.  
     - Communicates with the Inventory Service to verify availability.
     - If ingredients are available, confirms the order.
     - If ingredients are out of stock, suggests an alternative coffee or informs the user.

   **Sample Endpoints:**
   - `GET /order` – Place a new order.
   

   **Example Request:**
   ```json
   GET /order/place?coffeeType=cappuccino&quantity=1
   ```

   **Example Response:**
   ```json
   {
    "coffeeType": "cappuccino",
    "quantity": 1,
    "status": "Confirmed"
  }
   ```

---

2. **Service 2: Inventory Service**  
   - **Purpose:** Manages stock for ingredients (e.g., coffee beans, milk, syrups).  
   - **Responsibilities:**
     - Checks if the required ingredients are available for a given coffee type.
     - Updates stock levels after a successful order.
     - Provides available coffee alternatives when stock is insufficient.

   **Sample Endpoints:**
   - `GET /inventory/stock` – Returns the available stock of all the ingredients.
   - `POST /inventory/used` – Reduce the stock of the specify ingredients.


   **Example Response (GET /inventory/stock):**
   ```json
   {
    "espressoShot": {
      "name": "Espresso Shot",
      "quantity": 10
    },
    "milk": {
      "name": "Milk",
      "quantity": 1000
    },
    "hotWater": {
      "name": "Hot Water",
      "quantity": 99999999
    },
    "milkFoam": {
      "name": "Milk Foam",
      "quantity": 500
    }
  }
   ```

   **Example Request (POST /inventory/used):**
   ```json
   {
     "espressoShot": 1,
     "milk": 200,
     "milkFoam" : 50
   }
   ```

   **Example Response:**
   ```json
   true/false
   ```

---

### **Communication Flow**  
1. **Customer places an order** by sending a request to the **Order Service**.
2. **Order Service** calls **Inventory Service** to **check stock availability**.
3. If the stock is **available**, Order Service confirms the order.
4. If the stock is **unavailable**, Order Service responds with suggested alternatives or an out-of-stock message.

---

### **Workflow Example**  
1. **Customer:**  
   - “I want a Large Cappuccino.”

2. **Order Service:**  
   - Checks with Inventory Service:  
     - “Do we have ingredients for Large Cappuccino?”

3. **Inventory Service:**  
   - Responds:  
     - “Milk is out of stock. Suggested alternatives: Espresso, Americano.”

4. **Order Service:**  
   - Sends a response to the customer:  
     - “Cappuccino is unavailable. Would you like an Espresso or Americano instead?”

---


### **Technology Stack**
- **Backend:**  
  - Java (Spring Boot).
- **API Communication:**  
  - REST or gRPC for service-to-service communication.
