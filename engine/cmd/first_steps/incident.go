package first_steps

import (
	"go.uber.org/zap"
	"tragedy-looper/engine/internal/models"
)

const (
	// MurderIncidentType 谋杀事件：与被杀者同一地区的其他角色死亡
	MurderIncidentType models.IncidentType = "MurderIncident"

	// FarawayMurderIncidentType 远程谋杀：带有至少2个阴谋值的角色死亡
	FarawayMurderIncidentType models.IncidentType = "FarawayMurderIncident"

	// SuicideIncidentType 自杀事件：角色因疑神值过高自我了断
	SuicideIncidentType models.IncidentType = "SuicideIncident"

	// HospitalIncidentType 医院事故：医院阴谋值1时医院所有人死亡，阴谋值2时主角死亡
	HospitalIncidentType models.IncidentType = "HospitalIncident"

	// MissingIncidentType 人员失踪：移动角色到任意位置并在该位置放置1个阴谋值
	MissingIncidentType models.IncidentType = "MissingIncident"

	// IncreasingUneaseIncidentType 不安扩散：在任意角色上放置2个疑神值，然后在另一个角色上放置1个阴谋值
	IncreasingUneaseIncidentType models.IncidentType = "IncreasingUneaseIncident"

	// SpreadingIncidentType 阴谋扩散：移除任意角色2个好感值，然后在另一个角色上放置2个好感值
	SpreadingIncidentType models.IncidentType = "SpreadingIncident"
)

type MurderEffect struct {
	CulpritName models.CharacterName
}

// MurderIncident 谋杀事件：角色主动杀死其他角色
type MurderIncident struct{}

func (incident *MurderIncident) Type() models.IncidentType {
	return MurderIncidentType
}

func (incident *MurderIncident) Execute(logger zap.Logger, gameState *models.GameState, target models.IncidentEffectTarget) error {
	//murderEffect, ok := target.(*MurderEffect)
	//if !ok {
	//	return fmt.Errorf("target is not a MurderEffect %s", tools.GetInterfaceType(target))
	//}
	//if murderEffect.CulpritName == "" {
	//	return fmt.Errorf("culprit is nil")
	//}
	//
	//gameState.KillCharacter(murderEffect.CulpritName)
	return nil
}

func (incident *MurderIncident) IsTriggerable(logger zap.Logger, gameState *models.GameState, target models.IncidentEffectTarget) bool {
	//murderEffect, ok := target.(*MurderEffect)
	//if !ok {
	//	logger.DPanic("target is not a MurderEffect",
	//		zap.Error(fmt.Errorf("target is not a MurderEffect %s", tools.GetInterfaceType(target))))
	//	return false
	//}
	//if murderEffect.CulpritName == "" {
	//	logger.DPanic("culprit is nil")
	//	return false
	//}
	//
	//culpritCharacter := gameState.GetCharacter(murderEffect.CulpritName)
	//if culpritCharacter == nil {
	//	logger.DPanic("culprit is nil")
	//	return false
	//}
	//
	//return culpritCharacter.IsAlive() && culpritCharacter.IsParanoia()
	return false
}

// FarawayMurderEffect 远程谋杀效果
type FarawayMurderEffect struct {
	TargetName models.CharacterName
}

// FarawayMurderIncident 远程谋杀：带有至少2个阴谋值的角色死亡
type FarawayMurderIncident struct{}

func (incident *FarawayMurderIncident) Type() models.IncidentType {
	return FarawayMurderIncidentType
}

func (incident *FarawayMurderIncident) Execute(logger zap.Logger, gameState *models.GameState, target models.IncidentEffectTarget) error {
	//effect, ok := target.(*FarawayMurderEffect)
	//if !ok {
	//	return fmt.Errorf("target is not a FarawayMurderEffect %s", tools.GetInterfaceType(target))
	//}
	//
	//// 确保目标角色存在
	//if effect.TargetName == "" {
	//	return fmt.Errorf("target character is nil")
	//}
	//
	//// 执行远程谋杀
	//gameState.KillCharacter(effect.TargetName)
	return nil
}

