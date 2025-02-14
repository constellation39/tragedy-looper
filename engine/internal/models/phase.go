package models

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
