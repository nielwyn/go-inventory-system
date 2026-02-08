# 📚 Go Inventory System - Learning Guide

A structured learning path to understand this production-ready Go backend project. Follow this order for the best learning experience.

---

## 🎯 Learning Path Overview

```
┌─────────────────────────────────────────────────────────────────────┐
│                    PHASE 1: FOUNDATION                              │
│  [1] Project Setup → [2] Configuration → [3] Utilities (pkg/)       │
└─────────────────────────────────────────────────────────────────────┘
                              ↓
┌─────────────────────────────────────────────────────────────────────┐
│                    PHASE 2: DATA LAYER                              │
│  [4] Models → [5] Database Connection → [6] Repositories            │
└─────────────────────────────────────────────────────────────────────┘
                              ↓
┌─────────────────────────────────────────────────────────────────────┐
│                    PHASE 3: BUSINESS LOGIC                          │
│  [7] Services (Auth & Inventory)                                    │
└─────────────────────────────────────────────────────────────────────┘
                              ↓
┌─────────────────────────────────────────────────────────────────────┐
│                    PHASE 4: API LAYER                               │
│  [8] Middleware → [9] Handlers → [10] Main Entry Point              │
└─────────────────────────────────────────────────────────────────────┘
                              ↓
┌─────────────────────────────────────────────────────────────────────┐
│                    PHASE 5: DEVOPS                                  │
│  [11] Docker → [12] Makefile → [13] CI/CD                           │
└─────────────────────────────────────────────────────────────────────┘
```

---

## 📖 Detailed Learning Sections

### Phase 1: Foundation (Start Here!)

#### 1️⃣ Project Structure & Setup
**Files:** `go.mod`, `README.md`, `Makefile`

**What you'll learn:**
- Go modules and dependency management
- Project organization patterns
- Available development commands

**Key concepts:**
- `go.mod` - Defines module path and dependencies
- Clean Architecture folder structure (`cmd/`, `internal/`, `pkg/`)
- Makefile automation for common tasks

**Study order:**
1. [go.mod](go.mod) - See all dependencies used
2. [README.md](README.md) - Understand project features
3. [Makefile](Makefile) - Learn available commands

---

#### 2️⃣ Configuration Management
**File:** [config/config.go](config/config.go)

**What you'll learn:**
- Environment variable management with `godotenv`
- Type-safe configuration structs
- Configuration validation

**Key concepts:**
- Struct-based configuration (`Config`, `ServerConfig`, `DatabaseConfig`)
- Environment variable loading with defaults
- DSN (Data Source Name) generation for PostgreSQL

**Practice:** Try modifying default values and adding new config options.

---

#### 3️⃣ Utility Packages (pkg/)
**Files:** `pkg/logger/`, `pkg/response/`, `pkg/validator/`

**What you'll learn:**
- Reusable package design
- Structured logging with Zap
- Standardized API responses
- Custom validation rules

**Study order:**
1. [pkg/logger/logger.go](pkg/logger/logger.go) - Uber Zap logger setup
2. [pkg/response/response.go](pkg/response/response.go) - Standard API response format
3. [pkg/validator/validator.go](pkg/validator/validator.go) - Custom validation (e.g., `non_negative`)

**Key concepts:**
- Global logger pattern
- Consistent response structure `{success, message, data}`
- Custom validators with `go-playground/validator`

---

### Phase 2: Data Layer

#### 4️⃣ Data Models
**Files:** [internal/models/user.go](internal/models/user.go), [internal/models/item.go](internal/models/item.go)

**What you'll learn:**
- GORM model definitions
- Struct tags for JSON and DB mapping
- Request/Response DTOs (Data Transfer Objects)
- Validation annotations

**Key concepts:**
```go
// GORM model with JSON serialization
type Item struct {
    ID          uint           `gorm:"primaryKey" json:"id"`
    Name        string         `gorm:"not null" json:"name"`
    DeletedAt   gorm.DeletedAt `gorm:"index" json:"-"` // Soft delete
}

// Request DTO with validation
type CreateItemRequest struct {
    Name     string  `json:"name" binding:"required,min=1,max=200"`
    Quantity int     `json:"quantity" binding:"non_negative"`
}
```

**Study order:**
1. [internal/models/user.go](internal/models/user.go) - User model & auth DTOs
2. [internal/models/item.go](internal/models/item.go) - Item model & CRUD DTOs

