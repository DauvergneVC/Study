// Envio y recepcion de valores, para uso concurrente.

package main

import "fmt"

func main() {
	// Asi declaro un canal, con el tipo de dato que va a transportar, string en este caso.
	// ESTE NO TIENE BUFFER, ya que la goroutine se bloquea hasta que el mensaje sea recibido
	messages := make(chan string)
	go func() {
		// Esto envia un string al canal messages
		messages <- "Hola desde channel"
	}()
	msg := <-messages
	fmt.Println(msg)

	// CANAL CON BUFFER
	// Al pasarle el 2, le digo que cree un canal con buffer, espacio para 2 valores
	messages2 := make(chan string, 2)
	messages2 <- "Hola desde channel 1"
	messages2 <- "Hola desde channel 2"
	fmt.Println(<-messages2)
	fmt.Println(<-messages2)
}
