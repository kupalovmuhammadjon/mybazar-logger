package logger

import (
	"time"

	rabbitmq "github.com/kupalovmuhammadjon/rabbitmq-go"
)

// logger is the implementation of the Logger interface.
// It interacts with RabbitMQ to publish log messages to a specified queue.
type logger struct {
	rabbitmq         rabbitmq.RabbitMQ // RabbitMQ client for managing messages.
	queue            string            // Name of the RabbitMQ queue where logs will be sent.
	orderQueue       string            // Name of the RabbitMQ queue where logs will be sent.
	bitrixOrderQueue string            // Name of the RabbitMQ queue where logs will be sent.
	functionName     string            // Name of the function generating logs.
	apiEndpoint      string            // API endpoint associated with the logs.
}

// logRequest represents the structure of a log message sent to RabbitMQ.
// It includes metadata such as error level, error messages, API endpoint, and other details.
type logRequest struct {
	Timestamp       time.Time `json:"timestamp"`
	ErrorLevel      string    `json:"error_level"`
	Errorcode       int       `json:"error_code"`
	ClientMessageUz string    `json:"client_message_uz"`
	ClientMessageRu string    `json:"client_message_ru"`
	ErrorMessage    string    `json:"error_message"`
	DetailsUz       string    `json:"details_uz,omitempty"` // Optional details in Uzbek.
	DetailsRu       string    `json:"details_ru,omitempty"` // Optional details in Russian.
	ApiEndpoint     string    `json:"api_endpoint"`
	Method          string    `json:"method"`
	FunctionName    string    `json:"function_name"`
	StatusCode      int       `json:"status_code"`
	RequestPayload  string    `json:"request_payload"`
	EventType       string    `json:"event_type"`                 // Event type, usually based on the function name.
	ResponseData    string    `json:"response_data,omitempty"`    // Optional response data.
	MerchantApiKey  string    `json:"merchant_api_key,omitempty"` // Merchant API key, required if sending to merchants.
}

// LogRequest is a simplified structure used by the user to send log data.
// It will be converted into a `logRequest` structure with additional metadata.
type LogRequest struct {
	Errorcode       Errorcode `json:"error_code"`
	ClientMessageUz string    `json:"client_message_uz"`
	ClientMessageRu string    `json:"client_message_ru"`
	ErrorMessage    string    `json:"error_message"`
	DetailsUz       string    `json:"details_uz,omitempty"` // Optional details in Uzbek.
	DetailsRu       string    `json:"details_ru,omitempty"` // Optional details in Russian.
	ApiEndpoint     string    `json:"api_endpoint"`
	Method          string    `json:"method"`
	StatusCode      int       `json:"status_code"`
	RequestPayload  any       `json:"request_payload"`
	EventType       string    `json:"event_type"`                 // Event type, usually based on the function name.
	ResponseData    string    `json:"response_data,omitempty"`    // Optional response data.
	MerchantApiKey  string    `json:"merchant_api_key,omitempty"` // Merchant API key, required if sending to merchants.
}

type Order struct {
	OrderText  string `json:"order_text"`
	MerchantId string `json:"merchant_id"`
}

type BitrixOrder struct {
	OrderIds []string `json:"order_ids"`
}
