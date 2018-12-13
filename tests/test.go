package main

import (
	"fmt"
	"strings"

	//"time"
)

func main() {
	a:="a234d"
	b:="b233s"
	if strings.Compare(a,b) == -1 {
		fmt.Println("a < b")
	} else {
		fmt.Println("a >= b")
	}
}
