package main

import (
	"fmt"
	"math/rand"
	"time"
)

func moverenx(numeroaleatorio int) {
	//disponibilidad := posDisponiblesx(numeroaleatorio)

	//	if disponibilidad != false {
	for i := 0; i < numeroaleatorio; i++ {

		//if i == numeroaleatorio-1 {
		//	fmt.Print(numeroaleatorio)
		//} else {
		fmt.Print(" ")
		//}
		if i == numeroaleatorio-1 {
			for l := 30; l < 254; l++ {
				fmt.Printf("%c", l)
				for m := 0; m < 20; m++ {
					//esperamos un timepo para poder ver la letra
					time.Sleep(1000 * time.Nanosecond)
				}
				fmt.Print("\b")
			}
		}
	}
	/*return true
	} else {
		return false
	}*/

}

//creamos la semilla aleatoria para rand con el tiempo en decenas
func aleatoriosemmill() int {
	start := time.Now()
	r := start.Nanosecond() / 10000000
	return r
}

//funcion para generar numero aleatorio
func numeroaletorio(maximo int) int {
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

//hacemos que pasen la mañoria de letras en un solo lugar
func main() {

	moverenx(10)

	for i := 30; i < 254; i++ {
		fmt.Printf("%c", i)
		for i := 0; i < 20; i++ {
			//esperamos un timepo para poder ver la letra
			time.Sleep(1000 * time.Nanosecond)
		}
		fmt.Print("\b")
	}
	r, j := 0, 1
	for j != 0 {
		r = numeroaletorio(73)
		moverenx(r)
	}

}
