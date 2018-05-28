package main

import (
	"fmt"
	"time"
)

//hacemos que pasen la mañoria de letras en un solo lugar
func main() {
	for i := 30; i < 254; i++ {
		fmt.Printf("%c", i)
		for i := 0; i < 20; i++ {
			time.Sleep(1000 * time.Nanosecond)
		}
		fmt.Print("\b")
	}

}
