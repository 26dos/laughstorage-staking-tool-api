# LaughStorage Staking Tool API

一个基于 GoFrame 框架开发的区块链质押工具 API 服务，支持 Filecoin 网络上的质押管理和提案审核功能。

## 📋 目录

- [环境要求](#环境要求)
- [快速开始](#快速开始)
- [配置说明](#配置说明)
- [API 文档](#api-文档)
- [开发指南](#开发指南)
- [部署指南](#部署指南)
- [常见问题](#常见问题)

## 🛠️ 环境要求

### 必需软件

- **Go**: 1.23.7 或更高版本
- **MySQL**: 8.x 或更高版本
- **Redis**: 6.0 或更高版本
- **Git**: 用于代码版本控制

### Go 环境安装

#### Linux/macOS

```bash
# 下载并安装 Go
wget https://golang.org/dl/go1.23.7.linux-amd64.tar.gz
sudo tar -C /usr/local -xzf go1.23.7.linux-amd64.tar.gz

# 配置环境变量
echo 'export PATH=$PATH:/usr/local/go/bin' >> ~/.bashrc
echo 'export GOPATH=$HOME/go' >> ~/.bashrc
echo 'export GOROOT=/usr/local/go' >> ~/.bashrc
source ~/.bashrc

# 验证安装
go version
```

#### Windows

1. 从 [Go 官网](https://golang.org/dl/) 下载 Windows 安装包
2. 运行安装程序并按照提示完成安装
3. 重启命令行工具并验证：`go version`

### GoFrame CLI 工具

本项目使用 GoFrame 框架进行开发，需要安装 GoFrame CLI 工具：

```bash
# 安装 GoFrame CLI
go install github.com/gogf/gf/v2/cmd/gf/v2@latest

# 验证安装
gf version
```

更多信息请访问 [GoFrame 官网](https://goframe.org/)

## 🚀 快速开始

### 1. 克隆项目

```bash
git clone <your-repository-url>
cd api
```

### 2. 安装依赖

```bash
go mod tidy
```

### 3. 配置数据库

1. 创建 MySQL 数据库
2. 导入数据库结构：
   ```bash
   mysql -u your_username -p your_database < script.sql
   ```

### 4. 配置环境

复制并修改配置文件：

```bash
cp config.toml.example config.toml
# 编辑 config.toml 文件，配置数据库、Redis 等连接信息
```

### 5. 运行项目

```bash
# 开发模式运行
go run main.go

# 或者使用 Makefile
make build
./laughstorage_staking_tool
```

服务将在 `http://127.0.0.1:9002` 启动

## ⚙️ 配置说明

### 配置文件结构

项目使用 `config.toml` 作为主配置文件，包含以下主要配置项：

#### 服务器配置

```toml
[server]
address = "127.0.0.1:9002"  # 服务监听地址
openapiPath = ""           # OpenAPI 文档路径
swaggerPath = ""           # Swagger 文档路径
appDebug = true           # 调试模式
```

#### 数据库配置

```toml
[database.default]
type = "mysql"
host = "your_mysql_host"
port = "3306"
user = "your_username"
pass = "your_password"
name = "your_database"
timezone = "UTC"
parseTime = true
debug = true
```

#### Redis 配置

```toml
[redis.default]
address = "127.0.0.1:6379"
db = 4
pass = "your_redis_password"
```

#### JWT 配置

```toml
[jwt]
exp = 72000                                    # Token 有效期（秒）
pre = "laughstorage_staking_tool_"            # Redis 中 Token 前缀
issuer = "laughstorage_staking_tool_"         # Token 发行者
key = "your_jwt_secret_key"                   # JWT 密钥
encryptKey = "your_encryption_key"            # 加密密钥
```

#### SIWE (Sign-In with Ethereum) 配置

```toml
[siwe]
domain = "localhost:9000"  # Web3 登录验证域名，必须与前端一致
```

#### 阿里云配置 (KYC 认证)

```toml
[ali]
accessKeyId = "your_aliyun_access_key_id"
accessKeySecret = "your_aliyun_access_key_secret"
ossBucketName = "your_oss_bucket_name"
ossUrl = "your_oss_url"
region = "cn-hongkong"
```

#### 应用配置

```toml
[app]
administrator = "0x..."                    # 管理员钱包地址（拥有审核提案权限，多个用英文逗号分隔）
rpc = "https://filecoin-calibration.chainup.net/rpc/v1"  # Filecoin RPC 地址
stakingContract = "0x..."                  # 质押合约地址
```

#### 分配器配置

```allocator
[allocator]
contract = ""   # 分配器合约地址
account  = ""   # 公证人地址
accountPrivateKey = ""  # 公证人私钥
```

### 重要配置说明

1. **administrator**: 这是最重要的配置项，指定拥有审核提案权限的钱包地址
2. **domain**: SIWE 配置中的域名必须与前端应用的域名完全一致
3. **数据库**: 确保数据库连接信息正确，并已导入 `script.sql` 文件
4. **Redis**: 用于缓存和会话管理，确保连接正常

## 🛠️ 开发指南

### 项目结构

```
api/
├── api/                    # API 路由定义
├── internal/              # 内部包
│   ├── cmd/              # 命令行入口
│   ├── controller/       # 控制器层
│   ├── dao/             # 数据访问层
│   ├── logic/           # 业务逻辑层
│   ├── middleware/      # 中间件
│   ├── model/           # 数据模型
│   └── service/         # 服务层
├── resource/            # 静态资源
├── manifest/           # 部署配置
└── main.go            # 程序入口
```

### 常用开发命令

```bash
# 生成控制器代码
gf ge  ctrl

# 生成数据访问层代码
gf gen dao

# 构建项目
gf build

```

### 代码生成

项目使用 GoFrame 的代码生成功能：

```bash
# 根据 API 定义生成控制器
gf gen ctrl

# 根据数据库表生成 DAO/DO/Entity
gf gen dao

```

## 🚀 部署指南

### 生产环境配置

1. 修改 `config.toml` 中的生产环境配置
2. 确保数据库和 Redis 连接稳定
3. 配置反向代理（如 Nginx）
4. 设置 SSL 证书
5. 配置日志轮转

## 🔧 常见问题

### Q: 启动时提示数据库连接失败

A: 检查 `config.toml` 中的数据库配置，确保：

- 数据库服务正在运行
- 连接信息正确
- 数据库已创建并导入了 `script.sql`

### Q: JWT Token 验证失败

A: 检查 JWT 配置：

- 确保 `key` 和 `encryptKey` 配置正确
- 检查 Token 是否过期
- 验证 Redis 连接正常

### Q: KYC 认证失败

A: 检查阿里云配置：

- 确保 AccessKey 有效
- 检查 OSS Bucket 权限
- 验证区域配置正确

### Q: Web3 登录失败

A: 检查 SIWE 配置：

- 确保 `domain` 与前端域名一致
- 检查网络连接
- 验证钱包连接正常

## 📝 日志

日志文件位于 `log/` 目录下，按日期自动分割。可以通过修改 `config.toml` 中的日志配置来调整日志级别和输出格式。

## 🤝 贡献指南

1. Fork 项目
2. 创建功能分支 (`git checkout -b feature/AmazingFeature`)
3. 提交更改 (`git commit -m 'Add some AmazingFeature'`)
4. 推送到分支 (`git push origin feature/AmazingFeature`)
5. 打开 Pull Request

## 📄 许可证

本项目采用 MIT 许可证 - 查看 [LICENSE](LICENSE) 文件了解详情。

## 📞 支持

如果您遇到问题或有任何疑问，请：

1. 查看 [常见问题](#常见问题) 部分
2. 检查项目 [Issues](https://github.com/your-repo/issues)
3. 创建新的 Issue 描述您的问题

---

**注意**: 在生产环境中部署前，请确保：

- 所有敏感信息已正确配置
- 数据库已备份
- 监控和日志系统已配置
- 安全措施已实施
