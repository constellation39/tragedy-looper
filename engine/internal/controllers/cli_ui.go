package controllers

import (
	"fmt"
	"github.com/pterm/pterm"
)

// TerminalUI 实现使用pterm的终端UI
type TerminalUI struct{}

func (t *TerminalUI) Select(title string, options []string) (int, error) {
	prompt := pterm.DefaultInteractiveSelect
	prompt.DefaultText = title
	prompt.Options = options
	prompt.MaxHeight = 10

	result, err := prompt.Show()
	if err != nil {
		return -1, err
	}

	// 直接匹配结果字符串
	for i, opt := range options {
		if opt == result {
			return i, nil
		}
	}
	return -1, fmt.Errorf("invalid selection")
}

// WebUI 保留空结构以便未来扩展
type WebUI struct{}
