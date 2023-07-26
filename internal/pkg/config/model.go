package config

import "time"

// options
type Options struct {
	App    `yaml:",inline"` // App 运行设置
	System `yaml:",inline"` // App 系统级配置

	Mysql `yaml:"mysql"` // 数据库配置
	Redis `yaml:"redis"` // 数据库配置
}

// app env config
type App struct {
	Addr string `yaml:"addr"` // App 地址
	Port string `yaml:"port"` // App 端口
}

type System struct {
	Logdir   string `yaml:"logdir"`   // 日志路径
	Datadir  string `yaml:"datadir"`  // 数据路径
	LogLevel string `yaml:"loglevel"` // 日志等级
	TimeZone string `yaml:"timezone"` // 时区设置
}

type Mysql struct {
	Addr     string `yaml:"addr"`     // 地址
	Port     string `yaml:"port"`     // 端口
	User     string `yaml:"user"`     // 用户
	Passwd   string `yaml:"passwd"`   // 密码
	Database string `yaml:"database"` // 库名
	Options  string `yaml:"options"`  // 选项

	MaxOpenConns    int           `yaml:"MaxOpenConns"`
	MaxIdleConns    int           `yaml:"MaxIdleConns"`
	ConnMaxIdleTime time.Duration `yaml:"ConnMaxIdleTime"`
	ConnMaxLifetime time.Duration `yaml:"ConnMaxLifetime"`
}

type Redis struct {
	Addrs      []string `yaml:"addrs"`      // 地址
	Passwd     string   `yaml:"passwd"`     // 密码
	MasterName string   `yaml:"mastername"` // 选项
}
