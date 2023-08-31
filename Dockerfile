FROM golang:alpine AS builder
WORKDIR /app
COPY . .
RUN go mod download
RUN apk add build-base && GO111MODULE="on" CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o app cmd/main.go
FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/app .
COPY --from=builder /app/config /app/config
CMD ["./app"]