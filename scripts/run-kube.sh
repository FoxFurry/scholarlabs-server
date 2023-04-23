#/bin/sh
GATEWAYCONFIGHASH=$(openssl dgst -sha256 -hex ./services/gateway/.env | awk '{print $2}')
USERCONFIGHASH=$(openssl dgst -sha256 -hex ./services/user/.env | awk '{print $2}')
COURSECONFIGHASH=$(openssl dgst -sha256 -hex ./services/course/.env | awk '{print $2}')

for f in $(find ./infra/k8s/ -name '*.yaml');
do
  GATEWAYCONFIGHASH=${GATEWAYCONFIGHASH} \
  USERCONFIGHASH=${USERCONFIGHASH} \
  COURSECONFIGHASH=${COURSECONFIGHASH} \
  envsubst < $f | kubectl apply -f -;
done

helm install --values ./infra/helm/values.yaml --namespace=loki-stack --create-namespace loki grafana/loki-stack
