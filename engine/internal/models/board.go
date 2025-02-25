package models

import (
	"fmt"
	"go.uber.org/zap"
	"sort"
)

// Board 代表游戏的主要状态
type Board struct {
	logging          *zap.Logger // 添加日志记录器
	characters       []*Character
	locations        map[LocationType]*Location // 所有位置的映射表
	forbiddenActions map[any]CardType           // 被禁止的动作
	actionCards      []Card                     // 在此位置上打出的行动卡
}

// NewBoard 初始化一个新的游戏板
func NewBoard(logging *zap.Logger, characters []*Character) *Board {
	board := &Board{
		logging:    logging,
		characters: characters,
		locations:  nil,
	}
	return board
}

// Reset 初始化游戏板上的所有位置及其方向关系
func (board *Board) Reset() error {
	board.locations = make(map[LocationType]*Location)

	// 初始化所有位置
	hospital := &Location{
		LocationType: LocationHospital,
		Characters:   make(map[CharacterName]*Character),
	}
	city := &Location{
		LocationType: LocationCity,
		Characters:   make(map[CharacterName]*Character),
	}
	school := &Location{
		LocationType: LocationSchool,
		Characters:   make(map[CharacterName]*Character),
	}
	shrine := &Location{
		LocationType: LocationShrine,
		Characters:   make(map[CharacterName]*Character),
	}

	// 记录位置初始化
	board.logging.Debug("All locations have been initialized")

	// 设置地图位置关系
	hospital.right = shrine
	hospital.bottom = city
	hospital.diagonal = school

	city.right = school
	city.top = hospital
	city.diagonal = shrine

	school.top = shrine
	school.left = city
	school.diagonal = hospital

	shrine.right = hospital
	shrine.bottom = school
	shrine.diagonal = city

	board.logging.Debug("The location relationships have been set")

	// 将位置加入到board的locations映射中
	board.locations[LocationHospital] = hospital
	board.locations[LocationCity] = city
	board.locations[LocationSchool] = school
	board.locations[LocationShrine] = shrine

	board.logging.Debug("The locations were added to the game board mapping",
		zap.Reflect("hospital", hospital),
		zap.Reflect("city", city),
		zap.Reflect("school", school),
		zap.Reflect("shrine", shrine))

	// 将角色放置在起始位置
	for _, char := range board.characters {
		loc := board.locations[char.StartLocation]
		loc.Characters[char.Name] = char
		board.logging.Debug("The character has been placed at the starting location",
			zap.String("character", string(char.Name)),
			zap.String("location", string(char.StartLocation)))
	}

	board.logging.Debug("All character positions have been successfully Reset")
	return nil
}

// IsAdjacent 检查两个位置是否相邻
func (l *Location) IsAdjacent(other *Location) bool {
	return l.left == other ||
		l.right == other ||
		l.top == other ||
		l.bottom == other ||
		l.diagonal == other
}

// ResolveActionCards 按照规则顺序处理所有行动卡
func (board *Board) ResolveActionCards(gs *GameState) error {
	board.logging.Debug("Starting to process action cards")

	// 重置禁止动作
	board.forbiddenActions = make(map[any]CardType)
	board.logging.Debug("Forbidden actions have been Reset")

	// 收集所有行动卡
	allCards := board.collectAllActionCards()
	board.logging.Debug("Action cards have been collected", zap.Int("cardCount", len(allCards)))

	// 揭示所有卡牌
	for _, card := range allCards {
		card.Reveal()
	}
	board.logging.Debug("All cards have been revealed")

	// 按优先级排序
	sort.Slice(allCards, func(i, j int) bool {
		return allCards[i].Priority() < allCards[j].Priority()
	})
	board.logging.Debug("All cards have been sorted by priority")

	// 处理卡牌
	for _, card := range allCards {
		board.logging.Debug("Processing card",
			zap.String("cardID", card.Id()),
			zap.String("cardType", string(card.Type())))

		if err := board.applyCardEffect(card); err != nil {
			board.logging.Error("Failed to process card",
				zap.String("cardID", card.Id()),
				zap.Error(err))
			return fmt.Errorf("failed to process card: %w", err)
		}
	}

	// 处理后续效果
	err := board.handleCardsAfterResolution(allCards)
	if err != nil {
		board.logging.Error("Failed to process subsequent card effects", zap.Error(err))
		return err
	}

	board.logging.Debug("All action cards have been successfully processed")
	return nil
}

