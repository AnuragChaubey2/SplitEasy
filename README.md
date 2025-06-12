project-name/                 ← e.g., splitly-backend or settlr-core
│
├── cmd/                      ← Entry points (main.go for now)
│   └── main.go
│
├── internal/                 ← Business logic (split by domain/module)
│   ├── auth/                 ← Register, login, JWT generation
│   ├── user/                 ← User profile, fetching
│   ├── group/                ← Create group, add/remove members
│   ├── expense/              ← Add/view expenses, split logic
│   ├── settlement/           ← Manual debt settlement logic
│   ├── balance/              ← Debt calculations, who owes whom
│   ├── receipt/              ← File uploads and MinIO/S3 integration
│   ├── middleware/           ← JWT auth middleware, request logging
│   ├── config/               ← Viper config loader, env vars
│   ├── router/               ← Gin route registration (grouped by module)
│   ├── database/             ← DB init, connection, migrations
│   ├── models/               ← GORM structs, schema, shared types
│   └── utils/                ← Common helpers (error handling, UUID, etc.)
│
├── pkg/                      ← Public packages (if you want to reuse)
│   └── fxrates/              ← (optional) FX rates utils in future
│
├── scripts/                  ← Setup scripts (e.g., DB init, local tooling)
│
├── .env                      ← Environment variables for dev
├── .gitignore
├── go.mod
├── go.sum
├── Dockerfile
├── docker-compose.yml       ← (if needed for local dev: Postgres, MinIO)
├── README.md
└── openapi.yaml             ← Swagger/OpenAPI spec (if using Swagger)




PHASE 1: Initial Setup and Bootstrapping
1. Project Initialization
Initialize Go module

go mod init github.com/yourname/splitly-backend

Create project folder structure:

cmd/main.go

internal/, config/, router/, etc.

Create .env and add essential variables:

DATABASE_URL

JWT_SECRET

MINIO_ACCESS_KEY, etc.

2. Setup Configuration Loader
Create internal/config/config.go

Use viper to read from .env

Load and validate basic config (DB, JWT, MinIO)

3. Setup Database Connection
Create internal/database/db.go

Use GORM to connect to PostgreSQL

Auto-migration support

Define models in internal/models/

User, Group, GroupMember, Expense, etc.

4. Setup Project Entry Point
Create cmd/main.go

Initialize config

Connect to DB

Initialize router

Start HTTP server

PHASE 2: Authentication System
1. Build Auth Module
Create internal/auth/

handler.go: Register, Login endpoints

service.go: Business logic (hash password, verify)

repository.go: Create user, get user by email

Use JWT for token generation

Store secret in .env

2. Add Middleware
internal/middleware/auth.go: JWT parsing and user injection into context

Add global middleware registration in internal/router/router.go

PHASE 3: User and Group Management
1. User Profile
Add /me route to return user info from JWT

2. Group Module
internal/group/:

handler.go: Create group, invite members, get group info

repository.go: Manage group and member DB ops

Models:

Group, GroupMember

PHASE 4: Expense Handling
1. Expense Module
internal/expense/:

handler.go: POST and GET expenses for a group

service.go: Core logic to split expenses (equal only in MVP)

repository.go: CRUD operations

2. ExpenseParticipant Table
Store individual shares per user

Logic: on expense creation

Calculate shares

Update balances (on-the-fly for now, not persisted)

PHASE 5: Balance Calculation
1. Balance Module
internal/balance/:

service.go: Compute who owes who how much

Read from expenses and settlements

No persisted balances in MVP, only calculated dynamically

2. Balance API
/groups/:id/balances:

Show list of net balances per user in group