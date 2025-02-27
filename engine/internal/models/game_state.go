package models

import (
	"fmt"
	"go.uber.org/zap"
	"time"
)

// state 管理整个游戏的状态

type GameState struct {
	logging       *zap.Logger
	Script        *Script   // 当前剧本
	CurrentPhase  DayPhase  // 当前阶段
	CurrentLoop   int       // 当前循环
	CurrentDay    int       // 当前日期
	CurrentPlayer Player    // 当前玩家
	TimeSpiral    time.Time // 时间螺旋阶段（TODO：具体实现时间螺旋的逻辑）

	IsGameOver   bool         // 游戏是否结束
	WinnerType   string       // 获胜方类型
	Board        *Board       // 游戏板状态
	Protagonists Protagonists // 主人公
	Mastermind   *Mastermind  // 幕后主使

	GuessMade bool // 是否进行了最终猜测

	Roles []*Role

	IncidentsOccurred map[string]bool
	TimingAbility     map[RoleAbilityTiming][]RoleAbility // 当前阶段可用的角色能力
	Characters        []*Character                        // 游戏中的所有角色
	RoleTypes         map[RoleType]*Character             // 角色身份对应的角色
	ActiveRoles       map[string]*RoleAbility             // 当前激活的角色能力
	Incidents         []Incident
}

func NewGameState(logging *zap.Logger) *GameState {
	return &GameState{
		logging:      logging,
		Script:       nil,
		CurrentPhase: PhaseDayStart,
		CurrentLoop:  0,
		CurrentDay:   0,
		TimeSpiral:   time.Time{},
		IsGameOver:   false,
		WinnerType:   "",
		Board:        nil,
		Protagonists: nil,
		Mastermind:   nil,
		GuessMade:    false,
		Characters:   nil,
		Incidents:    nil,
		Roles:        nil,

		IncidentsOccurred: make(map[string]bool),
		TimingAbility:     make(map[RoleAbilityTiming][]RoleAbility),
		RoleTypes:         make(map[RoleType]*Character),
		ActiveRoles:       make(map[string]*RoleAbility),
	}
}
func (gs *GameState) Character(characterName CharacterName) *Character {
	for _, c := range gs.Characters {
		if c.Name == characterName {
			return c
		}
	}
	return nil
}

func (gs *GameState) Location(locationType LocationType) *Location {
	return gs.Board.GetLocation(locationType)
}

