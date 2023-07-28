package models

import (
	"testing"
)

func TestInsertSubscription(t *testing.T) {
	type testCase struct {
		server, topic string
		want          int
		err           bool
	}

	testCases := []testCase{
		{server: "localhost:1883", topic: "test/topic", want: 1, err: false},
		{server: "localhost:1883", topic: "test/topic", want: 0, err: true},
		{server: "localhost:1884", topic: "test/topic", want: 2, err: false},
	}

	for _, tc := range testCases {
		s := &SubscriptionModel{DB: db}
		got, err := s.Insert(tc.server, tc.topic)
		if tc.err && err == nil {
			t.Errorf("Insert(): expected error, but got none")
		} else if got != tc.want {
			t.Errorf("Insert(): wanted %d, got %d", tc.want, got)
		}

	}

}

func TestSetOnlineStatus(t *testing.T) {
	type testCase struct {
		id, rowsAffected int
		online           bool
		err              bool
	}

	testCases := []testCase{
		{id: 1, online: true, rowsAffected: 1, err: false},
		{id: 1, online: false, rowsAffected: 1, err: false},
	}
	s := &SubscriptionModel{DB: db}
	id, _ := s.Insert("localhost:1885", "some/topic")

	for _, tc := range testCases {
		rows, err := s.SetOnlineStatus(id, tc.online)
		if tc.err && err == nil {
			t.Errorf("SetOnlineStatus(): expected error, but got none")
		}
		if tc.rowsAffected != rows {
			t.Errorf("SetOnlineStatus() rows affected: expected %d, got %d\n", tc.rowsAffected, rows)
		}
	}
}

func TestGetOnlineStatus(t *testing.T) {
	type testCase struct {
		setTo bool
		err   bool
	}

	testCases := []testCase{
		{setTo: true, err: false},
		{setTo: false, err: false},
	}
	s := &SubscriptionModel{DB: db}
	id, _ := s.Insert("localhost:1883", "some/topic")

	for _, tc := range testCases {
		s.SetOnlineStatus(id, tc.setTo)
		status, err := s.GetOnlineStatus(id)
		if tc.err && err == nil {
			t.Errorf("GetOnlineStatus(): expected error, but got: %s", err)
		}
		if status != tc.setTo {
			t.Errorf("GetOnlineStatus(): expected %t, got %t\n", tc.setTo, status)
		}
	}
}
