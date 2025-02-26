# Build Stage
FROM golang:1.22-alpine AS build

WORKDIR /app

# Copy Go modules first for caching
COPY go.mod go.sum ./
RUN go mod download

# Copy source code
COPY . .

# Build the Go application
RUN go build -o main main.go

####################################################################
# Production Stage
FROM alpine:latest

WORKDIR /app

# Copy the built binary from the build stage
COPY --from=build /app/main .

# Copy the .env file to the production container
COPY .env .

# Expose the application port
EXPOSE 9000

# Run the application
ENTRYPOINT ["./main"]