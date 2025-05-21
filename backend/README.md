# ArtVideoTube Backend

This project is a backend API built using Go, following the principles of Clean Architecture. It provides endpoints to manage items, allowing users to retrieve and create items through HTTP requests.

## Project Structure

```
backend/
├── cmd
│   └── main.go               # Entry point of the application
├── internal
│   ├── domain
│   │   └── item              # Domain models and core business logic
│   ├── presentation
│   │   └── http
│   │       ├── handler       # HTTP handlers for API endpoints
│   │       └── router        # Router setup for HTTP endpoints
│   ├── repository
│   │   └── item              # Data access layer (e.g., GORM-based repository)
│   ├── service
│   │   └── item              # Application services (business logic orchestration)
│   └── usecase
│       └── item              # Use cases (application-specific business rules)
├── go.mod                    # Module dependencies
├── go.sum                    # Module checksums
└── README.md                 # Project documentation
```

## Setup Instructions

1. **Clone the repository:**
   ```bash
   git clone <repository-url>
   cd backend
   ```

2. **Install dependencies:**
   ```bash
   go mod tidy
   ```

3. **Run the application:**
   ```bash
   go run cmd/main.go
   ```

## Usage

### Endpoints

- **Get Items**
  - **Endpoint:** `GET /items`
  - **Description:** Retrieves a list of items.

- **Create Item**
  - **Endpoint:** `POST /items`
  - **Description:** Creates a new item. Requires a JSON body with item details:
    ```json
    {
      "id": "1",
      "title": "Sample Item",
      "description": "This is a sample item.",
      "url": "https://example.com/item"
    }
    ```

## Architecture

This project follows the principles of Clean Architecture:

- **Domain Layer:** Contains the core business logic and domain models.
- **Use Case Layer:** Defines application-specific business rules and orchestrates interactions between the domain and other layers.
- **Service Layer:** Implements business logic and interacts with the repository layer.
- **Repository Layer:** Handles data access (e.g., database operations using GORM).
- **Presentation Layer:** Manages HTTP requests and responses, including routing and handlers.

## License

This project is licensed under the MIT License. See the LICENSE file for details.