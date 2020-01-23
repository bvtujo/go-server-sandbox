cnf ?= config.env
include $(cnf)
export $(shell sed 's/=.*//' $(cnf))

dpl ?= deploy.env
include $(dpl)
export $(shell sed 's/=.*//' $(dpl))

build : 
	docker build  -t $(APP_NAME) .

run : 
	docker run -t -i --rm -p $(PORT):$(PORT) --name="$(APP_NAME)" $(APP_NAME)

up : build run

stop :
	docker stop $(APP_NAME); docker rm $(APP_NAME)

gobuild :
	go build main.go
