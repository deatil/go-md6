# go-md6

md6 Hash with cgo


### Env

CGO_ENABLED=1


### Download

```cmd
go get -u github.com/deatil/go-md6
```


### Get Starting

```go
import "github.com/deatil/go-md6"

func main() {
    data := []byte("...")

    h := md6.New256()
    h.Write(data)
    sum := h.Sum(nil)
}
```


### New Functions

md6 have `New` functions: `New(d int) (hash.Hash, error)`,
`New224() hash.Hash`,
`New256() hash.Hash`,
`New384() hash.Hash`,
`New512() hash.Hash`


### Sum Functions

md6 have `Sum` functions: `Sum(d int, data []byte) (out []byte, err error)`,
`Sum224(data []byte) (out [Size224]byte)`,
`Sum256(data []byte) (out [Size256]byte)`,
`Sum384(data []byte) (out [Size384]byte)`,
`Sum512(data []byte) (out [Size512]byte)`


### LICENSE

*  The library LICENSE is `Apache2`, using the library need keep the LICENSE.


### Copyright

*  Copyright deatil(https://github.com/deatil).
