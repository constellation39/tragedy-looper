package commands

import "fmt"

// GoodwillCommand 使用好感度能力命令
type GoodwillCommand struct {
	CharacterName string
	AbilityID     string
	Target        string
}

func (c *GoodwillCommand) Type() CommandType {
	return CmdUseGoodwill
}

func (c *GoodwillCommand) Validate() error {
	if c.CharacterName == "" || c.AbilityID == "" {
		return fmt.Errorf("character name and ability ID cannot be empty")
	}
	return nil
}

func (c *GoodwillCommand) Execute(ctx CommandContext) error {
	panic("implement me")
}
