package dataset

import (
	"fmt"

	utils "github.com/awile/datamkr/pkg/cli/util"
	"github.com/awile/datamkr/pkg/client"
	"github.com/awile/datamkr/pkg/config"
	"github.com/spf13/cobra"
)

type DatasetListOptions struct {
	factory       config.ConfigFactory
	datamkrClient client.Interface
}

func NewDatasetListOptions(factory *config.DatamkrConfigFactory) *DatasetListOptions {
	return &DatasetListOptions{factory: factory}
}

func NewDatasetListCmd(configFactory *config.DatamkrConfigFactory) *cobra.Command {
	datasetListOptions := NewDatasetListOptions(configFactory)

	cmd := &cobra.Command{
		Use:     "list",
		Short:   "Print list of dataset definitions",
		Long:    "Print list of dataset definitions.",
		Example: "datamkr dataset list",
		Run: func(cmd *cobra.Command, args []string) {
			utils.CheckErr(datasetListOptions.Complete(cmd, args))
			utils.CheckErr(datasetListOptions.Validate())
			utils.CheckErr(datasetListOptions.Run())
		},
	}

	return cmd
}

func (opt *DatasetListOptions) Complete(cmd *cobra.Command, args []string) error {
	currentConfig, err := opt.factory.GetConfig()
	if err != nil {
		return err
	}
	opt.datamkrClient = client.NewWithConfig(currentConfig)

	return nil
}

func (opt *DatasetListOptions) Validate() error {
	return nil
}

func (opt *DatasetListOptions) Run() error {
	datasetService := opt.datamkrClient.Datasets()
	datasets, err := datasetService.List()
	if err != nil {
		return err
	}

	if len(datasets) == 0 {
		fmt.Print("No datasets found, to add a dataset (example):\n\n   datamkr dataset add users --field name=id,type=uuid --field name=email,type=email\n\n")
		return nil
	}

	for _, d := range datasets {
		fmt.Println((d))
	}
	return nil
}
