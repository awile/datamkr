package config

import (
	"io"

	"github.com/stretchr/testify/mock"
)

type MockIOWriter struct {
	mock.Mock
}

func (m *MockIOWriter) Write(p []byte) (n int, err error) {
	args := m.Called(p)
	return args.Int(0), args.Error(1)
}

type MockDatamkrConfigFactory struct {
	mock.Mock
}

func (m *MockDatamkrConfigFactory) GetConfig() (*DatamkrConfig, error) {
	args := m.Called()
	return args.Get(0).(*DatamkrConfig), args.Error(1)
}

func (m *MockDatamkrConfigFactory) ConfigToByteString() ([]byte, error) {
	args := m.Called()
	return []byte(args.String(0)), args.Error(1)
}

func (m *MockDatamkrConfigFactory) HasConfigInDirectory() (bool, error) {
	args := m.Called()
	return args.Bool(0), args.Error(1)
}

func (m *MockDatamkrConfigFactory) InitDatamkrConfigFile(configFile io.Writer) error {
	args := m.Called(configFile)
	return args.Error(0)
}

func (m *MockDatamkrConfigFactory) CreateNewConfigFile() io.Writer {
	mockIOWriter := new(MockIOWriter)
	m.Called()
	return mockIOWriter
}
