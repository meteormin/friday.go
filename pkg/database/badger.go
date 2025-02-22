package database

import (
	"fmt"
	"github.com/dgraph-io/badger/v4"
)

type BadgerConfig struct {
	Path       string `json:"path" yaml:"path"`
	CacheSize  int64  `json:"cacheSize" yaml:"cacheSize"`
	EncryptKey []byte `json:"encryptKey" yaml:"encryptKey"`
	InMemory   bool   `json:"in-memory" yaml:"in-memory"`
}

func NewBadger(cfg BadgerConfig) (*badger.DB, error) {
	opts := badger.DefaultOptions(cfg.Path).
		WithInMemory(cfg.InMemory).
		WithEncryptionKey(cfg.EncryptKey).
		WithIndexCacheSize(cfg.CacheSize << 20)
	return badger.Open(opts)
}

// PutFile PutFile은 지정한 로컬 파일의 내용을 읽어 key에 해당하는 값으로 Badger DB에 저장합니다.
func PutFile(db *badger.DB, key string, data []byte) error {
	// Badger 트랜잭션을 통해 데이터 저장
	err := db.Update(func(txn *badger.Txn) error {
		return txn.Set([]byte(key), data)
	})
	if err != nil {
		return fmt.Errorf("failed save badger: %w", err)
	}
	return nil
}

// GetFile GetFile은 Badger DB에서 지정된 key의 값을 읽어 로컬 파일로 저장합니다.
func GetFile(db *badger.DB, key string, destPath string) ([]byte, error) {
	var data []byte
	err := db.View(func(txn *badger.Txn) error {
		item, err := txn.Get([]byte(key))
		if err != nil {
			return fmt.Errorf("failed get [%s] key: %w", key, err)
		}
		data, err = item.ValueCopy(nil)
		return err
	})
	if err != nil {
		return nil, err
	}

	return data, nil
}

// DeleteFile DeleteFile은 Badger DB에서 지정된 key를 삭제합니다.
func DeleteFile(db *badger.DB, key string) error {
	err := db.Update(func(txn *badger.Txn) error {
		return txn.Delete([]byte(key))
	})
	if err != nil {
		return fmt.Errorf("키 [%s] 삭제 실패: %w", key, err)
	}
	return nil
}
