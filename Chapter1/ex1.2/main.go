package main

import (
	"fmt"
	"os"
)

func main() {
	for i := 1; i < len(os.Args); i++ {
		fmt.Printf("Index : %d, value : %s \n", i, os.Args[i])
	}
}
