package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type task struct {
	name        string
	description string
	complete    bool
}

func main() {
	tasks := make(map[int]task)
	lastID := 0

	for {
		fmt.Println("Selecciuone la opción deseada:")
		fmt.Println("1. Ver tareas")
		fmt.Println("2. Agregar tarea")
		fmt.Println("3. Marcar tarea como completada")
		fmt.Println("4. Elimintar tarea")
		fmt.Println("5. Salir")

		fmt.Print("Opción: ")

		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		choice, _ := strconv.Atoi(scanner.Text())

		switch choice {
		case 1:
			if len(tasks) == 0 {
				fmt.Println("-----------------------------")
				fmt.Println("No hay tareas pendientes :3 ")
				fmt.Println("-----------------------------")
			} else {
				for id, task := range tasks {
					fmt.Printf("%d. %s: %s (Completado: %t)\n", id, task.name, task.description, task.complete)
				}
			}
		case 2:
			fmt.Print("Nombre de la tarea?: ")
			scanner.Scan()
			name := scanner.Text()
			fmt.Print("Descripción de la tarea: ")
			scanner.Scan()
			description := scanner.Text()

			lastID++

			tasks[lastID] = task{name: name, description: description, complete: false}
			fmt.Println("-----------------------------")
			fmt.Printf("Tarea agregada con el ID: %d\n", lastID)
			fmt.Println("-----------------------------")
		case 3:
			fmt.Println("-----------------------------")
			fmt.Print("Inserte el ID de la tarea que quiera completar: ")

			scanner.Scan()
			id, _ := strconv.Atoi(scanner.Text())
			fmt.Println("-----------------------------")

			if task, ok := tasks[id]; ok {
				task.complete = true
				tasks[id] = task
				fmt.Println("-----------------------------")
				fmt.Printf("Tarea con el id: %d, fue completada\n", id)
				fmt.Println("-----------------------------")

			} else {
				fmt.Println("-----------------------------")
				fmt.Println("El ID no existe oe sonso :v")
				fmt.Println("-----------------------------")
			}
		case 4:
			fmt.Println("-----------------------------")
			fmt.Print("Inserte el ID de la tarea que quiera eliminar: ")

			scanner.Scan()
			id, _ := strconv.Atoi(scanner.Text())
			fmt.Println("-----------------------------")

			if _, ok := tasks[id]; ok {
				delete(tasks, id)
				fmt.Println("-----------------------------")
				fmt.Printf("La tarea con ID: %d ha sido eliminada\n", id)
				fmt.Println("-----------------------------")

			} else {
				fmt.Println("-----------------------------")
				fmt.Println("El ID no existe oe sonso :v")
				fmt.Println("-----------------------------")
			}
		case 5:
			fmt.Println("-----------------------------")
			fmt.Println("Saliendo papi :D")
			os.Exit(0)
		}
	}
}
