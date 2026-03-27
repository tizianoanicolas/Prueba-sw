package main

import "fmt" // Importamos el paquete "fmt" para funciones de formato (Print, Scan)

func main() {
	// Declaración de variables
	var nombre string
	var precio float64

	// --- 3. Usando Scanln para leer la entrada del usuario ---
	fmt.Print("Por favor, ingresa tu nombre: ") // Print sin salto de línea para que la entrada quede al lado
	fmt.Scanln(&nombre)                         // Lee la línea completa hasta el Enter y la guarda en 'nombre'
	                                            // El '&' es importante, indica la dirección de memoria de la variable

	// Pedir y leer un precio decimal
	fmt.Print("Ingresa un precio (ej. 19.234452): ")
	fmt.Scanln(&precio)

	// Imprimir nombre no usando nada
	fmt.Println("Nombre:", nombre)

	// Imprimir precio con formato
	fmt.Printf("Precio: %.2f\n", precio) // ".2f" formatea el float a 2 decimales
}