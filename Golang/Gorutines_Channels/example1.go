package main

import (
	"fmt"
	"time"
)

func main() {
	jobs := make(chan int, 100)
	results := make(chan int, 100)

	// crea 3 trabajadores que toman trabajos del canal jobs y envían resultados al canal results
	for w := 1; w <= 3; w++ {
		go Worker(w, jobs, results)
	}
	// envia 5 trabajos al canal jobs
	for j := 1; j <= 5; j++ {
		jobs <- j
	}
	// cierra el canal jobs para indicar que no hay más trabajos
	close(jobs)
	for a := 1; a <= 5; a++ {
		fmt.Println("Resultado", <-results)
	}
}

// recibe 2 canales_______ Lectura (<- antes)__ Escritura (chan<-)
func Worker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Printf("Trabajador %d empezo trabajo %d\n", id, j)
		time.Sleep(time.Second)
		fmt.Printf("Trabajador %d termino trabajo %d\n", id, j)
		results <- j * 2
	}
}
