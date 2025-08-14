package ossinsight

// ListTrendingReposQuery 获取趋势仓库列表的查询参数
//
//	@author centonhuang
//	@update 2025-06-09 20:39:02
type ListTrendingReposQuery struct {
	Period ListTrendReposPeriod `json:"period"`
	// Language 表示仓库的语言信息，使用 ListTrendReposLanguage 结构体进行序列化
	Language ListTrendReposLanguage `json:"language"`
}

// RepoRankingQuery 仓库排名查询参数
type RepoRankingQuery struct {
	RankingType RankingType `json:"ranking_type"`
	TimeRange   TimeRange   `json:"time_range"`
	Language    ListTrendReposLanguage `json:"language,omitempty"`
}

// DeveloperRankingQuery 开发者排名查询参数
type DeveloperRankingQuery struct {
	RankingType RankingType `json:"ranking_type"`
	TimeRange   TimeRange   `json:"time_range"`
	Language    ListTrendReposLanguage `json:"language,omitempty"`
}

// CollectionsQuery 收藏夹查询参数
type CollectionsQuery struct {
	CollectionType CollectionType `json:"collection_type"`
	TimeRange      TimeRange      `json:"time_range"`
}

// CollectionReposQuery 收藏夹仓库查询参数
type CollectionReposQuery struct {
	CollectionID string `json:"collection_id"`
	TimeRange    TimeRange `json:"time_range"`
}

// RepoDetailQuery 仓库详情查询参数
type RepoDetailQuery struct {
	Owner string `json:"owner"`
	Repo  string `json:"repo"`
}

// DeveloperDetailQuery 开发者详情查询参数
type DeveloperDetailQuery struct {
	Username string `json:"username"`
}

// LanguageStatsQuery 语言统计查询参数
type LanguageStatsQuery struct {
	TimeRange TimeRange `json:"time_range"`
}
