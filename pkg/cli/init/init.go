package init

import (
    "fmt"
    "log"
    "github.com/spf13/cobra"
    "github.com/awile/datamkr/pkg/config"
)

func AddInitCmd() *cobra.Command {

    cmd := &cobra.Command{
        Use:     "init",
        Short:   "Creates a project config file",
        Long:    "Creates a project config file.",
        Example: "datamkr init",
        Run: func(cmd *cobra.Command, args []string) {
            hasConfigFile, err := config.CreateDatamkrConfigFile()
            if err != nil {
                log.Fatal(err)
            }

            msg := fmt.Sprintf("Creating config file: %s", config.ConfigFileName)
            if hasConfigFile {
                msg = fmt.Sprintf("Config file already exists at %s", config.ConfigFileName)
            }
            fmt.Println(msg)
        },
    }

    return cmd
}

