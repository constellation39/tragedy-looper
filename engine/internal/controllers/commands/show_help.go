// engine/internal/controllers/commands/show_help.go
package commands  // Add this as the FIRST line

import "fmt"       // Add import statement

// Keep existing Execute function
func (c *ShowHelpCommand) Execute(ctx CommandContext) error {
    fmt.Println("可用命令:")
    fmt.Println("  start - 开始新游戏")
    fmt.Println("  help - 显示帮助")
    return nil
}

// Add these missing method declarations to fulfill Command interface
func (c *ShowHelpCommand) Type() CommandType { return "help" }
func (c *ShowHelpCommand) Validate() error   { return nil }
func (c *ShowHelpCommand) RequiredInputs() []string { return nil }
