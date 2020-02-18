//task list
//update delete search
//works with time
//string
package main

import (
	"fmt"
	"strings"
	"time"
)

type Task struct {
	ID int
	Name string
	StartTime time.Time
	EndTime time.Time
	isDone bool
	Status string
}


type TaskList struct {
	TodoList []*Task
}


func (lists *TaskList) Add(task Task)  {
	lists.TodoList = append(lists.TodoList, &task)

}


func (lists TaskList) Print()  {
	for _, value := range lists.TodoList {
		fmt.Println("ID: ", value.ID)
		fmt.Println("Name: ", value.Name)
		fmt.Println("Start time: ", value.StartTime)
		fmt.Println("End time: ", value.EndTime)
		fmt.Println("Status: ", value.Status)
		fmt.Println("Task id done: ", value.isDone)
		fmt.Println("---------------------------------")
	}
}


func (Lists *TaskList) Update(id int, tl Task) {
	for _, value := range Lists.TodoList {
		if id == value.ID {
			value.ID = tl.ID
			value.Name = tl.Name
			value.EndTime = tl.StartTime
			value.EndTime = tl.EndTime
			value.Status = tl.Status
			value.isDone = tl.isDone
		}
	}
}


func (Lists TaskList) Delete(id int) {
	for index, value := range Lists.TodoList {
		if id == value.ID {
			Lists.TodoList = append(Lists.TodoList[:index], Lists.TodoList[index + 1:]...)


		}
	}
}


func (Lists TaskList) Search(letter string) TaskList  {
	var arr TaskList
	for _, value := range Lists.TodoList {
		if strings.Contains(value.Name, letter) {
			arr.TodoList = append(arr.TodoList, value)
		}
	}
	return arr

}


//1-add
//2-print and check out
//3- check out by ID
//4- if isID false, Update task
//

func main () {

	var tlist TaskList

	startTime, _ := time.Parse("01.02.2006 15:04:05", "02.02.2020 14:00:00")
	endTime, _ := time.Parse("01.02.2006 15:04:05", "02.08.2020 14:00:00")

	t1 := Task {
		ID:        32,
		Name:      "Otabek",
		StartTime: startTime,
		EndTime: endTime,
		isDone:    false,
		Status:    "ok",
	}

	startTime2, _ := time.Parse("01.02.2006 15:04:05", "02.01.2020 09:00:00")
	endTime2, _ := time.Parse("01.02.2006 15:04:05", "02.07.2020 08:00:00")

	t2 := Task{
		ID:        11,
		Name:      "Oybek",
		StartTime: startTime2,
		EndTime:   endTime2,
		isDone:    false,
		Status:    "good luck",
	}
	//1.id
	//2.name
	//starttime
	//endtime
	//isdone
	//status

	tlist.Add(t1)
	tlist.Add(t2)

	fmt.Println("Tasks added and printing now if ID number nis an error\n", "---------------------------------")

	tlist.Print()
	fmt.Println("Tasks will update now")

	t3 := Task{
		ID:        51,
		Name:      "Mashhurbek",
		StartTime: startTime2,
		EndTime:   endTime2,
		isDone:    false,
		Status:    "Very GOOD",
	}

	tlist.Update(11, t3)

	tlist.Print()

	tlist.Delete(32)

	tlist.Print()
}



