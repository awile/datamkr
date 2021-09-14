package providers

import (
	"math/rand"
	"strconv"

	"github.com/awile/datamkr/pkg/dataset"
)

type intFieldMaker struct {
	definition dataset.DatasetDefinitionField
}

func NewIntWithDefinition(definition dataset.DatasetDefinitionField) FieldProviderInterface {
	return &intFieldMaker{definition: definition}
}

func (ifm *intFieldMaker) MakeField() ProviderField {
	randInt := rand.Intn(1000)
	return &intField{value: randInt}
}

type intField struct {
	value int
}

func (i *intField) String() string {
	return strconv.Itoa(i.value)
}

func (i *intField) Value() interface{} {
	return i.value
}
