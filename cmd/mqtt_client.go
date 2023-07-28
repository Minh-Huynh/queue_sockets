package main

import (
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/Minh-Huynh/queue_sockets/int/models"
	mqtt "github.com/eclipse/paho.mqtt.golang"
)

type BrokerMessage struct {
	brokerId int
	payload  []byte
}

type MqttClient struct {
	id              int
	server          string
	topic           string
	client          mqtt.Client
	outgoingChannel chan BrokerMessage
	subscription    models.SubscriptionModel
	done            chan struct{}
}

func (m MqttClient) SetUp(username, password string) mqtt.Client {
	opts := mqtt.NewClientOptions()
	opts.AddBroker(m.server)
	opts.SetClientID(strconv.FormatInt(time.Now().UTC().UnixNano(), 10))
	opts.SetKeepAlive(2 * time.Second)
	opts.SetUsername(username)
	opts.SetPassword(password)
	opts.SetDefaultPublishHandler(f)
	opts.SetAutoReconnect(true)
	opts.SetDefaultPublishHandler(m.messagePubHandler)
	opts.OnConnect = m.connectHandler
	opts.OnConnectionLost = m.connectLostHandler
	return mqtt.NewClient(opts)
}

func (m *MqttClient) Connect() {
	if token := m.client.Connect(); token.Wait() && token.Error() != nil {
		log.Printf("MQTT Connect error: %s\n", token.Error())
	}

	if token := m.client.Subscribe(m.topic, 0, nil); token.Wait() && token.Error() != nil {
		log.Printf("MQTT Subscribe error: %s\n", token.Error())
	}

	id, err := m.subscription.Insert(m.server, m.topic)
	if err != nil {
		log.Printf("Error saving subscription to database: %s\n", err)
	}
	log.Printf("Setting %s:%s's ID to %d\n", m.server, m.topic, id)
	m.id = id
	<-m.done
}

func (m MqttClient) messagePubHandler(client mqtt.Client, msg mqtt.Message) {
	log.Printf("Received message from topic %s\n %s\n\n", msg.Topic(), msg.Payload()[:400])
	log.Printf("Forwarding to multiplexer via publish channel\n")
	m.outgoingChannel <- BrokerMessage{brokerId: m.id, payload: msg.Payload()}

}

func (m *MqttClient) connectHandler(client mqtt.Client) {
	fmt.Printf("Connected to %s:%s\n", m.server, m.topic)
	m.subscription.SetOnlineStatus(m.id, true)
}

func (m MqttClient) connectLostHandler(client mqtt.Client, err error) {
	fmt.Printf("Disconnected from %s:%s\n", m.server, m.topic)
	m.subscription.SetOnlineStatus(m.id, false)
}
