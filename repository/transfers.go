package repository

import (
	"database/sql"
	"log"

	"github.com/eliasfeijo/desafio-imersao/database"
	"github.com/eliasfeijo/desafio-imersao/model"
)

type TransfersRepository interface {
	CreateTransfer(from int64, to int64, amount float64) (int64, error)
	FindTransferById(id int64) (*model.Transfer, error)
}

type transfersRepository struct {
	db *sql.DB
}

var instanceTransfers TransfersRepository

func NewTransfers() TransfersRepository {
	if instanceTransfers == nil {
		db := database.GetConn()
		i := &transfersRepository{db: db}
		instanceTransfers = i
	}
	return instanceTransfers
}

func (repository *transfersRepository) CreateTransfer(fromId int64, toId int64, amount float64) (int64, error) {
	stmt, err := repository.db.Prepare("INSERT INTO transfers (from_id, to_id, amount) VALUES (?, ?, ?)")
	if err != nil {
		log.Fatalf("Error preparing insert query: %v", err)
		return 0, err
	}
	defer stmt.Close()

	result, err := stmt.Exec(fromId, toId, amount)
	if err != nil {
		log.Fatal(err)
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		log.Fatal(err)
		return 0, err
	}
	return id, nil
}

func (repository *transfersRepository) FindTransferById(id int64) (*model.Transfer, error) {
	stmt, err := repository.db.Prepare("SELECT id, from_id, to_id, amount, created_at FROM transfers WHERE id = ?")
	if err != nil {
		log.Fatalf("Error preparing select query: %v", err)
		return nil, err
	}
	defer stmt.Close()

	transfer := model.Transfer{}
	err = stmt.QueryRow(id).Scan(&transfer.ID, &transfer.FromId, &transfer.ToId, &transfer.Amount, &transfer.CreatedAt)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	return &transfer, nil
}
