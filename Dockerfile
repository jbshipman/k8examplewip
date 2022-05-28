## Specify the base image we need for our
## go application
FROM golang:1.17.10

## create an /app directory within our
## image that will hold our application source
## files
RUN mkdir /app

## Copy everything in the root directory
## into our /app directory
ADD . /app

## Specify that we now wish to execute 
## any further commands inside our /app
## directory
WORKDIR /app

## Run go build to compile the binary
## executable of our Go program
RUN go build -o main .

## Our start command which kicks off
## our newly created binary executable

CMD ["/app/main"]