package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var t, a, b int
	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)

	_, _ = fmt.Fscanln(r, &t)
	for i := 0; i < t; i++ {
		_, _ = fmt.Fscanln(r, &a, &b)
		_, _ = fmt.Fprintln(w, a+b)
	}

	_ = w.Flush()
}
