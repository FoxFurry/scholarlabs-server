CONFIGHASH := $(shell openssl dgst -sha256 -hex .env | awk '{print $$2}')

build:
	docker build -t foxfurry/scholarlabs-course -f ./services/course/Dockerfile .

migrate:
	export $(cat .env | xargs) && go run main.go migrate

run: build
	docker run --env-file=.env foxfurry/scholarlabs-course

push:
	docker push foxfurry/scholarlabs-course

proto-generate:
	protoc \
	--go_opt=paths=source_relative \
	--go_out=. \
	--go-grpc_out=. --go-grpc_opt=paths=source_relative \
	./proto/course.proto

configmap:
	kubectl create configmap scholarlabs-course-$(CONFIGHASH) --from-env-file=.env -n=scholarlabs