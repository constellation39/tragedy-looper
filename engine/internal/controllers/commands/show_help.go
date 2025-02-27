// engine/internal/controllers/commands/show_help.go
package commands

import "fmt"

// Command struct definition
type ShowHelpCommand struct{}

// Execute function with corrected implementation
func (c *ShowHelpCommand) Execute(ctx CommandContext) error {
    fmt.Println("可用命令:")
    fmt.Println("  start - 开始新游戏")
    fmt.Println("  help - 显示帮助")
    fmt.Println("  move <character> <location> - 移动角色")
    fmt.Println("  pass - 跳过当前操作")
    return nil
}

// Interface method implementations
func (c *ShowHelpCommand) Type() CommandType { return CmdHelp }
func (c *ShowHelpCommand) Validate() error { return nil }
func (c *ShowHelpCommand) RequiredInputs() []string { return []string{} }
