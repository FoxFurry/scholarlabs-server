CONFIGHASH := $(shell openssl dgst -sha256 -hex .env | awk '{print $$2}')

all: build-all push-all-latest run-kube

# All services

build-all:


push-all-latest:


# Kubernetes

run-kube:
	./scripts/run-kube.sh
