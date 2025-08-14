# OSS Insight MCP Server

[English](#english) | [中文](#chinese)

---

## English

### Overview

OSS Insight MCP Server is a comprehensive Model Context Protocol (MCP) server that provides complete encapsulation of the OSS Insight API. OSS Insight is a powerful GitHub analytics platform that offers real-time insights and analytics about GitHub repositories, developers, and trends.

### Features

#### 🚀 Trend Analysis
- **ListTrendingRepos**: Get trending GitHub repositories based on various time periods and programming languages
- Support for past 24 hours, past week, past month, past 3 months time ranges
- Support for 40+ programming languages filtering

#### 📊 Ranking System
- **GetRepoRanking**: Get repository rankings by stars, forks, commits, contributors, pull requests, issues
- **GetDeveloperRanking**: Get developer rankings by stars, forks, commits, pull requests, issues
- Support for multiple time ranges: past 24 hours, past week, past month, past 3 months, past year, all time

#### 📚 Collection Management
- **GetCollections**: Get GitHub collections (curated lists of repositories)
- **GetHotCollections**: Get hot collections that are currently trending
- **GetCollectionRepos**: Get repositories from specific collections

#### 🔍 Detailed Information
- **GetRepoDetail**: Get comprehensive repository information
- **GetDeveloperDetail**: Get comprehensive developer information

#### 📈 Statistical Analysis
- **GetLanguageStats**: Get programming language statistics and insights

### Architecture

```
.
├── cmd/                    # Command line tools
│   ├── client.go          # MCP client implementation
│   ├── root.go            # Root command configuration
│   └── server.go          # MCP server implementation
├── internal/               # Internal packages
│   ├── config/            # Configuration management
│   ├── constant/          # Constant definitions
│   ├── logger/            # Logging management
│   ├── middleware/        # Middleware components
│   ├── repo/              # Data repository layer
│   │   ├── mcp/           # MCP-related components
│   │   │   ├── context.go # Context management
│   │   │   ├── handler.go # MCP tool handlers
│   │   │   └── schema.go  # MCP tool schemas
│   │   └── ossinsight/    # OSS Insight API client
│   │       ├── client.go  # HTTP client implementation
│   │       ├── constant.go # API constants and endpoints
│   │       ├── query.go   # Query parameter structures
│   │       ├── response.go # Response data structures
│   │       └── util.go    # Utility functions
│   └── service/           # Service layer
│       └── service.go     # Tool registration and service management
├── main.go                # Main application entry point
├── Makefile               # Build and development automation
├── dockerfile             # Docker configuration
├── go.mod                 # Go module dependencies
├── go.sum                 # Go module checksums
└── README.md              # Project documentation
```

### Installation & Setup

#### Prerequisites
- Go 1.21+
- Valid OSS Insight API key

#### Quick Start

1. **Clone the repository**
   ```bash
   git clone https://github.com/your-username/ossinsight-mcp.git
   cd ossinsight-mcp
   ```

2. **Install dependencies**
   ```bash
   make deps
   # or
   go mod download
   go mod tidy
   ```

3. **Configure API key**
   ```bash
   # Edit internal/config/config.go
   var APIKey = "your-api-key-here"
   ```

4. **Build and run**
   ```bash
   # Build the application
   make build
   
   # Start the MCP server
   make server
   
   # In another terminal, test the client
   make client
   ```

#### Using Makefile

```bash
# Show all available commands
make help

# Run all tests
make test

# Build for multiple platforms
make release

# Run with Docker
make docker-build
make docker-run
```

### API Tools Reference

#### ListTrendingRepos
Get trending repositories that are gaining popularity in the GitHub ecosystem.

**Parameters:**
- `period` (required): Time period - `past_24_hours`, `past_week`, `past_month`, `past_3_months`
- `language` (required): Programming language - 40+ languages supported
- `limit` (optional): Number of repositories to return (1-100, default: 10)

**Example:**
```json
{
  "period": "past_24_hours",
  "language": "Go",
  "limit": 20
}
```

#### GetRepoRanking
Get repository rankings based on various metrics.

**Parameters:**
- `ranking_type` (required): `stars`, `forks`, `commits`, `contributors`, `pull_requests`, `issues`
- `time_range` (required): Time range for ranking data
- `language` (optional): Programming language filter
- `limit` (optional): Number of repositories to return (1-100, default: 10)

#### GetDeveloperRanking
Get developer rankings based on various metrics.

**Parameters:**
- `ranking_type` (required): `stars`, `forks`, `commits`, `pull_requests`, `issues`
- `time_range` (required): Time range for ranking data
- `language` (optional): Programming language filter
- `limit` (optional): Number of developers to return (1-100, default: 10)

#### GetCollections
Get GitHub collections (curated lists of repositories).

**Parameters:**
- `collection_type` (required): `all`, `hot`, `user`
- `time_range` (required): Time range for collections data
- `limit` (optional): Number of collections to return (1-100, default: 10)

#### GetHotCollections
Get hot collections that are currently trending.

**Parameters:**
- `time_range` (required): Time range for hot collections data
- `limit` (optional): Number of collections to return (1-100, default: 10)

#### GetCollectionRepos
Get repositories from a specific collection.

**Parameters:**
- `collection_id` (required): ID of the collection
- `time_range` (required): Time range for collection repositories data
- `limit` (optional): Number of repositories to return (1-100, default: 10)

#### GetRepoDetail
Get detailed information about a specific repository.

**Parameters:**
- `owner` (required): Repository owner (username or organization)
- `repo` (required): Repository name

#### GetDeveloperDetail
Get detailed information about a specific developer.

**Parameters:**
- `username` (required): GitHub username

#### GetLanguageStats
Get programming language statistics and insights.

**Parameters:**
- `time_range` (required): Time range for language statistics data
- `limit` (optional): Number of languages to return (1-100, default: 10)

### Supported Parameters

#### Time Ranges
- `past_24_hours`: Past 24 hours
- `past_week`: Past week
- `past_month`: Past month
- `past_3_months`: Past 3 months
- `past_year`: Past year
- `all_time`: All time

#### Programming Languages
Supports 40+ programming languages including:
- `Go`, `JavaScript`, `Python`, `Java`, `PHP`
- `C++`, `C#`, `TypeScript`, `Rust`, `Kotlin`
- `Ruby`, `Swift`, `Dart`, `Scala`, `Julia`
- And many more...

#### Ranking Types

**Repository Rankings:**
- `stars`: Rank by star count
- `forks`: Rank by fork count
- `commits`: Rank by commit count
- `contributors`: Rank by contributor count
- `pull_requests`: Rank by pull request count
- `issues`: Rank by issue count

**Developer Rankings:**
- `stars`: Rank by stars received
- `forks`: Rank by forks received
- `commits`: Rank by commit count
- `pull_requests`: Rank by pull request count
- `issues`: Rank by issue count

### Error Handling

All tools include comprehensive error handling:
- Parameter validation errors
- API call errors
- Rate limit information
- Detailed error logging

### Rate Limiting

OSS Insight API includes rate limiting:
- 600 requests per hour per IP
- 1000 requests per minute globally

All responses include rate limit information to help clients manage API usage.

### Testing

```bash
# Run all tests
make test

# Run tests with coverage
make test-coverage

# Run quick tests (core functionality only)
make quick-test

# Run tests for specific package
make test-package PACKAGE=internal/repo/ossinsight
```

### Development

#### Adding New Tools

1. Add new API endpoints in `internal/repo/ossinsight/constant.go`
2. Add query parameter structures in `internal/repo/ossinsight/query.go`
3. Add response data structures in `internal/repo/ossinsight/response.go`
4. Add client methods in `internal/repo/ossinsight/client.go`
5. Add tool schemas in `internal/repo/mcp/schema.go`
6. Add handler logic in `internal/repo/mcp/handler.go`
7. Register new tools in `internal/service/service.go`

#### Code Quality

```bash
# Format code
make fmt

# Lint code (requires golangci-lint)
make lint

# Run race condition tests
make race

# Run benchmark tests
make bench
```

### Deployment

#### Docker

```bash
# Build Docker image
make docker-build

# Run Docker container
make docker-run
```

#### Binary Release

```bash
# Create release builds for multiple platforms
make release
```

### Contributing

1. Fork the repository
2. Create a feature branch
3. Make your changes
4. Add tests for new functionality
5. Ensure all tests pass
6. Submit a pull request

### License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

### Support

For questions or suggestions, please open an issue on GitHub.

---

## Chinese

### 概述

OSS Insight MCP Server 是一个基于 Model Context Protocol (MCP) 的完整服务器，提供了对 OSS Insight API 的全面封装。OSS Insight 是一个强大的 GitHub 分析平台，提供关于 GitHub 仓库、开发者和趋势的实时洞察和分析。

### 功能特性

#### 🚀 趋势分析
- **ListTrendingRepos**: 获取趋势 GitHub 仓库列表，基于各种时间周期和编程语言
- 支持过去 24 小时、过去一周、过去一个月、过去三个月等时间范围
- 支持 40+ 种编程语言过滤

#### 📊 排名系统
- **GetRepoRanking**: 获取仓库排名，支持按星标、分支、提交、贡献者、拉取请求、问题等指标排名
- **GetDeveloperRanking**: 获取开发者排名，支持按星标、分支、提交、拉取请求、问题等指标排名
- 支持多种时间范围：过去 24 小时、过去一周、过去一个月、过去三个月、过去一年、所有时间

#### 📚 收藏夹管理
- **GetCollections**: 获取 GitHub 收藏夹列表，支持所有类型、热门、用户创建等分类
- **GetHotCollections**: 获取热门收藏夹，发现当前流行的精选仓库列表
- **GetCollectionRepos**: 获取特定收藏夹中的仓库列表

#### 🔍 详细信息
- **GetRepoDetail**: 获取仓库详细信息，包括星标、分支、提交、贡献者、许可证等
- **GetDeveloperDetail**: 获取开发者详细信息，包括仓库、贡献、关注者、公司等

#### 📈 统计分析
- **GetLanguageStats**: 获取编程语言统计信息，包括流行度、增长率、使用统计等

### 项目架构

```
.
├── cmd/                    # 命令行工具
│   ├── client.go          # MCP 客户端实现
│   ├── root.go            # 根命令配置
│   └── server.go          # MCP 服务器实现
├── internal/               # 内部包
│   ├── config/            # 配置管理
│   ├── constant/          # 常量定义
│   ├── logger/            # 日志管理
│   ├── middleware/        # 中间件组件
│   ├── repo/              # 数据仓库层
│   │   ├── mcp/           # MCP 相关组件
│   │   │   ├── context.go # 上下文管理
│   │   │   ├── handler.go # MCP 工具处理器
│   │   │   └── schema.go  # MCP 工具模式定义
│   │   └── ossinsight/    # OSS Insight API 客户端
│   │       ├── client.go  # HTTP 客户端实现
│   │       ├── constant.go # API 常量和端点
│   │       ├── query.go   # 查询参数结构
│   │       ├── response.go # 响应数据结构
│   │       └── util.go    # 工具函数
│   └── service/           # 服务层
│       └── service.go     # 工具注册和服务管理
├── main.go                # 主程序入口点
├── Makefile               # 构建和开发自动化
├── dockerfile             # Docker 配置
├── go.mod                 # Go 模块依赖
├── go.sum                 # Go 模块校验和
└── README.md              # 项目文档
```

### 安装和设置

#### 前置要求
- Go 1.21+
- 有效的 OSS Insight API 密钥

#### 快速开始

1. **克隆仓库**
   ```bash
   git clone https://github.com/your-username/ossinsight-mcp.git
   cd ossinsight-mcp
   ```

2. **安装依赖**
   ```bash
   make deps
   # 或者
   go mod download
   go mod tidy
   ```

3. **配置 API 密钥**
   ```bash
   # 编辑 internal/config/config.go
   var APIKey = "your-api-key-here"
   ```

4. **构建和运行**
   ```bash
   # 构建应用程序
   make build
   
   # 启动 MCP 服务器
   make server
   
   # 在另一个终端中测试客户端
   make client
   ```

#### 使用 Makefile

```bash
# 显示所有可用命令
make help

# 运行所有测试
make test

# 为多个平台构建
make release

# 使用 Docker 运行
make docker-build
make docker-run
```

### API 工具参考

#### ListTrendingRepos
获取在 GitHub 生态系统中获得流行的趋势仓库。

**参数:**
- `period` (必需): 时间周期 - `past_24_hours`, `past_week`, `past_month`, `past_3_months`
- `language` (必需): 编程语言 - 支持 40+ 种语言
- `limit` (可选): 返回仓库数量 (1-100，默认: 10)

**示例:**
```json
{
  "period": "past_24_hours",
  "language": "Go",
  "limit": 20
}
```

#### GetRepoRanking
基于各种指标获取仓库排名。

**参数:**
- `ranking_type` (必需): `stars`, `forks`, `commits`, `contributors`, `pull_requests`, `issues`
- `time_range` (必需): 排名数据的时间范围
- `language` (可选): 编程语言过滤
- `limit` (可选): 返回仓库数量 (1-100，默认: 10)

#### GetDeveloperRanking
基于各种指标获取开发者排名。

**参数:**
- `ranking_type` (必需): `stars`, `forks`, `commits`, `pull_requests`, `issues`
- `time_range` (必需): 排名数据的时间范围
- `language` (可选): 编程语言过滤
- `limit` (可选): 返回开发者数量 (1-100，默认: 10)

#### GetCollections
获取 GitHub 收藏夹（精选仓库列表）。

**参数:**
- `collection_type` (必需): `all`, `hot`, `user`
- `time_range` (必需): 收藏夹数据的时间范围
- `limit` (可选): 返回收藏夹数量 (1-100，默认: 10)

#### GetHotCollections
获取当前流行的热门收藏夹。

**参数:**
- `time_range` (必需): 热门收藏夹数据的时间范围
- `limit` (可选): 返回收藏夹数量 (1-100，默认: 10)

#### GetCollectionRepos
获取特定收藏夹中的仓库。

**参数:**
- `collection_id` (必需): 收藏夹 ID
- `time_range` (必需): 收藏夹仓库数据的时间范围
- `limit` (可选): 返回仓库数量 (1-100，默认: 10)

#### GetRepoDetail
获取特定仓库的详细信息。

**参数:**
- `owner` (必需): 仓库所有者（用户名或组织）
- `repo` (必需): 仓库名称

#### GetDeveloperDetail
获取特定开发者的详细信息。

**参数:**
- `username` (必需): GitHub 用户名

#### GetLanguageStats
获取编程语言统计信息和洞察。

**参数:**
- `time_range` (必需): 语言统计数据的时间范围
- `limit` (可选): 返回语言数量 (1-100，默认: 10)

### 支持的参数

#### 时间范围
- `past_24_hours`: 过去 24 小时
- `past_week`: 过去一周
- `past_month`: 过去一个月
- `past_3_months`: 过去三个月
- `past_year`: 过去一年
- `all_time`: 所有时间

#### 编程语言
支持 40+ 种编程语言，包括：
- `Go`, `JavaScript`, `Python`, `Java`, `PHP`
- `C++`, `C#`, `TypeScript`, `Rust`, `Kotlin`
- `Ruby`, `Swift`, `Dart`, `Scala`, `Julia`
- 以及更多...

#### 排名类型

**仓库排名:**
- `stars`: 按星标数排名
- `forks`: 按分支数排名
- `commits`: 按提交数排名
- `contributors`: 按贡献者数排名
- `pull_requests`: 按拉取请求数排名
- `issues`: 按问题数排名

**开发者排名:**
- `stars`: 按获得的星标数排名
- `forks`: 按获得的分支数排名
- `commits`: 按提交数排名
- `pull_requests`: 按拉取请求数排名
- `issues`: 按问题数排名

### 错误处理

所有工具都包含完善的错误处理机制：
- 参数验证错误
- API 调用错误
- 速率限制信息
- 详细的错误日志

### 速率限制

OSS Insight API 包含速率限制：
- 每小时每 IP 600 个请求
- 每分钟全局 1000 个请求

所有响应都包含速率限制信息，帮助客户端管理 API 使用。

### 测试

```bash
# 运行所有测试
make test

# 运行带覆盖率的测试
make test-coverage

# 运行快速测试（仅核心功能）
make quick-test

# 运行特定包的测试
make test-package PACKAGE=internal/repo/ossinsight
```

### 开发

#### 添加新工具

1. 在 `internal/repo/ossinsight/constant.go` 中添加新的 API 端点
2. 在 `internal/repo/ossinsight/query.go` 中添加查询参数结构
3. 在 `internal/repo/ossinsight/response.go` 中添加响应数据结构
4. 在 `internal/repo/ossinsight/client.go` 中添加客户端方法
5. 在 `internal/repo/mcp/schema.go` 中添加工具模式定义
6. 在 `internal/repo/mcp/handler.go` 中添加处理逻辑
7. 在 `internal/service/service.go` 中注册新工具

#### 代码质量

```bash
# 格式化代码
make fmt

# 代码检查（需要 golangci-lint）
make lint

# 运行竞态条件测试
make race

# 运行基准测试
make bench
```

### 部署

#### Docker

```bash
# 构建 Docker 镜像
make docker-build

# 运行 Docker 容器
make docker-run
```

#### 二进制发布

```bash
# 为多个平台创建发布构建
make release
```

### 贡献

1. Fork 仓库
2. 创建功能分支
3. 进行更改
4. 为新功能添加测试
5. 确保所有测试通过
6. 提交拉取请求

### 许可证

本项目采用 MIT 许可证。详见 [LICENSE](LICENSE) 文件。

### 支持

如有问题或建议，请在 GitHub 上提交 issue。