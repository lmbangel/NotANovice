.PHONY: serve goose-up goose-down goose-status build up down stop clean

serve: 
	@go run main.go

goose-up:
	@goose -dir migrations sqlite3 quiz.db up

goose-down:
	@goose -dir migrations sqlite3 quiz.db down

goose-status:
	@goose -dir migrations sqlite3 quiz.db status

build:
	@docker compose build

up: 
	@docker compose up -d

down: 
	@docker compose down

stop:
	@docker compose stop

clean:
	@docker compose rm -f
