package checkers

import (
	"log"
	"sync"
	"wallet-checker/internal/checkers/ethereum"
	"wallet-checker/internal/config"
)

var (
	binance = ethereum.NewChecker(ethereum.NetworkConfig{
		Name:   "Binance",
		Symbol: "bnb",
	}, config.Endpoints["binance"])
	eth = ethereum.NewChecker(ethereum.NetworkConfig{
		Name:   "Ethereum",
		Symbol: "eth",
	}, config.Endpoints["ethereum"])
	polygon = ethereum.NewChecker(ethereum.NetworkConfig{
		Name:   "Polygon",
		Symbol: "matic",
	}, config.Endpoints["polygon"])
	arbitrum = ethereum.NewChecker(ethereum.NetworkConfig{
		Name:   "Arbitrum",
		Symbol: "eth",
	}, config.Endpoints["arbitrum"])

	networks = []Interface{
		binance,
		eth,
		polygon,
		arbitrum,
	}
)

func CheckWalletBySecretKey(secretKey string) {
	var wg sync.WaitGroup

	for _, network := range networks {
		wg.Add(1)
		go func(secretKey string, network Interface) {
			result, err := network.Check(secretKey)
			if err != nil {
				log.Println(err)
			}

			log.Printf("\n------------------\nNetwork: %s\nWallet Address: %s\nBalance (in Wei): %d wei\nBalance: %f %s\n------------------\n\n", network.GetName(), result.Address.String(), result.Balance, float64(result.Balance.Uint64())/float64(network.GetWei()), network.GetSymbol())

			wg.Done()
		}(secretKey, network)
	}

	wg.Wait()
}
