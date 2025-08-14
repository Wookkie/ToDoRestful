FROM golang:1.23.11-alpine AS builder
WORKDIR /app
COPY . .
RUN go mod download
RUN go build -o todo cmd/main.go
RUN ls

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/todo .
CMD ["./todo"]
EXPOSE 8080