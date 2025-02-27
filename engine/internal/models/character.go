package models

import "fmt"

// CharacterName 角色名称类型
type CharacterName string

// Character 角色定义
type Character struct {
	*CharacterData
	*CharacterState
}

func (c *Character) Intrigue() int {
	// 返回当前角色的阴谋值
	return c.GetAttribute(IntrigueAttribute)
}

func (c *Character) SetIntrigue(i int) {
	// 设置角色的阴谋值
	c.SetAttribute(IntrigueAttribute, i)
}

func (c *Character) Paranoia() int {
	// 返回当前角色的不安值
	return c.GetAttribute(ParanoiaAttribute)
}

func (c *Character) SetParanoia(i int) {
	// 设置角色的不安值
	c.SetAttribute(ParanoiaAttribute, i)
}

func (c *Character) Goodwill() int {
	// 返回当前角色的好感度
	return c.GetAttribute(GoodwillAttribute)
}

func (c *Character) SetGoodwill(i int) {
	// 设置角色的好感度
	c.SetAttribute(GoodwillAttribute, i)
}

func (c *Character) GetAttribute(attr AttributeType) int {
	return c.CharacterState.Attributes.Get(attr)
}

func (c *Character) SetAttribute(attr AttributeType, value int) {
	switch attr {
	case GoodwillAttribute:
		if value < 0 {
			value = 0
		}
		if value > c.GoodwillLimit {
			value = c.GoodwillLimit
		}
	case ParanoiaAttribute:
		if value < 0 {
			value = 0
		}
		if value > c.ParanoiaLimit {
			value = c.ParanoiaLimit
		}
	case IntrigueAttribute:
		if value < 0 {
			value = 0
		}
	}
	c.CharacterState.Attributes.Set(attr, value)
}

func (c *Character) Location() LocationType {
	// 返回当前角色所在位置
	return c.CharacterState.CurrentLocation
}

func (c *Character) ToLocation(board *Board, movementDirection MovementDirection) {
	// 根据移动方向移动角色到新位置
	currentLoc := board.GetLocation(c.CharacterState.CurrentLocation)
	if currentLoc == nil {
		return
	}

	// 获取目标位置
	nextLoc, err := currentLoc.getNextLocation(movementDirection)
	if err != nil {
		return
	}

	// 从当前位置移除角色
	delete(currentLoc.Characters, c.Name)

	// 更新角色位置
	c.CharacterState.CurrentLocation = nextLoc.LocationType

	// 将角色添加到新位置
	nextLoc.Characters[c.Name] = c
}

type CharacterTag string

// CharacterData 角色静态数据
type CharacterData struct {
	Name                CharacterName // 角色名称
	Tags                []CharacterTag
	StartLocation       LocationType            // 初始位置
	ForbidMovement      []LocationType          // 禁止移动Id() string
	Traits              []CharacterTrait        // 特征
	GoodwillLimit       int                     // 好感度上限
	ParanoiaLimit       int                     // 不安上限
	GoodwillAbilityList []*CharacterAbilityData // 好感度能力
}

func (cd *CharacterData) ExistsTag(needTag CharacterTag) bool {
	for _, tag := range cd.Tags {
		if tag == needTag {
			return true
		}
	}
	return false
}

// CharacterState 角色动态状态
type CharacterState struct {
	CurrentLocation LocationType // 当前位置
	Attributes      Attributes   // 角色属性值
	IsAlive         bool         // 是否存活
	Role            *Role        // 当前角色身份
}

// CharacterTrait 角色特征
type CharacterTrait struct {
	Effects []TraitEffect
}

// TraitEffect 特征效果接口
type TraitEffect interface {
	Apply(*Character, *GameState)
	Remove(*Character, *GameState)
}

// CharacterAbilityData 能力数据
type CharacterAbilityData struct {
	Name         string                 // 能力
	Cost         int                    // 好感度
	CanBeRefused bool                   // 是否可被拒绝
	Effect       func(*GameState) error // 能力效果
}

// NewCharacter 创建新角色
func NewCharacter(data *CharacterData, role *Role) *Character {
	return &Character{
		CharacterData: data,
		CharacterState: &CharacterState{
			CurrentLocation: data.StartLocation,
			IsAlive:         true,
			Attributes: Attributes{
				GoodwillAttribute: 0,
				ParanoiaAttribute: 0,
				IntrigueAttribute: 0,
			},
			Role: role,
		},
	}
}

// GetCurrentLocation 获取角色当前位置
func (c *Character) GetCurrentLocation() LocationType {
	return c.CurrentLocation
}

