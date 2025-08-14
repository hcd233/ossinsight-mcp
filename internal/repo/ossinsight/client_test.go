package ossinsight

import (
	"context"
	"testing"
	"time"
)

func TestNewClient(t *testing.T) {
	client := NewClient()
	if client == nil {
		t.Fatal("NewClient() returned nil")
	}
	if client.baseURL != BaseURL {
		t.Errorf("Expected baseURL %s, got %s", BaseURL, client.baseURL)
	}
	if client.httpClient == nil {
		t.Fatal("HTTP client is nil")
	}
	if client.userAgent == "" {
		t.Fatal("User agent is empty")
	}
}

func TestNewClientWithOptions(t *testing.T) {
	customURL := "https://custom.api.com"
	customUserAgent := "custom-agent/1.0"
	
	client := NewClient(
		WithBaseURL(customURL),
		WithUserAgent(customUserAgent),
	)
	
	if client.baseURL != customURL {
		t.Errorf("Expected baseURL %s, got %s", customURL, client.baseURL)
	}
	if client.userAgent != customUserAgent {
		t.Errorf("Expected userAgent %s, got %s", customUserAgent, client.userAgent)
	}
}

func TestListTrendingRepos(t *testing.T) {
	client := NewClient()
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	query := &ListTrendingReposQuery{
		Period:   ListTrendReposPeriodPast24Hours,
		Language: LanguageGo,
	}

	response, rateLimit, err := client.ListTrendingRepos(ctx, query, 5)
	if err != nil {
		t.Fatalf("ListTrendingRepos failed: %v", err)
	}

	if response == nil {
		t.Fatal("Response is nil")
	}

	if response.Data.Rows == nil {
		t.Fatal("Response data rows is nil")
	}

	if len(response.Data.Rows) > 5 {
		t.Errorf("Expected max 5 rows, got %d", len(response.Data.Rows))
	}

	// Check rate limit info
	if rateLimit != nil {
		t.Logf("Rate limit info: Limit=%d, Remaining=%d", rateLimit.Limit, rateLimit.Remaining)
	}

	// Log first repo info if available
	if len(response.Data.Rows) > 0 {
		firstRepo := response.Data.Rows[0]
		t.Logf("First trending repo: %s", firstRepo.RepoName)
	}
}

func TestGetCollections(t *testing.T) {
	client := NewClient()
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	query := &CollectionsQuery{
		CollectionType: CollectionTypeHot,
		TimeRange:      TimeRangePastMonth,
	}

	response, rateLimit, err := client.GetCollections(ctx, query, 5)
	if err != nil {
		t.Fatalf("GetCollections failed: %v", err)
	}

	if response == nil {
		t.Fatal("Response is nil")
	}

	if response.Data.Rows == nil {
		t.Fatal("Response data rows is nil")
	}

	if len(response.Data.Rows) > 5 {
		t.Errorf("Expected max 5 rows, got %d", len(response.Data.Rows))
	}

	// Check rate limit info
	if rateLimit != nil {
		t.Logf("Rate limit info: Limit=%d, Remaining=%d", rateLimit.Limit, rateLimit.Remaining)
	}

	// Log first collection info if available
	if len(response.Data.Rows) > 0 {
		firstCollection := response.Data.Rows[0]
		t.Logf("First collection: %s", firstCollection.CollectionName)
	}
}

func TestGetHotCollections(t *testing.T) {
	client := NewClient()
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	query := &CollectionsQuery{
		CollectionType: CollectionTypeHot,
		TimeRange:      TimeRangePastMonth,
	}

	response, rateLimit, err := client.GetHotCollections(ctx, query, 5)
	if err != nil {
		t.Fatalf("GetHotCollections failed: %v", err)
	}

	if response == nil {
		t.Fatal("Response is nil")
	}

	if response.Data.Rows == nil {
		t.Fatal("Response data rows is nil")
	}

	if len(response.Data.Rows) > 5 {
		t.Errorf("Expected max 5 rows, got %d", len(response.Data.Rows))
	}

	// Check rate limit info
	if rateLimit != nil {
		t.Logf("Rate limit info: Limit=%d, Remaining=%d", rateLimit.Limit, rateLimit.Remaining)
	}

	// Log first hot collection info if available
	if len(response.Data.Rows) > 0 {
		firstCollection := response.Data.Rows[0]
		t.Logf("First hot collection: %s (Hot Score: %s)", firstCollection.CollectionName, firstCollection.HotScore)
	}
}

func TestToQuery(t *testing.T) {
	query := &ListTrendingReposQuery{
		Period:   ListTrendReposPeriodPast24Hours,
		Language: LanguageGo,
	}

	values := toQuery(query)
	
	if values.Get("period") != string(ListTrendReposPeriodPast24Hours) {
		t.Errorf("Expected period %s, got %s", ListTrendReposPeriodPast24Hours, values.Get("period"))
	}
	
	if values.Get("language") != string(LanguageGo) {
		t.Errorf("Expected language %s, got %s", LanguageGo, values.Get("language"))
	}
}

