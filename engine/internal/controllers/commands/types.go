package commands

import "tragedy-looper/engine/internal/models"

type CommandType string

const (
	// CmdStartGame 启动游戏
	CmdStartGame CommandType = "start"
	// CmdPlaceCard 放置卡牌到角色或地点
	CmdPlaceCard CommandType = "place"
	// CmdPassAction 跳过当前操作/不执行行动
	CmdPassAction CommandType = "pass"

	// CmdShowCards 查看手牌
	CmdShowCards CommandType = "cards"
	// CmdShowBoard 查看场上情况
	CmdShowBoard CommandType = "board"
	// CmdStatus 查看角色/地点状态
	CmdStatus CommandType = "status"
	// CmdViewRules 查看当前脚本规则/角色信息
	CmdViewRules CommandType = "rules"
	// CmdViewIncidents 查看已发生的事件
	CmdViewIncidents CommandType = "incidents"
	// CmdViewHistory 查看历史记录/事件日志
	CmdViewHistory CommandType = "history"

	// CmdUseGoodwill 使用好感度能力
	CmdUseGoodwill CommandType = "goodwill"
	// CmdSelectChar 选择角色作为操作目标
	CmdSelectChar CommandType = "selectChar"
	// CmdSelectLocation 选择地点作为操作目标
	CmdSelectLocation CommandType = "selectLoc"
	// CmdFinalGuess 进行最终猜测（主角玩家）
	CmdFinalGuess CommandType = "guess"

	// CmdNextPhase 进入下一阶段
	CmdNextPhase CommandType = "next"
	// CmdEndTurn 结束当前回合
	CmdEndTurn CommandType = "end"

	// CmdMakeNote 添加个人笔记/标记
	CmdMakeNote CommandType = "note"

	// CmdHelp 显示帮助信息
	CmdHelp CommandType = "help"
	// CmdQuit 退出游戏
	CmdQuit CommandType = "quit"
)

// Command 命令接口
type Command interface {
	// Type 返回命令类型
	Type() CommandType

	// Validate 验证命令参数是否有效
	// 返回错误表示命令无效，返回nil表示命令有效
	Validate() error

	// RequiredInputs 返回命令需要的输入描述
	// 返回一个描述各个输入项的字符串切片，用于提示用户
	RequiredInputs() []string

	// Execute 执行命令
	Execute(context CommandContext) error
}

// CommandContext 命令执行上下文
type CommandContext struct {
	GameState     *models.GameState
	CurrentPlayer models.Player // 当前玩家字段
}
