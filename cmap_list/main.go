package main

import (
	"structures/cmap_list/cmap"
)

func main() {
	m := cmap.New()
	// m.Debug()

	m.Set("k1", 1)
	m.Set("k2", 2)
	m.Set("k3", 3)
	m.Set("k4", 4)
	m.Set("k5", 5)
	m.Debug()

	m.Set("k6", 6)
	m.Set("k7", 7)
	m.Set("k8", 8)
	m.Set("k9", 9)
	m.Set("k10", 10)
	m.Debug()

	// m.Set("k6", 6)
	// m.Set("k7", 7)
	// m.Set("k8", 8)
	// m.Debug()
}
