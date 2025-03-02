package config

import (
	"bytes"
	"crypto/rand"
	"fmt"
	"github.com/spf13/viper"
	"log"
	"os"
	"path"
)

func LoadWithViper(in string, appInfo App) *Config {

	if _, err := os.Stat(in); err == nil {
		viper.SetConfigFile(in)
		if err = viper.ReadInConfig(); err != nil {
			log.Fatalf("Error reading config file, %s", err)
		}
	} else {
		err = viper.ReadConfig(bytes.NewBufferString(in))
		if err != nil {
			log.Fatalf("Error reading config file, %s", err)
		}
	}

	var cfg Config
	if err := viper.Unmarshal(&cfg); err != nil {
		log.Fatalf("Unable to decode into struct, %v", err)
	}

	if cfg.App.Name == "" {
		cfg.App.Name = appInfo.Name
	}

	if cfg.App.Version == "" {
		cfg.App.Version = appInfo.Version
	}

	if cfg.InMemoryMode {
		cfg.Logging.FilePath = ""
		cfg.Database.Path = ""
		cfg.Badger.InMemory = true
		cfg.Badger.Path = ""

		key, err := randomSecretKey()
		if err != nil {
			log.Fatal(err)
		}

		cfg.Badger.EncryptKey = key
	} else {
		if err := cfg.Path.mkdir(); err != nil {
			log.Fatal(err)
		}

		key, err := generateEncryptionKey(cfg.Path.Secret)
		if err != nil {
			log.Fatal(err)
		}

		cfg.Badger.EncryptKey = key
	}

	return &cfg
}

func randomSecretKey() ([]byte, error) {
	key := make([]byte, 32) // AES-256에 필요한 32바이트 키
	_, err := rand.Read(key)
	return key, err
}

// generateEncryptionKey 32바이트(256비트)의 랜덤 암호화 키를 생성합니다.
func generateEncryptionKey(secretPath string) ([]byte, error) {
	secretFile := path.Join(secretPath, ".secret")
	if _, err := os.Stat(secretFile); err == nil {
		key, err := os.ReadFile(secretFile)
		if err == nil {
			return key, nil
		}
	}

	key, err := randomSecretKey()
	if err != nil {
		return nil, fmt.Errorf("failed generating encrypt ley: %w", err)
	}

	err = os.MkdirAll(secretPath, 0755)
	if err != nil {
		return nil, err
	}

	err = os.WriteFile(secretFile, key, 0644)
	if err != nil {
		return nil, err
	}

	return key, nil
}
