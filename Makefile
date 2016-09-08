default: rikka-build-test

# Rikka

rikka-build-test: rikka-build rikka-test

rikka-build:
	go build .

rikka-test:
	-./rikka -port 8000 -fsDebugSleep 5000

# Docker

IMAGE_NAME = 7sdream/rikka
OLD_VERSION = $(shell docker images | sed -n 's:$(IMAGE_NAME) \+\([0-9.]\+\).*:\1:gp')
NEW_VERSION = $(shell sed -n 's:\t*Version = "\([0-9.]\+\).*":\1:p' api/consts.go)
GIT_COMMIT = $(strip $(shell git rev-parse --short HEAD))

version:
	@echo "Will delete $(OLD_VERSION)"
	@echo "Will build $(NEW_VERSION)"

confirm:
	@if [ "$(OLD_VERSION)" = "$(NEW_VERSION)" ]; then $(error No newer version than $(OLD_VERSION)); fi
	@bash -c "read -s -n 1 -p 'Press any key to continue'"

delete:
	docker rmi $(IMAGE_NAME):$(OLD_VERSION)
	docker rmi $(IMAGE_NAME):latest

build: version confirm clean delete 
	docker build \
		--build-arg VCS_URL=$(shell git config --get remote.origin.url) \
  		--build-arg VCS_REF=$(GIT_COMMIT) \
  		--build-arg BUILD_DATE=$(shell date -u +"%Y-%m-%dT%H:%M:%SZ") \
		-t $(IMAGE_NAME):$(NEW_VERSION) .
	docker tag -f $(IMAGE_NAME):$(NEW_VERSION) $(IMAGE_NAME):latest

push: build 
	dockler push $(IMAGE_NAME)

# Clean

clean:
	rm -f ./rikka
	rm -rf files
	rm -f debug
	rm -f rikkac/rikkac 
