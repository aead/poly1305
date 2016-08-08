[![Godoc Reference](https://godoc.org/github.com/aead/poly1305?status.svg)](https://godoc.org/github.com/aead/poly1305)

## The poly1305 message authentication code

Poly1305 is a cryptographic message authentication code (MAC) created by Daniel J. Bernstein. 
It can be used to verify the data integrity and the authenticity of a message and has been
standardized in [RFC 7539](https://tools.ietf.org/html/rfc7539 "RFC 7539").

### Requirements
Following Go versions are supported:
 - 1.5.3
 - 1.5.4
 - 1.6.x
 - 1.7 (currently rc5) 

Notice, that the code is only tested on amd64 and x86.

### Installation
Install in your GOPATH: `go get -u github.com/aead/chacha20`
