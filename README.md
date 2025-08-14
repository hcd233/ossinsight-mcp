# OSS Insight MCP Server

OSS Insight MCP Server 是一个基于 Model Context Protocol (MCP) 的服务器，提供了对 OSS Insight API 的完整封装。OSS Insight 是一个强大的 GitHub 分析平台，提供关于 GitHub 仓库、开发者和趋势的实时洞察和分析。

## 功能特性

### 🚀 趋势分析
- **ListTrendingRepos**: 获取趋势 GitHub 仓库列表，基于各种时间周期和编程语言
- 支持过去 24 小时、过去一周、过去一个月、过去三个月等时间范围
- 支持 40+ 种编程语言过滤

### 📊 排名系统
- **GetRepoRanking**: 获取仓库排名，支持按星标、分支、提交、贡献者、拉取请求、问题等指标排名
- **GetDeveloperRanking**: 获取开发者排名，支持按星标、分支、提交、拉取请求、问题等指标排名
- 支持多种时间范围：过去 24 小时、过去一周、过去一个月、过去三个月、过去一年、所有时间

### 📚 收藏夹管理
- **GetCollections**: 获取 GitHub 收藏夹列表，支持所有类型、热门、用户创建等分类
- **GetHotCollections**: 获取热门收藏夹，发现当前流行的精选仓库列表
- **GetCollectionRepos**: 获取特定收藏夹中的仓库列表

### 🔍 详细信息
- **GetRepoDetail**: 获取仓库详细信息，包括星标、分支、提交、贡献者、许可证等
- **GetDeveloperDetail**: 获取开发者详细信息，包括仓库、贡献、关注者、公司等

### 📈 统计分析
- **GetLanguageStats**: 获取编程语言统计信息，包括流行度、增长率、使用统计等

## 安装和运行

### 前置要求
- Go 1.21+
- 有效的 OSS Insight API 密钥

### 安装依赖
```bash
go mod download
```

### 配置 API 密钥
在 `internal/config/config.go` 中设置你的 OSS Insight API 密钥：
```go
var APIKey = "your-api-key-here"
```

### 启动服务器
```bash
# 启动 MCP 服务器
go run main.go server --host localhost --port 8080

# 或者构建后运行
go build -o ossinsight-mcp
./ossinsight-mcp server --host localhost --port 8080
```

### 测试客户端
```bash
# 测试 MCP 客户端
go run main.go client --endpoint http://localhost:8080/mcp
```

## API 工具详细说明

### ListTrendingRepos
获取趋势仓库列表，帮助发现 GitHub 生态系统中流行和快速增长的仓库。

**参数:**
- `period` (必需): 时间周期 - `past_24_hours`, `past_week`, `past_month`, `past_3_months`
- `language` (必需): 编程语言 - 支持 40+ 种语言，包括 `Go`, `JavaScript`, `Python`, `Java` 等
- `limit` (可选): 返回仓库数量，默认 10，范围 1-100

**示例:**
```json
{
  "period": "past_24_hours",
  "language": "Go",
  "limit": 20
}
```

### GetRepoRanking
获取仓库排名，基于各种指标识别 GitHub 生态系统中最受欢迎和最活跃的仓库。

**参数:**
- `ranking_type` (必需): 排名类型 - `stars`, `forks`, `commits`, `contributors`, `pull_requests`, `issues`
- `time_range` (必需): 时间范围 - `past_24_hours`, `past_week`, `past_month`, `past_3_months`, `past_year`, `all_time`
- `language` (可选): 编程语言过滤
- `limit` (可选): 返回仓库数量，默认 10，范围 1-100

**示例:**
```json
{
  "ranking_type": "stars",
  "time_range": "past_month",
  "language": "Python",
  "limit": 15
}
```

### GetDeveloperRanking
获取开发者排名，基于各种指标识别 GitHub 生态系统中最活跃和最有影响力的开发者。

**参数:**
- `ranking_type` (必需): 排名类型 - `stars`, `forks`, `commits`, `pull_requests`, `issues`
- `time_range` (必需): 时间范围
- `language` (可选): 编程语言过滤
- `limit` (可选): 返回开发者数量，默认 10，范围 1-100

### GetCollections
获取 GitHub 收藏夹列表，发现由社区创建的精选仓库集合。