func TestConstants(t *testing.T) {
	// Test ranking types
	expectedRankingTypes := []string{
		RankingTypeStars,
		RankingTypeForks,
		RankingTypeCommits,
		RankingTypeContributors,
		RankingTypePullRequests,
		RankingTypeIssues,
	}
	
	for _, rankingType := range expectedRankingTypes {
		if rankingType == "" {
			t.Errorf("Ranking type is empty")
		}
	}

	// Test time ranges
	expectedTimeRanges := []string{
		TimeRangePast24Hours,
		TimeRangePastWeek,
		TimeRangePastMonth,
		TimeRangePast3Months,
		TimeRangePastYear,
		TimeRangeAllTime,
	}
	
	for _, timeRange := range expectedTimeRanges {
		if timeRange == "" {
			t.Errorf("Time range is empty")
		}
	}

	// Test collection types
	expectedCollectionTypes := []string{
		CollectionTypeAll,
		CollectionTypeHot,
		CollectionTypeUser,
	}
	
	for _, collectionType := range expectedCollectionTypes {
		if collectionType == "" {
			t.Errorf("Collection type is empty")
		}
	}
}

// Test API endpoints that may not be available yet
func TestGetRepoRanking_Optional(t *testing.T) {
	client := NewClient()
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	query := &RepoRankingQuery{
		RankingType: RankingTypeStars,
		TimeRange:   TimeRangePastMonth,
		Language:    LanguageGo,
	}

	response, rateLimit, err := client.GetRepoRanking(ctx, query, 5)
	if err != nil {
		t.Logf("GetRepoRanking failed (this may be expected): %v", err)
		return
	}

	if response == nil {
		t.Fatal("Response is nil")
	}

	if response.Data.Rows == nil {
		t.Fatal("Response data rows is nil")
	}

	// Check rate limit info
	if rateLimit != nil {
		t.Logf("Rate limit info: Limit=%d, Remaining=%d", rateLimit.Limit, rateLimit.Remaining)
	}

	// Log first repo info if available
	if len(response.Data.Rows) > 0 {
		firstRepo := response.Data.Rows[0]
		t.Logf("First ranked repo: %s (Rank: %s)", firstRepo.RepoName, firstRepo.Rank)
	}
}

func TestGetDeveloperRanking_Optional(t *testing.T) {
	client := NewClient()
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	query := &DeveloperRankingQuery{
		RankingType: RankingTypeCommits,
		TimeRange:   TimeRangePastMonth,
	}

	response, rateLimit, err := client.GetDeveloperRanking(ctx, query, 5)
	if err != nil {
		t.Logf("GetDeveloperRanking failed (this may be expected): %v", err)
		return
	}

	if response == nil {
		t.Fatal("Response is nil")
	}

	if response.Data.Rows == nil {
		t.Fatal("Response data rows is nil")
	}

	// Check rate limit info
	if rateLimit != nil {
		t.Logf("Rate limit info: Limit=%d, Remaining=%d", rateLimit.Limit, rateLimit.Remaining)
	}

	// Log first developer info if available
	if len(response.Data.Rows) > 0 {
		firstDev := response.Data.Rows[0]
		t.Logf("First ranked developer: %s (Rank: %s)", firstDev.Login, firstDev.Rank)
	}
}

func TestGetLanguageStats_Optional(t *testing.T) {
	client := NewClient()
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	query := &LanguageStatsQuery{
		TimeRange: TimeRangePastMonth,
	}

	response, rateLimit, err := client.GetLanguageStats(ctx, query, 10)
	if err != nil {
		t.Logf("GetLanguageStats failed (this may be expected): %v", err)
		return
	}

	if response == nil {
		t.Fatal("Response is nil")
	}

	if response.Data.Rows == nil {
		t.Fatal("Response data rows is nil")
	}

	// Check rate limit info
	if rateLimit != nil {
		t.Logf("Rate limit info: Limit=%d, Remaining=%d", rateLimit.Limit, rateLimit.Remaining)
	}

	// Log first language info if available
	if len(response.Data.Rows) > 0 {
		firstLang := response.Data.Rows[0]
		t.Logf("First language: %s (Repo Count: %s)", firstLang.Language, firstLang.RepoCount)
	}
}

func TestGetRepoDetail_Optional(t *testing.T) {
	client := NewClient()
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	query := &RepoDetailQuery{
		Owner: "golang",
		Repo:  "go",
	}

	response, rateLimit, err := client.GetRepoDetail(ctx, query)
	if err != nil {
		t.Logf("GetRepoDetail failed (this may be expected): %v", err)
		return
	}

	if response == nil {
		t.Fatal("Response is nil")
	}

	if response.Data.Rows == nil {
		t.Fatal("Response data rows is nil")
	}

	// Check rate limit info
	if rateLimit != nil {
		t.Logf("Rate limit info: Limit=%d, Remaining=%d", rateLimit.Limit, rateLimit.Remaining)
	}

	// Log repo info if available
	if len(response.Data.Rows) > 0 {
		repo := response.Data.Rows[0]
		t.Logf("Repo detail: %s (Stars: %s, Forks: %s)", repo.RepoName, repo.Stars, repo.Forks)
	}
}

func TestGetDeveloperDetail_Optional(t *testing.T) {
	client := NewClient()
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	query := &DeveloperDetailQuery{
		Username: "torvalds",
	}

	response, rateLimit, err := client.GetDeveloperDetail(ctx, query)
	if err != nil {
		t.Logf("GetDeveloperDetail failed (this may be expected): %v", err)
		return
	}

	if response == nil {
		t.Fatal("Response is nil")
	}

	if response.Data.Rows == nil {
		t.Fatal("Response data rows is nil")
	}

	// Check rate limit info
	if rateLimit != nil {
		t.Logf("Rate limit info: Limit=%d, Remaining=%d", rateLimit.Limit, rateLimit.Remaining)
	}

	// Log developer info if available
	if len(response.Data.Rows) > 0 {
		dev := response.Data.Rows[0]
		t.Logf("Developer detail: %s (Followers: %s, Following: %s)", dev.Login, dev.Followers, dev.Following)
	}
}