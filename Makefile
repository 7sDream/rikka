default: rikka-build-test

# Rikka

rikka-build-test: rikka-build rikka-test

rikka-build:
	go build .

rikka-test:
	-./rikka -port 8000 -fsDebugSleep 5000

qiniu: rikka-build
	-./rikka -port 8000 -plugin qiniu -bname rikka -bhost od8qjnxw1.bkt.clouddn.com

# Docker

IMAGE_NAME = 7sdream/rikka
OLD_VERSION = $(shell docker images | sed -n 's:$(IMAGE_NAME) \+\([0-9.]\+\).*:\1:gp')
HAS_LATEST = $(shell docker images | sed -n 's:$(IMAGE_NAME) \+\(latest\).*:\1:gp')
NEW_VERSION = $(shell sed -n 's:\t*Version = "\([0-9.]\+\).*":\1:p' api/consts.go)
GIT_COMMIT = $(strip $(shell git rev-parse --short HEAD))

version:
	@echo "Will delete $(OLD_VERSION) $(HAS_LATEST)"
	@echo "Will build $(NEW_VERSION) latest"

confirm:
	@bash -c "read -s -n 1 -p 'Press any key to continue, Ctrl+C to stop'"

delete:
ifneq ($(OLD_VERSION),)
	docker rmi $(IMAGE_NAME):$(OLD_VERSION)
endif
ifneq ($(HAS_LATEST),)
	docker rmi $(IMAGE_NAME):latest
endif
	-docker rmi $(IMAGE_NAME):$(NEW_VERSION)

build: version confirm delete clean
	@docker build \
		--build-arg VERSION=$(NEW_VERSION) \
		--build-arg VCS_URL=$(shell git config --get remote.origin.url) \
  		--build-arg VCS_REF=$(GIT_COMMIT) \
  		--build-arg BUILD_DATE=$(shell date -u +"%Y-%m-%dT%H:%M:%SZ") \
		-t $(IMAGE_NAME):$(NEW_VERSION) .
	docker tag $(IMAGE_NAME):$(NEW_VERSION) $(IMAGE_NAME):latest

push: build 
	docker push $(IMAGE_NAME)

# Clean

clean:
	rm -f ./rikka
	rm -rf files/
	rm -f debug
	rm -f rikkac/rikkac 
