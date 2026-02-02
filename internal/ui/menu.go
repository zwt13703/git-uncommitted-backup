package ui

import (
	"bufio"
	"fmt"
	"os"

	"git-uncommitted-backup/internal/config"
	"git-uncommitted-backup/internal/console"
	"git-uncommitted-backup/internal/i18n"
	"git-uncommitted-backup/internal/service"
)

// PrintHeader 打印程序头部信息
func PrintHeader(cfg *config.Config) {
	fmt.Println()
	fmt.Printf("%s╔══════════════════════════════════════════════════════════════╗%s\n", console.ColorCyan, console.ColorReset)
	fmt.Printf("%s║%s%50s%s║%s\n", console.ColorCyan, console.ColorBold+console.ColorWhite, i18n.T(i18n.KeyTitle), console.ColorCyan, console.ColorReset)
	fmt.Printf("%s╠══════════════════════════════════════════════════════════════╣%s\n", console.ColorCyan, console.ColorReset)
	fmt.Printf("%s║%s  %s: %-13s %s  %s: %-15s              %s║%s\n",
		console.ColorCyan, console.ColorYellow, i18n.T(i18n.KeyVersion), config.Version,
		console.ColorCyan, i18n.T(i18n.KeyAuthor), config.Author, console.ColorCyan, console.ColorReset)
	fmt.Printf("%s╠══════════════════════════════════════════════════════════════╣%s\n", console.ColorCyan, console.ColorReset)
	fmt.Printf("%s║%s  %s: %-49s%s║%s\n",
		console.ColorCyan, console.ColorYellow, "GitHub", config.RepoURL, console.ColorCyan, console.ColorReset)
	fmt.Printf("%s╚══════════════════════════════════════════════════════════════╝%s\n", console.ColorCyan, console.ColorReset)
	fmt.Println()
}

// PrintMenu 打印菜单
func PrintMenu() {
	fmt.Printf("%s%s══════════════════════════════════════════════════════════════%s\n", console.ColorBlue, console.ColorBold, console.ColorReset)
	fmt.Printf("%s  [1]%s %s\n", console.ColorGreen, console.ColorWhite, i18n.T(i18n.KeyMenuView))
	fmt.Printf("%s  [2]%s %s\n", console.ColorGreen, console.ColorWhite, i18n.T(i18n.KeyMenuSave))
	fmt.Printf("%s  [3]%s %s\n", console.ColorRed, console.ColorWhite, i18n.T(i18n.KeyMenuReset))
	fmt.Printf("%s  [4]%s %s\n", console.ColorYellow, console.ColorWhite, i18n.T(i18n.KeyMenuExit))
	fmt.Printf("%s%s══════════════════════════════════════════════════════════════%s\n", console.ColorBlue, console.ColorBold, console.ColorReset)
}

// RunMainLoop 运行主循环
func RunMainLoop(gitService *service.GitService, fileService *service.FileService) {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		PrintMenu()
		fmt.Printf("%s%s: %s", console.ColorCyan, i18n.T(i18n.KeySelectOption), console.ColorReset)

		if !scanner.Scan() {
			break
		}

		choice := scanner.Text()

		switch choice {
		case "1":
			viewUncommittedFiles(gitService)
		case "2":
			saveUncommittedFiles(gitService, fileService, scanner)
		case "3":
			resetUncommittedFiles(gitService, scanner)
		case "4":
			console.PrintGreen(i18n.T("thank_you"))
			os.Exit(0)
		default:
			console.PrintRed(i18n.T(i18n.KeyInvalidChoice))
		}
		fmt.Println()
	}
}

// viewUncommittedFiles 查看未提交的文件
func viewUncommittedFiles(gitService *service.GitService) {
	fmt.Println()
	console.PrintCyan(i18n.T(i18n.KeyCheckingChanges))

	// 检查是否是 Git 仓库
	if !gitService.IsGitRepository() {
		console.PrintRed(i18n.T(i18n.KeyNotGitRepo))
		console.PrintYellow(i18n.T(i18n.KeyPleaseRunInGitRepo))
		return
	}

	files, err := gitService.GetUncommittedFiles()
	if err != nil {
		console.PrintRed(i18n.T(i18n.KeyNotGitRepo) + ": " + err.Error())
		return
	}

	if len(files) == 0 {
		console.PrintGreen(i18n.T(i18n.KeyNoChanges))
		return
	}

	fmt.Printf("\n%s%s: %d%s\n", console.ColorYellow, i18n.T(i18n.KeyFoundFiles), len(files), console.ColorReset)
	fmt.Printf("%s%s─────────────────────────────────────────────────────────%s\n", console.ColorBlue, console.ColorBold, console.ColorReset)
	for i, file := range files {
		fmt.Printf("%s  [%d]%s %s\n", console.ColorGreen, i+1, console.ColorWhite, file)
	}
	fmt.Printf("%s%s─────────────────────────────────────────────────────────%s\n", console.ColorBlue, console.ColorBold, console.ColorReset)
	fmt.Printf("%s%s: %d%s\n", console.ColorYellow, i18n.T("total"), len(files), console.ColorReset)
}

