package main

type Todo struct {
	Id        int    `json:"id" gorm:"primary_key; autoIncrement"`
	Task      string `json:"task"`
	Completed bool   `json:"completed"`
}

type updateTodo struct {
	Task      string `json:"task"`
	Completed bool   `json:"completed"`
}
