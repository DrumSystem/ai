package config

import (
	"github.com/zeromicro/go-zero/rest"
)

type Config struct {
	rest.RestConf
	OpenAI struct {
		ApiKey      string
		Model       string
		BaseURL     string
		MaxTokens   int
		Temperature float32
	}
	VectorDB      VectorDBConfig // 新增向量数据库配置
	UniPDFLicense string
}

// VectorDBConfig 向量数据库配置
type VectorDBConfig struct {
	Host           string
	Port           int
	DBName         string
	User           string
	Password       string
	Table          string
	MaxConn        int
	EmbeddingModel string
}
