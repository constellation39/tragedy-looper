package controllers

import (
	"tragedy-looper/engine/internal/models"
)

var MastermindCLI *CLI
var ProtagonistCLI *CLI

func SetMastermindCLI(cli *CLI) {
	MastermindCLI = cli
}

func SetProtagonistCLI(cli *CLI) {
	ProtagonistCLI = cli
}

type PlayerController struct {
	gameState *models.GameState
	cli       *CLI
}

func NewPlayerController(gs *models.GameState) *PlayerController {
	return &PlayerController{gameState: gs}
}

// SetCLI 设置控制器使用的CLI
func (pc *PlayerController) SetCLI(cli *CLI) {
	pc.cli = cli
}

// HandleMastermindActions 处理幕后主使的行动
func (pc *PlayerController) HandleMastermindActions(m *models.Mastermind) error {
	// 验证手牌数量
	if err := m.PlaceActionCards(pc.gameState); err != nil {
		return err
	}
	
	// 启动交互式卡牌放置阶段
	return pc.cli.StartPlacementPhase(m, pc.gameState)
}

// HandleProtagonistActions 处理主角玩家的行动
func (pc *PlayerController) HandleProtagonistActions(protagonist *models.Protagonist) error {
	// 验证手牌数量
	if err := protagonist.PlaceActionCards(pc.gameState); err != nil {
		return err
	}
	
	// 启动交互式卡牌放置阶段
	return pc.cli.StartPlacementPhase(protagonist, pc.gameState)
}

// HandleProtagonistsActions 处理所有主角玩家的行动
func (pc *PlayerController) HandleProtagonistsActions(protagonists models.Protagonists) error {
	for _, protagonist := range protagonists {
		if err := pc.HandleProtagonistActions(protagonist); err != nil {
			return err
		}
	}
	return nil
}
