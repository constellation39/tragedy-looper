package commands

import "tragedy-looper/engine/internal/models"

type CommandType string

const (
	// 卡牌操作相关
	CmdPlaceCard CommandType = "place" // 放置卡牌
	CmdShowCards CommandType = "cards" // 查看手牌
	CmdShowBoard CommandType = "board" // 查看场上情况

	// 能力相关
	CmdUseGoodwill CommandType = "goodwill" // 使用好感度能力

	// 游戏流程相关
	CmdNextPhase CommandType = "next" // 进入下一阶段
	CmdEndTurn   CommandType = "end"  // 结束当前回合

	// 信息查询
	CmdStatus CommandType = "status" // 查看状态
	CmdHelp   CommandType = "help"   // 显示帮助
	CmdQuit   CommandType = "quit"   // 退出游戏
)

// Command 命令接口
type Command interface {
	Type() CommandType
	Execute(context CommandContext) error
}

// CommandContext 命令执行上下文
type CommandContext struct {
	IsLeader      bool
	IsMastermind  bool
	GameState     *models.GameState
	CurrentPlayer models.Player // 新增当前玩家字段
}
