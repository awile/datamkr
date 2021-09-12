package init

import (
	"fmt"
	"log"
	"os"

	utils "github.com/awile/datamkr/pkg/cli/util"
	"github.com/awile/datamkr/pkg/client"
	"github.com/awile/datamkr/pkg/config"
	"github.com/awile/datamkr/pkg/dataset"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

type InitOptions struct {
	HasConfig     bool
	factory       config.ConfigFactory
	datamkrClient client.Interface
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

	currentConfig, err := options.factory.GetConfig()
	if err != nil {
		return err
	}
	options.datamkrClient = client.NewWithConfig(currentConfig)

	datasetPath := "./datasets"
	if _, dirErr := os.Stat(datasetPath); os.IsNotExist(dirErr) {
		err = os.MkdirAll(datasetPath, os.ModePerm)
		if err != nil {
			return err
		}
		options.createDemoDataDefinition()
	}

	fmt.Println("Config file created.")
	return nil
}

func (options *InitOptions) createDemoDataDefinition() error {
	datasetClient := options.datamkrClient.Datasets()

	datasetFields := map[string]dataset.DatasetDefinitionField{
		"id":        {Type: "uuid"},
		"name":      {Type: "name"},
		"email":     {Type: "email"},
		"password":  {Type: "string"},
		"isDeleted": {Type: "boolean"},
	}
	datasetDefinition := dataset.DatasetDefinition{Fields: datasetFields}
	return datasetClient.Add("demo", datasetDefinition)
}
