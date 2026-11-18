package service

import (
	"ToDo-App/internal/entity"
	"errors"
	"fmt"
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
