package postgres

import (
	"fmt"

	"github.com/burxondv/word_task/storage/repo"

	"github.com/jmoiron/sqlx"
)

type wordRepo struct {
	db *sqlx.DB
}

func NewWord(db *sqlx.DB) repo.WordStorageI {
	return &wordRepo{
		db: db,
	}
}

func (wr *wordRepo) Create(body map[string]int) error {
	query := `
		INSERT INTO words_task (
			word,
			point
		) VALUES($1, $2)
	`
	for key, value := range body {
		_, err := wr.db.Exec(query, key, value)

		if err != nil {
			return err
		}
	}

	return nil
}

func (wr *wordRepo) GetAll(params *repo.GetWordParam) (*repo.GetWordResult, error) {
	result := repo.GetWordResult{
		Words: make([]*repo.GetWord, 0),
	}

	offset := (params.Page - 1) * params.Limit

	limit := fmt.Sprintf(" LIMIT %d OFFSET %d", params.Limit, offset)

	filter := ""
	if params.Search != "" {
		filter += " WHERE word ilike '%" + params.Search + "%' "
	}

	query := `
		SELECT
			word,
			point
		FROM words_task
		` + filter + `
		ORDER BY point desc
		` + limit

	rows, err := wr.db.Query(query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var wd repo.GetWord
		err := rows.Scan(
			&wd.Word,
			&wd.Point,
		)
		if err != nil {
			return nil, err
		}

		result.Words = append(result.Words, &wd)
	}

	queryCount := `SELECT count(*) FROM words_task` + filter
	err = wr.db.QueryRow(queryCount).Scan(&result.Count)
	if err != nil {
		return nil, err
	}

	return &result, nil
}
