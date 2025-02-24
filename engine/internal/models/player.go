package models

import "fmt"

type Player interface {
	PlaceCards(card Card) error
	RecycleCards(card Card) error
}

// PlayerBase 表示游戏中的玩家基础属性
type PlayerBase struct {
	ID             string // 玩家唯一标识
	Name           string // 玩家名称
	HandCards      []Card // 玩家当前持有的卡牌
	OnceCards      []Card // 已使用的一次性卡牌（每循环使用一次的卡牌）
	MaxCardsPerDay int    // 每天可以使用的最大卡牌数量（固定为3张）
}

func (p *PlayerBase) PlaceCards(card Card) error {
	for i := 0; i < len(p.HandCards); i++ {
		if p.HandCards[i] == card {
			p.HandCards = append(p.HandCards[:i], p.HandCards[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("玩家没有该卡牌")
}

func (p *PlayerBase) RecycleCards(card Card) error {
	if card.IsOncePerLoop() {
		p.OnceCards = append(p.OnceCards, card)
	} else {
		p.HandCards = append(p.HandCards, card)
	}
	return nil
}

// Mastermind 表示幕后主使玩家
type Mastermind struct {
	PlayerBase // 继承基础玩家属性
}

// NewMastermind 创建新的幕后主使玩家
func NewMastermind() *Mastermind {
	mastermind := &Mastermind{
		PlayerBase: PlayerBase{
			ID:             "Mastermind",
			Name:           "Mastermind",
			HandCards:      nil,
			OnceCards:      nil,
			MaxCardsPerDay: 3,
		},
	}
	mastermind.HandCards = InitMastermindCard(mastermind)
	return mastermind
}

// PlaceActionCards 允许幕后主使放置行动卡
func (mastermind *Mastermind) PlaceActionCards(state *GameState) error {
	for i := 0; i < 3; i++ {
		mastermindHandCards := mastermind.HandCards[i]
		switch i {
		case 0:
			character := state.Script.Characters[0]
			err := state.Board.SetCard(character, mastermindHandCards)
			if err != nil {
				return err
			}
			err = mastermind.PlaceCards(mastermindHandCards)
			if err != nil {
				return err
			}
		case 1:
			locationShrine := state.Board.GetLocation(LocationShrine)
			err := state.Board.SetCard(locationShrine, mastermindHandCards)
			if err != nil {
				return err
			}
			err = mastermind.PlaceCards(mastermindHandCards)
			if err != nil {
				return err
			}
		case 2:
			locationHospital := state.Board.GetLocation(LocationHospital)
			err := state.Board.SetCard(locationHospital, mastermindHandCards)
			if err != nil {
				return err
			}
			err = mastermind.PlaceCards(mastermindHandCards)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

type Protagonists []*Protagonist

func (protagonists Protagonists) MakeFinalGuess(script *Script) (bool, error) {
	return false, nil
}

func (protagonists Protagonists) GetLeader() *Protagonist {
	return nil
}

func (protagonists Protagonists) PlaceActionCards(state *GameState) error {
	for _, protagonist := range protagonists {
		err := protagonist.PlaceActionCards(state)
		if err != nil {
			return err
		}
	}
	return nil
}

// Protagonist 表示主角玩家
type Protagonist struct {
	PlayerBase      // 继承基础玩家属性
	IsLeader   bool // 是否为当前领袖
}

// NewProtagonist 创建新的主角玩家
func NewProtagonist(id string, isLeader bool) *Protagonist {
	protagonist := &Protagonist{
		PlayerBase: PlayerBase{
			ID:             id,
			Name:           fmt.Sprintf("Protagonist-%s", id),
			HandCards:      nil,
			OnceCards:      nil,
			MaxCardsPerDay: 1,
		},
		IsLeader: isLeader,
	}
	protagonist.HandCards = InitProtagonistCard(protagonist)
	return protagonist
}

// PlaceActionCards 允许主角玩家根据其控制的牌组数量打出牌
func (protagonist *Protagonist) PlaceActionCards(state *GameState) error {
	// 根据是否为领袖决定可以打出的牌数
	//cardsToPlay := 1
	//if protagonist.IsLeader {
	//	cardsToPlay = 2  // 领袖可以从两副牌组中各打出一张
	//}
	// 待实现：选择和打出卡牌的具体逻辑

	switch protagonist.ID {
	case "A":
		character := state.Script.Characters[0]
		card := protagonist.HandCards[0]
		err := state.Board.SetCard(character, card)
		if err != nil {
			return err
		}
		err = protagonist.PlaceCards(card)
		if err != nil {
			return err
		}
	case "B":
		character := state.Script.Characters[1]
		card := protagonist.HandCards[1]
		err := state.Board.SetCard(character, card)
		if err != nil {
			return err
		}
		err = protagonist.PlaceCards(card)
		if err != nil {
			return err
		}
	case "C":
		character := state.Script.Characters[2]
		card := protagonist.HandCards[2]
		err := state.Board.SetCard(character, card)
		if err != nil {
			return err
		}
		err = protagonist.PlaceCards(card)
		if err != nil {
			return err
		}
	}

	return nil
}

// SetLeader 更新主角玩家的领袖状态
func (protagonist *Protagonist) SetLeader(isLeader bool) {
	protagonist.IsLeader = isLeader
}

// PassDeck 在领袖更替时将一副牌组传递给另一位主角玩家
func (protagonist *Protagonist) PassDeck(receiver *Protagonist) error {
	if !protagonist.IsLeader {
		return fmt.Errorf("只有领袖可以传递牌组")
	}
	// 待实现：传递牌组的具体逻辑
	return nil
}
