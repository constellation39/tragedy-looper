// engine/internal/controllers/commands/show_cards.go
package commands

import "fmt"

type ShowCardsCommand struct{}

func NewShowCardsCommand() *ShowCardsCommand {
	return &ShowCardsCommand{}
}

func (c *ShowCardsCommand) Type() CommandType { return CmdShowCards }

func (c *ShowCardsCommand) Validate() error { return nil }

func (c *ShowCardsCommand) RequiredInputs() []string { return nil }

func (c *ShowCardsCommand) Execute(ctx CommandContext) error {
	gameState := ctx.GetGameState()
	currentPlayer := gameState.GetCurrentPlayer()
	cards := currentPlayer.GetCards()
	fmt.Println("手牌:")
	for _, card := range cards {
		fmt.Printf("  %s (%s) - %s\n", card.GetTitle(), card.GetType(), card.GetDescription())
	}
	return nil
}
