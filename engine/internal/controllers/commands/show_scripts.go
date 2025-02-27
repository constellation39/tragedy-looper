// engine/internal/controllers/commands/show_scripts.go
package commands

import "fmt"

type ShowScriptsCommand struct{}

func NewShowScriptsCommand() *ShowScriptsCommand {
	return &ShowScriptsCommand{}
}

func (c *ShowScriptsCommand) Type() CommandType { return CmdShowScripts }

func (c *ShowScriptsCommand) Validate() error { return nil }

func (c *ShowScriptsCommand) RequiredInputs() []string { return nil }

func (c *ShowScriptsCommand) Execute(ctx CommandContext) error {
	gameState := ctx.GetGameState()
	scripts := gameState.GetScripts()
	fmt.Println("脚本:")
	for _, script := range scripts {
		fmt.Printf("  %s - %s\n", script.GetTitle(), script.GetDescription())
	}
	return nil
}
