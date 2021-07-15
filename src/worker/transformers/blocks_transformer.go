package transformers

import (
	"github.com/geometry-labs/icon-blocks/config"
	"github.com/geometry-labs/icon-blocks/global"
	"github.com/geometry-labs/icon-blocks/kafka"
	"github.com/geometry-labs/icon-blocks/models"
	"github.com/geometry-labs/icon-blocks/worker/utils"
	"go.uber.org/zap"
	"gopkg.in/Shopify/sarama.v1"
)

func StartBlocksTransformer() {
	go blocksTransformer()
}

func blocksTransformer() {
	consumer_topic_name := "transactions"
	producer_topic_name := "transactions-ws"

	// TODO: Need to move all of the config validations to config.go
	// Check topic names
	if utils.StringInSlice(consumer_topic_name, config.Config.ConsumerTopics) == false {
		zap.S().Panic("Transactions Worker: no ", consumer_topic_name, " topic found in CONSUMER_TOPICS=", config.Config.ConsumerTopics)
	}
	if utils.StringInSlice(producer_topic_name, config.Config.ProducerTopics) == false {
		zap.S().Panic("Transactions Worker: no ", producer_topic_name, " topic found in PRODUCER_TOPICS=", config.Config.ConsumerTopics)
	}

	consumer_topic_chan := make(chan *sarama.ConsumerMessage)
	producer_topic_chan := kafka.KafkaTopicProducers[producer_topic_name].TopicChan
	postgresLoaderChan := global.GetGlobal().Transactions.GetWriteChan()

	// Register consumer channel
	broadcaster_output_chan_id := kafka.Broadcasters[consumer_topic_name].AddBroadcastChannel(consumer_topic_chan)
	defer func() {
		kafka.Broadcasters[consumer_topic_name].RemoveBroadcastChannel(broadcaster_output_chan_id)
	}()

	zap.S().Debug("Transactions Worker: started working")
	for {
		// Read from kafka
		consumer_topic_msg := <-consumer_topic_chan
		blockRaw, err := models.ConvertToTransactionRaw(consumer_topic_msg.Value)
		if err != nil {
			zap.S().Error("Transactions Worker: Unable to proceed cannot convert kafka msg value to Block")
		}

		// Transform logic
		transformedBlock, _ := transform(blockRaw)

		// Produce to Kafka
		producer_topic_msg := &sarama.ProducerMessage{
			Topic: producer_topic_name,
			Key:   sarama.ByteEncoder(consumer_topic_msg.Key),
			Value: sarama.ByteEncoder(consumer_topic_msg.Value),
		}

		producer_topic_chan <- producer_topic_msg

		// Load to Postgres
		postgresLoaderChan <- transformedBlock

		zap.S().Debug("Transactions worker: last seen block #", string(consumer_topic_msg.Key))
	}
}

// Business logic goes here
func transform(txRaw *models.TransactionRaw) (*models.Transaction, error) {
	//time.Sleep(time.Minute)
	return &models.Transaction{
		Type:                      txRaw.Type,
		Version:                   txRaw.Version,
		FromAddress:               txRaw.FromAddress,
		ToAddress:                 txRaw.ToAddress,
		Value:                     txRaw.Value,
		StepLimit:                 txRaw.StepLimit,
		Timestamp:                 txRaw.Timestamp,
		BlockTimestamp:            txRaw.BlockTimestamp,
		Nid:                       txRaw.Nid,
		Nonce:                     txRaw.Nonce,
		Hash:                      txRaw.Hash,
		TransactionIndex:          txRaw.TransactionIndex,
		BlockHash:                 txRaw.BlockHash,
		BlockNumber:               txRaw.BlockNumber,
		Fee:                       txRaw.Fee,
		Signature:                 txRaw.Signature,
		DataType:                  txRaw.DataType,
		Data:                      txRaw.Data,
		ReceiptCumulativeStepUsed: txRaw.ReceiptCumulativeStepUsed,
		ReceiptStepUsed:           txRaw.ReceiptStepUsed,
		ReceiptStepPrice:          txRaw.ReceiptStepPrice,
		ReceiptScoreAddress:       txRaw.ReceiptScoreAddress,
		ReceiptLogs:               txRaw.ReceiptLogs,
		ReceiptStatus:             txRaw.ReceiptStatus,
		ItemId:                    txRaw.ItemId,
		ItemTimestamp:             txRaw.ItemTimestamp,
	}, nil
}
