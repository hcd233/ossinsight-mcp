package mcp

import (
	"context"
	"net/http"
	"strings"

	"github.com/google/uuid"
	"github.com/hcd233/ossinsight-mcp/internal/constant"
)

// HTTPContextFunc 设置请求上下文
//
//	@param ctx
//	@param r
//	@return context.Context
//	@author centonhuang
//	@update 2025-06-13 17:15:07
func HTTPContextFunc(ctx context.Context, r *http.Request) context.Context {
	ctx = withTraceID(ctx, r)
	ctx = withBearerAuth(ctx, r)
	return ctx
}

// withTraceID 设置请求追踪ID
//
//	@param ctx
//	@return context.Context
//	@author centonhuang
//	@update 2025-06-12 21:14:43
func withTraceID(ctx context.Context, r *http.Request) context.Context {
	traceID := r.Header.Get("X-Trace-ID")
	if traceID == "" {
		traceID = uuid.NewString()
	}
	return context.WithValue(ctx, constant.CtxKeyTraceID, traceID)
}

// withBearerAuth 设置请求认证
//
//	@param ctx
//	@param r
//	@return context.Context
//	@author centonhuang
//	@update 2025-06-13 16:23:26
func withBearerAuth(ctx context.Context, r *http.Request) context.Context {
	apiKey := r.Header.Get("Authorization")
	if apiKey == "" {
		return ctx
	}

	apiKey = strings.TrimSpace(strings.TrimPrefix(apiKey, "Bearer "))
	return context.WithValue(ctx, constant.CtxKeyAPIKey, apiKey)
}
