!#/bin/bash



# Generate a private key with no password
openssl genrsa -out private.pem 4096

# Extract the public key from the private key
openssl rsa -in private.pem -outform PEM -pubout -out public.pem