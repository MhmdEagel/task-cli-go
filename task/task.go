package task

import (
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"math/rand"
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

func GenerateId() string {
	b := make([]byte, 4) //equals 8 characters
	rand.Read(b)
	s := hex.EncodeToString(b)
	return s
}

func InitTaskList() TaskList {
	var emptyTaskList = TaskList{Todo: []Task{}, Done: []Task{}, InProgres: []Task{}}
	return emptyTaskList
}
func NewTask(title string) *Task {
	return &Task{Id: GenerateId(), Title: title, CreatedAt: time.Now().String(), UpdatedAt: time.Now().String()}
}
func ShowAllTask(taskList *TaskList) {
	fmt.Println("=========\nList Task\n=========")
	ShowTodoTask(taskList)
	ShowDoneTask(taskList)
	ShowInProgresTask(taskList)
}

func ShowTodoTask(taskList *TaskList) {
	fmt.Println("\n[*] Todo Task")
	if len(taskList.Todo) > 0 {
		for i, val := range taskList.Todo {
			fmt.Printf("[%d] id: %s title: %s\n", i+1, val.Id, val.Title)
		}
	} else {
		fmt.Println("List Empty")
	}
}
func ShowDoneTask(taskList *TaskList) {
	fmt.Println("\n[*] Done Task")
	if len(taskList.Done) > 0 {
		for i, val := range taskList.Done {
			fmt.Printf("[%d] id: %s title: %s\n", i+1, val.Id, val.Title)
		}
	} else {
		fmt.Println("List Empty")
	}
}
func ShowInProgresTask(taskList *TaskList) {
	fmt.Println("\n[*] In-Progres Task")
	if len(taskList.InProgres) > 0 {
		for i, val := range taskList.InProgres {
			fmt.Printf("[%d] id: %s title: %s\n", i+1, val.Id, val.Title)
		}
	} else {
		fmt.Println("List Empty")
	}

}

func AddTask(taskList *[]Task, newTask Task) {
	*taskList = append(*taskList, newTask)
}
func DeleteTask(taskList *TaskList, id string) bool {
	newTaskList := TaskList{}
	for _, v := range taskList.Todo {
		if v.Id != id {
			newTaskList.Todo = append(newTaskList.Todo, v)
		}
	}
	for _, v := range taskList.Done {
		if v.Id != id {
			newTaskList.Done = append(newTaskList.Done, v)
		}
	}
	for _, v := range taskList.InProgres {
		if v.Id != id {
			newTaskList.InProgres = append(newTaskList.InProgres, v)
		}
	}

	if len(newTaskList.Done) == len(taskList.Done) && len(newTaskList.InProgres) == len(taskList.InProgres) && len(newTaskList.Todo) == len(taskList.Todo) {
		return false
	}
	*taskList = newTaskList
	return true
}
func UpdateTask(taskList *TaskList, id, title string) bool {
	for i, v := range taskList.Todo {
		if v.Id == id {
			(taskList.Todo)[i].Title = title
			(taskList.Todo)[i].UpdatedAt = time.Now().String()
			return true
		}
	}
	for i, v := range taskList.Done {
		if v.Id == id {
			(taskList.Done)[i].Title = title
			(taskList.Done)[i].UpdatedAt = time.Now().String()
			return true
		}
	}
	for i, v := range taskList.InProgres {
		if v.Id == id {
			(taskList.InProgres)[i].Title = title
			(taskList.InProgres)[i].UpdatedAt = time.Now().String()
			return true
		}
	}
	return false
}
func SaveTask(taskList *TaskList) {
	allTask, err := json.Marshal(taskList)
	allTaskJSON := []byte(allTask)
	if err != nil {
		panic(err)
	}
	os.WriteFile("./data/tasks.json", allTaskJSON, 0644)
}

func FindTaskInTodoAndInProgres(taskList *TaskList, id string) *Task {
	for i, v := range taskList.Todo {
		if v.Id == id {
			return &taskList.Todo[i]
		}
	}
	for i, v := range taskList.InProgres {
		if v.Id == id {
			return &taskList.InProgres[i]
		}
	}
	return nil
}

func FindTaskInTodo(taskList *TaskList, id string) *Task {
	for i, v := range taskList.Todo {
		if v.Id == id {
			return &taskList.Todo[i]
		}
	}
	return nil
}

func DoneTask(taskList *TaskList, id string) bool {
	foundTask := FindTaskInTodoAndInProgres(taskList, id)
	if foundTask != nil {
		DeleteTask(taskList, foundTask.Id)
		taskList.Done = append(taskList.Done, *foundTask)
		SaveTask(taskList)
		return true
	}
	return false
}

func InProgresTask(taskList *TaskList, id string) bool {
	foundTask := FindTaskInTodo(taskList, id)
	if foundTask != nil {
		DeleteTask(taskList, foundTask.Id)
		taskList.InProgres = append(taskList.InProgres, *foundTask)
		SaveTask(taskList)
		return true
	}
	return false
}

func ReadTask() TaskList {
	file, err := os.ReadFile("./data/tasks.json")
	if errors.Is(err, os.ErrNotExist) {
		initTaskList, _ := json.Marshal(InitTaskList())
		initTaskListJSON := []byte(initTaskList)
		os.WriteFile("./data/tasks.json", initTaskListJSON, 0644)
		// attempt to read again if file not exist
		file, _ := os.ReadFile("./data/tasks.json")
		taskList := TaskList{}
		json.Unmarshal(file, &taskList)
		return taskList
	}
	taskList := TaskList{}
	json.Unmarshal(file, &taskList)
	return taskList
}
