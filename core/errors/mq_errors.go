package errors

import "fmt"

type MqNotConfiguredError struct{}

func (e *MqNotConfiguredError) Error() string {
	return "Message queue configuration not found!"
}

type MqTypeNotSupportedError struct {
	Type string
}

func (e *MqTypeNotSupportedError) Error() string {
	return fmt.Sprintf("Message queue type (%s) not supported.", e.Type)
}

type ConfigurationNotFoundError struct {
	Type string
}

func (e *ConfigurationNotFoundError) Error() string {
	return fmt.Sprintf("Message Queue of type (%s) not configured.", e.Type)
}

type MissingConfigurationKeyError struct {
	Type string
	Key  string
}

func (e *MissingConfigurationKeyError) Error() string {
	return fmt.Sprintf("Message Queue of type (%s) missing required configuration key %s.", e.Type, e.Key)
}

type MqError struct {
	Type    string
	Message string
	Err     error
}

func (e *MqError) Error() string {
	return fmt.Sprintf("%s | type -> %s | %s", e.Message, e.Type, e.Err.Error())
}
