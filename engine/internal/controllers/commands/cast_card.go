// engine/internal/controllers/commands/cast_card.go
package commands

import (
	"fmt"
)

type CastCardCommand struct {
	CardID   string
	TargetID string
}

func NewCastCardCommand(cardID, targetID string) *CastCardCommand {
	return &CastCardCommand{
		CardID:   cardID,
		TargetID: targetID,
	}
}

func (c *CastCardCommand) Type() CommandType { return CmdCastCard }

func (c *CastCardCommand) Validate() error {
	if c.CardID == "" {
		return fmt.Errorf("card ID is required")
	}
	if c.TargetID == "" {
		return fmt.Errorf("target ID is required")
	}
	return nil
}

func (c *CastCardCommand) RequiredInputs() []string {
	return []string{"card_id", "target_id"}
}

func (c *CastCardCommand) Execute(ctx CommandContext) error {
	gameState := ctx.GetGameState()
	currentPlayer := gameState.GetCurrentPlayer()
	cards := currentPlayer.GetCards()
	var card interface {
		GetID() string
		GetTitle() string
		GetType() string
		GetDescription() string
		CanCastOn(target interface{}) bool
		Cast(target interface{}) error
	}
	for _, c := range cards {
		if c.GetID() == c.CardID {
			card = c
			break
		}
	}
	if card == nil {
		return fmt.Errorf("card not found")
	}
	// 根据卡牌类型和目标ID获取目标对象
	var target interface{}
	switch card.GetType() {
	case "character":
		target = gameState.GetBoard().GetCharacterByID(c.TargetID)
	case "location":
		target = gameState.GetBoard().GetLocationByID(c.TargetID)
	case "incident":
		target = gameState.GetBoard().GetIncidentByID(c.TargetID)
	default:
		return fmt.Errorf("unknown card type")
	}
	if target == nil {
		return fmt.Errorf("target not found")
	}
	if !card.CanCastOn(target) {
		return fmt.Errorf("card cannot be cast on target")
	}
	err := card.Cast(target)
	if err != nil {
		return err
	}
	fmt.Printf("使用卡牌 %s 在目标 %s 上\n", card.GetTitle(), c.TargetID)
	return nil
}
