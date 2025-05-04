package main

import "fmt"

func main() {

	shortGoLang := "Watch go crash course";
	fullGoLang := "Watch Nana's Golang Full Course";
	rewardDEssert := "Reward myself with a cheesecake"

	taskItems := []string {shortGoLang, fullGoLang, rewardDEssert}

	fmt.Println("###### Welcome to our Todolist App! ######")

	fmt.Println("List of my Todos")
	
	for index, task:=range(taskItems) {
		fmt.Printf("%d. %s\n", index+1, task)
	}

}
