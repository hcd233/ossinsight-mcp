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

// RepoRankingResponse 仓库排名响应
type RepoRankingResponse struct {
	Response
	Data RepoRankingData `json:"data"`
}

// DeveloperRankingResponse 开发者排名响应
type DeveloperRankingResponse struct {
	Response
	Data DeveloperRankingData `json:"data"`
}

// CollectionsResponse 收藏夹列表响应
type CollectionsResponse struct {
	Response
	Data CollectionsData `json:"data"`
}

// HotCollectionsResponse 热门收藏夹响应
type HotCollectionsResponse struct {
	Response
	Data HotCollectionsData `json:"data"`
}

// CollectionReposResponse 收藏夹仓库列表响应
type CollectionReposResponse struct {
	Response
	Data CollectionReposData `json:"data"`
}

// RepoDetailResponse 仓库详情响应
type RepoDetailResponse struct {
	Response
	Data RepoDetailData `json:"data"`
}

// DeveloperDetailResponse 开发者详情响应
type DeveloperDetailResponse struct {
	Response
	Data DeveloperDetailData `json:"data"`
}

// LanguageStatsResponse 语言统计响应
type LanguageStatsResponse struct {
	Response
	Data LanguageStatsData `json:"data"`
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
	RepoID            string `json:"repo_id"`
	RepoName          string `json:"repo_name"`
	PrimaryLanguage   string `json:"primary_language"`
	Description       string `json:"description"`
	Stars             string `json:"stars"`
	Forks             string `json:"forks"`
	PullRequests      string `json:"pull_requests"`
	Pushes            string `json:"pushes"`
	TotalScore        string `json:"total_score"`
	ContributorLogins string `json:"contributor_logins"`
	CollectionNames   string `json:"collection_names"`
}

// RepoRankingData 仓库排名数据
type RepoRankingData struct {
	Data
	Rows []RepoRankingRow `json:"rows"`
}

// RepoRankingRow 仓库排名数据行
type RepoRankingRow struct {
	Rank            string `json:"rank"`
	RepoID          string `json:"repo_id"`
	RepoName        string `json:"repo_name"`
	PrimaryLanguage string `json:"primary_language"`
	Description     string `json:"description"`
	Stars           string `json:"stars"`
	Forks           string `json:"forks"`
	Commits         string `json:"commits"`
	Contributors    string `json:"contributors"`
	PullRequests    string `json:"pull_requests"`
	Issues          string `json:"issues"`
}

// DeveloperRankingData 开发者排名数据
type DeveloperRankingData struct {
	Data
	Rows []DeveloperRankingRow `json:"rows"`
}

// DeveloperRankingRow 开发者排名数据行
type DeveloperRankingRow struct {
	Rank            string `json:"rank"`
	UserID          string `json:"user_id"`
	Login           string `json:"login"`
	Name            string `json:"name"`
	Company         string `json:"company"`
	Location        string `json:"location"`
	Stars           string `json:"stars"`
	Forks           string `json:"forks"`
	Commits         string `json:"commits"`
	PullRequests    string `json:"pull_requests"`
	Issues          string `json:"issues"`
	Followers       string `json:"followers"`
	Following       string `json:"following"`
}

// CollectionsData 收藏夹数据
type CollectionsData struct {
	Data
	Rows []CollectionRow `json:"rows"`
}

// CollectionRow 收藏夹数据行
type CollectionRow struct {
	CollectionID   string `json:"collection_id"`
	CollectionName string `json:"collection_name"`
	Description    string `json:"description"`
	CreatorLogin   string `json:"creator_login"`
	CreatorName    string `json:"creator_name"`
	RepoCount      string `json:"repo_count"`
	FollowerCount  string `json:"follower_count"`
	CreatedAt      string `json:"created_at"`
	UpdatedAt      string `json:"updated_at"`
}

// HotCollectionsData 热门收藏夹数据
type HotCollectionsData struct {
	Data
	Rows []HotCollectionRow `json:"rows"`
}

