package commands

import (
	"fmt"
)

type ViewRulesCommand struct{}

func NewViewRulesCommand() *ViewRulesCommand {
	return &ViewRulesCommand{}
}

func (c *ViewRulesCommand) Type() CommandType { return CmdViewRules }

func (c *ViewRulesCommand) Validate() error {
	return nil
}

func (c *ViewRulesCommand) RequiredInputs() []string {
	return []string{}
}

func (c *ViewRulesCommand) Execute(ctx CommandContext) error {
	gameState := ctx.GameState

	if gameState.Script == nil {
		return fmt.Errorf("没有选择剧本")
	}

	fmt.Println("剧本规则:")
	fmt.Printf("剧本名称: %s\n", gameState.Script.Title)
	fmt.Printf("最大循环数: %d\n", gameState.Script.MaxLoops)
	fmt.Printf("每循环天数: %d\n", gameState.Script.DaysPerLoop)

	if gameState.Script.MainPlot != nil {
		fmt.Printf("\n主要剧情: %s\n", gameState.Script.MainPlot.Name)
		fmt.Printf("描述: %s\n", gameState.Script.MainPlot.Description)
	}

	if len(gameState.Script.SubPlots) > 0 {
		fmt.Println("\n子剧情:")
		for i, subplot := range gameState.Script.SubPlots {
			fmt.Printf("%d. %s\n", i+1, subplot.Name)
			fmt.Printf("   描述: %s\n", subplot.Description)
		}
	}

	fmt.Println("\n角色信息:")
	for _, char := range gameState.Characters {
		fmt.Printf("- %s (%s)\n", char.Name, char.StartLocation)
	}

	return nil
}
