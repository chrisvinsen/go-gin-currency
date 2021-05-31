# Foreign Exchange Currency - Gin Framework (Go Language)

Using Gin Web Frameworks written in Go (Golang)

## Manual Local Installation

Clone the repository
```bash
git clone https://github.com/chrisvinsen/go-gin-currency.git
```

Make sure you have GO111MODULE enabled, you can check it by writing:
```bash
go env
```
If GO111MODULE is OFF, turn it on by writing:
```bash
export GO111MODULE="on"
```
Run locally by writing:
```bash
go run server.go
```
or
```bash
go build -o app/go-gin-currency
./app/go-gin-currency
```
Open in your web browser:
[http://localhost:8080/](http://localhost:8080/) 

## Docker Installation

Clone the repository
```bash
git clone https://github.com/chrisvinsen/go-gin-currency.git
```

Build Images from Dockerfile
```bash
docker build . -f Dockerfile -t go-gin-currency
```

Running on Docker Container with Port 8080
```bash
docker run -p 8080:8080 go-gin-currency
```
Open in your web browser:
[http://localhost:8080/](http://localhost:8080/) 
