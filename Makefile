SOURCEDIR=.

PROJECT_NAME=hackzurich2019-be
PROJECT_ROOT=`git rev-parse --show-toplevel`
PROJECT_TOOL_PATH=/go/src/github.com/repa40x/$(PROJECT_NAME)

SWAGGER_VERSION=v0.20.1
SWAGGER_CMD=docker run --rm -it -v $(PROJECT_ROOT):$(PROJECT_TOOL_PATH) -w $(PROJECT_TOOL_PATH) quay.io/goswagger/swagger:$(SWAGGER_VERSION)

DOCKER_REPO=zrh.ocir.io/zrywfei8cmcr/hz2019/$(PROJECT_NAME)
VERSION=0.13
TAG=$(DOCKER_REPO):$(VERSION)

.DEFAULT_GOAL: validate

.PHONY: install-swagger validate vendor

install-swagger:
	docker pull quay.io/goswagger/swagger:$(SWAGGER_VERSION)

validate:
	$(SWAGGER_CMD) validate ./openapi/swagger.yaml

generate-full:
	$(SWAGGER_CMD) generate server -A $(PROJECT_NAME) -f ./openapi/swagger.yaml --target ./generated

generate:
	$(SWAGGER_CMD) generate server -A $(PROJECT_NAME) -f ./openapi/swagger.yaml --target ./generated --exclude-main

vendor:
	go mod vendor

build:
	go build

run:
	godotenv -f ./local/.env ./hackzurich2019-be --scheme=http --host=0.0.0.0 --port=8080

docker-login:
	docker login -u $(DOCKER_USER) -p $(DOCKER_TOKEN) $(DOCKER_URL)

docker-build:
	docker build --pull -f ./deployment/Dockerfile -t $(TAG) .

docker-push:
	docker push $(TAG)

docker-rm:
     docker image rm $(TAG)

docker-run:
	docker rm -f $(PROJECT_NAME) || true && docker run -dt --name $(PROJECT_NAME) --rm -p 8080:8080 $(TAG)

docker-stop:
	docker rm -f $(PROJECT_NAME) || true