GO111MODULE=on

.PHONY: build
build: gochat

.PHONY: gochat
gochat:
	go build $(RACE) -o bin/gochat.bin -tags=etcd main.go

.PHONY: stop
stop:
	pkill gochat

.PHONY: run
run: build run-logic run-ctcp run-cwebsocket run-task run-api

.PHONY: run-test
run: build run-logic run-ctcp run-cwebsocket run-task run-api run-site

.PHONY: run-logic
run-logic:
	nohup bin/gochat.bin -module logic &

.PHONY: run-ctcp
run-ctcp:
	nohup bin/gochat.bin -module connect_tcp &

.PHONY: run-cwebsocket
run-cwebsocket:
	nohup bin/gochat.bin -module connect_websocket &

.PHONY: run-task
run-task:
	nohup bin/gochat.bin -module task &

.PHONY: run-api
run-api:
	nohup bin/gochat.bin -module api &

.PHONY: run-site
run-site:
	nohup bin/gochat.bin -module site &

.PHONY: clean
clean:
	rm bin/gochat.bin
