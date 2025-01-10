package kafka

type Kafka struct{}

func (k Kafka) WriteCount(count int64) {
}

func (k Kafka) WriteStatus(endpoint string, statusCode int) {
}
