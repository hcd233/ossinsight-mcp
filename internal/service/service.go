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

	// 趋势仓库工具
	schema := mcp.ListTrendingReposSchema()
	s.AddTool(schema, handler.ListTrendingRepos)
	logger.Logger().Info("[Server] register tool", zap.String("tool", schema.Name), zap.String("description", schema.Description))

	// 仓库排名工具
	repoRankingSchema := mcp.GetRepoRankingSchema()
	s.AddTool(repoRankingSchema, handler.GetRepoRanking)
	logger.Logger().Info("[Server] register tool", zap.String("tool", repoRankingSchema.Name), zap.String("description", repoRankingSchema.Description))

	// 开发者排名工具
	developerRankingSchema := mcp.GetDeveloperRankingSchema()
	s.AddTool(developerRankingSchema, handler.GetDeveloperRanking)
	logger.Logger().Info("[Server] register tool", zap.String("tool", developerRankingSchema.Name), zap.String("description", developerRankingSchema.Description))

	// 收藏夹列表工具
	collectionsSchema := mcp.GetCollectionsSchema()
	s.AddTool(collectionsSchema, handler.GetCollections)
	logger.Logger().Info("[Server] register tool", zap.String("tool", collectionsSchema.Name), zap.String("description", collectionsSchema.Description))

	// 热门收藏夹工具
	hotCollectionsSchema := mcp.GetHotCollectionsSchema()
	s.AddTool(hotCollectionsSchema, handler.GetHotCollections)
	logger.Logger().Info("[Server] register tool", zap.String("tool", hotCollectionsSchema.Name), zap.String("description", hotCollectionsSchema.Description))

	// 收藏夹仓库列表工具
	collectionReposSchema := mcp.GetCollectionReposSchema()
	s.AddTool(collectionReposSchema, handler.GetCollectionRepos)
	logger.Logger().Info("[Server] register tool", zap.String("tool", collectionReposSchema.Name), zap.String("description", collectionReposSchema.Description))

	// 仓库详情工具
	repoDetailSchema := mcp.GetRepoDetailSchema()
	s.AddTool(repoDetailSchema, handler.GetRepoDetail)
	logger.Logger().Info("[Server] register tool", zap.String("tool", repoDetailSchema.Name), zap.String("description", repoDetailSchema.Description))

	// 开发者详情工具
	developerDetailSchema := mcp.GetDeveloperDetailSchema()
	s.AddTool(developerDetailSchema, handler.GetDeveloperDetail)
	logger.Logger().Info("[Server] register tool", zap.String("tool", developerDetailSchema.Name), zap.String("description", developerDetailSchema.Description))

	// 语言统计工具
	languageStatsSchema := mcp.GetLanguageStatsSchema()
	s.AddTool(languageStatsSchema, handler.GetLanguageStats)
	logger.Logger().Info("[Server] register tool", zap.String("tool", languageStatsSchema.Name), zap.String("description", languageStatsSchema.Description))

	logger.Logger().Info("[Server] register tool success")
}