**参数:**
- `collection_type` (必需): 收藏夹类型 - `all`, `hot`, `user`
- `time_range` (必需): 时间范围
- `limit` (可选): 返回收藏夹数量，默认 10，范围 1-100

### GetHotCollections
获取热门收藏夹，发现当前流行和受欢迎的精选仓库列表。

**参数:**
- `time_range` (必需): 时间范围
- `limit` (可选): 返回收藏夹数量，默认 10，范围 1-100

### GetCollectionRepos
获取特定收藏夹中的仓库列表，探索精选在特定收藏夹中的仓库。

**参数:**
- `collection_id` (必需): 收藏夹 ID
- `time_range` (必需): 时间范围
- `limit` (可选): 返回仓库数量，默认 10，范围 1-100

### GetRepoDetail
获取仓库详细信息，提供关于仓库的全面数据，包括星标、分支、提交、贡献者等。

**参数:**
- `owner` (必需): 仓库所有者（用户名或组织）
- `repo` (必需): 仓库名称

**示例:**
```json
{
  "owner": "golang",
  "repo": "go"
}
```

### GetDeveloperDetail
获取开发者详细信息，提供关于开发者的全面数据，包括他们的仓库、贡献、关注者等。

**参数:**
- `username` (必需): GitHub 用户名

### GetLanguageStats
获取编程语言统计信息，提供关于语言流行度、增长率和使用统计的洞察。

**参数:**
- `time_range` (必需): 时间范围
- `limit` (可选): 返回语言数量，默认 10，范围 1-100

## 支持的时间范围

- `past_24_hours`: 过去 24 小时
- `past_week`: 过去一周
- `past_month`: 过去一个月
- `past_3_months`: 过去三个月
- `past_year`: 过去一年
- `all_time`: 所有时间

## 支持的编程语言

支持 40+ 种编程语言，包括但不限于：
- `Go`, `JavaScript`, `Python`, `Java`, `PHP`
- `C++`, `C#`, `TypeScript`, `Rust`, `Kotlin`
- `Ruby`, `Swift`, `Dart`, `Scala`, `Julia`
- 以及更多...

## 支持的排名类型

### 仓库排名
- `stars`: 按星标数排名
- `forks`: 按分支数排名
- `commits`: 按提交数排名
- `contributors`: 按贡献者数排名
- `pull_requests`: 按拉取请求数排名
- `issues`: 按问题数排名

### 开发者排名
- `stars`: 按获得的星标数排名
- `forks`: 按获得的分支数排名
- `commits`: 按提交数排名
- `pull_requests`: 按拉取请求数排名
- `issues`: 按问题数排名

## 错误处理

所有工具都包含完善的错误处理机制：
- 参数验证错误
- API 调用错误
- 速率限制信息
- 详细的错误日志

## 速率限制

OSS Insight API 包含速率限制：
- 每小时每 IP 600 个请求
- 每分钟全局 1000 个请求

所有响应都包含速率限制信息，帮助客户端管理 API 使用。

## 开发

### 项目结构
```
.
├── cmd/                    # 命令行工具
│   ├── client.go          # MCP 客户端
│   ├── root.go            # 根命令
│   └── server.go          # MCP 服务器
├── internal/               # 内部包
│   ├── config/            # 配置管理
│   ├── constant/          # 常量定义
│   ├── logger/            # 日志管理
│   ├── middleware/        # 中间件
│   ├── repo/              # 数据仓库
│   │   ├── mcp/           # MCP 相关
│   │   └── ossinsight/    # OSS Insight API 客户端
│   └── service/           # 服务层
├── main.go                # 主程序
└── README.md              # 说明文档
```

### 添加新工具
1. 在 `internal/repo/ossinsight/constant.go` 中添加新的 API 端点
2. 在 `internal/repo/ossinsight/query.go` 中添加查询参数结构
3. 在 `internal/repo/ossinsight/response.go` 中添加响应数据结构
4. 在 `internal/repo/ossinsight/client.go` 中添加客户端方法
5. 在 `internal/repo/mcp/schema.go` 中添加工具 schema
6. 在 `internal/repo/mcp/handler.go` 中添加处理逻辑
7. 在 `internal/service/service.go` 中注册新工具

## 许可证

本项目采用 MIT 许可证。详见 [LICENSE](LICENSE) 文件。

## 贡献

欢迎提交 Issue 和 Pull Request 来改进这个项目。

## 联系方式

如有问题或建议，请通过 GitHub Issues 联系我们。