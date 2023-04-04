CONFIGHASH := $(shell openssl dgst -sha256 -hex .env | awk '{print $$2}')

build:
	docker build -t foxfurry/scholarlabs-gateway .

run: build
	docker run --env-file=.env foxfurry/scholarlabs-gateway

push:
	docker push foxfurry/scholarlabs-gateway

configmap:
	kubectl create configmap scholarlabs-gateway-$(CONFIGHASH) --from-env-file=.env -n=scholarlabs