build-all:
	docker build -t foxfurry/scholarlabs-gateway -f ./gateway/Dockerfile .

push-all-latest:
	docker push foxfurry/scholarlabs-gateway

run-kube:
	kubectl apply -R -f ./infra

