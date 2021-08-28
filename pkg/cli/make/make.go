package make

import (
	utils "github.com/awile/datamkr/pkg/cli/util"
	"github.com/awile/datamkr/pkg/client"
	"github.com/awile/datamkr/pkg/config"
	"github.com/awile/datamkr/pkg/dataset"
	"github.com/spf13/cobra"
)

type MakeOptions struct {
	DatasetDefinition dataset.DatasetDefinition

	factory       config.ConfigFactory
	datamkrClient client.Interface
}

func NewMakeOptions(factory *config.DatamkrConfigFactory) *MakeOptions {
	return &MakeOptions{factory: factory}
}

func newMakeCmd(configFactory *config.DatamkrConfigFactory) *cobra.Command {
	makeOptions := NewMakeOptions(configFactory)

	cmd := &cobra.Command{
		Use:     "make",
		Short:   "Generate a dataset based on given definition",
		Long:    "Generate dataset based on given definition",
		Example: "datamkr make <dataset_name>",
		Run: func(cmd *cobra.Command, args []string) {
			utils.CheckErr(makeOptions.Complete(cmd, args))
			utils.CheckErr(makeOptions.Validate())
			utils.CheckErr(makeOptions.Run())
		},
	}

	return cmd
}

func (opt *MakeOptions) Complete(cmd *cobra.Command, args []string) error {
	return nil
}

func (opt *MakeOptions) Validate() error {
	return nil
}

func (opt *MakeOptions) Run() error {
	return nil
}
