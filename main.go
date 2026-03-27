package main

import "fmt"

func main() {
	var nombre string // Forma 1: declaramos la variable sin asignarle un valor
	nombre = "Ana" // Luego le asignamos un valor
	fmt.Println(nombre)

        nombre = "Juan" // Le reasignamos otro valor, son dinámicas, una vez definidas podemos cambiar su valor
        fmt.Println(nombre)

	var ciudad string = "Córdoba" // Forma 2: declaramos e inicializamos al mismo tiempo
	fmt.Println(ciudad)

	edad := 25 // Forma 3: declaración corta, solo dentro de funciones
	fmt.Println(edad)
}