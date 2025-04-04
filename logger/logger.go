package logger

import (
	"encoding/json"
	"errors"
	"fmt"
	"time"

	rabbitmq "github.com/kupalovmuhammadjon/rabbitmq-go"
	amqp "github.com/rabbitmq/amqp091-go"
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

	OrderNotification(order Order) error

	SendOrderToBitrix(order BitrixOrder) error
}

// NewLogger initializes and returns a new Logger instance.
// Parameters:
// - rabbitMQ: RabbitMQ interface.
// - queueName: Name of the RabbitMQ queue where logs will be sent.
// - functionName: Name of the function generating logs.
// - apiEndpoint: API endpoint associated with the logs.
func NewLogger(rabbitMQ rabbitmq.RabbitMQ, queueName, funtionName, apiEndpoint string, orderQueue, bitrixOrderQueue *string) (Logger, error) {

	err := rabbitMQ.DeclareQueue(queueName, true, true, false, false, amqp.Table{})
	if err != nil {
		return nil, fmt.Errorf("failed to declare queue: %s", err)
	}

	var oQueue string
	var bitrixOQueue string
	if orderQueue != nil {
		oQueue = *orderQueue
	}

	if bitrixOrderQueue != nil {
		bitrixOQueue = *bitrixOrderQueue
	}

	return &logger{
		rabbitmq:         rabbitMQ,
		queue:            queueName,
		orderQueue:       oQueue,
		bitrixOrderQueue: bitrixOQueue,
		functionName:     funtionName,
		apiEndpoint:      apiEndpoint,
	}, nil
}

// Info logs an informational message.
func (l *logger) Info(log LogRequest) error {
	fullLog, err := l.populateLogRequest(log, "info")
	if err != nil {
		return err
	}

	if err := validateLogRequest(fullLog); err != nil {
		return err
	}

	return l.rabbitmq.PublishMessage(l.queue, "", fullLog)
}

// Warn logs a warning message.
func (l *logger) Warn(log LogRequest) error {
	fullLog, err := l.populateLogRequest(log, "warning")
	if err != nil {
		return err
	}

	if err := validateLogRequest(fullLog); err != nil {
		return err
	}

	return l.rabbitmq.PublishMessage(l.queue, "", fullLog)
}

// Error logs an error message.
func (l *logger) Error(log LogRequest) error {
	fullLog, err := l.populateLogRequest(log, "error")
	if err != nil {
		return err
	}

	if err := validateLogRequest(fullLog); err != nil {
		return err
	}

	return l.rabbitmq.PublishMessage(l.queue, "", fullLog)
}

// Critical logs a critical error message.
func (l *logger) Critical(log LogRequest) error {
	fullLog, err := l.populateLogRequest(log, "critical")
	if err != nil {
		return err
	}

	if err := validateLogRequest(fullLog); err != nil {
		return err
	}

	return l.rabbitmq.PublishMessage(l.queue, "", fullLog)
}

func (l *logger) OrderNotification(order Order) error {
	return l.rabbitmq.PublishMessage(l.orderQueue, "", order)
}

func (l *logger) SendOrderToBitrix(order BitrixOrder) error {
	return l.rabbitmq.PublishMessage(l.bitrixOrderQueue, "", order)
}

// validateLogRequest ensures that required fields in the log request are present.
func validateLogRequest(log logRequest) error {
	if log.Errorcode == 0 {
		return errors.New("error_code is required")
	}

	if log.ClientMessageUz == "" && log.ClientMessageRu == "" {
		return errors.New("at least one client message (Uz or Ru) is required")
	}

	if log.ErrorLevel == "" || ((log.ErrorLevel == "error" || log.ErrorLevel == "critical") && log.RequestPayload == "") {
		return errors.New("request payload is required for this error level")
	}

	return nil
}

// populateLogRequest populates a `logRequest` with additional metadata like timestamp, error level, and function name.
func (l *logger) populateLogRequest(log LogRequest, errorLevel string) (logRequest, error) {

	var (
		body []byte
		err  error
	)

	switch msg := log.RequestPayload.(type) {
	case []byte:
		body = msg
	case string:
		body = []byte(msg)
	default:
		body, err = json.Marshal(msg)
		if err != nil {
			return logRequest{}, err
		}
	}

	logRequest := logRequest{
		Timestamp:       time.Now(),
		ErrorLevel:      errorLevel,
		Errorcode:       int(log.Errorcode),
		ClientMessageUz: log.ClientMessageUz,
		ClientMessageRu: log.ClientMessageRu,
		ErrorMessage:    log.ErrorMessage,
		DetailsUz:       log.DetailsUz,
		DetailsRu:       log.DetailsRu,
		ApiEndpoint:     log.ApiEndpoint,
		Method:          log.Method,
		FunctionName:    l.functionName,
		StatusCode:      log.StatusCode,
		RequestPayload:  string(body),
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

	return logRequest, nil
}