// PrintGameState 详细打印游戏状态信息
func (gs *GameState) PrintGameState() {
	if gs.logging == nil {
		return
	}

	gs.logging.Debug("=================== Game State Overview ===================")

	// 1. 基本游戏信息
	gs.logging.Debug("Basic Game Information",
		zap.Int("Current Loop", gs.CurrentLoop),
		zap.Int("Days Per Loop", gs.Script.DaysPerLoop),
		zap.Int("Current Day", gs.CurrentDay),
		zap.String("Current Phase", string(gs.CurrentPhase)),
		zap.Bool("Is Game Over", gs.IsGameOver),
		zap.String("Winner", gs.WinnerType),
		zap.Bool("Final Guess Made", gs.GuessMade))

	// 2. 剧本信息
	if gs.Script != nil {
		mainPlotName := ""
		if gs.Script.MainPlot != nil {
			mainPlotName = gs.Script.MainPlot.Name
		}
		subPlotNames := make([]string, 0)
		for _, subplot := range gs.Script.SubPlots {
			subPlotNames = append(subPlotNames, subplot.Name)
		}

		gs.logging.Debug("Script Information",
			zap.String("Main Plot", mainPlotName),
			zap.Strings("Sub Plots", subPlotNames),
			zap.Int("Maximum Loops", gs.Script.MaxLoops),
			zap.Int("Days Per Loop", gs.Script.DaysPerLoop))
	}

	// 3. 详细的角色信息
	gs.logging.Debug("---------- Character Status ----------")
	for _, char := range gs.Characters {
		if char == nil {
			continue
		}

		roleType := "Person"
		if char.Role() != nil {
			roleType = string(char.Role().Type)
		}

		// 角色基本信息
		gs.logging.Debug("Character Details",
			zap.String("Name", string(char.Name)),
			zap.String("Role", roleType),
			zap.String("Current Location", string(char.Location())),
			zap.String("Starting Location", string(char.StartLocation)),
			zap.Bool("Is Alive", char.IsAlive()))

		// 角色计数器信息
		gs.logging.Debug("Character Counters",
			zap.String("Character", string(char.Name)),
			zap.Int("Goodwill", char.Goodwill()),
			zap.Int("Goodwill Limit", char.GoodwillLimit),
			zap.Int("Paranoia", char.Paranoia()),
			zap.Int("Paranoia Limit", char.ParanoiaLimit),
			zap.Int("Intrigue", char.Intrigue()))

		// 角色能力信息
		abilities := make([]string, 0)
		for _, ability := range char.GoodwillAbilityList {
			abilities = append(abilities, fmt.Sprintf("%s (Cost: %d, Can Be Refused: %v)",
				ability.Name, ability.Cost, ability.CanBeRefused))
		}
		gs.logging.Debug("Character Abilities",
			zap.String("Character", string(char.Name)),
			zap.Strings("Goodwill Abilities", abilities))
	}

	// 4. 位置信息
	gs.logging.Debug("---------- Location Status ----------")
	if gs.Board != nil {
		locations := []LocationType{LocationHospital, LocationCity, LocationSchool, LocationShrine}
		for _, locType := range locations {
			loc := gs.Board.GetLocation(locType)
			if loc == nil {
				continue
			}

			// 获取该位置的所有角色名称
			characterNames := make([]string, 0)
			for charName := range loc.Characters {
				characterNames = append(characterNames, string(charName))
			}

			gs.logging.Debug("Location Details",
				zap.String("Location", string(locType)),
				zap.Int("Intrigue Count", loc.Intrigue()),
				zap.Int("Character Count", len(loc.Characters)),
				zap.Strings("Characters Present", characterNames))

		}
	}

	// 5. 主角方信息
	gs.logging.Debug("---------- Protagonists Status ----------")
	for _, protagonist := range gs.Protagonists {
		gs.logging.Debug("Protagonist Details",
			zap.String("ID", protagonist.ID),
			zap.String("Name", protagonist.Name),
			zap.Bool("Is Leader", protagonist.IsLeader),
			zap.Int("Cards In Hand", len(protagonist.HandCards)),
			zap.Int("Once Cards Used", len(protagonist.OnceCards)))
	}

	// 6. 幕后主使信息
	if gs.Mastermind != nil {
		gs.logging.Debug("---------- Mastermind Status ----------",
			zap.Int("Cards In Hand", len(gs.Mastermind.HandCards)),
			zap.Int("Once Cards Used", len(gs.Mastermind.OnceCards)))
	}

	// 7. Action Cards Status
	gs.logging.Debug("---------- Action Cards Status ----------")

	// 当前回合的行动卡
	if gs.Board != nil && gs.Board.actionCards != nil {
		gs.logging.Debug("Current Action Cards on Board")

		// 按优先级分类打印卡牌
		// 1. Forbid Movement cards
		// 2. Movement cards
		// 3. Other Forbid cards
		// 4. Other remaining action cards

		// 创建分类映射
		cardsByType := make(map[CardType][]Card)
		for _, card := range gs.Board.actionCards {
			cardsByType[card.Type()] = append(cardsByType[card.Type()], card)
		}

		// 按照规则顺序打印
		priorityOrder := []CardType{
			ForbidMovementType, // 1st: Forbid Movement
			MovementType,       // 2nd: Movement
			ForbidIntrigueType, // 3rd: Other Forbid cards
			ForbidParanoiaType,
			ForbidGoodwillType,
			IntrigueType, // 4th: Other remaining cards
			ParanoiaType,
			GoodwillType,
		}

		for _, cardType := range priorityOrder {
			if cards, exists := cardsByType[cardType]; exists {
				for _, card := range cards {
					targetInfo := "Unknown"
					if target := card.Target(); target != nil {
						switch t := target.(type) {
						case *Character:
							targetInfo = fmt.Sprintf("Character: %s", t.Name)
						case *Location:
							targetInfo = fmt.Sprintf("Location: %s", t.LocationType)
						}
					}

					gs.logging.Debug("Action Card",
						zap.String("Type", string(card.Type())),
						zap.String("ID", card.Id()),
						zap.String("Target", targetInfo),
						zap.Int("Priority", card.Priority()),
						zap.Bool("OncePerLoop", card.IsOncePerLoop()),
						zap.Reflect("Owner", card.Owner()))
				}
			}
		}
	}

	// 打印主角方的手牌和已使用的一次性卡牌
	gs.logging.Debug("---------- Protagonists Cards ----------")
	for _, protagonist := range gs.Protagonists {
		// 打印手牌
		gs.logging.Debug(fmt.Sprintf("Protagonist %s Hand Cards", protagonist.Name))
		for _, card := range protagonist.HandCards {
			gs.logging.Debug("Card",
				zap.String("Type", string(card.Type())),
				zap.String("ID", card.Id()),
				zap.Bool("OncePerLoop", card.IsOncePerLoop()))
		}

		// 打印已使用的一次性卡牌
		gs.logging.Debug(fmt.Sprintf("Protagonist %s Used Once-Per-Loop Cards", protagonist.Name))
		for _, card := range protagonist.OnceCards {
			gs.logging.Debug("Used Once Card",
				zap.String("Type", string(card.Type())),
				zap.String("ID", card.Id()))
		}
	}

	// 打印幕后主使的手牌和已使用的一次性卡牌
	gs.logging.Debug("---------- Mastermind Cards ----------")
	if gs.Mastermind != nil {
		// 打印手牌
		gs.logging.Debug("Mastermind Hand Cards")
		for _, card := range gs.Mastermind.HandCards {
			gs.logging.Debug("Card",
				zap.String("Type", string(card.Type())),
				zap.String("ID", card.Id()),
				zap.Bool("OncePerLoop", card.IsOncePerLoop()))
		}

		// 打印已使用的一次性卡牌
		gs.logging.Debug("Mastermind Used Once-Per-Loop Cards")
		for _, card := range gs.Mastermind.OnceCards {
			gs.logging.Debug("Used Once Card",
				zap.String("Type", string(card.Type())),
				zap.String("ID", card.Id()))
		}
	}
	gs.logging.Debug("===================== End of State =====================")
}
