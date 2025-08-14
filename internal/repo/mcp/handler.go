package mcp

import (
	"context"
	"encoding/json"

	"github.com/hcd233/ossinsight-mcp/internal/constant"
	"github.com/hcd233/ossinsight-mcp/internal/logger"
	"github.com/hcd233/ossinsight-mcp/internal/repo/ossinsight"
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/samber/lo"
	"go.uber.org/zap"
)

// OssinsightMCPHandler Ossinsight MCP 处理器
//
//	@author centonhuang
//	@update 2025-06-12 20:46:51
type OssinsightMCPHandler struct {
	cli *ossinsight.Client
}

// NewOssinsightMCPHandler 创建 Ossinsight MCP 处理器
//
//	@param ossinsightClient
//	@return *OssinsightMCPHandler
//	@author centonhuang
//	@update 2025-06-12 20:47:00
func NewOssinsightMCPHandler(ossinsightClient *ossinsight.Client) *OssinsightMCPHandler {
	return &OssinsightMCPHandler{cli: ossinsightClient}
}

// ListTrendingRepos 获取趋势仓库列表
//
//	@receiver h *OssinsightMCPHandler
//	@param ctx
//	@param request
//	@return *mcp.CallToolResult
//	@return error
//	@author centonhuang
//	@update 2025-06-12 20:57:36
func (h *OssinsightMCPHandler) ListTrendingRepos(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	logger := logger.ContextLogger(ctx)

	logger.Info("[ListTrendingRepos] receive request", zap.Any("params", request.Params))
	response := &mcp.CallToolResult{Result: mcp.Result{
		Meta: map[string]any{
			constant.MetaKeyTraceID: ctx.Value(constant.MetaKeyTraceID),
		},
	}}
	period, err := request.RequireString("period")
	if err != nil {
		logger.Error("[ListTrendingRepos] period is required", zap.Error(err))
		response.IsError = true
		return response, constant.ErrRequiredArgumentNotFound.Errorf("period")
	}

	language, err := request.RequireString("language")
	if err != nil {
		logger.Error("[ListTrendingRepos] language is required", zap.Error(err))
		response.IsError = true
		return response, constant.ErrRequiredArgumentNotFound.Errorf("language")
	}

	limit, err := request.RequireInt("limit")
	if err != nil {
		logger.Error("[ListTrendingRepos] limit is required", zap.Error(err))
		response.IsError = true
		return response, constant.ErrRequiredArgumentNotFound.Errorf("limit")
	}

	req := &ossinsight.ListTrendingReposQuery{
		Period:   period,
		Language: language,
	}

	rsp, _, err := h.cli.ListTrendingRepos(ctx, req, limit)
	if err != nil {
		logger.Error("[ListTrendingRepos] list trending repos failed", zap.Error(err))
		response.IsError = true
		return response, err
	}

	content := mcp.NewTextContent(string(lo.Must1(json.MarshalIndent(rsp.Data.Rows, "", "  "))))

	logger.Info("[ListTrendingRepos] list trending repos success", zap.Any("content", content))
	response.Content = []mcp.Content{content}
	response.IsError = false
	return response, nil
}

