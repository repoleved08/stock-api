# Stock Management API

This is a simple Golang API that performs CRUD (Create, Read, Update, Delete) operations on stock data. The API uses the `gorilla/mux` router and interacts with stock details, including `id`, `name`, `price`, and `company`.

## Features

- Add stock information
- Retrieve stock information
- Update stock information
- Delete stock information

## API Endpoints

### Base URL

```
http://localhost:8000/api/stock
```

### Endpoints

- **Create Stock (POST)**  
  Adds a new stock entry.  
  **URL:** `/api/stock`  
  **Method:** `POST`  
  **Request Body:**
  ```json
  {
      "name": "string",
      "price": float,
      "company": "string"
  }
  ```

- **Get All Stock (GET)**  
  Retrieves all stock entries.  
  **URL:** `/api/stock`  
  **Method:** `GET`

- **Get Stock by ID (GET)**  
  Retrieves a specific stock entry by `id`.  
  **URL:** `/api/stock/{id}`  
  **Method:** `GET`

- **Update Stock (PUT)**  
  Updates stock information for a specific `id`.  
  **URL:** `/api/stock/{id}`  
  **Method:** `PUT`  
  **Request Body:**
  ```json
  {
      "name": "string",
      "price": float,
      "company": "string"
  }
  ```

- **Delete Stock (DELETE)**  
  Deletes a specific stock entry by `id`.  
  **URL:** `/api/stock/{id}`  
  **Method:** `DELETE`

## Installation

1. **Clone the repository**
   ```bash
   git clone https://github.com/your-username/stock-api.git
   cd stock-api
   ```

2. **Install dependencies**
   Ensure you have Go installed. Run the following command to install the necessary packages:
   ```bash
   go mod tidy
   ```

3. **Run the API**
   ```bash
   go run main.go
   ```

4. **Access the API**
   The API will be running at `http://localhost:8000`.

## Project Structure

```
.
├── main.go         # Entry point of the application
├── routes.go       # API routing definitions
├── handlers.go     # Request handlers for CRUD operations
└── models.go       # Stock struct and any data models
```

## Example Request for Creating Stock

```bash
curl -X POST http://localhost:8000/api/stock \
  -H "Content-Type: application/json" \
  -d '{
        "name": "Apple Inc.",
        "price": 145.30,
        "company": "Apple"
      }'
```

## Technologies Used

- [Golang](https://golang.org/) — Programming language used to build the API.
- [Gorilla Mux](https://github.com/gorilla/mux) — HTTP router and URL matcher for Golang.

## Future Improvements

- Add authentication (e.g., JWT) for securing the API.
- Implement pagination and filtering for the stock listings.
- Add database support (e.g., PostgreSQL, MySQL).