func (incident *FarawayMurderIncident) IsTriggerable(logger zap.Logger, gameState *models.GameState, target models.IncidentEffectTarget) bool {
	//effect, ok := target.(*FarawayMurderEffect)
	//if !ok {
	//	logger.DPanic("target is not a FarawayMurderEffect",
	//		zap.Error(fmt.Errorf("invalid target type %s", tools.GetInterfaceType(target))))
	//	return false
	//}
	//
	//// 获取目标角色
	//targetCharacter := gameState.GetCharacter(effect.TargetName)
	//if targetCharacter == nil {
	//	logger.DPanic("target character not found")
	//	return false
	//}
	//
	//// 检查目标是否活着且阴谋值至少为2
	//return targetCharacter.IsAlive() && targetCharacter.GetIntrigue() >= 2
	return false
}

// SuicideEffect 自杀事件效果
type SuicideEffect struct {
	CharacterName models.CharacterName
}

// SuicideIncident 自杀事件：角色因疑神值过高自我了断
type SuicideIncident struct{}

func (incident *SuicideIncident) Type() models.IncidentType {
	return SuicideIncidentType
}

func (incident *SuicideIncident) Execute(logger zap.Logger, gameState *models.GameState, target models.IncidentEffectTarget) error {
	//effect, ok := target.(*SuicideEffect)
	//if !ok {
	//	return fmt.Errorf("target is not a SuicideEffect %s", tools.GetInterfaceType(target))
	//}
	//gameState.KillCharacter(effect.CharacterName)
	//return nil
	return nil
}

func (incident *SuicideIncident) IsTriggerable(logger zap.Logger, gameState *models.GameState, target models.IncidentEffectTarget) bool {
	//effect, ok := target.(*SuicideEffect)
	//if !ok {
	//	logger.DPanic("target is not a SuicideEffect", zap.Error(fmt.Errorf("invalid target type")))
	//	return false
	//}
	//character := gameState.GetCharacter(effect.CharacterName)
	//return character != nil && character.IsAlive() && character.GetParanoia() >= 4
	return false
}

// HospitalIncidentEffect 医院事故效果
type HospitalIncidentEffect struct {
	Location models.LocationType
}

// HospitalIncident 医院事故：阴谋值积累引发的特殊事件
type HospitalIncident struct{}

func (incident *HospitalIncident) Type() models.IncidentType {
	return HospitalIncidentType
}

func (incident *HospitalIncident) Execute(logger zap.Logger, gameState *models.GameState, target models.IncidentEffectTarget) error {
	//effect, ok := target.(*HospitalIncidentEffect)
	//if !ok {
	//	return fmt.Errorf("target is not a HospitalIncidentEffect %s", tools.GetInterfaceType(target))
	//}
	//
	//location := gameState.GetLocation(effect.Location)
	//if location.GetIntrigue() >= 2 {
	//	// 杀死所有主角
	//	gameState.KillAllProtagonists()
	//} else if location.GetIntrigue() >= 1 {
	//	// 杀死医院中的所有角色
	//	gameState.KillAllCharactersInLocation(effect.Location)
	//}
	//return nil
	return nil
}

func (incident *HospitalIncident) IsTriggerable(logger zap.Logger, gameState *models.GameState, target models.IncidentEffectTarget) bool {
	//effect, ok := target.(*HospitalIncidentEffect)
	//if !ok {
	//	logger.DPanic("target is not a HospitalIncidentEffect")
	//	return false
	//}
	//location := gameState.GetLocation(effect.Location)
	//return location != nil && location.GetIntrigue() > 0
	return true
}

// MissingEffect 失踪事件效果
type MissingEffect struct {
	CharacterName models.CharacterName
}

// MissingIncident 人员失踪：角色永久退出当前循环
type MissingIncident struct{}

func (incident *MissingIncident) Type() models.IncidentType {
	return MissingIncidentType
}

