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
	mcp.WithDescription(""),
	mcp.WithString("period",
		mcp.Description(""),
		mcp.Enum(
			ossinsight.ListTrendReposPeriodPast24Hours,
			ossinsight.ListTrendReposPeriodPastWeek,
			ossinsight.ListTrendReposPeriodPastMonth,
			ossinsight.ListTrendReposPeriodPast3Months,
		),
	),
	mcp.WithString("language",
		mcp.Description(""),
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
)
