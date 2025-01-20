# API-GO: Product Management with JWT Verification

This project is an API for product management developed in **Go** using the **Gin** framework. The architecture is modular, ensuring better organization and maintainability.

<br><br><br>

---

## 🚀 Project Initialization


### Pre-requisites

Make sure you have installed:
- **Golang**: [Install Golang](https://go.dev/doc/install)
- **Docker**: [Install Docker](https://www.docker.com/products/docker-desktop)

---

### How to Run

1. Clone the project and navigate to the folder:
   ```bash
   git clone https://github.com/usuario/api-go.git
   cd api-go
   ```

2. Configure the database connection:
   Edit the settings in the `db/conn.go` file to reflect your credentials.

3. Start the services with Docker:
   ```bash
   docker-compose up
   ```

4. Access the API at the address:
   - `http://localhost:8000`

---

## 📚 Project Structure

The structure is organized as follows:

```plaintext
API-GO/
├── cmd/
│   └── main.go               # Main file to start the application
├── controller/
│   └── product_controller.go # Controllers responsible for the routes
├── db/
│   └── conn.go               # Database configuration and connection
├── model/
│   ├── product.go            # Product model
│   └── response.go           # Response structures for the API
├── repository/
│   └── product_repository.go # Layer for interacting with the database
├── usecase/
│   └── product_usecase.go    # Business logic and application logic
├── .gitignore                # Files ignored in version control
├── docker-compose.yml        # Docker Compose configuration
├── Dockerfile                # Docker container configuration
├── go.mod                    # Project dependencies
├── go.sum                    # Dependency hashes
└── README.md                 # Project documentation
```

---

## 📖 Endpoints

### **Ping**
- **GET `/ping`**
  - Returns a test message.
  - **Response**:
    ```json
    {
      "message": "initial tests"
    }
    ```

---

### **Products**
- **GET `/products`**
  - Lists all registered products.
  - **Example Response**:
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

- **POST `/product`**
  - Creates a new product.
  - **Body**:
    ```json
    {
      "name": "Sample Product",
      "price": 100.00
    }
    ```
  - **Example Response**:
    ```json
    {
      "id": 1,
      "name": "Sample Product",
      "price": 100.00
    }
    ```

- **GET `/product/:productId`**
  - Returns a specific product based on the ID.
  - **Example Response**:
    ```json
    {
      "id": 1,
      "name": "Sample Product",
      "price": 100.00
    }
    ```

- **PUT `/product`**
  - Updates an existing product.
  - **Body**:
    ```json
    {
      "id": 1,
      "name": "Updated Product",
      "price": 150.00
    }
    ```
  - **Example Response**:
    ``` json
    {
      "id": 1,
      "name": "Updated Product",
      "price": 150.00
    }
    ```

- **DELETE `/product/:productId`**
  - Removes a product based on the ID.
  - **Example Response**:
    ``` json
    {
      "message": "Product successfully removed"
    }
    ```

---

## 🛠️ Technologies Used

- **Golang**: Main programming language.
- **Gin**: Lightweight and fast framework for API development.
- **Docker**: For container creation and management.
- **PostgreSQL** (or another relational database): Configured in the `db` module.

<br>

---
