package main

import "fmt"

func main() {

	shortGoLang := "Watch go crash course";
	fullGoLang := "Watch Nana's Golang Full Course";
	rewardDEssert := "Reward myself with a cheesecake"
	taskItems := []string {shortGoLang, fullGoLang, rewardDEssert}

	fmt.Println("###### Welcome to our Todolist App! ######")

	printTasks(taskItems)

	taskItems = addTasks(taskItems, "Go for a run")
	
	fmt.Println()
	printTasks(taskItems)

}

func printTasks(taskItems []string){
	fmt.Println("List of my Todos")
	for index, task:=range(taskItems) {
		fmt.Printf("%d. %s\n", index+1, task)
	}
}

func addTasks(taskItems []string, newtask string) []string {
	updateTaskItems := append(taskItems, newtask)
	return updateTaskItems
}