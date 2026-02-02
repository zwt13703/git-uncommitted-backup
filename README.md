# Git Uncommitted Backup

[![Version](https://img.shields.io/badge/version-1.0.0-blue.svg)](https://github.com/zwt13703/git-uncommitted-backup)
[![Go](https://img.shields.io/badge/Go-1.21+-00ADD8?style=flat&logo=go)](https://golang.org/)
[![License](https://img.shields.io/badge/license-MIT-green.svg)](LICENSE)

Git Uncommitted Backup - A command-line tool to help you backup and view uncommitted code in your Git repository.

## Features

- ✅ View all uncommitted files in the current branch
- ✅ Backup uncommitted code to a specified directory with full directory structure
- ✅ Reset uncommitted changes (execute `git reset --hard`)
- ✅ Support for Chinese and English interfaces
- ✅ Colorful command-line interface for intuitive operation
- ✅ Automatic Git repository environment detection

## Use Cases

When you've written code in the wrong branch but haven't committed yet:

1. View the list of uncommitted files
2. Backup the uncommitted code to a specified directory
3. Reset the uncommitted changes in the current branch
4. Switch to the correct branch
5. Copy the backup files back and commit

## Screenshots

```
╔══════════════════════════════════════════════════════════════╗
║                    Git Uncommitted Backup                    ║
╠══════════════════════════════════════════════════════════════╣
║  Version: 1.0.0							                   ║
╠══════════════════════════════════════════════════════════════╣
║  GitHub: https://github.com/zwt13703/git-uncommitted-backup  ║
╚══════════════════════════════════════════════════════════════╝

══════════════════════════════════════════════════════════════
  [1] View uncommitted files
  [2] Backup uncommitted files to directory
  [3] Reset uncommitted changes (git reset --hard)
  [4] Exit program
══════════════════════════════════════════════════════════════
```

## Download

### Windows

Download the latest `git-uncommitted-backup.exe` file and double-click to run.

### Build from Source

```bash
# Clone the repository
git clone https://github.com/zwt13703/git-uncommitted-backup.git
cd git-uncommitted-backup

# Build
go build -o git-uncommitted-backup.exe
```

## Usage

### 1. Run the Program

```bash
# Windows
git-uncommitted-backup.exe

# Or run from source
go run main.go
```

### 2. Select an Option

The program will display an interactive menu:

- **[1] View uncommitted files** - List all uncommitted files
- **[2] Backup uncommitted files to directory** - Backup uncommitted code
- **[3] Reset uncommitted changes** - Execute `git reset --hard`, note this operation is irreversible
- **[4] Exit program** - Exit the tool

### 3. Typical Workflow

```bash
# Step 1: Run the tool
git-uncommitted-backup.exe

# Step 2: Select [2] to backup uncommitted code
# Enter the backup directory, e.g., C:\backup\my_code

# Step 3: Select [3] to reset uncommitted changes in the current branch
# Type YES to confirm

# Step 4: Switch to the correct branch
git checkout correct-branch

# Step 5: Copy the backup files back
xcopy C:\backup\my_code\* . /E /I /Y

# Step 6: Commit the code
git add .
git commit -m "your commit message"
```

## Project Structure

```
git-uncommitted-backup/
├── main.go                 # Program entry
├── go.mod                  # Go module definition
├── README.md               # Project documentation (English)
├── README.zh-CN.md         # Project documentation (Chinese)
└── internal/               # Internal packages
    ├── config/             # Configuration management
    │   └── config.go
    ├── console/            # Console output
    │   └── console.go
    ├── i18n/               # Internationalization
    │   └── i18n.go
    ├── service/            # Business logic
    │   ├── git.go          # Git operations
    │   └── file.go         # File operations
    └── ui/                 # User interface
        └── menu.go         # Menu logic
```

## System Requirements

- Windows 10 or higher
- Git installed and configured
- No additional dependencies required

## Internationalization

The tool supports Chinese and English interfaces, which automatically switches based on system environment:

- Chinese environment (zh-CN): Simplified Chinese interface
- English environment (en-US): English interface

## Security Notes

- **Reset Operation Warning**: When selecting [3] to reset uncommitted code, you need to type `YES` to confirm to prevent accidental operations
- **Non-Git Repository Detection**: When running in a non-Git repository, the program will provide friendly error messages

## Notes

1. Before performing the [3] reset operation, make sure you have backed up important uncommitted code
2. The program must be run in a Git repository directory
3. Backup operations preserve the original directory structure

## Development

### Adding New Translations

Edit `internal/i18n/i18n.go` file:

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

## Version History

### v1.0.0 (2026-02-02)
- ✨ Initial release
- ✨ Support viewing uncommitted files
- ✨ Support backing up uncommitted code
- ✨ Support resetting uncommitted code
- ✨ Chinese and English interface support
- ✨ Colorful command-line interface
- ✨ Non-Git repository error handling

## Contributing

Issues and Pull Requests are welcome!

## License

MIT License

## Repository

https://github.com/zwt13703/git-uncommitted-backup

---

**Disclaimer**: Please use this tool with caution, especially ensure important data is backed up before performing reset operations.