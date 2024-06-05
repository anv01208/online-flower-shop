FROM golang:1.19.4
WORKDIR /app
COPY . .
RUN go mod download
RUN go build -o authSvc ./cmd/main.go
EXPOSE 8080
CMD ["./authSvc"]