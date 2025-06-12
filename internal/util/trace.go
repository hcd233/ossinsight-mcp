package util

import (
	"context"
	"net/http"

	"github.com/google/uuid"
	"github.com/hcd233/ossinsight-mcp/internal/constant"
)

// WithTraceID 设置请求追踪ID
//
//	@param ctx
//	@return context.Context
//	@author centonhuang
//	@update 2025-06-12 21:14:43
func WithTraceID(ctx context.Context, r *http.Request) context.Context {
	traceID := r.Header.Get("X-Trace-ID")
	if traceID == "" {
		traceID = uuid.New().String()
	}
	return context.WithValue(ctx, constant.CtxKeyTraceID, traceID)
}
