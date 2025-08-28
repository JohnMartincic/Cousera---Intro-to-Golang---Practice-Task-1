package greeter

import "fmt"

func Main() {

	var userName string

	fmt.Println("What is your name?")
	fmt.Scanln(userName)
	fmt.Println("Welcome, " + userName + "!")

}