MIGRATION_DIR := "storage/postgre/migrations"

.PHONY: migrate-generate
migrate-generate:
	$(GOPATH)/bin/goose -dir $(MIGRATION_DIR) create $(name) go

.PHONY: migrate-up
migrate-up:
	$(GOPATH)/bin/goose -dir $(MIGRATION_DIR) create $(name) up

.PHONY: migrate-down
migrate-down:
	$(GOPATH)/bin/goose -dir $(MIGRATION_DIR) create $(name) down

generate: mock-gen

mock-gen:
	@rm -rf ./test/mocks/packages
	@go generate ./...

generate: certificat-gen

certificat-gen:
	@mkdir certificat
	@openssl req -x509 -newkey rsa:4096 -sha256 -days 3650 -nodes -keyout certificat/local.key -out certificat/local.crt -subj "/CN=localhost" -addext "subjectAltName=DNS:localhost,IP:127.0.0.1"
