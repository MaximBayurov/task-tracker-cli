package main

import (
	"github.com/MaximBayurov/task-tracker-cli/internal/args"
	"github.com/MaximBayurov/task-tracker-cli/internal/manager"
	"github.com/MaximBayurov/task-tracker-cli/internal/storage"
	"log"
)

func main() {
	_ = storage.Init("todo-list.json")
	manager.Init()

	arguments := args.GetParsed()

	m, err := manager.GetInstance()
	if err != nil {
		log.Fatal(err)
	}
	err = m.HandleFromArgs(arguments)
	if err != nil {
		log.Fatal(err)
	}
}
