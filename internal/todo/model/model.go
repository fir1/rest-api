package model

import "time"

type Status int

const (
	StatusPending Status = iota + 1
	StatusInProgress
	StatusDone
)

func (s Status) IsValid() bool {
	switch s {
	case StatusPending:
		return true
	case StatusInProgress:
		return true
	case StatusDone:
		return true
	}
	return false
}

type ToDo struct {
	ID          int        `db:"id"`
	Name        string     `db:"name"`
	Description string     `db:"description"`
	Status      Status     `db:"status"`
	CreatedOn   time.Time  `db:"created_on"`
	UpdatedOn   *time.Time `db:"updated_on"`
	DeletedOn   *time.Time `db:"deleted_on"`
}
