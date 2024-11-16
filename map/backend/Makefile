override GO_VERSION=1.23.3
override GOLANGCI_LINT_VERSION=v1.62.0

GOOS?=$(shell go env GOOS || echo linux)
GOARCH?=$(shell go env GOARCH || echo amd64)
CGO_ENABLED?=0

.PHONY: generate_server
generate_server:
	@rm -r gen/server.go || true
	@docker run --rm \
		-v ${PWD}:/project \
		-w /project \
		-e GOOS=${GOOS} \
		-e GOARCH=${GOARCH} \
		-e CGO_ENABLED=${CGO_ENABLED} \
		-e GO111MODULE=on \
		golang:${GO_VERSION} \
			go install github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen@latest && \
			go get github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen && \
			go run github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen --config=gen/cfg.yaml oapi/api.yaml

.PHONY: lint-go
lint-go:
	@docker run --rm \
		-v "${CURDIR}":/project \
		-w /project \
		golangci/golangci-lint:${GOLANGCI_LINT_VERSION} \
			golangci-lint run -v

.PHONY: mqtt-start
mqtt-start:
	@docker-compose up -f docker/mosquito/docker-compose.yaml --build