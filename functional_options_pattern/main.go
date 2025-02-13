package main

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
