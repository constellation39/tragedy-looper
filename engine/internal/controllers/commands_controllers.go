package controllers

import (
	"bufio"
	"fmt"
	"github.com/pterm/pterm"
	"go.uber.org/zap"
	"os"
	"strings"
	"tragedy-looper/engine/internal/controllers/commands"
	"tragedy-looper/engine/internal/models"
)

// CLI 控制器现在支持多种输入模式
type CLI struct {
	logging          *zap.Logger
	inputReader      *bufio.Reader
	cmdParser        *commands.CommandParser
	ui               UI       // UI接口
	cachedCards      []string // 缓存当前可用的卡片
	availableScripts []string // 可用剧本列表
	selectedScript   string   // 当前选中剧本
}

// NewCLI 创建新的CLI控制器（默认使用terminal UI）
func NewCLI(logger *zap.Logger) *CLI {
	return &CLI{
		logging:     logger,
		inputReader: bufio.NewReader(os.Stdin),
		cmdParser:   commands.NewCommandParser(),
		ui:          &TerminalUI{},
	}
}

// 选择目标（角色/位置）
func (cli *CLI) selectTarget(state *models.GameState) (models.TargetType, error) {
	// 第一步：选择目标类型
	category, err := cli.ui.Select("选择目标类型", []string{"角色", "位置"})
	if err != nil {
		return nil, err
	}

	// 第二步：展示对应类型的候选目标
	var options []string
	targetMap := make(map[int]models.TargetType)

	switch category {
	case 0: // 角色
		for i, c := range state.Script.Characters {
			options = append(options, string(c.Name))
			targetMap[i] = findCharacter(state, string(c.Name))
		}
	case 1: // 位置
		for i, loc := range state.Board.Locations() {
			options = append(options, string(loc))
			targetMap[i] = findLocation(state, string(loc))
		}
	default:
		return nil, fmt.Errorf("无效选择")
	}

	// 添加返回上级选项
	options = append(options, "返回上级菜单")

	selectedIdx, err := cli.ui.Select("选择目标 (输入序号)", options)
	if err != nil {
		return nil, err
	}

	if selectedIdx >= len(options)-1 { // 选择了返回
		return cli.selectTarget(state) // 递归调用重新选择类型
	}

	return targetMap[selectedIdx], nil
}

// 新增角色目标选择方法
func (cli *CLI) chooseCharacterTarget(state *models.GameState, filter func(models.TargetType) bool) (models.TargetType, error) {
	var options []string
	var targetMap = make(map[int]*models.Character)

	validChars := 0
	for _, c := range state.Script.Characters {
		char := findCharacter(state, string(c.Name))
		if filter == nil || filter(char) {
			options = append(options, string(c.Name))
			targetMap[validChars] = char
			validChars++
		}
	}
	options = append(options, "取消返回")

	selectedIdx, err := cli.ui.Select("选择角色", options)
	if err != nil || selectedIdx >= len(targetMap) {
		return nil, err
	}
	return targetMap[selectedIdx], nil
}

// 新增位置目标选择方法
func (cli *CLI) chooseLocationTarget(state *models.GameState, filter func(models.TargetType) bool) (models.TargetType, error) {
	var options []string
	var targetMap = make(map[int]*models.Location)

	validLocs := 0
	for _, loc := range state.Board.Locations() {
		location := findLocation(state, string(loc))
		if filter == nil || filter(location) {
			options = append(options, string(loc))
			targetMap[validLocs] = location
			validLocs++
		}
	}
	options = append(options, "取消返回")

	selectedIdx, err := cli.ui.Select("选择位置", options)
	if err != nil || selectedIdx >= len(targetMap) {
		return nil, err
	}
	return targetMap[selectedIdx], nil
}

// 新增多选目标方法
func (cli *CLI) selectMultipleTargets(state *models.GameState, prompt string, filter func(models.TargetType) bool) ([]models.TargetType, error) {
	var selectedTargets []models.TargetType

	for {
		// 第一步：选择本次要选取的目标类型
		typeChoice, err := cli.ui.Select("选择要操作的目标类型", []string{"添加角色目标", "添加位置目标", "完成选择"})
		if err != nil {
			return nil, err
		}

		switch typeChoice {
		case 0: // 添加角色
			target, err := cli.chooseCharacterTarget(state, filter)
			if err != nil {
				return nil, err
			}
			if target != nil {
				selectedTargets = append(selectedTargets, target)
			}

		case 1: // 添加位置
			target, err := cli.chooseLocationTarget(state, filter)
			if err != nil {
				return nil, err
			}
			if target != nil {
				selectedTargets = append(selectedTargets, target)
			}

		case 2: // 完成
			return selectedTargets, nil
		}

		// 显示当前已选择的
		cli.ui.ShowInfo(fmt.Sprintf("当前已选目标: %s", formatSelectedTargets(selectedTargets)))
	}
}

