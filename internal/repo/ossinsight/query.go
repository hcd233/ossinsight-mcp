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
