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
