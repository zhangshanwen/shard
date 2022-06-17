Project = shard

NOW = $(shell date  '+%Y-%m-%d %H:%M:%S')
# GIT = $(shell git rev-parse HEAD)
GIT = $(shell git --no-pager log --decorate=short --pretty=oneline -n1)

### build flag
time_flag = 'github.com/zhangshanwen/$(Project)/admin_api/v1/version.buildTime=$(NOW)'
project_flag = 'github.com/zhangshanwen/$(Project)/initialize/conf.Project=$(Project)'
git_flag = 'github.com/zhangshanwen/$(Project)/admin_api/v1/version.git=$(GIT)'

### linux env
set_linux_env = CGO_ENABLED=0 GOOS=linux GOARCH=amd64

### go build
normal_build = go build  -ldflags "-X  $(time_flag) -X  $(project_flag)  -X $(git_flag)"   -o bin/$(Project)  cmd/api.go

.PHONY: check dist build run all check

all: build

check: test all build clean fmt todo legacy


build:
	`$(normal_build)`
linux:
	`$(set_linux_env) $(normal_build)`


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
