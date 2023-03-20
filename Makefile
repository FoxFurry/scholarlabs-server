all: build-all push-all-latest configmap run-kube

build-all:
	docker build -t foxfurry/scholarlabs-gateway -f ./gateway/Dockerfile .

push-all-latest:
	docker push foxfurry/scholarlabs-gateway

configmap:
	$(eval ENVHASH := $(shell openssl dgst -sha256 -hex .env | awk '{print $$2}'))
	kubectl create configmap scholarlabs-$(ENVHASH) --from-env-file=.env -n=scholarlabs

run-kube:
	kubectl create -R -f ./infra -n=scholarlabs
