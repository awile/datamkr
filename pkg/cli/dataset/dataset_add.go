package dataset

import (
	"errors"

	utils "github.com/awile/datamkr/pkg/cli/util"
	"github.com/awile/datamkr/pkg/client"
	"github.com/awile/datamkr/pkg/config"
	"github.com/spf13/cobra"
)

type DatasetAddOptions struct {
	DatasetName string

	factory       config.ConfigFactory
	datamkrClient client.Interface
}

func NewDatasetAddOptions(factory *config.DatamkrConfigFactory) *DatasetAddOptions {
	return &DatasetAddOptions{factory: factory}
}

func NewDatasetAddCmd(configFactory *config.DatamkrConfigFactory) *cobra.Command {
	datasetAddOptions := NewDatasetAddOptions(configFactory)

	cmd := &cobra.Command{
		Use:     "add",
		Short:   "Add a new dataset definition",
		Long:    "Add a new dataset definition.",
		Example: "datamkr dataset add <dataset_name>",
		Run: func(cmd *cobra.Command, args []string) {
			utils.CheckErr(datasetAddOptions.Complete(cmd, args))
			utils.CheckErr(datasetAddOptions.Validate())
			utils.CheckErr(datasetAddOptions.Run())
		},
	}

	return cmd
}

func (opt *DatasetAddOptions) Complete(cmd *cobra.Command, args []string) error {
	if len(args) == 0 {
		return errors.New("Must give dataset a name:\n\n    datamkr dataset add <dataset_name>\n\n")
	} else {
		opt.DatasetName = args[0]
	}

	currentConfig, err := opt.factory.GetConfig()
	if err != nil {
		return err
	}
	opt.datamkrClient = client.NewWithConfig(currentConfig)

	return nil
}

func (opt *DatasetAddOptions) Validate() error {
	if opt.DatasetName == "" {
		return errors.New("Must give dataset a name:\n\n    datamkr dataset add <dataset_name>\n\n")
	}
	return nil
}

func (opt *DatasetAddOptions) Run() error {
	datasetClient := opt.datamkrClient.Datasets()
	return datasetClient.Add(opt.DatasetName)
}
