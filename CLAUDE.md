# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Build & Test Commands

```bash
go build ./...                    # Build all packages
go test ./...                     # Run all tests
go test ./framework/...           # Run framework tests only
go test ./framework -run TestBind # Run a single test
go test -v ./...                  # Verbose output
```

No Makefile present. No linting tools configured. Go 1.24+ required.

## Architecture

Gob is a Go rapid-development framework built on a **Service Container** (IoC/DI) pattern.

### Core Abstractions (in `framework/`)

- **Container** (`container.go`) — `GobContainer` implements dependency injection. Services are bound via `Bind(provider)` and retrieved via `Make(key)` (singleton) or `MakeNew(key, params)` (new instance). Thread-safe with `sync.RWMutex`.

- **ServiceProvider** (`provider.go`) — Interface every service must implement: `Name()`, `IsDefer()`, `Boot(container)`, `Params(container)`, `Register(container)`. `IsDefer()` controls whether instantiation happens at bind-time (false) or on first `Make()` (true).

- **ServiceConfig** (`config.go`) — Interface for structured config registration: `ConfigName()`, `ConfigStruct()`, `Defaults()`, `Validate()`. Providers implement this to auto-register their config section.

### Key Directories

- `framework/contract/` — Interface definitions (contracts) keyed by string constants like `gob:app`, `gob:config`, `gob:id`. Each contract file defines both the key constant and the service interface.
- `framework/provider/` — Concrete implementations grouped by service name (app/, config/, id/). Each provider directory contains: `provider.go` (ServiceProvider impl), `service.go` (service instance), and optionally `config.go` (ServiceConfig impl) + tests.
- `framework/cobra/` — Forked cobra CLI library extended with container integration (`SetContainer`/`GetContainer`) and cron command scheduling (`AddCronCommand`).
- `framework/util/` — Shared utilities (file ops, console, goroutine helpers, zip, syscalls).
- `framework/testing_util.go` — `MockServiceProvider`, `TestContainer` helpers for writing tests.
- `config/dev/` — YAML config files per service. The `.env` file at root controls config mode and environment.

### Config System (koanf-based)

Three modes controlled by `.env` → `configMode`:
- **root** — single `config.yaml` in config folder
- **folder** — separate YAML files per service in config folder
- **deploy** — config folder with env subdirectory (`config/dev/`, `config/test/`, `config/prod/`)

Config priority (low→high): code defaults → main config YAML → sub-config file → environment variables. Supports `env(key)` placeholder syntax in YAML for variable substitution.

### Service Registration Pattern

```go
// 1. Define contract in framework/contract/
const MyKey = "gob:my"
type My interface { ... }

// 2. Implement in framework/provider/my/
type MyProvider struct{}
func (p *MyProvider) Name() string { return contract.MyKey }
func (p *MyProvider) Register(c Container) NewInstance { return NewMyService }

// 3. Bind in main.go
container.Bind(&my.MyProvider{})
```

### Current State (v2 rewrite in progress)

The framework is mid-refactor. Many features (HTTP engine, ORM, distributed cron, CLI commands) are commented out in `main.go` and `gob_command_distributed.go`. Active providers: app, config, id. See `TODO.md` for the full roadmap.
