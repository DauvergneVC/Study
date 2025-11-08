package main

import "time"

func main() {
	ch1 := make(chan string)
	ch2 := make(chan string)

	// Los time.sleep provocn que se envie antes al ch1 que al ch2
	go func() {
		time.Sleep(2 * time.Second)
		ch1 <- "1"
	}()

	go func() {
		time.Sleep(2 * time.Second)
		ch2 <- "2"
	}()

	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-ch1:
			println("Recibdo:", msg1)
		case msg2 := <-ch2:
			println("Recibdo:", msg2)
		// ESTE SIRVE PARA TERMINAR LA ejecucion desues de que pasa el timepo
		case <-time.After(500 * time.Millisecond):
			println("Timeout")
		}
	}
}
