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
			tasks, err = service.CreateTask(tasks, t1)
			if err != nil {
				fmt.Println(err)
			}
			fmt.Println(tasks)

		case "2":
			fmt.Print("task which you want to delete: ")
			fmt.Scan(&input)

			tasks, err = service.DeleteTask(tasks, inputID)
			if err != nil {
				fmt.Println(err)
			}
			fmt.Println(tasks)

		case "3":
			var err error

			err = service.ListTasks(tasks)
			if err != nil {
				fmt.Println(err)
			}

		case "4":
			fmt.Print("task you want to find: ")
			fmt.Scan(&input)

			Task, err := service.FindTask(tasks, inputID)
			if err != nil {
				fmt.Println(err)
			}
			fmt.Println(Task)

		case "5":
			fmt.Print("task you want to update: ")
			fmt.Scan(&input)

			fmt.Print("1) update title")
			fmt.Print("2) update description")
			fmt.Scan(&action)

			tasks, err = service.UpdateTask(tasks, inputID, action)
			if err != nil {
				fmt.Println(err)
			}
			fmt.Println(tasks)

		case "6":
			fmt.Print("task you want to complete: ")
			fmt.Scan(&input)

			Task, err := service.CompleteTask(tasks, inputID)
			if err != nil {
				fmt.Println(err)
			}
			fmt.Println(Task)

		case "7":
			fmt.Print("task you want to uncomplete: ")
			fmt.Scan(&input)

			Task, err := service.UncompleteTask(tasks, inputID)
			if err != nil {
				fmt.Println(err)
			}
			fmt.Println(Task)

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

			FilteredTasks := service.FilterTasks(tasks, operator)
			fmt.Println(FilteredTasks)

		case "9":
			res := service.CountTasks(tasks)
			fmt.Println(res)

		case "10":
			fmt.Println("Programm wird beendet...")
			os.Exit(0)

		default:
			fmt.Println("invalid input")

		}
	}

}
