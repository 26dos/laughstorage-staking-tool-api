# LaughStorage Staking Tool API

A blockchain staking tool API service developed based on the GoFrame framework, supporting staking management and proposal review functionality on the Filecoin network.

## 📋 Table of Contents

- [Requirements](#requirements)
- [Quick Start](#quick-start)
- [Configuration](#configuration)
- [API Documentation](#api-documentation)
- [Development Guide](#development-guide)
- [Deployment Guide](#deployment-guide)
- [FAQ](#faq)

## 🛠️ Requirements

### Required Software

- **Go**: 1.23.7 or higher
- **MySQL**: 8.x or higher
- **Redis**: 6.0 or higher
- **Git**: For code version control

### Go Environment Installation

#### Linux/macOS

```bash
# Download and install Go
wget https://golang.org/dl/go1.23.7.linux-amd64.tar.gz
sudo tar -C /usr/local -xzf go1.23.7.linux-amd64.tar.gz

# Configure environment variables
echo 'export PATH=$PATH:/usr/local/go/bin' >> ~/.bashrc
echo 'export GOPATH=$HOME/go' >> ~/.bashrc
echo 'export GOROOT=/usr/local/go' >> ~/.bashrc
source ~/.bashrc

# Verify installation
go version
```

#### Windows

1. Download the Windows installer from [Go official website](https://golang.org/dl/)
2. Run the installer and follow the prompts to complete installation
3. Restart command line tools and verify: `go version`

### GoFrame CLI Tool

This project uses the GoFrame framework for development and requires the GoFrame CLI tool:

```bash
# Install GoFrame CLI
go install github.com/gogf/gf/v2/cmd/gf/v2@latest

# Verify installation
gf version
```

For more information, visit the [GoFrame official website](https://goframe.org/)

## 🚀 Quick Start

### 1. Clone the Project

```bash
git clone <your-repository-url>
cd api
```

### 2. Install Dependencies

```bash
go mod tidy
```

### 3. Configure Database

1. Create a MySQL database
2. Import database structure:
   ```bash
   mysql -u your_username -p your_database < script.sql
   ```

### 4. Configure Environment

Copy and modify the configuration file:

```bash
cp config.toml.example config.toml
# Edit config.toml file to configure database, Redis, and other connection information
```

### 5. Run the Project

```bash
# Run in development mode
go run main.go

# Or use Makefile
make build
./laughstorage_staking_tool
```

The service will start at `http://127.0.0.1:9002`

## ⚙️ Configuration

### Configuration File Structure

The project uses `config.toml` as the main configuration file, containing the following main configuration items:

#### Server Configuration

```toml
[server]
address = "127.0.0.1:9002"  # Service listening address
openapiPath = ""           # OpenAPI documentation path
swaggerPath = ""           # Swagger documentation path
appDebug = true           # Debug mode
```

#### Database Configuration

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

#### Redis Configuration

```toml
[redis.default]
address = "127.0.0.1:6379"
db = 4
pass = "your_redis_password"
```

#### JWT Configuration

```toml
[jwt]
exp = 72000                                    # Token validity period (seconds)
pre = "laughstorage_staking_tool_"            # Token prefix in Redis
issuer = "laughstorage_staking_tool_"         # Token issuer
key = "your_jwt_secret_key"                   # JWT secret key
encryptKey = "your_encryption_key"            # Encryption key
```

#### SIWE (Sign-In with Ethereum) Configuration

```toml
[siwe]
domain = "localhost:9000"  # Web3 login verification domain, must match the frontend
```

#### Alibaba Cloud Configuration (KYC Authentication)

```toml
[ali]
accessKeyId = "your_aliyun_access_key_id"
accessKeySecret = "your_aliyun_access_key_secret"
ossBucketName = "your_oss_bucket_name"
ossUrl = "your_oss_url"
region = "cn-hongkong"
```

#### Application Configuration

```toml
[app]
administrator = "0x..."                    # Administrator wallet address (has proposal review permissions, multiple addresses separated by commas)
rpc = "https://filecoin-calibration.chainup.net/rpc/v1"  # Filecoin RPC address
stakingContract = "0x..."                  # Staking contract address
```

#### Allocator Configuration

```toml
[allocator]
contract = ""   # Allocator contract address
account  = ""   # Notary address
accountPrivateKey = ""  # Notary private key
```

### Important Configuration Notes

1. **administrator**: This is the most important configuration item, specifying the wallet address that has proposal review permissions
2. **domain**: The domain in SIWE configuration must exactly match the frontend application domain
3. **Database**: Ensure database connection information is correct and `script.sql` file has been imported
4. **Redis**: Used for caching and session management, ensure connection is normal

## 🛠️ Development Guide

### Project Structure

```
api/
├── api/                    # API route definitions
├── internal/              # Internal packages
│   ├── cmd/              # Command line entry
│   ├── controller/       # Controller layer
│   ├── dao/             # Data access layer
│   ├── logic/           # Business logic layer
│   ├── middleware/      # Middleware
│   ├── model/           # Data models
│   └── service/         # Service layer
├── resource/            # Static resources
├── manifest/           # Deployment configuration
└── main.go            # Program entry
```

### Common Development Commands

```bash
# Generate controller code
gf gen ctrl

# Generate data access layer code
gf gen dao

# Build project
gf build
```

### Code Generation

The project uses GoFrame's code generation functionality:

```bash
# Generate controllers based on API definitions
gf gen ctrl

# Generate DAO/DO/Entity based on database tables
gf gen dao
```

## 🚀 Deployment Guide

### Production Environment Configuration

1. Modify production environment configuration in `config.toml`
2. Ensure stable database and Redis connections
3. Configure reverse proxy (such as Nginx)
4. Set up SSL certificates
5. Configure log rotation

## 🔧 FAQ

### Q: Database connection fails when starting

A: Check database configuration in `config.toml`, ensure:

- Database service is running
- Connection information is correct
- Database has been created and `script.sql` has been imported

### Q: JWT Token validation fails

A: Check JWT configuration:

- Ensure `key` and `encryptKey` are configured correctly
- Check if Token has expired
- Verify Redis connection is normal

### Q: KYC authentication fails

A: Check Alibaba Cloud configuration:

- Ensure AccessKey is valid
- Check OSS Bucket permissions
- Verify region configuration is correct

### Q: Web3 login fails

A: Check SIWE configuration:

- Ensure `domain` matches frontend domain
- Check network connection
- Verify wallet connection is normal

## 📝 Logs

Log files are located in the `log/` directory and are automatically split by date. You can adjust log level and output format by modifying log configuration in `config.toml`.

## 🤝 Contributing

1. Fork the project
2. Create a feature branch (`git checkout -b feature/AmazingFeature`)
3. Commit your changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request

## 📄 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## 📞 Support

If you encounter issues or have any questions, please:

1. Check the [FAQ](#faq) section
2. Review project [Issues](https://github.com/your-repo/issues)
3. Create a new Issue describing your problem

---

**Note**: Before deploying to production environment, please ensure:

- All sensitive information has been properly configured
- Database has been backed up
- Monitoring and logging systems have been configured
- Security measures have been implemented
