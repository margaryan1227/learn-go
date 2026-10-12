package files

import (
	"fmt"
	"os"
)

func WriteFile(path string, data []byte) {
	file, err := os.Create(path)
	if err != nil {
		fmt.Println(err)
	}

	defer file.Close()
	_, err = file.Write(data)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Successfully wrote " + path)
}

func ReadFile(path string) ([]byte, error) {
	file, err := os.ReadFile(path)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return file, nil
}
