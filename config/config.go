package config

var RunMode string
var ServerPort string

var KafkaBroker string
var Topic string

func InitEnvironmentVariables() {
	ServerPort ="8080"

	KafkaBroker = "my-cluster-kafka-bootstrap:9092"
	Topic = "kt-test-event"
}
