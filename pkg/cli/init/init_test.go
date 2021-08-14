package init

import (
	"errors"
	"testing"

	"github.com/awile/datamkr/pkg/config"
	"github.com/stretchr/testify/mock"
)

func TestInitOptionsRun(t *testing.T) {
	mockFactory := new(config.MockDatamkrConfigFactory)

	mockFactory.On("CreateNewConfigFile", mock.Anything)

	mockFactory.On("InitDatamkrConfigFile", mock.Anything).Return(nil)

	initOptions := &InitOptions{
		HasConfig: false,
		factory:   mockFactory,
	}

	err := initOptions.Run()
	if err != nil {
		t.Fatal("Failed to run initOptions.Run")
	}

	mockFactory.AssertNumberOfCalls(t, "InitDatamkrConfigFile", 1)
	mockFactory.AssertNumberOfCalls(t, "CreateNewConfigFile", 1)
	mockFactory.AssertExpectations(t)
}

func TestInitOptionsRunNoConfig(t *testing.T) {
	mockFactory := new(config.MockDatamkrConfigFactory)

	initOptions := &InitOptions{
		HasConfig: true,
		factory:   mockFactory,
	}

	err := initOptions.Run()
	if err != nil {
		t.Fatal("Failed to run initOptions.Run")
	}

	mockFactory.AssertNotCalled(t, "InitDatamkrConfigFile")
	mockFactory.AssertNotCalled(t, "CreateNewConfigFile")
}

func TestRunInitConfigError(t *testing.T) {
	mockFactory := new(config.MockDatamkrConfigFactory)

	mockFactory.On("CreateNewConfigFile", mock.Anything)

	mockFactory.On("InitDatamkrConfigFile", mock.Anything).Return(errors.New("failed to init config file"))

	initOptions := &InitOptions{
		HasConfig: false,
		factory:   mockFactory,
	}

	err := initOptions.Run()
	if err != nil {
		mockFactory.AssertNumberOfCalls(t, "InitDatamkrConfigFile", 1)
		mockFactory.AssertNumberOfCalls(t, "CreateNewConfigFile", 1)
		mockFactory.AssertExpectations(t)
	} else {
		t.Fatal("initOptions.Run should have returned an error")
	}
}

func TestValidate(t *testing.T) {
	mockFactory := new(config.MockDatamkrConfigFactory)

	initOptions := &InitOptions{
		HasConfig: false,
		factory:   mockFactory,
	}

	err := initOptions.Validate()
	if err != nil {
		t.Fatal("initOptions.Validat should not have returned an error")
	}
}
