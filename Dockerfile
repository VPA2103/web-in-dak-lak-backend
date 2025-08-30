# Stage 1: Build app
FROM golang:1.24 AS builder

WORKDIR /app

# Copy go.mod và go.sum trước để cache dependency
COPY go.mod go.sum ./
RUN go mod download

# Copy toàn bộ source
COPY . .

# Build binary
RUN go build -o backend .

# Stage 2: Run app
FROM gcr.io/distroless/base-debian12
WORKDIR /app
COPY --from=builder /app/backend .

EXPOSE 8080
CMD ["./backend"]
