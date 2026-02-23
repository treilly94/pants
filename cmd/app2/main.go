package main

import (
	"fmt"
	"pants/pkg/package1"
)

func main() {
	tmp := package1.Add()
	fmt.Println(tmp)
	fmt.Println("app 2")
}
