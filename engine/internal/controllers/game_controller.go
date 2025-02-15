package controllers

import (
	"errors"
	"fmt"
	"go.uber.org/zap"
	"tragedy-looper/engine/internal/models"
)

type GameController struct {
	logger *zap.Logger
	state  *models.GameState
	script *models.Script
}

func NewGameController(logger *zap.Logger, script *models.Script) *GameController {
	return &GameController{
		state:  models.NewGameState(logger),
		script: script,
		logger: logger,
	}
}

func (gc *GameController) StartGame() error {
	gc.logger.Info("Game initialization started in controller")

	// 1. 先设置脚本
	err := gc.setupGame()
	if err != nil {
		gc.logger.Error("Game setup failed", zap.Error(err))
		return err
	}

	gc.logger.Info("Enter the main game loop from the controller")

	// 2. 进入游戏主循环
	err = gc.gameLoop()
	if err != nil {
		gc.logger.Error("Game loop failed", zap.Error(err))
		return err
	}
	return nil
}

// setupGame 设置游戏，选择剧本并设定角色与事件
func (gc *GameController) setupGame() error {
	if gc.script == nil {
		gc.logger.Error("Setup failed: script is nil")
		return errors.New("script is nil")
	}
	if gc.state.Script != nil {
		gc.logger.Error("Setup failed: script already set in state")
		return errors.New("script is not nil")
	}

	// 设置脚本到 state
	gc.state.Script = gc.script

	gc.logger.Info("Setup Mastermind")
	gc.state.Mastermind = models.NewMastermind()

	gc.logger.Info("Setup Protagonists")
	gc.state.Protagonists = models.Protagonists{
		models.NewProtagonist("A", true),
		models.NewProtagonist("B", false),
		models.NewProtagonist("C", false),
	}

	gc.logger.Debug("Setup script",
		zap.String("MainPlot", gc.script.MainPlot.Name),
		zap.Int("CharactersCount", len(gc.script.Characters)),
	)

	gc.state.Characters = gc.script.Characters
	gc.state.Incidents = gc.script.Incidents

	for _, character := range gc.script.Characters {
		role := character.Role()
		if role == nil {
			gc.logger.Error("Setup failed: character has no role",
				zap.String("Character", string(character.Name)))
			return errors.New("character role is nil")
		}
		gc.state.Roles = append(gc.state.Roles, role)
	}

	gc.logger.Info("Initialize the game board")
	gc.state.Board = models.NewBoard(gc.logger, gc.state.Characters)

	gc.logger.Info("Game setup complete")
	return nil
}

// gameLoop 游戏主循环
func (gc *GameController) gameLoop() error {
	for {
		// 检查是否达到循环上限
		if gc.state.CurrentLoop >= gc.script.MaxLoops {
			gc.logger.Info("Reached the maximum number of loops, entering final guess phase",
				zap.Int("CurrentLoop", gc.state.CurrentLoop),
				zap.Int("MaxLoops", gc.script.MaxLoops))
			err := gc.enterFinalGuess()
			if err != nil {
				gc.logger.Error("Final guess phase failed", zap.Error(err))
				return err
			}
			break
		}

		// 准备新的循环
		err := gc.prepareLoop()
		if err != nil {
			gc.logger.Error("Prepare loop failed",
				zap.Int("CurrentLoop", gc.state.CurrentLoop),
				zap.Error(err))
			return err
		}

		// 时间螺旋阶段
		gc.logger.Info("Enter time spiral phase",
			zap.Int("CurrentLoop", gc.state.CurrentLoop))
		err = gc.timeSpiralPhase()
		if err != nil {
			gc.logger.Error("Time spiral phase failed",
				zap.Int("CurrentLoop", gc.state.CurrentLoop),
				zap.Error(err))
			return err
		}

		// 角色归位
		gc.logger.Info("Reset character positions",
			zap.Int("CurrentLoop", gc.state.CurrentLoop))
		err = gc.state.Board.Reset() // 原来的 reset() 在 Board 上
		if err != nil {
			gc.logger.Error("Reset character positions failed",
				zap.Int("CurrentLoop", gc.state.CurrentLoop),
				zap.Error(err))
			return err
		}

		// 重置计数器
		gc.logger.Info("Reset counters", zap.Int("CurrentLoop", gc.state.CurrentLoop))
		err = gc.state.Board.ResetCounters()
		if err != nil {
			gc.logger.Error("Reset counters failed",
				zap.Int("CurrentLoop", gc.state.CurrentLoop),
				zap.Error(err))
			return err
		}

		// 返回所有卡牌
		gc.logger.Info("Return all cards", zap.Int("CurrentLoop", gc.state.CurrentLoop))
		err = gc.state.Board.ReturnAllCards(gc.state)
		if err != nil {
			gc.logger.Error("Return cards failed",
				zap.Int("CurrentLoop", gc.state.CurrentLoop),
				zap.Error(err))
			return err
		}

		// 每日流程
		gc.logger.Info("Start daily phase", zap.Int("CurrentLoop", gc.state.CurrentLoop))
		err = gc.dailyPhases()
		if err != nil {
			gc.logger.Error("Daily phase failed",
				zap.Int("CurrentLoop", gc.state.CurrentLoop),
				zap.Error(err))
			return err
		}

		// 检查胜利条件
		if gc.checkWinCondition() {
			gc.logger.Info("Protagonists have met the win condition",
				zap.Int("CurrentLoop", gc.state.CurrentLoop))
			gc.state.IsGameOver = true
			gc.state.WinnerType = "Protagonists"
			break
		} else if gc.state.CurrentLoop >= gc.script.MaxLoops {
			gc.logger.Info("Reached the maximum number of loops, entering final guess",
				zap.Int("CurrentLoop", gc.state.CurrentLoop))
			err = gc.enterFinalGuess()
			if err != nil {
				gc.logger.Error("Final guess phase failed", zap.Error(err))
				return err
			}
			break
		} else {
			gc.logger.Info("Enter the next loop",
				zap.Int("CurrentLoop", gc.state.CurrentLoop),
				zap.Int("NextLoop", gc.state.CurrentLoop+1))
			gc.state.CurrentLoop++
		}
	}

	gc.logger.Info("Game loop ended",
		zap.String("Winner", gc.state.WinnerType),
		zap.Bool("GameOver", gc.state.IsGameOver))
	return nil
}

