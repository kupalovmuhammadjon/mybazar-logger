package logger

import (
	"encoding/json"
	"errors"
	"fmt"
	rabbitmq "mybazar_logger/rabbitMQ"
	"time"

	"github.com/streadway/amqp"
)

// Logger is the main interface for logging operations.
// It provides methods to log messages with different severity levels and manage the RabbitMQ connection.
type Logger interface {
	// Info logs informational messages.
	Info(log LogRequest) error

	// Warn logs warning messages.
	Warn(log LogRequest) error

	// Error logs error messages.
	Error(log LogRequest) error

	// Critical logs critical errors.
	Critical(log LogRequest) error

	/*
	CloseRabbitMQConnection closes the connection to RabbitMQ.
	When you initialize logger service it initializes the connection so it should be closed
	*/ 
	CloseRabbitMQConnection() error
}

// logger is the implementation of the Logger interface.
// It interacts with RabbitMQ to publish log messages to a specified queue.
type logger struct {
	rabbitmq     rabbitmq.RabbitMQ // RabbitMQ client for managing messages.
	queue        string            // Name of the RabbitMQ queue where logs will be sent.
	functionName string            // Name of the function generating logs.
	apiEndpoint  string            // API endpoint associated with the logs.
}

// logRequest represents the structure of a log message sent to RabbitMQ.
// It includes metadata such as error level, error messages, API endpoint, and other details.
type logRequest struct {
	Timestamp       time.Time `json:"timestamp"`
	ErrorLevel      string    `json:"error_level"`
	ErrorCode       int       `json:"error_code"`
	ClientMessageUz string    `json:"client_message_uz"`
	ClientMessageRu string    `json:"client_message_ru"`
	ErrorMessage    string    `json:"error_message"`
	DetailsUz       string    `json:"details_uz,omitempty"` // Optional details in Uzbek.
	DetailsRu       string    `json:"details_ru,omitempty"` // Optional details in Russian.
	ApiEndpoint     string    `json:"api_endpoint"`
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
	ErrorCode       int    `json:"error_code"`
	ClientMessageUz string `json:"client_message_uz"`
	ClientMessageRu string `json:"client_message_ru"`
	ErrorMessage    string `json:"error_message"`
	DetailsUz       string `json:"details_uz,omitempty"` // Optional details in Uzbek.
	DetailsRu       string `json:"details_ru,omitempty"` // Optional details in Russian.
	ApiEndpoint     string `json:"api_endpoint"`
	StatusCode      int    `json:"status_code"`
	RequestPayload  string `json:"request_payload"`
	EventType       string `json:"event_type"`                 // Event type, usually based on the function name.
	ResponseData    string `json:"response_data,omitempty"`    // Optional response data.
	MerchantApiKey  string `json:"merchant_api_key,omitempty"` // Merchant API key, required if sending to merchants.
}

// NewLogger initializes and returns a new Logger instance.
// Parameters:
// - rabbitMQUrl: RabbitMQ connection string.
// - queueName: Name of the RabbitMQ queue where logs will be sent.
// - functionName: Name of the function generating logs.
// - apiEndpoint: API endpoint associated with the logs.
func NewLogger(rabbitMQUrl, queueName string, functionName string, apiEndpoint string) (Logger, error) {
	rmq, err := rabbitmq.NewRabbitMQ(rabbitMQUrl)
	if err != nil {
		return nil, fmt.Errorf("failed to initialize RabbitMQ: %s", err)
	}

	// Declare the queue in RabbitMQ.
	err = rmq.DeclareQueue(queueName, true, true, false, false, amqp.Table{})
	if err != nil {
		return nil, fmt.Errorf("failed to declare queue: %s", err)
	}

	return &logger{
		rabbitmq:     rmq,
		queue:        queueName,
		functionName: functionName,
		apiEndpoint:  apiEndpoint,
	}, nil
}

// Info logs an informational message.
func (l *logger) Info(log LogRequest) error {
	fullLog := l.populateLogRequest(log, "info")
	return l.publish(fullLog)
}

// Warn logs a warning message.
func (l *logger) Warn(log LogRequest) error {
	fullLog := l.populateLogRequest(log, "warning")
	return l.publish(fullLog)
}

// Error logs an error message.
func (l *logger) Error(log LogRequest) error {
	fullLog := l.populateLogRequest(log, "error")
	return l.publish(fullLog)
}

// Critical logs a critical error message.
func (l *logger) Critical(log LogRequest) error {
	fullLog := l.populateLogRequest(log, "critical")
	return l.publish(fullLog)
}

// CloseRabbitMQConnection closes the RabbitMQ connection.
func (l *logger) CloseRabbitMQConnection() error {
	return l.rabbitmq.Close()
}

// publish validates the log request, serializes it, and sends it to RabbitMQ.
func (l *logger) publish(log logRequest) error {
	// Validate the log request.
	if err := validateLogRequest(log); err != nil {
		return err
	}

	// Serialize the log request into JSON format.
	message, err := json.Marshal(log)
	if err != nil {
		return fmt.Errorf("failed to serialize log request: %s", err)
	}

	// Publish the log message to RabbitMQ.
	err = l.rabbitmq.PublishMessage(l.queue, "", message)
	if err != nil {
		return fmt.Errorf("failed to publish log message: %s", err)
	}

	fmt.Printf("[%s] Log sent: %s\n", log.ErrorLevel, string(message))
	return nil
}

// validateLogRequest ensures that required fields in the log request are present.
func validateLogRequest(log logRequest) error {
	if log.ErrorCode == 0 {
		return errors.New("error_code is required")
	}
	if log.ClientMessageUz == "" && log.ClientMessageRu == "" {
		return errors.New("at least one client message (Uz or Ru) is required")
	}

	return nil
}

// populateLogRequest populates a `logRequest` with additional metadata like timestamp, error level, and function name.
func (l *logger) populateLogRequest(log LogRequest, errorLevel string) logRequest {
	logRequest := logRequest{
		Timestamp:       time.Now(),
		ErrorLevel:      errorLevel,
		ErrorCode:       log.ErrorCode,
		ClientMessageUz: log.ClientMessageUz,
		ClientMessageRu: log.ClientMessageRu,
		ErrorMessage:    log.ErrorMessage,
		DetailsUz:       log.DetailsUz,
		DetailsRu:       log.DetailsRu,
		ApiEndpoint:     log.ApiEndpoint,
		FunctionName:    l.functionName,
		StatusCode:      log.StatusCode,
		RequestPayload:  log.RequestPayload,
		EventType:       log.EventType,
		ResponseData:    log.ResponseData,
		MerchantApiKey:  log.MerchantApiKey,
	}
	// Fallbacks for missing API endpoint or status code.
	if log.ApiEndpoint == "" {
		logRequest.ApiEndpoint = l.apiEndpoint
	}
	if log.StatusCode == 0 {
		logRequest.StatusCode = 200
	}

	return logRequest
}
