# Config 子配置实现 - 完成文档

> 本文档梳理 gob V2 配置系统子配置模块的实现状态和使用说明。

---

## 当前进度概览

| 功能项 | 状态 | 文件位置 |
|--------|------|----------|
| ServiceConfig 契约定义 | ✅ 已完成 | `framework/config.go` |
| kSubConfig 字段声明与初始化 | ✅ 已完成 | `framework/provider/config/service.go:37,116` |
| RegisterSubConfig 方法 | ✅ 已完成 | `framework/provider/config/service.go:166-224` |
| GetSubConfig 方法 | ✅ 已完成 | `framework/provider/config/service.go:232-237` |
| 并发保护（读写锁） | ✅ 已完成 | `framework/provider/config/service.go` |
| 配置优先级机制 | ✅ 已完成 | 支持 4 个优先级层级 |
| AppConfig 实现 | ✅ 已完成 | `framework/provider/app/config.go` |
| AppProvider 集成 | ✅ 已完成 | `framework/provider/app/provider.go:29-41` |
| 单元测试覆盖 | ✅ 已完成 | 9 个测试用例通过 |

---

## 架构设计

### ServiceConfig 接口

```go
// framework/config.go
type ServiceConfig interface {
    ConfigName() string                    // 配置名称，如 "app", "log"
    ConfigStruct() interface{}               // 配置结构体
    Defaults() map[string]interface{}        // 默认值
    Validate(config interface{}) error       // 验证配置
}
```

### Config 接口

```go
// framework/contract/config.go
type Config interface {
    GetEnv() *koanf.Koanf
    GetEnvStruct() *contract.ConfigEnvStruct
    GetConfig() *koanf.Koanf
    GetSubConfig(key string) *koanf.Koanf    // 子配置获取
}
```

### ConfigService 结构

```go
// framework/provider/config/service.go
type ConfigService struct {
    kConfig    *koanf.Koanf                // 主配置
    kSubConfig map[string]*koanf.Koanf     // 子配置映射
    rwLock     sync.RWMutex               // 并发保护
    // ... 其他字段
}
```

---

## 配置优先级规则

```
环境变量 > 子配置文件 > 主配置文件 > 代码默认值
```

| 优先级 | 配置源 | 格式说明 |
|--------|--------|----------|
| 1（最高） | 环境变量 | `APP_{KEY}_{FIELD}`，如 `APP_DEBUG=true` |
| 2 | 子配置文件 | `config/app.yaml`（folder/deploy 模式） |
| 3 | 主配置文件 | `config.yaml` 中的 `app.*` 配置 |
| 4（最低） | 代码默认值 | `ServiceConfig.Defaults()` |

---

## 使用示例

### 创建 ServiceConfig

```go
// framework/provider/app/config.go
type AppConfig struct {
    Debug   bool   `koanf:"debug"`
    Version string `koanf:"version"`
}

func (a *AppConfig) ConfigName() string {
    return "app"
}

func (a *AppConfig) Defaults() map[string]interface{} {
    return map[string]interface{}{
        "debug":   false,
        "version": "1.0.0",
    }
}

func (a *AppConfig) Validate(config interface{}) error {
    cfg, ok := config.(*AppConfig)
    if !ok {
        return fmt.Errorf("invalid config type")
    }
    if cfg.Version == "" {
        return fmt.Errorf("version is required")
    }
    return nil
}
```

### 在 Provider 中注册

```go
// framework/provider/app/provider.go
func (provider *AppProvider) Boot(container framework.Container) error {
    configService, err := container.Make(contract.ConfigKey)
    if err == nil {
        if configSvc, ok := configService.(*config.ConfigService); ok {
            return configSvc.RegisterSubConfig(&AppConfig{})
        }
    }
    return nil
}
```

### 使用子配置

```go
// 在服务中获取子配置
config := container.MustMake(contract.ConfigKey).(contract.Config)
subConfig := config.GetSubConfig("app")

if subConfig != nil {
    debug := subConfig.Bool("debug")
    version := subConfig.String("version")
    // ...
}
```

---

## 配置文件示例

### root 模式（单一文件）

```yaml
# config.yaml
app:
  debug: true
  version: "2.0.0"
```

### folder 模式（模块分离 + 覆盖）

```yaml
# config.yaml - 主配置（基础值）
app:
  version: "2.0.0"

# config/app.yaml - 覆盖配置
debug: true  # 只覆盖 debug，version 从主配置读取
```

### deploy 模式（环境特定）

```yaml
# config/deploy_env/prod/app.yaml
debug: false  # 生产环境配置
```

---

## 已修改文件清单

| 文件 | 修改内容 |
|------|----------|
| `framework/config.go` | ServiceConfig 接口定义 |
| `framework/contract/config.go` | 添加 GetSubConfig 到 Config 接口 |
| `framework/provider/config/service.go` | 子配置完整实现（kSubConfig、RegisterSubConfig、GetSubConfig） |
| `framework/provider/app/config.go` | AppConfig 实现 ServiceConfig |
| `framework/provider/app/provider.go` | Boot 方法注册子配置 |
| `framework/provider/config/service_test.go` | 9 个单元测试用例 |
| `framework/provider/app/provider_test.go` | 集成测试更新 |

---

## 设计决策说明

### 注册时机：Boot() 方法
- **理由**：ConfigService 需要先注册完成
- **避免**：Register() 阶段的循环依赖
- **方式**：使用 `Make()` 而非 `MustMake()` 避免 panic

### Key 映射：ConfigName() 独立
- **理由**：解耦 Provider Name 和配置路径
- **灵活**：支持嵌套配置如 `database.mysql.*`

### 热重载：不自动刷新
- **理由**：避免运行时配置变更的不可预期行为
- **安全**：配置变更后重启应用更安全
- **扩展**：可按需提供 Reload 方法

---

## 参考资料

- [12-Factor App - Config](https://12factor.net/config)
- [Koanf 文档](https://github.com/knadh/koanf)
- [Spring Boot Externalized Configuration](https://docs.spring.io/spring-boot/reference/features/external-config.html)
