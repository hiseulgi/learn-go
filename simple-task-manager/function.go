// * Function for main.go
// this file contains all function for main.go

package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"simple-task-manager/library/task"
	"strconv"
)

// * Clear Terminal
// this function is used to clear terminal
func clear() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

// * Catch Error
// this function is used to catch error
func catch() {
	if r := recover(); r != nil {
		fmt.Println("====================")
		fmt.Println("error:", r)
	}
}

// * Blocking program
// this function is used to block program until user press enter
func blocking() {
	fmt.Println("Press enter to continue!")
	fmt.Scanln()
}

// * Choice Menu Function
/* -------------------------------------------------------------------------- */

// * Create Task Menu
// this function is used to create task menu
func createMenu(scanner *bufio.Scanner, tasks *task.TaskSlice) {
	defer blocking()
	defer catch()
	clear()
	fmt.Println("1. Add New Task")
	fmt.Println("====================")
	fmt.Print("Enter new task name: ")
	scanner.Scan()
	taskName := scanner.Text()
	fmt.Println("====================")
	data := task.Task{Name: taskName, IsCompleted: false}
	if err := tasks.CreateTask(data); err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(taskName, "has been added!")

}

// * Mark / Unmark Task Menu
// this function is used to mark / unmark task menu
func markMenu(tasks *task.TaskSlice) {
	defer blocking()
	defer catch()

	clear()
	fmt.Println("2. Mark / Unmark Task")
	fmt.Println("====================")
	taskIndexStr := ""
	fmt.Print("Enter task index: ")
	fmt.Scanln(&taskIndexStr)
	taskIndex, err := strconv.Atoi(taskIndexStr)

	// check input type is integer
	if err != nil {
		fmt.Println("Invalid input. Please enter an integer")
		return
	}

	// check negative index
	if taskIndex < 0 {
		fmt.Println("Invalid input. Please enter a positive integer")
		return
	}
	fmt.Println("====================")

	// check there is no error on marking task process
	if err := tasks.MarkTask(taskIndex); err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("Task with", taskIndex, "has been mark as completed!")
}

// * Tasks List Menu
// this function is used to list all task menu
func listMenu(tasks *task.TaskSlice) {
	defer blocking()
	defer catch()

	clear()
	fmt.Println("3. Tasks List")
	fmt.Println("====================")
	tasks.ReadTask()
	fmt.Println("====================")
}

// * Delete Task Menu
// this function is used to delete task menu
func deleteMenu(tasks *task.TaskSlice) {
	defer blocking()
	defer catch()

	clear()
	fmt.Println("4. Delete Task")
	fmt.Println("====================")
	taskIndexStr := ""
	fmt.Print("Enter task index: ")
	fmt.Scanln(&taskIndexStr)
	taskIndex, err := strconv.Atoi(taskIndexStr)

	// check input type is integer
	if err != nil {
		fmt.Println("Invalid input. Please enter an integer")
		return
	}

	// check negative index
	if taskIndex < 0 {
		fmt.Println("Invalid input. Please enter a positive integer")
		return
	}

	// delete confirmation alert
	deleteConfirmation := 0
	fmt.Println("are you sure want to delete task with index", taskIndex, "?")
	fmt.Println("1. Yes")
	fmt.Println("2. No")
	fmt.Print("Enter your choice: ")
	fmt.Scanln(&deleteConfirmation)

	if deleteConfirmation != 1 {
		fmt.Println("Task with", taskIndex, "has not been deleted!")
		return
	}

	fmt.Println("====================")

	// check there is no error on marking task process
	if err := tasks.DeleteTask(taskIndex); err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("Task with", taskIndex, "has been deleted!")

}
