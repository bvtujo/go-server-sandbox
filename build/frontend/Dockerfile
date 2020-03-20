FROM golang:1.13
ADD . /fe
WORKDIR /fe
RUN go build -o bin/frontend cmd/frontend
EXPOSE 80
ENTRYPOINT ["./bin/frontend"]