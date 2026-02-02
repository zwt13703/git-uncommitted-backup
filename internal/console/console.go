package console

import (
	"os"
	"syscall"
	"unsafe"
)

const (
	ColorReset  = "\033[0m"
	ColorRed    = "\033[31m"
	ColorGreen  = "\033[32m"
	ColorYellow = "\033[33m"
	ColorBlue   = "\033[34m"
	ColorPurple = "\033[35m"
	ColorCyan   = "\033[36m"
	ColorWhite  = "\033[37m"
	ColorBold   = "\033[1m"
)

var (
	kernel32                       = syscall.NewLazyDLL("kernel32.dll")
	procSetConsoleOutputCP         = kernel32.NewProc("SetConsoleOutputCP")
	procGetConsoleMode             = kernel32.NewProc("GetConsoleMode")
	procSetConsoleMode             = kernel32.NewProc("SetConsoleMode")
	ENABLE_VIRTUAL_TERMINAL_PROCESSING = uint32(0x0004)
)

// Setup 初始化控制台支持 ANSI 颜色和 UTF-8
func Setup() {
	// 设置控制台输出编码为 UTF-8 (CP65001)
	procSetConsoleOutputCP.Call(65001)

	// 在 Windows 上启用 ANSI 转义序列支持
	stdoutHandle := syscall.Handle(os.Stdout.Fd())
	var mode uint32

	// 获取当前控制台模式
	procGetConsoleMode.Call(uintptr(stdoutHandle), uintptr(unsafe.Pointer(&mode)))

	// 设置新的模式，启用虚拟终端处理
	procSetConsoleMode.Call(uintptr(stdoutHandle), uintptr(unsafe.Pointer(&mode)), uintptr(mode|ENABLE_VIRTUAL_TERMINAL_PROCESSING))
}

// PrintRed 红色输出
func PrintRed(text string) {
	printColored(ColorRed, text)
}

// PrintGreen 绿色输出
func PrintGreen(text string) {
	printColored(ColorGreen, text)
}

// PrintYellow 黄色输出
func PrintYellow(text string) {
	printColored(ColorYellow, text)
}

// PrintBlue 蓝色输出
func PrintBlue(text string) {
	printColored(ColorBlue, text)
}

// PrintCyan 青色输出
func PrintCyan(text string) {
	printColored(ColorCyan, text)
}

// PrintPurple 紫色输出
func PrintPurple(text string) {
	printColored(ColorPurple, text)
}

// PrintWhite 白色输出
func PrintWhite(text string) {
	printColored(ColorWhite, text)
}

// PrintBold 粗体输出
func PrintBold(text string) {
	printColored(ColorBold, text)
}

func printColored(color, text string) {
	println(color + text + ColorReset)
}