FROM golang
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod tidy
COPY . .
RUN go build -o main app/cmd/app/main.go
CMD ["./main"]