// enterFinalGuess 进入最终猜测阶段
func (gc *GameController) enterFinalGuess() error {
	gc.logger.Info("Enter final guess phase")
	if gc.state.GuessMade {
		gc.logger.Error("Final guess has already been made")
		return errors.New("最终猜测已经进行过")
	}

	gc.logger.Debug("Protagonists are making the final guess")
	correctGuess, err := gc.state.Protagonists.MakeFinalGuess(gc.script)
	if err != nil {
		gc.logger.Error("An error occurred during the final guess", zap.Error(err))
		return err
	}

	gc.state.GuessMade = true
	gc.state.IsGameOver = true

	if correctGuess {
		gc.logger.Info("Protagonists win because the guess is correct")
		gc.state.WinnerType = "Protagonists"
	} else {
		gc.logger.Info("Mastermind wins because the final guess is wrong")
		gc.state.WinnerType = "Mastermind"
	}

	return nil
}

// prepareLoop 准备新的循环
func (gc *GameController) prepareLoop() error {
	gc.logger.Info("=================== Preparing New Loop ===================",
		zap.Int("Loop", gc.state.CurrentLoop+1))

	// 来源: 知识库中的 "Preparing the Loop" 部分
	gc.logger.Info("1. Time Spiral Phase - Protagonists discussion time")
	// TODO: 实现时间螺旋阶段的具体逻辑

	gc.logger.Info("2. Returning characters to starting positions")
	err := gc.state.Board.Reset()
	if err != nil {
		return err
	}

	gc.logger.Info("3. Removing and replacing counters")
	err = gc.state.Board.ResetCounters()
	if err != nil {
		return err
	}

	gc.logger.Info("4. Returning all action cards to hands")
	err = gc.state.Board.ReturnAllCards(gc.state)
	if err != nil {
		return err
	}

	// 重置游戏状态
	gc.state.IsGameOver = false
	gc.state.WinnerType = ""
	gc.state.CurrentLoop++
	gc.state.CurrentDay = 1

	// 重置角色状态
	for _, character := range gc.state.Characters {
		character.ResetState()
	}

	gc.logger.Info("Loop preparation completed",
		zap.Int("NewLoop", gc.state.CurrentLoop))

	// 打印初始状态
	gc.state.PrintGameState()

	return nil
}

// timeSpiralPhase 时间螺旋阶段
func (gc *GameController) timeSpiralPhase() error {
	// TODO: 实现时间螺旋的逻辑, 例如主角讨论策略
	return nil
}

