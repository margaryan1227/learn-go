package main

import (
	"app-2/src/account"
	"app-2/src/files"
	"fmt"

	"github.com/fatih/color"
)

func main() {
	color.Magenta("%s", "Initializing account-manager-cli 0.0.1...")
	vault, err := account.NewVault()
	if err != nil {
		return
	}

Menu:
	for {
		input := menu()

	Input:
		switch input {
		case "c":
			createAccount(vault)
			break Input
		case "d":
			deleteAccount()
			break Input
		case "r":
			readAccount(vault)
			break Input
		case "q":
			break Menu
		default:
			fmt.Println("Invalid option")
			break Input
		}
	}
}

func menu() string {
	input := ""
	color.Cyan("%s: ", "Select the option to c(reate), d(elete), r(ead) account or q(iut)")
	fmt.Scanf("%s", &input)

	return input
}

func readAccount(vault *account.Vault) {
	url := promptData("Insert url you want to read")

	result := vault.FindAccountByUrl(url)
	color.Cyan("%s", "Accounts found account: ", result)
}

func createAccount(vault *account.Vault) {
	acc1, err := account.NewAccount(
		promptData("Insert login"),
		promptData("Insert password"),
		promptData("Insert url"),
	)
	if err != nil {
		color.Red("%s", "Invalid login or url")
		return
	}

	acc1.Print()

	vault.AddAccount(*acc1)
	bytes, err := vault.ToJSONBytes()
	if err != nil {
		color.Red("%s", "Can't marshal account to json")
	}
	files.WriteFile("accounts.json", bytes)
}

func deleteAccount() {}

func promptData(prompt string) string {
	var res string
	color.Cyan("%s: ", prompt)
	fmt.Scanf("%s", &res)

	return res
}
