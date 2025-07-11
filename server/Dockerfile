FROM golang:1.24.0-alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o server cmd/api/main.go

FROM alpine
WORKDIR /root/
COPY --from=builder /app/server .
CMD ["./server"]
