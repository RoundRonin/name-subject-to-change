# ========= Project Setup ========= #
APP_NAME := back
CMD_DIR := cmd/$(APP_NAME)
DOCS_DIR := docs
MIGRATIONS_DIR := internal/db/migrations
DB_URL ?= $(POSTGRES_DSN)  # example: postgres://user:pass@localhost:5432/dbname?sslmode=disable

# ========= Tooling ========= #
SWAG := swag
MIGRATE := migrate
SQLC := sqlc
MOCKGEN := go generate
WATCH := air

# ========= Targets ========= #
.PHONY: all run build clean swag sqlc mocks migrate-up migrate-down migrate-force watch

# Default target
all: swag sqlc build

# Build binary
build:
	@echo "🔨 Building $(APP_NAME)..."
	go build -o bin/$(APP_NAME) ./$(CMD_DIR)

# Run app (without rebuild)
run:
	@echo "🚀 Running $(APP_NAME)..."
	go run ./$(CMD_DIR)

# Watch for changes and auto-reload
watch:
	@echo "👀 Watching for changes (requires air)..."
	$(WATCH)

# Clean build artifacts
clean:
	@echo "🧹 Cleaning..."
	rm -rf bin/$(APP_NAME) $(DOCS_DIR)

# ========= Code Generation ========= #

# Generate Swagger documentation
swag:
	@echo "📚 Generating Swagger docs..."
	$(SWAG) init -g $(CMD_DIR)/main.go -o $(DOCS_DIR)

# Generate SQLC code
sqlc:
	@echo "🧩 Generating SQLC code..."
	$(SQLC) generate

# Generate mocks
mocks:
	@echo "🎭 Generating mocks..."
	$(MOCKGEN) ./...

# ========= Database Migrations ========= #

# Apply all up migrations (or specify version)
migrate-up:
	@echo "📈 Applying migrations (up)..."
	$(MIGRATE) -path $(MIGRATIONS_DIR) -database "$(DB_URL)" up $(if $(v),$(v),)

# Rollback migrations (down)
migrate-down:
	@echo "📉 Reverting migrations (down)..."
	$(MIGRATE) -path $(MIGRATIONS_DIR) -database "$(DB_URL)" down $(if $(v),$(v),)

# Force migrate state (use carefully)
migrate-force:
	@echo "⚠️  Forcing migration version..."
	$(MIGRATE) -path $(MIGRATIONS_DIR) -database "$(DB_URL)" force $(v)

# ========= Combined Dev Workflow ========= #

# Prepare everything for development
dev: swag sqlc mocks run

# Regenerate everything + rebuild
rebuild: clean swag sqlc build
