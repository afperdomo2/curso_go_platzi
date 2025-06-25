package main

import (
	"fmt"
	"time"
)

func main() {
	// Crear dos canales
	canal1 := make(chan string)
	canal2 := make(chan string)

	// Goroutine para canal1
	go func() {
		for i := 1; i <= 3; i++ {
			canal1 <- fmt.Sprintf("Mensaje %d desde canal 1", i)
			time.Sleep(300 * time.Millisecond)
		}
		close(canal1)
	}()

	// Goroutine para canal2
	go func() {
		for i := 1; i <= 3; i++ {
			canal2 <- fmt.Sprintf("Mensaje %d desde canal 2", i)
			time.Sleep(500 * time.Millisecond)
		}
		close(canal2)
	}()

	// Recibir mensajes de ambos canales
	for canal1 != nil || canal2 != nil {
		select {

		case msg, ok := <-canal1: // Recibe mensajes del canal1
			// ok, indica si el canal está abierto o cerrado
			// Si el canal está cerrado, ok será false
			if ok {
				fmt.Println("📨 " + msg)
			} else {
				// Si el canal está cerrado, lo establecemos a nil
				// para que no se vuelva a recibir mensajes de él
				canal1 = nil
				fmt.Println("❌ Canal 1 cerrado")
			}

		case msg, ok := <-canal2:
			if ok {
				fmt.Println("📬 " + msg)
			} else {
				canal2 = nil
				fmt.Println("❌ Canal 2 cerrado")
			}
		}
	}

	fmt.Println("✅ Todos los mensajes han sido recibidos")
}
