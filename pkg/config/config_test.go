package config

// func TestNewDatamkrConfigFactory(t *testing.T) {
// 	configFactory, err := NewDatamkrConfigFactory()
// 	if err != nil {
// 		t.Fatal("Failed to create DatamkrConfigFactory")
// 	}
// 	hasConfig, err := configFactory.HasConfigInDirectory()
// 	if err != nil {
// 		t.Fatal("Failed to check if config file exists")
// 	}
// 	assert.Equal(t, hasConfig, false)
// }

// func TestInitDatamkrConfigFile(t *testing.T) {
// 	mockIOWriter := new(MockIOWriter)

// 	mockIOWriter.On("Write", mock.Anything).Return(1, nil)

// 	configFactory, err := NewDatamkrConfigFactory()
// 	if err != nil {
// 		t.Fatal("Failed to create new config")
// 	}

// 	err = configFactory.InitDatamkrConfigFile(mockIOWriter)
// 	if err != nil {
// 		t.Fatal("Failed to create new config")
// 	}

// 	mockIOWriter.AssertNumberOfCalls(t, "Write", 2)
// 	mockIOWriter.AssertExpectations(t)
// }
