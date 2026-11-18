package service

import (
	"ToDo-App/internal/entity"
	"errors"
	"fmt"
	"os"
)

func CreateTask(tasks []entity.Task, task entity.Task) ([]entity.Task, error) {
	if task.ID <= 0 || task.Title == "" || task.Description == "" {
		return tasks, errors.New("task missing")
	}

	for i := range tasks {
		if task.ID == tasks[i].ID {
			return tasks, errors.New("task with ID already exists")
		}
	}
	tasks = append(tasks, task)
	return tasks, nil
}

func DeleteTask(tasks []entity.Task, id int) ([]entity.Task, error) {
	if id <= 0 {
		return tasks, errors.New("invalid id")
	}

	for i := range tasks {
		if tasks[i].ID == id {
			tasks = append(tasks[:i], tasks[i+1:]...)
			return tasks, nil
		}
	}
	return tasks, errors.New("task does not exist")
}

func ListTasks(tasks []entity.Task) error {

	if len(tasks) == 0 {
		return errors.New("list is empty")
	}

	for _, v := range tasks {
		fmt.Printf("ID: %v | TITLE: %v | DESCRIPTION: %v | STATUS: %v\n", v.ID, v.Title, v.Description, v.Status)
	}
	return nil
}

func FindTask(tasks []entity.Task, id int) (entity.Task, error) {
	for i := range tasks {
		if tasks[i].ID == id {
			return tasks[i], nil
		}
	}
	return entity.Task{}, errors.New("task not found")
}

func UpdateTask(tasks []entity.Task, id int, action string) ([]entity.Task, error) {
	if id <= 0 {
		return tasks, errors.New("invalid id")
	}

	for i := range tasks {
		if tasks[i].ID == id {
			switch action {
			case "1":
				var newTitle string

				fmt.Print("New Title: ")
				fmt.Scan(&newTitle)

				if newTitle == "" {
					return tasks, errors.New("invalid input")
				}

				tasks[i].Title = newTitle
				return tasks, nil

			case "2":
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

func CompleteTask(tasks []entity.Task, id int) (entity.Task, error) {
	if id <= 0 {
		return entity.Task{}, errors.New("invalid id")
	}

	for i := range tasks {
		if tasks[i].ID == id && !tasks[i].Status {
			tasks[i].Status = true
			return tasks[i], nil
		}
	}
	return entity.Task{}, errors.New("task not found")

}

func UncompleteTask(tasks []entity.Task, id int) (entity.Task, error) {
	if id <= 0 {
		return entity.Task{}, errors.New("invalid id")
	}

	for i := range tasks {
		if tasks[i].ID == id && tasks[i].Status {
			tasks[i].Status = false
			return tasks[i], nil
		}
	}
	return entity.Task{}, errors.New("task not found")

}

func FilterTasks(tasks []entity.Task, status bool) []entity.Task {
	filteredTasks := []entity.Task{}

	for i, _ := range tasks {
		if tasks[i].Status == status {
			filteredTasks = append(filteredTasks, tasks[i])
		}
	}
	return filteredTasks
}

func CountTasks(tasks []entity.Task) int {
	return len(tasks)
}

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
			tasks, err = CreateTask(tasks, t1)
			if err != nil {
				fmt.Println(err)
			}
			fmt.Println(tasks)

		case "2":
			fmt.Print("task which you want to delete: ")
			fmt.Scan(&input)

			tasks, err = DeleteTask(tasks, inputID)
			if err != nil {
				fmt.Println(err)
			}
			fmt.Println(tasks)

		case "3":
			var err error

			err = ListTasks(tasks)
			if err != nil {
				fmt.Println(err)
			}

		case "4":
			fmt.Print("task you want to find: ")
			fmt.Scan(&input)

			Task, err := FindTask(tasks, inputID)
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

			tasks, err = UpdateTask(tasks, inputID, action)
			if err != nil {
				fmt.Println(err)
			}
			fmt.Println(tasks)

		case "6":
			fmt.Print("task you want to complete: ")
			fmt.Scan(&input)

			Task, err := CompleteTask(tasks, inputID)
			if err != nil {
				fmt.Println(err)
			}
			fmt.Println(Task)

		case "7":
			fmt.Print("task you want to uncomplete: ")
			fmt.Scan(&input)

			Task, err := UncompleteTask(tasks, inputID)
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

			FilteredTasks := FilterTasks(tasks, operator)
			fmt.Println(FilteredTasks)

		case "9":
			res := CountTasks(tasks)
			fmt.Println(res)

		case "10":
			fmt.Println("Programm wird beendet...")
			os.Exit(0)

		default:
			fmt.Println("invalid input")

		}
	}

}
