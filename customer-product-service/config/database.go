package config

type DatabaseConfig struct {
	Driver	string
	Host	string
	Port	int
	User	string
	Pass	string
	Name	string
	Charset	string
	MaxIdleConns	int
	MaxOpenConns	int
	ConnMaxLifetime	int // in seconds
}