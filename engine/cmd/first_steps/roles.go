package first_steps

import (
	"tragedy-looper/engine/internal/models"
)

// 角色类型定义
const (
	// KeyPerson - 关键人物,其死亡会导致玩家立即失败并结束循环
	KeyPerson models.RoleType = "KeyPerson"
	// Killer - 杀手,可以杀死关键人物或主角
	Killer models.RoleType = "Killer"
	// Brain - 黑幕,可以增加谜团指示物
	Brain models.RoleType = "Brain"
	// Cultist - 异教徒,可以无视"禁止谜团"效果
	Cultist models.RoleType = "Cultist"
	// Friend - 密友,其死亡会在循环结束时导致失败
	Friend models.RoleType = "Friend"
	// ConspiracyTheorist - 造谣者,可以增加疑心指示物
	ConspiracyTheorist models.RoleType = "ConspiracyTheorist"
	// SerialKiller - 连环杀手,当与另一个角色单独在一起时会杀死对方
	SerialKiller models.RoleType = "SerialKiller"
	// Curmudgeon - 暴徒,没有特殊能力
	Curmudgeon models.RoleType = "Curmudgeon"
)

// KeyPersonRoleAbility represents the ability of the KeyPerson role
type KeyPersonRoleAbility struct{}

func (roleAbility *KeyPersonRoleAbility) RoleType() models.RoleType {
	return KeyPerson
}

func (roleAbility *KeyPersonRoleAbility) IsTriggerable(gameState *models.GameState, target models.RoleAbilityTarget) (bool, error) {
	return true, nil
}
func (roleAbility *KeyPersonRoleAbility) Execute(gameState *models.GameState, target models.RoleAbilityTarget) error {
	return nil
}

func (roleAbility *KeyPersonRoleAbility) GetTiming() models.RoleAbilityTiming {
	return models.RoleTimingCharacterDeath
}

func (roleAbility *KeyPersonRoleAbility) GetMandatory() models.GoodwillRefusal {
	return models.GoodwillRefusalMust
}

// KillerAbility represents the Killer's ability to kill the Protagonists
type KillerAbility struct{}

func (roleAbility *KillerAbility) RoleType() models.RoleType {
	return Killer
}

func (roleAbility *KillerAbility) IsTriggerable(gameState *models.GameState, target models.RoleAbilityTarget) (bool, error) {
	return true, nil
}

func (roleAbility *KillerAbility) Execute(gameState *models.GameState, target models.RoleAbilityTarget) error {
	// Kill the Protagonists
	return nil
}

func (roleAbility *KillerAbility) GetTiming() models.RoleAbilityTiming {
	return models.RoleTimingDayEnd
}

func (roleAbility *KillerAbility) GetMandatory() models.GoodwillRefusal {
	return models.GoodwillRefusalOptional
}

// BrainAbility represents the Brain's ability to add Intrigue
type BrainAbility struct{}

func (roleAbility *BrainAbility) RoleType() models.RoleType {
	return Brain
}

func (roleAbility *BrainAbility) IsTriggerable(gameState *models.GameState, target models.RoleAbilityTarget) (bool, error) {
	return true, nil
}

func (roleAbility *BrainAbility) Execute(gameState *models.GameState, target models.RoleAbilityTarget) error {
	return nil
}

func (roleAbility *BrainAbility) GetTiming() models.RoleAbilityTiming {
	return models.RoleTimingMastermind
}

func (roleAbility *BrainAbility) GetMandatory() models.GoodwillRefusal {
	return models.GoodwillRefusalOptional
}

// FriendDeathCheckAbility represents the Friend's ability causing Protagonists to lose if dead at loop end
type FriendDeathCheckAbility struct{}

func (roleAbility *FriendDeathCheckAbility) RoleType() models.RoleType {
	return Friend
}

func (roleAbility *FriendDeathCheckAbility) IsTriggerable(gameState *models.GameState, target models.RoleAbilityTarget) (bool, error) {
	return true, nil
}

