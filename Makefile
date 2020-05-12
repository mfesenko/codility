.PHONY: test install-golint lint

install-golint:
	GO111MODULE=on go get -v -u golang.org/x/lint/golint

test:
	go test -count=1 -cover ./...

lint: install-golint
	${HOME}/go/bin/golint -set_exit_status ./...
