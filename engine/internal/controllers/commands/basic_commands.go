package commands

import (
	"fmt"
	"tragedy-looper/engine/internal/models"
)

// PassActionCommand 跳过当前操作/不执行行动
type PassActionCommand struct{}

func NewPassActionCommand() *PassActionCommand {
	return &PassActionCommand{}
}

func (c *PassActionCommand) Type() CommandType {
	return CmdPassAction
}

func (c *PassActionCommand) Validate() error {
	return nil // 无需参数，总是有效
}

func (c *PassActionCommand) RequiredInputs() []string {
	return []string{} // 无需输入
}

func (c *PassActionCommand) Execute(ctx CommandContext) error {
	fmt.Println("跳过当前操作")
	return nil
}

// ShowCardsCommand 查看手牌
type ShowCardsCommand struct{}

func NewShowCardsCommand() *ShowCardsCommand {
	return &ShowCardsCommand{}
}

func (c *ShowCardsCommand) Type() CommandType {
	return CmdShowCards
}

func (c *ShowCardsCommand) Validate() error {
	return nil // 无需参数，总是有效
}

func (c *ShowCardsCommand) RequiredInputs() []string {
	return []string{} // 无需输入
}

func (c *ShowCardsCommand) Execute(ctx CommandContext) error {
	if ctx.CurrentPlayer == nil {
		return fmt.Errorf("当前玩家信息缺失")
	}

	fmt.Println("=== 手牌列表 ===")
	for _, card := range ctx.CurrentPlayer.GetHandCards() {
		fmt.Printf("- %s (ID: %s)\n", card.Name(), card.Id())
	}
	fmt.Println("===============")

	return nil
}

// ShowBoardCommand 查看场上情况
type ShowBoardCommand struct{}

func NewShowBoardCommand() *ShowBoardCommand {
	return &ShowBoardCommand{}
}

func (c *ShowBoardCommand) Type() CommandType {
	return CmdShowBoard
}

func (c *ShowBoardCommand) Validate() error {
	return nil // 无需参数，总是有效
}

func (c *ShowBoardCommand) RequiredInputs() []string {
	return []string{} // 无需输入
}

func (c *ShowBoardCommand) Execute(ctx CommandContext) error {
	if ctx.GameState == nil {
		return fmt.Errorf("游戏状态信息缺失")
	}

	fmt.Println("=== 游戏板 ===")
	fmt.Printf("当前循环: %d, 当前日期: %d, 当前阶段: %s\n",
		ctx.GameState.CurrentLoop, ctx.GameState.CurrentDay, ctx.GameState.CurrentPhase)

	// 显示位置及角色
	for _, loc := range ctx.GameState.Board.Locations() {
		location := ctx.GameState.Location(loc)
		if location == nil {
			continue
		}
		fmt.Printf("[%s] 阴谋:%d\n", loc, location.CurIntrigue)
		for name := range location.Characters {
			fmt.Printf("  - %s\n", name)
		}
	}
	fmt.Println("=============")

	return nil
}

// StatusCommand 查看角色/地点状态
type StatusCommand struct {
	Target string
}

func NewStatusCommand(target string) *StatusCommand {
	return &StatusCommand{
		Target: target,
	}
}

func (c *StatusCommand) Type() CommandType {
	return CmdStatus
}

func (c *StatusCommand) Validate() error {
	if c.Target == "" {
		return fmt.Errorf("目标不能为空")
	}
	return nil
}

func (c *StatusCommand) RequiredInputs() []string {
	return []string{
		"目标: 要查看状态的角色或位置",
	}
}

func (c *StatusCommand) Execute(ctx CommandContext) error {
	if ctx.GameState == nil {
		return fmt.Errorf("游戏状态信息缺失")
	}

	// 尝试作为角色查找
	character := ctx.GameState.Character(models.CharacterName(c.Target))
	if character != nil {
		fmt.Printf("=== 角色状态: %s ===\n", c.Target)
		fmt.Printf("位置: %s\n", character.Location())
		if character.Role() != nil {
			fmt.Printf("角色类型: %s\n", character.Role().RoleType())
		}
		fmt.Println("========================")
		return nil
	}

	// 尝试作为位置查找
	location := ctx.GameState.Location(models.LocationType(c.Target))
	if location != nil {
		fmt.Printf("=== 位置状态: %s ===\n", c.Target)
		fmt.Printf("阴谋标记: %d\n", location.CurIntrigue)
		if len(location.Characters) > 0 {
			fmt.Println("角色:")
			for name := range location.Characters {
				fmt.Printf("  - %s\n", name)
			}
		}
		fmt.Println("========================")
		return nil
	}

	return fmt.Errorf("找不到目标: %s", c.Target)
}

