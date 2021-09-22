APP_NAME = reconciler
IMG_REPO := $(DOCKER_PUSH_REPOSITORY)$(DOCKER_PUSH_DIRECTORY)
IMG_NAME := $(IMG_REPO)/$(APP_NAME)
TAG := $(DOCKER_TAG)
COMPONENTS := $(shell (go run scripts/reconcilernames.go))

ifndef VERSION
	VERSION = ${shell git describe --tags --always}
endif

ifeq ($(VERSION),stable)
	VERSION = stable-${shell git rev-parse --short HEAD}
endif

.DEFAULT_GOAL=all
FLAGS = -ldflags '-s -w'

.PHONY: resolve
resolve:
	go mod tidy

.PHONY: lint
lint:
	./scripts/lint.sh

.PHONY: build
build: build-linux build-darwin build-linux-arm

.PHONY: build-linux
build-linux:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./bin/reconciler-linux $(FLAGS) ./cmd/reconciler
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./bin/mothership-linux $(FLAGS) ./cmd/mothership

.PHONY: build-linux-arm
build-linux-arm:
	CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -o ./bin/reconciler-arm $(FLAGS) ./cmd/reconciler
	CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -o ./bin/mothership-arm $(FLAGS) ./cmd/mothership

.PHONY: build-darwin
build-darwin:
	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -o ./bin/reconciler-darwin $(FLAGS) ./cmd/reconciler
	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -o ./bin/mothership-darwin $(FLAGS) ./cmd/mothership

.PHONY: docker-build
docker-build:
	docker build -t $(APP_NAME)/mothership:latest -f Dockerfile.mr .
	docker build -t $(APP_NAME)/component:latest -f Dockerfile.cr .

.PHONY: docker-push
docker-push:
ifdef IMG_REPO
	docker tag $(APP_NAME)/mothership $(IMG_NAME)/mothership:$(TAG)
	docker tag $(APP_NAME)/component $(IMG_NAME)/component:$(TAG)
	docker push $(IMG_NAME)/mothership:$(TAG)
	docker push $(IMG_NAME)/component:$(TAG)
endif

.PHONY: bump-primage
bump-primage:
	./scripts/bumpimage.sh

.PHONY: deploy
deploy:
	@./scripts/kcversion.sh
	kubectl create namespace reconciler --dry-run=client -o yaml | kubectl apply -f -
	helm template reconciler --namespace reconciler --set "global.components={$(COMPONENTS)}" ./resources/reconciler > reconciler.yaml
	kubectl apply -f reconciler.yaml
	rm reconciler.yaml

.PHONY: test
test:
	go test -race -coverprofile=cover.out ./...
	@echo "Total test coverage: $$(go tool cover -func=cover.out | grep total | awk '{print $$3}')"
	@rm cover.out

.PHONY: test-all
test-all: export RECONCILER_INTEGRATION_TESTS = 1
test-all: test

.PHONY: clean
clean:
	rm -rf bin

.PHONY: all
all: resolve build test lint docker-build docker-push

.PHONY: release
release: all
