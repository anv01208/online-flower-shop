FROM golang:1.19.4
WORKDIR /app
COPY . .
RUN go mod download
RUN go build -o apiGateway ./cmd/main.go
EXPOSE 8080
CMD ["./apiGateway"]