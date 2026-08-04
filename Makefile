# Nestory Makefile

.PHONY: rebuild run

build-backend:
	docker compose up backend-builder --build

run-backend:
	docker compose up -d backend

rebuild:
	docker compose up --build

run:
	docker compose up -d

down:
	docker compose down

