package main

import (
	"bufio"
	"fmt"
	"os"
)

//user table

// déclaration de la fonction affichage()
func user(username string, password string) {
	fmt.Println("#################################")
	fmt.Println("\tBonjour" + username)
	fmt.Println("#################################")
}

func getUserInput() string {

	var (
		EntryOk  = false // variable qui permet de vérifier si l'utilisateur a rentré la valeur qu'on attend de lui
		username string
		err      error
		scanner  = bufio.NewScanner(os.Stdin)
	)

	for EntryOk == false {
		fmt.Println("type username: ")
		scanner.Scan()
		username = scanner.Text()

		if err != nil {
			fmt.Println("Only strings !")
		} else {
			EntryOk = true
			fmt.Println("Succes")
		}
	}

	return username

}

func main() {

	getUserInput()

	fmt.println("Your username is : ")

	os.Exit(0)

}
