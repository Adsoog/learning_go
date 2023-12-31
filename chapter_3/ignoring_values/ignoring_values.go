package main

import "fmt"

func main() {

	gato1, _ := michis() // Haber aqui algunos errores, si en la funcion michis pides a michi1 y michi2 tiene que pasarlos cuando la llames, para solucionar esto simplemente borra michi1 y michi 2 y deja que la funcion retorne los michis
	fmt.Print(gato1)

}

func michis() (string, string) {
	return "atutui", "pansudini"
}

func getPoint(x int, y int) (int, int) {
	return 3, 4
}
