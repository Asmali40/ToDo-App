package main

import (
	"errors"
	"fmt"
	"os"
)

type Task struct {
	ID          int
	Title       string
	Description string
	Status      bool
}

func CreateTask(tasks []Task, task Task) ([]Task, error) {
	if task.ID <= 0 || task.Title == "" || task.Description == "" {
		return []Task{}, errors.New("task missing")
	}

	for _, v := range tasks {
		if task.ID == v.ID {
			return []Task{}, errors.New("task with ID already exists")
		}
	}
	tasks = append(tasks, task)
	return tasks, nil
}

func DeleteTask(tasks []Task) {

}

func ListTasks(tasks []Task) error {

	if len(tasks) == 0 {
		return errors.New("list is empty")
	}

	for _, v := range tasks {
		fmt.Printf("ID: %v | TITLE: %v | DESCRIPTION: %v | STATUS: %v\n", v.ID, v.Title, v.Description, v.Status)
	}
	return nil
}

func FindTask(tasks []Task, id int) (Task, error) {
	for _, v := range tasks {
		if v.ID == id {
			return v, nil
		}
	}
	return Task{}, errors.New("task not found")
}

func UpdateTask(tasks []Task, id int) {

}

func CompleteTask(tasks []Task, id int) (Task, error) {
	if id <= 0 {
		return Task{}, errors.New("invalid id")
	}

	for _, v := range tasks {
		if v.ID == id && !v.Status {
			v.Status = true
			return v, nil
		}
	}
	return Task{}, errors.New("task not found")

}

func UncompleteTask(tasks []Task, id int) (Task, error) {
	if id <= 0 {
		return Task{}, errors.New("invalid id")
	}

	for _, v := range tasks {
		if v.ID == id && v.Status {
			v.Status = false
			return v, nil
		}
	}
	return Task{}, errors.New("task not found")

}

func FilterTasks(tasks []Task, status bool) []Task {
	filteredTasks := []Task{}

	for _, v := range tasks {
		if v.Status == status {
			filteredTasks = append(filteredTasks, v)
		}
	}
	return filteredTasks
}

func CountTasks(tasks []Task) int {
	return len(tasks)
}

func ClearCompleted(tasks []Task) {

}

func main() {
	tasks := []Task{}
	var input int

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
		fmt.Println("6) uncomplete task")
		fmt.Println("7) filter tasks")
		fmt.Println("8) count tasks")
		fmt.Println("9) clear completed tasks")
		fmt.Println("10) quit")

		var input string
		fmt.Scan(&input)

		switch input {
		case 1:
			var newId int
			var newTitle string
			var newDescription string

			fmt.Print("ID: ")
			fmt.Scan(&newId)
			fmt.Print("Title: ")
			fmt.Scan(&newTitle)
			fmt.Print("Description: ")
			fmt.Scan(&newDescription)

			t1 := Task{
				ID:          newId,
				Title:       newTitle,
				Description: newDescription,
				Status:      true,
			}

			CreateTask(tasks, t1)

		case 2:
			DeleteTask()
		case 3:
			ListTasks()
		case 4:
			FindTask()
		case 5:
			UpdateTask()
		case 6:
			CompleteTask()
		case 7:
			UncompleteTask()
		case 8:
			FilterTasks()
		case 9:
			CountTasks()
		case 10:
			ClearCompleted()
		case "Quit":
			os.Exit()
		default:
			fmt.Println("invalid input")

		}
	}

}
