package commands

import (
	"fmt"
)

type GoodwillCommand struct {
	CharacterName string
	AbilityID     string
}

func NewGoodwillCommand(characterName, abilityID string) *GoodwillCommand {
	return &GoodwillCommand{
		CharacterName: characterName,
		AbilityID:     abilityID,
	}
}

func (c *GoodwillCommand) Type() CommandType { return CmdUseGoodwill }

func (c *GoodwillCommand) Validate() error {
	if c.CharacterName == "" {
		return fmt.Errorf("角色名称不能为空")
	}
	if c.AbilityID == "" {
		return fmt.Errorf("能力ID不能为空")
	}
	return nil
}

func (c *GoodwillCommand) RequiredInputs() []string {
	return []string{"characterName", "abilityID"}
}

func (c *GoodwillCommand) Execute(ctx CommandContext) error {
	//gameState := ctx.GameState
	//
	//// Find the character
	//character := gameState.Character(c.CharacterName)
	//if character == nil {
	//	return fmt.Errorf("找不到角色: %s", c.CharacterName)
	//}
	//
	//// Check if the ability exists
	//ability := character.GetAbility(c.AbilityID)
	//if ability == nil {
	//	return fmt.Errorf("角色没有此能力: %s", c.AbilityID)
	//}
	//
	//// Check if player has enough goodwill with the character
	//if character.Goodwill() < ability.GoodwillCost {
	//	return fmt.Errorf("好感度不足，需要 %d 点", ability.GoodwillCost)
	//}
	//
	//// Use the ability
	//character.UseGoodwill(ability.GoodwillCost)
	//fmt.Printf("使用了 %s 的能力: %s\n", c.CharacterName, ability.Name)
	//
	//// Apply ability effects
	//// This would need to be expanded based on the ability implementation
	//if ability.Effect != nil {
	//	ability.Effect(character, gameState)
	//}

	return nil
}
