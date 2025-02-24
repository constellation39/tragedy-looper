package commands

import "fmt"

// PlaceCardCommand 放置卡牌命令
type PlaceCardCommand struct {
	CardID string
	Target string
}

func NewPlaceCardCommand(cardID, target string) *PlaceCardCommand {
	return &PlaceCardCommand{
		CardID: cardID,
		Target: target,
	}
}

func (c *PlaceCardCommand) Type() CommandType {
	return CmdPlaceCard
}

func (c *PlaceCardCommand) Validate() error {
	// 根据知识库规则验证
	// 1. 检查是否在正确的阶段
	// 2. 检查目标位置是否已被其他主角占用
	// 3. 检查卡牌放置规则
	if c.CardID == "" || c.Target == "" {
		return fmt.Errorf("cardID and target cannot be empty")
	}
	return nil
}

func (c *PlaceCardCommand) Execute(ctx CommandContext) error {
	panic("implement me")
}

// UseGoodwillCommand 使用好感度能力命令
type UseGoodwillCommand struct {
	CharacterName string
	AbilityID     string
	Target        string
}

func (c *UseGoodwillCommand) Type() CommandType {
	return CmdUseGoodwill
}

func (c *UseGoodwillCommand) Validate() error {
	if c.CharacterName == "" || c.AbilityID == "" {
		return fmt.Errorf("character name and ability ID cannot be empty")
	}
	return nil
}

func (c *UseGoodwillCommand) Execute(ctx CommandContext) error {
	panic("implement me")
}
