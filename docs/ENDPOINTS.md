# 📖 API Endpoints

---

## 🟢 Ping
**GET `/ping`**  
Returns a test message.

```json
{
  "message": "initial tests"
}
```

---

## 📦 Products

### 🔹 List All Products
**GET `/products`**  
Returns all registered products.

```json
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
```

---

### 🔹 Create a Product
**POST `/product`**  

**Body**
```json
{
  "name": "Sample Product",
  "price": 100.00
}
```

**Response**
```json
{
  "id": 1,
  "name": "Sample Product",
  "price": 100.00
}
```

---

### 🔹 Get Product by ID
**GET `/product/:productId`**  
Returns a specific product by its ID.

```json
{
  "id": 1,
  "name": "Sample Product",
  "price": 100.00
}
```

---

### 🔹 Update Product
**PUT `/product`**  

**Body**
```json
{
  "id": 1,
  "name": "Updated Product",
  "price": 150.00
}
```

**Response**
```json
{
  "id": 1,
  "name": "Updated Product",
  "price": 150.00
}
```

---

### 🔹 Delete Product
**DELETE `/product/:productId`**  
Removes a product by its ID.

```json
{
  "message": "Product successfully removed"
}
```
