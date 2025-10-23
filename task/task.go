package task

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

type TaskList struct {
	Todo      []Task
	Done      []Task
	InProgres []Task
}

type Task struct {
	Id        string
	Title     string
	CreatedAt string
	UpdatedAt string
}

func InitTaskList() TaskList {
	var emptyTaskList = TaskList{Todo: []Task{}, Done: []Task{}, InProgres: []Task{}}
	return emptyTaskList
}
func NewTask(id, title string) *Task {
	return &Task{Id: id, Title: title, CreatedAt: time.Now().String(), UpdatedAt: time.Now().String()}
}
func ShowAllTask(taskList *[]Task) {
	fmt.Println("=========\nList Task\n=========")
	for i, v := range *taskList {
		fmt.Printf("[%d] %s \n", i+1, v.Title)
	}
}
func AddTask(taskList *[]Task, newTask Task) {
	*taskList = append(*taskList, newTask)
}
func DeleteTask(taskList *[]Task, id string) bool {
	newTaskList := []Task{}
	for _, v := range *taskList {
		if v.Id != id {
			newTaskList = append(newTaskList, v)
		}
	}

	if len(newTaskList) == len(*taskList) {
		return false
	}
	*taskList = newTaskList
	return true
}
func UpdateTask(taskList *[]Task, id, title string) bool {
	for i, v := range *taskList {
		if v.Id == id {
			(*taskList)[i].Title = title
			(*taskList)[i].UpdatedAt = time.Now().String()
			return true
		}
	}
	return false
}
func SaveTask(taskList *[]Task) {
	allTask, err := json.Marshal(taskList)
	allTaskJSON := []byte(allTask)
	if err != nil {
		panic(err)
	}
	os.WriteFile("./data/tasks.json", allTaskJSON, 0644)
}
