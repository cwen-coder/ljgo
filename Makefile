GIT_COMMIT=$(shell git rev-parse HEAD)
BUILD_TIME=$(shell date "+%Y-%m-%dT%H:%M:%S%z")
	VERSION_FLAG=-ldflags="-X github.com/cwen-coder/ljgo/app/command.GitCommit=$(GIT_COMMIT) -X github.com/cwen-coder/ljgo/app/command.BuildTime=$(BUILD_TIME)"

all: b

b:
	        CGO_ENABLED=0 go build -o bin/ljgo $(VERSION_FLAG) main.go

ins:
	        go install -a $(VERSION_FLAG) .

install:
	        go install $(VERSION_FLAG) .

