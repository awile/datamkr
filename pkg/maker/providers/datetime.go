package providers

import (
	"time"

	"github.com/awile/datamkr/pkg/dataset"
)

type datetimeFieldMaker struct {
	definition dataset.DatasetDefinitionField
}

func NewDatetimeWithDefinition(definition dataset.DatasetDefinitionField) FieldProviderInterface {
	return &datetimeFieldMaker{definition: definition}
}

func (t *datetimeFieldMaker) MakeField() ProviderField {
	now := time.Now()
	return &datetimeField{value: now}
}

type datetimeField struct {
	value time.Time
}

func (str *datetimeField) String() string {
	return str.value.Format(time.RFC3339)
}

func (str *datetimeField) Value() interface{} {
	return str.value
}
