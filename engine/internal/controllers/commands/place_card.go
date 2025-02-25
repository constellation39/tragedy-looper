package commands

import (
	"fmt"
	"strings"
	"tragedy-looper/engine/internal/models"
)

// PlaceCardCommand 放置卡牌命令
type PlaceCardCommand struct {
	CardID string
	Target string
}

// NewPlaceCardCommand 创建新的放置卡牌命令
func NewPlaceCardCommand(cardID, target string) *PlaceCardCommand {
	return &PlaceCardCommand{
		CardID: cardID,
		Target: target,
	}
}

func (c *PlaceCardCommand) Type() CommandType {
	return CmdPlaceCard
}

func (c *PlaceCardCommand) Execute(ctx CommandContext) error {
	if ctx.CurrentPlayer == nil {
		return fmt.Errorf("当前玩家信息缺失")
	}

	if ctx.GameState == nil {
		return fmt.Errorf("游戏状态信息缺失")
	}

	// 查找卡牌
	var card models.Card
	for _, handCard := range ctx.CurrentPlayer.GetHandCards() {
		if handCard.Id() == c.CardID {
			card = handCard
			break
		}
	}

	if card == nil {
		return fmt.Errorf("无效的卡牌ID: %s", c.CardID)
	}

	// 解析目标
	target := parseTarget(ctx.GameState, c.Target)
	if target == nil {
		return fmt.Errorf("无效的目标: %s", c.Target)
	}

	// 执行卡牌放置
	if err := ctx.GameState.Board.SetCard(target, card); err != nil {
		return err
	}

	return ctx.CurrentPlayer.PlaceCards(card)
}

// parseTarget 解析目标字符串为游戏目标对象
func parseTarget(state *models.GameState, targetInput string) models.TargetType {
	parts := strings.SplitN(targetInput, "_", 2)
	if len(parts) != 2 {
		return nil
	}

	switch parts[0] {
	case "char":
		// 查找角色
		for _, c := range state.Script.Characters {
			if c.Name == models.CharacterName(parts[1]) {
				return c
			}
		}
	case "loc":
		// 查找位置
		return state.Board.GetLocation(models.LocationType(parts[1]))
	}

	return nil
}
