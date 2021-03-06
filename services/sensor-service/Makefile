name := sensor-service
docker_tag ?= latest
docker_image ?= moorara/$(name)
docker_test_image ?= $(name)-test
docker_dir := /tmp/workspace


clean:
	@ rm -rf *.log coverage

run:
	@ go run main.go

build:
	@ ./scripts/build.sh --main main.go --binary $(name)

docker:
	@ docker build -t $(docker_image):$(docker_tag) .

docker-test:
	@ docker build --file Dockerfile.test --tag $(docker_test_image) .

up:
	@ VERSION=$(docker_tag) docker-compose up -d $(name)

down:
	@ VERSION=$(docker_tag) docker-compose down

test:
	@ go test -race ./...

coverage:
	@ ./scripts/test-unit-cover.sh

test-component: # docker up
	@ ./scripts/test-component.sh --version $(docker_tag)

test-component-docker: # docker docker-test
	@ ./scripts/test-component.sh --version $(docker_tag) --docker

push:
	@ docker push $(docker_image):$(docker_tag)

save-images:
	@ mkdir -p $(docker_dir)
	@ docker image save -o $(docker_dir)/docker.tar $(docker_image):$(docker_tag)
	@ docker image save -o $(docker_dir)/docker-test.tar $(docker_test_image)


.PHONY: clean
.PHONY: run build
.PHONY: docker docker-test up down
.PHONY: test coverage test-component test-component-docker
.PHONY: push save-images
