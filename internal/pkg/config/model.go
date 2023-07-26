package config

// options
type Options struct {
	System `yaml:",inline"` // App 系统级配置
}

type System struct {
	Logdir   string `yaml:"logdir"`   // 日志路径
	Datadir  string `yaml:"datadir"`  // 数据路径
	LogLevel string `yaml:"loglevel"` // 日志等级
	TimeZone string `yaml:"timezone"` // 时区设置
}
