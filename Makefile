.PHONY: setup
setup:
	curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(GOBIN) v1.47.2

.PHONY: env
env:
	cp .envrc-develop .envrc
	direnv allow .

.PHONY: lint
lint:
	golangci-lint run