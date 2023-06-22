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
	@openssl genrsa -out certificat/private.key 2048
	@openssl rsa -in certificat/private.key -outform PEM -pubout -out certificat/public.key
	@openssl req -new -x509 -sha256 -key certificat/private.key -out certificat/certificate.crt -subj "/CN=localhost" -addext "subjectAltName=DNS:localhost,IP:127.0.0.1" -days 3650

generate: proto-gen

proto-gen:
	@protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative proto/metrics.proto