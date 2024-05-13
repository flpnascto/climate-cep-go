server:
	@docker compose build --no-cache && docker compose up -d

shutdown:
	@docker compose down

run:
	@go run cmd/server/main.go

test:
	@go test -v ./...

requests:
	@echo "\n\n=== Valid Request ====================="; curl localhost:8080/01001000
	@echo "\n\n=== Request with invalid cep format ==="; curl localhost:8080/123
	@echo "\n\n=== Request with not found cep ========"; curl localhost:8080/12345678

.PHONY: server shutdown run test requests