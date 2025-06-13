package middleware

import (
	"context"

	"github.com/hcd233/ossinsight-mcp/internal/config"
	"github.com/hcd233/ossinsight-mcp/internal/constant"
	"github.com/hcd233/ossinsight-mcp/internal/logger"
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

// BearerAuthMiddleware 认证中间件
//
//	@return server.ToolHandlerMiddleware
//	@author centonhuang
//	@update 2025-06-13 16:46:43
func BearerAuthMiddleware() server.ToolHandlerMiddleware {
	return func(thf server.ToolHandlerFunc) server.ToolHandlerFunc {
		return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			logger := logger.ContextLogger(ctx)

			apiKey, ok := ctx.Value(constant.CtxKeyAPIKey).(string)
			if !ok {
				logger.Error("[BearerAuthMiddleware] require credentials")
				return nil, constant.ErrRequireCredentials.Errorf()
			}
			if apiKey != config.APIKey {
				logger.Error("[BearerAuthMiddleware] invalid credentials")
				return nil, constant.ErrInvalidCredentials.Errorf()
			}
			return thf(ctx, request)
		}
	}
}
