FROM golang:1.13
ADD . /api
WORKDIR /api
ENV GOPROXY=direct
RUN go build -o bin/api cmd/api
EXPOSE 8080
ENTRYPOINT ["./bin/api"]

