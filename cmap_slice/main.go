package main

import (
	"fmt"
	"structures/cmap_slice/cmap"
)

func main() {
	cmap := cmap.NewCmap()
	cmap.Set("111", 11111)

	v, ok := cmap.Get("111")
	fmt.Println(v, ok)
}