// IsAtLocation 检查角色是否在指定位置
func (c *Character) IsAtLocation(location LocationType) bool {
	return c.CurrentLocation == location
}

// CanMoveTo 检查是否可以移动到指定位置
func (c *Character) CanMoveTo(location LocationType) bool {
	// 检查是否在禁止移动列表中
	for _, forbidden := range c.ForbidMovement {
		if forbidden == location {
			return false
		}
	}
	return true
}

// MoveTo 移动到指定位置
func (c *Character) MoveTo(location LocationType) error {
	if !c.CanMoveTo(location) {
		return fmt.Errorf("无法移动到该位置: %v", location)
	}
	c.CurrentLocation = location
	return nil
}

// ===== 计数器相关方法 =====

// HasSufficientGoodwill 检查是否有足够的好感度
func (c *Character) HasSufficientGoodwill(cost int) bool {
	return c.Goodwill() >= cost
}

// HasReachedGoodwillLimit 检查是否达到好感度上限
func (c *Character) HasReachedGoodwillLimit() bool {
	return c.Goodwill() >= c.GoodwillLimit
}

// HasReachedParanoiaLimit 检查是否达到不安上限
func (c *Character) HasReachedParanoiaLimit() bool {
	return c.Paranoia() >= c.ParanoiaLimit
}

// ===== 特征相关方法 =====

type TraitType string

// HasTrait 检查是否具有指定特征
func (c *Character) HasTrait(traitType TraitType) bool {
	for _, trait := range c.Traits {
		_ = trait.Effects
	}
	return false
}

// ApplyTraits 应用所有特征效果
func (c *Character) ApplyTraits(gs *GameState) {
	for _, trait := range c.Traits {
		for _, effect := range trait.Effects {
			effect.Apply(c, gs)
		}
	}
}

// RemoveTraits 移除所有特征效果
func (c *Character) RemoveTraits(gs *GameState) {
	for _, trait := range c.Traits {
		for _, effect := range trait.Effects {
			effect.Remove(c, gs)
		}
	}
}

// ===== 角色状态相关方法 =====

// IsAlive 检查角色是否存活
func (c *Character) IsAlive() bool {
	return c.CharacterState.IsAlive
}

// Kill 使角色死亡
func (c *Character) Kill() error {
	if !c.CharacterState.IsAlive {
		return fmt.Errorf("角色已死")
	}
	c.CharacterState.IsAlive = false
	return nil
}

// Revive 使角色复活
func (c *Character) Revive() error {
	if c.CharacterState.IsAlive {
		return fmt.Errorf("角色没死")
	}
	c.CharacterState.IsAlive = true
	return nil
}

// ===== 角色身份相关方法 =====

// Role 获取当前角色身份
func (c *Character) Role() *Role {
	return c.CharacterState.Role
}

// SetRole 设置角色身份
func (c *Character) SetRole(role *Role) {
	c.CharacterState.Role = role
}

// HasRole 检查是否具有指定角色身份
func (c *Character) HasRole(roleType RoleType) bool {
	return c.CharacterState.Role != nil && c.CharacterState.Role.Type == roleType
}

// GetGoodwillAbility 获取好感度能力
func (c *Character) GetGoodwillAbility() []*CharacterAbilityData {
	return c.GoodwillAbilityList
}

// CanUseGoodwillAbility 检查是否可以使用好感度能力
func (c *Character) CanUseGoodwillAbility() bool {
	//return c.GoodwillAbilityList != nil &&
	//c.HasSufficientGoodwill(c.GoodwillAbilityList.Cost)
	return false
}

// UseGoodwillAbility 使用好感度能力
func (c *Character) UseGoodwillAbility(gs *GameState, ability RoleAbility) error {
	if !c.CanUseGoodwillAbility() {
		return fmt.Errorf("无法使用好感度能力")
	}

	//if c.GoodwillAbilityList.CanBeRefused && ability.GetMandatory() == GoodwillRefusalMandatory {
	//	return fmt.Errorf("角色拒绝了能力")
	//}
	//
	//c.GoodwillAbilityList.Effect(gs)
	return nil
}

// ResetState 重置角色状态(新循环开始时)
func (c *Character) ResetState() {
	c.CharacterState = &CharacterState{
		CurrentLocation: c.StartLocation,
		IsAlive:         true,
		Role:            c.CharacterState.Role, // 保持角色身份
		Attributes: Attributes{
			GoodwillAttribute: 0,
			ParanoiaAttribute: 0,
			IntrigueAttribute: 0,
		},
	}
}
