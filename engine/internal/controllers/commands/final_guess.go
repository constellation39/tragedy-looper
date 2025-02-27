package commands

import (
	"fmt"
)

type FinalGuessCommand struct{}

func NewFinalGuessCommand() *FinalGuessCommand {
	return &FinalGuessCommand{}
}

func (c *FinalGuessCommand) Type() CommandType { return CmdFinalGuess }

func (c *FinalGuessCommand) Validate() error {
	return nil
}

func (c *FinalGuessCommand) RequiredInputs() []string {
	return []string{}
}

func (c *FinalGuessCommand) Execute(ctx CommandContext) error {
	gameState := ctx.GameState

	// Check if the current player is a protagonist
	isProtagonist := false
	for _, protagonist := range gameState.Protagonists {
		if protagonist.ID == ctx.CurrentPlayer.GetID() {
			isProtagonist = true
			break
		}
	}

	if !isProtagonist {
		return fmt.Errorf("只有主角玩家可以进行最终猜测")
	}

	// Check if final guess is allowed at this point
	if gameState.CurrentLoop < gameState.Script.MaxLoops-1 {
		return fmt.Errorf("只能在最后一个循环进行最终猜测")
	}

	// Mark that a guess has been made
	gameState.GuessMade = true

	fmt.Println("请输入你的最终猜测:")
	fmt.Println("1. 主要剧情")
	fmt.Println("2. 子剧情")
	fmt.Println("3. 角色身份")

	// In a real implementation, this would gather input from the user
	// and process the final guess

	return nil
}
