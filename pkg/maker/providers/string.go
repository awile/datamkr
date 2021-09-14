package providers

import (
	"bytes"
	"math/rand"
	"time"

	"github.com/awile/datamkr/pkg/dataset"
)

type stringFieldMaker struct {
	definition dataset.DatasetDefinitionField
}

func NewStringWithDefinition(definition dataset.DatasetDefinitionField) FieldProviderInterface {
	return &stringFieldMaker{definition: definition}
}

func (sfm *stringFieldMaker) MakeField() ProviderField {
	// 65: A, 122: z
	rand.Seed(time.Now().UnixNano())
	length := rand.Intn(30) + 5

	upperLowerDiff := 97 - 65
	var randomStr bytes.Buffer
	for i := 0; i < length; i++ {
		randInt := rand.Intn(26) + 65
		if rand.Intn(2) == 1 {
			randInt += upperLowerDiff
		}
		char := string(byte(randInt))
		randomStr.WriteString(char)
	}
	return &stringField{value: randomStr.String()}
}

type stringField struct {
	value string
}

func (str *stringField) String() string {
	return str.value
}

func (str *stringField) Value() interface{} {
	return str.value
}
