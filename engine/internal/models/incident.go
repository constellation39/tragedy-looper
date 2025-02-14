package models

import (
	"go.uber.org/zap"
)

type IncidentType string

type IncidentEffectTarget interface {
}

type Incident interface {
	Type() IncidentType
	Execute(logger zap.Logger, gameState *GameState, target IncidentEffectTarget) error
	IsTriggerable(logger zap.Logger, gameState *GameState, target IncidentEffectTarget) bool
}
