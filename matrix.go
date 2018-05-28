package main

import (
	"fmt"
	"math/rand"
	"time"
)

var consolapox [11]int //mover este

func posDisponiblesx(posx int) bool {
	p := posx

	if posx != 0 {
		posx = posx - 1

		if consolapox[posx] == 0 {
			consolapox[posx] = p
			return true
		} else {
			return false
		}
	} else {
		if consolapox[posx] == 0 {
			consolapox[posx] = p
			return true
		} else {
			return false
		}
	}

}

func moverenx(numeroaleatorio int) bool {
	disponibilidad := posDisponiblesx(numeroaleatorio)

	if disponibilidad != false {
		for i := 0; i < numeroaleatorio; i++ {

			//if i == numeroaleatorio-1 {
			//	fmt.Print(numeroaleatorio)
			//} else {
			fmt.Print(" ")
			//}
		}
		return true
	} else {
		return false
	}

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
func main() {
	var m int
	var d bool
	for i := 0; i < 10; i++ { //mover este
		m = numeroaletorio(10) //mover este
		d = moverenx(m)

		if d != false {
			fmt.Print("o")
			if i == 9 { //mover este
				for i := 0; i < len(consolapox)-1; i++ {

					consolapox[i] = 0
				}

			}
		}

	}

}
