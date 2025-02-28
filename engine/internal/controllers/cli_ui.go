package controllers

import (
	"fmt"
	"github.com/pterm/pterm"
)

// UI 接口用于多种交互模式
type UI interface {
	Select(title string, options []string) (int, error)
	MultiSelect(title string, options []string) ([]int, error)
	ShowInfo(msg string) // 新增显示信息的方法
}

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

// MultiSelect 实现多选功能
func (t *TerminalUI) MultiSelect(title string, options []string) ([]int, error) {
	prompt := pterm.DefaultInteractiveMultiselect
	prompt.DefaultText = title
	prompt.Options = options
	prompt.MaxHeight = 10

	results, err := prompt.Show()
	if err != nil {
		return nil, err
	}

	// 找出选中项的索引
	var selectedIndices []int
	for _, result := range results {
		for i, opt := range options {
			if opt == result {
				selectedIndices = append(selectedIndices, i)
				break
			}
		}
	}

	return selectedIndices, nil
}

// ShowInfo 实现消息显示功能
func (t *TerminalUI) ShowInfo(msg string) {
	fmt.Println(msg)
}

// WebUI 保留空结构以便未来扩展
type WebUI struct{}

// ShowInfo 实现消息显示功能
func (w *WebUI) ShowInfo(msg string) {
	// WebUI暂不实现信息显示功能
}
