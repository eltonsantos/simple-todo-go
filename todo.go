package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	tasks := []string{}
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("=== Simple Todo Go | Developed by Elton Santos - 08.04.2023 ===")

	for {
		fmt.Println("\n=== My Simple Todo List ===")
		fmt.Println("1. Add Task")
		fmt.Println("2. Show Tasks")
		fmt.Println("3. Remove Task")
		fmt.Println("4. Exit")
		fmt.Print("\nChoose a option: ")

		choice, _ := reader.ReadString('\n')
		choice = strings.TrimSpace(choice)

		switch choice {
		case "1":
			fmt.Print("Add a task: ")
			task, _ := reader.ReadString('\n')
			task = strings.TrimSpace(task)
			tasks = append(tasks, task)
			fmt.Println("Task added successfully!")
		case "2":
			fmt.Println("\n=== All Tasks ===")
			if len(tasks) == 0 {
				fmt.Println("There is no task registered.")
			} else {
				for i, task := range tasks {
					fmt.Printf("%d. %s\n", i+1, task)
				}
			}
		case "3":
			fmt.Println("\n=== All Tasks ===")
			if len(tasks) == 0 {
				fmt.Println("There is no task registered.")
			} else {
				fmt.Print("Enter the number of the task to be removed: ")
				var taskIndex int
				_, err := fmt.Scanf("%d", &taskIndex)
				if err != nil || taskIndex < 1 || taskIndex > len(tasks) {
					fmt.Println("Invalid task number.")
					continue
				}
				tasks = append(tasks[:taskIndex-1], tasks[taskIndex:]...)
				fmt.Println("Task removed successfully!")
			}
		case "4":
			fmt.Println("Exiting...")
			return
		default:
			fmt.Println("\nInvalid option! Try again.")
		}
	}
}
