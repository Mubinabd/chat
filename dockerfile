FROM golang:1.22.5 AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o chat ./cmd
FROM alpine:latest
RUN apk add --no-cache ca-certificates
WORKDIR /app
COPY --from=builder /app/chat .
COPY .env .env 
RUN chmod +x chat
EXPOSE 8080
CMD ["./chat"]