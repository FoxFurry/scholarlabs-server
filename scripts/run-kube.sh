#/bin/sh
GATEWAYCONFIGHASH=$(openssl dgst -sha256 -hex ./services/gateway/.env | awk '{print $2}')
USERCONFIGHASH=$(openssl dgst -sha256 -hex ./services/user/.env | awk '{print $2}')

for f in $(find . -name '*.yaml');
do
  GATEWAYCONFIGHASH=${GATEWAYCONFIGHASH} USERCONFIGHASH=${USERCONFIGHASH} envsubst < $f | kubectl apply -n=scholarlabs -f -;
done
