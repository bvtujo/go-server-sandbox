FROM golang:1.13
ADD . /devx/points/frontend
WORKDIR /devx/points/frontend
ENV GOPROXY=direct
RUN go build -o bin/frontend main.go
RUN cp -r cmd/frontend/assets ./bin/assets
EXPOSE 80
ENTRYPOINT ["./bin/frontend"]
