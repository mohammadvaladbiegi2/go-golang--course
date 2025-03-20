package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
)

type User struct {
	ID       string
	email    string
	password string
}

var categorys []string
var userStoreg []User

func main() {
	command := flag.String("command", "no-command", "create a command")
	flag.Parse()

	for {
		runcommand(*command)
		scanner := bufio.NewScanner(os.Stdin)
		fmt.Println("Please enter the new command")
		scanner.Scan()
		*command = scanner.Text()
	}

}

func runcommand(command string) {
	switch command {
	case "creaet-task":
		createTask()
	case "register":
		register()
	case "login":
		login()
	case "create-category":
		createCategory()
	case "list-category":
		listCategories()
	case "exit":
		os.Exit(0)
	default:
		fmt.Println("Enter a valid command\nList of commands:")
		fmt.Println("create-task => Create a new task")
		fmt.Println("register => Sign up a new user")
		fmt.Println("login => Log in to your account")
		fmt.Println("create-category => Create a new category")
		fmt.Println("list-category => List all categories")
		fmt.Println("exit => Exit the application")
	}
}

func createCategory() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Please enter the category of task")
	scanner.Scan()
	category := scanner.Text()
	categorys = append(categorys, category)
	fmt.Println("Category created successfully ✅✅")
	fmt.Printf("List of categories => %v\n", categorys)
}

func listCategories() {
	fmt.Printf("List of categories => %v\n", categorys)
}

func createTask() {
	scanner := bufio.NewScanner(os.Stdin)
	var title_task string
	var category string
	var doudate string
	fmt.Println("Please enter the title of the task")
	scanner.Scan()
	title_task = scanner.Text()

	fmt.Println("Please enter the category of the task")
	scanner.Scan()
	category = scanner.Text()

	fmt.Println("Please enter the due date of the task")
	scanner.Scan()
	doudate = scanner.Text()

	fmt.Printf("Task created with this info: title => %s, category => %s, due date => %s\n", title_task, category, doudate)
}

func register() {
	scanner := bufio.NewScanner(os.Stdin)
	var email string
	var password string
	var id string
	fmt.Println("Signup\nPlease enter the email")
	scanner.Scan()
	email = scanner.Text()

	fmt.Println("Please enter the password ")
	scanner.Scan()
	password = scanner.Text()

	id = email
	userStoreg = append(userStoreg, User{email: email, ID: id, password: password})

	fmt.Println("welcom")
}

func login() {
	scanner := bufio.NewScanner(os.Stdin)
	var email string
	var password string
	fmt.Println("Signup\nPlease enter the email")
	scanner.Scan()
	email = scanner.Text()

	fmt.Println("Please enter the password ")
	scanner.Scan()
	password = scanner.Text()

	for _, v := range userStoreg {
		if v.email == email && v.password == password {
			fmt.Println("welcom")
			return
		}
	}

	fmt.Println("user not found please register first")
}
