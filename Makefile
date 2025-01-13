BUILDER = ./bin/callisto-builder

.PHONY: znet-start
znet-start:
	$(BUILDER) znet start --profiles=monitoring

.PHONY: znet-remove
znet-remove:
	$(BUILDER) znet remove

.PHONY: test
test:
	docker run --name callisto-test-db -e POSTGRES_USER=callisto -e POSTGRES_PASSWORD=password -e POSTGRES_DB=callisto -d -p 6433:5432 postgres
	$(BUILDER) test
	docker stop callisto-test-db
	docker rm callisto-test-db

.PHONY: build
build:
	$(BUILDER) build

.PHONY: build/arm64
build/arm64:
	$(BUILDER) build/arm64

.PHONY: build/amd64
build/amd64:
	$(BUILDER) build/amd64

.PHONY: images
images:
	$(BUILDER) images

.PHONY: release-images
release-images:
	$(BUILDER) release/images