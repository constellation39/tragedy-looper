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

// Validate 验证命令参数
func (c *GoodwillCommand) Validate() error {
	if c.CharacterName == "" {
		return fmt.Errorf("角色名称不能为空")
	}
	if c.AbilityID == "" {
		return fmt.Errorf("能力ID不能为空")
	}
	// Target可以为空，因为它可能由后续的selectChar或selectLoc命令设置
	return nil
}

// RequiredInputs 返回命令需要的输入描述
func (c *GoodwillCommand) RequiredInputs() []string {
	inputs := []string{
		"角色名称: 使用好感度能力的角色",
		"能力ID: 要使用的好感度能力标识符",
	}
	
	if c.Target == "" {
		inputs = append(inputs, "目标: 需要使用selectChar或selectLoc命令来选择能力目标")
	}
	
	return inputs
}

func (c *GoodwillCommand) Execute(ctx CommandContext) error {
	// 先验证参数
	if err := c.Validate(); err != nil {
		return err
	}
	
	if c.Target == "" {
		return fmt.Errorf("目标不能为空，请先选择目标")
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
