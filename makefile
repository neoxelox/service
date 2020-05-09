PACKAGES ?= ./...
COMPOSER ?= -f ./deploy/docker-compose.yml --project-directory . -p mst
FLAGS ?=
DEVTOOLS ?= $(shell $(GOPATH)/bin/golangci-lint --version 2>/dev/null | grep 1.24)


start:
	docker-compose $(COMPOSER) up -d
	docker logs -f mst_container
.SILENT: start

stop:
	docker-compose $(COMPOSER) stop
	docker stop $(shell docker ps -a -q)
.SILENT: stop

build:
	docker build -t mst:latest -f ./build/Dockerfile .
.PHONY: build
.SILENT: build

restart: stop build start
.SILENT: restart

remove: stop
	docker rm $(shell docker ps -a -q)
.SILENT: remove

prune: stop remove
	docker system prune --force -a
	docker volume prune --force
.SILENT: prune

test:
	go test -race -count=1 $(FLAGS) $(PACKAGES) -cover | tee coverage.out
	echo "\e[1m====================================="
	grep -Po "[0-9]+\.[0-9]+(?=%)" coverage.out | awk '{ SUM += $$1; PKGS += 1} END { print "  Total Coverage (" PKGS " pkg/s) : " SUM/PKGS "%"}'
	echo "=====================================\e[0m"
	rm -f coverage.out
.SILENT: test

cover:
	go test -race -count=1 $(PACKAGES) -coverprofile=coverage.out && go tool cover -html=coverage.out
	rm -f coverage.out
.SILENT: cover

lint: devtools
	$(GOPATH)/bin/golangci-lint run $(PACKAGES) -c ./configs/golangci.yml
.SILENT: lint

devtools:
ifeq ($(strip $(DEVTOOLS)),)
	echo "\e[1mDEVTOOLS not present, installing...\e[0m"
	curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(GOPATH)/bin v1.24.0
endif
.SILENT: devtools

clean:
	rm -f *.exe *.out