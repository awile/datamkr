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

func (ufm *uuidFieldMaker) MakeField() interface{} {
	return uuid.NewString()
}
