// engine/internal/controllers/commands/set_attribute.go
package commands

import (
	"fmt"
)

type SetAttributeCommand struct {
	TargetType string
	TargetID   string
	Attribute  string
	Value      string
}

func NewSetAttributeCommand(targetType, targetID, attribute, value string) *SetAttributeCommand {
	return &SetAttributeCommand{
		TargetType: targetType,
		TargetID:   targetID,
		Attribute:  attribute,
		Value:      value,
	}
}

func (c *SetAttributeCommand) Type() CommandType { return CmdSetAttribute }

func (c *SetAttributeCommand) Validate() error {
	if c.TargetType == "" {
		return fmt.Errorf("target type is required")
	}
	if c.TargetID == "" {
		return fmt.Errorf("target ID is required")
	}
	if c.Attribute == "" {
		return fmt.Errorf("attribute is required")
	}
	if c.Value == "" {
		return fmt.Errorf("value is required")
	}
	return nil
}

func (c *SetAttributeCommand) RequiredInputs() []string {
	return []string{"target_type", "target_id", "attribute", "value"}
}

func (c *SetAttributeCommand) Execute(ctx CommandContext) error {
	gameState := ctx.GetGameState()
	switch c.TargetType {
	case "character":
		character := gameState.GetBoard().GetCharacterByID(c.TargetID)
		if character == nil {
			return fmt.Errorf("character not found")
		}
		// 根据属性名称设置属性
		switch c.Attribute {
		case "goodwill":
			// 设置角色好感度
			goodwill := 0
			fmt.Sscanf(c.Value, "%d", &goodwill)
			character.SetGoodwill(goodwill)
			fmt.Printf("设置角色 %s 的好感度为 %d\n", character.GetName(), goodwill)
		case "paranoia":
			// 设置角色疑心
			paranoia := 0
			fmt.Sscanf(c.Value, "%d", &paranoia)
			character.SetParanoia(paranoia)
			fmt.Printf("设置角色 %s 的疑心为 %d\n", character.GetName(), paranoia)
		default:
			return fmt.Errorf("unknown attribute: %s", c.Attribute)
		}
	case "location":
		location := gameState.GetBoard().GetLocationByID(c.TargetID)
		if location == nil {
			return fmt.Errorf("location not found")
		}
		// 根据属性名称设置属性
		switch c.Attribute {
		case "intrigue":
			// 设置地点阴谋值
			intrigue := 0
			fmt.Sscanf(c.Value, "%d", &intrigue)
			location.SetIntrigue(intrigue)
			fmt.Printf("设置地点 %s 的阴谋值为 %d\n", location.GetTitle(), intrigue)
		default:
			return fmt.Errorf("unknown attribute: %s", c.Attribute)
		}
	default:
		return fmt.Errorf("unknown target type: %s", c.TargetType)
	}
	return nil
}
