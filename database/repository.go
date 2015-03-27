package database

import (
	"fmt"
	"github.com/alexdevid/go-hello-rest/models"
)

var currentId int

var TodoList models.Todos

// Give us some seed data
func init() {
	Create(models.Todo{Name: "Write presentation"})
	Create(models.Todo{Name: "Host meetup"})
}

func Find(id int) models.Todo {
	for _, t := range TodoList {
		if t.Id == id {
			return t
		}
	}
	// return empty Todo if not found
	return models.Todo{}
}

//this is bad, I don't think it passes race condtions
func Create(t models.Todo) models.Todo {
	currentId += 1
	t.Id = currentId
	TodoList = append(TodoList, t)
	return t
}

func Destroy(id int) error {
	for i, t := range TodoList {
		if t.Id == id {
			TodoList = append(TodoList[:i], TodoList[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("Could not find Todo with id of %d to delete", id)
}
