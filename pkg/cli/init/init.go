package init

import (
	"fmt"
	"log"

	utils "github.com/awile/datamkr/pkg/cli/util"
	"github.com/awile/datamkr/pkg/config"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

type InitOptions struct {
	HasConfig bool
	factory   config.ConfigFactory
}

func NewInitOptions() *InitOptions {
	dcf, err := config.NewDatamkrConfigFactory()
	if err != nil {
		log.Fatal(err)
	}
	return &InitOptions{factory: dcf}
}

func NewInitCmd(configFactory *config.DatamkrConfigFactory) *cobra.Command {
	initOptions := NewInitOptions()

	cmd := &cobra.Command{
		Use:     "init",
		Short:   "Creates a project config file",
		Long:    "Creates a project config file.",
		Example: "datamkr init",
		Run: func(cmd *cobra.Command, args []string) {
			utils.CheckErr(initOptions.Complete(cmd, args))
			utils.CheckErr(initOptions.Validate())
			utils.CheckErr(initOptions.Run())
		},
	}

	return cmd
}

func (options *InitOptions) Complete(cmd *cobra.Command, args []string) error {
	version := viper.Get("version")
	options.HasConfig = version != nil
	return nil
}

func (options *InitOptions) Validate() error {
	return nil
}

func (options *InitOptions) Run() error {
	if options.HasConfig {
		fmt.Println("Config file already exists at ./datamkr.yml")
		return nil
	}
	configFile := options.factory.CreateNewConfigFile()
	err := options.factory.InitDatamkrConfigFile(configFile)
	if err != nil {
		return err
	}
	fmt.Println("Config file created.")
	return nil
}
