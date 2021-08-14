package init

import (
	"testing"

	"github.com/awile/datamkr/pkg/config"
	"github.com/stretchr/testify/mock"
)

func TestInitOptionsRun(t *testing.T) {
	mockFactory := new(config.MockDatamkrConfigFactory)

	mockFactory.On("HasConfigInDirectory", mock.Anything).Return(false, nil)

	mockFactory.On("CreateNewConfigFile", mock.Anything)

	mockFactory.On("InitDatamkrConfigFile", mock.Anything).Return(nil)

	initOptions := &InitOptions{factory: mockFactory}

	err := initOptions.Run()
	if err != nil {
		t.Fatal("Failed to run initOptions.Run")
	}

	mockFactory.AssertNumberOfCalls(t, "HasConfigInDirectory", 1)
	mockFactory.AssertNumberOfCalls(t, "InitDatamkrConfigFile", 1)
	mockFactory.AssertExpectations(t)
}
