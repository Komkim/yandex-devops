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
	@rm -rf ./certificat
	@mkdir certificat
#	@openssl genrsa -out certificat/private.key 2048
#	@openssl rsa -in certificat/private.key -outform PEM -pubout -out certificat/public.key
#	@openssl req -new -key certificat/private.key -out certificat/certificate.crt

	@openssl genrsa -out certificat/local.key 2048
	@openssl rsa -in certificat/local.key -outform PEM -pubout -out certificat/public.key
	@openssl req -new -x509 -sha256 -key certificat/local.key -out certificat/certificate.crt -subj "/CN=localhost" -addext "subjectAltName=DNS:localhost,IP:127.0.0.1" -days 3650

	@#openssl req -x509 -nodes -newkey rsa:2048 -keyout private.key -out certificat.crt -days 3650

	@#openssl req -x509 -newkey rsa:4096 -sha256 -days 3650 -keyout certificat/public.key


	@#openssl req -x509 -newkey rsa:4096 -sha256 -days 3650 -nodes -keyout certificat/public.key -out certificat/certificate.crt -subj "/CN=localhost" -addext "subjectAltName=DNS:localhost,IP:127.0.0.1"
