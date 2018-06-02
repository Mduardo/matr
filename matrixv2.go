package main

import (
	"fmt"
	"math/rand"
	"time"
)

func aleatoriosemmill() int {
	start := time.Now()
	r := start.Nanosecond() / 10000000
	return r
}

func GenAleatorio(maximo int) int {
	for i := 0; i < 9; i++ {
		//esperamos 8 nanosegundos
		time.Sleep(1000 * time.Nanosecond)
	}
	//establecemos la semilla por la cual debe empezar
	rand.Seed(int64(aleatoriosemmill()))
	//establecemos el limite y retornamos el numero encontrado por Intn
	l := rand.Intn(maximo)
	return l
}

func NumberValidation(num int) bool {
	if num < 34 {
		return false
	}
	return true
}

func main() {
	val := false
	num := 0
	for i := 0; i < 1000; i++ {
		for l := 0; l < 20; l++ {
			num = GenAleatorio(220)
			val = NumberValidation(num)
			if val == false {
				num = GenAleatorio(220)
				val = NumberValidation(num)
				if val == false {
					num = GenAleatorio(220)

				}
			}
			fmt.Printf("%c \t", num)
		}
	}

}
