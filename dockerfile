## We specify the base image we need for our
## go application

FROM golang:latest

## We create an /app directory within our
## image that will hold our application source
## files
WORKDIR /app


## We copy everything in the root directory
COPY . .

## We remove any existing go.mod and go.sum files
RUN rm go.mod
RUN rm go.sum

RUN go mod init urlapp
RUN go mod tidy


## we run go build to compile the binary
## executable of our Go program
RUN go build -o main cmd/main/main.go

## we run go test to run unit test cases for the application
## also get the code coverage result
RUN go test -v -cover -coverpkg=./... ./test


## Our start command which kicks off
## our newly created binary executable
CMD ["/app/main"]



