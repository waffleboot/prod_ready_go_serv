.PHONY: run

run:
	GCUK_KEY_FILE=certs/key.file GCUK_CERT_FILE=certs/cert.file GCUK_SERVICE_ADDR=:8080 go run .