// saveUncommittedFiles 另存未提交的文件
func saveUncommittedFiles(gitService *service.GitService, fileService *service.FileService, scanner *bufio.Scanner) {
	fmt.Println()
	console.PrintCyan(i18n.T(i18n.KeyCheckingChanges))

	// 检查是否是 Git 仓库
	if !gitService.IsGitRepository() {
		console.PrintRed(i18n.T(i18n.KeyNotGitRepo))
		console.PrintYellow(i18n.T(i18n.KeyPleaseRunInGitRepo))
		return
	}

	files, err := gitService.GetUncommittedFiles()
	if err != nil {
		console.PrintRed(i18n.T(i18n.KeyNotGitRepo) + ": " + err.Error())
		return
	}

	if len(files) == 0 {
		console.PrintGreen(i18n.T(i18n.KeyNoChanges))
		return
	}

	fmt.Printf("\n%s%s: %d%s\n", console.ColorYellow, i18n.T(i18n.KeyFoundFiles), len(files), console.ColorReset)

	fmt.Printf("%s%s: %s", console.ColorCyan, i18n.T(i18n.KeyEnterBackupDir), console.ColorReset)
	if !scanner.Scan() {
		return
	}

	destinationDir := scanner.Text()
	if destinationDir == "" {
		console.PrintRed(i18n.T(i18n.KeyDirCannotBeEmpty))
		return
	}

	fmt.Printf("\n%s%s: %s%s\n", console.ColorCyan, i18n.T(i18n.KeyCopyingTo), destinationDir, console.ColorReset)

	successCount, skippedCount := fileService.BackupFiles(files, destinationDir)

	fmt.Printf("\n%s%s══════════════════════════════════════════════════════════════%s\n", console.ColorBlue, console.ColorBold, console.ColorReset)
	fmt.Printf("%s  %s! %s: %d/%d, %s: %d%s\n",
		console.ColorGreen, i18n.T(i18n.KeyComplete), i18n.T(i18n.KeySuccess), successCount, len(files),
		i18n.T(i18n.KeySkipped), skippedCount, console.ColorReset)
	fmt.Printf("%s%s══════════════════════════════════════════════════════════════%s\n", console.ColorBlue, console.ColorBold, console.ColorReset)
}

// resetUncommittedFiles 驳回未提交的文件
func resetUncommittedFiles(gitService *service.GitService, scanner *bufio.Scanner) {
	fmt.Println()
	console.PrintCyan(i18n.T(i18n.KeyCheckingChanges))

	// 检查是否是 Git 仓库
	if !gitService.IsGitRepository() {
		console.PrintRed(i18n.T(i18n.KeyNotGitRepo))
		console.PrintYellow(i18n.T(i18n.KeyPleaseRunInGitRepo))
		return
	}

	files, err := gitService.GetUncommittedFiles()
	if err != nil {
		console.PrintRed(i18n.T(i18n.KeyNotGitRepo) + ": " + err.Error())
		return
	}

	if len(files) == 0 {
		console.PrintGreen(i18n.T(i18n.KeyNoChanges))
		return
	}

	fmt.Printf("\n%s%s: %d%s\n", console.ColorYellow, i18n.T(i18n.KeyFoundFiles), len(files), console.ColorReset)
	fmt.Printf("%s%s─────────────────────────────────────────────────────────%s\n", console.ColorBlue, console.ColorBold, console.ColorReset)
	for i, file := range files {
		fmt.Printf("%s  [%d]%s %s\n", console.ColorGreen, i+1, console.ColorWhite, file)
	}
	fmt.Printf("%s%s─────────────────────────────────────────────────────────%s\n", console.ColorBlue, console.ColorBold, console.ColorReset)

	fmt.Printf("\n%s⚠ %s%s\n", console.ColorRed, i18n.T(i18n.KeyWarningReset), console.ColorReset)
	fmt.Printf("%s%s (YES): %s", console.ColorCyan, i18n.T(i18n.KeyConfirmReset), console.ColorReset)
	if !scanner.Scan() {
		return
	}

	confirm := scanner.Text()
	if confirm != "YES" {
		console.PrintYellow(i18n.T(i18n.KeyCancelled))
		return
	}

	fmt.Printf("\n%s%s%s\n", console.ColorCyan, i18n.T(i18n.KeyExecutingReset), console.ColorReset)

	if err := gitService.ResetHard(); err != nil {
		console.PrintRed(err.Error())
		return
	}
}
