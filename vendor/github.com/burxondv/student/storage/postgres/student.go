package postgres

import (
	"fmt"

	"github.com/burxondv/student/storage/repo"
	"github.com/jmoiron/sqlx"
)

type studentRepo struct {
	db *sqlx.DB
}

func NewStudent(db *sqlx.DB) repo.StudentStorageI {
	return &studentRepo{
		db: db,
	}
}

func (sr *studentRepo) Create(student []*repo.Student) error {
	for _, s := range student {
		query := `
			INSERT INTO student(
				first_name,
				last_name,
				username,
				email,
				phone_number
			) VALUES ($1, $2, $3, $4, $5)
		`
		_, err := sr.db.Exec(
			query,
			s.FirstName,
			s.LastName,
			s.Username,
			s.Email,
			s.PhoneNumber,
		)

		if err != nil {
			return err
		}
	}

	return nil
}

func (sr *studentRepo) Get(params *repo.GetStudentParam) (*repo.GetStudentResult, error) {
	result := repo.GetStudentResult{
		Students: make([]*repo.Student, 0),
	}

	offset := (params.Page - 1) * params.Limit

	limit := fmt.Sprintf(" LIMIT %d OFFSET %d ", params.Limit, offset)

	filter := " WHERE true "
	if params.Search != "" {
		filter += " AND username ilike '%" + params.Search + "%' "
	}

	orderBy := "ORDER BY created_at desc "
	if params.SortByData != "" {
		orderBy = fmt.Sprintf(" ORDER BY created_at %s", params.SortByData)
	}

	query := `
		SELECT
			id,
			first_name,
			last_name,
			username,
			email,
			phone_number,
			created_at
		FROM student
		` + filter + orderBy + limit

	rows, err := sr.db.Query(query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var std repo.Student
		err := rows.Scan(
			&std.ID,
			&std.FirstName,
			&std.LastName,
			&std.Username,
			&std.Email,
			&std.PhoneNumber,
			&std.CreatedAt,
		)
		if err != nil {
			return nil, err
		}
		result.Students = append(result.Students, &std)
	}

	queryCount := `SELECT count(*) FROM student` + filter
	err = sr.db.QueryRow(queryCount).Scan(&result.Count)
	if err != nil {
		return nil, err
	}

	return &result, nil
}
