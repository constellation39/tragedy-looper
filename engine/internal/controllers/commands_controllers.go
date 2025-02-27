package controllers

import (
	"bufio"
	"fmt"
	"go.uber.org/zap"
	"os"
	"strings"
	"tragedy-looper/engine/internal/controllers/commands"
	"tragedy-looper/engine/internal/models"
)

// CLI 控制器现在支持多种输入模式
type CLI struct {
	logging     *zap.Logger
	inputReader *bufio.Reader
	cmdParser   *commands.CommandParser
	ui          UI       // UI接口
	cachedCards []string // 缓存当前可用的卡片
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
	characters := make([]string, 0, len(state.Script.Characters))
	locations := make([]string, 0, len(state.Board.Locations()))

	for _, c := range state.Script.Characters {
		characters = append(characters, fmt.Sprintf("char_%s", c.Name))
	}
	for _, loc := range state.Board.Locations() {
		locations = append(locations, fmt.Sprintf("loc_%s", loc))
	}

	options := append([]string{"-- 角色 --"}, characters...)
	options = append(options, "-- 位置 --")
	options = append(options, locations...)

	selectedIdx, err := cli.ui.Select("选择目标", options)
	if err != nil {
		return nil, err
	}

	if selectedIdx <= len(characters) {
		return findCharacter(state, strings.TrimPrefix(options[selectedIdx], "char_")), nil
	}
	return findLocation(state, strings.TrimPrefix(options[selectedIdx], "loc_")), nil
}

// 新增多选目标方法
func (cli *CLI) selectMultipleTargets(state *models.GameState, prompt string, filter func(models.TargetType) bool) ([]models.TargetType, error) {
	var selectedTargets []models.TargetType
	var availableOptions []string
	var targetMap = make(map[int]models.TargetType)

	// 准备可选目标列表
	characters := make([]string, 0, len(state.Script.Characters))
	locations := make([]string, 0, len(state.Board.Locations()))

	// 处理角色选项
	idx := 0
	for _, c := range state.Script.Characters {
		charTarget := findCharacter(state, string(c.Name))
		if filter == nil || filter(charTarget) {
			option := fmt.Sprintf("char_%s", c.Name)
			characters = append(characters, option)
			targetMap[idx] = charTarget
			idx++
		}
	}

	// 处理位置选项
	for _, loc := range state.Board.Locations() {
		locTarget := findLocation(state, string(loc))
		if filter == nil || filter(locTarget) {
			option := fmt.Sprintf("loc_%s", loc)
			locations = append(locations, option)
			targetMap[idx] = locTarget
			idx++
		}
	}

	// 合并选项列表
	availableOptions = append(availableOptions, characters...)
	availableOptions = append(availableOptions, locations...)
	availableOptions = append(availableOptions, "完成选择") // 添加完成选项

	// 处理循环选择
	for {
		if len(availableOptions) <= 1 { // 只剩"完成选择"时退出
			break
		}

		// 显示已选择的目标
		selectionPrompt := prompt
		if len(selectedTargets) > 0 {
			selectionPrompt += fmt.Sprintf("\n已选择: %v", formatSelectedTargets(selectedTargets))
		}

		selectedIdx, err := cli.ui.Select(selectionPrompt, availableOptions)
		if err != nil {
			return nil, err
		}

		// 检查是否完成选择
		if selectedIdx == len(availableOptions)-1 { // "完成选择"选项
			break
		}

		// 添加到已选列表
		selectedTarget := targetMap[selectedIdx]
		selectedTargets = append(selectedTargets, selectedTarget)

		// 从可选列表中移除已选项
		availableOptions = append(availableOptions[:selectedIdx], availableOptions[selectedIdx+1:]...)
		for i := selectedIdx; i < len(targetMap)-1; i++ {
			targetMap[i] = targetMap[i+1]
		}
		delete(targetMap, len(targetMap)-1)

		// 询问是否继续选择
		if len(availableOptions) <= 1 { // 只剩"完成选择"时自动退出
			break
		}
	}

	return selectedTargets, nil
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
	fmt.Print("> ")
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
		fmt.Println("命令解析错误:", err)
		return nil
	}

	// 创建命令上下文
	ctx := commands.CommandContext{
		GameState: state,
		// 假设当前玩家可从游戏状态获取，或通过参数传入
		CurrentPlayer: getCurrentPlayer(state),
	}

	// 执行命令
	if err := cmd.Execute(ctx); err != nil {
		fmt.Println("命令执行错误:", err)
	}

	return nil
}

// getCurrentPlayer 从游戏状态获取当前玩家
// 此处为示例实现，需根据实际游戏状态逻辑进行调整
func getCurrentPlayer(state *models.GameState) models.Player {
	// 假设当前玩家可从游戏状态获取
	if len(state.Players) > 0 {
		return state.Players[state.CurrentPlayerIndex]
	}
	return nil
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
	// 这个功能现在可以使用status命令实现
	// 但为了保持API兼容性，这里保留此方法
	fmt.Println("=== 角色列表 ===")
	for _, c := range state.Script.Characters {
		character := state.Character(c.Name)
		if character == nil {
			continue
		}
		fmt.Printf("- %s (位置: %s)\n", c.Name, character.Location())
		if character.Role() != nil {
			fmt.Printf("  角色: %s\n", character.Role().RoleType())
		}
	}
	fmt.Println("===============")
}

// ShowLocations 展示位置信息
func (cli *CLI) ShowLocations(state *models.GameState) {
	// 这个功能现在可以使用status命令实现
	// 但为了保持API兼容性，这里保留此方法
	fmt.Println("=== 位置列表 ===")
	for _, loc := range state.Board.Locations() {
		location := state.Location(loc)
		if location == nil {
			continue
		}
		fmt.Printf("- %s (阴谋标记: %d)\n", loc, location.CurIntrigue)
		if len(location.Characters) > 0 {
			fmt.Print("  角色: ")
			for name := range location.Characters {
				fmt.Printf("%s ", name)
			}
			fmt.Println()
		}
	}
	fmt.Println("===============")
}

// ShowGameInfo 展示游戏信息
func (cli *CLI) ShowGameInfo(state *models.GameState) {
	fmt.Println("=== 游戏信息 ===")
	fmt.Printf("剧本: %s\n", state.Script.Title)
	fmt.Printf("循环: %d/%d\n", state.CurrentLoop, state.Script.MaxLoops)
	fmt.Printf("日期: %d\n", state.CurrentDay)
	fmt.Printf("阶段: %s\n", state.CurrentPhase)
	if state.IsGameOver {
		fmt.Printf("游戏结束, 获胜方: %s\n", state.WinnerType)
	}
	fmt.Println("===============")
}

// 运行CLI
func (cli *CLI) Run(state *models.GameState) error {
	fmt.Println("=== 悲剧循环游戏控制台 ===")
	fmt.Println("输入 'help' 获取可用命令列表")

	for !state.IsGameOver {
		if err := cli.processInput(state); err != nil {
			return err
		}
	}

	fmt.Println("游戏结束!")
	fmt.Printf("获胜方: %s\n", state.WinnerType)

	return nil
}