// HotCollectionRow 热门收藏夹数据行
type HotCollectionRow struct {
	CollectionID   string `json:"collection_id"`
	CollectionName string `json:"collection_name"`
	Description    string `json:"description"`
	CreatorLogin   string `json:"creator_login"`
	CreatorName    string `json:"creator_name"`
	RepoCount      string `json:"repo_count"`
	FollowerCount  string `json:"follower_count"`
	HotScore       string `json:"hot_score"`
	CreatedAt      string `json:"created_at"`
	UpdatedAt      string `json:"updated_at"`
}

// CollectionReposData 收藏夹仓库数据
type CollectionReposData struct {
	Data
	Rows []CollectionRepoRow `json:"rows"`
}

// CollectionRepoRow 收藏夹仓库数据行
type CollectionRepoRow struct {
	RepoID          string `json:"repo_id"`
	RepoName        string `json:"repo_name"`
	PrimaryLanguage string `json:"primary_language"`
	Description     string `json:"description"`
	Stars           string `json:"stars"`
	Forks           string `json:"forks"`
	Commits         string `json:"commits"`
	Contributors    string `json:"contributors"`
	PullRequests    string `json:"pull_requests"`
	Issues          string `json:"issues"`
	AddedAt         string `json:"added_at"`
}

// RepoDetailData 仓库详情数据
type RepoDetailData struct {
	Data
	Rows []RepoDetailRow `json:"rows"`
}

// RepoDetailRow 仓库详情数据行
type RepoDetailRow struct {
	RepoID          string `json:"repo_id"`
	RepoName        string `json:"repo_name"`
	PrimaryLanguage string `json:"primary_language"`
	Description     string `json:"description"`
	Stars           string `json:"stars"`
	Forks           string `json:"forks"`
	Commits         string `json:"commits"`
	Contributors    string `json:"contributors"`
	PullRequests    string `json:"pull_requests"`
	Issues          string `json:"issues"`
	CreatedAt       string `json:"created_at"`
	UpdatedAt       string `json:"updated_at"`
	PushedAt        string `json:"pushed_at"`
	Size            string `json:"size"`
	License         string `json:"license"`
	Topics          string `json:"topics"`
	Homepage        string `json:"homepage"`
	Language        string `json:"language"`
}

// DeveloperDetailData 开发者详情数据
type DeveloperDetailData struct {
	Data
	Rows []DeveloperDetailRow `json:"rows"`
}

// DeveloperDetailRow 开发者详情数据行
type DeveloperDetailRow struct {
	UserID        string `json:"user_id"`
	Login         string `json:"login"`
	Name          string `json:"name"`
	Company       string `json:"company"`
	Location      string `json:"location"`
	Email         string `json:"email"`
	Bio           string `json:"bio"`
	Blog          string `json:"blog"`
	Twitter       string `json:"twitter_username"`
	Stars         string `json:"stars"`
	Forks         string `json:"forks"`
	Commits       string `json:"commits"`
	PullRequests  string `json:"pull_requests"`
	Issues        string `json:"issues"`
	Followers     string `json:"followers"`
	Following     string `json:"following"`
	CreatedAt     string `json:"created_at"`
	UpdatedAt     string `json:"updated_at"`
	PublicRepos   string `json:"public_repos"`
	PublicGists   string `json:"public_gists"`
	Hireable      string `json:"hireable"`
}

// LanguageStatsData 语言统计数据
type LanguageStatsData struct {
	Data
	Rows []LanguageStatsRow `json:"rows"`
}

// LanguageStatsRow 语言统计数据行
type LanguageStatsRow struct {
	Language       string `json:"language"`
	RepoCount      string `json:"repo_count"`
	Stars          string `json:"stars"`
	Forks          string `json:"forks"`
	Commits        string `json:"commits"`
	Contributors   string `json:"contributors"`
	PullRequests   string `json:"pull_requests"`
	Issues         string `json:"issues"`
	GrowthRate     string `json:"growth_rate"`
	PopularityRank string `json:"popularity_rank"`
}
