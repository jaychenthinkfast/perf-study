# 函数选项模式
函数选项模式是一种在 Go 语言中广泛使用的设计模式，它允许在创建结构体实例时，使用可选参数灵活地进行配置，而不会导致构造函数参数过长或过于复杂。

#### **🌟 示例场景**

我们要创建一个 **数据库连接配置（DatabaseConfig）**，并提供不同的可选参数，如：

* **数据库地址（Address）**
* **端口（Port）**
* **用户名（Username）**
* **密码（Password）**
* **超时时间（Timeout）**

#### **📌 代码实现**

```go
import (
	"fmt"
	"time"
)

// DatabaseConfig 结构体用于存储数据库连接配置
type DatabaseConfig struct {
	Address  string        // 数据库地址
	Port     int           // 端口号
	Username string        // 用户名
	Password string        // 密码
	Timeout  time.Duration // 超时时间
}

// Option 类型表示一个函数选项
type Option func(*DatabaseConfig)

// WithAddress 允许设置数据库地址
func WithAddress(address string) Option {
	return func(cfg *DatabaseConfig) {
		cfg.Address = address
	}
}

// WithPort 允许设置端口号
func WithPort(port int) Option {
	return func(cfg *DatabaseConfig) {
		cfg.Port = port
	}
}

// WithUsername 允许设置数据库用户名
func WithUsername(username string) Option {
	return func(cfg *DatabaseConfig) {
		cfg.Username = username
	}
}

// WithPassword 允许设置数据库密码
func WithPassword(password string) Option {
	return func(cfg *DatabaseConfig) {
		cfg.Password = password
	}
}

// WithTimeout 允许设置数据库连接超时时间
func WithTimeout(timeout time.Duration) Option {
	return func(cfg *DatabaseConfig) {
		cfg.Timeout = timeout
	}
}

// NewDatabaseConfig 创建一个新的 DatabaseConfig，并应用所有选项
func NewDatabaseConfig(opts ...Option) *DatabaseConfig {
	// 设置默认值
	cfg := &DatabaseConfig{
		Address:  "localhost",
		Port:     5432,
		Username: "root",
		Password: "",
		Timeout:  5 * time.Second, // 默认超时时间 5 秒
	}

	// 应用用户提供的选项
	for _, opt := range opts {
		opt(cfg)
	}

	return cfg
}

func main() {
	// 使用函数选项模式创建配置
	dbConfig := NewDatabaseConfig(
		WithAddress("192.168.1.100"),
		WithPort(3306),
		WithUsername("admin"),
		WithPassword("secret"),
		WithTimeout(10*time.Second),
	)

	// 打印配置
	fmt.Printf("Database Config: %+v\n", dbConfig)
}
```

---

### **🔍 代码解析**

1. **`DatabaseConfig` 结构体**
   * 存储数据库的配置信息，如地址、端口、用户名、密码和超时设置。
2. **`Option` 类型**
   * `type Option func(*DatabaseConfig)`
   * 这是一个函数类型，接收一个 `*DatabaseConfig` 指针并对其进行修改。
3. **提供多个 `WithXXX` 选项函数**
   * `WithAddress(address string) Option`
   * `WithPort(port int) Option`
   * `WithUsername(username string) Option`
   * `WithPassword(password string) Option`
   * `WithTimeout(timeout time.Duration) Option`
   * 这些函数返回一个 `Option` 类型的函数，应用到 `DatabaseConfig` 上。
4. **`NewDatabaseConfig` 构造函数**
   * 先创建 `DatabaseConfig`，并设置默认值。
   * 遍历传入的 `opts` 选项并应用到 `cfg`。
5. **示例使用**
   * 通过 `NewDatabaseConfig()` 传入不同的选项进行自定义配置。

---

### **📌 为什么使用函数选项模式？**

✅ **更灵活**：可以传入任意数量的选项，不需要维护多个构造函数。

✅ **避免参数爆炸**：如果使用传统的构造函数，参数可能会越来越多，难以维护。

✅ **默认值友好**：可以提供合理的默认值，不指定某个选项时，仍然有默认行为。

---

### **🚀 运行结果**

```go
Database Config: &{Address:192.168.1.100 Port:3306 Username:admin Password:secret Timeout:10s}
```

这就成功地创建了一个带有自定义选项的数据库配置！🎯
