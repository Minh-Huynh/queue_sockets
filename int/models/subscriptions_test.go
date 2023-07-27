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
