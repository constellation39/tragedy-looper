// engine/internal/controllers/commands/show_help.go
package commands

import "fmt"

type ShowHelpCommand struct{}

func NewShowHelpCommand() *ShowHelpCommand {
	return &ShowHelpCommand{}
}

func (c *ShowHelpCommand) Type() CommandType { return CmdShowHelp }

func (c *ShowHelpCommand) Validate() error { return nil }

func (c *ShowHelpCommand) RequiredInputs() []string { return nil }

func (c *ShowHelpCommand) Execute(ctx CommandContext) error {
	fmt.Println("可用命令:")
	fmt.Println("  help - 显示帮助")
	fmt.Println("  status - 显示游戏状态")
	fmt.Println("  board - 显示游戏板")
	fmt.Println("  cards - 显示手牌")
	fmt.Println("  scripts - 显示脚本")
	fmt.Println("  roles - 显示角色")
	fmt.Println("  plots - 显示剧情")
	fmt.Println("  move <character_id> <location_id> - 移动角色")
	fmt.Println("  set <target_type> <target_id> <attribute> <value> - 设置目标属性")
	fmt.Println("  cast <card_id> <target_id> - 使用卡牌")
	fmt.Println("  pass - 跳过当前操作")
	fmt.Println("  quit - 退出游戏")
	return nil
}
