GOCMD=go
GOBUILD=$(GOCMD) build


BINARY_NAME=mongo-app
BUILD_PATH=./bin/$(BINARY_NAME)

build:
	$(GOBUILD) -o bin/$(BINARY_NAME) main.go
   
run:
	$(BUILD_PATH) $(MYSQL_HOST) 
 
