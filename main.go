package main

import (
	"fmt"

	v1 "gopkg.in/aristorinjuang/go-version.v1"
	v2 "gopkg.in/aristorinjuang/go-version.v2"
)

func main() {
	fmt.Println(v1.Version(), v2.Version())
}
