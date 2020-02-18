package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
)

type TodoList struct {
	Dates []string
	Times []string
	Things []string
}
func (self *TodoList) Print() {
	if len(self.Things) != 0 {
		for i, _ := range self.Dates {
			fmt.Printf("Date: %v, Time: %v, Thing: %v\n", self.Dates[i], self.Times[i], self.Things[i])
		}
	} else {
		fmt.Println("Todo list is empty.")
	}
}
func (self *TodoList) Add(date string, time string, thing string) {
	self.Dates = append(self.Dates, date)
	self.Times = append(self.Times, time)
	self.Things = append(self.Things, thing)
}
func (self *TodoList) Remove(i int) {
	self.Dates = append(self.Dates[:i], self.Dates[i+1:]...)
	self.Times = append(self.Times[:i], self.Times[i+1:]...)
	self.Things = append(self.Things[:i], self.Things[i+1:]...)
}
func (self *TodoList) Find(thing string) (int, error) {
	for i, v := range self.Things {
		if v == thing {
			return i, nil
		}
	}
	return -1, fmt.Errorf("%s not found in todo list.", thing)
}

func main() {
	todoList := TodoList{}
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Println("Add something to your todo list, Remove something from your todo list, or view your todo list, or quit this program?")
		scanner.Scan()
		choice := scanner.Text()
		if strings.Contains(strings.ToLower(choice), "view") {
			todoList.Print()
		} else if strings.Contains(strings.ToLower(choice), "add") {
			fmt.Println("Enter the date:")
			scanner.Scan()
			date := scanner.Text()
			fmt.Println("Enter the time:")
			scanner.Scan()
			time := scanner.Text()
			fmt.Println("Enter the thing:")
			scanner.Scan()
			thing := scanner.Text()
			todoList.Add(date, time, thing)
		} else if strings.Contains(strings.ToLower(choice), "remove") {
			fmt.Println("Enter the thing you want to remove:")
			scanner.Scan()
			thing := scanner.Text()
			if index, err := todoList.Find(thing); err != nil {
				fmt.Printf("%s not found in todo list.\n", thing)
			} else {
				todoList.Remove(index)
				fmt.Printf("%s succesfully removed!\n", thing)
			}
		} else if strings.Contains(strings.ToLower(choice), "quit") {
			fmt.Println("Program quitting.")
			os.Exit(0)
		}
		fmt.Println("")
	}
}