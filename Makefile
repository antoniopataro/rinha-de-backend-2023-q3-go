build: clean deps
	go build -o ./bin/main ./cmd/main.go

clean:
	rm -rf ./bin

deps:
	go mod tidy
	
dev:
	go run ./cmd/main.go

docker-dev: docker-down
	cp docker-compose.dev.yml docker-compose.yml
	docker compose up --build -d
	
docker-down:
	docker compose down -v --remove-orphans

docker-prod: docker-down
	cp docker-compose.prod.yml docker-compose.yml
	docker compose up --build -d

start: build
	./bin/main

stress-test:
	sh ./scripts/stress-test/run-test.sh
