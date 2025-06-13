// Package cmd 命令行
//
//	@author centonhuang
//	@update 2025-06-13 16:26:27
package cmd

import (
	"os"

	"github.com/hcd233/ossinsight-mcp/internal/logger"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "ossinsight-mcp",
	Short: "Welcome to Ossinsight MCP Server",
	Long:  `Welcome to Ossinsight MCP Server`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		logger.Logger().Error("[Root CMD] execute failed", zap.Error(err))
		os.Exit(1)
	}
}
