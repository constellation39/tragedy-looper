package commands

import (
	"fmt"
)

type MakeNoteCommand struct {
	Content string
}

func NewMakeNoteCommand(content string) *MakeNoteCommand {
	return &MakeNoteCommand{Content: content}
}

func (c *MakeNoteCommand) Type() CommandType { return CmdMakeNote }

func (c *MakeNoteCommand) Validate() error {
	if c.Content == "" {
		return fmt.Errorf("笔记内容不能为空")
	}
	return nil
}

func (c *MakeNoteCommand) RequiredInputs() []string {
	return []string{"content"}
}

func (c *MakeNoteCommand) Execute(ctx CommandContext) error {
	// In a real implementation, this would save a note to the player's notes
	// This is a simplified version that just acknowledges the note

	fmt.Printf("已添加笔记: %s\n", c.Content)

	// If the GameState had a notes tracking system, we would add it there
	// For example:
	// ctx.GameState.PlayerNotes[ctx.CurrentPlayer.GetID()] = append(
	//    ctx.GameState.PlayerNotes[ctx.CurrentPlayer.GetID()],
	//    c.Content
	// )

	return nil
}
