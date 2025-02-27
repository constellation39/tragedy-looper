// engine/internal/controllers/commands/show_help.go
func (c *ShowHelpCommand) Execute(ctx CommandContext) error {
    fmt.Println("可用命令:")
    fmt.Println("  start - 开始新游戏")  // 新增帮助条目
    fmt.Println("  help - 显示帮助")
    // ... 其他已有命令
    return nil
}
