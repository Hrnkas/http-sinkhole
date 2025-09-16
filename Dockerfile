# Build stage
FROM golang:1.21-alpine AS builder

WORKDIR /app

# Copy go mod files first for better layer caching
COPY go.mod ./
# Only copy go.sum if it exists (optional for projects with no dependencies)
COPY go.su[m] ./ 
RUN go mod download

# Copy source code
COPY . .

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o http-sinkhole .

# Final stage
FROM alpine:latest

# Add ca-certificates for any HTTPS calls (if needed in future)
RUN apk --no-cache add ca-certificates

WORKDIR /root/

# Copy the binary from builder stage
COPY --from=builder /app/http-sinkhole .

# Expose the default port
EXPOSE 8080

# Command to run the executable
CMD ["./http-sinkhole"]