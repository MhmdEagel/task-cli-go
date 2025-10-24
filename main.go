package main

import (
	"fmt"
	"os"
	"tast-tracker/task"
)

func main() {
	args := os.Args
	prefix := &args[1]
	var secondArg *string
	var thirdArg *string

	var taskList task.TaskList = task.ReadTask()

	if *prefix == "task-cli" {
		command := &args[2]
		switch *command {
		case "list":
			if len(args) >= 4 {
				showOption := &args[3]
				switch *showOption {
				case "todo":
					task.ShowTodoTask(&taskList)
				case "done":
					task.ShowDoneTask(&taskList)
				case "in-progress":
					task.ShowInProgresTask(&taskList)
				default:
					fmt.Println("Invalid argument")
				}
			} else {
				task.ShowAllTask(&taskList)
			}
		case "add":
			if len(args) >= 4 {
				secondArg = &args[3]
				newTask := task.NewTask(*secondArg)
				task.AddTask(&taskList.Todo, *newTask)
				task.SaveTask(&taskList)
				fmt.Println("Task successfully added.")
			} else {
				fmt.Println("Missing argument")
			}
		case "update":
			secondArg = &args[3]
			thirdArg = &args[4]
			if len(args) >=5 {
				task.UpdateTask(&taskList, *secondArg, *thirdArg)
				task.SaveTask(&taskList)
				fmt.Println("Task successfully updated.")
			} else {
				fmt.Println("Missing argument.")
			}
		case "delete":
			if len(args) >= 4 {
				secondArg = &args[3]
				result := task.DeleteTask(&taskList, *secondArg)
				if result {
					task.SaveTask(&taskList)
					fmt.Println("Task deleted")
				} else {
					fmt.Println("Task not found.")
				}
			} else {
				fmt.Println("Missing argument. Expected Id.")
			}
		case "mark-in-progres":
			if len(args) >= 4 {
				secondArg = &args[3]
				result := task.InProgresTask(&taskList, *secondArg)

				if result {
					fmt.Println("Process done successfully.")
				} else {
					fmt.Println("Task not found.")
				}
			} else {
				fmt.Println("Missing argument. Expected Id.")
			}
		case "mark-done":
			if len(args) >= 4 {
				secondArg = &args[3]
				result := task.DoneTask(&taskList, *secondArg)
				if result {
					fmt.Println("Process done successfully.")
				} else {
					fmt.Println("Task not found.")
				}
			} else {
				fmt.Println("Missing argument. Expected Id.")
			}
		default:
			fmt.Println("Invalid Command")
			fmt.Println("Available Command:")
			fmt.Println("[*] show - Showing all task")
			fmt.Println("[*] add - Add a new task")
			fmt.Println("[*] update - Change title of a task")
			fmt.Println("[*] delete - Delete a task")
			fmt.Println("[*] mark-done - Mark a task to be done")
			fmt.Println("[*] mark-in-progres - Mark a task to be in-progres")
		}
	} else {
		fmt.Println("Invalid prefix. Use task-cli")
	}
}
