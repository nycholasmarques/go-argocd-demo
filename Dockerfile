# Build
FROM golang:1.22 AS builder
WORKDIR /app
COPY . .
RUN go mod init go-argocd-demo && go build -o server .

# Run
FROM debian:bookworm-slim
WORKDIR /app
COPY --from=builder /app/server .
EXPOSE 8081
CMD ["./server"]