// 格式化已选择的目标为字符串，用于显示
func formatSelectedTargets(targets []models.TargetType) string {
	var result []string
	for _, target := range targets {
		switch t := target.(type) {
		case *models.Character:
			result = append(result, string(t.Name))
		case *models.Location:
			result = append(result, string(t.LocationType))
		}
	}
	return strings.Join(result, ", ")
}

// 辅助函数：查找角色
func findCharacter(state *models.GameState, name string) *models.Character {
	return state.Character(models.CharacterName(name))
}

// 辅助函数：查找位置
func findLocation(state *models.GameState, locType string) *models.Location {
	return state.Location(models.LocationType(locType))
}

// 处理命令行输入
func (cli *CLI) processInput(state *models.GameState) error {
	// 使用pterm的提示样式
	pterm.FgLightGreen.Print("> ")
	input, err := cli.inputReader.ReadString('\n')
	if err != nil {
		return err
	}

	input = strings.TrimSpace(input)
	if input == "" {
		return nil
	}

	// 解析命令
	cmd, err := cli.cmdParser.Parse(input)
	if err != nil {
		pterm.Error.Println("命令解析错误:", err)
		return nil
	}

	// 立即处理剧本选择命令
	if cmd.Type() == commands.CmdSelectScript {
		selectCmd := cmd.(*commands.SelectScriptCommand)
		cli.handleScriptSelection(selectCmd.ScriptName, state)
		return nil
	}

	// 创建执行上下文
	ctx := commands.CommandContext{
		GameState:     state,
		CurrentPlayer: getCurrentPlayer(state),
	}

	// 检查start命令时需先选择剧本
	if cmd.Type() == commands.CmdStartGame && cli.selectedScript == "" {
		pterm.Error.Println("错误: 请先使用selectScript选择剧本")
		return nil
	}

	// 根据命令类型动态处理
	switch cmd.Type() {
	case commands.CmdPlaceCard, commands.CmdUseGoodwill:
		return cli.handleTargetCommand(cmd, ctx)
	default:
		return cmd.Execute(ctx)
	}
}

// 处理需要目标选择的指令
func (cli *CLI) handleTargetCommand(cmd commands.Command, ctx commands.CommandContext) error {
	switch cmd.Type() {
	case commands.CmdPlaceCard:
		return cli.handleSingleTargetCommand(cmd, ctx)
	case commands.CmdUseGoodwill:
		return cli.handleMultiTargetCommand(cmd, ctx)
	default:
		return fmt.Errorf("未知指令类型: %s", cmd.Type())
	}
}

// 处理单目标指令
func (cli *CLI) handleSingleTargetCommand(cmd commands.Command, ctx commands.CommandContext) error {
	//target, err := cli.selectTarget(ctx.GameState)
	//if err != nil {
	//	return err
	//}
	//
	//// 设置目标到命令并执行
	//switch c := cmd.(type) {
	//case *commands.PlaceCardCommand:
	//	c.SetTarget(targetIdentifier(target))
	//	return c.Execute(ctx)
	//case *commands.GoodwillCommand: // 如果有单目标使用场景
	//	c.SetTarget(targetIdentifier(target))
	//	return c.Execute(ctx)
	//}

	return fmt.Errorf("不支持的指令类型: %s", cmd.Type())
}

// 处理多目标指令
func (cli *CLI) handleMultiTargetCommand(cmd commands.Command, ctx commands.CommandContext) error {
	//targets, err := cli.selectMultipleTargets(ctx.GameState, "请选择多个目标（空格确认）", nil)
	//if err != nil {
	//	return err
	//}
	//
	//switch c := cmd.(type) {
	//case *commands.GoodwillCommand:
	//	c.SetTargets(mapTargetIdentifiers(targets))
	//	return c.Execute(ctx)
	//}

	return fmt.Errorf("不支持的指令类型: %s", cmd.Type())
}

// 将目标转换为标识字符串
func targetIdentifier(target models.TargetType) string {
	switch t := target.(type) {
	case *models.Character:
		return fmt.Sprintf("char_%s", t.Name)
	case *models.Location:
		return fmt.Sprintf("loc_%s", t.LocationType)
	default:
		return "unknown_target"
	}
}

// 批量转换目标标识
func mapTargetIdentifiers(targets []models.TargetType) []string {
	var ids []string
	for _, t := range targets {
		ids = append(ids, targetIdentifier(t))
	}
	return ids
}

// 新增剧本选择处理逻辑
func (cli *CLI) handleScriptSelection(name string, state *models.GameState) {
	// 验证剧本是否存在于可用列表中
	scriptFound := false
	for _, script := range cli.availableScripts {
		if script == name {
			scriptFound = true
			break
		}
	}

	if !scriptFound {
		pterm.Error.Printf("错误: 剧本 [%s] 不存在，请选择有效的剧本\n", name)
		return
	}

	// 这里应该根据name加载对应剧本（需扩展实际剧本加载逻辑）
	cli.selectedScript = name
	pterm.Success.Printf("剧本 [%s] 已选择\n", name)
}

// getCurrentPlayer 从游戏状态获取当前玩家
// 此处为示例实现，需根据实际游戏状态逻辑进行调整
func getCurrentPlayer(state *models.GameState) models.Player {
	return state.CurrentPlayer
}

