// engine/internal/controllers/commands/show_status.go
package commands

import "fmt"

type ShowStatusCommand struct{}

func NewShowStatusCommand() *ShowStatusCommand {
	return &ShowStatusCommand{}
}

func (c *ShowStatusCommand) Type() CommandType { return CmdShowStatus }

func (c *ShowStatusCommand) Validate() error { return nil }

func (c *ShowStatusCommand) RequiredInputs() []string { return nil }

func (c *ShowStatusCommand) Execute(ctx CommandContext) error {
	gameState := ctx.GetGameState()
	fmt.Printf("游戏状态: %s\n", gameState.GetState())
	fmt.Printf("当前天数: %d\n", gameState.GetCurrentDay())
	fmt.Printf("当前玩家: %s\n", gameState.GetCurrentPlayer().GetName())
	fmt.Printf("当前阶段: %s\n", gameState.GetCurrentPhase())
	fmt.Printf("剩余行动点: %d\n", gameState.GetCurrentPlayer().GetActionPoints())
	return nil
}
