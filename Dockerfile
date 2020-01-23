FROM golang:1.13

ADD . /app
WORKDIR /app
ENTRYPOINT go run main.go
EXPOSE 80