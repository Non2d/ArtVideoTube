# Go Render API

This project is a simple API built using Go. It provides endpoints to manage items, allowing users to retrieve and create items through HTTP requests.

## Project Structure

```
go-render-api
├── cmd
│   └── main.go          # Entry point of the application
├── internal
│   ├── handlers
│   │   └── api.go      # API request handlers
│   └── models
│       └── model.go    # Data models
├── go.mod               # Module dependencies
├── go.sum               # Module checksums
└── README.md            # Project documentation
```

## Setup Instructions

1. **Clone the repository:**
   ```
   git clone <repository-url>
   cd go-render-api
   ```

2. **Install dependencies:**
   ```
   go mod tidy
   ```

3. **Run the application:**
   ```
   go run cmd/main.go
   ```

## Usage

- **Get Items:**
  - Endpoint: `GET /api/items`
  - Description: Retrieves a list of items.

- **Create Item:**
  - Endpoint: `POST /api/items`
  - Description: Creates a new item. Requires a JSON body with item details.

## License

This project is licensed under the MIT License. See the LICENSE file for details.