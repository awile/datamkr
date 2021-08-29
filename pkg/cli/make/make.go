package make

import (
	"errors"
	"fmt"

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
	Name              string
	NumberOfRows      int32

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
	if opt.Name == "" {
		opt.Name = args[0]
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

	csvWriter, err := storageClient.GetStorageService(opt.Target)
	if err != nil {
		return err
	}

	var args storage.StorageArgs
	args.FileName = fmt.Sprintf("%s.csv", opt.Name)
	args.IsWriter = true

	err = csvWriter.Init(args)
	if err != nil {
		return err
	}
	defer csvWriter.Close()

	fieldKeys := make([]string, len(opt.DatasetDefinition.Fields))
	var idx int = 0
	for fieldKey := range opt.DatasetDefinition.Fields {
		fieldKeys[idx] = fieldKey
		idx++
	}
	err = csvWriter.Write(fieldKeys)
	if err != nil {
		return err
	}

	orderedRows := make([][]string, opt.NumberOfRows)
	for i := 0; i < int(opt.NumberOfRows); i++ {
		row, err := makerClient.MakeRow(opt.DatasetDefinition)
		if err != nil {
			return err
		}
		orderedRow := make([]string, len(fieldKeys))
		for j, fieldKey := range fieldKeys {
			value := row[fieldKey]
			orderedRow[j] = value.(string)
		}
		orderedRows[i] = orderedRow
	}
	err = csvWriter.WriteAll(orderedRows)
	if err != nil {
		return err
	}

	return nil
}
