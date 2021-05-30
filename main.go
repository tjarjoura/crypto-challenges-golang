package main

import (
	"fmt"
	"os"
)

type challenge struct {
	fn     func(args ...string)
	n_args int
}

func chal1(hexstr string) {
}

func main() {
	challenges := []challenge{
			{fn: chal1, n_args: 1},
	}
}
