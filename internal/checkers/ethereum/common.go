package ethereum

import (
	"context"
	"crypto/ecdsa"
	"encoding/hex"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"math/big"
)

type NetworkConfig struct {
	Name   string
	Symbol string
}

type CheckResult struct {
	Address common.Address
	Balance *big.Int
}

type Checker struct {
	Name    string
	Symbol  string
	NodeURL string

	wei uint64
}

func NewChecker(networkConfig NetworkConfig, nodeURL string) *Checker {
	return &Checker{
		Name:    networkConfig.Name,
		Symbol:  networkConfig.Symbol,
		NodeURL: nodeURL,
		wei:     1000000000000000000,
	}
}

func (c Checker) Check(secretKey string) (CheckResult, error) {
	client, err := ethclient.Dial(c.NodeURL)
	if err != nil {
		return CheckResult{}, err
	}
	defer client.Close()

	keyBytes, err := hex.DecodeString(secretKey)
	if err != nil {
		return CheckResult{}, err
	}

	privateKey, err := crypto.ToECDSA(keyBytes)
	if err != nil {
		return CheckResult{}, err
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return CheckResult{}, err
	}

	address := crypto.PubkeyToAddress(*publicKeyECDSA)

	balance, err := client.BalanceAt(context.Background(), common.HexToAddress(address.Hex()), nil)
	if err != nil {
		return CheckResult{}, err
	}

	res := CheckResult{
		Address: address,
		Balance: balance,
	}

	return res, nil
}

func (c Checker) GetName() string {
	return c.Name
}

func (c Checker) GetSymbol() string {
	return c.Symbol
}

func (c Checker) GetWei() uint64 {
	return c.wei
}
