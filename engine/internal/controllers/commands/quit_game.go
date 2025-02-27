// engine/internal/controllers/commands/quit_game.go
package commands

import "fmt"

type QuitGameCommand struct{}

func NewQuitGameCommand() *QuitGameCommand {
	return &QuitGameCommand{}
}

func (c *QuitGameCommand) Type() CommandType { return CmdQuitGame }

func (c *QuitGameCommand) Validate() error { return nil }

func (c *QuitGameCommand) RequiredInputs() []string { return nil }

func (c *QuitGameCommand) Execute(ctx CommandContext) error {
	fmt.Println("退出游戏")
	return nil
}
