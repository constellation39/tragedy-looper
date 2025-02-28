package controllers

import (
	"github.com/pterm/pterm"
	"slices"
)

// UI 接口用于多种交互模式
type UI interface {
	Select(title string, options []string) (int, error)
	MultiSelect(title string, options []string) ([]int, error)
	ShowInfo(msg string) // 新增显示信息的方法
}

// TerminalUI 实现使用pterm的终端UI
type TerminalUI struct {
	selectPrinter      pterm.InteractiveSelectPrinter
	multiSelectPrinter pterm.InteractiveMultiselectPrinter
}

// Select 使用pterm实现选择功能
func (t *TerminalUI) Select(title string, options []string) (int, error) {
	result, err := pterm.DefaultInteractiveSelect.
		WithOptions(options).
		WithDefaultText(title).
		Show()
	if err != nil {
		return -1, err
	}
	return slices.Index(options, result), nil
}

// MultiSelect 使用pterm实现多选功能
func (t *TerminalUI) MultiSelect(title string, options []string) ([]int, error) {
	selected, err := pterm.DefaultInteractiveMultiselect.
		WithOptions(options).
		WithDefaultText(title).
		Show()
	if err != nil {
		return nil, err
	}

	var indexes []int
	for _, s := range selected {
		indexes = append(indexes, slices.Index(options, s))
	}
	return indexes, nil
}

// ShowInfo 使用pterm的带样式的信息显示
func (t *TerminalUI) ShowInfo(msg string) {
	pterm.Info.WithPrefix(pterm.Prefix{
		Text:  "提示",
		Style: pterm.NewStyle(pterm.FgLightCyan),
	}).Println(msg)
}

// WebUI 保留空结构以便未来扩展
type WebUI struct{}

// ShowInfo 实现消息显示功能
func (w *WebUI) ShowInfo(msg string) {
	// WebUI暂不实现信息显示功能
}
