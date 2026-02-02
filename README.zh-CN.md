# Git 未提交代码备份工具

[![Version](https://img.shields.io/badge/version-1.0.0-blue.svg)](https://github.com/zwt13703/git-uncommitted-backup)
[![Go](https://img.shields.io/badge/Go-1.21+-00ADD8?style=flat&logo=go)](https://golang.org/)
[![License](https://img.shields.io/badge/license-MIT-green.svg)](LICENSE)

Git 未提交代码备份工具 - 一个帮助你备份和查看 Git 仓库中未提交代码的命令行工具。

## 功能特性

- ✅ 查看当前分支下所有未提交的文件
- ✅ 将未提交的代码连同目录结构一起备份到指定目录
- ✅ 驳回未提交的代码（执行 `git reset --hard`）
- ✅ 支持中英文界面
- ✅ 彩色命令行界面，操作直观
- ✅ 自动检测 Git 仓库环境

## 适用场景

当你在错误的分支写了代码但还没有 commit 时，可以使用此工具：

1. 查看当前未提交的文件列表
2. 将未提交的代码备份到指定目录
3. 驳回当前分支的未提交更改
4. 切换到正确的分支
5. 将备份的文件复制回来并提交

## 截图

```
╔══════════════════════════════════════════════════════════════╗
║                    Git   未提交代码备份工具                    ║
╠══════════════════════════════════════════════════════════════╣
║  版本: 1.0.0							                       ║
╠══════════════════════════════════════════════════════════════╣
║  GitHub: https://github.com/zwt13703/git-uncommitted-backup  ║
╚══════════════════════════════════════════════════════════════╝

══════════════════════════════════════════════════════════════
  [1] 查看当前未提交的代码文件
  [2] 将未提交的代码另存到指定目录
  [3] 驳回未提交的代码 (git reset --hard)
  [4] 退出程序
══════════════════════════════════════════════════════════════
```

## 下载

### Windows

下载最新版本的 `git-uncommitted-backup.exe` 文件，双击运行即可。

### 源码编译

```bash
# 克隆仓库
git clone https://github.com/zwt13703/git-uncommitted-backup.git
cd git-uncommitted-backup

# 编译
go build -o git-uncommitted-backup.exe
```

## 使用方法

### 1. 运行程序

```bash
# Windows
git-uncommitted-backup.exe

# 或从源码运行
go run main.go
```

### 2. 选择操作

程序启动后会显示交互式菜单：

- **[1] 查看当前未提交的代码文件** - 列出所有未提交的文件
- **[2] 将未提交的代码另存到指定目录** - 备份未提交的代码
- **[3] 驳回未提交的代码** - 执行 `git reset --hard`，注意此操作不可逆
- **[4] 退出程序** - 退出工具

### 3. 典型工作流程

```bash
# 步骤 1: 运行工具
git-uncommitted-backup.exe

# 步骤 2: 选择 [2] 备份未提交的代码
# 输入备份目录，例如: C:\backup\my_code

# 步骤 3: 选择 [3] 驳回当前分支的未提交更改
# 输入 YES 确认

# 步骤 4: 切换到正确的分支
git checkout correct-branch

# 步骤 5: 将备份的文件复制回来
xcopy C:\backup\my_code\* . /E /I /Y

# 步骤 6: 提交代码
git add .
git commit -m "your commit message"
```

## 项目结构

```
git-uncommitted-backup/
├── main.go                 # 程序入口
├── go.mod                  # Go 模块定义
├── README.md               # 项目文档（英文）
├── README.zh-CN.md         # 项目文档（中文）
└── internal/               # 内部包
    ├── config/             # 配置管理
    │   └── config.go
    ├── console/            # 控制台输出
    │   └── console.go
    ├── i18n/               # 国际化
    │   └── i18n.go
    ├── service/            # 业务逻辑
    │   ├── git.go          # Git 操作
    │   └── file.go         # 文件操作
    └── ui/                 # 用户界面
        └── menu.go         # 菜单逻辑
```

## 系统要求

- Windows 10 或更高版本
- Git 已安装并配置
- 无需额外依赖

## 国际化

工具支持中英文界面，语言会根据系统环境自动切换：

- 中文环境 (zh-CN): 简体中文界面
- 英文环境 (en-US): 英文界面

## 安全性说明

- **驳回操作警告**: 选择 [3] 驳回未提交代码时，需要输入 `YES` 确认，防止误操作
- **非 Git 仓库检测**: 在非 Git 仓库中运行时，程序会给出友好的错误提示

## 注意事项

1. 在执行 [3] 驳回操作前，请确保已经备份了重要的未提交代码
2. 程序必须在 Git 仓库目录下运行
3. 备份操作会保持原有的目录结构

## 开发

### 添加新的翻译

编辑 `internal/i18n/i18n.go` 文件：

```go
const (
    KeyYourNewKey = "your_new_key"
)

func loadChinese() map[string]string {
    return map[string]string{
        KeyYourNewKey: "你的中文翻译",
    }
}

func loadEnglish() map[string]string {
    return map[string]string{
        KeyYourNewKey: "Your English Translation",
    }
}
```

## 版本历史

### v1.0.0 (2026-02-02)
- ✨ 初始版本发布
- ✨ 支持查看未提交文件
- ✨ 支持备份未提交代码
- ✨ 支持驳回未提交代码
- ✨ 中英文界面支持
- ✨ 彩色命令行界面
- ✨ 非 Git 仓库错误处理

## 贡献

欢迎提交 Issue 和 Pull Request！

## 许可证

MIT License

## 仓库地址

https://github.com/zwt13703/git-uncommitted-backup

---

**免责声明**: 使用本工具时请谨慎操作，特别是在执行驳回操作前请确保已备份重要数据。