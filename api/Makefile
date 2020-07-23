setup:
	go get github.com/cespare/reflex

run:
	cd app/http/ && go run main.go

watch:
	reflex -s -r '\.go$$' make run

test:
	cd app/http/ && go test -v

.PHONY: setup run watch test
