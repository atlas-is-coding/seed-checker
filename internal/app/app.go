package app

import (
	"errors"
	"path"
	"strings"
	"sync"
	"wallet-checker/internal/checkers"
	"wallet-checker/internal/config"
	"wallet-checker/pkg/duplicates"
	"wallet-checker/pkg/utils"
)

type App struct {
	cfg *config.Config
}

func NewApp(cfg *config.Config) (*App, error) {
	exists, err := utils.IsDirectoryExists(cfg.App.WalletsDir)
	if err != nil {
		return nil, err
	}

	if !exists {
		if err := utils.CreateDirectory(cfg.App.WalletsDir); err != nil {
			return nil, err
		}

		exists := utils.IsFileExists(path.Join(cfg.App.WalletsDir, cfg.App.WalletsFile))
		if !exists {
			if err := utils.CreateFile(path.Join(cfg.App.WalletsDir, cfg.App.WalletsFile)); err != nil {
				return nil, err
			}

			return nil, errors.New("wallets file doesn`t exists. It was created but you should fill it with private keys")
		}
	}

	exists = utils.IsFileExists(path.Join(cfg.App.WalletsDir, cfg.App.WalletsFile))
	if !exists {
		if err := utils.CreateFile(path.Join(cfg.App.WalletsDir, cfg.App.WalletsFile)); err != nil {
			return nil, err
		}

		return nil, errors.New("wallets file doesn`t exists. It was created but you should fill it with secret keys")
	}

	return &App{
		cfg: cfg,
	}, nil
}

func (a App) Run() error {
	var wg sync.WaitGroup

	file, err := utils.ReadFile(path.Join(a.cfg.App.WalletsDir, a.cfg.App.WalletsFile))
	if err != nil {
		return err
	}

	secretKeys := string(file)

	for _, secretKey := range duplicates.RemoveFromSlice(strings.Split(secretKeys, "\n")) {
		wg.Add(1)
		go func(secretKey string) {
			checkers.CheckWalletBySecretKey(secretKey)
			wg.Done()
		}(secretKey)
	}

	wg.Wait()
	return nil
}
