package cmd

import (
	"fmt"

	"github.com/hcd233/ossinsight-mcp/internal/logger"
	"github.com/hcd233/ossinsight-mcp/internal/middleware"
	"github.com/hcd233/ossinsight-mcp/internal/repo/mcp"
	"github.com/hcd233/ossinsight-mcp/internal/service"
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
			"OSSInsight API Server",
			"1.0.0",
			server.WithInstructions("This is OSSInsight MCP Server. OSSInsight is a powerful analytics platform for GitHub, providing real-time insights and analytics about GitHub repositories, developers, and trends. It helps you discover trending repositories, analyze repository growth, and understand developer behaviors across the GitHub ecosystem."),
			server.WithToolCapabilities(true),
			// server.WithResourceCapabilities(false, true),
			// server.WithPromptCapabilities(true),
			server.WithToolHandlerMiddleware(middleware.BearerAuthMiddleware()),
		)

		service.RegisterTools(s)

		svr := server.NewStreamableHTTPServer(s,
			server.WithHTTPContextFunc(mcp.HTTPContextFunc),
		)

		logger.Logger().Info("[Server CMD] start ossinsight-mcp server", zap.String("addr", addr))
		if err := svr.Start(addr); err != nil {
			logger.Logger().Error("[Root CMD] serve stdio failed", zap.Error(err))
		}
	},
}

func init() {
	rootCmd.AddCommand(serverCmd)

	serverCmd.Flags().StringP("host", "", "localhost", "Listen host")
	serverCmd.Flags().StringP("port", "p", "8080", "Listen port")
}
