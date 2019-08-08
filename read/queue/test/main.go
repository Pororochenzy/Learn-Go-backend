package main

import (
	"encoding/json"
	"fmt"
	"queue"
	"time"

	"github.com/Shopify/sarama"
)

var Queue *queue.KafkaQueue

func main() {

	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Partitioner = sarama.NewRandomPartitioner
	config.Producer.Return.Successes = true

	c, err := queue.NewQueue([]string{"172.24.143.242:9092"}, config)
	if err != nil {
		panic(err)
	}
	Queue = c
	defer Queue.Close()

	fmt.Println(Alarm(AlarmMsg{
		IP:           "172.16.20.11",
		SOURCE_AGENT: "网管自动化",
		NAME:         "TEST",
		EventTime:    time.Now(),
		ORG_TYPE:     "1",
		SUMMARY:      "Test",
		SEVERITY:     "0",
	}))
}

type AlarmMsg struct {
	IP           string
	SOURCE_AGENT string
	NAME         string
	EventTime    time.Time `json:"eventTime"`
	ORG_TYPE     string
	SUMMARY      string
	SOURCE_TYPE  string
	SEVERITY     string
}

func Alarm(msg AlarmMsg) error {

	b, err := json.Marshal(msg)
	if err != nil {
		return err
	}

	return Queue.SendMsg(&sarama.ProducerMessage{
		Topic: "alarm_cosie_event",
		Value: sarama.ByteEncoder(b),
	})
}
