## Steps to run the application in a docker
    docker build -t my-go-app .

    docker run -p 8081:8081 -it my-go-app


## local setup to run the application
    rm go.mod go.sum

    go mod init urlapp

    go mod tidy

    go run cmd/main/main.go



## Steps to generate and view the code coverage report in html
    go test -coverprofile=coverage.out -coverpkg=./... ./test

    go tool cover -html=coverage.out


## Project Tree Structure
    .
    ├── cmd
    │   └── main
    │       └── main.go
    ├── config
    │   └── config.yaml
    ├── dockerfile
    ├── go.mod
    ├── go.sum
    ├── initapp
    │   └── Initapp.go
    ├── internal
    │   └── utilities
    │       ├── readConfig.go
    │       ├── urlShortener.go
    │       └── urlUtils.go
    ├── README.md
    └── test
        ├── urlShorterner_test.go
        └── urlUtils_test.go

    7 directories, 12 files


## Application usage

### To shorten the url
    http request : http://localhost:8081/shorten?url=https://github.com/iamrachanagupta/Go

    http response : http://localhost:8081/shorturl/MTM2MTYx


### To get the original link 
    Visit the http response url : http://localhost:8081/shorturl/MTM2MTYx

    It should redirect to original link:  https://github.com/iamrachanagupta/Go



## Setting up the log level
Edit the config/config.yaml file

Change the loglevel as per Golang logrus package standand:

    PanicLevel Level = 0

	FatalLevel Level = 1

	ErrorLevel Level = 2

	WarnLevel Level  = 3

	InfoLevel Level  = 4

	DebugLevel Level = 5

	TraceLevel Level = 6
