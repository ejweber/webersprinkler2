# I am experimenting with certificate-based client authentication using the following resources:
# https://dev.to/techschoolguru/how-to-secure-grpc-connection-with-ssl-tls-in-go-4ph
# https://dev.to/techschoolguru/how-to-create-sign-ssl-tls-certificates-2aai

# Generate CA's private key and self-signed certificate.
#   openssl req is generally used to create CSRs (certificate signing requests).
#   -x509 - Output a self-signed certificate instead of a CSR.
#   -newkey - Create a new private key instead of using an existing one.
#   -days - The number of days the certificate is valid for.
#   -nodes - Do not encrypt the private key.
#   -keyout - Write the private key to a file.
#   -out - Write the certificate to a file.
openssl req -x509 -newkey rsa:4096 -days 365 -nodes -keyout ca-key.pem -out ca-cert.pem \
    -subj "/C=US/ST=Kansas/L=Wichita/O=WeberSprinkler/OU=/CN=/emailAddress=webericjames@gmail.com"

# Echo CA's self-signed certificate for inspection if desired.
#   openssl x509 is a multipurpose utility used here to inspect a certificate.
#   -in - The filename for the certificate to inspect.
#   -noout - Don't output an encoded version of the certificate.
#   -text - Print the certificate in text form.
# openssl x509 -in ca-cert.pem -noout -text

# Generate server's private key and CSR.
#   openssl req is generally used to create CSRs (certificate signing requests).
#   -newkey - Create a new private key instead of using an existing one.
#   -nodes - Do not encrypt the private key.
#   -keyout - Write the private key to a file.
#   -out - Write the certificate to a file.
openssl req -newkey rsa:4096 -nodes -keyout server-key.pem -out server-req.pem \
    -subj "/C=US/ST=Kansas/L=Wichita/O=WeberSprinkler/OU=/CN=/emailAddress=webericjames@gmail.com"

# Use CA's private key to sign server's CSR and get back the signed certificate.
#   openssl x509 is a multipurpose utility used here to sign a certificate.
#   -in - The filename for the certificate to sign.
#   -days - The number of days the certificate is valid for.
#   -CA - The certificate to use for signing.
#   -CAkey - The file containing the CA's private key.
#   -CAcreateserial - Randomly choose a serial number to start with and store it in a file to be incremented.
#   - out - Write the certificate to a file.
openssl x509 -req -in server-req.pem -days 60 -CA ca-cert.pem -CAkey ca-key.pem -out server-cert.pem

# Echo CA's self-signed certificate for inspection if desired.
#   openssl x509 is a multipurpose utility used here to inspect a certificate.
#   -in - The filename for the certificate to inspect.
#   -noout - Don't output an encoded version of the certificate.
#   -text - Print the certificate in text form.
openssl x509 -in server-cert.pem -noout -text

# -----
# TODO: Work through the following commands for client certificate generation.

# Generate client's private key and CSR.
openssl req -newkey rsa:4096 -nodes -keyout client-key.pem -out client-req.pem -subj ""

# 5. Use CA's private key to sign client's CSR and get back the signed certificate
# openssl x509 -req -in client-req.pem -days 60 -CA ca-cert.pem -CAkey ca-key.pem -CAcreateserial -out client-cert.pem -extfile client-ext.cnf
 
echo "Client's signed certificate"
# openssl x509 -in client-cert.pem -noout -text