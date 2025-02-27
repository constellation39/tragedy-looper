// engine/internal/controllers/commands/show_roles.go
package commands

import "fmt"

type ShowRolesCommand struct{}

func NewShowRolesCommand() *ShowRolesCommand {
	return &ShowRolesCommand{}
}

func (c *ShowRolesCommand) Type() CommandType { return CmdShowRoles }

func (c *ShowRolesCommand) Validate() error { return nil }

func (c *ShowRolesCommand) RequiredInputs() []string { return nil }

func (c *ShowRolesCommand) Execute(ctx CommandContext) error {
	gameState := ctx.GetGameState()
	currentPlayer := gameState.GetCurrentPlayer()
	roles := currentPlayer.GetRoles()
	fmt.Println("角色:")
	for _, role := range roles {
		fmt.Printf("  %s - %s\n", role.GetTitle(), role.GetDescription())
	}
	return nil
}
