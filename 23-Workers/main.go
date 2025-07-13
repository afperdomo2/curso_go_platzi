package main

import (
	"fmt"
	"time"
)

// ! workers
// Este ejemplo muestra cómo crear un worker que procesa tareas de un canal.
// Los workers son goroutines que realizan tareas específicas y pueden ser utilizados
// para procesar datos en paralelo, mejorando la eficiencia y el rendimiento de las
// aplicaciones en Go.

func worker(id int, tasksChannel <-chan int, resultsChannel chan<- int) {
	for task := range tasksChannel {
		fmt.Printf("👷 Worker %d procesando tarea %d...\n", id, task)
		time.Sleep(3 * time.Second) // Simula un trabajo que toma tiempo
		fmt.Printf("✅ Worker %d completó tarea %d\n", id, task)
		// Envía el resultado al canal de resultados
		resultsChannel <- task * 2 // Envía resultado
	}
}

func main() {
	fmt.Println("🚀 Iniciando worker...")
	const numWorkers = 3

	tasksChannel := make(chan int, numWorkers)   // Canal para tareas
	resultsChannel := make(chan int, numWorkers) // Canal para resultados

	// Enviamos tareas al canal
	for i := range numWorkers {
		go worker(i, tasksChannel, resultsChannel) // Inicia workers
	}

	// Enviamos tareas al canal de tareas
	fmt.Println("📥 Enviando tareas a los workers...")
	for i := range numWorkers {
		tasksChannel <- i // Envía tarea al canal
		fmt.Printf("📤 Tarea %d enviada al worker\n", i)
	}
	close(tasksChannel) // Cierra el canal de tareas

	// Recibimos resultados de los workers
	fmt.Println("📥 Esperando resultados de los workers...")
	for range numWorkers {
		result := <-resultsChannel // Recibe resultado del canal
		fmt.Printf("📤 Resultado recibido: %d\n", result)
	}
	close(resultsChannel) // Cierra el canal de resultados
}
