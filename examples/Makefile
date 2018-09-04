# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
BINARY_NAME=app
BINARY_UNIX=$(BINARY_NAME)_unix
TEMPLATES_DIR=templates
PKG_TEMPLATES=$(GOPATH)/src/github.com/yetkinozturk/streamview/$(TEMPLATES_DIR)


all: build
build: 
		$(GOBUILD) -o $(BINARY_NAME) -v
		cp -rf $(PKG_TEMPLATES) $(TEMPLATES_DIR)
test: 
		$(GOTEST) -v ./...
clean: 
		$(GOCLEAN)
		rm -f $(BINARY_NAME)
		rm -f $(BINARY_UNIX)
		rm -rf $(TEMPLATES_DIR)
run:
		$(GOBUILD) -o $(BINARY_NAME) -v ./...
		./$(BINARY_NAME)
deps:
		$(GOGET) github.com/gorilla/websocket
		$(GOGET) github.com/yetkinozturk/streamview


# Cross compilation
build-linux:
		CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GOBUILD) -o $(BINARY_UNIX) -v