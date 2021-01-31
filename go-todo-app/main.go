package main

import (
	"fmt"
	"go-tools/go-todo-app/app/controllers"
	"go-tools/go-todo-app/app/models"
)

func main() {
	fmt.Println(models.Db)

	controllers.StartMainServer()
}
