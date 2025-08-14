# 📖 API Endpoints

---

## 🟢 Ping
**GET `/ping`**  
Returns a test message.

-------------------------
{
  "message": "initial tests"
}
-------------------------

---

## 📦 Products

### 🔹 List All Products
**GET `/products`**  
Returns all registered products.

-------------------------
[
  {
    "id": 1,
    "name": "Product 1",
    "price": 100.00
  },
  {
    "id": 2,
    "name": "Product 2",
    "price": 150.00
  }
]
-------------------------

---

### 🔹 Create a Product
**POST `/product`**  

**Body**
-------------------------
{
  "name": "Sample Product",
  "price": 100.00
}
-------------------------

**Response**
-------------------------
{
  "id": 1,
  "name": "Sample Product",
  "price": 100.00
}
-------------------------

---

### 🔹 Get Product by ID
**GET `/product/:productId`**  
Returns a specific product by its ID.

-------------------------
{
  "id": 1,
  "name": "Sample Product",
  "price": 100.00
}
-------------------------

---

### 🔹 Update Product
**PUT `/product`**  

**Body**
-------------------------
{
  "id": 1,
  "name": "Updated Product",
  "price": 150.00
}
-------------------------

**Response**
-------------------------
{
  "id": 1,
  "name": "Updated Product",
  "price": 150.00
}
-------------------------

---

### 🔹 Delete Product
**DELETE `/product/:productId`**  
Removes a product by its ID.

-------------------------
{
  "message": "Product successfully removed"
}
-------------------------
