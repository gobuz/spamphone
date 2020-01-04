.PHONY: build clean deploy

build:
#	dep ensure -v
	env GOOS=linux go build -ldflags="-s -w" -o bin/public_spam cmd/main.go
