.PHONY:

APPLICATION_NAME ?= srcm

build:
	docker build -t ${APPLICATION_NAME} -f Dockerfile .

run:
	docker run ${APPLICATION_NAME} -p 8081:8081