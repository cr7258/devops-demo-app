FROM golang:alpine AS builder
WORKDIR /app
COPY . .
RUN go build -o http-server.go

FROM alpine
WORKDIR /app
COPY --from=builder /app/http-server /app/http-server
CMD ["./http-server"]