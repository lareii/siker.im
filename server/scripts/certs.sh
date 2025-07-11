#!/bin/bash
set -euo pipefail

ORG="siker.im"
INTERNAL_IP="127.0.0.1"
EXTERNAL_IP="x.x.x.x"
DNS_NAME="mongodb"
COUNTRY="TR"
STATE="Istanbul"
CITY="Istanbul"
VALID_DAYS=99999
CA_NAME="siker.imCA"

CERT_DIR=".certs"
mkdir -p "$CERT_DIR"
cd "$CERT_DIR"

# private ca
openssl req -x509 -new -nodes \
  -days "$VALID_DAYS" \
  -subj "/C=$COUNTRY/ST=$STATE/L=$CITY/O=$ORG/CN=$CA_NAME" \
  -keyout ca.key -out ca.crt

# openssl config
cat > server.cnf <<EOF
[ req ]
distinguished_name = dn
prompt = no
default_md = sha256
req_extensions = v3_req

[ dn ]
C = $COUNTRY
ST = $STATE
L = $CITY
O = $ORG
CN = $DNS_NAME

[ v3_req ]
subjectAltName = @alt_names
keyUsage = critical, digitalSignature, keyEncipherment
extendedKeyUsage = serverAuth

[ alt_names ]
DNS.1 = $DNS_NAME
IP.1 = $IP
IP.2 = $EXTERNAL_IP
EOF

# server key and csr
openssl req -new -nodes -newkey rsa:4096 \
  -keyout server.key -out server.csr \
  -config server.cnf

# server certificate signed by CA
openssl x509 -req -in server.csr -CA ca.crt -CAkey ca.key -CAcreateserial \
  -out server.crt -days "$VALID_DAYS" -extensions v3_req -extfile server.cnf

cat server.key server.crt > mongo.pem
chmod 644 mongo.pem

# client certificate config
cat > client.cnf <<EOF
[ req ]
distinguished_name = dn
prompt = no
default_md = sha256
req_extensions = v3_req

[ dn ]
C = $COUNTRY
ST = $STATE
L = $CITY
O = $ORG
CN = mongo-client

[ v3_req ]
keyUsage = critical, digitalSignature, keyEncipherment
extendedKeyUsage = clientAuth
EOF

# client key and csr
openssl req -new -nodes -newkey rsa:4096 \
  -keyout client.key -out client.csr \
  -config client.cnf

# client certificate signed by CA
openssl x509 -req -in client.csr -CA ca.crt -CAkey ca.key -CAcreateserial \
  -out client.crt -days "$VALID_DAYS" -extensions v3_req -extfile client.cnf

cat client.key client.crt > mongo-client.pem
chmod 400 mongo-client.pem

# list generated files
ls -l mongo*.pem ca.crt
