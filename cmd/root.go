/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"os"

	"github.com/hcd233/ossinsight-mcp/internal/logger"
	"github.com/hcd233/ossinsight-mcp/internal/service"
	"github.com/hcd233/ossinsight-mcp/internal/util"
	"github.com/mark3labs/mcp-go/server"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "ossinsight-mcp",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(_ *cobra.Command, _ []string) {
		s := server.NewMCPServer(
			"Demo Server",
			"1.0.0",
			server.WithToolCapabilities(true),
			// server.WithResourceCapabilities(false, true),
			// server.WithPromptCapabilities(true),
		)

		service.RegisterTools(s)

		logger.Logger().Info("[Root CMD] start ossinsight-mcp server")
		if err := server.NewStreamableHTTPServer(s, server.WithHTTPContextFunc(util.WithTraceID)).Start("0.0.0.0:8080"); err != nil {
			logger.Logger().Error("[Root CMD] serve stdio failed", zap.Error(err))
		}
	},
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
