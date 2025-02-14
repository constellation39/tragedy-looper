package models

// RoleType 定义角色类型
type RoleType string

// RoleAbilityTiming 表示能力的触发时机
type RoleAbilityTiming string

const (
	// RoleTimingDayStart 在日初阶段触发的能力
	RoleTimingDayStart RoleAbilityTiming = "Day start"
	// RoleTimingDayEnd 在日结阶段触发的能力
	RoleTimingDayEnd RoleAbilityTiming = "Day end"
	// RoleTimingLoopEnd 在循环结束时触发的能力
	RoleTimingLoopEnd RoleAbilityTiming = "Loop end"
	// RoleTimingLoopStart 在循环开始时触发的能力
	RoleTimingLoopStart RoleAbilityTiming = "Loop start"
	// RoleTimingCardResolve 在卡牌结算时触发的能力
	RoleTimingCardResolve RoleAbilityTiming = "Card resolve"
	// RoleTimingMastermind 在主谋能力阶段触发的能力
	RoleTimingMastermind RoleAbilityTiming = "Mastermind ability"
	// RoleTimingCharacterDeath 在角色死亡时触发的能力
	RoleTimingCharacterDeath RoleAbilityTiming = "Character death"
	// RoleTimingGoodwillUse 使用Goodwill能力时的时机
	RoleTimingGoodwillUse RoleAbilityTiming = "Goodwill use"
	// RoleTimingIncidentTrigger 事件触发时的能力时机
	RoleTimingIncidentTrigger RoleAbilityTiming = "Incident trigger"
	// RoleTimingAlways 始终生效的能力
	RoleTimingAlways RoleAbilityTiming = "Always"
)

type GoodwillRefusal string

const (
	GoodwillRefusalMust      GoodwillRefusal = "Must"      // 不拒绝
	GoodwillRefusalOptional  GoodwillRefusal = "Optional"  // 可以选拒绝
	GoodwillRefusalMandatory GoodwillRefusal = "Mandatory" // 必须拒绝
)

type Role struct {
	Type      RoleType
	Name      string
	Abilities []RoleAbility
}

type RoleAbilityTarget interface {
}

type RoleAbility interface {
	RoleType() RoleType
	IsTriggerable(gameState *GameState, target RoleAbilityTarget) (bool, error)
	Execute(gameState *GameState, target RoleAbilityTarget) error
	GetTiming() RoleAbilityTiming
	GetMandatory() GoodwillRefusal
}

type RolePerson struct {
}

const RolePersonType RoleType = "RolePerson"

func (r *RolePerson) RoleType() RoleType {
	return RolePersonType
}

func (r *RolePerson) IsTriggerable(gameState *GameState, target RoleAbilityTarget) (bool, error) {
	return false, nil
}

func (r *RolePerson) Execute(gameState *GameState, target RoleAbilityTarget) error {
	return nil
}

func (r *RolePerson) GetTiming() RoleAbilityTiming {
	return RoleTimingGoodwillUse
}

func (r *RolePerson) GetMandatory() GoodwillRefusal {
	return GoodwillRefusalMust
}
