GOCMD=go
GOBUILD=$(GOCMD) build
BINARY_NAME=go-admin

all: deploy-test

deploy-test:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GOBUILD) -o ./build/$(BINARY_NAME) -tags=jsoniter -v ./
	ansible-playbook -i ./deploy/hosts ./deploy/deploy.yml

fmt:
	go fmt ./ecommerce/...
	go fmt ./login/...
	go fmt ./pages/...