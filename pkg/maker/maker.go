package maker

import (
	"math/rand"
	"time"

	"github.com/awile/datamkr/pkg/config"
	"github.com/awile/datamkr/pkg/dataset"
	"github.com/awile/datamkr/pkg/maker/providers"
)

type MakerClientInterface interface {
	MakeRow(definition dataset.DatasetDefinition) (map[string]interface{}, error)
}

type MakerClient struct {
	providers map[string]providers.FieldProviderInterface
	config    *config.DatamkrConfig
}

func NewWithConfig(config *config.DatamkrConfig) MakerClientInterface {
	rand.Seed(time.Now().UnixNano())
	var mc MakerClient

	mc.config = config
	mc.providers = make(map[string]providers.FieldProviderInterface)

	return &mc
}

func (mc *MakerClient) MakeRow(definition dataset.DatasetDefinition) (map[string]interface{}, error) {
	var err error
	row := make(map[string]interface{}, len(definition.Fields))
	for key, fieldDefinition := range definition.Fields {
		fieldProvider, err := mc.getProvider(fieldDefinition)
		if err != nil {
			break
		}
		row[key] = fieldProvider.MakeField()
	}
	return row, err
}

func (mc *MakerClient) getProvider(fieldDefinition dataset.DatasetDefinitionField) (providers.FieldProviderInterface, error) {
	provider := mc.providers[fieldDefinition.Type]
	if provider != nil {
		return provider, nil
	}
	provider, err := providers.NewFieldProvider(fieldDefinition)
	if err != nil {
		return nil, err
	}
	mc.providers[fieldDefinition.Type] = provider
	return provider, nil
}
