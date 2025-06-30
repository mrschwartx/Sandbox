package config

import "fmt"

type Server struct {
	Host string `yaml:"host"`
	Port string `yaml:"port"`
}

func (t Server) GetAddr() string {
	return fmt.Sprintf("%s:%s", t.Host, t.Port)
}

type Logging struct {
	Level          string `yaml:"level"`
	Type           string `yaml:"type"`
	LogFileEnabled bool   `yaml:"logFileEnabled"`
	LogFilePath    string `yaml:"logFilePath"`
}

type SQL struct {
	User        string `yaml:"user"`
	Password    string `yaml:"password"`
	Host        string `yaml:"host"`
	Name        string `yaml:"name"`
	Port        string `yaml:"port"`
	MaxIdleConn int    `yaml:"maxIdleConn"`
	MaxOpenConn int    `yaml:"maxOpenConn"`
}

func (s SQL) GetUrl() string {
	return fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		s.User, s.Password, s.Host, s.Port, s.Name)
}

func (s SQL) GetDataSource() string {
	return fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s sslmode=disable",
		s.User, s.Password, s.Host, s.Port, s.Name)
}

type Redis struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	DB       int    `yaml:"db"`
	Password string `yaml:"password"`
	// Timeout settings (in seconds)
	ConnDialTimeoutSec int `yaml:"connDialTimeoutSec"` // Dial timeout (default: 5)
	ReadTimeoutSec     int `yaml:"readTimeoutSec"`     // Read timeout: 0=default (3s), -1=none, -2=disable deadline
	WriteTimeoutSec    int `yaml:"writeTimeoutSec"`    // Write timeout: same rules as ReadTimeout
	// Connection pool settings
	ConnPoolSize       int `yaml:"connPoolSize"`       // Max connections (default: 10 per CPU)
	ConnPoolTimeoutSec int `yaml:"connPoolTimeoutSec"` // Wait timeout when pool is full (default: ReadTimeout + 1s)
	MinIdleConn        int `yaml:"minIdleConn"`        // Min idle connections (default: 0)
	MaxIdleConn        int `yaml:"maxIdleConn"`        // Max idle connections (default: 0)
}

func (r Redis) GetAddr() string {
	return r.Host + ":" + r.Port
}
