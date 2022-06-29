package main

import (
	"github.com/gofiber/fiber/v2"
)

func CreateTodo(c *fiber.Ctx) error {
	var todos Todo
	if err := c.BodyParser(&todos); err != nil {
		return err
	}
	db.Create(&todos)
	return c.JSON(todos)
}

func GetTodo(c *fiber.Ctx) error {
	var todos []Todo
	db.Find(&todos)
	return c.JSON(todos)
}

func UpdateTodo(c *fiber.Ctx) error {
	var todos Todo
	todoId := c.Params("id")

	db.Find(&todos, "id = ?", todoId)

	var UpdateTodo updateTodo

	if err := c.BodyParser(&UpdateTodo); err != nil {
		return err
	}

	todos.Task = UpdateTodo.Task
	todos.Completed = UpdateTodo.Completed

	db.Save(&todos)

	return c.JSON(todos)
}

func DeleteTodo(c *fiber.Ctx) error {
	var todos Todo

	id := c.Params("id")

	db.Find(&todos, "id = ?", id)

	db.Delete(&todos, "id = ?", id)

	return c.JSON("Todo deleted")
}