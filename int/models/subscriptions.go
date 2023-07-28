package models

import (
	"database/sql"
	"log"
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
	statement := `INSERT INTO subscriptions (server, topic) SELECT ?, ? WHERE NOT EXISTS
	(SELECT server, topic FROM subscriptions WHERE server=? AND topic=?);`

	_, err := s.DB.Exec(statement, server, topic, server, topic)
	if err != nil {
		log.Printf("ERROR: %s", err)
		return 0, err
	}

	getIdStatement := `SELECT rowId from subscriptions WHERE server=? AND topic=?`

	var id int
	row := s.DB.QueryRow(getIdStatement, server, topic)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return int(id), nil
}

func (s *SubscriptionModel) SetOnlineStatus(id int, value bool) (int, error) {
	statement := `UPDATE subscriptions SET online = ? WHERE rowid=?;`
	var status int
	if value == true {
		status = 1
	} else {
		status = 0
	}
	result, err := s.DB.Exec(statement, status, id)
	if err != nil {
		return 0, err
	}
	rows, err := result.RowsAffected()
	if err != nil {
		return 0, err
	}
	return int(rows), nil
}

func (s *SubscriptionModel) GetOnlineStatus(id int) (bool, error) {
	statement := `SELECT online from subscriptions WHERE rowid=? ;`
	row := s.DB.QueryRow(statement, id)
	var status bool
	if err := row.Scan(&status); err != nil {
		return false, err
	}
	return status, nil
}
