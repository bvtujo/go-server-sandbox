FROM golang:1.13
ENV GOPROXY=direct
ADD . /app
WORKDIR /app
RUN go build -o bin/main src/main.go
ENTRYPOINT ["./bin/main"]
EXPOSE 80