package main

import "fmt"
import "strconv"

import (
	"bufio"
	"os"
)

func basicVariables() {
	fmt.Println("Hello World")

	var x, y, z int
	var boolVar bool
	var stringVar string // initialize variables
	name := "Oscar" // to declare with default value

	nameTwo := "Rafael"
	nameTwo = "Reyes" //change value of var

	numToString, _ := strconv.Atoi("20") //returns two values, second value is the error, and with underscore we are ignoring it

	// var setOfStrings []string

	fmt.Println(x + y + z)
	fmt.Println(boolVar)
	fmt.Println(stringVar)
	fmt.Println("Hello " + name)
	fmt.Println(nameTwo + ", Mi edad es: " + strconv.Itoa(21) )

	// interpolate
	fmt.Printf("My age is: %d \n", numToString) // %t %b %e %f %s
}

func inputExample() {
	var superhero string

	fmt.Println("Favorite superhero:")

	fmt.Scanf("%s\n", &superhero)

	fmt.Printf("%s sucks!!\n", superhero)

	// Using readers
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Say your name")

	yourName, err := reader.ReadString('\n')

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("For %s!!! \n", yourName)
 	}
}

func arrays() {
	var integers [5]int
	strings := [5]string { "wut", "wut1", "wut2"}

	for i := 0; i < len(integers); i++ {
		integers[i] = i;
	}

	fmt.Println(strings)
	fmt.Println(integers)
}

func slices() {
	integers := [5]int { 1, 2, 3 }
	intSlice := []int { 0, 1, 3 }
	dest := make([]int, len(intSlice), cap(intSlice) * 2)

	maked := make([]int, 3, 10)
	

	copy(dest, intSlice)

	fmt.Println(intSlice, dest)
	fmt.Println(maked, len(maked), cap(maked), append(intSlice, 9))
	fmt.Println(integers[:2])
}

func pointers() {
	var pointer *int
	integer := 5

	pointer = &integer

	*pointer = 6

	fmt.Println(pointer, *pointer)
}

type User struct {
	age int
	name, lastname string
	size int
}

type Admin struct {
	User
	isAdmin bool
}

type Util interface {
	Permissions() int
	Name() string
}

func (this Admin) Permissions() int {
	return 5
}

func (this Admin) Name() string {
	return this.name
}

type Editor struct {
	User
}

func (this Editor) Permissions() int {
	return 3
}

func (this Editor) Name() string {
	return this.name
}

func auth(user Util) string {
	if user.Permissions() > 4 {
		return "The user: " + user.Name() + " has admin permissions"
	} else if user.Permissions() == 3 {
		return "The user: " + user.Name() + " has editor permissions"
	}

	return "Unauthorized"
}

func (this User) sayHi(to string) string {
	return "Hello " + to + ", my name is " + this.name + " " + this.lastname
}

func (this Admin) sayHi(to string) string {
	return "Hello "  + to + ", my name is " + this.name + " " + this.lastname + " and I'm an admin"
}

func (this *User) setName(name string) {
	this.name = name
}

func structs() {
	Oscar := Admin{User{name:"Don Señoron", lastname:"Rodriguez", size:2, age:23}, true}

	oscarPointer := new(User)

	fmt.Println(Oscar.name, Oscar.isAdmin)
	(*oscarPointer).name = "Rafael"
	(*oscarPointer).setName("Oscarin")
	fmt.Println(Oscar.sayHi("Wut"))
	fmt.Println(User{age: 23, name: "Oscar", lastname: "Reyes", size: 1}, (*oscarPointer).name, (*oscarPointer).sayHi("Daniel"))
}

func interfaces() {
	admin := Admin{User{23,"Oscar", "Reyes", 2}, true}
	editor := Editor{User{23,"Rafael", "Gaucin", 2}}

	users := []Util{Admin{User{23,"Oscar", "Reyes", 2}, true}, Editor{User{23,"Rafael", "Gaucin", 2}}}

	for _, user := range users {
		fmt.Println(auth(user))
	}

	fmt.Println(">>>>>>>>>>>>>>>>>>")

	fmt.Println(auth(admin))
	fmt.Println(auth(editor))
}

func main() {
	interfaces()	
}