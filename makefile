build: up exec

up:
	docker-compose up -d

down:
	docker-compose down -v

exec:
	docker exec -it calc bash

run:
	go run main.go

add:

	curl -X POST localhost:8080/evaluate -d '{"expression":"What is 5 plus 3"}'

subtract:

	curl -X POST localhost:8080/evaluate -d '{"expression":"What is 5 minus 3"}'

multiply:

	curl -X POST localhost:8080/evaluate -d '{"expression":"What is 5 multiplied by 3"}'

divide:

	curl -X POST localhost:8080/evaluate -d '{"expression":"What is 88 divided 4"}'

validate:

	curl -X POST localhost:8080/validate -d '{"expression":"What is 88 divided 4"}'

validate_err:

	curl -X POST localhost:8080/validate -d '{"expression":"What is 88 divided 4 plus minus"}'

errors:

	curl localhost:8080/errors

test:

	go test ./...