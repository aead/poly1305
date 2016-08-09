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

### Performance
Benchmarks are run on a Intel i7-6500U (Sky Lake) on linux/amd64 with Go 1.6.3
```
On amd64:
BenchmarkSum_8-4               	334.30 MB/s
BenchmarkSumUnaligned_8-4      	334.61 MB/s
BenchmarkSum_4K-4              2507.44 MB/s
BenchmarkSumUnaligned_4K-4     2506.90 MB/s
BenchmarkWrite_8-4              595.91 MB/s
BenchmarkWriteUnaligned_8-4     598.25 MB/s
BenchmarkWrite_4K-4            2493.24 MB/s
BenchmarkWriteUnaligned_4K-4   2511.37 MB/s
```