// 以下是原来各种展示函数的实现，现在可以直接使用command包中的命令

// ShowCards 展示卡牌信息
func (cli *CLI) ShowCards(state *models.GameState) {
	cmd := commands.NewShowCardsCommand()
	ctx := commands.CommandContext{
		GameState:     state,
		CurrentPlayer: getCurrentPlayer(state),
	}
	_ = cmd.Execute(ctx)
}

// ShowBoard 展示游戏板
func (cli *CLI) ShowBoard(state *models.GameState) {
	cmd := commands.NewShowBoardCommand()
	ctx := commands.CommandContext{
		GameState: state,
	}
	_ = cmd.Execute(ctx)
}

// ShowCharacters 展示角色信息
func (cli *CLI) ShowCharacters(state *models.GameState) {
	//// 使用pterm表格呈现角色信息
	//data := [][]string{
	//	{"角色", "位置", "状态", "属性"},
	//}
	//
	//for _, c := range state.Script.Characters {
	//	char := state.Character(c.Name)
	//	if char == nil {
	//		continue
	//	}
	//
	//	row := []string{
	//		string(c.Name),
	//		string(char.Location()),
	//		fmt.Sprintf("❤%d", char.PhysicalHealth),
	//		fmt.Sprintf("疑%d 善%d", char.State.Paranoia, char.State.Goodwill),
	//	}
	//	data = append(data, row)
	//}
	//
	//pterm.DefaultTable.
	//	WithHasHeader(true).
	//	WithBoxed(true).
	//	WithData(data).
	//	Render()
}

// ShowLocations 展示位置信息
func (cli *CLI) ShowLocations(state *models.GameState) {
	for _, loc := range state.Board.Locations() {
		location := state.Location(loc)
		if location == nil {
			continue
		}

		pterm.DefaultSection.Printf("位置: %s\n", loc)
		pterm.DefaultBarChart.
			WithBars([]pterm.Bar{
				{Label: "阴谋标记", Value: location.Intrigue()},
			}).
			Render()

		if len(location.Characters) > 0 {
			var charNames []string
			for name := range location.Characters {
				charNames = append(charNames, string(name))
			}
			cli.ui.ShowInfo(fmt.Sprintf("在场角色：%s", strings.Join(charNames, ", ")))
		}
	}
}

// ShowGameInfo 展示游戏信息
func (cli *CLI) ShowGameInfo(state *models.GameState) {
	//panels := pterm.Panels{
	//	{
	//		{Data: pterm.Sprintf("剧本: %s", pterm.LightMagenta(state.Script.Title))},
	//		{Data: pterm.Sprintf("当前阶段: %s", pterm.Cyan(state.CurrentPhase))},
	//	},
	//	{
	//		{Data: pterm.Sprintf("循环进度: %s",
	//			pterm.DefaultProgressbar.
	//				WithTotal(state.Script.MaxLoops).
	//				WithCurrent(state.CurrentLoop).
	//				WithTitle("循环"))},
	//		{Data: pterm.Sprintf("当前玩家: %s", state.CurrentPlayer.Name())},
	//	},
	//}
	//
	//pterm.DefaultPanel.
	//	WithPanels(panels).
	//	WithPadding(1).
	//	WithSameColumnWidth(true).
	//	Render()
	//
	//if state.IsGameOver {
	//	pterm.DefaultHeader.
	//		WithBackgroundStyle(pterm.NewStyle(pterm.BgRed)).
	//		WithTextStyle(pterm.NewStyle(pterm.FgWhite)).
	//		Println(fmt.Sprintf("游戏结束! 获胜方: %s", state.WinnerType))
	//}
}

// 运行CLI
func (cli *CLI) Run(state *models.GameState) error {
	// 初始化可用剧本列表（应该从配置文件或目录加载）
	cli.availableScripts = []string{"新手教学", "第一幕", "校园谜案"}

	// 使用pterm的标题样式
	pterm.DefaultHeader.
		WithBackgroundStyle(pterm.NewStyle(pterm.BgLightBlue)).
		WithTextStyle(pterm.NewStyle(pterm.FgBlack)).
		Println("悲剧循环游戏控制台")

	// 使用带编号的列表显示剧本
	scriptsList := pterm.DefaultBulletList.WithItems(make([]pterm.BulletListItem, 0))
	for i, script := range cli.availableScripts {
		scriptsList.Items = append(scriptsList.Items, pterm.BulletListItem{
			Level: 0,
			Text:  fmt.Sprintf("%d. %s", i+1, script),
		})
	}
	scriptsList.Render()

	pterm.Info.Println("使用 selectScript <剧本名称> 选择剧本")
	pterm.Info.Println("输入 'help' 获取可用命令列表")

	for !state.IsGameOver {
		if err := cli.processInput(state); err != nil {
			pterm.Error.Println(err)
			return err
		}
	}

	pterm.Success.Println("游戏结束!")
	pterm.Success.Printf("获胜方: %s\n", state.WinnerType)

	return nil
}
