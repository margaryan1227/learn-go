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

type VaultWithDB struct {
	Vault
	db files.JsonDB
}

func NewVault(db *files.JsonDB) (*VaultWithDB, error) {
	file, err := db.Read()
	if err != nil {
		return &VaultWithDB{
			Accounts:  []Account{},
			UpdatedAt: time.Now(),
			db:        *db,
		}, nil
	}

	var vault VaultWithDB

	err = json.Unmarshal(file, &vault)
	if err != nil {
		color.Red("%s", "Error reading file: ", err)
		return &VaultWithDB{
			Accounts:  []Account{},
			UpdatedAt: time.Now(),
			db:        *db,
		}, nil
	}

	return &VaultWithDB{
		Accounts:  vault.Accounts,
		UpdatedAt: vault.UpdatedAt,
		db:        *db,
	}, err
}

func (vault *VaultWithDB) AddAccount(account Account) {
	vault.Accounts = append(vault.Accounts, account)
	data, err := vault.ToJSONBytes()
	if err != nil {
		color.Red("%s", "Error converting accounts to json: ", err)
	}

	vault.db.Write(data)
}

func (vault *VaultWithDB) FindAccountByUrl(urlString string) []Account {
	var accounts = make([]Account, 2)

	for _, account := range vault.Accounts {
		if strings.Contains(account.Url, urlString) {
			accounts = append(accounts, account)
		}
	}

	return accounts
}

func (vault *VaultWithDB) DeleteByUrl(url string) []Account {
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
	vault.db.Write(data)

	return deletedAccounts
}

func (vault *Vault) ToJSONBytes() ([]byte, error) {
	file, err := json.Marshal(vault)
	if err != nil {
		return nil, err
	}

	return file, nil
}
