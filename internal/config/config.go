package config

import "os"

const (
	Version      = "1.0.0"
	RepoURL      = "https://github.com/zwt13703/git-uncommitted-backup"
	Author       = "zwt13703"
	ProjectName  = "Git 未提交代码备份工具"
	EnglishName  = "Git Uncommitted Backup"
)

type Config struct {
	Language string
}

func Load() *Config {
	lang := os.Getenv("LANG")
	if lang == "" {
		lang = "zh-CN" // 默认中文
	}

	return &Config{
		Language: lang,
	}
}