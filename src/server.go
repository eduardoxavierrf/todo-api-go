package server

import (
	"fmt"
	"log"
	"net/http"

	"github.com/eduardoxavierrf/todo-api-go/src/database"
	models "github.com/eduardoxavierrf/todo-api-go/src/model"
)

func Init() {
	db := database.Init()
	defer db.Close()

	db.AutoMigrate(&models.Task{})

	fmt.Println("Starting server at :3000")
	log.Fatal(http.ListenAndServe(":3000", nil))
}
