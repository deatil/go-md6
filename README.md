# go-md6

md6 摘要, 使用 cgo 方式实现


### 环境要求

CGO_ENABLED=1


### 使用方法

1. 下载

```cmd
git clone github.com/deatil/go-md6
```

2. 开始使用

```go
import "github.com/deatil/go-md6"

func main() {
    data := []byte("...")

    h := md6.New256()
    h.Write(data)
    sum := h.Sum(nil)
}
```


### 开源协议

*  `go-md6` 遵循 `Apache2` 开源协议发布，开源版在保留本软件版权的情况下可免费使用。


### 版权

*  该系统所属版权归 deatil(https://github.com/deatil) 所有。
