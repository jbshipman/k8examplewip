FROM golang:1.17.10

RUN mkdir /app

ADD . ./app

WORKDIR /app

EXPOSE 8180

RUN go build -o main .

CMD ["/app/main"]