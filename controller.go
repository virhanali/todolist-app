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
	return c.JSON(fiber.Map{"messages": "Todo created", "data": todos})
}

func GetTodo(c *fiber.Ctx) error {
	var todos []Todo
	db.Find(&todos)
	return c.JSON(fiber.Map{
		"data":    todos,
		"message": "Todos fetched",
	})
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

	if todos.Id == 0 {
		return c.Status(404).JSON(fiber.Map{
			"message": "Todo not found",
		})
	}
	db.Save(&todos)

	return c.JSON(fiber.Map{"messages": "Todo updated", "data": todos})
}

func DeleteTodo(c *fiber.Ctx) error {
	var todos Todo

	id := c.Params("id")

	db.Find(&todos, "id = ?", id)

	if todos.Id == 0 {
		return c.Status(404).JSON(fiber.Map{
			"message": "Todo not found",
		})
	}

	db.Delete(&todos, "id = ?", id)

	return c.JSON(fiber.Map{"messages": "Todo deleted"})
}
