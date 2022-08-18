package service

import (
	"github.com/eliasfeijo/desafio-imersao/model"
	"github.com/eliasfeijo/desafio-imersao/repository"
)

type Transfers interface {
	CreateTransfer(accountNumberFrom string, accountNumberTo string, amount float64) (*model.Transfer, error)
}

type transfers struct {
	transfersRepository    repository.TransfersRepository
	bankAccountsRepository repository.BankAccountsRepository
}

func NewTransfers(transfersRepository repository.TransfersRepository, bankAccountsRepository repository.BankAccountsRepository) Transfers {
	return &transfers{transfersRepository: transfersRepository, bankAccountsRepository: bankAccountsRepository}
}

func (t transfers) CreateTransfer(accountNumberFrom string, accountNumberTo string, amount float64) (*model.Transfer, error) {

	from, err := t.bankAccountsRepository.FindBankAccountByNumber(accountNumberFrom)
	if err != nil {
		return nil, err
	}

	to, err := t.bankAccountsRepository.FindBankAccountByNumber(accountNumberTo)
	if err != nil {
		return nil, err
	}

	id, err := t.transfersRepository.CreateTransfer(from.ID, to.ID, amount)
	if err != nil {
		return nil, err
	}

	transfer, err := t.transfersRepository.FindTransferById(id)
	if err != nil {
		return nil, err
	}
	return transfer, nil
}
