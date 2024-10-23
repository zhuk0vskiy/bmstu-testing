package utils

import "ppo/pkg/logger"

type mockLogger struct{}

func NewMockLogger() logger.ILogger {
	return &mockLogger{}
}

func (m *mockLogger) Infof(message string, args ...interface{}) {}

func (m *mockLogger) Warnf(message string, args ...interface{}) {}

func (m *mockLogger) Errorf(message string, args ...interface{}) {}

func (m *mockLogger) Fatalf(message string, args ...interface{}) {}
