package providers

import (
	"math/rand"
	"strconv"

	"github.com/awile/datamkr/pkg/dataset"
)

type booleanFieldMaker struct {
	definition dataset.DatasetDefinitionField
}

func NewBooleanWithDefinition(definition dataset.DatasetDefinitionField) FieldProviderInterface {
	return &booleanFieldMaker{definition: definition}
}

func (bfm *booleanFieldMaker) MakeField() interface{} {
	return strconv.FormatBool(rand.Intn(100) >= 50)
}
