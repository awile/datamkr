package dataset

import (
	"github.com/awile/datamkr/pkg/config"
	"github.com/spf13/cobra"
)

func NewDatasetCmd(configFactory *config.DatamkrConfigFactory) *cobra.Command {

	cmd := &cobra.Command{
		Use:   "dataset",
		Short: "Manage dataset definitions",
		Long:  "Manage dataset definitions",
	}

	cmd.AddCommand(NewDatasetListCmd(configFactory))
	cmd.AddCommand(NewDatasetAddCmd(configFactory))

	return cmd
}
