package add

import (
	"errors"
	"fmt"
	"strings"

	utils "github.com/awile/datamkr/pkg/cli/util"
	"github.com/awile/datamkr/pkg/client"
	"github.com/awile/datamkr/pkg/config"
	"github.com/awile/datamkr/pkg/dataset"
	"github.com/awile/datamkr/pkg/storage"
	"github.com/spf13/cobra"
)

type DatasetAddOptions struct {
	DatasetName string
	Definition  dataset.DatasetDefinition
	Fieldslist  []string
	From        string
	Table       string

	storageType   string
	factory       config.ConfigFactory
	datamkrClient client.Interface
}

func NewAddOptions(factory *config.DatamkrConfigFactory) *DatasetAddOptions {
	return &DatasetAddOptions{factory: factory}
}

func NewAddCmd(configFactory *config.DatamkrConfigFactory) *cobra.Command {
	datasetAddOptions := NewAddOptions(configFactory)

	cmd := &cobra.Command{
		Use:     "add",
		Short:   "Add a new dataset definition",
		Long:    "Add a new dataset definition.",
		Example: "datamkr add <dataset_name>",
		Run: func(cmd *cobra.Command, args []string) {
			utils.CheckErr(datasetAddOptions.Complete(cmd, args))
			utils.CheckErr(datasetAddOptions.Validate())
			utils.CheckErr(datasetAddOptions.Run())
		},
	}

	cmd.Flags().StringArrayVar(
		&datasetAddOptions.Fieldslist,
		"field",
		datasetAddOptions.Fieldslist,
		"Define definition fields (e.g. --field name=id,key2=uuid --field name=email,type=email)",
	)
	cmd.Flags().StringVar(&datasetAddOptions.From, "from", "", "database to import dataset definition from")
	cmd.Flags().StringVarP(&datasetAddOptions.Table, "table", "t", "", "DB table to use (only valid for: --from <db_connection_string>)")
	return cmd
}

func (opt *DatasetAddOptions) Complete(cmd *cobra.Command, args []string) error {
	currentConfig, err := opt.factory.GetConfig()
	if err != nil {
		return err
	}
	opt.datamkrClient = client.NewWithConfig(currentConfig)

	if len(args) == 0 {
		return errors.New("Must give dataset a name:\n\n    datamkr add <dataset_name>\n\n")
	} else {
		opt.DatasetName = args[0]
	}

	storageAlias, aliasExists := currentConfig.GetStorageAlias(opt.From)
	if aliasExists {
		if storageAlias.Type == "postgres" {
			opt.From = storageAlias.ConnectionString
			opt.Table = storageAlias.Table
		}
	}

	var datasetDefinition dataset.DatasetDefinition
	if len(opt.Fieldslist) > 0 {
		datasetDefinition.Fields = parseDatasetDefinitionFields(opt.Fieldslist)
	}
	opt.Definition = datasetDefinition

	if strings.Contains(opt.From, "postgresql://") {
		opt.storageType = "postgres"
	}
	return nil
}

func (opt *DatasetAddOptions) Validate() error {
	if opt.DatasetName == "" {
		return errors.New("Must give dataset a name:\n\n    datamkr add <dataset_name>\n\n")
	}
	if opt.storageType == "postgres" && opt.Table == "" {
		return fmt.Errorf("Must provide which postgres table to use: --table <table_name>\n")
	}
	return nil
}

func (opt *DatasetAddOptions) Run() error {
	datasetClient := opt.datamkrClient.Datasets()

	if opt.storageType != "" {
		storageClient := opt.datamkrClient.Storage()
		readerOptions := storage.CreateReaderOptions()
		if opt.storageType == "postgres" {
			readerOptions.Id = opt.From
			readerOptions.SecondaryId = opt.Table
		}

		storageReader := storageClient.GetStorageServiceReader(opt.storageType, readerOptions)
		if storageReader == nil {
			return fmt.Errorf("%s is not a valid target\n", opt.storageType)
		}

		err := storageReader.Init()
		if err != nil {
			return err
		}
		defer storageReader.Close()

		datasetDefinition, err := storageReader.GetDatasetDefinition()
		if err != nil {
			return err
		}
		opt.Definition = datasetDefinition
	}

	err := datasetClient.Add(opt.DatasetName, opt.Definition)
	if err != nil {
		return err
	}
	fmt.Printf("Dataset %s created\n", opt.DatasetName)
	return nil
}

func parseDatasetDefinitionFields(fields []string) map[string]dataset.DatasetDefinitionField {
	definitionFields := make(map[string]dataset.DatasetDefinitionField)
	for _, fieldStr := range fields {
		var name string
		var field dataset.DatasetDefinitionField
		for _, keyValuePair := range strings.Split(fieldStr, ",") {
			pair := strings.Split(keyValuePair, "=")
			key := strings.TrimSpace(pair[0])
			value := strings.TrimSpace(pair[1])
			if key == "name" {
				name = value
			} else if key == "type" {
				field.Type = value
			}
		}

		definitionFields[name] = field
	}
	return definitionFields
}
