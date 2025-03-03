package models

// GamePhase 表示整个游戏的主要阶段
type GamePhase string

const (
	PhaseGameStart       GamePhase = "GameStart"       // 游戏开始
	PhasePreparation     GamePhase = "Preparation"     // 准备阶段
	PhaseScriptSelection GamePhase = "ScriptSelection" // 选择剧本
	PhaseCharacterSetup  GamePhase = "CharacterSetup"  // 设定角色与事件
	PhaseLoop            GamePhase = "Loop"            // 循环阶段(包含多个循环)
	PhaseFinalGuess      GamePhase = "FinalGuess"      // 最终猜测阶段
	PhaseGameEnd         GamePhase = "GameEnd"         // 游戏结束
)

// LoopPhase 表示一个循环内的各个阶段
type LoopPhase string

const (
	PhaseLoopStart      LoopPhase = "LoopStart"      // 循环开始
	PhaseTimeSpiral     LoopPhase = "TimeSpiral"     // Time Spiral讨论阶段
	PhaseCharacterReset LoopPhase = "CharacterReset" // 角色归位阶段
	PhaseCountersReset  LoopPhase = "CountersReset"  // 移除和替换计数器阶段
	PhaseReturnCards    LoopPhase = "ReturnCards"    // 玩家取回卡牌阶段
	PhaseDay            LoopPhase = "Day"            // 每日流程阶段(包含多天)
	PhaseLoopEnd        LoopPhase = "LoopEnd"        // 循环结束，检查胜利条件
)

// DayPhase 表示一天中的各个阶段
type DayPhase string

const (
	PhaseDayStart            DayPhase = "DayStart"            // 日出阶段
	PhaseMastermindAction    DayPhase = "MastermindAction"    // Mastermind行动阶段
	PhaseProtagonistsAction  DayPhase = "ProtagonistsAction"  // 主角方行动阶段
	PhaseResolveCards        DayPhase = "ResolveCards"        // 卡牌结算阶段
	PhaseMastermindAbilities DayPhase = "MastermindAbilities" // Mastermind能力阶段
	PhaseLeaderGoodwill      DayPhase = "LeaderGoodwill"      // 领袖好感度能力阶段
	PhaseIncidents           DayPhase = "Incidents"           // 事件发生阶段
	PhaseSwitchLeader        DayPhase = "SwitchLeader"        // 切换领袖阶段
	PhaseDayEnd              DayPhase = "DayEnd"              // 日落阶段
)
