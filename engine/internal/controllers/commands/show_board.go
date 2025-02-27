// engine/internal/controllers/commands/show_board.go
package commands

import (
	"fmt"
	"tragedy-looper/engine/internal/models"
)

type ShowBoardCommand struct{}

func NewShowBoardCommand() *ShowBoardCommand {
	return &ShowBoardCommand{}
}

func (c *ShowBoardCommand) Type() CommandType { return CmdShowBoard }

func (c *ShowBoardCommand) Validate() error { return nil }

func (c *ShowBoardCommand) RequiredInputs() []string { return nil }

func (c *ShowBoardCommand) Execute(ctx CommandContext) error {
	gameState := ctx.GetGameState()
	board := gameState.GetBoard()
	characters := board.GetCharacters()
	locations := board.GetLocations()
	incidents := board.GetIncidents()
	fmt.Println("角色:")
	for _, character := range characters {
		locationTitle := "未知"
		for _, location := range locations {
			if location.GetID() == character.GetLocationID() {
				locationTitle = location.GetTitle()
				break
			}
		}
		fmt.Printf("  %s (%s) - 位置: %s\n", character.GetName(), character.GetTitle(), locationTitle)
	}
	fmt.Println("地点:")
	for _, location := range locations {
		fmt.Printf("  %s (%s)\n", location.GetTitle(), location.GetType())
	}
	fmt.Println("事件:")
	for _, incident := range incidents {
		culpritName := "未知"
		for _, character := range characters {
			if character.GetID() == incident.GetCulpritID() {
				culpritName = character.GetName()
				break
			}
		}
		locationTitle := "未知"
		for _, location := range locations {
			if location.GetID() == incident.GetLocationID() {
				locationTitle = location.GetTitle()
				break
			}
		}
		var days []int
		for i := 1; i <= models.DaysCount; i++ {
			if incident.GetDay() == i {
				days = append(days, i)
			}
		}
		daysStr := ""
		for i, day := range days {
			if i > 0 {
				daysStr += ", "
			}
			daysStr += fmt.Sprintf("%d", day)
		}
		fmt.Printf("  %s - 犯人: %s, 地点: %s, 天数: %s\n", incident.GetTitle(), culpritName, locationTitle, daysStr)
	}
	return nil
}
