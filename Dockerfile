FROM golang:1.20-alpine

WORKDIR /app

COPY . .

RUN go build main.go

# Replace [to_address] with your address
CMD ["./main", "address", "[to_address]"]
