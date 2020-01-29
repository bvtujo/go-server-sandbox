cnf ?= config.env
include $(cnf)
export $(shell sed 's/=.*//' $(cnf))

dpl ?= deploy.env
include $(dpl)
export $(shell sed 's/=.*//' $(dpl))

build :
	go build -o bin/main src/main.go
dbuild : 
	docker build  -t $(APP_NAME) .

run : 
	docker run -t -i --rm -p $(PORT):$(PORT) --name="$(APP_NAME)" $(APP_NAME)

up : dbuild run

stop :
	docker stop $(APP_NAME); docker rm $(APP_NAME)


test :
	go test -v ./...