// dailyPhases 每日流程开始
func (gc *GameController) dailyPhases() error {
	gc.logger.Info("Start daily phase",
		zap.Int("CurrentLoop", gc.state.CurrentLoop),
		zap.Int("CurrentDay", gc.state.CurrentDay),
		zap.Int("MaxDays", gc.script.DaysPerLoop))

	for {
		// 是否为最后一天
		if gc.state.CurrentDay > gc.script.DaysPerLoop {
			gc.logger.Info("Reached the last day of the loop",
				zap.Int("CurrentLoop", gc.state.CurrentLoop),
				zap.Int("CompletedDays", gc.state.CurrentDay-1))
			break
		}
		err := gc.processDay()
		if err != nil {
			gc.logger.Error("Processing today's routine failed",
				zap.Int("Day", gc.state.CurrentDay),
				zap.Int("CurrentLoop", gc.state.CurrentLoop),
				zap.Error(err))
			return err
		}
		if gc.state.IsGameOver {
			gc.logger.Info("Game ended during the daily phase",
				zap.String("Winner", gc.state.WinnerType),
				zap.Int("EndDay", gc.state.CurrentDay),
				zap.Int("EndLoop", gc.state.CurrentLoop))
			return nil
		}

		gc.state.CurrentDay++
		gc.logger.Debug("Enter the next day",
			zap.Int("NextDay", gc.state.CurrentDay),
			zap.Int("CurrentLoop", gc.state.CurrentLoop))
	}

	gc.logger.Info("Daily phase of the current loop completed",
		zap.Int("CurrentLoop", gc.state.CurrentLoop),
		zap.Int("TotalProcessedDays", gc.state.CurrentDay-1))
	return nil
}

// processDay 执行一天的流程
func (gc *GameController) processDay() error {
	gc.logger.Info("=================== New Day Started ===================",
		zap.Int("Day", gc.state.CurrentDay),
		zap.Int("CurrentLoop", gc.state.CurrentLoop))

	// 顺序处理每个阶段
	phases := []models.DayPhase{
		models.PhaseDayStart,
		models.PhaseMastermindAction,
		models.PhaseProtagonistsAction,
		models.PhaseResolveCards,
		models.PhaseMastermindAbilities,
		models.PhaseLeaderGoodwill,
		models.PhaseIncidents,
		models.PhaseSwitchLeader,
		models.PhaseDayEnd,
	}

	for _, phase := range phases {
		gc.logger.Info("----------- Phase Started -----------",
			zap.String("Phase", string(phase)))

		err := gc.processDayPhase(phase)
		if err != nil {
			gc.logger.Error("Day phase failed",
				zap.String("phase", string(phase)),
				zap.Error(err))
			return err
		}

		// 在每个阶段结束后打印游戏状态
		gc.state.PrintGameState()

		if gc.state.IsGameOver {
			gc.logger.Info("Game ended during day phases",
				zap.String("Winner", gc.state.WinnerType))
			return nil
		}
	}

	gc.logger.Info("=================== Day Completed ===================",
		zap.Int("Day", gc.state.CurrentDay),
		zap.Int("CurrentLoop", gc.state.CurrentLoop))

	return nil
}

// processDayPhase 处理每日各阶段
func (gc *GameController) processDayPhase(phase models.DayPhase) error {
	gc.logger.Info("Starting to process the day phase",
		zap.String("phase", string(phase)),
		zap.Int("day", gc.state.CurrentDay),
		zap.Int("CurrentLoop", gc.state.CurrentLoop))

	var err error
	switch phase {
	case models.PhaseDayStart:
		err = gc.handleDayStart()
	case models.PhaseMastermindAction:
		err = gc.handleMastermindAction()
	case models.PhaseProtagonistsAction:
		err = gc.handleProtagonistsAction()
	case models.PhaseResolveCards:
		err = gc.handleResolveCards()
	case models.PhaseMastermindAbilities:
		err = gc.handleMastermindAbilities()
	case models.PhaseLeaderGoodwill:
		err = gc.handleLeaderGoodwill()
	case models.PhaseIncidents:
		err = gc.handleIncidents()
	case models.PhaseSwitchLeader:
		err = gc.handleSwitchLeader()
	case models.PhaseDayEnd:
		err = gc.handleDayEnd()
	default:
		return fmt.Errorf("unknown game phase: %d", phase)
	}
	if err != nil {
		return err
	}

	// 触发与此阶段相关的角色能力
	err = gc.triggerPhaseAbilities(phase)
	if err != nil {
		gc.logger.Error("Trigger phase abilities failed",
			zap.String("phase", string(phase)),
			zap.Error(err))
		return err
	}
	return nil
}

// 各阶段的处理方法
func (gc *GameController) handleDayStart() error {
	return gc.triggerAbilities(models.RoleTimingDayStart)
}

// handleMastermindAction Mastermind放置行动卡
func (gc *GameController) handleMastermindAction() error {
	gc.logger.Info("Mastermind placing action cards...")
	// 来源: 知识库中提到 "Mastermind plays 3 action cards face down"
	gc.logger.Debug("Mastermind must place exactly 3 cards face down")
	err := gc.state.Mastermind.PlaceActionCards(gc.state)
	if err != nil {
		return err
	}
	return nil
}

