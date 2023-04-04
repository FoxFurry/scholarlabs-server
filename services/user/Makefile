CONFIGHASH := $(shell openssl dgst -sha256 -hex .env | awk '{print $$2}')

build:
	docker build -t foxfurry/scholarlabs-user -f ./services/user/Dockerfile .

migrate:
	export $(cat .env | xargs) && go run main.go migrate

run: build
	docker run --env-file=.env foxfurry/scholarlabs-user

push:
	docker push foxfurry/scholarlabs-user

proto-generate:
	protoc \
	--go_opt=paths=source_relative \
	--go_out=. \
	--go-grpc_out=. --go-grpc_opt=paths=source_relative \
	./proto/user.proto

configmap:
	kubectl create configmap scholarlabs-user-$(CONFIGHASH) --from-env-file=.env -n=scholarlabs