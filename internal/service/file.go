package service

import (
	"fmt"
	"io"
	"os"
	"path/filepath"

	"git-uncommitted-backup/internal/console"
	"git-uncommitted-backup/internal/i18n"
)

// FileService 文件服务
type FileService struct{}

// NewFileService 创建文件服务
func NewFileService() *FileService {
	return &FileService{}
}

// BackupFiles 备份文件到指定目录
func (f *FileService) BackupFiles(files []string, destinationDir string) (int, int) {
	successCount := 0
	skippedCount := 0

	for _, file := range files {
		// 检查文件是否存在（可能被删除了）
		if _, err := os.Stat(file); os.IsNotExist(err) {
			console.PrintYellow("  ⚠ " + i18n.Tf(i18n.KeySkippedDeletedFile+": %s", file))
			skippedCount++
			continue
		}

		// 构建目标路径，保持原有的目录结构
		destPath := filepath.Join(destinationDir, file)

		// 创建目标目录
		destDir := filepath.Dir(destPath)
		if err := os.MkdirAll(destDir, 0755); err != nil {
			console.PrintRed("  ✗ " + fmt.Sprintf(i18n.T(i18n.KeyCreateDirFailed)+" %s: %v", destDir, err))
			continue
		}

		// 复制文件
		if err := copyFile(file, destPath); err != nil {
			console.PrintRed("  ✗ " + fmt.Sprintf(i18n.T(i18n.KeyCopyFailed)+" %s: %v", file, err))
			continue
		}

		console.PrintGreen("  ✓ " + i18n.T(i18n.KeyCopied) + ": " + file)
		successCount++
	}

	return successCount, skippedCount
}

// copyFile 复制文件
func copyFile(src, dst string) error {
	srcFile, err := os.Open(src)
	if err != nil {
		return err
	}
	defer srcFile.Close()

	destFile, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer destFile.Close()

	_, err = io.Copy(destFile, srcFile)
	return err
}