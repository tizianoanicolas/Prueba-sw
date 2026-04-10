package main

import "fmt" 
import "math" 

func main() {
	
	var lado1 float64
	var lado2 float64

	
	fmt.Print("Por favor, ingresa el primer lado: ") 
	fmt.Scanln(&lado1)                              
	

	
	fmt.Print("Ingresa el segundo lado: ")
	fmt.Scanln(&lado2)

	var area float64 = (lado1 * lado2) / 2
	var perimetro float64 = lado1 + lado2 + math.Sqrt(lado1*lado1+lado2*lado2)

	
	fmt.Printf("El área del triángulo es: %.2f\n", area)
	fmt.Printf("El perímetro del triángulo es: %.2f\n", perimetro)
}
