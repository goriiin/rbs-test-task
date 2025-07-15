
.PHONY:up down

.env:
	@if [ ! -f /docker/env ]; then \
		echo "File .env not found"; \
		cp ./docker/.env.example ./docker/.env; \
	else \
		echo "File .env already exists."; \
	fi

prune: down
	docker system prune -a -f --volumes

up: .env
	docker compose -f docker/docker-compose.yml build --pull back_go
	docker compose -f docker/docker-compose.yml up --force-recreate -d

down:
	docker compose -f docker/docker-compose.yml down -v
