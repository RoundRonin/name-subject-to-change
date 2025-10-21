.PHONY: 
	run migrate sqlc mocks

run:
	\tgo run ./cmd/back/main.go

migrate:
	\tmigrate -path internal/db/migrations -database $$POSTGRES_DSN up

sqlc:
	\tsqlc generate

mocks:
	\tgo generate ./mocks