---

#### 5️⃣ Database Connection
**File:** [internal/database/database.go](internal/database/database.go)

**What you'll learn:**
- GORM PostgreSQL connection setup
- Connection pooling configuration
- Auto-migration
- Health check (Ping)

**Key concepts:**
- `gorm.Open()` for database connection
- Connection pool settings (`SetMaxIdleConns`, `SetMaxOpenConns`)
- `AutoMigrate()` for schema management

---

#### 6️⃣ Repository Pattern
**Files:** [internal/repository/user_repository.go](internal/repository/user_repository.go), [internal/repository/inventory_repository.go](internal/repository/inventory_repository.go)

**What you'll learn:**
- Repository pattern (data access abstraction)
- Interface-based design
- CRUD operations with GORM
- Soft deletes

**Key concepts:**
```go
// Interface for dependency injection
type UserRepository interface {
    Create(user *models.User) error
    FindByUsername(username string) (*models.User, error)
    FindByID(id uint) (*models.User, error)
}

// Implementation
type userRepository struct {
    db *gorm.DB
}
```

**Study order:**
1. [internal/repository/user_repository.go](internal/repository/user_repository.go) - User CRUD
2. [internal/repository/inventory_repository.go](internal/repository/inventory_repository.go) - Item CRUD

---

### Phase 3: Business Logic

#### 7️⃣ Service Layer
**Files:** [internal/service/auth_service.go](internal/service/auth_service.go), [internal/service/inventory_service.go](internal/service/inventory_service.go)

**What you'll learn:**
- Business logic encapsulation
- JWT token generation & validation
- Password hashing with bcrypt
- Service interface pattern

**Key concepts:**
```go
// Auth Service responsibilities:
- Register: Validate uniqueness → Hash password → Create user
- Login: Find user → Verify password → Generate JWT
- ValidateToken: Parse & verify JWT signature
```

**Study order:**
1. [internal/service/auth_service.go](internal/service/auth_service.go) - Authentication logic
2. [internal/service/inventory_service.go](internal/service/inventory_service.go) - Business rules for items

**Important patterns:**
- Services depend on Repository interfaces (not implementations)
- Business validation happens here (e.g., "SKU already exists")

---

### Phase 4: API Layer

#### 8️⃣ Middleware
**Files:** [internal/middleware/auth.go](internal/middleware/auth.go), [internal/middleware/logger.go](internal/middleware/logger.go), [internal/middleware/cors.go](internal/middleware/cors.go)

**What you'll learn:**
- Gin middleware pattern
- JWT authentication middleware
- Request logging
- CORS configuration

**Key concepts:**
```go
// Middleware signature
func Auth(authService service.AuthService) gin.HandlerFunc {
    return func(c *gin.Context) {
        // 1. Extract token from header
        // 2. Validate token
        // 3. Set user_id in context
        c.Set("user_id", userID)
        c.Next()
    }
}
```

**Study order:**
1. [internal/middleware/auth.go](internal/middleware/auth.go) - JWT validation
2. [internal/middleware/logger.go](internal/middleware/logger.go) - Request logging
3. [internal/middleware/cors.go](internal/middleware/cors.go) - CORS handling

---

#### 9️⃣ HTTP Handlers
**Files:** [internal/handlers/auth.go](internal/handlers/auth.go), [internal/handlers/inventory.go](internal/handlers/inventory.go), [internal/handlers/health.go](internal/handlers/health.go)

**What you'll learn:**
- Gin handler functions
- Request parsing & validation
- Response formatting
- Error handling

**Key concepts:**
```go
func (h *InventoryHandler) CreateItem(c *gin.Context) {
    // 1. Bind JSON request
    var req models.CreateItemRequest
    if err := c.ShouldBindJSON(&req); err != nil {
        response.Error(c, http.StatusBadRequest, ...)
        return
    }
    
    // 2. Call service
    item, err := h.inventoryService.CreateItem(&req)
    
    // 3. Return response
    response.Success(c, http.StatusCreated, "Item created", item)
}
```

**Study order:**
1. [internal/handlers/health.go](internal/handlers/health.go) - Simple health checks
2. [internal/handlers/auth.go](internal/handlers/auth.go) - Register & Login endpoints
3. [internal/handlers/inventory.go](internal/handlers/inventory.go) - Full CRUD handlers

---

