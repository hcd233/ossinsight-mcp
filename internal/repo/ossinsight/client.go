package ossinsight

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

// Client represents the OSSInsight API client
type Client struct {
	baseURL    string
	httpClient *http.Client
	userAgent  string
}

// ClientOption represents a client configuration option
type ClientOption func(*Client)

// WithBaseURL sets a custom base URL for the client
func WithBaseURL(baseURL string) ClientOption {
	return func(c *Client) {
		c.baseURL = baseURL
	}
}

// WithHTTPClient sets a custom HTTP client
func WithHTTPClient(httpClient *http.Client) ClientOption {
	return func(c *Client) {
		c.httpClient = httpClient
	}
}

// WithUserAgent sets a custom user agent
func WithUserAgent(userAgent string) ClientOption {
	return func(c *Client) {
		c.userAgent = userAgent
	}
}

// NewClient creates a new OSSInsight API client
func NewClient(options ...ClientOption) *Client {
	client := &Client{
		baseURL: BaseURL,
		httpClient: &http.Client{
			Timeout: 30 * time.Second,
		},
		userAgent: "ossinsight-go-sdk/1.0",
	}

	for _, option := range options {
		option(client)
	}

	return client
}

// makeRequest makes an HTTP request to the OSSInsight API
func (c *Client) makeRequest(ctx context.Context, endpoint string, query url.Values) ([]byte, *RateLimitInfo, error) {
	u, err := url.JoinPath(c.baseURL, endpoint)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to join path: %w", err)
	}

	// 添加查询参数
	if query != nil {
		u = u + "?" + query.Encode()
	}

	req, err := http.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("User-Agent", c.userAgent)

	req = req.WithContext(ctx)

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to make request: %w", err)
	}
	defer resp.Body.Close()

	// Parse rate limit information from headers
	rateLimitInfo := &RateLimitInfo{}
	if limit := resp.Header.Get(headerRateLimitLimit); limit != "" {
		if val, err := strconv.ParseInt(limit, 10, 64); err == nil {
			rateLimitInfo.Limit = val
		}
	}
	if remaining := resp.Header.Get(headerRateLimitRemaining); remaining != "" {
		if val, err := strconv.ParseInt(remaining, 10, 64); err == nil {
			rateLimitInfo.Remaining = val
		}
	}
	if limitMinute := resp.Header.Get(headerRateLimitLimitMinute); limitMinute != "" {
		if val, err := strconv.ParseInt(limitMinute, 10, 64); err == nil {
			rateLimitInfo.LimitMinute = val
		}
	}
	if remainingMinute := resp.Header.Get(headerRateLimitRemainingMinute); remainingMinute != "" {
		if val, err := strconv.ParseInt(remainingMinute, 10, 64); err == nil {
			rateLimitInfo.RemainingMinute = val
		}
	}

	if resp.StatusCode != http.StatusOK {
		return nil, rateLimitInfo, fmt.Errorf("API request failed with status %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, rateLimitInfo, fmt.Errorf("failed to read response body: %w", err)
	}

	return body, rateLimitInfo, nil
}

// ListTrendingRepos 获取趋势仓库列表
//
//	@receiver c *Client
//	@param owner
//	@param repo
//	@return *ListTrendingReposResponse
//	@return *RateLimitInfo
//	@return error
//	@author centonhuang
//	@update 2025-06-09 20:34:05
func (c *Client) ListTrendingRepos(ctx context.Context, query *ListTrendingReposQuery) (*ListTrendingReposResponse, *RateLimitInfo, error) {
	bts, rateLimitInfo, err := c.makeRequest(ctx, listTrendingReposEndpoint, toQuery(query))
	if err != nil {
		return nil, rateLimitInfo, err
	}

	var response ListTrendingReposResponse
	if err := json.Unmarshal(bts, &response); err != nil {
		return nil, rateLimitInfo, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	return &response, rateLimitInfo, nil
}
