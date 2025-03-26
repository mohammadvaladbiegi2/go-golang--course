package main

import (
	"bufio"
	"crypto/md5"
	"encoding/hex"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type User struct {
	ID       int
	email    string
	password string
}

var userStoreg []User

type Task struct {
	ID       int
	Title    string
	DouDate  string
	Category string
	IsDone   bool
	UserID   int
}

type CategoryStruct struct {
	title  string
	UserID int
}

var categorys []CategoryStruct
var taskStoreg []Task
var userID int

const UserStoragePath = "user-storage.txt"

func main() {
	loadUsers()
	command := flag.String("command", "no-command", "create a command")
	flag.Parse()
	authentications()
	if userID == 0 {
		authentications()
	}
	for {
		scanner := bufio.NewScanner(os.Stdin)
		fmt.Println("Please enter the new command or tab to see list of command")
		scanner.Scan()
		*command = scanner.Text()
		runcommand(*command)
	}

}

func cerateFile(user string) {

	file, err := os.OpenFile(UserStoragePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	defer file.Close() // این کد در اخر تابع اجرا میشه

	if err != nil {
		fmt.Println("خطا در ایجاد فایل:", err)
		return
	}

	_, err = file.WriteString(user)
	file.Close()
	if err != nil {
		fmt.Println("خطا در نوشتن فایل:", err)
	}
}

func runcommand(command string) {
	switch command {
	case "create-task":
		createTask()
	case "create-category":
		createCategory()
	case "list-category":
		listCategories()
	case "exit":
		os.Exit(0)
	default:
		fmt.Println("Enter a valid command\nList of commands:")
		fmt.Println("create-task => Create a new task")
		fmt.Println("create-category => Create a new category")
		fmt.Println("list-category => List all categories")
		fmt.Println("exit => Exit the application")
	}
}

func createCategory() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("create a  category for task")
	scanner.Scan()
	category_title := scanner.Text()
	categorys = append(categorys, CategoryStruct{title: category_title, UserID: userID})
	fmt.Println("Category created successfully ✅✅")
}

func listCategories() {
	fmt.Printf("List of categories => %v\n", categorys)
}

func createTask() {
	if len(categorys) == 0 {
		createCategory()
	}
	scanner := bufio.NewScanner(os.Stdin)
	var title_task string
	var category string
	var doudate string
	fmt.Println("Please enter the title of the task")
	scanner.Scan()
	title_task = scanner.Text()

	fmt.Println("Please enter the category of the task")
	fmt.Println("list of categorys => ")
	for _, value := range categorys {
		fmt.Printf("%s  ", value.title)
	}
	fmt.Println("")
	scanner.Scan()
	category = scanner.Text()
	for {
		if isValidCategory(scanner.Text()) {
			break
		}
		fmt.Println("Please enter the valid category")
		fmt.Println("list of categorys => ")
		for _, value := range categorys {
			fmt.Printf("%s  ", value.title)
		}
		scanner.Scan()
		category = scanner.Text()

	}

	fmt.Println("Please enter the due date of the task")
	scanner.Scan()
	doudate = scanner.Text()

	task := Task{
		ID:       len(taskStoreg) + 1,
		Title:    title_task,
		DouDate:  doudate,
		Category: category,
		IsDone:   false,
		UserID:   userID,
	}
	taskStoreg = append(taskStoreg, task)
	fmt.Printf("Task created with this info: title => %s, category => %s, due date => %s\n, ID of creator => %d\n", title_task, category, doudate, userID)
}

func isValidCategory(category string) bool {
	for _, value := range categorys {
		if value.title == category {
			return true
		}
	}

	return false
}

func authentications() {
	scanner := bufio.NewScanner(os.Stdin)
	var email string
	var password string
	var id int
	fmt.Println("Signup / login \nPlease enter the email")
	scanner.Scan()
	email = scanner.Text()

	fmt.Println("Please enter the password ")
	scanner.Scan()
	password = scanner.Text()

	for _, user := range userStoreg {
		if user.email == email && user.password == hashPassword(password) {
			id = len(userStoreg) + 1
			userID = id
			fmt.Println("welcom")
			return
		} else if user.email == email || user.password == hashPassword(password) {
			fmt.Println("user or password not crect try agin")
			return
		}
	}

	id = len(userStoreg) + 1
	userID = id
	user := fmt.Sprintf("ID : %d,email : %s,password : %s\n", id, email, hashPassword(password))
	cerateFile(user)
	fmt.Println("welcom")
}

func loadUsers() {
	file, err := os.Open(UserStoragePath)
	if err != nil {
		fmt.Println("خطا در باز کردن فایل:", err)
		return
	}
	defer file.Close() // این کد در اخر تابع اجرا میشه

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, ",")
		if len(parts) != 3 {
			fmt.Println("خطا در پردازش خط:", line)
			continue
		}

		idStr := strings.TrimPrefix(parts[0], "ID : ")
		email := strings.TrimPrefix(parts[1], "email : ")
		password := strings.TrimPrefix(parts[2], "password : ")

		id, err := strconv.Atoi(strings.TrimSpace(idStr))
		if err != nil {
			fmt.Println("خطا در تبدیل ID:", err)
			continue
		}

		user := User{
			ID:       id,
			email:    strings.TrimSpace(email),
			password: strings.TrimSpace(password),
		}
		userStoreg = append(userStoreg, user)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("خطا در خواندن فایل:", err)
	}
}

func hashPassword(text string) string {
	hash := md5.Sum([]byte(text))
	return hex.EncodeToString(hash[:])
}
