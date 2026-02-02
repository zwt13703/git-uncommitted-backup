package i18n

import (
	"fmt"
	"strings"
)

var currentLang = "zh-CN"
var translations = make(map[string]map[string]string)

// 定义所有翻译 key
const (
	KeyTitle                     = "title"
	KeyVersion                   = "version"
	KeyAuthor                    = "author"
	KeyMenuView                  = "menu_view"
	KeyMenuSave                  = "menu_save"
	KeyMenuReset                 = "menu_reset"
	KeyMenuExit                  = "menu_exit"
	KeySelectOption              = "select_option"
	KeyInvalidChoice             = "invalid_choice"
	KeyCheckingChanges           = "checking_changes"
	KeyNoChanges                 = "no_changes"
	KeyFoundFiles                = "found_files"
	KeyEnterBackupDir            = "enter_backup_dir"
	KeyDirCannotBeEmpty          = "dir_cannot_be_empty"
	KeyCopyingTo                 = "copying_to"
	KeyComplete                  = "complete"
	KeySuccess                   = "success"
	KeySkipped                   = "skipped"
	KeyWarningReset              = "warning_reset"
	KeyConfirmReset              = "confirm_reset"
	KeyCancelled                 = "cancelled"
	KeyExecutingReset            = "executing_reset"
	KeyResetSuccess              = "reset_success"
	KeyOutput                    = "output"
	KeyNotGitRepo                = "not_git_repo"
	KeyPleaseRunInGitRepo        = "please_run_in_git_repo"
)

func Init(lang string) {
	currentLang = lang
	loadTranslations()
}

func T(key string) string {
	if langMap, ok := translations[currentLang]; ok {
		if text, ok := langMap[key]; ok {
			return text
		}
	}
	// 如果找不到翻译，返回 key 本身
	return key
}

func Tf(key string, args ...interface{}) string {
	return fmt.Sprintf(T(key), args...)
}

func loadTranslations() {
	translations["zh-CN"] = loadChinese()
	translations["en-US"] = loadEnglish()

	// 如果当前语言不是支持的语言，默认使用中文
	if _, ok := translations[currentLang]; !ok {
		currentLang = "zh-CN"
	}
}

func loadChinese() map[string]string {
	return map[string]string{
		KeyTitle:                     "Git 未提交代码管理工具",
		KeyVersion:                   "版本",
		KeyAuthor:                    "作者",
		KeyMenuView:                  "查看当前未提交的代码文件",
		KeyMenuSave:                  "将未提交的代码另存到指定目录",
		KeyMenuReset:                 "驳回未提交的代码 (git reset --hard)",
		KeyMenuExit:                  "退出程序",
		KeySelectOption:              "请选择操作 [1-4]",
		KeyInvalidChoice:             "无效的选择，请重新输入。",
		KeyCheckingChanges:           "正在检测未提交的更改...",
		KeyNoChanges:                 "当前没有未提交的更改。",
		KeyFoundFiles:                "找到 %d 个未提交的文件",
		KeyEnterBackupDir:            "请输入目标备份目录",
		KeyDirCannotBeEmpty:          "目标目录不能为空！",
		KeyCopyingTo:                 "正在复制文件到",
		KeyComplete:                  "完成",
		KeySuccess:                   "成功",
		KeySkipped:                   "跳过",
		KeyWarningReset:              "警告: 此操作将删除所有未提交的更改，且无法恢复！",
		KeyConfirmReset:              "请输入 YES 确认驳回所有未提交的更改",
		KeyCancelled:                 "操作已取消。",
		KeyExecutingReset:            "正在执行 git reset --hard...",
		KeyResetSuccess:              "成功驳回所有未提交的更改！",
		KeyOutput:                    "输出",
		KeyNotGitRepo:                "错误: 不是 Git 仓库",
		KeyPleaseRunInGitRepo:        "请在 Git 仓库目录下运行此程序。",
	}
}

func loadEnglish() map[string]string {
	return map[string]string{
		KeyTitle:                     "Git Uncommitted Code Manager",
		KeyVersion:                   "Version",
		KeyAuthor:                    "Author",
		KeyMenuView:                  "View uncommitted files",
		KeyMenuSave:                  "Backup uncommitted files to directory",
		KeyMenuReset:                 "Reset uncommitted changes (git reset --hard)",
		KeyMenuExit:                  "Exit program",
		KeySelectOption:              "Select option [1-4]",
		KeyInvalidChoice:             "Invalid choice, please try again.",
		KeyCheckingChanges:           "Checking for uncommitted changes...",
		KeyNoChanges:                 "No uncommitted changes found.",
		KeyFoundFiles:                "Found %d uncommitted file(s)",
		KeyEnterBackupDir:            "Enter backup directory",
		KeyDirCannotBeEmpty:          "Backup directory cannot be empty!",
		KeyCopyingTo:                 "Copying files to",
		KeyComplete:                  "Complete",
		KeySuccess:                   "Success",
		KeySkipped:                   "Skipped",
		KeyWarningReset:              "WARNING: This will discard all uncommitted changes and cannot be undone!",
		KeyConfirmReset:              "Type YES to confirm resetting all uncommitted changes",
		KeyCancelled:                 "Operation cancelled.",
		KeyExecutingReset:            "Executing git reset --hard...",
		KeyResetSuccess:              "Successfully reset all uncommitted changes!",
		KeyOutput:                    "Output",
		KeyNotGitRepo:                "Error: Not a git repository",
		KeyPleaseRunInGitRepo:        "Please run this program in a git repository.",
	}
}

// SetLanguage 设置语言
func SetLanguage(lang string) {
	if strings.HasPrefix(lang, "zh") {
		currentLang = "zh-CN"
	} else if strings.HasPrefix(lang, "en") {
		currentLang = "en-US"
	}
}