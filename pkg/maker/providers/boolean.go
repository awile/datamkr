package providers

import (
	"fmt"
	"math/rand"

	"github.com/awile/datamkr/pkg/dataset"
)

type booleanFieldMaker struct {
	definition dataset.DatasetDefinitionField
}

func NewBooleanWithDefinition(definition dataset.DatasetDefinitionField) FieldProviderInterface {
	return &booleanFieldMaker{definition: definition}
}

func (bfm *booleanFieldMaker) MakeField() ProviderField {
	boolean := rand.Intn(100) >= 50
	return &booleanField{value: boolean}
}

type booleanField struct {
	value bool
}

func (b *booleanField) String() string {
	return fmt.Sprintf("%t", b.value)
}

func (b *booleanField) Value() interface{} {
	return b.value
}
