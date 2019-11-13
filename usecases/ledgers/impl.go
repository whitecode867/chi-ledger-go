package ledgers

import "chi-ledger-go/standard"

type Ledgers interface {
}

type ledgersImpl struct {
	LedgersMongoDBRepository standard.DatabaseRepository
}

type LedgersRepositories struct {
	LedgersMongoDBRepository standard.DatabaseRepository
}

func NewLedgersUseCase(repo LedgersRepositories) Ledgers {
	return &ledgersImpl{
		LedgersMongoDBRepository: repo.LedgersMongoDBRepository,
	}
}
