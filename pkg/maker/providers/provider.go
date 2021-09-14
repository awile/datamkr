package providers

import (
	"fmt"

	"github.com/awile/datamkr/pkg/dataset"
)

type FieldProviderInterface interface {
	MakeField() ProviderField
}

func NewFieldProvider(fieldDefinition dataset.DatasetDefinitionField) (FieldProviderInterface, error) {
	switch fieldDefinition.Type {
	case "uuid":
		return NewUuidWithDefinition(fieldDefinition), nil
	case "string":
		return NewStringWithDefinition(fieldDefinition), nil
	case "name":
		return NewNameWithDefinition(fieldDefinition), nil
	case "email":
		return NewEmailWithDefinition(fieldDefinition), nil
	case "boolean":
		return NewBooleanWithDefinition(fieldDefinition), nil
	default:
		return nil, fmt.Errorf("Field type %s does not exist", fieldDefinition.Type)
	}
}

type ProviderField interface {
	String() string
	Value() interface{}
}
