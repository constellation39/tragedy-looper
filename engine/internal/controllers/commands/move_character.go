package commands

import (
	"fmt"
)

type MoveCharacterCommand struct {
	CharacterID string
	LocationID  string
}

func NewMoveCharacterCommand(characterID, locationID string) *MoveCharacterCommand {
	return &MoveCharacterCommand{
		CharacterID: characterID,
		LocationID:  locationID,
	}
}

func (c *MoveCharacterCommand) Type() CommandType { return CmdMoveCharacter }

func (c *MoveCharacterCommand) Validate() error {
	if c.CharacterID == "" {
		return fmt.Errorf("character ID is required")
	}
	if c.LocationID == "" {
		return fmt.Errorf("location ID is required")
	}
	return nil
}

func (c *MoveCharacterCommand) RequiredInputs() []string {
	return []string{"character_id", "location_id"}
}

func (c *MoveCharacterCommand) Execute(ctx CommandContext) error {
	//gameState := ctx.GameState
	//board := gameState.GetBoard()
	//character := board.GetCharacterByID(c.CharacterID)
	//location := board.GetLocationByID(c.LocationID)
	//if character == nil {
	//	return fmt.Errorf("character not found")
	//}
	//if location == nil {
	//	return fmt.Errorf("location not found")
	//}
	//if !character.CanMoveTo(location) {
	//	return fmt.Errorf("character cannot move to location")
	//}
	//character.SetLocationID(c.LocationID)
	//fmt.Printf("角色 %s 移动到 %s\n", character.GetName(), location.GetTitle())
	return nil
}
