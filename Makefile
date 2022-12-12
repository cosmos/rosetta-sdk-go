.PHONY: deps lint format check-format test test-coverage add-license \
	check-comments check-license shorten-lines salus release mocks

# To run the the following packages as commands,
# it is necessary to use `go run <pkg>`. Running `go get` does
# not install any binaries that could be used to run
# the commands directly.
ADDLICENSE_INSTALL=go install github.com/google/addlicense@latest
ADDLICENSE_CMD=addlicense
ADDLICENSE_IGNORE=-ignore ".github/**/*" -ignore ".idea/**/*"
ADDLICENCE_SCRIPT=${ADDLICENSE_CMD} -c "Coinbase, Inc." -l "apache" -v ${ADDLICENSE_IGNORE}
GOIMPORTS_INSTALL=go install golang.org/x/tools/cmd/goimports@latest
GOIMPORTS_CMD=goimports
GOLINES_INSTALL=go install github.com/segmentio/golines@latest
GOLINES_CMD=golines
GOVERALLS_INSTALL=go install github.com/mattn/goveralls@latest
GOVERALLS_CMD=goveralls
GOLINT_CMD=go run golang.org/x/lint/golint
GO_PACKAGES=./asserter/... ./server/... ./errors/...
GO_FOLDERS=$(shell echo ${GO_PACKAGES} | sed -e "s/\.\///g" | sed -e "s/\/\.\.\.//g")
TEST_SCRIPT=go test ${GO_PACKAGES}
LINT_SETTINGS=golint,misspell,gocyclo,gocritic,whitespace,goconst,gocognit,bodyclose,unconvert,lll,unparam

build:
	go build ./...
	
deps:
	go get ./...

install:
	go install ./...

check-comments:
	${GOLINT_CMD} -set_exit_status ${GO_FOLDERS} .

lint:
	golangci-lint run --timeout 2m0s -v -E ${LINT_SETTINGS},gomnd

format:
	gofmt -s -w -l .
	${GOIMPORTS_INSTALL}
	${GOIMPORTS_CMD} -w .

check-format:
	! gofmt -s -l . | read;
	${GOIMPORTS_INSTALL}
	! ${GOIMPORTS_CMD} -l . | read;

test:
	${TEST_SCRIPT}

test-cover:
	${GOVERALLS_INSTALL}
	if [ "${COVERALLS_TOKEN}" ]; then ${TEST_SCRIPT} -coverprofile=c.out -covermode=count; ${GOVERALLS_CMD} -coverprofile=c.out -repotoken ${COVERALLS_TOKEN}; fi

add-license:
	${ADDLICENSE_INSTALL}
	${ADDLICENCE_SCRIPT} .

check-license:
	${ADDLICENSE_INSTALL}
	${ADDLICENCE_SCRIPT} -check .

shorten-lines:
	${GOLINES_INSTALL}
	${GOLINES_CMD} -w --shorten-comments ${GO_FOLDERS} examples

salus:
	docker run --rm -t -v ${PWD}:/home/repo coinbase/salus

release: check-license check-format test lint salus