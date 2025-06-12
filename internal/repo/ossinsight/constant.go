package ossinsight

const (
	// BaseURL is the base URL for OSSInsight API
	BaseURL = "https://api.ossinsight.io/v1"

	// Rate limit constants
	defaultRateLimitPerHour   = 600  // requests per hour per IP
	defaultRateLimitPerMinute = 1000 // requests per minute globally

	// HTTP headers
	headerRateLimitLimit           = "x-ratelimit-limit"
	headerRateLimitRemaining       = "x-ratelimit-remaining"
	headerRateLimitLimitMinute     = "x-ratelimit-limit-minute"
	headerRateLimitRemainingMinute = "x-ratelimit-remaining-minute"

	listTrendingReposEndpoint = "/trends/repos/"
)

// ListTrendReposPeriod 趋势仓库周期
//
//	@author centonhuang
//	@update 2025-06-09 20:39:20
type ListTrendReposPeriod = string

const (
	// ListTrendReposPeriodPast24Hours ListTrendReposPeriod 最近24小时
	//
	//	@author centonhuang
	//	@update 2025-06-09 20:39:45
	ListTrendReposPeriodPast24Hours ListTrendReposPeriod = "past_24_hours"

	// ListTrendReposPeriodPastWeek ListTrendReposPeriod 最近1周
	//
	//	@author centonhuang
	//	@update 2025-06-09 20:39:48
	ListTrendReposPeriodPastWeek ListTrendReposPeriod = "past_week"

	// ListTrendReposPeriodPastMonth ListTrendReposPeriod 最近1个月
	//
	//	@author centonhuang
	//	@update 2025-06-09 20:39:39
	ListTrendReposPeriodPastMonth ListTrendReposPeriod = "past_month"

	// ListTrendReposPeriodPast3Months ListTrendReposPeriod 最近3个月
	//	@update 2025-06-09 20:39:42
	ListTrendReposPeriodPast3Months ListTrendReposPeriod = "past_3_months"
)

// ListTrendReposLanguage 编程语言类型
//
//	@author centonhuang
//	@update 2025-01-15 00:00:00
type ListTrendReposLanguage = string