// applyCardEffect 应用卡牌效果
func (board *Board) applyCardEffect(card Card) error {
	board.logging.Debug("Applying card effect",
		zap.String("cardID", card.Id()),
		zap.String("cardType", string(card.Type())))

	switch c := card.(type) {
	case *MovementCard:
		return board.handleMovementCard(c)
	case *ForbidMovementCard:
		board.forbiddenActions[card.Target()] = ForbidMovementType
		board.logging.Debug("The forbidden movement effect has been applied",
			zap.Any("target", card.Target()))
		return nil
	case *IntrigueCard:
		return board.handleIntrigueCard(c)
	case *ForbidIntrigueCard:
		board.forbiddenActions[card.Target()] = ForbidIntrigueType
		board.logging.Debug("The forbidden intrigue effect has been applied",
			zap.Any("target", card.Target()))
		return nil
	case *GoodwillCard:
		return board.handleGoodwillCard(c)
	case *ForbidGoodwillCard:
		board.forbiddenActions[card.Target()] = ForbidGoodwillType
		board.logging.Debug("The forbidden goodwill effect has been applied",
			zap.Any("target", card.Target()))
		return nil
	case *ParanoiaCard:
		return board.handleParanoiaCard(c)
	case *ForbidParanoiaCard:
		board.forbiddenActions[card.Target()] = ForbidParanoiaType
	default:
		err := fmt.Errorf("unknown card type")
		board.logging.Error("Unknown card type", zap.Error(err))
		return err
	}
	return nil
}

// handleMovementCard 处理移动卡牌
func (board *Board) handleMovementCard(movementCard *MovementCard) error {
	board.logging.Debug("Processing movement card",
		zap.Any("target", movementCard.Target()),
		zap.String("direction", string(movementCard.Direction)))

	if _, ok := board.forbiddenActions[movementCard.Target()]; ok {
		board.logging.Debug("Target is prohibited from moving",
			zap.Any("target", movementCard.Target()))
		return nil
	}

	target := movementCard.Target()
	if target == nil {
		err := fmt.Errorf("the movement card's target is nil")
		board.logging.Error("Invalid movement target", zap.Error(err))
		return err
	}

	target.ToLocation(board, movementCard.Direction)
	board.logging.Debug("Target has been successfully moved",
		zap.Any("target", target),
		zap.String("direction", string(movementCard.Direction)))
	return nil
}

// handleIntrigueCard 处理阴谋卡牌
func (board *Board) handleIntrigueCard(intrigueCard *IntrigueCard) error {
	board.logging.Debug("Processing intrigue card",
		zap.Any("target", intrigueCard.Target()),
		zap.Int("value", intrigueCard.Value))

	if _, ok := board.forbiddenActions[intrigueCard.Target()]; ok {
		board.logging.Debug("Target is prohibited from intrigue action",
			zap.Any("target", intrigueCard.Target()))
		return nil
	}

	target := intrigueCard.Target()
	if target == nil {
		err := fmt.Errorf("intrigue card target is nil")
		board.logging.Error("Invalid intrigue target", zap.Error(err))
		return err
	}

	oldValue := target.Intrigue()
	target.SetIntrigue(oldValue + intrigueCard.Value)
	board.logging.Debug("The intrigue effect has been applied",
		zap.Any("target", target),
		zap.Int("oldValue", oldValue),
		zap.Int("newValue", target.Intrigue()))
	return nil
}

