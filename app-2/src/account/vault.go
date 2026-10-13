package account

import (
	"app-2/src/files"
	"encoding/json"
	"strings"
	"time"

	"github.com/fatih/color"
)

type Vault struct {
	Accounts  []Account `json:"accounts"`
	UpdatedAt time.Time `json:"updated_at"`
}

func NewVault() (*Vault, error) {
	file, err := files.ReadFile("accounts.json")
	if err != nil {
		return &Vault{
			Accounts:  []Account{},
			UpdatedAt: time.Now(),
		}, nil
	}

	var vault Vault

	err = json.Unmarshal(file, &vault)
	if err != nil {
		color.Red("%s", "Error reading file: ", err)
		return &Vault{
			Accounts:  []Account{},
			UpdatedAt: time.Now(),
		}, nil
	}

	return &vault, err
}

func (vault *Vault) AddAccount(account Account) {
	vault.Accounts = append(vault.Accounts, account)
	data, err := vault.ToJSONBytes()
	if err != nil {
		color.Red("%s", "Error converting accounts to json: ", err)
	}

	files.WriteFile("accounts.json", data)
}

func (vault *Vault) FindAccountByUrl(urlString string) []Account {
	var accounts = make([]Account, 2)

	for _, account := range vault.Accounts {
		if strings.Contains(account.Url, urlString) {
			accounts = append(accounts, account)
		}
	}

	return accounts
}

func (vault *Vault) DeleteByUrl(url string) []Account {

	deletedAccounts := make([]Account, 0)
	for idx, account := range vault.Accounts {
		if strings.Contains(account.Url, url) {
			deletedAccounts = append(deletedAccounts, account)
			vault.Accounts = append(vault.Accounts[:idx], vault.Accounts[idx+1:]...)
		}
	}

	data, err := vault.ToJSONBytes()
	if err != nil {
		color.Red("%s", "Error converting accounts to json: ", err)
		return deletedAccounts
	}
	files.WriteFile("accounts.json", data)

	return deletedAccounts
}

func (vault *Vault) ToJSONBytes() ([]byte, error) {
	file, err := json.Marshal(vault)
	if err != nil {
		return nil, err
	}

	return file, nil
}
