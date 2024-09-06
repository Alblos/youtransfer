FROM golang:1.23.1-alpine3.20

WORKDIR /app

COPY . .

RUN go get -d -v ./...

RUN go build -o main .

EXPOSE 8080

CMD ["/app/main"]