// handleGoodwillCard 处理好感卡牌
func (board *Board) handleGoodwillCard(goodwillCard *GoodwillCard) error {
	board.logging.Debug("Processing goodwill card",
		zap.Any("target", goodwillCard.Target()),
		zap.Int("value", goodwillCard.Value))

	if _, ok := board.forbiddenActions[goodwillCard.Target()]; ok {
		board.logging.Debug("Target is prohibited from goodwill action",
			zap.Any("target", goodwillCard.Target()))
		return nil
	}

	target := goodwillCard.Target()
	if target == nil {
		err := fmt.Errorf("invalid goodwill target")
		board.logging.Error("Invalid goodwill target", zap.Error(err))
		return err
	}

	oldValue := target.Goodwill()
	target.SetGoodwill(target.Goodwill() + goodwillCard.Value)
	board.logging.Debug("The goodwill effect has been applied",
		zap.Any("target", target),
		zap.Int("oldValue", oldValue),
		zap.Int("newValue", target.Goodwill()))
	return nil
}
func (board *Board) handleParanoiaCard(c *ParanoiaCard) error {
	board.logging.Debug("Processing paranoia card",
		zap.Any("target", c.Target()),
		zap.Int("value", c.Value))
	if _, ok := board.forbiddenActions[c.Target()]; ok {
		board.logging.Debug("Target is prohibited from paranoia action",
			zap.Any("target", c.Target()))
		return nil
	}
	target := c.Target()
	if target == nil {
		err := fmt.Errorf("invalid paranoia target")
		board.logging.Error("Invalid paranoia target", zap.Error(err))
		return err
	}
	oldValue := target.Paranoia()
	target.SetParanoia(target.Paranoia() + c.Value)
	board.logging.Debug("The paranoia effect has been applied",
		zap.Any("target", target),
		zap.Int("oldValue", oldValue),
		zap.Int("newValue", target.Paranoia()))
	return nil
}

// collectAllActionCards 收集所有已打出的行动卡
func (board *Board) collectAllActionCards() []Card {
	board.logging.Debug("Collecting all action cards",
		zap.Int("cardCount", len(board.actionCards)))
	return board.actionCards
}

// handleCardsAfterResolution 处理卡牌结算后的效果
func (board *Board) handleCardsAfterResolution(cards []Card) error {
	board.logging.Debug("Processing card effects after resolution",
		zap.Int("cardCount", len(cards)))

	for _, card := range cards {
		board.logging.Debug("Returning the card to the hand",
			zap.String("cardID", card.Id()),
			zap.String("cardType", string(card.Type())))

		err := ReturnToHand(card)
		if err != nil {
			board.logging.Error("Failed to return the card to the hand",
				zap.String("cardID", card.Id()),
				zap.Error(err))
			return err
		}
	}

	board.logging.Debug("All card effects after resolution have been successfully processed")
	return nil
}

// SetCard 在指定位置添加卡牌
func (board *Board) SetCard(target TargetType, card Card) error {
	board.logging.Debug("Adding card to target",
		zap.String("cardID", card.Id()),
		zap.Any("target", target))

	if !card.IsValidTarget(target) {
		err := fmt.Errorf("invalid card target")
		board.logging.Error("Invalid card target",
			zap.String("cardID", card.Id()),
			zap.Error(err))
		return err
	}

	err := card.SetTarget(target)
	if err != nil {
		board.logging.Error("Failed to set card target",
			zap.String("cardID", card.Id()),
			zap.Error(err))
		return err
	}

	board.actionCards = append(board.actionCards, card)

	board.logging.Debug("Card has been successfully added to target")
	return nil
}

// GetLocation 获取指定类型的位置
func (board *Board) GetLocation(locationType LocationType) *Location {
	location := board.locations[locationType]
	board.logging.Debug("Location has been retrieved",
		zap.String("locationType", string(locationType)),
		zap.Bool("found", location != nil))
	return location
}

// MoveTo 移动角色到指定位置
func (board *Board) MoveTo(character *Character, location LocationType) error {
	board.logging.Debug("Attempting to move character",
		zap.String("character", string(character.Name)),
		zap.String("targetLocation", string(location)))

	if !character.CanMoveTo(location) {
		err := fmt.Errorf("character cannot move to this location")
		board.logging.Error("Character cannot move to this location",
			zap.String("character", string(character.Name)),
			zap.String("targetLocation", string(location)),
			zap.Error(err))
		return err
	}

	delete(board.locations, character.CurrentLocation)
	board.locations[location].Characters[character.Name] = character

	err := character.MoveTo(location)
	if err != nil {
		board.logging.Error("Failed to move character",
			zap.String("character", string(character.Name)),
			zap.String("targetLocation", string(location)),
			zap.Error(err))
		return err
	}

	board.logging.Debug("Character has been successfully moved",
		zap.String("character", string(character.Name)),
		zap.String("from", string(character.CurrentLocation)),
		zap.String("to", string(location)))
	return nil
}

