package main

import (
	"database/sql"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

type BrokerMessage struct {
	brokerId int
	payload  []byte
}

type MqttClient struct {
	mqtt.Client
	outgoingChannel chan BrokerMessage
	db              *sql.DB
	online          bool
}

func (m *MqttClient) SetUp(address string, clientId string, username, password string) (int, error) {
	//create subscription in DB if none exists
	//retrieve the id of the subscription
	return 0, nil

}

func (m MqttClient) Connect() (int, error) {
	//try to connect
	//set the 'online' column in DB if connect succeeds
	return 0, nil
}
