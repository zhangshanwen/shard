Project = shard

NOW = $(shell date  '+%Y-%m-%d %H:%M:%S')
# GIT = $(shell git rev-parse HEAD)
GIT = $(shell git --no-pager log --decorate=short --pretty=oneline -n1)

.PHONY: check dist build run all check

all: build

check: test all build clean fmt todo legacy


build:
	go build  -ldflags "-X  'github.com/zhangshanwen/$(Project)/api/v1.buildTime=$(NOW)'  -X 'github.com/zhangshanwen/$(Project)/api/v1.git=$(GIT)'"   -o bin/$(Project)  cmd/api.go

run:build
	./bin/$(Project)


clean:
	find . -name "*.DS_Store" -type f -delete
	rm -rf bin

test:
	go test -cover -race ./...


fmt:
	go fmt  ./...

todo:
	grep -rnw "TODO" internal

# Legacy code should be removed by the time of release
legacy:
	grep -rnw "\(LEGACY\|Deprecated\)" internal
