package main

import "fmt"

func main() {

	shortGoLang := "Watch go crash course";
	fullGoLang := "Watch Nana's Golang Full Course";
	rewardDEssert := "Reward myself with a cheesecake"

	taskItems := []string {shortGoLang, fullGoLang, rewardDEssert}

	fmt.Println("###### Welcome to our Todolist App! ######")

	fmt.Println("List of my Todos")
	for _, task:=range(taskItems) {
		fmt.Println(task)
	}

}
