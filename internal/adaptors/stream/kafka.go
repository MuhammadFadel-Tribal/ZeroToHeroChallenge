package stream

import (
	"encoding/json"
	"fmt"
	"time"
	"zerotoherochallenge/config"
	"zerotoherochallenge/internal/models"

	"github.com/segmentio/kafka-go"
	"go.uber.org/zap"
	"golang.org/x/net/context"
)

func NewKafkaProducer(transaction *models.TransactionModel) {
	conn, _ := kafka.DialLeader(context.Background(), "tcp", "kafka:29092", "topic_transaction", 0)
	conn.SetWriteDeadline(time.Now().Add(time.Second * 10))
	obj, _ := json.Marshal(&transaction)
	conn.WriteMessages(kafka.Message{Value: []byte(obj)})
}

func NewKafkaConsumer(log *zap.SugaredLogger) {
	configurations := config.LoadConfig(log)
	config := kafka.ReaderConfig{
		Brokers:  []string{configurations.Kafka.Broker},
		Topic:    configurations.Kafka.Topic,
		MaxBytes: configurations.Kafka.MaxBytes,
	}
	reader := kafka.NewReader(config)
	for {
		message, error := reader.ReadMessage(context.Background())
		if error != nil {
			log.Fatalf(time.Now().String()+":: Error happened during calling kafka server %v", error)
			continue
		}
		fmt.Println(time.Now().String() + "::message of transaction consumed:: " + string(message.Value))
	}
}
