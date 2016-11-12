package main

import "fmt"

var currentId int
var todos Todos

func init() {
	CreateTodo(Todo{Name: "Tarefa 1"})
	CreateTodo(Todo{Name: "Tarefa 2"})
}

func FindTodo(id int) Todo {
	for _, todo := range todos {
		if todo.Id == id {
			return todo
		}
	}
	return Todo{}
}

func CreateTodo(todo Todo) Todo {
	currentId += 1
	todo.Id = currentId
	todos = append(todos, todo)
	return todo
}

func RemoveTodo(id int) error {
	for i, t := range todos {
		if t.Id == id {
			todos = append(todos[:i], todos[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("Could not find Todo with id of %d", id)
}
