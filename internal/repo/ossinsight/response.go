package ossinsight

// Column 列定义
//
//	@author centonhuang
//	@update 2025-06-09 20:14:56
type Column struct {
	Col      string `json:"col"`
	DataType string `json:"data_type"`
	Nullable bool   `json:"nullable"`
}

// Result 查询结果元数据
//
//	@author centonhuang
//	@update 2025-06-09 20:14:58
type Result struct {
	Code      int64    `json:"code"`
	Message   string   `json:"message"`
	StartMs   int64    `json:"start_ms"`
	EndMs     int64    `json:"end_ms"`
	Latency   string   `json:"latency"`
	RowCount  int64    `json:"row_count"`
	RowAffect int64    `json:"row_affect"`
	Limit     int64    `json:"limit"`
	Databases []string `json:"databases"`
}

// Response 响应结构
//
//	@author centonhuang
//	@update 2025-06-09 20:27:11
type Response struct {
	Type string `json:"type"`
	Data any    `json:"data"`
}

// ListTrendingReposResponse 趋势仓库列表响应
//
//	@author centonhuang
//	@update 2025-06-09 20:27:49
type ListTrendingReposResponse struct {
	Response
	Data TrendingReposData `json:"data"`
}

// // ListCollectionsResponse 收藏夹列表响应
// //
// //	@author centonhuang
// //	@update 2025-06-09 20:27:51
// type ListCollectionsResponse struct {
// 	Response
// 	Data ListCollectionsData `json:"data"`
// }

// // ListHotCollectionsResponse 热门收藏夹列表响应
// //
// //	@author centonhuang
// //	@update 2025-06-09 20:27:53
// type ListHotCollectionsResponse struct {
// 	Response
// 	Data ListHotCollectionsData `json:"data"`
// }

// // ListCollectionReposResponse 收藏夹仓库列表响应

// // @author centonhuang
// // @update 2025-06-09 20:27:55
// type ListCollectionReposResponse struct {
// 	Response
// 	Data ListCollectionReposData `json:"data"`
// }

// // RepoRankingResponse 仓库排名响应
// //
// //	@author centonhuang
// //	@update 2025-06-09 20:27:57
// type RepoRankingResponse struct {
// 	Response
// 	Data RepoRankingData `json:"data"`
// }

// RateLimitInfo contains rate limit information from response headers
type RateLimitInfo struct {
	Limit           int64 `json:"limit"`
	Remaining       int64 `json:"remaining"`
	LimitMinute     int64 `json:"limit_minute"`
	RemainingMinute int64 `json:"remaining_minute"`
}

// Data 基本数据字段
//
//	@author centonhuang
//	@update 2025-06-09 20:14:49
type Data struct {
	Columns []Column `json:"columns"`
	Result  Result   `json:"result"`
	Rows    []any    `json:"rows"`
}

// TrendingReposData 趋势仓库数据
//
//	@author centonhuang
//	@update 2025-06-09 20:14:53
type TrendingReposData struct {
	Data
	Rows []TrendingRepoRow `json:"rows"`
}

// // ListCollectionsData 列表收藏夹数据
// //
// //	@author centonhuang
// //	@update 2025-06-09 20:16:48
// type ListCollectionsData struct {
// 	Data
// 	Rows []CollectionRow `json:"rows"`
// }

// // ListHotCollectionsData 热门收藏夹数据
// //
// //	@author centonhuang
// //	@update 2025-06-09 20:16:50
// type ListHotCollectionsData struct {
// 	Data
// 	Rows []HotCollectionRow `json:"rows"`
// }

// // ListCollectionReposData 收藏夹仓库列表数据
// //
// //	@author centonhuang
// //	@update 2025-06-09 20:16:52
// type ListCollectionReposData struct {
// 	Data
// 	Rows []CollectionRepoRow `json:"rows"`
// }

// // RepoRankingByIssuesData 仓库Issues排名数据
// //
// //	@author centonhuang
// //	@update 2025-06-09 20:22:00
// type RepoRankingData struct {
// 	Data
// 	Rows []RepoRankingRow `json:"rows"`
// }

// TrendingRepoRow 趋势仓库数据
//
//	@author centonhuang
//	@update 2025-06-09 20:24:00
type TrendingRepoRow struct {
	RepoID            int64   `json:"repo_id"`
	RepoName          string  `json:"repo_name"`
	PrimaryLanguage   string  `json:"primary_language"`
	Description       string  `json:"description"`
	Stars             int64   `json:"stars"`
	Forks             int64   `json:"forks"`
	PullRequests      int64   `json:"pull_requests"`
	Pushes            int64   `json:"pushes"`
	TotalScore        float64 `json:"total_score"`
	ContributorLogins string  `json:"contributor_logins"`
	CollectionNames   string  `json:"collection_names"`
}

// // CollectionRow 收藏夹数据
// //
// //	@author centonhuang
// //	@update 2025-06-09 20:24:03
// type CollectionRow struct {
// 	ID   int64  `json:"id"`
// 	Name string `json:"name"`
// }

// // HotCollectionRow 热门收藏夹数据
// //
// //	@author centonhuang
// //	@update 2025-06-09 20:24:09
// type HotCollectionRow struct {
// 	ID                    int64  `json:"id"`
// 	Name                  string `json:"name"`
// 	Repos                 int64  `json:"repos"`
// 	RepoID                int64  `json:"repo_id"`
// 	RepoName              string `json:"repo_name"`
// 	RepoCurrentPeriodRank *int   `json:"repo_current_period_rank"`
// 	RepoPastPeriodRank    *int   `json:"repo_past_period_rank"`
// 	RepoRankChanges       *int64 `json:"repo_rank_changes"`
// }

// // CollectionRepoRow 收藏夹仓库列表数据
// //
// //	@author centonhuang
// //	@update 2025-06-09 20:24:11
// type CollectionRepoRow struct {
// 	RepoID   int64  `json:"repo_id"`
// 	RepoName string `json:"repo_name"`
// }

// // RepoRankingRow 仓库排名数据
// //
// //	@author centonhuang
// //	@update 2025-06-09 20:24:13
// type RepoRankingRow struct {
// 	RepoID              int64   `json:"repo_id"`
// 	RepoName            string  `json:"repo_name"`
// 	CurrentPeriodGrowth int64   `json:"current_period_growth"`
// 	CurrentPeriodRank   int64   `json:"current_period_rank"`
// 	PastPeriodGrowth    int64   `json:"past_period_growth"`
// 	PastPeriodRank      int64   `json:"past_period_rank"`
// 	GrowthPop           float64 `json:"growth_pop"`
// 	RankPop             int64   `json:"rank_pop"`
// 	Total               int64   `json:"total"`
// }
