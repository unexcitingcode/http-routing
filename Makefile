CWD := $(shell pwd)
IMAGE := http-routing
DEV_IMAGE := $(IMAGE)-dev
BUILD_DOCKER := @docker run --rm \
    -v "$(CWD):/opt/$(IMAGE)" \
    -v "$(CWD)/tmp/.cache:/root/.cache" \
    -v "$(CWD)/tmp/go:/opt/go/packages" \
    -v /var/run/docker.sock:/var/run/docker.sock \
    -w "/opt/$(IMAGE)" \
    -t \
    $(DEV_IMAGE)
INTERACTIVE_DOCKER := @docker run --rm \
    -v "$(CWD):/opt/$(IMAGE)" \
    -v "$(CWD)/tmp/.cache:/root/.cache" \
    -v "$(CWD)/tmp/go:/opt/go/packages" \
    -v /var/run/docker.sock:/var/run/docker.sock \
    -w "/opt/$(IMAGE)" \
    -it \
    $(DEV_IMAGE)

.PHONY: build
build: image
	$(INTERACTIVE_DOCKER) go build

.PHONY: image
image:
	docker build -t "$(DEV_IMAGE)" .

.PHONY: dev-build
dev-build: image
	cp -r ./scripts/pre-commit.sh ./.git/hooks/pre-commit

.PHONY: ci-build
ci-build: image
	mkdir -p "$(CWD)/tmp/.cache/docker"
	docker save "$(DEV_IMAGE)" > "$(CWD)/tmp/.cache/docker/$(DEV_IMAGE).tar"
	$(BUILD_DOCKER) pre-commit install-hooks --color always
	sudo chmod -R ugo+r tmp
	ls -al
	ls -al tmp/.cache
	ls -al tmp/.cache/pre-commit

.PHONY: ci-load
ci-load:
	docker load < "$(CWD)/tmp/.cache/docker/$(DEV_IMAGE).tar"

.PHONY: test
test:
	$(INTERACTIVE_DOCKER) go test -v

.PHONY: pre-commit
pre-commit:
	$(BUILD_DOCKER) pre-commit run --color always

.PHONY: pre-commit-all
pre-commit-all:
	$(BUILD_DOCKER) pre-commit run --color always --all-files

.PHONY: shell
shell:
	$(INTERACTIVE_DOCKER) bash
