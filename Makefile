# All services

build-all:
	docker build -t foxfurry/scholarlabs-gateway -f ./services/gateway/Dockerfile .
	docker build -t foxfurry/scholarlabs-user -f ./services/user/Dockerfile .
	docker build -t foxfurry/scholarlabs-course -f ./services/course/Dockerfile .

push-all:
	docker push foxfurry/scholarlabs-gateway
	docker push foxfurry/scholarlabs-user
	docker push foxfurry/scholarlabs-course

configmap-all:
	$(MAKE) -C ./services/gateway configmap
	$(MAKE) -C ./services/user configmap
	$(MAKE) -C ./services/course configmap


run-kube:
	./scripts/run-kube.sh

nuke:
	kubectl delete namespace scholarlabs
	kubectl create namespace scholarlabs
	kubectl delete namespace loki-stack
	kubectl create namespace loki-stack
	$(MAKE) configmap-all
	$(MAKE) run-kube

all: build-all push-all nuke