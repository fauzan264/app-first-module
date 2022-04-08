package main

import (
	"fmt"

	go_first_module "github.com/fauzan264/go-first-module"
)

func main() {
	for i := 0; i < 10; i++ {
		fmt.Println(go_first_module.First(), " ", i)
	}
}
