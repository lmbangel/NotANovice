.PHONY: serve goose-up goose-down goose-status build up down stop clean clean-proto


proto-python:
	@cd services/quiz_agent && python3 -m grpc_tools.protoc -I. --python_out=. --grpc_python_out=. quiz.proto

proto: proto-python

clean-proto:
	@rm -f python/*_pb2.py python/*_pb2_grpc.py

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
	@rm -f python/*_pb2.py python/*_pb2_grpc.py

run-rebuild: stop down clean build up
