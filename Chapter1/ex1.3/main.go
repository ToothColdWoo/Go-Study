package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	startTime := time.Now()
	/* var s, sep string
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	println(s) */
	fmt.Println(strings.Join(os.Args[1:], " "))
	elapsedTime := time.Since(startTime)
	fmt.Printf("Elapsed Time : %s", elapsedTime)
}
