package main

import (
	"context"
	"fmt"
	"log"

	"github.com/hcd233/ossinsight-mcp/internal/repo/ossinsight"
)

func main() {
	fmt.Println("Testing OSS Insight MCP Tools...")

	// 创建客户端
	client := ossinsight.NewClient()

	ctx := context.Background()

	// 测试趋势仓库工具
	fmt.Println("\n=== Testing ListTrendingRepos ===")
	trendingQuery := &ossinsight.ListTrendingReposQuery{
		Period:   ossinsight.ListTrendReposPeriodPast24Hours,
		Language: ossinsight.LanguageGo,
	}
	trendingResp, _, err := client.ListTrendingRepos(ctx, trendingQuery, 5)
	if err != nil {
		log.Printf("ListTrendingRepos failed: %v", err)
	} else {
		fmt.Printf("Trending repos found: %d\n", len(trendingResp.Data.Rows))
	}

	// 测试仓库排名工具
	fmt.Println("\n=== Testing GetRepoRanking ===")
	repoRankingQuery := &ossinsight.RepoRankingQuery{
		RankingType: ossinsight.RankingTypeStars,
		TimeRange:   ossinsight.TimeRangePastMonth,
		Language:    ossinsight.LanguageGo,
	}
	repoRankingResp, _, err := client.GetRepoRanking(ctx, repoRankingQuery, 5)
	if err != nil {
		log.Printf("GetRepoRanking failed: %v", err)
	} else {
		fmt.Printf("Repo rankings found: %d\n", len(repoRankingResp.Data.Rows))
	}

	// 测试开发者排名工具
	fmt.Println("\n=== Testing GetDeveloperRanking ===")
	devRankingQuery := &ossinsight.DeveloperRankingQuery{
		RankingType: ossinsight.RankingTypeCommits,
		TimeRange:   ossinsight.TimeRangePastMonth,
	}
	devRankingResp, _, err := client.GetDeveloperRanking(ctx, devRankingQuery, 5)
	if err != nil {
		log.Printf("GetDeveloperRanking failed: %v", err)
	} else {
		fmt.Printf("Developer rankings found: %d\n", len(devRankingResp.Data.Rows))
	}

	// 测试热门收藏夹工具
	fmt.Println("\n=== Testing GetHotCollections ===")
	hotCollectionsQuery := &ossinsight.CollectionsQuery{
		CollectionType: ossinsight.CollectionTypeHot,
		TimeRange:      ossinsight.TimeRangePastMonth,
	}
	hotCollectionsResp, _, err := client.GetHotCollections(ctx, hotCollectionsQuery, 5)
	if err != nil {
		log.Printf("GetHotCollections failed: %v", err)
	} else {
		fmt.Printf("Hot collections found: %d\n", len(hotCollectionsResp.Data.Rows))
	}

	// 测试语言统计工具
	fmt.Println("\n=== Testing GetLanguageStats ===")
	languageStatsQuery := &ossinsight.LanguageStatsQuery{
		TimeRange: ossinsight.TimeRangePastMonth,
	}
	languageStatsResp, _, err := client.GetLanguageStats(ctx, languageStatsQuery, 10)
	if err != nil {
		log.Printf("GetLanguageStats failed: %v", err)
	} else {
		fmt.Printf("Language stats found: %d\n", len(languageStatsResp.Data.Rows))
	}

	fmt.Println("\n=== Testing completed ===")
}