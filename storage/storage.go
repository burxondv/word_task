package storage

import (
	"github.com/burxondv/word_task/storage/postgres"
	"github.com/burxondv/word_task/storage/repo"
	"github.com/jmoiron/sqlx"
)

type StorageI interface {
	Word() repo.WordStorageI
}

type storagePg struct {
	wordRepo repo.WordStorageI
}

func NewStoragePg(db *sqlx.DB) StorageI {
	return &storagePg{
		wordRepo: postgres.NewWord(db),
	}
}

func (s *storagePg) Word() repo.WordStorageI {
	return s.wordRepo
}