package checkers

import (
	"log"
	"sync"
	"wallet-checker/internal/checkers/ethereum"
	"wallet-checker/internal/config"
)

type CheckerInterface interface {
	Check(secretKey string) (ethereum.CheckResult, error)
	GetName() string
	GetSymbol() string
	GetWei() uint64
}

var (
	binance = ethereum.NewChecker(ethereum.NetworkConfig{
		Name:   "Binance",
		Symbol: "$BNB",
	}, config.Endpoints["binance"])
	eth = ethereum.NewChecker(ethereum.NetworkConfig{
		Name:   "Ethereum",
		Symbol: "$ETH",
	}, config.Endpoints["ethereum"])
	polygon = ethereum.NewChecker(ethereum.NetworkConfig{
		Name:   "Polygon",
		Symbol: "$MATIC",
	}, config.Endpoints["polygon"])
	arbitrum = ethereum.NewChecker(ethereum.NetworkConfig{
		Name:   "Arbitrum",
		Symbol: "$ETH",
	}, config.Endpoints["arbitrum"])
	fantom = ethereum.NewChecker(ethereum.NetworkConfig{
		Name:   "Fantom",
		Symbol: "$FTM",
	}, config.Endpoints["fantom"])
	optimism = ethereum.NewChecker(ethereum.NetworkConfig{
		Name:   "Optimism",
		Symbol: "$ETH-OPT",
	}, config.Endpoints["optimism"])
	avalance = ethereum.NewChecker(ethereum.NetworkConfig{
		Name:   "Avalance",
		Symbol: "$AVAX",
	}, config.Endpoints["avalance"])

	networks = []CheckerInterface{
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
		go func(secretKey string, network CheckerInterface) {
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
