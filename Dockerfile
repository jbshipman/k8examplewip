FROM golang:1.17.10

RUN mkdir /app

COPY go.mod ./app
COPY go.sum ./app

RUN go mod download

WORKDIR /app

EXPOSE 8180

RUN go build -o main .

CMD ["/app/main"]