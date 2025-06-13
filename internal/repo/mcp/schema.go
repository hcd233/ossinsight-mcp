package mcp

import (
	"github.com/hcd233/ossinsight-mcp/internal/repo/ossinsight"
	"github.com/mark3labs/mcp-go/mcp"
)

// ListTrendingReposSchema 获取趋势仓库列表的工具
//
//	@author centonhuang
//	@update 2025-06-12 20:41:39
var ListTrendingReposSchema = mcp.NewTool(
	"ListTrendingRepos",
	mcp.WithDescription("Get trending GitHub repositories based on various time periods and programming languages. This tool helps discover popular and rapidly growing repositories in the GitHub ecosystem."),
	mcp.WithString("period",
		mcp.Description("Time period for fetching trending repositories. Choose from past 24 hours, past week, past month, or past 3 months to see repositories that gained popularity during that timeframe."),
		mcp.Enum(
			ossinsight.ListTrendReposPeriodPast24Hours,
			ossinsight.ListTrendReposPeriodPastWeek,
			ossinsight.ListTrendReposPeriodPastMonth,
			ossinsight.ListTrendReposPeriodPast3Months,
		),
	),
	mcp.WithString("language",
		mcp.Description("Programming language filter for trending repositories. Select a specific programming language to see trending repositories in that language, or choose 'All' to see repositories across all programming languages."),
		mcp.Enum(
			ossinsight.LanguageAll,
			ossinsight.LanguageJavaScript,
			ossinsight.LanguageJava,
			ossinsight.LanguagePython,
			ossinsight.LanguagePHP,
			ossinsight.LanguageCPlusPlus,
			ossinsight.LanguageCSharp,
			ossinsight.LanguageTypeScript,
			ossinsight.LanguageShell,
			ossinsight.LanguageC,
			ossinsight.LanguageRuby,
			ossinsight.LanguageRust,
			ossinsight.LanguageGo,
			ossinsight.LanguageKotlin,
			ossinsight.LanguageHCL,
			ossinsight.LanguagePowerShell,
			ossinsight.LanguageCMake,
			ossinsight.LanguageGroovy,
			ossinsight.LanguagePLpgSQL,
			ossinsight.LanguageTSQL,
			ossinsight.LanguageDart,
			ossinsight.LanguageSwift,
			ossinsight.LanguageHTML,
			ossinsight.LanguageCSS,
			ossinsight.LanguageElixir,
			ossinsight.LanguageHaskell,
			ossinsight.LanguageSolidity,
			ossinsight.LanguageAssembly,
			ossinsight.LanguageR,
			ossinsight.LanguageScala,
			ossinsight.LanguageJulia,
			ossinsight.LanguageLua,
			ossinsight.LanguageClojure,
			ossinsight.LanguageErlang,
			ossinsight.LanguageCommonLisp,
			ossinsight.LanguageEmacsLisp,
			ossinsight.LanguageOCaml,
			ossinsight.LanguageMATLAB,
			ossinsight.LanguageObjectiveC,
			ossinsight.LanguagePerl,
			ossinsight.LanguageFortran,
		),
	),
	mcp.WithNumber("limit",
		mcp.Description("Maximum number of repositories to return. Range from 1 to 100."),
		mcp.DefaultNumber(10),
		mcp.Min(1),
		mcp.Max(100),
	),
)
