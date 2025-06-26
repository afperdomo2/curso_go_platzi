package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Simula una consulta a un servidor externo
func consultarServidor(nombre string, canal chan<- string) {
	fmt.Printf("🔍 Consultando %s...\n", nombre)
	// Simula un tiempo de respuesta aleatorio
	tiempo := time.Duration(rand.Intn(1000)+500) * time.Millisecond
	time.Sleep(tiempo)
	canal <- fmt.Sprintf("Respuesta de %s en %v", nombre, tiempo)
}

// Implementa un patrón de "carrera de consultas" donde se realizan peticiones
// simultáneas a dos servidores virtuales y se procesa solo la primera respuesta recibida.
//
// Este patrón es útil para mejorar la latencia en aplicaciones distribuidas,
// ya que permite obtener la respuesta del servidor más rápido sin esperar
// a que todos los servidores respondan.

func main() {
	// 1. Crea dos canales para recibir respuestas de cada servidor
	servidor1 := make(chan string)
	servidor2 := make(chan string)

	// 2. Inicia dos goroutines que simulan consultas a servidores distintos
	go consultarServidor("Servidor A", servidor1)
	go consultarServidor("Servidor B", servidor2)

	// 3. Utiliza una declaración select para manejar el primer resultado que llegue
	select {
	case respuesta := <-servidor1:
		fmt.Println("✅ Primer respuesta recibida:", respuesta)
	case respuesta := <-servidor2:
		fmt.Println("✅ Primer respuesta recibida:", respuesta)
	// 4. Incluye un timeout de 2 segundos para evitar bloqueos indefinidos
	case <-time.After(2 * time.Second):
		fmt.Println("⏰ Timeout: ningún servidor respondió a tiempo")
	}
}
