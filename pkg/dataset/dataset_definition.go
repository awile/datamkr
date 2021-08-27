package dataset

type DatasetDefinitionField struct {
	Type string `yaml:"type"`
}

type DatasetDefinition struct {
	Fields map[string]DatasetDefinitionField `yaml:"fields"`
}