// ViewRulesCommand 查看当前脚本规则/角色信息
type ViewRulesCommand struct{}

func NewViewRulesCommand() *ViewRulesCommand {
	return &ViewRulesCommand{}
}

func (c *ViewRulesCommand) Type() CommandType {
	return CmdViewRules
}

func (c *ViewRulesCommand) Validate() error {
	return nil // 无需参数，总是有效
}

func (c *ViewRulesCommand) RequiredInputs() []string {
	return []string{} // 无需输入
}

func (c *ViewRulesCommand) Execute(ctx CommandContext) error {
	if ctx.GameState == nil {
		return fmt.Errorf("游戏状态信息缺失")
	}

	fmt.Println("=== 脚本规则 ===")
	fmt.Printf("剧本: %s\n", ctx.GameState.Script.Title)
	fmt.Printf("最大循环数: %d\n", ctx.GameState.Script.MaxLoops)
	fmt.Println("===============")

	return nil
}

// ViewIncidentsCommand 查看已发生的事件
type ViewIncidentsCommand struct{}

func NewViewIncidentsCommand() *ViewIncidentsCommand {
	return &ViewIncidentsCommand{}
}

func (c *ViewIncidentsCommand) Type() CommandType {
	return CmdViewIncidents
}

func (c *ViewIncidentsCommand) Validate() error {
	return nil // 无需参数，总是有效
}

func (c *ViewIncidentsCommand) RequiredInputs() []string {
	return []string{} // 无需输入
}

func (c *ViewIncidentsCommand) Execute(ctx CommandContext) error {
	if ctx.GameState == nil {
		return fmt.Errorf("游戏状态信息缺失")
	}

	fmt.Println("=== 已发生事件 ===")
	// TODO: 实际实现逻辑
	fmt.Println("=================")

	return nil
}

// ViewHistoryCommand 查看历史记录/事件日志
type ViewHistoryCommand struct{}

func NewViewHistoryCommand() *ViewHistoryCommand {
	return &ViewHistoryCommand{}
}

func (c *ViewHistoryCommand) Type() CommandType {
	return CmdViewHistory
}

func (c *ViewHistoryCommand) Validate() error {
	return nil // 无需参数，总是有效
}

func (c *ViewHistoryCommand) RequiredInputs() []string {
	return []string{} // 无需输入
}

func (c *ViewHistoryCommand) Execute(ctx CommandContext) error {
	if ctx.GameState == nil {
		return fmt.Errorf("游戏状态信息缺失")
	}

	fmt.Println("=== 历史记录 ===")
	// TODO: 实际实现逻辑
	fmt.Println("===============")

	return nil
}

// FinalGuessCommand 进行最终猜测（主角玩家）
type FinalGuessCommand struct{}

func NewFinalGuessCommand() *FinalGuessCommand {
	return &FinalGuessCommand{}
}

func (c *FinalGuessCommand) Type() CommandType {
	return CmdFinalGuess
}

func (c *FinalGuessCommand) Validate() error {
	return nil // 初始无参数，参数将通过交互获取
}

func (c *FinalGuessCommand) RequiredInputs() []string {
	return []string{
		"将通过交互方式选择最终猜测的角色身份",
	}
}

func (c *FinalGuessCommand) Execute(ctx CommandContext) error {
	if ctx.GameState == nil {
		return fmt.Errorf("游戏状态信息缺失")
	}

	if ctx.CurrentPlayer == nil {
		return fmt.Errorf("当前玩家信息缺失")
	}

	// 检查是否是主角玩家
	if ctx.CurrentPlayer.Role() != "protagonist" {
		return fmt.Errorf("只有主角玩家才能进行最终猜测")
	}

	// 检查是否是最后一个循环
	if ctx.GameState.CurrentLoop != ctx.GameState.Script.MaxLoops {
		return fmt.Errorf("最终猜测只能在最后一个循环中进行")
	}

	fmt.Println("=== 开始最终猜测 ===")
	// 获取所有角色列表
	characters := make([]string, 0, len(ctx.GameState.Script.Characters))
	for _, c := range ctx.GameState.Script.Characters {
		characters = append(characters, string(c.Name))
	}

	// TODO: 此处应该实现一个交互式的角色猜测流程
	// 例如提示玩家为每个角色选择身份

	fmt.Println("最终猜测已记录，等待游戏结束后公布结果")
	fmt.Println("===================")

	return nil
}

