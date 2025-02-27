package commands

import (
	"fmt"
)

type StatusCommand struct {
	Target string
}

func NewStatusCommand(target string) *StatusCommand {
	return &StatusCommand{Target: target}
}

func (c *StatusCommand) Type() CommandType { return CmdStatus }

func (c *StatusCommand) Validate() error {
	if c.Target == "" {
		return fmt.Errorf("target is required")
	}
	return nil
}

func (c *StatusCommand) RequiredInputs() []string {
	return []string{"target"}
}

func (c *StatusCommand) Execute(ctx CommandContext) error {
	gameState := ctx.GameState

	// Check if target is a character
	character := gameState.Character(c.Target)
	if character != nil {
		fmt.Printf("角色: %s\n", character.Name)
		fmt.Printf("位置: %s\n", character.Location())
		fmt.Printf("存活状态: %t\n", character.IsAlive())
		fmt.Printf("好感度: %d/%d\n", character.Goodwill(), character.GoodwillLimit)
		fmt.Printf("偏执度: %d/%d\n", character.Paranoia(), character.ParanoiaLimit)
		fmt.Printf("阴谋指示物: %d\n", character.Intrigue())
		return nil
	}

	// Check if target is a location
	locType := c.Target
	location := gameState.Location(locType)
	if location != nil {
		fmt.Printf("位置: %s\n", location.LocationType)
		fmt.Printf("阴谋指示物: %d\n", location.Intrigue())
		fmt.Println("当前角色:")
		for charName := range location.Characters {
			fmt.Printf("- %s\n", charName)
		}
		return nil
	}

	return fmt.Errorf("找不到目标: %s", c.Target)
}
