package maker

import (
	"bytes"
	"fmt"
	"math/rand"
	"time"

	"github.com/awile/datamkr/pkg/dataset"
	"github.com/google/uuid"
)

type FieldMakerInterface interface {
	MakeField() interface{}
}

func NewFieldMaker(fieldDefinition dataset.DatasetDefinitionField) (FieldMakerInterface, error) {
	switch fieldDefinition.Type {
	case "uuid":
		return &UuidFieldMaker{definition: fieldDefinition}, nil
	case "string":
		return &StringFieldMaker{definition: fieldDefinition}, nil
	default:
		return nil, fmt.Errorf("Field type %s does not exist", fieldDefinition.Type)
	}
}

type UuidFieldMaker struct {
	definition dataset.DatasetDefinitionField
}

func (ufm *UuidFieldMaker) MakeField() interface{} {
	return uuid.NewString()
}

type StringFieldMaker struct {
	definition dataset.DatasetDefinitionField
}

func (sfm *StringFieldMaker) MakeField() interface{} {
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
	return randomStr.String()
}
