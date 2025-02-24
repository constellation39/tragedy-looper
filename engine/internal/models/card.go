package models

import "fmt"

// CardType 表示卡牌类型的枚举
type CardType string

const (
	MovementType       CardType = "Movement"       // 移动卡
	ForbidMovementType CardType = "ForbidMovement" // 禁止移动卡
	IntrigueType       CardType = "Intrigue"       // 阴谋卡
	ForbidIntrigueType CardType = "ForbidIntrigue" // 禁止阴谋卡
	ParanoiaType       CardType = "Paranoia"       // 不安卡
	ForbidParanoiaType CardType = "ForbidParanoia" // 禁止不安卡
	GoodwillType       CardType = "Goodwill"       // 好感卡
	ForbidGoodwillType CardType = "ForbidGoodwill" // 禁止好感卡
)

type TargetType interface {
	Intrigue() int
	SetIntrigue(int)

	Paranoia() int
	SetParanoia(int)

	Goodwill() int
	SetGoodwill(int)

	Location() LocationType
	ToLocation(*Board, MovementDirection)
}

// Card 接口定义了所有卡牌的基本行为
type Card interface {
	Owner() Player                        // 获取卡牌的所有者
	Id() string                           // 获取卡牌唯一标识
	Type() CardType                       // 获取卡牌类型
	Priority() int                        // 获取卡牌优先级
	IsOncePerLoop() bool                  // 是否为每循环一次的卡牌
	IsValidTarget(target TargetType) bool // 检查目标是否有效
	State() *CardState                    // 获取卡牌当前状态
	Target() TargetType                   // 卡牌目标
	SetTarget(target TargetType) error    // 设置卡牌目标
	Reveal()                              // 翻开卡牌
}

// CardState 表示卡牌的动态状态
type CardState struct {
	owner    Player     // 卡牌所有者
	faceDown bool       // 是否面朝下
	target   TargetType // 卡牌目标（角色或位置）
	used     bool       // 是否已使用（用于每循环一次的卡牌）
}

// BaseCardData 包含卡牌的静态信息
type BaseCardData struct {
	id          string   // 卡牌唯一标识
	cardType    CardType // 卡牌类型
	priority    int      // 优先级
	oncePerLoop bool     // 是否为每循环一次的卡牌
}

// BaseCard 组合了静态数据和动态状态
type BaseCard struct {
	data  BaseCardData // 静态数据
	state CardState    // 动态状态
}

func (c *BaseCard) Owner() Player {
	return c.state.owner
}

func (c *BaseCard) Id() string {
	return c.data.id
}

func (c *BaseCard) Type() CardType {
	return c.data.cardType
}

func (c *BaseCard) Priority() int {
	return c.data.priority
}

func (c *BaseCard) IsOncePerLoop() bool {
	return c.data.oncePerLoop
}

func (c *BaseCard) State() *CardState {
	return &c.state
}

func (c *BaseCard) Target() TargetType {
	return c.state.target
}

// Reveal 翻开卡牌
func (c *BaseCard) Reveal() {
	c.state.faceDown = false
}

// ReturnToHand 将卡牌返回手牌
func ReturnToHand(card Card) error {
	return card.Owner().RecycleCards(card)
}

// NewBaseCard 创建一个新的基础卡牌
func NewBaseCard(data BaseCardData, owner Player) BaseCard {
	return BaseCard{
		data: data,
		state: CardState{
			owner:    owner,
			faceDown: true,
			target:   nil,
			used:     false,
		},
	}
}

// MovementDirection defines the possible movement directions
type MovementDirection string

const (
	HorizontalMovement MovementDirection = "Horizontal" // 左右移动
	VerticalMovement   MovementDirection = "Vertical"   // 上下移动
	DiagonalMovement   MovementDirection = "diagonal"   // 斜向移动
)

type MovementCard struct {
	BaseCard
	Direction MovementDirection // 移动方向
}

// NewMovementCard 创建新的移动卡
func NewMovementCard(owner Player, direction MovementDirection, oncePerLoop bool) *MovementCard {
	return &MovementCard{
		BaseCard: NewBaseCard(BaseCardData{
			id:          "move",
			cardType:    MovementType,
			priority:    2,
			oncePerLoop: oncePerLoop,
		}, owner),
		Direction: direction,
	}
}

// SetTarget 设置移动卡的目标
func (c *MovementCard) SetTarget(target TargetType) error {
	c.state.target = target
	return nil
}

// IsValidTarget 检查目标是否有效
func (c *MovementCard) IsValidTarget(target TargetType) bool {
	return true
}

// IntrigueCard 情报卡实现
type IntrigueCard struct {
	BaseCard
	Value int // 情报值
}

// NewIntrigueCard 创建新的情报卡
func NewIntrigueCard(owner Player, value int, oncePerLoop bool) *IntrigueCard {
	return &IntrigueCard{
		BaseCard: NewBaseCard(BaseCardData{
			id:          fmt.Sprintf("intrigue_%d", value),
			cardType:    IntrigueType,
			priority:    4,
			oncePerLoop: oncePerLoop,
		}, owner),
		Value: value,
	}
}

func (i *IntrigueCard) IsValidTarget(target TargetType) bool {
	return true
}

