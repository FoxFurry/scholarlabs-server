#/bin/sh
ENVHASH=$(openssl dgst -sha256 -hex .env | awk '{print $2}')

for f in $(find . -name '*.yaml');
do
  ENVHASH=${ENVHASH} envsubst < $f | kubectl apply -n=scholarlabs -f -;
done
