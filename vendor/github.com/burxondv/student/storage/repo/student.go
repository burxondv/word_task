package repo

import (
	"time"
)

type Student struct {
	ID          int64
	FirstName   string
	LastName    string
	Username    string
	Email       string
	PhoneNumber string
	CreatedAt   time.Time
}

type GetStudentParam struct {
	Limit      int32
	Page       int32
	Search     string
	SortByData string
}

type GetStudentResult struct {
	Students []*Student
	Count   int32
}

type StudentStorageI interface {
	Create(u []*Student) error
	Get(param *GetStudentParam) (*GetStudentResult, error)
}
