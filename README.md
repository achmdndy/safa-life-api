# Safa Life REST API 🌙

**Safa Life** is a modern Islamic lifestyle REST API designed to help Muslims live a life guided by the Qur'an and Sunnah. Built with Go using Clean Architecture principles, this robust backend provides comprehensive Islamic services through a scalable and maintainable codebase.

## ✨ Features

- 🏥 **Health Check API**: Application health monitoring endpoints
- 📊 **Prometheus Metrics**: Built-in metrics collection for monitoring
- 🔍 **Jaeger Tracing**: Distributed tracing for request flow analysis
- 🔐 **Clean Architecture**: Domain-driven design with clear separation of concerns
- 🚀 **High Performance**: Built with Gin framework for optimal performance
- 📦 **Container Ready**: Docker support for easy deployment
- 🧪 **Test Coverage**: Comprehensive testing setup

## 🛠️ Tech Stack

- **Framework**: [Gin](https://gin-gonic.com/) - High-performance HTTP web framework
- **ORM**: [GORM](https://gorm.io/) - Developer-friendly ORM library
- **Configuration**: [Viper](https://github.com/spf13/viper) - Complete configuration solution
- **CLI**: [Cobra](https://cobra.dev/) - Powerful CLI framework
- **Database**: PostgreSQL with Redis for caching
- **Monitoring**: Prometheus metrics with Jaeger tracing via OpenTelemetry
- **Architecture**: Clean Architecture with Domain-Driven Design

## 🚀 Quick Start

### Prerequisites

- Go 1.21 or higher
- PostgreSQL 14+
- Redis 6+
- Docker & Docker Compose (recommended)

### Using Docker (Recommended)

1. **Clone the repository**
   ```bash
   git clone https://github.com/achmdndy/safa-life-api.git
   cd safa-life-api
   ```

2. **Start all services**
   ```bash
   docker-compose up -d
   ```

3. **Verify the application**
   ```bash
   curl http://localhost:8080/health
   ```

### Manual Installation

1. **Clone and setup**
   ```bash
   git clone https://github.com/achmdndy/safa-life-api.git
   cd safa-life-api
   go mod download
   ```

2. **Configure environment**
   ```bash
   # Edit configs/config.yaml with your database settings
   ```

3. **Start dependencies**
   ```bash
   # Start PostgreSQL and Redis
   docker-compose up -d postgres redis
   ```

4. **Run the application**
   ```bash
   go run src/cmd/server/main.go start
   ```

## 📁 Project Structure

```
safa-life-api/
├── src/
│   ├── application/           # Application layer
│   │   ├── container/        # Application container
│   │   └── health/           # Health use cases and interfaces
│   ├── cmd/                  # Application entry points
│   │   ├── commands/         # CLI commands
│   │   ├── core/            # Core configuration
│   │   └── server/          # HTTP server entry point
│   ├── domain/              # Domain layer (business logic)
│   │   ├── health/          # Health domain entities
│   │   └── shared/          # Shared domain interfaces
│   ├── infrastructure/      # Infrastructure layer
│   │   ├── configs/         # Configuration implementations
│   │   ├── container/       # Infrastructure container
│   │   ├── health/          # Health repository implementations
│   │   └── monitoring/      # Monitoring setup (Prometheus, Jaeger)
│   └── presentation/        # Presentation layer
│       ├── container/       # Presentation container
│       ├── core/           # Response utilities
│       ├── handlers/       # HTTP handlers
│       ├── middlewares/    # HTTP middlewares
│       └── routes/         # Route definitions
├── configs/                 # Configuration files
├── bin/                    # Compiled binaries
├── docs/                   # Documentation
└── docker-compose.yml     # Complete development setup
```

## 🔧 Configuration

The application uses a layered configuration system:

```yaml
# configs/config.yaml
app_name: "Safa Life API"
port: 8080

db:
  host: "localhost"
  user: "postgres"
  password: "password"
  name: "safa_life"
  port: 5432
  ssl_mode: "disable"
  timezone: "UTC"

redis:
  host: "localhost"
  port: 6379
  password: ""
  db: 0

server:
  host: "0.0.0.0"
  port: "8080"
  env: "development"
```

Environment variables can override configuration values using the pattern: `SAFA_SECTION_KEY`

## 📚 API Documentation

### Base URL
```
Development: http://localhost:8080
Production: Not available yet
```

### Health Endpoints
```bash
# Application health check
GET /health

# Root endpoint with API information
GET /
```

### Monitoring Endpoints
```bash
# Prometheus metrics
GET /metrics

# Health check for monitoring
GET /v1/health
```

## 📊 Monitoring & Observability

### Prometheus Metrics
- **HTTP Request Metrics**: Total requests, duration, status codes
- **Application Info**: Version and service information
- **Go Runtime Metrics**: Memory, GC, goroutines

### Jaeger Tracing
- **Distributed Tracing**: Request flow across services
- **OpenTelemetry Integration**: Standard observability
- **Automatic Instrumentation**: HTTP request tracing

### Access Monitoring Tools
- **Jaeger UI**: http://localhost:16686
- **Prometheus**: http://localhost:9090
- **Grafana**: http://localhost:3000 (admin/admin)

## 🧪 Testing

```bash
# Run all tests
go test ./...

# Run tests with coverage
go test -cover ./...

# Run specific package tests
go test ./src/presentation/handlers/health

# Build and test
go build -o bin/app ./src/cmd/server/main.go
./bin/app start
```

## 🏗️ Architecture

This project follows **Clean Architecture** principles:

### Domain Layer (`src/domain/`)
- Contains business entities and rules
- Independent of external frameworks
- Defines repository interfaces

### Application Layer (`src/application/`)
- Contains use cases and application logic
- Orchestrates domain objects
- Defines application interfaces

### Infrastructure Layer (`src/infrastructure/`)
- Implements external concerns
- Database connections and repositories
- Monitoring and configuration setup

### Presentation Layer (`src/presentation/`)
- HTTP handlers and middleware
- Request/response formatting
- Route definitions

## 🚀 Deployment

### Environment Variables

| Variable | Description | Default |
|----------|-------------|---------|
| `SAFA_SERVER_PORT` | HTTP server port | `8080` |
| `SAFA_DB_HOST` | PostgreSQL host | `localhost` |
| `SAFA_DB_USER` | Database user | `postgres` |
| `SAFA_DB_PASSWORD` | Database password | `password` |
| `SAFA_DB_NAME` | Database name | `safa_life` |
| `SAFA_REDIS_HOST` | Redis host | `localhost` |

### Docker Deployment

```bash
# Build image
docker build -t safa-life-api .

# Run with docker-compose
docker-compose up -d

# Scale services
docker-compose up -d --scale api=3
```

### Production Checklist

- [ ] Set strong database passwords
- [ ] Configure SSL/TLS certificates
- [ ] Set up proper logging levels
- [ ] Configure monitoring alerts
- [ ] Set up backup strategies
- [ ] Configure rate limiting

## 🤝 Contributing

We welcome contributions! Please see our [Contributing Guide](CONTRIBUTING.md) for details.

### Development Guidelines

- Follow [Go Code Review Comments](https://github.com/golang/go/wiki/CodeReviewComments)
- Maintain Clean Architecture principles
- Write tests for new features
- Update documentation for API changes
- Use conventional commits

### Code of Conduct

Please read our [Code of Conduct](CODE_OF_CONDUCT.md) before contributing.

## 📄 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## 🙏 Acknowledgments

- [Gin Framework](https://gin-gonic.com/) for the excellent HTTP framework
- [GORM](https://gorm.io/) for the powerful ORM
- [OpenTelemetry](https://opentelemetry.io/) for observability standards
- [Prometheus](https://prometheus.io/) for metrics collection
- All contributors who help make Safa Life better

## 📞 Support

- 📧 Email: achmdndy@gmail.com
- 🐛 Issues: [GitHub Issues](https://github.com/achmdndy/safa-life-api/issues)
- 📖 Documentation: [Project Wiki](https://github.com/achmdndy/safa-life-api/wiki)

---

**✨ Built with Clean Architecture principles to create maintainable and scalable Islamic lifestyle applications.**