// engine/internal/controllers/commands/start_game.go
package commands

import (
    "fmt"
    "tragedy-looper/engine/internal/models"
)

type StartGameCommand struct{}

func NewStartGameCommand() *StartGameCommand {
    return &StartGameCommand{}
}

func (c *StartGameCommand) Type() CommandType { return CmdStartGame }

func (c *StartGameCommand) Validate() error { return nil }

func (c *StartGameCommand) RequiredInputs() []string { return nil }

func (c *StartGameCommand) Execute(ctx CommandContext) error {
    // 重置游戏状态
    gameState := ctx.GameState
    gameState.CurrentLoop = 1
    gameState.CurrentDay = 1
    gameState.CurrentPhase = models.DayPhaseMorning
    gameState.IsGameOver = false
    gameState.WinnerType = ""
    
    // 初始化游戏数据（根据实际需要补充具体初始化逻辑）
    fmt.Println("=== 新游戏开始 ===")
    fmt.Println("当前循环:", gameState.CurrentLoop)
    fmt.Println("当前天数:", gameState.CurrentDay)
    fmt.Println("当前阶段:", gameState.CurrentPhase)
    return nil
}
