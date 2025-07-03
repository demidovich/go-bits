package main

import (
	"fmt"
	"go-bits/msa/roundrobin"
)

func main() {
	servers := []string{
		"server1",
		"server2",
	}

	rr := roundrobin.New(servers)

	fmt.Println(rr.Next())
	fmt.Println(rr.Next())
	fmt.Println(rr.Next())
}
