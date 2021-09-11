package make

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

type MakeOptions struct {
	DatasetDefinition dataset.DatasetDefinition
	Target            string
	Type              string
	Name              string
	NumberOfRows      int32
	Fields            string

	factory       config.ConfigFactory
	datamkrClient client.Interface
}

func NewMakeOptions(factory *config.DatamkrConfigFactory) *MakeOptions {
	return &MakeOptions{factory: factory}
}

func NewMakeCmd(configFactory *config.DatamkrConfigFactory) *cobra.Command {
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

	cmd.Flags().Int32VarP(&makeOptions.NumberOfRows, "rows", "r", 10, "Number of rows to generate")
	cmd.Flags().StringVarP(&makeOptions.Name, "name", "n", "", "Name of file to generate")
	cmd.Flags().StringVarP(&makeOptions.Target, "target", "t", "csv", "Where generated data should go")
	cmd.Flags().StringVarP(&makeOptions.Fields, "fields", "f", "", "Fields to include (--fields a,b,c)")

	return cmd
}

func (opt *MakeOptions) Complete(cmd *cobra.Command, args []string) error {
	currentConfig, err := opt.factory.GetConfig()
	if err != nil {
		return err
	}
	opt.datamkrClient = client.NewWithConfig(currentConfig)
	datasetClient := opt.datamkrClient.Datasets()

	if len(args) == 0 {
		return errors.New("Must give dataset a name:\n\n    datamkr make <dataset_name>\n\n")
	}

	if opt.Target == "csv" {
		opt.Type = "csv"
	} else if strings.Contains(opt.Target, "postgresql://") {
		opt.Type = "postgres"
	}

	if opt.Name == "" {
		if opt.Type == "csv" {
			opt.Name = fmt.Sprintf("%s.csv", args[0])
		} else {
			opt.Name = args[0]
		}
	}

	datasetDefinition, err := datasetClient.Get(args[0])
	if err != nil {
		return err
	}
	opt.DatasetDefinition = datasetDefinition

	return nil
}

func (opt *MakeOptions) Validate() error {
	return nil
}

func (opt *MakeOptions) Run() error {
	makerClient := opt.datamkrClient.Maker()
	storageClient := opt.datamkrClient.Storage()

	writerOptions := storage.CreateWriterOptions()
	writerOptions.DatasetDefinition = opt.DatasetDefinition
	if opt.Target == "csv" {
		writerOptions.Id = opt.Name
	} else if strings.Contains(opt.Target, "postgresql://") {
		writerOptions.Id = opt.Target
	}

	storageWriter := storageClient.GetStorageWriterService(opt.Type, writerOptions)
	if storageWriter == nil {
		return fmt.Errorf("%s is not a valid target\n", opt.Target)
	}

	err := storageWriter.Init()
	if err != nil {
		return err
	}
	defer storageWriter.Close()

	for i := 0; i < int(opt.NumberOfRows); i++ {
		row, err := makerClient.MakeRow(opt.DatasetDefinition)
		if err != nil {
			return err
		}
		err = storageWriter.Write(row)
		if err != nil {
			return err
		}
	}

	return nil
}
