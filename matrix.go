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

/*func llamada(llam int) int {
	r := GenAleatorio(llam)
	return r
}*/

func main() {
	a, b, c, d, e, f, g, h, p := 0, 0, 0, 0, 0, 0, 0, 0, 0

	for i := 0; i < 20; i++ {

		for j := 0; j < 300; j++ {
			a = GenAleatorio(218)
			b = GenAleatorio(218)
			d = GenAleatorio(218)
			e = GenAleatorio(218)
			f = GenAleatorio(218)
			g = GenAleatorio(218)
			h = GenAleatorio(218)
			p = GenAleatorio(218)
			/**/ if a == 9 {
				a = GenAleatorio(5)
			}
			if c == 9 {
				c = GenAleatorio(5)
			}
			if b == 9 {
				b = GenAleatorio(5)
			}

			if d == 9 {
				d = GenAleatorio(5)
			}

			if e == 9 {
				e = GenAleatorio(5)
			}
			if f == 9 {
				f = GenAleatorio(5)
			}
			if g == 9 {
				g = GenAleatorio(5)
			}
			if h == 9 {
				h = GenAleatorio(5)
			}

			if p == 9 {
				p = GenAleatorio(5)
			} /**/

			fmt.Printf("%c%c\t%c%c\t%c\t%c%c\t%c%c", a, b, c, d, e, f, g, h, p)
			fmt.Printf("%c%c\t%c%c\t%c\t%c%c\t%c%c\n", e, c, p, h, a, f, b, g, d)

			//fmt.Printf("%c%c%c%c%c%c%c%c%c\n", a, b, c, d, e, f, g, h, p)

			//time.Sleep(1000 * time.Nanosecond)

			//fmt.Println()
		}
	}

}
