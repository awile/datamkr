package config

type DatamkrConfig struct {
	DatasetsDir    string                  `yaml:"datasetsDir"`
	StorageAliases map[string]StorageAlias `yaml:"storage,omitempty"`
	Version        int32
}

func (d *DatamkrConfig) GetStorageAlias(alias string) (StorageAlias, bool) {
	var storageAlias StorageAlias
	if len(d.StorageAliases) > 0 {
		storageAlias, found := d.StorageAliases[alias]
		return storageAlias, found
	}
	return storageAlias, false
}

type StorageAlias struct {
	ConnectionString string `yaml:"connection,omitempty"`
	Type             string `yaml:"type"`
}
