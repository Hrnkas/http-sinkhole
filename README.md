# HTTP Sinkhole

A simple HTTP service written in Go that accepts any request and always returns 200/OK.

This service acts as a "sinkhole" for HTTP traffic - it accepts all incoming requests regardless of:
- HTTP method (GET, POST, PUT, DELETE, PATCH, etc.)
- URL path
- Query parameters  
- Request headers
- Request body

All requests receive the same response: HTTP 200 OK with a simple "OK" message body.

## Features

- **Universal HTTP handler**: Accepts all HTTP methods and paths
- **Graceful shutdown**: Handles SIGINT and SIGTERM signals properly
- **Configurable port**: Uses PORT environment variable or defaults to 8080
- **Request logging**: Logs all incoming requests for debugging
- **Proper timeouts**: Includes read, write, and idle timeouts
- **Lightweight**: Minimal dependencies, just standard Go library

## Usage

### Building and Running

```bash
# Clone the repository
git clone https://github.com/Hrnkas/http-sinkhole.git
cd http-sinkhole

# Build the application
go build .

# Run with default port (8080)
./http-sinkhole

# Or run with custom port
PORT=3000 ./http-sinkhole
```

### Running with Go

```bash
# Run directly with Go
go run main.go

# Run with custom port
PORT=3000 go run main.go
```

### Testing

```bash
# Run the test suite
go test -v

# Test manually with curl
curl http://localhost:8080/
curl -X POST http://localhost:8080/api/test
curl -X DELETE http://localhost:8080/any/path/here
```

### Example Requests

All of these requests will return `200 OK`:

```bash
curl http://localhost:8080/
curl -X POST -d '{"data":"test"}' http://localhost:8080/api/users
curl -X PUT http://localhost:8080/v1/items/123
curl -X DELETE http://localhost:8080/admin/delete?force=true
curl -X PATCH -H "Authorization: Bearer token" http://localhost:8080/user/profile
```

## Use Cases

- **Testing**: Use as a mock server for testing HTTP clients
- **Load testing**: Target for performance testing without backend processing
- **Development**: Placeholder for APIs that aren't implemented yet
- **Debugging**: Capture and log HTTP traffic patterns
- **Prototyping**: Quick HTTP endpoint for early development phases

## Configuration

| Environment Variable | Default | Description |
|---------------------|---------|-------------|
| `PORT` | `8080` | Port number for the HTTP server |

## License

This project is open source and available under the MIT License.
