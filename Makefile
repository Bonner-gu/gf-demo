GOPATH:=$(shell go env GOPATH)
OUTPUT=main

default: test

.PHONY: test
test:
	go build -o ${OUTPUT} main.go;ls -l ${OUTPUT};md5sum ${OUTPUT};chmod 755 ${OUTPUT};./${OUTPUT}

.PHONY: build
build:
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o ${OUTPUT} main.go;ls -l ${OUTPUT};md5sum ${OUTPUT}
#
# .PHONY: docker
# docker:
# 	docker build . -t ${OUTPUT}:latest
#
# .PHONY: run
# run:build
# 	docker-compose down ; docker-compose up -d --build ; docker-compose ps ;docker-compose restart app ;docker-compose logs -f app
#
# .PHONY: clean
# clean:
# 	docker-compose down ; docker images|grep none|awk '{print $3}'|xargs docker rmi
#
# .PHONY: ps
# ps:
# 	docker-compose ps