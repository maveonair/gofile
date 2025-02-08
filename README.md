# Go File Generator Server

This is a simple Go application that starts an HTTP server with an endpoint to generate files of specified sizes on the fly. The server allows you to request files of sizes between 1MB and 1000MB.

## Endpoints

- `GET /file/{size}`: Generates a file of the specified size in megabytes and returns it to the requester. The size must be between 1 and 1000.

## Usage

### Using Docker

The easiest way to run the server is using Docker:

```sh
docker run -p 8080:8080 ghcr.io/maveonair/gofile
```

### Running the server locally

1. Clone the repository:

   ```sh
   git clone https://github.com/maveonair/gofile.git
   cd gofile
   ```

2. Run the server

   ```sh
   go run cmd/gofile/run main.go
   ```

The server will start on port `8080`.

### Example Requests

- To generate a 1MB file:

  ```sh
  curl -O http://localhost:8080/file/1
  ```

- To generate a 10MB file:

  ```sh
  curl -O http://localhost:8080/file/10
  ```

- To generate a 100MB file:

  ```sh
  curl -O http://localhost:8080/file/100
  ```

- To generate a 1000MB file:
  ```sh
  curl -O http://localhost:8080/file/1000
  ```
