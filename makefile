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
	nohup bin/gochat.bin -module logic > log/logic.log &

.PHONY: run-ctcp
run-ctcp:
	nohup bin/gochat.bin -module connect_tcp > log/connect_tcp.log &

.PHONY: run-cwebsocket
run-cwebsocket:
	nohup bin/gochat.bin -module connect_websocket > log/connect_websocket.log &

.PHONY: run-task
run-task:
	nohup bin/gochat.bin -module task > log/task.log &

.PHONY: run-api
run-api:
	nohup bin/gochat.bin -module api > log/api.log &

.PHONY: run-site
run-site:
	nohup bin/gochat.bin -module site > log/site.log &

.PHONY: clean
clean:
	rm bin/gochat.bin
