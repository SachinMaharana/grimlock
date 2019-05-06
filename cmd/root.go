package cmd

import (
	"os"

	"github.com/spf13/cobra"
	"go.uber.org/zap"
)

var (
	logger         *zap.Logger
	metricsBackend string
)

// RootCmd is the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "examples-grimlock",
	Short: "Grimlock - Tranformer",
	Long:  "Grimlock - Tranformers",
}

// Execute adds all the child commands to the root command and set lags appropriately.
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		logger.Fatal("Error.", zap.Error(err))
		os.Exit(-1)
	}
}

// func init() {
// 	RootCmd.PersistentFlags().StringVarP()
// }

func onInitialize() {

}

func logError(logger *zap.Logger, err error) error {
	if err != nil {
		logger.Error("Error running command", zap.Error(err))
	}
	return err
}
