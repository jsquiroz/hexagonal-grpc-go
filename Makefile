# This file will create all files related with grpc

# Version APP, You will need to change it when a new
# relase has been approved
CERTIFICATES_PATH=certs
PATH_OUT=pkg/infrastructure/delivery/grpc/proto/**
protoc:
	rm -f ${PATH_OUT}/*.go
	./proto.sh

cert:
	rm -f ${CERTIFICATES_PATH}/{*.cert,*.key,*.pem,*.srl,*.csr}
	openssl genrsa -out ${CERTIFICATES_PATH}/ca.key 4096
	openssl req -new -x509 -key ${CERTIFICATES_PATH}/ca.key -sha256 -subj "/C=MX/ST=CDMX/O=RR Corporativo." -days 365 -out ${CERTIFICATES_PATH}/ca.cert
	openssl genrsa -out ${CERTIFICATES_PATH}/service.key 4096
	openssl req -new -key ${CERTIFICATES_PATH}/service.key -out ${CERTIFICATES_PATH}/service.csr -config ${CERTIFICATES_PATH}/certificate.conf
	openssl x509 -req -in ${CERTIFICATES_PATH}/service.csr -CA ${CERTIFICATES_PATH}/ca.cert -CAkey ${CERTIFICATES_PATH}/ca.key -CAcreateserial \
		-out ${CERTIFICATES_PATH}/service.pem -days 365 -sha256 -extfile ${CERTIFICATES_PATH}/certificate.conf -extensions req_ext