// GetRepoRanking 获取仓库排名
func (h *OssinsightMCPHandler) GetRepoRanking(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	logger := logger.ContextLogger(ctx)

	logger.Info("[GetRepoRanking] receive request", zap.Any("params", request.Params))
	response := &mcp.CallToolResult{Result: mcp.Result{
		Meta: map[string]any{
			constant.MetaKeyTraceID: ctx.Value(constant.MetaKeyTraceID),
		},
	}}

	rankingType, err := request.RequireString("ranking_type")
	if err != nil {
		logger.Error("[GetRepoRanking] ranking_type is required", zap.Error(err))
		response.IsError = true
		return response, constant.ErrRequiredArgumentNotFound.Errorf("ranking_type")
	}

	timeRange, err := request.RequireString("time_range")
	if err != nil {
		logger.Error("[GetRepoRanking] time_range is required", zap.Error(err))
		response.IsError = true
		return response, constant.ErrRequiredArgumentNotFound.Errorf("time_range")
	}

	language := request.GetString("language", "")
	limit := request.GetInt("limit", 10)

	req := &ossinsight.RepoRankingQuery{
		RankingType: rankingType,
		TimeRange:   timeRange,
		Language:    language,
	}

	rsp, _, err := h.cli.GetRepoRanking(ctx, req, limit)
	if err != nil {
		logger.Error("[GetRepoRanking] get repo ranking failed", zap.Error(err))
		response.IsError = true
		return response, err
	}

	content := mcp.NewTextContent(string(lo.Must1(json.MarshalIndent(rsp.Data.Rows, "", "  "))))

	logger.Info("[GetRepoRanking] get repo ranking success", zap.Any("content", content))
	response.Content = []mcp.Content{content}
	response.IsError = false
	return response, nil
}

// GetDeveloperRanking 获取开发者排名
func (h *OssinsightMCPHandler) GetDeveloperRanking(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	logger := logger.ContextLogger(ctx)

	logger.Info("[GetDeveloperRanking] receive request", zap.Any("params", request.Params))
	response := &mcp.CallToolResult{Result: mcp.Result{
		Meta: map[string]any{
			constant.MetaKeyTraceID: ctx.Value(constant.MetaKeyTraceID),
		},
	}}

	rankingType, err := request.RequireString("ranking_type")
	if err != nil {
		logger.Error("[GetDeveloperRanking] ranking_type is required", zap.Error(err))
		response.IsError = true
		return response, constant.ErrRequiredArgumentNotFound.Errorf("ranking_type")
	}

	timeRange, err := request.RequireString("time_range")
	if err != nil {
		logger.Error("[GetDeveloperRanking] time_range is required", zap.Error(err))
		response.IsError = true
		return response, constant.ErrRequiredArgumentNotFound.Errorf("time_range")
	}

	language := request.GetString("language", "")
	limit := request.GetInt("limit", 10)

	req := &ossinsight.DeveloperRankingQuery{
		RankingType: rankingType,
		TimeRange:   timeRange,
		Language:    language,
	}

	rsp, _, err := h.cli.GetDeveloperRanking(ctx, req, limit)
	if err != nil {
		logger.Error("[GetDeveloperRanking] get developer ranking failed", zap.Error(err))
		response.IsError = true
		return response, err
	}

	content := mcp.NewTextContent(string(lo.Must1(json.MarshalIndent(rsp.Data.Rows, "", "  "))))

	logger.Info("[GetDeveloperRanking] get developer ranking success", zap.Any("content", content))
	response.Content = []mcp.Content{content}
	response.IsError = false
	return response, nil
}

// GetCollections 获取收藏夹列表
func (h *OssinsightMCPHandler) GetCollections(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	logger := logger.ContextLogger(ctx)

	logger.Info("[GetCollections] receive request", zap.Any("params", request.Params))
	response := &mcp.CallToolResult{Result: mcp.Result{
		Meta: map[string]any{
			constant.MetaKeyTraceID: ctx.Value(constant.MetaKeyTraceID),
		},
	}}

	collectionType, err := request.RequireString("collection_type")
	if err != nil {
		logger.Error("[GetCollections] collection_type is required", zap.Error(err))
		response.IsError = true
		return response, constant.ErrRequiredArgumentNotFound.Errorf("collection_type")
	}

	timeRange, err := request.RequireString("time_range")
	if err != nil {
		logger.Error("[GetCollections] time_range is required", zap.Error(err))
		response.IsError = true
		return response, constant.ErrRequiredArgumentNotFound.Errorf("time_range")
	}

	limit := request.GetInt("limit", 10)

	req := &ossinsight.CollectionsQuery{
		CollectionType: collectionType,
		TimeRange:      timeRange,
	}

	rsp, _, err := h.cli.GetCollections(ctx, req, limit)
	if err != nil {
		logger.Error("[GetCollections] get collections failed", zap.Error(err))
		response.IsError = true
		return response, err
	}

	content := mcp.NewTextContent(string(lo.Must1(json.MarshalIndent(rsp.Data.Rows, "", "  "))))

	logger.Info("[GetCollections] get collections success", zap.Any("content", content))
	response.Content = []mcp.Content{content}
	response.IsError = false
	return response, nil
}