const (
	// LanguageAll Language 所有语言 (默认值)
	//
	//	@author centonhuang
	//	@update 2025-01-15 00:00:00
	LanguageAll ListTrendReposLanguage = "All"

	// LanguageJavaScript ListTrendReposLanguage JS
	//
	//	@author centonhuang
	//	@update 2025-06-09 20:41:36
	LanguageJavaScript ListTrendReposLanguage = "JavaScript"

	// LanguageJava ListTrendReposLanguage Java
	//
	//	@author centonhuang
	//	@update 2025-06-09 20:41:39
	LanguageJava ListTrendReposLanguage = "Java"

	// LanguagePython ListTrendReposLanguage Python
	//
	//	@author centonhuang
	//	@update 2025-06-09 20:41:42
	LanguagePython ListTrendReposLanguage = "Python"

	// LanguagePHP ListTrendReposLanguage PHP
	//
	//	@author centonhuang
	//	@update 2025-06-09 20:41:45
	LanguagePHP ListTrendReposLanguage = "PHP"

	// LanguageCPlusPlus ListTrendReposLanguage C++
	//
	//	@author centonhuang
	//	@update 2025-06-09 20:41:48
	LanguageCPlusPlus ListTrendReposLanguage = "C++"

	// LanguageCSharp ListTrendReposLanguage C#
	//
	//	@author centonhuang
	//	@update 2025-06-09 20:41:51
	LanguageCSharp ListTrendReposLanguage = "C#"

	// LanguageTypeScript ListTrendReposLanguage TypeScript
	//
	//	@author centonhuang
	//	@update 2025-06-09 20:41:54
	LanguageTypeScript ListTrendReposLanguage = "TypeScript"

	// LanguageShell ListTrendReposLanguage Shell
	//
	//	@author centonhuang
	//	@update 2025-06-09 20:41:57
	LanguageShell ListTrendReposLanguage = "Shell"

	// LanguageC ListTrendReposLanguage C
	//
	//	@author centonhuang
	//	@update 2025-06-09 20:42:00
	LanguageC ListTrendReposLanguage = "C"

	// LanguageRuby ListTrendReposLanguage Ruby
	//
	//	@author centonhuang
	//	@update 2025-06-09 20:42:03
	LanguageRuby ListTrendReposLanguage = "Ruby"

	// LanguageRust ListTrendReposLanguage Rust
	//
	//	@author centonhuang
	//	@update 2025-06-09 20:42:06
	LanguageRust ListTrendReposLanguage = "Rust"

	// LanguageGo ListTrendReposLanguage Go
	//
	//	@author centonhuang
	//	@update 2025-06-09 20:42:12
	LanguageGo ListTrendReposLanguage = "Go"

	// LanguageKotlin ListTrendReposLanguage Kotlin
	//
	//	@author centonhuang
	//	@update 2025-06-09 20:42:15
	LanguageKotlin ListTrendReposLanguage = "Kotlin"

	// LanguageHCL ListTrendReposLanguage HCL
	//
	//	@author centonhuang
	//	@update 2025-06-09 20:42:18
	LanguageHCL ListTrendReposLanguage = "HCL"

	// LanguagePowerShell ListTrendReposLanguage PowerShell
	//
	//	@author centonhuang
	//	@update 2025-06-09 20:42:21
	LanguagePowerShell ListTrendReposLanguage = "PowerShell"

	// LanguageCMake ListTrendReposLanguage CMake
	//
	//	@author centonhuang
	//	@update 2025-06-09 20:42:24
	LanguageCMake ListTrendReposLanguage = "CMake"

	// LanguageGroovy ListTrendReposLanguage Groovy
	//
	//	@author centonhuang
	//	@update 2025-06-09 20:42:27
	LanguageGroovy ListTrendReposLanguage = "Groovy"

	// LanguagePLpgSQL ListTrendReposLanguage PLpgSQL
	//
	//	@author centonhuang
	//	@update 2025-06-09 20:42:30
	LanguagePLpgSQL ListTrendReposLanguage = "PLpgSQL"

	// LanguageTSQL ListTrendReposLanguage TSQL
	//
	//	@author centonhuang
	//	@update 2025-06-09 20:42:33
	LanguageTSQL ListTrendReposLanguage = "TSQL"

	// LanguageDart ListTrendReposLanguage Dart
	//
	//	@author centonhuang
	//	@update 2025-06-09 20:42:36
	LanguageDart ListTrendReposLanguage = "Dart"

	// LanguageSwift ListTrendReposLanguage Swift
	//
	//	@author centonhuang
	//	@update 2025-06-09 20:42:39
	LanguageSwift ListTrendReposLanguage = "Swift"

	// LanguageHTML ListTrendReposLanguage HTML
	//
	//	@author centonhuang
	//	@update 2025-06-09 20:42:42
	LanguageHTML ListTrendReposLanguage = "HTML"

	// LanguageCSS ListTrendReposLanguage CSS
	//
	//	@author centonhuang
	//	@update 2025-06-09 20:42:45
	LanguageCSS ListTrendReposLanguage = "CSS"

	// LanguageElixir ListTrendReposLanguage Elixir
	//
	//	@author centonhuang
	//	@update 2025-06-09 20:42:48
	LanguageElixir ListTrendReposLanguage = "Elixir"

	// LanguageHaskell ListTrendReposLanguage Haskell
	//
	//	@author centonhuang
	//	@update 2025-06-09 20:42:51
	LanguageHaskell ListTrendReposLanguage = "Haskell"

	// LanguageSolidity ListTrendReposLanguage Solidity
	//
	//	@author centonhuang
	//	@update 2025-06-09 20:42:54
	LanguageSolidity ListTrendReposLanguage = "Solidity"

	// LanguageAssembly ListTrendReposLanguage Assembly
	//
	//	@author centonhuang
	//	@update 2025-06-09 20:42:57
	LanguageAssembly ListTrendReposLanguage = "Assembly"

	// LanguageR ListTrendReposLanguage R
	//
	//	@author centonhuang
	//	@update 2025-06-09 20:43:00
	LanguageR ListTrendReposLanguage = "R"

	// LanguageScala ListTrendReposLanguage Scala
	//
	//	@author centonhuang
	//	@update 2025-06-09 20:43:03
	LanguageScala ListTrendReposLanguage = "Scala"

	// LanguageJulia ListTrendReposLanguage Julia
	//
	//	@author centonhuang
	//	@update 2025-06-09 20:43:06
	LanguageJulia ListTrendReposLanguage = "Julia"

	// LanguageLua ListTrendReposLanguage Lua
	//
	//	@author centonhuang
	//	@update 2025-06-09 20:43:09
	LanguageLua ListTrendReposLanguage = "Lua"

	// LanguageClojure ListTrendReposLanguage Clojure
	//
	//	@author centonhuang
	//	@update 2025-06-09 20:43:12
	LanguageClojure ListTrendReposLanguage = "Clojure"

	// LanguageErlang ListTrendReposLanguage Erlang
	//
	//	@author centonhuang
	//	@update 2025-06-09 20:43:15
	LanguageErlang ListTrendReposLanguage = "Erlang"

	// LanguageCommonLisp ListTrendReposLanguage Common Lisp
	//
	//	@author centonhuang
	//	@update 2025-06-09 20:43:18
	LanguageCommonLisp ListTrendReposLanguage = "Common Lisp"

	// LanguageEmacsLisp ListTrendReposLanguage Emacs Lisp
	//
	//	@author centonhuang
	//	@update 2025-06-09 20:43:21
	LanguageEmacsLisp ListTrendReposLanguage = "Emacs Lisp"

	// LanguageOCaml ListTrendReposLanguage OCaml
	//
	//	@author centonhuang
	//	@update 2025-06-09 20:43:24
	LanguageOCaml ListTrendReposLanguage = "OCaml"

	// LanguageMATLAB ListTrendReposLanguage MATLAB
	//
	//	@author centonhuang
	//	@update 2025-06-09 20:43:27
	LanguageMATLAB ListTrendReposLanguage = "MATLAB"

	// LanguageObjectiveC ListTrendReposLanguage Objective-C
	//
	//	@author centonhuang
	//	@update 2025-06-09 20:43:30
	LanguageObjectiveC ListTrendReposLanguage = "Objective-C"

	// LanguagePerl ListTrendReposLanguage Perl
	//
	//	@author centonhuang
	//	@update 2025-06-09 20:43:33
	LanguagePerl ListTrendReposLanguage = "Perl"

	// LanguageFortran ListTrendReposLanguage Fortran
	//
	//	@author centonhuang
	//	@update 2025-
	LanguageFortran ListTrendReposLanguage = "Fortran"
)
