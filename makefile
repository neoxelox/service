PACKAGES ?= ./...
FLAGS ?=
DEVTOOLS ?= $(shell $(GOPATH)/bin/golangci-lint --version)

start:
	docker-compose up -d
	docker logs -f mst_container
.SILENT: start

stop:
	docker-compose down
	docker stop $(shell docker ps -a -q)
.SILENT: stop

build:
	docker-compose up -d --build
	docker logs -f mst_container
.SILENT: build

remove: stop
	docker rm $(shell docker ps -a -q)
.SILENT: remove

prune: stop remove
	docker system prune --force
.SILENT: prune

test:
	go test -race -count=1 $(FLAGS) $(PACKAGES) -cover | tee coverage.out
	echo "\e[1m=====================================\e[21m"
	grep -Po "[0-9]+\.[0-9]+(?=%)" coverage.out | awk '{ SUM += $$1; PKGS += 1} END { print "  Total Coverage (" PKGS " pkg/s) : " SUM/PKGS "%"}'
	echo "\e[1m=====================================\e[21m"
	rm -f coverage.out
.SILENT: test

cover:
	go test -race -count=1 $(PACKAGES) -coverprofile=coverage.out && go tool cover -html=coverage.out
	rm -f coverage.out
.SILENT: cover

lint: devtools
	$(GOPATH)/bin/golangci-lint run $(PACKAGES)
.SILENT: lint

devtools:
ifndef DEVTOOLS
	echo "\e[1mDEVTOOLS not present, installing...\e[21m"
	curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin v1.24.0
endif
.SILENT: devtools

clean:
	rm -f *.exe *.out