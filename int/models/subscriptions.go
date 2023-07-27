package models

import (
	"database/sql"
)

type SubscriptionModel struct {
	DB *sql.DB
}

type Subscription struct {
	id     int
	server string
	topic  string
	online bool
}

func (s *SubscriptionModel) Insert(server, topic string) (int, error) {
	statement := `INSERT INTO subscriptions (server, topic) VALUES(?, ?);`

	result, err := s.DB.Exec(statement, server, topic)
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), nil
}
