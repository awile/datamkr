package providers

import (
	"github.com/awile/datamkr/pkg/dataset"
	"github.com/google/uuid"
)

type uuidFieldMaker struct {
	definition dataset.DatasetDefinitionField
}

func NewUuidWithDefinition(definition dataset.DatasetDefinitionField) FieldProviderInterface {
	return &uuidFieldMaker{definition: definition}
}

func (ufm *uuidFieldMaker) MakeField() ProviderField {
	return &uuidField{value: uuid.NewString()}
}

type uuidField struct {
	value string
}

func (uuid *uuidField) String() string {
	return uuid.value
}

func (uuid *uuidField) Value() interface{} {
	return uuid.value
}
