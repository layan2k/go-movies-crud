# Go Movies CRUD

This project is a simple CRUD (Create, Read, Update, Delete) application for managing a list of movies, built with Go.

## Purpose

The main goal of this project is to help you learn and practice Go programming by building a practical web application. You'll get hands-on experience with:

- Structs and slices
- HTTP routing and handlers
- JSON encoding/decoding
- Basic RESTful API design

## Features

- List all movies
- Add a new movie
- Update an existing movie
- Delete a movie

## Getting Started

1. **Clone the repository:**

    ```bash
    git clone https://github.com/layan2k/go-movies-crud.git
    cd go-movies-crud
    ```

2. **Run the application:**

    ```bash
    go run main.go
    ```

3. **Access the API:**
    - The server will start on `http://localhost:8080`
    - Use tools like [Postman](https://www.postman.com/) or `curl` to interact with the endpoints.

## Endpoints

| Method | Endpoint        | Description         |
|--------|----------------|---------------------|
| GET    | /movies        | List all movies     |
| POST   | /movies        | Add a new movie     |
| PUT    | /movies/{id}   | Update a movie      |
| DELETE | /movies/{id}   | Delete a movie      |

## License

This project is for educational purposes.
