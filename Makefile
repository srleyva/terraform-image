BINARY=terraform
VERSION=0.0.1
GITHUB_USERNAME=srleyva
GITHUB_REPO=${GITHUB_USERNAME}/${BINARY}
DOCKER_REPO=sleyva97/$(BINARY)

all: docker-push

docker-build:
	docker build . -t ${DOCKER_REPO}:${VERSION}

docker-push: docker-build
	docker push ${DOCKER_REPO}:${VERSION}

