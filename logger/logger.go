package logger

import (
	"encoding/json"
	"errors"
	"fmt"
	rabbitmq "mybazar_logger/rabbitMQ"
	"time"

	"github.com/streadway/amqp"
)

type Logger interface {
	Info(log LogRequest) error
	Warn(log LogRequest) error
	Error(log LogRequest) error
	Critical(log LogRequest) error
	CloseRabbitMQConnection() error
}

type logger struct {
	rabbitmq     rabbitmq.RabbitMQ
	queue        string
	functionName string
	apiEndpoint  string
}

type logRequest struct {
	Timestamp       time.Time `json:"timestamp"`
	ErrorLevel      string    `json:"error_level"`
	ErrorCode       int       `json:"error_code"`
	ClientMessageUz string    `json:"client_message_uz"`
	ClientMessageRu string    `json:"client_message_ru"`
	ErrorMessage    string    `json:"error_message"`
	DetailsUz       string    `json:"details_uz,omitempty"` // optional
	DetailsRu       string    `json:"details_ru,omitempty"` // optional
	ApiEndpoint     string    `json:"api_endpoint"`
	FunctionName    string    `json:"function_name"`
	StatusCode      int       `json:"status_code"`
	RequestPayload  string    `json:"request_payload"`
	EventType       string    `json:"event_type"`                 // based on function name
	ResponseData    string    `json:"response_data,omitempty"`    // optional if possible
	MerchantApiKey  string    `json:"merchant_api_key,omitempty"` // REQUIRED if send to merchant is on
}

type LogRequest struct {
	ErrorCode       int    `json:"error_code"`
	ClientMessageUz string `json:"client_message_uz"`
	ClientMessageRu string `json:"client_message_ru"`
	ErrorMessage    string `json:"error_message"`
	DetailsUz       string `json:"details_uz,omitempty"` // optional
	DetailsRu       string `json:"details_ru,omitempty"` // optional
	ApiEndpoint     string `json:"api_endpoint"`
	StatusCode      int    `json:"status_code"`
	RequestPayload  string `json:"request_payload"`
	EventType       string `json:"event_type"`                 // based on function name
	ResponseData    string `json:"response_data,omitempty"`    // optional if possible
	MerchantApiKey  string `json:"merchant_api_key,omitempty"` // REQUIRED if send to merchant is on
}

func NewLogger(rabbitMQUrl, queueName string, funtionName string, apiEndpoint string) (Logger, error) {
	rmq, err := rabbitmq.NewRabbitMQ(rabbitMQUrl)
	if err != nil {
		return nil, fmt.Errorf("failed to initialize RabbitMQ: %s", err)
	}

	err = rmq.DeclareQueue(queueName, true, true, false, false, amqp.Table{})
	if err != nil {
		return nil, fmt.Errorf("failed to declare queue: %s", err)
	}

	return &logger{
		rabbitmq:     rmq,
		queue:        queueName,
		functionName: funtionName,
		apiEndpoint:  apiEndpoint,
	}, nil
}

func (l *logger) Info(log LogRequest) error {
	fullLog := l.populateLogRequest(log, "info")
	return l.publish(fullLog)
}

func (l *logger) Warn(log LogRequest) error {
	fullLog := l.populateLogRequest(log, "warning")
	return l.publish(fullLog)
}

func (l *logger) Error(log LogRequest) error {
	fullLog := l.populateLogRequest(log, "error")
	return l.publish(fullLog)
}

func (l *logger) Critical(log LogRequest) error {
	fullLog := l.populateLogRequest(log, "critical")
	return l.publish(fullLog)
}

func (l *logger) CloseRabbitMQConnection() error {
	return l.rabbitmq.Close()
}

func (l *logger) publish(log logRequest) error {
	if err := validateLogRequest(log); err != nil {
		return err
	}

	message, err := json.Marshal(log)
	if err != nil {
		return fmt.Errorf("failed to serialize log request: %s", err)
	}

	err = l.rabbitmq.PublishMessage(l.queue, "", message)
	if err != nil {
		return fmt.Errorf("failed to publish log message: %s", err)
	}

	fmt.Printf("[%s] Log sent: %s\n", log.ErrorLevel, string(message))
	return nil
}

func validateLogRequest(log logRequest) error {
	if log.ErrorCode == 0 {
		return errors.New("error_code is required")
	}
	if log.ClientMessageUz == "" && log.ClientMessageRu == "" {
		return errors.New("at least one client message (Uz or Ru) is required")
	}

	return nil
}

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
	if log.ApiEndpoint == "" {
		logRequest.ApiEndpoint = l.apiEndpoint
	}
	if log.StatusCode == 0 {
		logRequest.StatusCode = 200
	}

	return logRequest
}
