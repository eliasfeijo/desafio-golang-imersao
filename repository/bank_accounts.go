package repository

import (
	"database/sql"
	"log"

	"github.com/eliasfeijo/desafio-imersao/database"
	"github.com/eliasfeijo/desafio-imersao/model"
)

type BankAccountsRepository interface {
	CreateBankAccount(number string) (*model.BankAccount, error)
}

type bankAccountsRepository struct {
	db *sql.DB
}

var instance BankAccountsRepository

func NewBankAccounts() BankAccountsRepository {
	if instance == nil {
		db := database.GetConn()
		i := &bankAccountsRepository{db: db}
		instance = i
	}
	return instance
}

func (repository *bankAccountsRepository) CreateBankAccount(number string) (*model.BankAccount, error) {
	stmt, err := repository.db.Prepare("INSERT INTO bank_accounts (number) VALUES (?)")
	if err != nil {
		log.Fatalf("Error preparing insert query: %v", err)
		return nil, err
	}
	defer stmt.Close()

	result, err := stmt.Exec(number)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	id, _ := result.LastInsertId()
	stmt, err = repository.db.Prepare("SELECT id, number, created_at FROM bank_accounts WHERE id = ?")
	if err != nil {
		log.Fatalf("Error preparing select query: %v", err)
		return nil, err
	}
	defer stmt.Close()

	bankAccount := model.BankAccount{}
	err = stmt.QueryRow(id).Scan(&bankAccount.ID, &bankAccount.Number, &bankAccount.CreatedAt)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	return &bankAccount, nil
}
