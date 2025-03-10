# Logger Package

The `logger` package is designed to manage and send log messages to RabbitMQ with different severity levels. It provides an easy way to log messages for various events such as information, warnings, errors, and critical failures. It includes the ability to log with additional metadata like request payload, response data, status codes, and more. 

This package leverages RabbitMQ for message queuing, making it ideal for distributed systems where logs need to be collected and processed asynchronously.

## Features

- Logs messages with different severity levels: `Info`, `Warn`, `Error`, and `Critical`.
- Sends logs to RabbitMQ for further processing or storage.
- Includes useful metadata in logs such as `error_code`, `client_message`, `request_payload`, `response_data`, and more.
- Handles different error codes for various error types (validation, authentication, resource, system, integration, etc.).
- Supports customizable log data, including API endpoint, method, status code, and merchant API key.

## Installation

To install the `logger` package, use the following Go command:

```bash
go get github.com/yourusername/logger
