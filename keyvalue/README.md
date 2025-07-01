# KeyValue хранилище

Дженерики и сегментация на бакеты.

```go
package main

import (
    "fmt"
    "go-bits/keyvalue"
)

func main() {
    kv, _ := keyvalue.New[string](keyvalue.Config{
        BucketsCount: 100,
    })

    kv.Set("a", "1")

    v, ok := kv.Get("a")
    fmt.Printf("size=%d, value=%s, found=%t\n", kv.Size(), v, ok)

    kv.Forget("a")
    fmt.Printf("size=%d\n", kv.Size())
}

// size=1, value=1, found=true
// size=0
```
