ifeq ($(OS),Windows_NT)     # is Windows_NT on XP, 2000, 7, Vista, 10...
	detected_OS := Windows
else
	detected_OS := $(shell uname)
endif

ifeq ($(detected_OS),Darwin)        # Mac OS X
	SED = "gsed"
else
	SED = "sed"
endif

default: rikka-build-test

# Rikka

rikka-build-test: rikka-build rikka-test

rikka-build:
	go build .

rikka-test:
	-./rikka -port 8000 -fsDebugSleep 5000 -https -level 0

qiniu: rikka-build
	-./rikka -port 8000 -plugin qiniu -bname rikka-qiniu -bhost odbw8jckg.bkt.clouddn.com

upai: rikka-build
	-./rikka -port 8000 -plugin upai -bname rikka-upai -bhost rikka-upai.b0.upaiyun.com

# Docker

IMAGE_NAME = 7sdream/rikka
OLD_VERSION = $(shell docker images | $(SED) -n 's:$(IMAGE_NAME) \+\([0-9.]\+\).*:\1:gp')
HAS_LATEST = $(shell docker images | $(SED) -n 's:$(IMAGE_NAME) \+\(latest\).*:\1:gp')
NEW_VERSION = $(shell $(SED) -n 's:\t*Version = "\([0-9.]\+\).*":\1:p' api/consts.go)
GIT_COMMIT = $(strip $(shell git rev-parse --short HEAD))

version:
	@echo "Current git commit: $(GIT_COMMIT)"
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
	docker build \
		--build-arg VERSION=$(NEW_VERSION) \
		--build-arg VCS_URL=$(shell git config --get remote.origin.url) \
		--build-arg VCS_REF=$(GIT_COMMIT) \
		--build-arg BUILD_DATE=$(shell date -u +"%Y-%m-%dT%H:%M:%SZ") \
		-t $(IMAGE_NAME):$(NEW_VERSION) .
	docker tag $(IMAGE_NAME):$(NEW_VERSION) $(IMAGE_NAME):latest

just-push:
	docker push $(IMAGE_NAME)

push: build just-push

# Clean

clean:
	rm -f ./rikka
	rm -rf files/
	rm -f debug
	rm -f rikkac/rikkac 
