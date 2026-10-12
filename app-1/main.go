package main

import "fmt"

type stringMap = map[string]string

func main() {
	bookmarks := make(stringMap, 5)

Menu:
	for {
		var input string

		fmt.Print("Enter a command please v(iew), a(dd), p(rint), r(emove), q(uit):")
		fmt.Scanf("%s", &input)

	Switch:
		switch input {
		case "v":
			viewBookmark(bookmarks)
		case "a":
			addBookmark(bookmarks)
		case "p":
			printBookmarks(bookmarks)
		case "r":
			removeBookmark(bookmarks)
		case "q":
			break Menu
		default:
			fmt.Println("Invalid command")
			break Switch
		}
	}
}

func addBookmark(bookmarks stringMap) {
	var name, value string
	fmt.Print("Enter your bookmark name:")
	fmt.Scanf("%s", &name)
	fmt.Print("Enter your bookmark value:")
	fmt.Scanf("%s", &value)

	bookmarks[name] = value
}

func removeBookmark(bookmarks stringMap) {
	var name string

	fmt.Print("Enter bookmark name you want to remove:")
	fmt.Scanf("%s", &name)

	delete(bookmarks, name)
}

func viewBookmark(bookmarks stringMap) {
	var name string

	fmt.Print("Enter bookmark name you want to view:")
	fmt.Scanf("%s", &name)

	fmt.Println(name + " -  " + bookmarks[name])
}

func printBookmarks(bookmarks stringMap) {
	for key, value := range bookmarks {
		fmt.Println(key + " -  " + value)
	}
}
