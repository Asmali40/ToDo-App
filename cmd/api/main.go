package main

import (
	"errors"
	"fmt"
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

func DeleteTask(tasks []Task, id int) ([]Task, error) {
	if id <= 0 {
		return []Task{}, errors.New("invalid id")
	}

	for i, v := range tasks {
		if v.ID == id {
			tasks = append(tasks[:i], tasks[i+1:]...)
			return tasks, nil
		}
	}
	return []Task{}, errors.New("task does not exist")
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

func UpdateTask(tasks []Task, id int, action string) ([]Task, error) {
	if id <= 0 {
		return tasks, errors.New("invalid id")
	}

	for i, v := range tasks {
		if v.ID == id {
			switch action {
			case "Title":
				var newTitle string

				fmt.Print("New Title: ")
				fmt.Scan(&newTitle)

				if newTitle == "" {
					return tasks, errors.New("invalid input")
				}

				tasks[i].Title = newTitle
				return tasks, nil

			case "Description":
				var newDescription string

				fmt.Print("New Description: ")
				fmt.Scan(&newDescription)

				if newDescription == "" {
					return tasks, errors.New("invalid input")
				}

				tasks[i].Description = newDescription
				return tasks, nil

			default:
				return tasks, errors.New("invalid input")

			}
		}
	}
	return tasks, errors.New("task with id does not exist")
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

func main() {
	tasks := []Task{}

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
		fmt.Println("9) quit")

		var input string
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

			t1 := Task{
				ID:          newId,
				Title:       newTitle,
				Description: newDescription,
				Status:      true,
			}

			res, err := CreateTask(tasks, t1)
			if err != nil {
				fmt.Println(err)
			}
			fmt.Println(res)

		case "2":

			var input int

			fmt.Print("task which you want to delete: ")
			fmt.Scan(&input)

			res, err := DeleteTask(tasks, input)
			if err != nil {
				fmt.Println(err)
			}
			fmt.Println(res)

		case "3":
			err := ListTasks(tasks)
			if err != nil {
				fmt.Println(err)
			}

		case "4":

			var input int

			fmt.Print("task you want to find: ")
			fmt.Scan(&input)

			res, err := FindTask(tasks, input)
			if err != nil {
				fmt.Println(err)
			}
			fmt.Println(res)

		case "5":
			var input int
			var action string

			fmt.Print("task you want to update: ")
			fmt.Scan(&input)

			fmt.Print("1) update title")
			fmt.Print("2) update description")
			fmt.Scan(&action)

			res, err := UpdateTask(tasks, input, action)
			if err != nil {
				fmt.Println(err)
			}
			fmt.Println(res)

		case "6":
			var input int

			fmt.Print("task you want to complete: ")
			fmt.Scan(&input)

			res, err := CompleteTask(tasks, input)
			if err != nil {
				fmt.Println(err)
			}
			fmt.Println(res)

		case "7":
			var input int

			fmt.Print("task you want to uncomplete: ")
			fmt.Scan(&input)

			res, err := UncompleteTask(tasks, input)
			if err != nil {
				fmt.Println(err)
			}
			fmt.Println(res)

		case "8":

			var input int
			var operator bool

			fmt.Print("1) show done tasks")
			fmt.Println("2) show open tasks")
			fmt.Scan(&input)

			if input == 1 {
				operator = true
			} else if input == 2 {
				operator = false
			} else {
				fmt.Println("invalid input")
			}

			res := FilterTasks(tasks, operator)
			fmt.Println(res)

		case "9":
			res := CountTasks(tasks)
			fmt.Println(res)

		case "Quit":
			fmt.Println("Programm wird beendet...")
			break

		default:
			fmt.Println("invalid input")

		}
	}

}
