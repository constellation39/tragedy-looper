// engine/internal/controllers/commands/show_plots.go
package commands

import "fmt"

type ShowPlotsCommand struct{}

func NewShowPlotsCommand() *ShowPlotsCommand {
	return &ShowPlotsCommand{}
}

func (c *ShowPlotsCommand) Type() CommandType { return CmdShowPlots }

func (c *ShowPlotsCommand) Validate() error { return nil }

func (c *ShowPlotsCommand) RequiredInputs() []string { return nil }

func (c *ShowPlotsCommand) Execute(ctx CommandContext) error {
	gameState := ctx.GetGameState()
	plots := gameState.GetPlots()
	fmt.Println("剧情:")
	for _, plot := range plots {
		fmt.Printf("  %s - %s\n", plot.GetTitle(), plot.GetDescription())
	}
	return nil
}
