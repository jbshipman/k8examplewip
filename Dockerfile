FROM golang:1.17.10-alpine3.16

LABEL version="1.0"
LABEL author="James Shipman"
LABEL tools="echo for api payload response"

ENV PORT=3030

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY *.go ./

RUN go build -o /littleapi

EXPOSE ${PORT}

CMD [ "/littleapi" ]