func (i *IntrigueCard) SetTarget(target TargetType) error {
	i.state.target = target
	return nil
}

// ParanoiaCard 不安
type ParanoiaCard struct {
	BaseCard
	Value int // 不安值
}

func NewParanoiaCard(owner Player, value int, oncePerLoop bool) *ParanoiaCard {
	return &ParanoiaCard{
		BaseCard: NewBaseCard(BaseCardData{
			id:          fmt.Sprintf("paranoia_%d", value),
			cardType:    ParanoiaType,
			priority:    4,
			oncePerLoop: oncePerLoop,
		}, owner),
		Value: value,
	}
}

func (p *ParanoiaCard) IsValidTarget(target TargetType) bool {
	return true
}

func (p *ParanoiaCard) SetTarget(target TargetType) error {
	p.state.target = target
	return nil
}

// GoodwillCard 好感卡实现
type GoodwillCard struct {
	BaseCard
	Value int // 好感值
}

// NewGoodwillCard 创建新的好感卡
func NewGoodwillCard(owner Player, value int, oncePerLoop bool) *GoodwillCard {
	return &GoodwillCard{
		BaseCard: NewBaseCard(BaseCardData{
			id:          fmt.Sprintf("goodwill_%d", value),
			cardType:    GoodwillType,
			priority:    4,
			oncePerLoop: oncePerLoop,
		}, owner),
		Value: value,
	}
}

// IsValidTarget 检查目标是否有效
func (c *GoodwillCard) IsValidTarget(target TargetType) bool {
	return true
}

// SetTarget 设置好感卡的目标
func (c *GoodwillCard) SetTarget(target TargetType) error {
	c.state.target = target
	return nil
}

// ForbidMovementCard 禁止移动卡实现
type ForbidMovementCard struct {
	BaseCard
}

// NewForbidMovementCard 创建新的禁止移动卡
func NewForbidMovementCard(owner Player, oncePerLoop bool) *ForbidMovementCard {
	return &ForbidMovementCard{
		BaseCard: NewBaseCard(BaseCardData{
			id:          "forbid_movement",
			cardType:    ForbidMovementType,
			priority:    1, // 根据规则，Forbid Movement cards先结算
			oncePerLoop: oncePerLoop,
		}, owner),
	}
}

// IsValidTarget 检查目标是否有效
func (c *ForbidMovementCard) IsValidTarget(target TargetType) bool {
	return true
}

// SetTarget 设置禁止移动卡的目标
func (c *ForbidMovementCard) SetTarget(target TargetType) error {
	c.state.target = target
	return nil
}

// ForbidIntrigueCard 禁止情报卡实现
type ForbidIntrigueCard struct {
	BaseCard
}

// NewForbidIntrigueCard 创建新的禁止情报卡
func NewForbidIntrigueCard(owner Player, oncePerLoop bool) *ForbidIntrigueCard {
	return &ForbidIntrigueCard{
		BaseCard: NewBaseCard(BaseCardData{
			id:          "forbid_intrigue",
			cardType:    ForbidIntrigueType,
			priority:    3, // 在Movement cards之后，其他卡之前
			oncePerLoop: oncePerLoop,
		}, owner),
	}
}

// IsValidTarget 检查目标是否有效
func (c *ForbidIntrigueCard) IsValidTarget(target TargetType) bool {
	return true
}

// SetTarget 设置禁止情报卡的目标
func (c *ForbidIntrigueCard) SetTarget(target TargetType) error {
	c.state.target = target
	return nil
}

// ForbidParanoiaCard 禁止不安卡实现
type ForbidParanoiaCard struct {
	BaseCard
}

func NewForbidParanoiaCard(owner Player, oncePerLoop bool) *ForbidParanoiaCard {
	return &ForbidParanoiaCard{
		BaseCard: NewBaseCard(BaseCardData{
			id:          "forbid_paranoia",
			cardType:    ForbidParanoiaType,
			priority:    3, // 与其他Forbid cards同优先级
			oncePerLoop: oncePerLoop,
		}, owner),
	}
}
func (c *ForbidParanoiaCard) IsValidTarget(target TargetType) bool {
	return true
}

func (c *ForbidParanoiaCard) SetTarget(target TargetType) error {
	c.state.target = target
	return nil
}

// ForbidGoodwillCard 禁止好感卡实现
type ForbidGoodwillCard struct {
	BaseCard
}

// NewForbidGoodwillCard 创建新的禁止好感卡
func NewForbidGoodwillCard(owner Player, oncePerLoop bool) *ForbidGoodwillCard {
	return &ForbidGoodwillCard{
		BaseCard: NewBaseCard(BaseCardData{
			id:          "forbid_goodwill",
			cardType:    ForbidGoodwillType,
			priority:    3, // 与其他Forbid cards同优先级
			oncePerLoop: oncePerLoop,
		}, owner),
	}
}

// IsValidTarget 检查目标是否有效
func (c *ForbidGoodwillCard) IsValidTarget(target TargetType) bool {
	return true
}

// SetTarget 设置禁止好感卡的目标
func (c *ForbidGoodwillCard) SetTarget(target TargetType) error {
	c.state.target = target
	return nil
}
