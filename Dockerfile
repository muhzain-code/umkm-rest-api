# Tahap build
FROM golang:1.25 AS builder
WORKDIR /app

# Copy modul dulu untuk caching
COPY go.mod go.sum ./
RUN go mod download

# Copy seluruh project
COPY . .

# Build binary dari folder cmd/server
RUN CGO_ENABLED=0 GOOS=linux go build -o server ./cmd/server

# Tahap final
FROM alpine:3.18
RUN apk add --no-cache ca-certificates tzdata

WORKDIR /root/

# Copy hasil build
COPY --from=builder /app/server .

# Copy .env ke container agar Go bisa membaca
COPY .env ./

EXPOSE 8080
CMD ["./server"]