// ResetCounters 重置所有计数器
func (board *Board) ResetCounters() error {
	board.logging.Debug("Starting to Reset all counters")

	// 检查是否初始化
	if board.locations == nil {
		err := fmt.Errorf("the board locations are not initialized")
		board.logging.Error("Failed to Reset counters: The board is not initialized", zap.Error(err))
		return err
	}

	// 重置所有位置的阴谋值
	for locType, location := range board.locations {
		oldValue := location.Intrigue()
		location.SetIntrigue(0)
		board.logging.Debug("Location intrigue counter has been Reset",
			zap.String("location", string(locType)),
			zap.Int("oldValue", oldValue),
			zap.Int("newValue", 0))
	}

	// 重置所有角色的计数器
	for _, character := range board.characters {
		board.logging.Debug("Resetting character counters",
			zap.String("character", string(character.Name)))

		// 重置不安值
		oldParanoia := character.Paranoia()
		character.SetParanoia(0)
		board.logging.Debug("Character paranoia has been Reset",
			zap.String("character", string(character.Name)),
			zap.Int("oldValue", oldParanoia),
			zap.Int("newValue", 0))

		// 重置阴谋值
		oldIntrigue := character.Intrigue()
		character.SetIntrigue(0)
		board.logging.Debug("Character intrigue has been Reset",
			zap.String("character", string(character.Name)),
			zap.Int("oldValue", oldIntrigue),
			zap.Int("newValue", 0))

		// 重置好感度
		oldGoodwill := character.Goodwill()
		character.SetGoodwill(0)
		board.logging.Debug("Character goodwill has been Reset",
			zap.String("character", string(character.Name)),
			zap.Int("oldValue", oldGoodwill),
			zap.Int("newValue", 0))
	}

	board.logging.Debug("All counters have been successfully Reset")
	return nil
}

// ReturnAllCards 返回所有卡牌
func (board *Board) ReturnAllCards(state *GameState) error {
	board.logging.Debug("Starting to return all cards")

	// 检查是否有actionCards
	if board.actionCards == nil {
		board.logging.Debug("No action cards to return")
		return nil
	}

	// 所有当前回合的卡牌返回手牌
	for _, card := range board.actionCards {
		board.logging.Debug("Returning action card to hand",
			zap.String("cardID", card.Id()),
			zap.String("cardType", string(card.Type())))

		err := ReturnToHand(card)
		if err != nil {
			board.logging.Error("Failed to return the card to the hand",
				zap.String("cardID", card.Id()),
				zap.Error(err))
			return err
		}
	}

	// 清空board上的卡牌
	cardCount := len(board.actionCards)
	board.actionCards = nil
	board.logging.Debug("The action cards on the board have been cleared",
		zap.Int("clearedCardCount", cardCount))

	// 清空禁止动作
	forbiddenCount := len(board.forbiddenActions)
	board.forbiddenActions = make(map[any]CardType)
	board.logging.Debug("Forbidden actions have been Reset",
		zap.Int("clearedActionCount", forbiddenCount))

	// 处理玩家已使用的卡牌
	for _, protagonist := range state.Protagonists {
		usedCards := protagonist.OnceCards
		if len(usedCards) > 0 {
			board.logging.Debug("Processing protagonist's used cards",
				zap.String("protagonistID", protagonist.ID),
				zap.Int("cardCount", len(usedCards)))

			for _, card := range usedCards {
				board.logging.Debug("Returning protagonist's used card",
					zap.String("protagonistID", protagonist.ID),
					zap.String("cardID", card.Id()))

				err := ReturnToHand(card)
				if err != nil {
					board.logging.Error("Failed to return protagonist's card",
						zap.String("protagonistID", protagonist.ID),
						zap.String("cardID", card.Id()),
						zap.Error(err))
					return err
				}
			}

			// 清空已使用卡牌列表
			protagonist.OnceCards = nil
			board.logging.Debug("The protagonist's used cards have been cleared",
				zap.String("protagonistID", protagonist.ID))
		}
	}

	board.logging.Debug("All cards have been successfully returned")
	return nil
}

func (board *Board) Locations() []LocationType {
	return []LocationType{
		"LocationHospital",
		"LocationCity",
		"LocationSchool",
		"LocationShrine",
	}
}
