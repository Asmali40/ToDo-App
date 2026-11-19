package main

import (
	"ToDo-App/internal/entity"
	"ToDo-App/internal/service"
	"fmt"
	"os"
)

func main() {
	var err error
	var input string
	var action string
	var operator bool
	var inputID int
	tasks := []entity.Task{}

	todoService := service.TodoService{}

	for {
		fmt.Println("---------------------------------------------------------")
		fmt.Println("------------------------TODO-APP-------------------------")
		fmt.Println("---------------------------------------------------------")
		fmt.Println()
		fmt.Println("---------------------------------------------------------")
		fmt.Println()
		fmt.Println("1) create task")
		fmt.Println("2) delete task")
		fmt.Println("3) list tasks")
		fmt.Println("4) find tasks")
		fmt.Println("5) update task")
		fmt.Println("6) complete task")
		fmt.Println("7) uncomplete task")
		fmt.Println("8) filter tasks")
		fmt.Println("9) count tasks")
		fmt.Println("10) quit")

		fmt.Scan(&input)

		switch input {
		case "1":
			var newId int
			var newTitle string
			var newDescription string

			fmt.Print("ID: ")
			fmt.Scan(&newId)
			fmt.Print("Title: ")
			fmt.Scan(&newTitle)
			fmt.Print("Description: ")
			fmt.Scan(&newDescription)

			t1 := entity.Task{
				ID:          newId,
				Title:       newTitle,
				Description: newDescription,
				Status:      false,
			}

			var err error
			err = todoService.CreateTask(t1)
			if err != nil {
				fmt.Println(err)
			}
			fmt.Println(tasks)

		case "2":
			fmt.Print("task which you want to delete: ")
			fmt.Scan(&input)

			err = todoService.DeleteTask(inputID)
			if err != nil {
				fmt.Println(err)
			}
			fmt.Println(tasks)

		case "3":
			var err error

			err = todoService.ListTasks()
			if err != nil {
				fmt.Println(err)
			}

		case "4":
			fmt.Print("task you want to find: ")
			fmt.Scan(&input)

			task, err := todoService.FindTask(inputID)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println(task)
			}

		case "5":
			fmt.Print("task you want to update: ")
			fmt.Scan(&input)

			fmt.Print("1) update title")
			fmt.Print("2) update description")
			fmt.Scan(&action)

			err = todoService.UpdateTask(inputID, action)
			if err != nil {
				fmt.Println(err)
			}
			fmt.Println(tasks)

		case "6":
			fmt.Print("task you want to complete: ")
			fmt.Scan(&input)

			err := todoService.CompleteTask(inputID)
			if err != nil {
				fmt.Println(err)
			}

		case "7":
			fmt.Print("task you want to uncomplete: ")
			fmt.Scan(&input)

			err := todoService.UncompleteTask(inputID)
			if err != nil {
				fmt.Println(err)
			}

		case "8":
			fmt.Print("1) show done tasks")
			fmt.Println("2) show open tasks")
			fmt.Scan(&input)

			if inputID == 1 {
				operator = true
			} else if inputID == 2 {
				operator = false
			} else {
				fmt.Println("invalid input")
			}

			FilteredTasks := todoService.FilterTasks(operator)
			fmt.Println(FilteredTasks)

		case "9":
			res := todoService.CountTasks()
			fmt.Println(res)

		case "10":
			fmt.Println("Programm wird beendet...")
			os.Exit(0)

		default:
			fmt.Println("invalid input")

		}
	}

}
