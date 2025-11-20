# Project: REST API with JSON

A real backend API that serves JSON data.

## How it Works

*   **JSON**: We use struct tags (e.g., `` `json:"title"` ``) to tell Go how to convert our structs to JSON.
*   **Routing**: We check `r.Method` to decide if it's a GET (list) or POST (create) request.
*   **Mutex**: Since web servers are concurrent, we use a `sync.Mutex` to safely modify our list of articles.

## API Endpoints

1.  **GET /articles**: Get all articles.
2.  **POST /articles**: Create a new article.
    *   Body: `{"title": "My Title", "content": "Some content"}`

## How to Run

1.  Start the server:
    ```bash
    go run main.go
    ```

2.  Test with `curl` (in another terminal):
    *   **Get all**: `curl http://localhost:8080/articles`
    *   **Create**: `curl -X POST -d "{\"title\":\"New Post\",\"content\":\"Go is great\"}" http://localhost:8080/articles`
