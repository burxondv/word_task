package storage

import (
	"github.com/burxondv/student/storage/postgres"
	"github.com/burxondv/student/storage/repo"
	"github.com/jmoiron/sqlx"
)

type StorageI interface {
	Student() repo.StudentStorageI
}

type storagePg struct {
	studentRepo repo.StudentStorageI
}

func NewStoragePg(db *sqlx.DB) StorageI {
	return &storagePg{
		studentRepo: postgres.NewStudent(db),
	}
}

func (s *storagePg) Student() repo.StudentStorageI {
	return s.studentRepo
}
