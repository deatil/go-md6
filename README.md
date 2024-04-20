# go-md6

md6 Hash with cgo


### Env

CGO_ENABLED=1


### Download

```cmd
git clone github.com/deatil/go-md6
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


### LICENSE

*  The library LICENSE is `Apache2`, using the library need keep the LICENSE.


### Copyright

*  Copyright deatil(https://github.com/deatil).
