package main

import (
	"fmt"
	"net/http"
)

var shortGoLang = "Watch go crash course"
var fullGoLang = "Watch Nana's Golang Full Course"
var rewardDEssert = "Reward myself with a cheesecake"
var taskItems = []string {shortGoLang, fullGoLang, rewardDEssert}

func main() {
	http.HandleFunc("/",helloUser)
	http.HandleFunc("/show-tasks",showTasks)
	http.ListenAndServe(":8081",nil)

}

func helloUser(writer http.ResponseWriter, request *http.Request){
	greeting := "Welcome to our Todolist App!"
	fmt.Fprintln(writer, greeting)
}

func showTasks(writer http.ResponseWriter, request *http.Request){
	for _,task:=range(taskItems){
		fmt.Fprintln(writer, task)
	}
}

func addTasks(taskItems []string, newtask string) []string {
	updateTaskItems := append(taskItems, newtask)
	return updateTaskItems
}