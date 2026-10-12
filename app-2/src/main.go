package main

import (
	"app-2/src/account"
	"app-2/src/files"
	"fmt"
)

func main() {
	menu()
}

func menu() {
	input := ""
Menu:
	for {
		fmt.Printf("%s: ", "Select the option to c(reate), d(elete), r(ead) account or q(iut)")
		fmt.Scanf("%s", &input)

	Input:
		switch input {
		case "c":
			createAccount()
			break Input
		case "d":
			deleteAccount()
			break Input
		case "r":
			readAccount()
			break Input
		case "q":
			break Menu
		default:
			fmt.Println("Invalid option")
			break Input
		}
	}
}

func createAccount() {
	acc1, err := account.NewAccount(
		promptData("Insert login"),
		promptData("Insert password"),
		promptData("Insert url"),
	)
	if err != nil {
		fmt.Println("invalid login or url")
		return
	}

	acc1.Print()

	bytes, err := acc1.ToJSONBytes()
	if err != nil {
		fmt.Println("Can't marshal account to json")
	}
	files.WriteFile("account1.json", bytes)
}

func readAccount() {

}

func deleteAccount() {}

func promptData(prompt string) string {
	var res string
	fmt.Printf("%s: ", prompt)
	fmt.Scanf("%s", &res)

	return res
}
