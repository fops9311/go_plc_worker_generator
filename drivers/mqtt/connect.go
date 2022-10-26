package mqttdriver

import (
	"fmt"
	"log"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

var Publish = make(chan Message)
var SubscribeSubdir string = "metric/#"
var Subsribe = make(chan Message)

var (
	BrokerHost = "192.168.1.37"
	BrokerPort = 1883
	ClientId   = "my-plc-server-pubsub"
	UserName   = "user1"
	Password   = "userpassword1"
)

type Message struct {
	Topic string
	Body  string
}

func Init() {
	var client mqtt.Client
	opts := mqtt.NewClientOptions()
	opts.AddBroker(fmt.Sprintf("tcp://%s:%d", BrokerHost, BrokerPort))
	opts.SetClientID(ClientId)
	opts.SetUsername(UserName)
	opts.SetPassword(Password)

	opts.OnConnect = connectHandler
	opts.OnConnectionLost = func(c mqtt.Client, err error) {
		log.Printf("[error] %v\nTry reconnect in 15 seconds...\n", err)
		connect(client)
	}

	client = mqtt.NewClient(opts)
	connect(client)
	subsribe(client)
	for m := range Publish {
		//log.Printf("topic=%s, message=%s\n", m.Topic, m.Body)
		t := client.Publish(m.Topic, 2, false, m.Body)
		if err := t.Error(); err != nil {
			log.Printf("[error] %v", err)
		}
	}
}

var messagePubHandler mqtt.MessageHandler = func(client mqtt.Client, msg mqtt.Message) {
	fmt.Printf("Received message: %s from topic: %s\n", msg.Payload(), msg.Topic())
}

var connectHandler mqtt.OnConnectHandler = func(client mqtt.Client) {
	fmt.Println("Connected")
}

var connectLostHandler mqtt.ConnectionLostHandler = func(client mqtt.Client, err error) {
	fmt.Printf("Connect lost: %v", err)
}

func connect(client mqtt.Client) {
	for {
		if token := client.Connect(); token.Wait() && token.Error() != nil {
			log.Printf("[error] %v\nTry reconnect in 15 seconds...\n", token.Error())
			<-time.NewTimer(time.Second * 15).C
			continue
		}
		return
	}
}
func subsribe(client mqtt.Client) {
	if token := client.Subscribe(SubscribeSubdir, 2, func(c mqtt.Client, m mqtt.Message) {
		Subsribe <- Message{
			Topic: m.Topic(),
			Body:  string(m.Payload()),
		}
	}); token.Wait() && token.Error() != nil {
		log.Println(token.Error())
	}
	fmt.Println("Subsribed")
}
