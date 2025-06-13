// package service 服务
//
//	@author centonhuang
//	@update 2025-06-13 16:50:01
package service

import (
	"github.com/hcd233/ossinsight-mcp/internal/logger"
	"github.com/hcd233/ossinsight-mcp/internal/repo/mcp"
	"github.com/hcd233/ossinsight-mcp/internal/repo/ossinsight"
	"github.com/mark3labs/mcp-go/server"
	"go.uber.org/zap"
)

// RegisterTools 注册工具
//
//	@param s
//	@author centonhuang
//	@update 2025-06-12 21:06:34
func RegisterTools(s *server.MCPServer) {
	handler := mcp.NewOssinsightMCPHandler(ossinsight.NewClient())

	schema := mcp.ListTrendingReposSchema()
	s.AddTool(schema, handler.ListTrendingRepos)
	logger.Logger().Info("[Server] register tool", zap.String("tool", schema.Name), zap.String("description", schema.Description))

	logger.Logger().Info("[Server] register tool success")
}
