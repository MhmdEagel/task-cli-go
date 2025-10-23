package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"reflect"
	"strconv"
	"tast-tracker/task"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	args := os.Args
	prefix := &args[1]
	command := &args[2]
	var secondArg *string
	var thirdArg *string

	file, err := os.ReadFile("./data/tasks.json")
	if errors.Is(err, os.ErrNotExist) {
		initTaskList, _ := json.Marshal(task.InitTaskList())
		initTaskListJSON := []byte(initTaskList)
		os.WriteFile("./data/tasks.json", initTaskListJSON, 0644)
	}

	var taskList task.TaskList = task.TaskList{}
	json.Unmarshal(file, &taskList)
	check(err)
	
	if *prefix == "task-cli" {
		switch *command {
		case "list":
			for i, val := range taskList.Todo {
				fmt.Printf("[%d] id: %s title: %s\n", i+1, val.Id, val.Title)
			}
			for i, val := range taskList.Done {
				fmt.Printf("[%d] id: %s title: %s\n", i+1, val.Id, val.Title)
			}
			for i, val := range taskList.InProgres {
				fmt.Printf("[%d] id: %s title: %s\n", i+1, val.Id, val.Title)
			}
		case "add":
			secondArg = &args[3]
			if reflect.TypeOf(secondArg).String() == "*string" {
				newId := fmt.Sprintf("%d", len(taskList.Todo)+1)
				newTask := task.NewTask(newId, *secondArg)
				task.AddTask(&taskList.Todo, *newTask)
				task.SaveTask(&taskList)
				fmt.Println("Task successfully added.")
			} else {
				fmt.Println("Invalid type value.")
			}
		case "update":
			secondArg = &args[3]
			thirdArg = &args[4]
			taskId, _ := strconv.ParseInt(*secondArg, 16, 64)
			if reflect.TypeOf(taskId).String() == "int64" {
				task.UpdateTask(&taskList, *secondArg, *thirdArg)
				task.SaveTask(&taskList)
				fmt.Println("Task successfully updated.")
			} else {
				fmt.Println("Invalid type value.")
			}
		case "delete":
			secondArg = &args[3]
			taskId, _ := strconv.ParseInt(*secondArg, 16, 64)
			if reflect.TypeOf(taskId).String() == "int64" {
				result := task.DeleteTask(&taskList.Todo, *secondArg)
				if result {
					task.SaveTask(&taskList)
				} else {
					fmt.Println("Task not found.")
				}
			}
		default :
			fmt.Println("Invalid Command")
			fmt.Println("Available Command:")
			fmt.Println("[*] show - Showing all task")
			fmt.Println("[*] add - Add a new task")
			fmt.Println("[*] update - Change title of a task")
			fmt.Println("[*] delete - Delete a task")
		}
	} else {
		fmt.Println("Invalid prefix. Use task-cli")
	}
}
