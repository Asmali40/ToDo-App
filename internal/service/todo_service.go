package service

import (
	"ToDo-App/internal/entity"
	"errors"
	"fmt"
)

type TodoService struct {
	tasks []entity.Task
}

func (t *TodoService) CreateTask(task entity.Task) (error) {
	if task.ID <= 0 || task.Title == "" || task.Description == "" {
		return errors.New("task missing")
	}

	for i := range t.tasks {
		if task.ID == t.tasks[i].ID {
			return errors.New("task with ID already exists")
		}
	}
	t.tasks = append(t.tasks, task)
	return nil
}

func (t *TodoService) DeleteTask(tasks []entity.Task, id int) (error) {
	if id <= 0 {
		return errors.New("invalid id")
	}

	for i := range tasks {
		if tasks[i].ID == id {
			tasks = append(tasks[:i], tasks[i+1:]...)
			return nil
		}
	}
	return errors.New("task does not exist")
}

func (t *TodoService) ListTasks(tasks []entity.Task) error {

	if len(tasks) == 0 {
		return errors.New("list is empty")
	}

	for _, v := range tasks {
		fmt.Printf("ID: %v | TITLE: %v | DESCRIPTION: %v | STATUS: %v\n", v.ID, v.Title, v.Description, v.Status)
	}
	return nil
}

func (t *TodoService) FindTask(id int) (entity.Task, error) {
	for _, v := range t.tasks {
		if v.ID == id {
			return v, nil
		}
	}
	return entity.Task{}, errors.New("task not found")
}

func (t *TodoService) UpdateTask(id int, action string) error {
	if id <= 0 {
		return errors.New("invalid id")
	}

	for _, v := range t.tasks {
		if v.ID == id {
			switch action {
			case "1":
				var newTitle string

				fmt.Print("New Title: ")
				fmt.Scan(&newTitle)

				if newTitle == "" {
					return errors.New("invalid input")
				}

				v.Title = newTitle
				return nil

			case "2":
				var newDescription string

				fmt.Print("New Description: ")
				fmt.Scan(&newDescription)

				if newDescription == "" {
					return errors.New("invalid input")
				}

				v.Description = newDescription
				return nil

			default:
				return errors.New("invalid input")

			}
		}
	}
	return errors.New("task with id does not exist")
}

func (t *TodoService) CompleteTask(id int) error {
	if id <= 0 {
		return errors.New("invalid id")
	}

	for _, v := range t.tasks {
		if v.ID == id && !v.Status {
			v.Status = true
			return nil
		}
	}
	return errors.New("task not found")

}

func (t *TodoService) UncompleteTask(id int) error {
	if id <= 0 {
		return errors.New("invalid id")
	}

	for _, v := range t.tasks {
		if v.ID == id && v.Status {
			v.Status = false
			return nil
		}
	}
	return errors.New("task not found")

}

func (t *TodoService) FilterTasks(status bool) []entity.Task {
	filteredTasks := []entity.Task{}

	for _, v := range t.tasks {
		if v.Status == status {
			filteredTasks = append(filteredTasks, v)
		}
	}
	return filteredTasks
}

func (t *TodoService) CountTasks() int {
	return len(t.tasks)
}

func CreateToDoService() *TodoService {
	return &TodoService{}
}