// handleProtagonistsAction 主角方放置行动卡
func (gc *GameController) handleProtagonistsAction() error {
	gc.logger.Info("Protagonists placing action cards...")
	// 来源: 知识库中提到 "Protagonists play one card each face down"
	gc.logger.Debug("Each Protagonist must place exactly 1 card face down")
	return gc.state.Protagonists.PlaceActionCards(gc.state)
}

// handleResolveCards 处理卡牌结算
func (gc *GameController) handleResolveCards() error {
	gc.logger.Info("Starting to resolve cards...")
	// 来源: 知识库中提到的卡牌结算顺序
	gc.logger.Debug("Cards will be resolved in order: 1.Forbid Movement, 2.Movement, 3.Other Forbid, 4.Other cards")
	return gc.state.Board.ResolveActionCards(gc.state)
}
func (gc *GameController) handleMastermindAbilities() error {
	return gc.triggerAbilities(models.RoleTimingMastermind)
}

func (gc *GameController) handleLeaderGoodwill() error {
	return gc.triggerAbilities(models.RoleTimingGoodwillUse)
}

func (gc *GameController) handleIncidents() error {
	gc.logger.Info("Start processing incidents phase")

	for _, incident := range gc.state.Script.Incidents {
		gc.logger.Debug("Check incident", zap.String("IncidentType", string(incident.Type())))
		if gc.canTriggerIncident(incident) {
			gc.logger.Info("Trigger incident", zap.String("IncidentType", string(incident.Type())))
			err := gc.executeIncident(incident)
			if err != nil {
				gc.logger.Error("Execute incident failed",
					zap.String("IncidentType", string(incident.Type())),
					zap.Error(err))
				return err
			}
		}
	}

	gc.logger.Info("Incidents phase completed")
	return nil
}

func (gc *GameController) handleSwitchLeader() error {
	gc.logger.Debug("Begin handling switch leader, TODO: implement switching logic")
	return nil
}

func (gc *GameController) handleDayEnd() error {
	return gc.triggerAbilities(models.RoleTimingDayEnd)
}

// 触发与阶段对应的角色能力
func (gc *GameController) triggerPhaseAbilities(phase models.DayPhase) error {
	switch phase {
	case models.PhaseDayStart:
		return gc.triggerAbilities(models.RoleTimingDayStart)
	case models.PhaseMastermindAbilities:
		return gc.triggerAbilities(models.RoleTimingMastermind)
	case models.PhaseDayEnd:
		return gc.triggerAbilities(models.RoleTimingDayEnd)
	default:
		// 其他阶段不触发或已有单独处理
		return nil
	}
}

// 触发某个 Timing 下的所有能力
func (gc *GameController) triggerAbilities(timing models.RoleAbilityTiming) error {
	gc.logger.Info("Ability trigger phase started",
		zap.String("Timing", string(timing)))

	var mustAbilities, mandatoryAbilities, optionalAbilities []models.RoleAbility

	for _, character := range gc.state.Characters {
		if !character.IsAlive() {
			continue
		}
		role := character.Role()
		if role == nil {
			continue
		}
		for _, ability := range role.Abilities {
			if ability.GetTiming() != timing {
				continue
			}
			isTriggerable, err := ability.IsTriggerable(gc.state, character)
			if err != nil {
				return err
			}
			if !isTriggerable {
				continue
			}
			switch ability.GetMandatory() {
			case models.GoodwillRefusalMust:
				mustAbilities = append(mustAbilities, ability)
			case models.GoodwillRefusalOptional:
				optionalAbilities = append(optionalAbilities, ability)
			case models.GoodwillRefusalMandatory:
				mandatoryAbilities = append(mandatoryAbilities, ability)
			}
		}
	}

	// 执行 "must" abilities
	for _, ability := range mustAbilities {
		err := ability.Execute(gc.state, nil)
		if err != nil {
			return err
		}
	}

	// 执行 "mandatory" abilities
	for _, ability := range mandatoryAbilities {
		err := ability.Execute(gc.state, nil)
		if err != nil {
			return err
		}
	}

	// 执行 "optional" abilities
	for _, ability := range optionalAbilities {
		err := ability.Execute(gc.state, nil)
		if err != nil {
			return err
		}
	}

	return nil
}

// canTriggerIncident 判断事件是否可以触发
func (gc *GameController) canTriggerIncident(incident models.Incident) bool {
	// 可以根据实际需求改成更精细的判断
	return true
}

// executeIncident 执行事件
func (gc *GameController) executeIncident(incident models.Incident) error {
	// 调用事件的 Execute 方法
	return incident.Execute(*gc.logger, gc.state, nil)
}

// checkWinCondition 检查胜利条件
func (gc *GameController) checkWinCondition() bool {
	// 根据游戏规则判断主角方是否获胜
	return false
}