func (roleAbility *FriendDeathCheckAbility) Execute(gameState *models.GameState, target models.RoleAbilityTarget) error {
	return nil
}

func (roleAbility *FriendDeathCheckAbility) GetTiming() models.RoleAbilityTiming {
	return models.RoleTimingLoopEnd
}

func (roleAbility *FriendDeathCheckAbility) GetMandatory() models.GoodwillRefusal {
	return models.GoodwillRefusalMust
}

// FriendGoodwillAbility represents the Friend's ability to gain Goodwill if role is revealed at loop start
type FriendGoodwillAbility struct{}

func (roleAbility *FriendGoodwillAbility) RoleType() models.RoleType {
	return Friend
}

func (roleAbility *FriendGoodwillAbility) IsTriggerable(gameState *models.GameState, target models.RoleAbilityTarget) (bool, error) {
	return true, nil
}

func (roleAbility *FriendGoodwillAbility) Execute(gameState *models.GameState, target models.RoleAbilityTarget) error {
	return nil
}

func (roleAbility *FriendGoodwillAbility) GetTiming() models.RoleAbilityTiming {
	return models.RoleTimingLoopStart
}

func (roleAbility *FriendGoodwillAbility) GetMandatory() models.GoodwillRefusal {
	return models.GoodwillRefusalMust
}

// ConspiracyTheoristAbility represents the Conspiracy Theorist's ability to add Paranoia
type ConspiracyTheoristAbility struct{}

func (roleAbility *ConspiracyTheoristAbility) RoleType() models.RoleType {
	return ConspiracyTheorist
}

func (roleAbility *ConspiracyTheoristAbility) IsTriggerable(gameState *models.GameState, target models.RoleAbilityTarget) (bool, error) {
	return true, nil
}

func (roleAbility *ConspiracyTheoristAbility) Execute(gameState *models.GameState, target models.RoleAbilityTarget) error {
	return nil
}

func (roleAbility *ConspiracyTheoristAbility) GetTiming() models.RoleAbilityTiming {
	return models.RoleTimingMastermind
}

func (roleAbility *ConspiracyTheoristAbility) GetMandatory() models.GoodwillRefusal {
	return models.GoodwillRefusalOptional
}

// SerialKillerAbility represents the Serial Killer's ability to kill when alone with another character
type SerialKillerAbility struct{}

func (roleAbility *SerialKillerAbility) RoleType() models.RoleType {
	return SerialKiller
}

func (roleAbility *SerialKillerAbility) IsTriggerable(gameState *models.GameState, target models.RoleAbilityTarget) (bool, error) {
	return true, nil
}

func (roleAbility *SerialKillerAbility) Execute(gameState *models.GameState, target models.RoleAbilityTarget) error {
	return nil
}

func (roleAbility *SerialKillerAbility) GetTiming() models.RoleAbilityTiming {
	return models.RoleTimingDayEnd
}

func (roleAbility *SerialKillerAbility) GetMandatory() models.GoodwillRefusal {
	return models.GoodwillRefusalMust
}

// CurmudgeonRole represents the Curmudgeon role with no special abilities
type CurmudgeonRole struct{}

func (roleAbility *CurmudgeonRole) RoleType() models.RoleType {
	return Curmudgeon
}

func (roleAbility *CurmudgeonRole) IsTriggerable(gameState *models.GameState, target models.RoleAbilityTarget) (bool, error) {
	// No abilities
	return false, nil
}

func (roleAbility *CurmudgeonRole) Execute(gameState *models.GameState, target models.RoleAbilityTarget) error {
	// No abilities
	return nil
}

func (roleAbility *CurmudgeonRole) GetTiming() models.RoleAbilityTiming {
	return models.RoleTimingAlways
}

func (roleAbility *CurmudgeonRole) GetMandatory() models.GoodwillRefusal {
	return models.GoodwillRefusalMust
}
