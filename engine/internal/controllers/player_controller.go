package controllers

import (
	"tragedy-looper/engine/internal/models"
)

type PlayerController struct {
	gameState *models.GameState
}

func NewPlayerController(gs *models.GameState) *PlayerController {
	return &PlayerController{gameState: gs}
}

func (pc *PlayerController) HandleMastermindActions(m *models.Mastermind) error {
	// 简化为直接验证
	return m.PlaceActionCards(pc.gameState)
}

func (pc *PlayerController) HandleProtagonistActions(protagonist *models.Protagonist) error {
	// 简化为直接验证
	return protagonist.PlaceActionCards(pc.gameState)
}

func (pc *PlayerController) HandleProtagonistsActions(protagonists models.Protagonists) error {
	// 简化为遍历验证
	for _, p := range protagonists {
		if err := p.PlaceActionCards(pc.gameState); err != nil {
			return err
		}
	}
	return nil
}
