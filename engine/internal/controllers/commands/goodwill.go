package commands

import (
	"fmt"
	"strings"
	"tragedy-looper/engine/internal/models"
)

// GoodwillCommand 使用好感度能力命令
type GoodwillCommand struct {
	CharacterName string
	AbilityID     string
	Target        string
}

// NewGoodwillCommand 创建新的好感度能力命令
func NewGoodwillCommand(characterName, abilityID string) *GoodwillCommand {
	return &GoodwillCommand{
		CharacterName: characterName,
		AbilityID:     abilityID,
		Target:        "",
	}
}

func (c *GoodwillCommand) Type() CommandType {
	return CmdUseGoodwill
}

func (c *GoodwillCommand) Validate() error {
	if c.CharacterName == "" || c.AbilityID == "" {
		return fmt.Errorf("character name and ability ID cannot be empty")
	}
	if c.Target == "" {
		return fmt.Errorf("target cannot be empty")
	}
	return nil
}

func (c *GoodwillCommand) Execute(ctx CommandContext) error {
	if err := c.Validate(); err != nil {
		return err
	}

	// 查找角色
	character := ctx.GameState.Character(models.CharacterName(c.CharacterName))
	if character == nil {
		return fmt.Errorf("找不到角色: %s", c.CharacterName)
	}

	// 解析目标
	target := parseGoodwillTarget(ctx.GameState, c.Target)
	if target == nil {
		return fmt.Errorf("无效的目标: %s", c.Target)
	}

	// TODO: 实现好感度能力逻辑
	fmt.Printf("角色 %s 使用好感度能力 %s 对 %s\n", c.CharacterName, c.AbilityID, c.Target)
	return nil
}

// parseGoodwillTarget 解析好感度能力目标
func parseGoodwillTarget(state *models.GameState, targetInput string) models.TargetType {
	parts := strings.SplitN(targetInput, "_", 2)
	if len(parts) != 2 {
		return nil
	}

	switch parts[0] {
	case "char":
		return state.Character(models.CharacterName(parts[1]))
	case "loc":
		return state.Location(models.LocationType(parts[1]))
	}

	return nil
}
