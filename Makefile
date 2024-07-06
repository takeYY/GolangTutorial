.PHONY: build
build:
	docker-compose build --no-cache

.PHONY: up
up:
	docker-compose up

.PHONY: up_d
up_d:
	docker-compose up -d

.PHONY: exec_api
exec_api:
	docker exec -it golang_tutorial_api bash

.PHONY: stop
stop:
	docker-compose stop

.PHONY: down_v
down_v:
	docker-compose down -v

.PHONY: test
test:
	docker-compose run --rm api go test -v ./src/...
