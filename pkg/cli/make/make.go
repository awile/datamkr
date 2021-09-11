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
		Example: "  datamkr make <dataset_name>",
		Run: func(cmd *cobra.Command, args []string) {
			utils.CheckErr(makeOptions.Complete(cmd, args))
			utils.CheckErr(makeOptions.Validate())
			utils.CheckErr(makeOptions.Run())
		},
	}

	cmd.Flags().Int32VarP(&makeOptions.NumberOfRows, "num", "n", 10, "Number of rows to generate")
	cmd.Flags().StringVarP(&makeOptions.Target, "to", "t", "", "Where generated data should go")
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
		return errors.New("Must give which dataset to use:\n\n    datamkr make <dataset_name>\n\nTo check datasets use command: datamkr dataset list\n")
	}

	if strings.Contains(opt.Target, "postgresql://") {
		opt.Type = "postgres"
	} else if len(opt.Target) > 4 && opt.Target[len(opt.Target)-4:] == ".csv" {
		opt.Type = "csv"
	} else {
		opt.Target = fmt.Sprintf("%s.csv", args[0])
		opt.Type = "csv"
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
	writerOptions.Id = opt.Target

	storageWriter := storageClient.GetStorageWriterService(opt.Type, writerOptions)
	if storageWriter == nil {
		return fmt.Errorf("%s is not a valid target\n", opt.Type)
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

	fmt.Printf("Created %d rows in %s\n", opt.NumberOfRows, opt.Target)
	return nil
}
