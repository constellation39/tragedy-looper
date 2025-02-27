package commands

import (
	"fmt"
)

type ViewHistoryCommand struct{}

func NewViewHistoryCommand() *ViewHistoryCommand {
	return &ViewHistoryCommand{}
}

func (c *ViewHistoryCommand) Type() CommandType { return CmdViewHistory }

func (c *ViewHistoryCommand) Validate() error {
	return nil
}

func (c *ViewHistoryCommand) RequiredInputs() []string {
	return []string{}
}

func (c *ViewHistoryCommand) Execute(ctx CommandContext) error {
	// Note: This implementation assumes a history log in the GameState
	// Since it's not shown in the provided models, this is a placeholder
	// that would need to be updated once the history tracking is implemented

	fmt.Println("游戏历史记录:")
	fmt.Printf("当前循环: %d\n", ctx.GameState.CurrentLoop)
	fmt.Printf("当前天数: %d\n", ctx.GameState.CurrentDay)
	fmt.Printf("当前阶段: %s\n", ctx.GameState.CurrentPhase)

	// Display incidents (as a simple history implementation)
	if len(ctx.GameState.Incidents) > 0 {
		fmt.Println("\n事件记录:")
		for i, incident := range ctx.GameState.Incidents {
			fmt.Printf("%d. 循环 %d, 天 %d: %s\n",
				i+1,
				incident.GetLoop(),
				incident.GetDay(),
				incident.GetName())
		}
	} else {
		fmt.Println("\n尚无事件记录")
	}

	return nil
}
