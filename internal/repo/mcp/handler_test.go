package mcp

import (
	"context"
	"testing"

	"github.com/hcd233/ossinsight-mcp/internal/constant"
	"github.com/hcd233/ossinsight-mcp/internal/repo/ossinsight"
	"github.com/mark3labs/mcp-go/mcp"
)

func TestNewOssinsightMCPHandler(t *testing.T) {
	client := ossinsight.NewClient()
	handler := NewOssinsightMCPHandler(client)
	
	if handler == nil {
		t.Fatal("NewOssinsightMCPHandler returned nil")
	}
	
	if handler.cli == nil {
		t.Fatal("Client is nil")
	}
}

func createTestContext() context.Context {
	ctx := context.Background()
	ctx = context.WithValue(ctx, constant.MetaKeyTraceID, "test-trace-id")
	return ctx
}

func TestListTrendingRepos(t *testing.T) {
	handler := NewOssinsightMCPHandler(ossinsight.NewClient())
	ctx := createTestContext()

	request := mcp.CallToolRequest{
		Params: mcp.CallToolParams{
			Arguments: map[string]any{
				"period":   ossinsight.ListTrendReposPeriodPast24Hours,
				"language": ossinsight.LanguageGo,
				"limit":    5,
			},
		},
	}

	result, err := handler.ListTrendingRepos(ctx, request)
	if err != nil {
		t.Fatalf("ListTrendingRepos failed: %v", err)
	}

	if result == nil {
		t.Fatal("Result is nil")
	}

	if result.IsError {
		t.Fatal("Result indicates error")
	}

	if len(result.Content) == 0 {
		t.Fatal("No content returned")
	}

	// Check meta information
	if result.Result.Meta == nil {
		t.Fatal("Meta information is nil")
	}

	if result.Result.Meta[constant.MetaKeyTraceID] != "test-trace-id" {
		t.Errorf("Expected trace ID 'test-trace-id', got %v", result.Result.Meta[constant.MetaKeyTraceID])
	}
}

func TestListTrendingReposMissingRequiredParams(t *testing.T) {
	handler := NewOssinsightMCPHandler(ossinsight.NewClient())
	ctx := createTestContext()

	// Test missing period
	request := mcp.CallToolRequest{
		Params: mcp.CallToolParams{
			Arguments: map[string]any{
				"language": ossinsight.LanguageGo,
				"limit":    5,
			},
		},
	}

	result, err := handler.ListTrendingRepos(ctx, request)
	if err == nil {
		t.Fatal("Expected error for missing period parameter")
	}

	if !result.IsError {
		t.Fatal("Result should indicate error")
	}

	// Test missing language
	request = mcp.CallToolRequest{
		Params: mcp.CallToolParams{
			Arguments: map[string]any{
				"period": ossinsight.ListTrendReposPeriodPast24Hours,
				"limit":  5,
			},
		},
	}

	result, err = handler.ListTrendingRepos(ctx, request)
	if err == nil {
		t.Fatal("Expected error for missing language parameter")
	}

	if !result.IsError {
		t.Fatal("Result should indicate error")
	}
}

func TestGetCollections(t *testing.T) {
	handler := NewOssinsightMCPHandler(ossinsight.NewClient())
	ctx := createTestContext()

	request := mcp.CallToolRequest{
		Params: mcp.CallToolParams{
			Arguments: map[string]any{
				"collection_type": ossinsight.CollectionTypeHot,
				"time_range":      ossinsight.TimeRangePastMonth,
				"limit":           5,
			},
		},
	}

	result, err := handler.GetCollections(ctx, request)
	if err != nil {
		t.Fatalf("GetCollections failed: %v", err)
	}

	if result == nil {
		t.Fatal("Result is nil")
	}

	if result.IsError {
		t.Fatal("Result indicates error")
	}

	if len(result.Content) == 0 {
		t.Fatal("No content returned")
	}
}

func TestGetHotCollections(t *testing.T) {
	handler := NewOssinsightMCPHandler(ossinsight.NewClient())
	ctx := createTestContext()

	request := mcp.CallToolRequest{
		Params: mcp.CallToolParams{
			Arguments: map[string]any{
				"time_range": ossinsight.TimeRangePastMonth,
				"limit":      5,
			},
		},
	}

	result, err := handler.GetHotCollections(ctx, request)
	if err != nil {
		t.Fatalf("GetHotCollections failed: %v", err)
	}

	if result == nil {
		t.Fatal("Result is nil")
	}

	if result.IsError {
		t.Fatal("Result indicates error")
	}

	if len(result.Content) == 0 {
		t.Fatal("No content returned")
	}
}

