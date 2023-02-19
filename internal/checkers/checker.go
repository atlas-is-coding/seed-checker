package checkers

import (
	"wallet-checker/internal/checkers/ethereum"
)

type Interface interface {
	Check(secretKey string) (ethereum.CheckResult, error)
	GetName() string
	GetSymbol() string
	GetWei() uint64
}
