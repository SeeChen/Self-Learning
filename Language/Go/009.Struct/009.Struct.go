package main

import "fmt"

// Here is how to declare a structure in Golang.
//
//	type struct_variable_type struct {
//			member definition
//			member definition
//			...
//			member definition
//	}
type User struct {
	Username string
	Email    string
	Age      int
}

// Calculator
type Calc struct {
	Operation func(int, int) int
}

// Nesting struct
// Golang allows us to nest structures
type Book struct {
	Name   string
	Pages  int
	Author string
}
type Novel struct {
	Description string
	BookInfo    Book
}

// Anonymous field (embedded struct) allows accessing Book's fields directly from Comics
// Example: comics.Name instead of comics.Book.Name
type Comics struct {
	Book
}

// Format the output
// This function uses a struct as a parameter
func formatUser(user User) string {

	var strUser = "%-10s: %s\n%-10s: %s\n%-10s: %d\n"
	return fmt.Sprintf(strUser, "Username", user.Username, "Email", user.Email, "Age", user.Age)
}

// Change username
func changeUsername(user User, newUsername string) User {
	fmt.Println("Original User data")
	fmt.Println(formatUser(user))

	user.Username = newUsername

	fmt.Println("New User data")
	fmt.Println(formatUser(user))
	return user
}

// Change the username using the structure pointer
func changeUsernamePtr(ptrUser *User, newUsername string) {
	ptrUser.Username = newUsername
}

// If a method needs to modify the struct's data or avoid copying,
// use a pointer receiver: (u *User) FuncName()
func (user User) Message(birth string) string {
	fmtString := "Hi, my name is %s, my email is %s, my age is %d, my birthday is %s.\n"
	return fmt.Sprintf(fmtString, user.Username, user.Email, user.Age, birth)
}

// Function to add two numbers
func addCalc(x int, y int) int {
	return x + y
}

// Function to subtract two numbers
func subCalc(x int, y int) int {
	return x - y
}

func main() {
	// For array structures, multiple data can only store one type.
	// However, Golang structures can store multiple data types.

	// Use the `var variable type` to define a variable
	var User1 User = User{
		Username: "SeeChen",
		Email:    "leeseechen@gmail.com",
		Age:      25}
	fmt.Println("User1:")
	fmt.Println(formatUser(User1))

	// Declare variables using `:=`
	User2 := User{"SeeChen2", "leeseechen@email.com", 100}
	fmt.Println("User2:")
	fmt.Println(formatUser(User2))

	// Or like the following method
	var User3 User
	User3.Username = "SeeChen 3"
	User3.Email = "seechenlee@gmail.com"
	User3.Age = 0
	fmt.Println("User3:")
	fmt.Println(formatUser(User3))

	// As with other value types, to change a struct value, must return the modified struct
	// Try change the username without returning a value
	changeUsername(User1, "SeeChen SeeChen")
	// But the main line has no changes
	fmt.Println("===")
	fmt.Println(formatUser(User1))

	// With return the value
	User1 = changeUsername(User1, "SeeChen SeeChen")
	fmt.Println("===")
	fmt.Println(formatUser(User1))

	// Or the same as other types of variables
	// Can use pointers to pass parameters
	changeUsernamePtr(&User2, "SeeChen 22")
	fmt.Println(formatUser(User2))

	// The functions in structures
	// In Golang, there are two ways to declare functions in structures.
	// 1. Using (struct) fucntion()
	fmt.Println(User1.Message("06 Aug 2000"))
	fmt.Println(User3.Message("06 Aug 2000"))

	// The second way is define a function in a structure and assign the function to the structure.
	// Create a new Structure Calc
	var calcAdd Calc
	var calcSub Calc
	calcAdd.Operation = addCalc
	calcSub.Operation = subCalc

	fmt.Printf("10 + 5 = %d\n", calcAdd.Operation(10, 5))
	fmt.Printf("10 - 5 = %d\n", calcSub.Operation(10, 5))

	// Nested Structs
	// In Golang, we can nested structs in our code.
	// This can make the program similar to inherited properties.
	novel := Novel{
		Description: "This is a novel",
		BookInfo: Book{
			Name:   "Novel",
			Pages:  100,
			Author: "Novel Author"}}
	fmt.Println(novel)
	// To access
	fmt.Println(novel.BookInfo.Author)

	// Anonymous fields (embedded structures)
	var comics Comics = Comics{Book{Name: "Comics", Pages: 200, Author: "Comics Author"}}
	fmt.Println(comics)
	fmt.Println(comics.Name)
}
