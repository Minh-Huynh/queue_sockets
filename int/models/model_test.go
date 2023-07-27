package models

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"testing"

	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

func TestMain(m *testing.M) {
	code, err := run(m)
	if err != nil {
		fmt.Println(err)
	}
	os.Exit(code)
}

func run(m *testing.M) (code int, err error) {
	db, err = sql.Open("sqlite3", "file:test.db?mode=memory")
	if err != nil {
		return -2, fmt.Errorf("could not connect to database: %w", err)
	}

	if err = db.Ping(); err != nil {
		log.Printf("db ping error: %s\n", err)
	}
	CreateSubscriptionTable(db)
	CreateUserSubscriptionTable(db)
	CreateUserTable(db)

	// truncates all test data after the tests are run
	defer func() {
		for _, t := range []string{"user", "subscription", "user_subscription"} {
			_, _ = db.Exec(fmt.Sprintf("DELETE FROM %s", t))
		}

		db.Close()
	}()

	return m.Run(), nil
}

func CreateSubscriptionTable(db *sql.DB) {
	subscriptions_table := `CREATE TABLE subscriptions(
  server varchar(64) NOT NULL,
  topic varchar(64) NOT NULL,
  online boolean DEFAULT false,
  UNIQUE(server, topic)
  );`
	_, err := db.Exec(subscriptions_table)
	if err != nil {
		log.Fatal(err)
	}
}
func CreateUserTable(db *sql.DB) {
	user_table := `CREATE TABLE users (
    name varchar(64) NOT NULL
);`
	_, err := db.Exec(user_table)
	if err != nil {
		log.Fatal(err)
	}
}
func CreateUserSubscriptionTable(db *sql.DB) {
	user_subscriptions_table := `CREATE TABLE user_subscriptions (
    user_id int NOT NULL REFERENCES user(rowid),
    subscription_id int NOT NULL REFERENCES subscription(rowid)
);`
	_, err := db.Exec(user_subscriptions_table)
	if err != nil {
		log.Fatal(err)
	}
	indexCreate := `CREATE UNIQUE INDEX user_subscription ON user_subscriptions(user_id, subscription_id);`
	_, err = db.Exec(indexCreate)
	if err != nil {
		log.Fatal(err)
	}
}
