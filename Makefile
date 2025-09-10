.PHONY: serve goose-up goose-down goose-status

serve: 
	@go run main.go

goose-up:
	@goose -dir migrations sqlite3 quiz.db up

goose-down:
	@goose -dir migrations sqlite3 quiz.db down

goose-status:
	@goose -dir migrations sqlite3 quiz.db status