.PHONY: build docker-build clean

APP_NAME := go-app-demo
IMAGE_TAG := latest
IMAGE_NAME := local/$(APP_NAME):$(IMAGE_TAG)

# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test

build:
	cd go-app-demo && $(GOBUILD) -o bin/$(APP_NAME) -v

clean:
	cd go-app-demo && $(GOCLEAN)
	rm -f go-app-demo/bin/$(APP_NAME)

docker-build:
	cd go-app-demo && docker build -t $(IMAGE_NAME) .

# Helper to deploy the standard app
deploy-app:
	kubectl apply -f go-app-demo/deployment.yaml

# Helper to delete the standard app
delete-app:
	kubectl delete -f go-app-demo/deployment.yaml
