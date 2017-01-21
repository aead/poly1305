[![Godoc Reference](https://godoc.org/github.com/aead/poly1305?status.svg)](https://godoc.org/github.com/aead/poly1305)

## The poly1305 message authentication code

Poly1305 is a cryptographic message authentication code (MAC) created by Daniel J. Bernstein. 
It can be used to verify the data integrity and the authenticity of a message and has been
standardized in [RFC 7539](https://tools.ietf.org/html/rfc7539 "RFC 7539").

### Requirements
Go version >= 1.5.3

### Installation
Install in your GOPATH: `go get -u github.com/aead/poly1305`

### Performance
The amd64 implementation of x/crypto/poly305 is based on this implementation - so
the performance should more or less be equal on amd64.

The reference (non-amd64) implementation is submitted to the go-team and (if it passes the review process)
will replace their ref. implementation.

The most significant performance improvement (compared to the /x/crypto/poly1305 impl.) can be observed,
if you compute the MAC over message chunks. For x/crypto/poly1305 you have to build a buffer, large enough
to hold the complete message. In some situations this leads to many memory allocations.
In these cases, the poly1305.Hash (impl. io.Writer) will lead to significant performance improvements.

Notice that, on arm machines the /x/crypto/poly1305 implementation may be faster, because
of an optimized assembly version.

amd64 (go1.7.4):
```
name                 speed
Sum_64-4             1.68GB/s ± 0%
SumUnaligned_64-4    1.68GB/s ± 0%
Sum_1K-4             2.50GB/s ± 0%
SumUnaligned_1K-4    2.50GB/s ± 0%
Write_64-4           1.96GB/s ± 0%
WriteUnaligned_64-4  1.96GB/s ± 0%
Write_1K-4           2.53GB/s ± 0%
WriteUnaligned_1K-4  2.53GB/s ± 0%
```

386 (go1.7.4):
```
name                 speed
Sum_64-2             165MB/s ± 0%
SumUnaligned_64-2    165MB/s ± 0%
Sum_1K-2             247MB/s ± 0%
SumUnaligned_1K-2    247MB/s ± 0%
Write_64-2           228MB/s ± 0%
WriteUnaligned_64-2  228MB/s ± 0%
Write_1K-2           256MB/s ± 0%
WriteUnaligned_1K-2  256MB/s ± 0%
```