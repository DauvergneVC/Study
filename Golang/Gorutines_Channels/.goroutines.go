// This is for this video : https://www.youtube.com/watch?v=tEBufGSmzAY&list=PL6Tl4DQ96MPOp8f_3ujLHU-n_ziRnhBCI&index=9&t=11s

package main

import (
	"fmt"
	"sync"
)

func main() {
	// // Con go indico que usare una goroutine
	// // Si no doy tiempo suficiente para que se ejecute la goroutine, el programa termina antes de que se imprima el mensaje
	// go PrintMessage("Hola desde goroutine!")
	// time.Sleep(time.Second) // Para dar tiempo
	// fmt.Println("Hola desde la funcion main")

	fmt.Println("_________________________________________")

	// Otra forma de hacerlo es con un wait group.
	var wg sync.WaitGroup
	wg.Add(1) // Añade una goroutine al wait group
	go PrintMessage(&wg, "Hola desde goroutine con waitgroup")
	wg.Wait() // PAra la ejecucion hasta terminar todas las goroutines
	fmt.Println("Hola desde la funcion main")
}

func PrintMessage(wg *sync.WaitGroup, message string) {
	defer wg.Done() // Indica que siempre debe ejecutarce al final
	fmt.Println(message)
}
