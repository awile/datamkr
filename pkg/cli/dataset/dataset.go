package dataset

import (
	utils "github.com/awile/datamkr/pkg/cli/util"
	"github.com/awile/datamkr/pkg/config"
	"github.com/spf13/cobra"
)

type DatasetOptions struct {
}

func NewDatasetOptions(factory *config.DatamkrConfigFactory) *DatasetOptions {
	return &DatasetOptions{}
}

func NewDatasetCmd(configFactory *config.DatamkrConfigFactory) *cobra.Command {
	datasetOptions := NewDatasetOptions(configFactory)

	cmd := &cobra.Command{
		Use:     "dataset",
		Short:   "Manage dataset definitions",
		Long:    "Manage dataset definitions.",
		Example: "datamkr dataset",
		Run: func(cmd *cobra.Command, args []string) {
			utils.CheckErr(datasetOptions.Complete(cmd, args))
			utils.CheckErr(datasetOptions.Validate())
			utils.CheckErr(datasetOptions.Run())
		},
	}

	cmd.AddCommand(NewDatasetListCmd(configFactory))

	return cmd
}

func (opt *DatasetOptions) Complete(cmd *cobra.Command, args []string) error {
	return nil
}

func (opt *DatasetOptions) Validate() error {
	return nil
}

func (opt *DatasetOptions) Run() error {
	return nil
}
