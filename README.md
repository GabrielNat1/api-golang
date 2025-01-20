# API-GO: Gerenciamento de Produtos

Este projeto é uma API para gerenciamento de produtos desenvolvida em **Go** utilizando o framework **Gin**. A arquitetura é modular, garantindo maior organização e manutenibilidade.

---

## 🚀 Inicialização do Projeto

### Pré-requisitos

Certifique-se de ter instalado:
- **Golang**: [Instalar Golang](https://go.dev/doc/install)
- **Docker**: [Instalar Docker](https://www.docker.com/products/docker-desktop)

---

### Como executar

1. Clone o projeto e navegue para a pasta:
   ```bash
   git clone https://github.com/usuario/api-go.git
   cd api-go
   ```

2. Configure a conexão com o banco de dados:
   Edite as configurações no arquivo `db/conn.go` para refletir suas credenciais.

3. Suba os serviços com Docker:
   ```bash
   docker-compose up
   ```

4. Acesse a API no endereço:
   - `http://localhost:8000`

---

## 📚 Estrutura do Projeto

A estrutura está organizada da seguinte maneira:

```plaintext
API-GO/
├── cmd/
│   └── main.go               # Arquivo principal para iniciar a aplicação
├── controller/
│   └── product_controller.go # Controladores responsáveis pelas rotas
├── db/
│   └── conn.go               # Configuração e conexão com o banco de dados
├── model/
│   ├── product.go            # Modelo do produto
│   └── response.go           # Estruturas de resposta para a API
├── repository/
│   └── product_repository.go # Camada de interação com o banco de dados
├── usecase/
│   └── product_usecase.go    # Regras de negócio e lógica da aplicação
├── .gitignore                # Arquivos ignorados no controle de versão
├── docker-compose.yml        # Configuração do Docker Compose
├── Dockerfile                # Configuração do container Docker
├── go.mod                    # Dependências do projeto
├── go.sum                    # Hash das dependências
└── README.md                 # Documentação do projeto
```

---

## 📖 Endpoints

### **Ping**
- **GET `/ping`**
  - Retorna uma mensagem de teste.
  - **Resposta**:
    ```json
    {
      "message": "primeiros testes"
    }
    ```

---

### **Produtos**
- **GET `/products`**
  - Lista todos os produtos cadastrados.
  - **Exemplo de Resposta**:
    ```json
    [
      {
        "id": 1,
        "name": "Produto 1",
        "price": 100.00
      },
      {
        "id": 2,
        "name": "Produto 2",
        "price": 150.00
      }
    ]
    ```

- **POST `/product`**
  - Cria um novo produto.
  - **Body**:
    ```json
    {
      "name": "Produto Exemplo",
      "price": 100.00
    }
    ```
  - **Exemplo de Resposta**:
    ```json
    {
      "id": 1,
      "name": "Produto Exemplo",
      "price": 100.00
    }
    ```

- **GET `/product/:productId`**
  - Retorna um produto específico com base no ID.
  - **Exemplo de Resposta**:
    ```json
    {
      "id": 1,
      "name": "Produto Exemplo",
      "price": 100.00
    }
    ```

- **PUT `/product`**
  - Atualiza um produto existente.
  - **Body**:
    ```json
    {
      "id": 1,
      "name": "Produto Atualizado",
      "price": 150.00
    }
    ```
  - **Exemplo de Resposta**:
    ```json
    {
      "id": 1,
      "name": "Produto Atualizado",
      "price": 150.00
    }
    ```

- **DELETE `/product/:productId`**
  - Remove um produto com base no ID.
  - **Exemplo de Resposta**:
    ```json
    {
      "message": "Produto removido com sucesso"
    }
    ```

---

## 🛠️ Tecnologias Utilizadas

- **Golang**: Linguagem de programação principal.
- **Gin**: Framework leve e rápido para desenvolvimento de APIs.
- **Docker**: Para criação e gerenciamento de containers.
- **PostgreSQL** (ou outro banco relacional): Configurado no módulo `db`.

---
