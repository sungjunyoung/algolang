package main

import (
	"fmt"
)

func main() {
	var a, b int
	_, _ = fmt.Scan(&a)
	_, _ = fmt.Scan(&b)

	fmt.Println(a + b)
}
