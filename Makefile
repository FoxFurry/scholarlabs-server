all: build-all push-all-latest configmap run-kube

build-all:
	docker build -t foxfurry/scholarlabs-gateway -f ./gateway/Dockerfile .

push-all-latest:
	docker push foxfurry/scholarlabs-gateway

configmap:
	kubectl delete configmap scholarlabs
	kubectl create configmap scholarlabs --from-env-file=.env

run-kube:
	kubectl delete pods --ignore-not-found=true --all
	kubectl apply -R -f ./infra
