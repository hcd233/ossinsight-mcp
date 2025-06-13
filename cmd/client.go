package cmd

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/hcd233/ossinsight-mcp/internal/config"
	"github.com/hcd233/ossinsight-mcp/internal/logger"
	"github.com/mark3labs/mcp-go/client"
	"github.com/mark3labs/mcp-go/client/transport"
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/samber/lo"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
)

var clientCmd = &cobra.Command{
	Use:   "client",
	Short: "Start MCP Client",
	Long:  `Start MCP Client`,
	Run: func(cmd *cobra.Command, _ []string) {
		endpoint := lo.Must1(cmd.Flags().GetString("endpoint"))

		cli, err := client.NewStreamableHttpClient(endpoint, transport.WithHTTPHeaders(map[string]string{
			"X-Trace-Id":    uuid.NewString(),
			"Authorization": fmt.Sprintf("Bearer %s", config.APIKey),
		}))
		if err != nil {
			logger.Logger().Error("[Client CMD] create client failed", zap.Error(err))
			return
		}

		defer cli.Close()

		ctx := context.Background()

		cli.Initialize(ctx, mcp.InitializeRequest{})

		tools, err := cli.ListTools(ctx, mcp.ListToolsRequest{})
		if err != nil {
			logger.Logger().Error("[Client CMD] list tools failed", zap.Error(err))
			return
		}

		logger.Logger().Info("[Client CMD] list tools", zap.Any("tools", tools))

		result, err := cli.CallTool(ctx, mcp.CallToolRequest{
			Params: mcp.CallToolParams{
				Name: "ListTrendingRepos",
				Arguments: map[string]any{
					"period":   "past_24_hours",
					"language": "Go",
					"limit":    10,
				},
			},
		})
		if err != nil {
			logger.Logger().Error("[Client CMD] call tool failed", zap.Error(err))
			return
		}

		logger.Logger().Info("[Client CMD] call tool", zap.Any("result", result))
	},
}

func init() {
	rootCmd.AddCommand(clientCmd)

	clientCmd.Flags().StringP("endpoint", "e", "http://localhost:8080/mcp", "MCP Server URL")
}
