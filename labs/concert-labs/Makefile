BUF_VERSION 				:= v1.55.1
SWAGGER_UI_VERSION 			:= v4.15.5

install.buf:
	go install github.com/bufbuild/buf/cmd/buf@latest

install.protoc:
	go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
	go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway@latest
	go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2@latest

gen.protoc: install.protoc
	go run github.com/bufbuild/buf/cmd/buf@$(BUF_VERSION) generate

gen.swagger-ui:
	chmod +x ./scripts/generate-swagger-ui.sh
	SWAGGER_UI_VERSION=$(SWAGGER_UI_VERSION) ./scripts/generate-swagger-ui.sh

buf.update:
	buf mod update

buf.generate:
	buf generate

##########################################################################################
### DEV
##########################################################################################
dev.boot: install.protoc
	go install github.com/vektra/mockery/v2@latest
	go install github.com/golang/mock/mockgen@latest
	go install github.com/bufbuild/buf/cmd/buf@$(BUF_VERSION)
	go get ./...
	go mod tidy

dev.gen:
	go generate ./...

dev.tidy:
	go mod tidy

##########################################################################################
### DOCKER
##########################################################################################
docker.up:
	docker compose up -d --build

docker.down:
	docker compose down
	docker volume prune -a