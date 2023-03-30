build:
	docker build -t foxfurry/scholarlabs-gateway -f ./services/gateway/Dockerfile .

run: build
	docker run --env-file=.env foxfurry/scholarlabs-gateway

push:
	docker push foxfurry/scholarlabs-gateway

configmap:
	kubectl create configmap scholarlabs-gateway-$(CONFIGHASH) --from-env-file=.env -n=scholarlabs