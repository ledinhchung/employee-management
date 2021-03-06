.PHONY: deps clean build

deps:
	go get -u ./...

clean: 
	rm -rf ./hello-world/hello-world
	
build:
	GOOS=linux GOARCH=amd64 go build -o get-employee/get-employee ./get-employee
	GOOS=linux GOARCH=amd64 go build -o add-employee/add-employee ./add-employee
	GOOS=linux GOARCH=amd64 go build -o send-paysplis/send-paysplis ./send-paysplis

local:
	sam local start-api
