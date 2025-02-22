FROM golang:1.22-alpine3.19 AS builder
WORKDIR /app
ENV GO111MODULE=on \
    GOPROXY=https://goproxy.cn,direct
COPY go.mod go.mod
COPY go.sum go.sum
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./bin/server .
FROM alpine:3.19
COPY --from=builder /app/bin/server /server
RUN chmod +x server
CMD ["./server"]
