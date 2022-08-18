package service

import (
	"github.com/eliasfeijo/desafio-imersao/model"
	"github.com/eliasfeijo/desafio-imersao/repository"
)

type Transfers interface {
	CreateTransfer(fromId int64, toId int64, amount float64) (*model.Transfer, error)
}

type transfers struct {
	repository repository.TransfersRepository
}

func NewTransfers(repository repository.TransfersRepository) Transfers {
	return &transfers{repository: repository}
}

func (t transfers) CreateTransfer(fromId int64, toId int64, amount float64) (*model.Transfer, error) {
	id, err := t.repository.CreateTransfer(fromId, toId, amount)
	if err != nil {
		return nil, err
	}
	transfer, err := t.repository.FindTransferById(id)
	if err != nil {
		return nil, err
	}
	return transfer, nil
}
