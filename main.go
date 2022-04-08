package main

import (
	"fmt"

	go_first_module "github.com/fauzan264/go-first-module/v2"
)

func main() {
	for i := 0; i < 10; i++ {
		fmt.Println(go_first_module.First(" Fauzan"), " ", i)
	}
}
