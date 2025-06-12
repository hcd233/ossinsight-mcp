package mcp

import (
	"context"
	"encoding/json"

	"github.com/hcd233/ossinsight-mcp/internal/constant"
	"github.com/hcd233/ossinsight-mcp/internal/repo/ossinsight"
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/samber/lo"
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
	response := &mcp.CallToolResult{Result: mcp.Result{
		Meta: map[string]any{
			constant.MetaKeyTraceID: ctx.Value(constant.MetaKeyTraceID),
		},
	}}
	period, err := request.RequireString("period")
	if err != nil {
		response.IsError = true
		return response, constant.ErrRequiredArgumentNotFound.Errorf("period")
	}

	language, err := request.RequireString("language")
	if err != nil {
		response.IsError = true
		return response, constant.ErrRequiredArgumentNotFound.Errorf("language")
	}

	req := &ossinsight.ListTrendingReposQuery{
		Period:   period,
		Language: language,
	}

	rsp, _, err := h.cli.ListTrendingRepos(ctx, req)
	if err != nil {
		response.IsError = true
		return response, err
	}

	content := mcp.NewTextContent(string(lo.Must1(json.MarshalIndent(rsp.Data.Rows, "", "  "))))

	response.Content = []mcp.Content{content}
	response.IsError = false
	return response, nil
}