#### 🔟 Main Entry Point (Everything Comes Together!)
**File:** [cmd/api/main.go](cmd/api/main.go)

**What you'll learn:**
- Application bootstrapping
- Dependency injection wiring
- Router setup with Gin
- Graceful shutdown

**Key flow:**
```
1. Load Config
2. Initialize Logger
3. Connect Database
4. Run Migrations
5. Create Repositories → Services → Handlers
6. Setup Router (routes + middleware)
7. Start HTTP Server
8. Handle graceful shutdown (SIGINT/SIGTERM)
```

**This is the best file to understand how everything connects!**

---

### Phase 5: DevOps

#### 1️⃣1️⃣ Docker
**Files:** [deployments/docker/Dockerfile](deployments/docker/Dockerfile), [deployments/docker/Dockerfile.multistage](deployments/docker/Dockerfile.multistage), [deployments/docker-compose.yml](deployments/docker-compose.yml)

**What you'll learn:**
- Containerizing Go applications
- Multi-stage builds for smaller images
- Docker Compose for local development

**Study order:**
1. `Dockerfile` - Basic container setup
2. `Dockerfile.multistage` - Optimized production build
3. `docker-compose.yml` - Full stack with PostgreSQL

---

#### 1️⃣2️⃣ Makefile
**File:** [Makefile](Makefile)

**Common commands:**
| Command | Description |
|---------|-------------|
| `make run` | Run locally |
| `make build` | Build binary |
| `make test` | Run tests |
| `make docker-run` | Start with Docker Compose |
| `make lint` | Run linters |

---

#### 1️⃣3️⃣ Database Migrations
**Files:** [migrations/001_initial_schema.sql](migrations/001_initial_schema.sql), [scripts/seed.sql](scripts/seed.sql)

**What you'll learn:**
- SQL schema design
- Seed data for development

---

## 🏗️ Architecture Summary

```
┌──────────────────────────────────────────────────────────────────┐
│                         HTTP Request                              │
└───────────────────────────────┬──────────────────────────────────┘
                                ↓
┌──────────────────────────────────────────────────────────────────┐
│                      Middleware Layer                             │
│              (CORS → Logger → Auth → Handler)                     │
└───────────────────────────────┬──────────────────────────────────┘
                                ↓
┌──────────────────────────────────────────────────────────────────┐
│                       Handler Layer                               │
│           (Parse request, call service, format response)          │
└───────────────────────────────┬──────────────────────────────────┘
                                ↓
┌──────────────────────────────────────────────────────────────────┐
│                       Service Layer                               │
│              (Business logic, validation, JWT)                    │
└───────────────────────────────┬──────────────────────────────────┘
                                ↓
┌──────────────────────────────────────────────────────────────────┐
│                      Repository Layer                             │
│                (Database operations via GORM)                     │
└───────────────────────────────┬──────────────────────────────────┘
                                ↓
┌──────────────────────────────────────────────────────────────────┐
│                        PostgreSQL                                 │
└──────────────────────────────────────────────────────────────────┘
```

---

## ✅ Learning Checklist

### Phase 1: Foundation
- [ ] Understand Go modules (`go.mod`)
- [ ] Know the project folder structure
- [ ] Can use Makefile commands
- [ ] Understand config loading

### Phase 2: Data Layer
- [ ] Can define GORM models
- [ ] Understand struct tags (`json`, `gorm`, `binding`)
- [ ] Know how to set up PostgreSQL connection
- [ ] Understand Repository pattern

### Phase 3: Business Logic
- [ ] Understand Service pattern
- [ ] Know how JWT works (generate/validate)
- [ ] Understand bcrypt password hashing

### Phase 4: API Layer
- [ ] Can write Gin middleware
- [ ] Can create HTTP handlers
- [ ] Understand request validation
- [ ] Know how to wire up routes

### Phase 5: DevOps
- [ ] Can build Docker images
- [ ] Can run with Docker Compose
- [ ] Understand multi-stage builds

---

## 🚀 Next Steps After Learning

1. **Add a feature** - Try adding a new entity (e.g., Category, Supplier)
2. **Write tests** - Add unit tests for services and integration tests
3. **Add pagination** - Implement pagination for `GetAllItems`
4. **Add search** - Implement search/filter for inventory
5. **Add rate limiting** - Protect API from abuse
6. **Deploy** - Try deploying to a cloud provider

---

Happy Learning! 🎉
