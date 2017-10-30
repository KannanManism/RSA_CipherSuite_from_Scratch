This repository consists of 3 files, namely

1. rsa-keygen.go -> Would generate the private and public key components as per RSA system. All algorithms have been implemented from
scratch (Extended Euclidean Algorithm, Miller-Rabin Test for primality and Square and Multiply algorithm for modular exponentiation)

2. rsa-encrypt.go -> Performs RSA encryption given the public key file (N,e) and message in decimal format.

3. rsa-decrypt.go -> Performs RSA decryption given the private key file (N,d) and ciphertext in decimal format.

Build Details :

1. go build rsa-keygen.go
2. go build rsa-encrypt.go
3. go build rsa-decrypt.go


Usage :

1. ./rsa-keygen <publickeyFileName> <privateKeyFileName>
2. ./rsa-encrypt <publickeyFileName> <message in decimal format>
3. ./rsa-decrypt <privateKeyFileName> <ciphertext in decimal format>


License :

Code belongs to Venkatesh Gopal (vgopal3@jhu.edu/vnktshgopalg@gmail.com). For modifications to the source code, please reach out to this email address.
