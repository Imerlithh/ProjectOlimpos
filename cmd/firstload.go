package cmd

import (
	"fmt"
	"os"
	"path/filepath"
)

func firstLoad() bool {
	configPath := filepath.Join("config", "config.yaml")

	if _, err := os.Stat(configPath); err == nil {
		return true // config.yaml mevcut
	}

	fmt.Println("Config bulunamadı. İlk kurulum başlatılıyor...")
	setupFirstRunForm()
	return false // config.yaml yok, form sunucusu devraldı
}
