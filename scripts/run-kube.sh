#/bin/sh
CONFIGHASH=$(openssl dgst -sha256 -hex .env | awk '{print $2}')

for f in $(find . -name '*.yaml');
do
  CONFIGHASH=${CONFIGHASH} envsubst < $f | kubectl apply -n=scholarlabs -f -;
done