// GetHotCollections 获取热门收藏夹
func (h *OssinsightMCPHandler) GetHotCollections(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	logger := logger.ContextLogger(ctx)

	logger.Info("[GetHotCollections] receive request", zap.Any("params", request.Params))
	response := &mcp.CallToolResult{Result: mcp.Result{
		Meta: map[string]any{
			constant.MetaKeyTraceID: ctx.Value(constant.MetaKeyTraceID),
		},
	}}

	timeRange, err := request.RequireString("time_range")
	if err != nil {
		logger.Error("[GetHotCollections] time_range is required", zap.Error(err))
		response.IsError = true
		return response, constant.ErrRequiredArgumentNotFound.Errorf("time_range")
	}

	limit := request.GetInt("limit", 10)

	req := &ossinsight.CollectionsQuery{
		CollectionType: ossinsight.CollectionTypeHot,
		TimeRange:      timeRange,
	}

	rsp, _, err := h.cli.GetHotCollections(ctx, req, limit)
	if err != nil {
		logger.Error("[GetHotCollections] get hot collections failed", zap.Error(err))
		response.IsError = true
		return response, err
	}

	content := mcp.NewTextContent(string(lo.Must1(json.MarshalIndent(rsp.Data.Rows, "", "  "))))

	logger.Info("[GetHotCollections] get hot collections success", zap.Any("content", content))
	response.Content = []mcp.Content{content}
	response.IsError = false
	return response, nil
}

// GetCollectionRepos 获取收藏夹仓库列表
func (h *OssinsightMCPHandler) GetCollectionRepos(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	logger := logger.ContextLogger(ctx)

	logger.Info("[GetCollectionRepos] receive request", zap.Any("params", request.Params))
	response := &mcp.CallToolResult{Result: mcp.Result{
		Meta: map[string]any{
			constant.MetaKeyTraceID: ctx.Value(constant.MetaKeyTraceID),
		},
	}}

	collectionID, err := request.RequireString("collection_id")
	if err != nil {
		logger.Error("[GetCollectionRepos] collection_id is required", zap.Error(err))
		response.IsError = true
		return response, constant.ErrRequiredArgumentNotFound.Errorf("collection_id")
	}

	timeRange, err := request.RequireString("time_range")
	if err != nil {
		logger.Error("[GetCollectionRepos] time_range is required", zap.Error(err))
		response.IsError = true
		return response, constant.ErrRequiredArgumentNotFound.Errorf("time_range")
	}

	limit := request.GetInt("limit", 10)

	req := &ossinsight.CollectionReposQuery{
		CollectionID: collectionID,
		TimeRange:    timeRange,
	}

	rsp, _, err := h.cli.GetCollectionRepos(ctx, req, limit)
	if err != nil {
		logger.Error("[GetCollectionRepos] get collection repos failed", zap.Error(err))
		response.IsError = true
		return response, err
	}

	content := mcp.NewTextContent(string(lo.Must1(json.MarshalIndent(rsp.Data.Rows, "", "  "))))

	logger.Info("[GetCollectionRepos] get collection repos success", zap.Any("content", content))
	response.Content = []mcp.Content{content}
	response.IsError = false
	return response, nil
}

