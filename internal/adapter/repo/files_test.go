package repo

import (
	"github.com/dgraph-io/badger/v4"
	"github.com/google/uuid"
	"github.com/meteormin/friday.go/internal/domain"
	"github.com/meteormin/friday.go/pkg/database"
	"github.com/stretchr/testify/assert"
	"math/rand"
	"testing"
	"time"
)

func setupBadger(t *testing.T) *badger.DB {
	t.Helper()
	storage, err := database.NewBadger(database.BadgerConfig{
		InMemory: true,
	})

	assert.NotNil(t, storage)
	assert.Nil(t, err)

	return storage
}

func generatePseudoRandomBytes(size int) []byte {
	rand.NewSource(time.Now().UnixNano()) // ✅ 시드 설정 (매 실행마다 다른 값)
	bytes := make([]byte, size)
	for i := range bytes {
		bytes[i] = byte(rand.Intn(256)) // 0~255 랜덤 값
	}
	return bytes
}

func TestFileRepositoryImpl_CreateFile(t *testing.T) {

	assert.NotNil(t, db)

	tx := db.Begin()
	repo := NewFileRepository("test", tx, setupBadger(t))

	newUUID, err := uuid.NewUUID()

	assert.Nil(t, err)

	_, err = repo.CreateFile(&domain.File{
		OriginName: "test",
		ConvName:   newUUID.String(),
		Size:       1024,
		MimeType:   "text/plain",
		FilePath:   "tmp",
	}, generatePseudoRandomBytes(1024))

	if err != nil {
		tx.Rollback()
		assert.Error(t, err)
	}

	tx.Rollback()
}

func TestFileRepositoryImpl_FindFile(t *testing.T) {
	assert.NotNil(t, db)

	tx := db.Begin()
	repo := NewFileRepository("test", tx, setupBadger(t))

	newUUID, err := uuid.NewUUID()

	assert.Nil(t, err)

	created, err := repo.CreateFile(&domain.File{
		OriginName: "test",
		ConvName:   newUUID.String(),
		Size:       1024,
		MimeType:   "text/plain",
		FilePath:   "tmp",
	}, generatePseudoRandomBytes(1024))

	if err != nil {
		tx.Rollback()
		assert.Error(t, err)
	}

	_, file, err := repo.FindFile(created.ID)

	if err != nil {
		tx.Rollback()
		assert.Error(t, err)
	}

	assert.NotNil(t, file)

	assert.Equal(t, "test", file.OriginName)
	assert.Equal(t, newUUID.String(), file.ConvName)
	assert.Equal(t, 1024, int(file.Size))
	assert.Equal(t, "text/plain", file.MimeType)
	assert.Equal(t, "tmp", file.FilePath)

	tx.Rollback()
}
