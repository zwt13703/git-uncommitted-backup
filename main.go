package main

import (
	"git-uncommitted-backup/internal/config"
	"git-uncommitted-backup/internal/console"
	"git-uncommitted-backup/internal/i18n"
	"git-uncommitted-backup/internal/service"
	"git-uncommitted-backup/internal/ui"
)

func main() {
	// 初始化控制台
	console.Setup()

	// 加载配置
	cfg := config.Load()

	// 初始化国际化
	i18n.Init(cfg.Language)

	// 打印头部信息
	ui.PrintHeader(cfg)

	// 创建服务
	gitService := service.NewGitService()
	fileService := service.NewFileService()

	// 运行主循环
	ui.RunMainLoop(gitService, fileService)
}