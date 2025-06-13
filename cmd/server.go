package cmd

import (
	"fmt"

	"github.com/hcd233/ossinsight-mcp/internal/logger"
	"github.com/hcd233/ossinsight-mcp/internal/service"
	"github.com/hcd233/ossinsight-mcp/internal/util"
	"github.com/mark3labs/mcp-go/server"
	"github.com/samber/lo"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
)

var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "Start MCP Server",
	Long:  `Start MCP Server`,
	Run: func(cmd *cobra.Command, _ []string) {
		host, port := lo.Must1(cmd.Flags().GetString("host")), lo.Must1(cmd.Flags().GetString("port"))
		addr := fmt.Sprintf("%s:%s", host, port)

		s := server.NewMCPServer(
			"Demo Server",
			"1.0.0",
			server.WithToolCapabilities(true),
			// server.WithResourceCapabilities(false, true),
			// server.WithPromptCapabilities(true),
		)

		service.RegisterTools(s)

		logger.Logger().Info("[Server CMD] start ossinsight-mcp server", zap.String("addr", addr))
		if err := server.NewStreamableHTTPServer(s, server.WithHTTPContextFunc(util.WithTraceID)).Start(addr); err != nil {
			logger.Logger().Error("[Root CMD] serve stdio failed", zap.Error(err))
		}
	},
}

func init() {
	rootCmd.AddCommand(serverCmd)

	serverCmd.Flags().StringP("host", "", "localhost", "Listen host")
	serverCmd.Flags().StringP("port", "p", "8080", "Listen port")
}
