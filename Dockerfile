FROM golang:1.20-alpine

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod tidy

COPY . .

RUN go build -o patients-golang-api ./cmd/api/main.go

EXPOSE 8080

CMD ["./patients-golang-api"]