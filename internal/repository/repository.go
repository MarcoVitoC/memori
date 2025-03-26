package repository

import "github.com/jackc/pgx/v5/pgxpool"

type Repository struct {
	Diary interface {
		Create(diary *Diary) error
	}
}

func NewRepository(db *pgxpool.Pool) Repository {
	return Repository{
		Diary: &DiaryRepository{db},
	}
}