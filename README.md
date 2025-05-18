# Curso de Go (Golang) - Platzi

Este repositorio contiene los ejercicios y proyectos desarrollados durante el curso de Go de Platzi.

## Estructura del Proyecto

- **HolaMundo/**: Ejemplo básico de un programa "Hola Mundo" en Go.
  - `main.go`: Archivo fuente con el código del programa.
  - `main.exe`: Archivo ejecutable compilado.
- **Paquetes/**: Ejemplos de manejo de paquetes y variables de entorno.
  - `main.go`: Acceso a variables de entorno multiplataforma.

## Programas Incluidos

### Hola Mundo

Un simple programa que muestra "Hola Mundo" en la consola.

```go
package main

import "fmt"

func main() {
 fmt.Println("Hola Mundo")
}
```

Para ejecutar este programa, navegue al directorio HolaMundo y ejecute:

```bash
go run main.go
```

O simplemente ejecute el archivo binario compilado:

```bash
./main.exe
```

### Variables de Entorno

Un programa que muestra cómo acceder a variables de entorno de manera multiplataforma (Windows/Linux).

```go
package main

import (
	"fmt"
	"os"
)

func main() {
	// En Windows, la variable HOME no existe por defecto
	// Se debe usar USERPROFILE en su lugar
	homeDir := os.Getenv("HOME")
	
	// Si HOME está vacío (Windows), intentamos con USERPROFILE
	if homeDir == "" {
		homeDir = os.Getenv("USERPROFILE")
		if homeDir == "" {
			fmt.Println("No se encontró la variable de entorno HOME ni USERPROFILE")
			return
		}
		fmt.Printf("La variable de entorno USERPROFILE es: %s\n", homeDir)
	} else {
		fmt.Printf("La variable de entorno HOME es: %s\n", homeDir)
	}
}
```

Para ejecutar este programa, navegue al directorio Paquetes y ejecute:

```bash
go run main.go
```

## Requisitos

- Go versión 1.15 o superior
- Editor de texto o IDE (como GoLand, Visual Studio Code con la extensión Go)

## Recursos Adicionales

- [Documentación oficial de Go](https://golang.org/doc/)
- [Tour de Go](https://tour.golang.org/welcome/1)
- [Curso de Go en Platzi](https://platzi.com/cursos/go-golang/)

---

Última actualización: 18 de mayo de 2025
