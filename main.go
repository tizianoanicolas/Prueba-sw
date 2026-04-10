package main

import "fmt"


type Tarea struct {
	Titulo      string
	Descripcion string
	Completada  bool
}

func main() {
	var listaTareas []Tarea
	var opcion int

	for {
		fmt.Println("\nSeleccione una opción:")
		fmt.Println("1. Agregar tarea")
		fmt.Println("2. Marcar tarea como completada")
		fmt.Println("3. Editar Tarea")
		fmt.Println("4. Eliminar Tarea")
		fmt.Println("5. Salir")
		fmt.Print("Ingrese la opción: ")
		fmt.Scanln(&opcion)

		if opcion == 5 {
			fmt.Println("Saliendo...")
			break
		}

		switch opcion {
		case 1:
			var t Tarea
			fmt.Println("Ingrese nombre de la tarea:")
			// Truco con Scanf para permitir espacios con solo fmt
			fmt.Scanf("%s\n", &t.Titulo) 
			fmt.Println("Ingrese descripción de la tarea:")
			fmt.Scanf("%s\n", &t.Descripcion)
			t.Completada = false
			listaTareas = append(listaTareas, t)
			fmt.Println("Tarea agregada correctamente")

		case 2:
			var idx int
			fmt.Print("Ingrese el índice de la tarea: ")
			fmt.Scanln(&idx)
			if idx >= 0 && idx < len(listaTareas) {
				listaTareas[idx].Completada = true
				fmt.Println("¡Tarea completada!")
			} else {
				fmt.Println("Índice inválido.")
			}

		case 3:
			var idx int
			fmt.Print("Ingrese el índice a editar: ")
			fmt.Scanln(&idx)
			if idx >= 0 && idx < len(listaTareas) {
				fmt.Print("Nuevo título: ")
				fmt.Scanf("%s\n", &listaTareas[idx].Titulo)
				fmt.Print("Nueva descripción: ")
				fmt.Scanf("%s\n", &listaTareas[idx].Descripcion)
				fmt.Println("Tarea actualizada.")
			}

		case 4:
			var idx int
			fmt.Print("Ingrese el índice a eliminar: ")
			fmt.Scanln(&idx)
			if idx >= 0 && idx < len(listaTareas) {
				// Eliminación usando re-slicing
				listaTareas = append(listaTareas[:idx], listaTareas[idx+1:]...)
				fmt.Println("Tarea eliminada.")
			}

		default:
			fmt.Println("Opción no reconocida.")
		}

		// Mostrar la lista después de cada operación
		fmt.Println("\nLista de tareas:")
		fmt.Println("===================================")
		for i, t := range listaTareas {
			estado := "Pendiente"
			if t.Completada {
				estado = "Hecho"
			}
			fmt.Printf("%d. %s\n - %s\n - Estado: [%s]\n", i, t.Titulo, t.Descripcion, estado)
			fmt.Println("-----------------------------------")
		}
		if len(listaTareas) == 0 {
			fmt.Println("No hay tareas pendientes.")
		}
		fmt.Println("===================================")
	}
}