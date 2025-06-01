FROM golang:1.24-alpine AS builder
RUN apk add --no-cache git ca-certificates
WORKDIR /app
COPY . .
RUN go build -ldflags="-s -w" -o file-server main.go

FROM alpine:3.19
RUN apk add --no-cache ca-certificates
WORKDIR /app
COPY --from=builder /app/file-server .
COPY cert.pem key.pem .
COPY mpds /app/mpds/

EXPOSE 8080 443

CMD ["./file-server"]
