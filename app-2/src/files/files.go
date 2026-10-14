package files

import (
	"os"

	"github.com/fatih/color"
)

type JsonDB struct {
	fileName string
}

func NewJsonDB(name string) *JsonDB {
	return &JsonDB{
		fileName: name,
	}
}

func (db *JsonDB) Write(data []byte) {
	file, err := os.Create(db.fileName)
	if err != nil {
		color.Red("%s", err)
	}

	defer file.Close()
	_, err = file.Write(data)

	if err != nil {
		color.Red("%s", err)
		return
	}

	color.Cyan("%s", "Successfully wrote "+db.fileName)
}

func (db *JsonDB) Read() ([]byte, error) {
	file, err := os.ReadFile(db.fileName)
	if err != nil {
		color.Red("%s", err)
		return nil, err
	}

	return file, nil
}
