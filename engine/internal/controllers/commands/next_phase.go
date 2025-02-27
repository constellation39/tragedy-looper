package commands

import (
	"fmt"
)

type NextPhaseCommand struct{}

func NewNextPhaseCommand() *NextPhaseCommand {
	return &NextPhaseCommand{}
}

func (c *NextPhaseCommand) Type() CommandType { return CmdNextPhase }

func (c *NextPhaseCommand) Validate() error {
	return nil
}

func (c *NextPhaseCommand) RequiredInputs() []string {
	return []string{}
}

func (c *NextPhaseCommand) Execute(ctx CommandContext) error {
	gameState := ctx.GameState

	// Define the phase progression
	phases := []string{
		string(gameState.CurrentPhase),
		// Here we'd include all possible phases in order
		// For example: "DayStart", "Mastermind", "Protagonists", "Incident", "DayEnd"
	}

	// Find current phase index
	currentIndex := 0
	for i, phase := range phases {
		if phase == string(gameState.CurrentPhase) {
			currentIndex = i
			break
		}
	}

	// Move to next phase or day
	if currentIndex < len(phases)-1 {
		// Go to next phase
		nextPhase := phases[currentIndex+1]
		gameState.CurrentPhase = nextPhase
		fmt.Printf("进入下一阶段: %s\n", nextPhase)
	} else {
		// End of day, go to next day or loop
		gameState.CurrentDay++
		if gameState.CurrentDay > gameState.Script.DaysPerLoop {
			gameState.CurrentDay = 1
			gameState.CurrentLoop++
			if gameState.CurrentLoop >= gameState.Script.MaxLoops {
				gameState.IsGameOver = true
				fmt.Println("游戏结束，已达到最大循环数")
				return nil
			}
			fmt.Printf("进入下一循环: %d\n", gameState.CurrentLoop)
		} else {
			fmt.Printf("进入下一天: %d\n", gameState.CurrentDay)
		}
		gameState.CurrentPhase = phases[0] // Reset to first phase
	}

	return nil
}
