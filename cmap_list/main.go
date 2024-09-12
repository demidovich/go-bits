package main

import (
	"fmt"
	"structures/cmap_list/cmap"
)

func main() {
	cmap := cmap.NewCmap()
	cmap.Set("111", 11111)
	cmap.Set("222", 22222)
	cmap.Set("222", 33333)

	v, ok := cmap.Get("222")
	fmt.Println(v, ok)
}
