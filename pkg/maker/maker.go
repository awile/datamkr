package maker

import (
	"github.com/awile/datamkr/pkg/config"
	"github.com/awile/datamkr/pkg/dataset"
)

type MakerClientInterface interface {
	MakeRow(definition dataset.DatasetDefinition) (map[string]interface{}, error)
}

type MakerClient struct {
	config *config.DatamkrConfig
}

func NewWithConfig(config *config.DatamkrConfig) *MakerClient {
	var mc MakerClient

	mc.config = config

	return &mc
}

func (mc *MakerClient) MakeRow(definition dataset.DatasetDefinition) (map[string]interface{}, error) {
	var err error
	row := make(map[string]interface{})
	for key, fieldDefinition := range definition.Fields {
		fieldMaker, err := NewFieldMaker(fieldDefinition)
		if err != nil {
			break
		}
		row[key] = fieldMaker.MakeField()
	}
	return row, err
}
