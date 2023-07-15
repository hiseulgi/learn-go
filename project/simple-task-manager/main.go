// TODO: Membuat penyimpanan tasks di file binary/json/database

package main

import (
	"bufio"
	"fmt"
	"os"
	"simple-task-manager/library/task"
)

func main() {
	// init task slice
	tasks := task.TaskSlice{}

	// init scanner
	scanner := bufio.NewScanner(os.Stdin)

looping:
	for {

		clear()

		choice := 0

		fmt.Println("Simple Task Manager")
		fmt.Println("====================")
		fmt.Println("1. Add New Task")
		fmt.Println("2. Mark / Unmark Task")
		fmt.Println("3. Tasks List")
		fmt.Println("4. Delete Task")
		fmt.Println("5. Exit Program")
		fmt.Println("====================")
		fmt.Print("Choose the option: ")
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			createMenu(scanner, &tasks)
		case 2:
			markMenu(&tasks)
		case 3:
			listMenu(&tasks)
		case 4:
			deleteMenu(&tasks)
		case 5:
			{
				clear()
				break looping
			}
		default:
			{
				func() {
					defer blocking()

					clear()
					fmt.Println("Invalid choice!")
				}()
			}
		}
	}
}
