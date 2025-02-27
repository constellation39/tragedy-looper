package commands

import (
	"fmt"
)

type ViewIncidentsCommand struct{}

func NewViewIncidentsCommand() *ViewIncidentsCommand {
	return &ViewIncidentsCommand{}
}

func (c *ViewIncidentsCommand) Type() CommandType { return CmdViewIncidents }

func (c *ViewIncidentsCommand) Validate() error {
	return nil
}

func (c *ViewIncidentsCommand) RequiredInputs() []string {
	return []string{}
}

func (c *ViewIncidentsCommand) Execute(ctx CommandContext) error {
	gameState := ctx.GameState

	if len(gameState.Incidents) == 0 {
		fmt.Println("尚无事件发生")
		return nil
	}

	fmt.Println("已发生的事件:")
	for i, incident := range gameState.Incidents {
		fmt.Printf("%d. 循环 %d, 天 %d: %s\n",
			i+1,
			incident.GetLoop(),
			incident.GetDay(),
			incident.GetName())
		fmt.Printf("   %s\n", incident.GetDescription())
	}

	return nil
}
