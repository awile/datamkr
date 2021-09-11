package config

type DatamkrConfig struct {
	DatasetsDir    string                  `yaml:"datasetsDir"`
	StorageAliases map[string]StorageAlias `yaml:"storage,omitempty"`
}

func (d *DatamkrConfig) GetStorageAlias(alias string) string {
	if len(d.StorageAliases) > 0 {
		storageAlias, found := d.StorageAliases[alias]
		if found {
			return storageAlias.Value()
		}
	}
	return ""
}

type StorageAlias struct {
	ConnectionString string `yaml:"connection,omitempty"`
}

func (sa *StorageAlias) Value() string {
	if sa.ConnectionString != "" {
		return sa.ConnectionString
	}
	return ""
}
