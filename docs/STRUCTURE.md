# 📚 Project Structure

-------------------------
API-GO/
├── cmd/
│   └── main.go               # Entry point of the application
├── controller/
│   └── product_controller.go # Routes and controllers
├── db/
│   └── conn.go               # Database configuration
├── model/
│   ├── product.go            # Product model
│   └── response.go           # API response structures
├── repository/
│   └── product_repository.go # Database interaction layer
├── usecase/
│   └── product_usecase.go    # Business logic
├── .gitignore                # Files ignored by git
├── docker-compose.yml        # Docker Compose configuration
├── Dockerfile                # Docker container configuration
├── go.mod                    # Go modules dependencies
├── go.sum                    # Dependency hashes
└── README.md                 # Project documentation
-------------------------