func TestGetLanguageStatsMissingRequiredParams(t *testing.T) {
	handler := NewOssinsightMCPHandler(ossinsight.NewClient())
	ctx := createTestContext()

	// Test missing time_range
	request := mcp.CallToolRequest{
		Params: mcp.CallToolParams{
			Arguments: map[string]any{
				"limit": 10,
			},
		},
	}

	result, err := handler.GetLanguageStats(ctx, request)
	if err == nil {
		t.Fatal("Expected error for missing time_range parameter")
	}

	if !result.IsError {
		t.Fatal("Result should indicate error")
	}
}

func TestGetRepoDetailMissingRequiredParams(t *testing.T) {
	handler := NewOssinsightMCPHandler(ossinsight.NewClient())
	ctx := createTestContext()

	// Test missing owner
	request := mcp.CallToolRequest{
		Params: mcp.CallToolParams{
			Arguments: map[string]any{
				"repo": "go",
			},
		},
	}

	result, err := handler.GetRepoDetail(ctx, request)
	if err == nil {
		t.Fatal("Expected error for missing owner parameter")
	}

	if !result.IsError {
		t.Fatal("Result should indicate error")
	}

	// Test missing repo
	request = mcp.CallToolRequest{
		Params: mcp.CallToolParams{
			Arguments: map[string]any{
				"owner": "golang",
			},
		},
	}

	result, err = handler.GetRepoDetail(ctx, request)
	if err == nil {
		t.Fatal("Expected error for missing repo parameter")
	}

	if !result.IsError {
		t.Fatal("Result should indicate error")
	}
}

func TestGetDeveloperDetailMissingRequiredParams(t *testing.T) {
	handler := NewOssinsightMCPHandler(ossinsight.NewClient())
	ctx := createTestContext()

	// Test missing username
	request := mcp.CallToolRequest{
		Params: mcp.CallToolParams{
			Arguments: map[string]any{},
		},
	}

	result, err := handler.GetDeveloperDetail(ctx, request)
	if err == nil {
		t.Fatal("Expected error for missing username parameter")
	}

	if !result.IsError {
		t.Fatal("Result should indicate error")
	}
}

func TestGetCollectionReposMissingRequiredParams(t *testing.T) {
	handler := NewOssinsightMCPHandler(ossinsight.NewClient())
	ctx := createTestContext()

	// Test missing collection_id
	request := mcp.CallToolRequest{
		Params: mcp.CallToolParams{
			Arguments: map[string]any{
				"time_range": ossinsight.TimeRangePastMonth,
				"limit":      5,
			},
		},
	}

	result, err := handler.GetCollectionRepos(ctx, request)
	if err == nil {
		t.Fatal("Expected error for missing collection_id parameter")
	}

	if !result.IsError {
		t.Fatal("Result should indicate error")
	}

	// Test missing time_range
	request = mcp.CallToolRequest{
		Params: mcp.CallToolParams{
			Arguments: map[string]any{
				"collection_id": "test-collection-id",
				"limit":         5,
			},
		},
	}

	result, err = handler.GetCollectionRepos(ctx, request)
	if err == nil {
		t.Fatal("Expected error for missing time_range parameter")
	}

	if !result.IsError {
		t.Fatal("Result should indicate error")
	}
}

// Optional tests for API endpoints that may not be available yet
func TestGetRepoRanking_Optional(t *testing.T) {
	handler := NewOssinsightMCPHandler(ossinsight.NewClient())
	ctx := createTestContext()

	request := mcp.CallToolRequest{
		Params: mcp.CallToolParams{
			Arguments: map[string]any{
				"ranking_type": ossinsight.RankingTypeStars,
				"time_range":   ossinsight.TimeRangePastMonth,
				"language":     ossinsight.LanguageGo,
				"limit":        5,
			},
		},
	}

	result, err := handler.GetRepoRanking(ctx, request)
	if err != nil {
		t.Logf("GetRepoRanking failed (this may be expected): %v", err)
		return
	}

	if result == nil {
		t.Fatal("Result is nil")
	}

	if result.IsError {
		t.Fatal("Result indicates error")
	}

	if len(result.Content) == 0 {
		t.Fatal("No content returned")
	}
}

