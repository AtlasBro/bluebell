.PHONY: all build run gotool clean help

BINARY="bluebell"

all: gotool build

build:
	set CGO_ENABLED=0  GOOS=linux GOARCH=amd64 go build -o ${BINARY}

run:
	@go run main.go conf/config.yaml

gotool:
	go fmt ./
	go vet ./

clean:
	@if [ -f ${BINARY} ] ; then rm ${BINARY} ; fi

help:
	@echo "make - ��ʽ�� Go ����, ���������ɶ������ļ�"
	@echo "make build - ���� Go ����, ���ɶ������ļ�"
	@echo "make run - ֱ������ Go ����"
	@echo "make clean - �Ƴ��������ļ��� vim swap files"
	@echo "make gotool - ���� Go ���� 'fmt' and 'vet'"
