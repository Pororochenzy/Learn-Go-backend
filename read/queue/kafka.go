package queue

import "github.com/Shopify/sarama"

// KafkaQueue kafka消息队列
type KafkaQueue struct {
	producer sarama.SyncProducer
}

func NewQueue(attr []string, config *sarama.Config) (*KafkaQueue, error) {

	client, err := sarama.NewSyncProducer(attr, config)

	if err != nil {
		return nil, err
	}

	return &KafkaQueue{
		producer: client,
	}, nil
}

func (p *KafkaQueue) SendMsg(msg *sarama.ProducerMessage) error {
	_, _, err := p.producer.SendMessage(msg)
	return err
}

func (p *KafkaQueue) Close() {
	p.producer.Close()
}
