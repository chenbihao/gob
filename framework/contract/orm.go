package contract

import (
	"context"
	"github.com/chenbihao/gob/framework"
	"net"
	"strconv"
	"time"

	"github.com/go-sql-driver/mysql"
	"gorm.io/gorm"
)

/*
## 服务介绍：
提供ORM服务的服务，可以用于获取数据库连接，获取表结构等。
## 支持命令：无
## 支持配置：无
*/

// ORMKey 代表 ORM的服务
const ORMKey = "gob:orm"

type TableColumn struct {
	Field   string `gorm:"column:Field"`   // 列名
	Type    string `gorm:"column:Type"`    // 数据类型
	Null    bool   `gorm:"column:Null"`    // 是否为空
	Key     bool   `gorm:"column:key"`     // 主键类型
	Default string `gorm:"column:Default"` // 默认值
	Comment string `json:"Comment"`        // 注释信息
}

// ORM 表示传入的参数
type ORM interface {
	// 获取 DB
	GetDB(option ...DBOption) (*gorm.DB, error)

	// CanConnect 是否可以连接
	CanConnect(ctx context.Context, db *gorm.DB) (bool, error)

	// Table 相关
	GetTables(ctx context.Context, db *gorm.DB) ([]string, error)
	HasTable(ctx context.Context, db *gorm.DB, table string) (bool, error)
	GetTableColumns(ctx context.Context, db *gorm.DB, table string) ([]TableColumn, error)
}

// DBOption 代表初始化的时候的选项
type DBOption func(container framework.Container, config *DBConfig) error

// DBConfig 代表数据库连接的所有配置
type DBConfig struct {
	// 以下配置关于gorm
	*gorm.Config // 集成gorm的配置

	// 以下配置关于dsn
	WriteTimeout         string `yaml:"write_timeout"`          // 写超时时间
	Loc                  string `yaml:"loc"`                    // 时区
	Port                 int    `yaml:"port"`                   // 端口
	ReadTimeout          string `yaml:"read_timeout"`           // 读超时时间
	Charset              string `yaml:"charset"`                // 字符集
	ParseTime            bool   `yaml:"parse_time"`             // 是否解析时间
	Protocol             string `yaml:"protocol"`               // 传输协议
	Dsn                  string `yaml:"dsn"`                    // 直接传递dsn，如果传递了，其他关于dsn的配置均无效
	Database             string `yaml:"database"`               // 数据库
	Collation            string `yaml:"collation"`              // 字符序
	Timeout              string `yaml:"timeout"`                // 连接超时时间
	Username             string `yaml:"username"`               // 用户名
	Password             string `yaml:"password"`               // 密码
	Driver               string `yaml:"driver"`                 // 驱动
	Host                 string `yaml:"host"`                   // 数据库地址
	AllowNativePasswords bool   `yaml:"allow_native_passwords"` // 是否允许nativePassword

	// 以下配置关于连接池
	ConnMaxIdle     int    `yaml:"conn_max_idle"`     // 最大空闲连接数
	ConnMaxOpen     int    `yaml:"conn_max_open"`     // 最大连接数
	ConnMaxLifetime string `yaml:"conn_max_lifetime"` // 连接最大生命周期
	ConnMaxIdletime string `yaml:"conn_max_idletime"` // 空闲最大生命周期
}

// FormatDsn 生成dsn
func (conf *DBConfig) FormatDsn() (string, error) {
	port := strconv.Itoa(conf.Port)
	timeout, err := time.ParseDuration(conf.Timeout)
	if err != nil {
		return "", err
	}
	readTimeout, err := time.ParseDuration(conf.ReadTimeout)
	if err != nil {
		return "", err
	}
	writeTimeout, err := time.ParseDuration(conf.WriteTimeout)
	if err != nil {
		return "", err
	}
	location, err := time.LoadLocation(conf.Loc)
	if err != nil {
		return "", err
	}
	driverConf := &mysql.Config{
		User:                 conf.Username,
		Passwd:               conf.Password,
		Net:                  conf.Protocol,
		Addr:                 net.JoinHostPort(conf.Host, port),
		DBName:               conf.Database,
		Collation:            conf.Collation,
		Loc:                  location,
		Timeout:              timeout,
		ReadTimeout:          readTimeout,
		WriteTimeout:         writeTimeout,
		ParseTime:            conf.ParseTime,
		AllowNativePasswords: conf.AllowNativePasswords,
	}
	return driverConf.FormatDSN(), nil
}
