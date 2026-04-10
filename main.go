package main

import "fmt" 
import "math/rand" 

func play() {
randomNumber := rand.Intn(101) 
var intentos int
var win int

fmt.Println("Adivina el número entre 0 y 100")
fmt.Println("Tienes 10 intentos para adivinar el número.")

for i := 0; i < 10; i++ {
	fmt.Print("Intento ", i+1, ": ")
	fmt.Scanln(&intentos)

	if intentos < randomNumber {
		fmt.Println("El número es mayor.")
	} else if intentos > randomNumber {
		fmt.Println("El número es menor.")
	} else if intentos == randomNumber {
		fmt.Println("¡Felicidades! Has adivinado el número.")
		win = 1
		return
	}
}
if win == 0 {
	fmt.Println("Lo siento, has agotado tus intentos. El número era: ", randomNumber)
}

}

func displayMenu() {
fmt.Println("=== MENU ===")
fmt.Println("1. Iniciar Juego")
fmt.Println("2. Salir")
fmt.Print("Elija una opción: ")
var choice int
fmt.Scanln(&choice)

for choice != 2 {
	switch choice {
	case 1:
		play()
		choice = 2
	default:
		fmt.Println("Opción no válida. Por favor, elija una opción válida.")
	}

}
}

func main() {
	displayMenu()

	
}
