package main

import (
	"fmt"
)

func main() {
	tasks := []string{}

	for {
		var input string
		fmt.Print("masukkan task")
		fmt.Scan(&input)
		if input == "exit" {
			break
		}

		tasks = append(tasks, input)
		fmt.Println("task berhasil ditambahkan")

	}

	fmt.Println("Tasks:")
	for index, task := range tasks {
		fmt.Println("%d. %s\n", index+1, task)
	}

}