func TestGetDeveloperRanking_Optional(t *testing.T) {
	handler := NewOssinsightMCPHandler(ossinsight.NewClient())
	ctx := createTestContext()

	request := mcp.CallToolRequest{
		Params: mcp.CallToolParams{
			Arguments: map[string]any{
				"ranking_type": ossinsight.RankingTypeCommits,
				"time_range":   ossinsight.TimeRangePastMonth,
				"limit":        5,
			},
		},
	}

	result, err := handler.GetDeveloperRanking(ctx, request)
	if err != nil {
		t.Logf("GetDeveloperRanking failed (this may be expected): %v", err)
		return
	}

	if result == nil {
		t.Fatal("Result is nil")
	}

	if result.IsError {
		t.Fatal("Result indicates error")
	}

	if len(result.Content) == 0 {
		t.Fatal("No content returned")
	}
}

func TestGetCollectionRepos_Optional(t *testing.T) {
	handler := NewOssinsightMCPHandler(ossinsight.NewClient())
	ctx := createTestContext()

	request := mcp.CallToolRequest{
		Params: mcp.CallToolParams{
			Arguments: map[string]any{
				"collection_id": "test-collection-id",
				"time_range":    ossinsight.TimeRangePastMonth,
				"limit":         5,
			},
		},
	}

	result, err := handler.GetCollectionRepos(ctx, request)
	if err != nil {
		t.Logf("GetCollectionRepos failed (this may be expected): %v", err)
		return
	}

	if result == nil {
		t.Fatal("Result is nil")
	}

	if result.IsError {
		t.Fatal("Result indicates error")
	}

	if len(result.Content) == 0 {
		t.Fatal("No content returned")
	}
}

func TestGetRepoDetail_Optional(t *testing.T) {
	handler := NewOssinsightMCPHandler(ossinsight.NewClient())
	ctx := createTestContext()

	request := mcp.CallToolRequest{
		Params: mcp.CallToolParams{
			Arguments: map[string]any{
				"owner": "golang",
				"repo":  "go",
			},
		},
	}

	result, err := handler.GetRepoDetail(ctx, request)
	if err != nil {
		t.Logf("GetRepoDetail failed (this may be expected): %v", err)
		return
	}

	if result == nil {
		t.Fatal("Result is nil")
	}

	if result.IsError {
		t.Fatal("Result indicates error")
	}

	if len(result.Content) == 0 {
		t.Fatal("No content returned")
	}
}

func TestGetDeveloperDetail_Optional(t *testing.T) {
	handler := NewOssinsightMCPHandler(ossinsight.NewClient())
	ctx := createTestContext()

	request := mcp.CallToolRequest{
		Params: mcp.CallToolParams{
			Arguments: map[string]any{
				"username": "torvalds",
			},
		},
	}

	result, err := handler.GetDeveloperDetail(ctx, request)
	if err != nil {
		t.Logf("GetDeveloperDetail failed (this may be expected): %v", err)
		return
	}

	if result == nil {
		t.Fatal("Result is nil")
	}

	if result.IsError {
		t.Fatal("Result indicates error")
	}

	if len(result.Content) == 0 {
		t.Fatal("No content returned")
	}
}

func TestGetLanguageStats_Optional(t *testing.T) {
	handler := NewOssinsightMCPHandler(ossinsight.NewClient())
	ctx := createTestContext()

	request := mcp.CallToolRequest{
		Params: mcp.CallToolParams{
			Arguments: map[string]any{
				"time_range": ossinsight.TimeRangePastMonth,
				"limit":      10,
			},
		},
	}

	result, err := handler.GetLanguageStats(ctx, request)
	if err != nil {
		t.Logf("GetLanguageStats failed (this may be expected): %v", err)
		return
	}

	if result == nil {
		t.Fatal("Result is nil")
	}

	if result.IsError {
		t.Fatal("Result indicates error")
	}

	if len(result.Content) == 0 {
		t.Fatal("No content returned")
	}
}