// GetRepoDetail 获取仓库详情
func (h *OssinsightMCPHandler) GetRepoDetail(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	logger := logger.ContextLogger(ctx)

	logger.Info("[GetRepoDetail] receive request", zap.Any("params", request.Params))
	response := &mcp.CallToolResult{Result: mcp.Result{
		Meta: map[string]any{
			constant.MetaKeyTraceID: ctx.Value(constant.MetaKeyTraceID),
		},
	}}

	owner, err := request.RequireString("owner")
	if err != nil {
		logger.Error("[GetRepoDetail] owner is required", zap.Error(err))
		response.IsError = true
		return response, constant.ErrRequiredArgumentNotFound.Errorf("owner")
	}

	repo, err := request.RequireString("repo")
	if err != nil {
		logger.Error("[GetRepoDetail] repo is required", zap.Error(err))
		response.IsError = true
		return response, constant.ErrRequiredArgumentNotFound.Errorf("repo")
	}

	req := &ossinsight.RepoDetailQuery{
		Owner: owner,
		Repo:  repo,
	}

	rsp, _, err := h.cli.GetRepoDetail(ctx, req)
	if err != nil {
		logger.Error("[GetRepoDetail] get repo detail failed", zap.Error(err))
		response.IsError = true
		return response, err
	}

	content := mcp.NewTextContent(string(lo.Must1(json.MarshalIndent(rsp.Data.Rows, "", "  "))))

	logger.Info("[GetRepoDetail] get repo detail success", zap.Any("content", content))
	response.Content = []mcp.Content{content}
	response.IsError = false
	return response, nil
}

// GetDeveloperDetail 获取开发者详情
func (h *OssinsightMCPHandler) GetDeveloperDetail(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	logger := logger.ContextLogger(ctx)

	logger.Info("[GetDeveloperDetail] receive request", zap.Any("params", request.Params))
	response := &mcp.CallToolResult{Result: mcp.Result{
		Meta: map[string]any{
			constant.MetaKeyTraceID: ctx.Value(constant.MetaKeyTraceID),
		},
	}}

	username, err := request.RequireString("username")
	if err != nil {
		logger.Error("[GetDeveloperDetail] username is required", zap.Error(err))
		response.IsError = true
		return response, constant.ErrRequiredArgumentNotFound.Errorf("username")
	}

	req := &ossinsight.DeveloperDetailQuery{
		Username: username,
	}

	rsp, _, err := h.cli.GetDeveloperDetail(ctx, req)
	if err != nil {
		logger.Error("[GetDeveloperDetail] get developer detail failed", zap.Error(err))
		response.IsError = true
		return response, err
	}

	content := mcp.NewTextContent(string(lo.Must1(json.MarshalIndent(rsp.Data.Rows, "", "  "))))

	logger.Info("[GetDeveloperDetail] get developer detail success", zap.Any("content", content))
	response.Content = []mcp.Content{content}
	response.IsError = false
	return response, nil
}

// GetLanguageStats 获取语言统计
func (h *OssinsightMCPHandler) GetLanguageStats(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	logger := logger.ContextLogger(ctx)

	logger.Info("[GetLanguageStats] receive request", zap.Any("params", request.Params))
	response := &mcp.CallToolResult{Result: mcp.Result{
		Meta: map[string]any{
			constant.MetaKeyTraceID: ctx.Value(constant.MetaKeyTraceID),
		},
	}}

	timeRange, err := request.RequireString("time_range")
	if err != nil {
		logger.Error("[GetLanguageStats] time_range is required", zap.Error(err))
		response.IsError = true
		return response, constant.ErrRequiredArgumentNotFound.Errorf("time_range")
	}

	limit := request.GetInt("limit", 10)

	req := &ossinsight.LanguageStatsQuery{
		TimeRange: timeRange,
	}

	rsp, _, err := h.cli.GetLanguageStats(ctx, req, limit)
	if err != nil {
		logger.Error("[GetLanguageStats] get language stats failed", zap.Error(err))
		response.IsError = true
		return response, err
	}

	content := mcp.NewTextContent(string(lo.Must1(json.MarshalIndent(rsp.Data.Rows, "", "  "))))

	logger.Info("[GetLanguageStats] get language stats success", zap.Any("content", content))
	response.Content = []mcp.Content{content}
	response.IsError = false
	return response, nil
}
