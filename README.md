# Go Inventory System 📦

A production-ready Go backend inventory management system with RESTful API, featuring authentication, CRUD operations, and complete DevOps setup. Built with clean architecture principles and modern best practices.

[![CI/CD Pipeline](https://github.com/nielwyn/go-inventory-system/actions/workflows/ci.yml/badge.svg)](https://github.com/nielwyn/go-inventory-system/actions/workflows/ci.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/nielwyn/go-inventory-system)](https://goreportcard.com/report/github.com/nielwyn/go-inventory-system)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

## ✨ Features

- **🔐 Authentication & Authorization**: JWT-based authentication with bcrypt password hashing
- **📊 CRUD Operations**: Complete inventory management with Create, Read, Update, Delete operations
- **🏗️ Clean Architecture**: Separation of concerns with handlers, services, repositories, and models
- **🗄️ PostgreSQL Database**: GORM ORM with auto-migrations and connection pooling
- **📝 Structured Logging**: JSON-formatted logs using Uber's Zap logger
- **🔍 Input Validation**: Custom validators for business rules
- **🐳 Docker Support**: Multi-stage Dockerfiles and Docker Compose setup
- **🚀 CI/CD Pipeline**: GitHub Actions workflow for automated testing and deployment
- **📊 Metrics**: Prometheus metrics endpoint for monitoring
- **🏥 Health Checks**: Health and readiness endpoints for Kubernetes/orchestration
- **⚡ Graceful Shutdown**: Proper signal handling and connection cleanup

## 🛠️ Technology Stack

- **Framework**: [Gin](https://github.com/gin-gonic/gin) v1.9.1 - High-performance HTTP web framework
- **ORM**: [GORM](https://gorm.io/) v1.25.5 - Developer-friendly ORM library
- **Database Driver**: [PostgreSQL](https://www.postgresql.org/) via gorm.io/driver/postgres v1.5.4
- **Authentication**: [JWT](https://github.com/golang-jwt/jwt) v5.2.0 - JSON Web Tokens
- **Logging**: [Zap](https://github.com/uber-go/zap) v1.26.0 - Structured, leveled logging
- **Configuration**: [godotenv](https://github.com/joho/godotenv) v1.5.1 - Environment variable management
- **Metrics**: [Prometheus](https://github.com/prometheus/client_golang) v1.18.0
- **Security**: [bcrypt](https://pkg.go.dev/golang.org/x/crypto/bcrypt) - Password hashing

## 📁 Project Structure

```
go-inventory-system/
├── cmd/
│   └── api/
│       └── main.go                    # Application entry point
├── internal/
│   ├── handlers/                      # HTTP request handlers
│   │   ├── health.go                  # Health check endpoints
│   │   ├── auth.go                    # Authentication endpoints
│   │   └── inventory.go               # Inventory CRUD endpoints
│   ├── models/                        # Data models
│   │   ├── user.go                    # User model and DTOs
│   │   └── item.go                    # Item model and DTOs
│   ├── database/
│   │   └── database.go                # Database connection and setup
│   ├── middleware/                    # HTTP middleware
│   │   ├── logger.go                  # Request logging
│   │   ├── auth.go                    # JWT authentication
│   │   └── cors.go                    # CORS handling
│   ├── repository/                    # Data access layer
│   │   ├── user_repository.go         # User data operations
│   │   └── inventory_repository.go    # Inventory data operations
│   └── service/                       # Business logic layer
│       ├── auth_service.go            # Authentication logic
│       └── inventory_service.go       # Inventory business logic
├── pkg/                               # Reusable packages
│   ├── logger/
│   │   └── logger.go                  # Logger initialization
│   ├── validator/
│   │   └── validator.go               # Custom validators
│   └── response/
│       └── response.go                # Standard API responses
├── config/
│   └── config.go                      # Configuration management
├── migrations/
│   └── 001_initial_schema.sql         # Database schema (reference)
├── scripts/
│   ├── setup.sh                       # Setup script
│   └── seed.sql                       # Sample data seeding
├── deployments/
│   ├── docker/
│   │   ├── Dockerfile                 # Basic Docker image
│   │   └── Dockerfile.multistage      # Optimized multi-stage build
│   └── docker-compose.yml             # Docker Compose configuration
├── .github/
│   └── workflows/
│       └── ci.yml                     # GitHub Actions CI/CD
├── tests/
│   └── integration/                   # Integration tests
├── .env.example                       # Environment variables template
├── .gitignore                         # Git ignore rules
├── .dockerignore                      # Docker ignore rules
├── Makefile                           # Development commands
├── go.mod                             # Go module dependencies
├── README.md                          # This file
└── LICENSE                            # MIT License
```

## 🚀 Quick Start

### Prerequisites

- Go 1.21 or higher
- PostgreSQL 12 or higher
- Docker and Docker Compose (optional, for containerized deployment)

### Local Development Setup

1. **Clone the repository**
   ```bash
   git clone https://github.com/nielwyn/go-inventory-system.git
   cd go-inventory-system
   ```

2. **Set up environment variables**
   ```bash
   cp .env.example .env
   # Edit .env with your configuration
   ```

3. **Install dependencies**
   ```bash
   make deps
   ```

4. **Run the setup script**
   ```bash
   ./scripts/setup.sh
   ```

5. **Start PostgreSQL** (if not using Docker)
   ```bash
   # Example using PostgreSQL service
   sudo service postgresql start
   ```

6. **Run the application**
   ```bash
   make run
   ```

The API will be available at `http://localhost:8080`

### Docker Deployment

**Option 1: Using Docker Compose (Recommended)**
```bash
# Copy environment file
cp .env.example .env

# Start all services (PostgreSQL + API)
make docker-run

# View logs
make docker-logs

# Stop services
make docker-down
```

**Option 2: Build and Run Docker Image Manually**
```bash
# Build Docker image
make docker-build

# Run with external PostgreSQL
docker run -p 8080:8080 \
  -e DB_HOST=your-postgres-host \
  -e DB_PASSWORD=your-password \
  -e JWT_SECRET=your-secret \
  go-inventory-system:latest
```

## 📚 API Documentation

### Base URL
```
http://localhost:8080
```

### Response Format

**Success Response:**
```json
{
  "success": true,
  "message": "Operation successful",
  "data": { ... }
}
```

**Error Response:**
```json
{
  "success": false,
  "message": "Error description"
}
```

### Endpoints

#### Health & Monitoring

| Method | Endpoint   | Description              | Auth Required |
|--------|-----------|--------------------------|---------------|
| GET    | /health   | Basic health check       | No            |
| GET    | /ready    | Readiness check with DB  | No            |
| GET    | /metrics  | Prometheus metrics       | No            |

**Example:**
```bash
curl http://localhost:8080/health
```

#### Authentication

| Method | Endpoint              | Description           | Auth Required |
|--------|-----------------------|----------------------|---------------|
| POST   | /api/v1/auth/register | Register new user    | No            |
| POST   | /api/v1/auth/login    | Login and get token  | No            |

**Register User:**
```bash
curl -X POST http://localhost:8080/api/v1/auth/register \
  -H "Content-Type: application/json" \
  -d '{
    "username": "johndoe",
    "email": "john@example.com",
    "password": "securepassword123"
  }'
```

**Login:**
```bash
curl -X POST http://localhost:8080/api/v1/auth/login \
  -H "Content-Type: application/json" \
  -d '{
    "username": "johndoe",
    "password": "securepassword123"
  }'
```

**Response:**
```json
{
  "success": true,
  "message": "Login successful",
  "data": {
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...",
    "user": {
      "id": 1,
      "username": "johndoe",
      "email": "john@example.com",
      "created_at": "2026-01-30T10:00:00Z",
      "updated_at": "2026-01-30T10:00:00Z"
    }
  }
}
```

#### Inventory Management (Protected)

All inventory endpoints require JWT authentication. Include the token in the Authorization header:
```
Authorization: Bearer <your-jwt-token>
```

| Method | Endpoint                      | Description        | Auth Required |
|--------|-------------------------------|-------------------|---------------|
| POST   | /api/v1/inventory/items       | Create new item   | Yes           |
| GET    | /api/v1/inventory/items       | Get all items     | Yes           |
| GET    | /api/v1/inventory/items/:id   | Get item by ID    | Yes           |
| PUT    | /api/v1/inventory/items/:id   | Update item       | Yes           |
| DELETE | /api/v1/inventory/items/:id   | Delete item       | Yes           |

**Create Item:**
```bash
curl -X POST http://localhost:8080/api/v1/inventory/items \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer <your-jwt-token>" \
  -d '{
    "name": "Laptop - Dell XPS 15",
    "sku": "LAPTOP-XPS15-001",
    "description": "High-performance laptop",
    "quantity": 25,
    "price": 1299.99,
    "category": "Electronics"
  }'
```

**Get All Items:**
```bash
curl http://localhost:8080/api/v1/inventory/items \
  -H "Authorization: Bearer <your-jwt-token>"
```

**Get Item by ID:**
```bash
curl http://localhost:8080/api/v1/inventory/items/1 \
  -H "Authorization: Bearer <your-jwt-token>"
```

**Update Item:**
```bash
curl -X PUT http://localhost:8080/api/v1/inventory/items/1 \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer <your-jwt-token>" \
  -d '{
    "quantity": 30,
    "price": 1199.99
  }'
```

**Delete Item:**
```bash
curl -X DELETE http://localhost:8080/api/v1/inventory/items/1 \
  -H "Authorization: Bearer <your-jwt-token>"
```

## ⚙️ Configuration

### Environment Variables

| Variable          | Description                    | Default        | Required |
|-------------------|--------------------------------|----------------|----------|
| SERVER_HOST       | Server host address            | 0.0.0.0        | No       |
| SERVER_PORT       | Server port                    | 8080           | No       |
| GIN_MODE          | Gin mode (debug/release)       | debug          | No       |
| DB_HOST           | PostgreSQL host                | localhost      | Yes      |
| DB_PORT           | PostgreSQL port                | 5432           | No       |
| DB_USER           | Database user                  | postgres       | Yes      |
| DB_PASSWORD       | Database password              | postgres       | Yes      |
| DB_NAME           | Database name                  | inventory_db   | Yes      |
| DB_SSLMODE        | PostgreSQL SSL mode            | disable        | No       |
| JWT_SECRET        | JWT signing secret             | -              | Yes      |
| JWT_EXPIRY_HOURS  | JWT token expiry in hours      | 24             | No       |
| LOG_LEVEL         | Log level (debug/info/error)   | debug          | No       |
| LOG_ENCODING      | Log encoding (json/console)    | json           | No       |

## 🧪 Development

### Available Commands

```bash
make help              # Show all available commands
make deps              # Download dependencies
make run               # Run the application locally
make build             # Build binary
make test              # Run tests with coverage
make clean             # Clean build artifacts
make docker-build      # Build Docker image
make docker-run        # Run with Docker Compose
make docker-down       # Stop Docker Compose
make docker-logs       # View Docker logs
make lint              # Run linter
make fmt               # Format code
make vet               # Run go vet
make setup             # Run setup script
make seed              # Seed database with sample data
make all               # Run all checks and build
```

### Database Seeding

To populate the database with sample inventory items:

```bash
# Ensure PostgreSQL is running
make seed
```

Or manually:
```bash
psql -h localhost -U postgres -d inventory_db -f scripts/seed.sql
```

## 🔒 Security Features

- **Password Hashing**: bcrypt with default cost factor
- **JWT Authentication**: Secure token-based authentication
- **SQL Injection Prevention**: GORM parameterized queries
- **CORS**: Configurable cross-origin resource sharing
- **Environment-based Secrets**: Sensitive data in environment variables
- **Input Validation**: Request validation with custom business rules
- **Soft Deletes**: Data integrity with GORM soft delete

## 📊 Monitoring & Observability

- **Structured Logging**: JSON-formatted logs with request context
- **Prometheus Metrics**: `/metrics` endpoint for monitoring
- **Health Checks**: `/health` and `/ready` endpoints for orchestration
- **Request Logging**: Automatic logging of all HTTP requests with latency

## 🐛 Troubleshooting

### Database Connection Issues

```bash
# Check PostgreSQL is running
sudo service postgresql status

# Check connection
psql -h localhost -U postgres -d inventory_db -c "SELECT 1"

# View logs
tail -f *.log
```

### Port Already in Use

```bash
# Find process using port 8080
lsof -i :8080

# Kill the process
kill -9 <PID>
```

### Docker Issues

```bash
# Clean Docker resources
docker-compose -f deployments/docker-compose.yml down -v
docker system prune -a
```

## 🤝 Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/AmazingFeature`)
3. Commit your changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request

## 📝 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## 👤 Author

**nielwyn**

- GitHub: [@nielwyn](https://github.com/nielwyn)

## 🙏 Acknowledgments

- Built with [Gin](https://github.com/gin-gonic/gin) web framework
- Database management with [GORM](https://gorm.io/)
- Logging powered by [Zap](https://github.com/uber-go/zap)
- Inspired by clean architecture principles

---

**⭐ If you find this project helpful, please consider giving it a star!**
