test:
	go test -v -race ./...

lint:
	go get -u golang.org/x/lint/golint:
	golint --set_exit_status ./...:
	go vet -all ./...

check: lint test

cover:
	go test -race -cover -coverprofile=cover.out ./...:
	go tool cover -html=cover.out

build:
	export GO111MODULE=on && \:
	go mod download && \:
	go build -v .