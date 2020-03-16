.PHONY: build-all
build-all:
	CGO_ENABLED=0 go build -o ./producer/producer ./producer/main.go
	CGO_ENABLED=0 go build -o ./writer/writer ./writer/main.go

.PHONY: build-producer
build-producer:
	CGO_ENABLED=0 go build -o ./producer/producer ./producer/main.go

.PHONY: build-writer
build-writer:
	CGO_ENABLED=0 go build -o ./writer/writer ./writer/main.go

.PHONY: start
start:
	sudo docker-compose down --rmi local
	sudo docker-compose build --no-cache
	sudo docker-compose up -d
