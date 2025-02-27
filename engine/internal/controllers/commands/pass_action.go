// engine/internal/controllers/commands/pass_action.go
package commands

import "fmt"

type PassActionCommand struct{}

func NewPassActionCommand() *PassActionCommand {
	return &PassActionCommand{}
}

func (c *PassActionCommand) Type() CommandType { return CmdPassAction }

func (c *PassActionCommand) Validate() error { return nil }

func (c *PassActionCommand) RequiredInputs() []string { return nil }

func (c *PassActionCommand) Execute(ctx CommandContext) error {
	fmt.Println("跳过当前操作")
	return nil
}
