# Lesson 10: Simple Web Server

Go has a production-ready web server built right into the standard library.

## Key Concepts

1.  **`net/http`**: The package for HTTP client and server implementations.
2.  **`http.HandleFunc`**: Maps a URL path (like `/`) to a function.
3.  **Handler Function**: Takes a `ResponseWriter` (to write the response) and a `Request` (containing info about the request).
4.  **`http.ListenAndServe`**: Starts the server and listens for requests.

## How to Run

1.  Run the server:
    ```bash
    go run main.go
    ```
2.  Open your browser to [http://localhost:8080](http://localhost:8080).