func (incident *MissingIncident) Execute(logger zap.Logger, gameState *models.GameState, target models.IncidentEffectTarget) error {
	//effect, ok := target.(*MissingEffect)
	//if !ok {
	//	return fmt.Errorf("target is not a MissingEffect %s", tools.GetInterfaceType(target))
	//}
	//gameState.RemoveCharacter(effect.CharacterName)
	//return nil
	return nil
}

func (incident *MissingIncident) IsTriggerable(logger zap.Logger, gameState *models.GameState, target models.IncidentEffectTarget) bool {
	//effect, ok := target.(*MissingEffect)
	//if !ok {
	//	logger.DPanic("target is not a MissingEffect")
	//	return false
	//}
	//character := gameState.GetCharacter(effect.CharacterName)
	//return character != nil && character.IsAlive()
	return false
}

// IncreasingUneaseEffect 不安扩散效果
type IncreasingUneaseEffect struct {
	TargetCharacterName         models.CharacterName
	IntrigueTargetCharacterName models.CharacterName
}

// IncreasingUneaseIncident 不安扩散：提升指定区域的疑神值
type IncreasingUneaseIncident struct{}

func (incident *IncreasingUneaseIncident) Type() models.IncidentType {
	return IncreasingUneaseIncidentType
}

func (incident *IncreasingUneaseIncident) Execute(logger zap.Logger, gameState *models.GameState, target models.IncidentEffectTarget) error {
	//effect, ok := target.(*IncreasingUneaseEffect)
	//if !ok {
	//	return fmt.Errorf("target is not an IncreasingUneaseEffect %s", tools.GetInterfaceType(target))
	//}
	//
	//targetChar := gameState.GetCharacter(effect.TargetCharacterName)
	//targetChar.AddParanoia(2)
	//
	//intrigueChar := gameState.GetCharacter(effect.IntrigueTargetCharacterName)
	//intrigueChar.AddIntrigue(1)
	return nil
}

func (incident *IncreasingUneaseIncident) IsTriggerable(logger zap.Logger, gameState *models.GameState, target models.IncidentEffectTarget) bool {
	//effect, ok := target.(*IncreasingUneaseEffect)
	//if !ok {
	//	logger.DPanic("target is not an IncreasingUneaseEffect")
	//	return false
	//}
	//return gameState.GetCharacter(effect.TargetCharacterName) != nil &&
	//	gameState.GetCharacter(effect.IntrigueTargetCharacterName) != nil
	return true
}

// SpreadingEffect 阴谋扩散效果
type SpreadingEffect struct {
	FromCharacterName models.CharacterName
	ToCharacterName   models.CharacterName
}

// SpreadingIncident 阴谋扩散：提升指定区域的阴谋值
type SpreadingIncident struct{}

func (incident *SpreadingIncident) Type() models.IncidentType {
	return SpreadingIncidentType
}

func (incident *SpreadingIncident) Execute(logger zap.Logger, gameState *models.GameState, target models.IncidentEffectTarget) error {
	//effect, ok := target.(*SpreadingEffect)
	//if !ok {
	//	return fmt.Errorf("target is not a SpreadingEffect %s", tools.GetInterfaceType(target))
	//}
	//
	//fromChar := gameState.GetCharacter(effect.FromCharacterName)
	//fromChar.RemoveGoodwill(2)
	//
	//toChar := gameState.GetCharacter(effect.ToCharacterName)
	//toChar.AddGoodwill(2)
	return nil
}

func (incident *SpreadingIncident) IsTriggerable(logger zap.Logger, gameState *models.GameState, target models.IncidentEffectTarget) bool {
	//effect, ok := target.(*SpreadingEffect)
	//if !ok {
	//	logger.DPanic("target is not a SpreadingEffect")
	//	return false
	//}
	//fromChar := gameState.GetCharacter(effect.FromCharacterName)
	//toChar := gameState.GetCharacter(effect.ToCharacterName)
	//return fromChar != nil && toChar != nil && fromChar.GetGoodwill() >= 2
	return true
}
