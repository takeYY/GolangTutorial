.PHONY: build
build:
	docker compose build --no-cache

.PHONY: up
up:
	docker compose up

.PHONY: up_d
up_d:
	docker compose up -d

.PHONY: exec_api
exec_api:
	docker exec -it golang_tutorial_api bash

.PHONY: stop
stop:
	docker compose stop

.PHONY: down_v
down_v:
	docker compose down -v

.PHONY: test
test:
	docker compose run --rm api gotest -v ./src/... -cover
	docker compose stop

.PHONY: gen_models
gen_models:
	docker compose run --rm api go run scripts/generate_models.go
	docker compose stop
