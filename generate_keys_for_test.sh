# shellcheck disable=SC1084
#!/bin/sh



# Generate a private key with no password
openssl genrsa -out private.pem 2048

# Extract the public key from the private key
openssl rsa -in private.pem -outform PEM -pubout -out public.pem