// NextPhaseCommand 进入下一阶段
type NextPhaseCommand struct{}

func NewNextPhaseCommand() *NextPhaseCommand {
	return &NextPhaseCommand{}
}

func (c *NextPhaseCommand) Type() CommandType {
	return CmdNextPhase
}

func (c *NextPhaseCommand) Validate() error {
	return nil // 无需参数，总是有效
}

func (c *NextPhaseCommand) RequiredInputs() []string {
	return []string{} // 无需输入
}

func (c *NextPhaseCommand) Execute(ctx CommandContext) error {
	if ctx.GameState == nil {
		return fmt.Errorf("游戏状态信息缺失")
	}

	// TODO: 实现进入下一阶段的逻辑
	fmt.Println("进入下一阶段")
	return nil
}

// EndTurnCommand 结束当前回合
type EndTurnCommand struct{}

func NewEndTurnCommand() *EndTurnCommand {
	return &EndTurnCommand{}
}

func (c *EndTurnCommand) Type() CommandType {
	return CmdEndTurn
}

func (c *EndTurnCommand) Validate() error {
	return nil // 无需参数，总是有效
}

func (c *EndTurnCommand) RequiredInputs() []string {
	return []string{} // 无需输入
}

func (c *EndTurnCommand) Execute(ctx CommandContext) error {
	if ctx.GameState == nil {
		return fmt.Errorf("游戏状态信息缺失")
	}

	// TODO: 实现结束当前回合的逻辑
	fmt.Println("结束当前回合")
	return nil
}

// MakeNoteCommand 添加个人笔记/标记
type MakeNoteCommand struct {
	Content string
}

func NewMakeNoteCommand(content string) *MakeNoteCommand {
	return &MakeNoteCommand{
		Content: content,
	}
}

func (c *MakeNoteCommand) Type() CommandType {
	return CmdMakeNote
}

func (c *MakeNoteCommand) Validate() error {
	if c.Content == "" {
		return fmt.Errorf("笔记内容不能为空")
	}
	return nil
}

func (c *MakeNoteCommand) RequiredInputs() []string {
	return []string{
		"内容: 笔记内容",
	}
}

func (c *MakeNoteCommand) Execute(ctx CommandContext) error {
	if ctx.CurrentPlayer == nil {
		return fmt.Errorf("当前玩家信息缺失")
	}

	// TODO: 实现保存笔记的逻辑
	fmt.Printf("已保存笔记: %s\n", c.Content)
	return nil
}

// HelpCommand 显示帮助信息
type HelpCommand struct{}

func NewHelpCommand() *HelpCommand {
	return &HelpCommand{}
}

func (c *HelpCommand) Type() CommandType {
	return CmdHelp
}

func (c *HelpCommand) Validate() error {
	return nil // 无需参数，总是有效
}

func (c *HelpCommand) RequiredInputs() []string {
	return []string{} // 无需输入
}

func (c *HelpCommand) Execute(ctx CommandContext) error {
	fmt.Println("=== 命令帮助 ===")
	fmt.Println("place <cardID> - 放置卡牌")
	fmt.Println("selectChar <characterName> - 选择角色作为目标")
	fmt.Println("selectLoc <locationType> - 选择位置作为目标")
	fmt.Println("pass - 跳过当前操作")
	fmt.Println("cards - 查看手牌")
	fmt.Println("board - 查看场上情况")
	fmt.Println("status <target> - 查看角色或位置状态")
	fmt.Println("rules - 查看当前脚本规则")
	fmt.Println("incidents - 查看已发生的事件")
	fmt.Println("history - 查看历史记录")
	fmt.Println("goodwill <characterName> <abilityID> - 使用好感度能力")
	fmt.Println("guess - 进行最终猜测")
	fmt.Println("next - 进入下一阶段")
	fmt.Println("end - 结束当前回合")
	fmt.Println("note <content> - 添加个人笔记")
	fmt.Println("help - 显示本帮助")
	fmt.Println("quit - 退出游戏")
	fmt.Println("===============")
	return nil
}

// QuitCommand 退出游戏
type QuitCommand struct{}

func NewQuitCommand() *QuitCommand {
	return &QuitCommand{}
}

func (c *QuitCommand) Type() CommandType {
	return CmdQuit
}

func (c *QuitCommand) Validate() error {
	return nil // 无需参数，总是有效
}

func (c *QuitCommand) RequiredInputs() []string {
	return []string{} // 无需输入
}

func (c *QuitCommand) Execute(ctx CommandContext) error {
	fmt.Println("正在退出游戏...")
	// TODO: 实现退出游戏的逻辑
	return nil
}
