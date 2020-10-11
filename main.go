package main

import (
	"fmt"
	"time"
)

func main() {
	i := 0
	for {
		time.Sleep(1 * time.Second)
		i++
		fmt.Printf("Program has completed %d cycles\n", i)
	}
}
