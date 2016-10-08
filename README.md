[![Godoc Reference](https://godoc.org/github.com/aead/poly1305?status.svg)](https://godoc.org/github.com/aead/poly1305)

## The poly1305 message authentication code

Poly1305 is a cryptographic message authentication code (MAC) created by Daniel J. Bernstein. 
It can be used to verify the data integrity and the authenticity of a message and has been
standardized in [RFC 7539](https://tools.ietf.org/html/rfc7539 "RFC 7539").

This code is now stable and can be used in productive environments.
Backward compatibility is now guaranteed.

### Requirements
Following Go versions are supported:
 - 1.5.3
 - 1.5.4
 - 1.6.x
 - 1.7.x

Notice, that the code is only tested on amd64 and x86.

### Installation
Install in your GOPATH: `go get -u github.com/aead/poly1305`

### Performance
Benchmarks are run on a Intel i7-6500U (Sky Lake) on linux/amd64 with Go 1.7
```
On amd64: (The 'old' one is the implementation at golang.org/x/crypto/poly1305)

name           old time/op    new time/op    delta
64-4             96.5ns ± 1%    38.1ns ± 0%   -60.52%  (p=0.000 n=9+9)
1K-4              906ns ± 0%     399ns ± 0%   -55.98%  (p=0.000 n=9+9)
64Unaligned-4    96.1ns ± 0%    37.8ns ± 0%   -60.65%  (p=0.000 n=9+9)
1KUnaligned-4     910ns ± 1%     399ns ± 0%   -56.10%  (p=0.000 n=9+9)

name           old speed      new speed      delta
64-4            663MB/s ± 1%  1678MB/s ± 0%  +153.03%  (p=0.000 n=9+9)
1K-4           1.13GB/s ± 0%  2.56GB/s ± 0%  +126.89%  (p=0.000 n=9+9)
64Unaligned-4   666MB/s ± 0%  1693MB/s ± 0%  +154.19%  (p=0.000 n=9+9)
1KUnaligned-4  1.13GB/s ± 1%  2.56GB/s ± 0%  +127.44%  (p=0.000 n=9+9)
```
