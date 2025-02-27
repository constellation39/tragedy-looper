package commands

import (
	"fmt"
	"strings"
	"tragedy-looper/engine/internal/models"
)

// PlaceCardCommand 放置卡牌命令
type PlaceCardCommand struct {
	CardID string
	Target string // 目标将由后续的selectChar或selectLoc命令设置
}

// NewPlaceCardCommand 创建新的放置卡牌命令
func NewPlaceCardCommand(cardID string) *PlaceCardCommand {
	return &PlaceCardCommand{
		CardID: cardID,
		Target: "", // 初始为空，等待后续命令设置
	}
}

func (c *PlaceCardCommand) Type() CommandType {
	return CmdPlaceCard
}

// SetTarget 设置目标
func (c *PlaceCardCommand) SetTarget(target string) {
	c.Target = target
}

func (c *PlaceCardCommand) Execute(ctx CommandContext) error {
	if ctx.CurrentPlayer == nil {
		return fmt.Errorf("当前玩家信息缺失")
	}

	if ctx.GameState == nil {
		return fmt.Errorf("游戏状态信息缺失")
	}

	if c.Target == "" {
		return fmt.Errorf("请先选择目标")
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

	// 检查目标是否有效
	if !card.IsValidTarget(target) {
		return fmt.Errorf("该目标不能放置此卡牌")
	}

	// 设置卡牌目标
	if err := card.SetTarget(target); err != nil {
		return err
	}

	// 执行卡牌放置
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
		return state.Character(models.CharacterName(parts[1]))
	case "loc":
		// 查找位置
		return state.Location(models.LocationType(parts[1]))
	}

	return nil
}
