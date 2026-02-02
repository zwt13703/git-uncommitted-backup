package service

import (
	"bufio"
	"fmt"
	"os/exec"
	"strings"

	"git-uncommitted-backup/internal/console"
	"git-uncommitted-backup/internal/i18n"
)

// GitService Git 服务
type GitService struct{}

// NewGitService 创建 Git 服务
func NewGitService() *GitService {
	return &GitService{}
}

// GetUncommittedFiles 获取未提交的文件列表
func (g *GitService) GetUncommittedFiles() ([]string, error) {
	// 使用 git status --porcelain 获取未提交的文件
	cmd := exec.Command("git", "status", "--porcelain")
	output, err := cmd.CombinedOutput()
	if err != nil {
		// 检查是否是 Git 仓库
		if strings.Contains(string(output), "not a git repository") {
			return nil, fmt.Errorf(i18n.T(i18n.KeyNotGitRepo))
		}
		return nil, fmt.Errorf(i18n.T(i18n.KeyNotGitRepo) + ": %v", err)
	}

	var files []string
	scanner := bufio.NewScanner(strings.NewReader(string(output)))
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}

		if len(line) < 3 {
			continue
		}

		filename := strings.TrimSpace(line[2:])
		if filename == "" {
			continue
		}

		status := line[:2]
		if status[0] != ' ' && status[0] != '?' {
			files = append(files, filename)
		} else if status[1] != ' ' {
			files = append(files, filename)
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("failed to read output: %v", err)
	}

	return files, nil
}

// IsGitRepository 检查是否是 Git 仓库
func (g *GitService) IsGitRepository() bool {
	cmd := exec.Command("git", "rev-parse", "--git-dir")
	err := cmd.Run()
	return err == nil
}

// ResetHard 执行 git reset --hard
func (g *GitService) ResetHard() error {
	cmd := exec.Command("git", "reset", "--hard")
	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("git reset --hard failed: %v\n%s", err, string(output))
	}
	console.PrintGreen(i18n.T(i18n.KeyResetSuccess))
	console.PrintCyan(i18n.T(i18n.KeyOutput) + ": " + string(output))
	return nil
}

// GetCurrentBranch 获取当前分支名
func (g *GitService) GetCurrentBranch() (string, error) {
	cmd := exec.Command("git", "branch", "--show-current")
	output, err := cmd.CombinedOutput()
	if err != nil {
		return "", fmt.Errorf("failed to get current branch: %v", err)
	}
	return strings.TrimSpace(string(output)), nil
}