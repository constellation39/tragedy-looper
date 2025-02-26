package controllers

import (
	"bufio"
	"fmt"
	"go.uber.org/zap"
	"os"
	"strconv"
	"strings"
	"tragedy-looper/engine/internal/controllers/commands"
	"tragedy-looper/engine/internal/models"
)

// UI 接口用于多种交互模式
type UI interface {
	Select(title string, options []string) (int, error)
	MultiSelect(title string, options []string) ([]int, error) // 新增多选接口
}

// CLI 控制器现在支持多种输入模式
type CLI struct {
	logging     *zap.Logger
	inputReader *bufio.Reader
	cmdParser   *commands.CommandParser
	ui          UI       // 新增UI接口
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

// 格式化目标参数
func (cli *CLI) formatTarget(target models.TargetType) string {
	switch t := target.(type) {
	case *models.Character:
		return fmt.Sprintf("char_%s", t.Name)
	case *models.Location:
		return fmt.Sprintf("loc_%s", t.LocationType)
	}
	return ""
}

func (cli *CLI) Init() {
	cli.logging.Info("欢迎来到Tragedy Looper!")
	cli.logging.Info("可用命令:")
	cli.logging.Info("=== 卡牌操作 ===")
	cli.logging.Info("- place <卡牌ID> <char_角色ID|loc_位置ID>  放置行动卡")
	cli.logging.Info("- resolve  结算所有已放置的卡牌")

	cli.logging.Info("\n=== 能力操作 ===")
	cli.logging.Info("- ability <角色名> <能力ID> <目标>  使用角色能力")
	cli.logging.Info("- goodwill <角色名> <能力ID> <目标>  使用好感度能力(仅领袖)")

	cli.logging.Info("\n=== 信息查看 ===")
	cli.logging.Info("- status [角色名]  查看状态")
	cli.logging.Info("- cards  查看手牌")
	cli.logging.Info("- board  查看场上情况")

	cli.logging.Info("\n=== 其他 ===")
	cli.logging.Info("- help  显示帮助")
	cli.logging.Info("- exit/quit  退出游戏")
}

// handleGoodwillCommand 处理好感度能力执行
func (cli *CLI) handleGoodwillCommand(cmd *commands.GoodwillCommand, state *models.GameState, player models.Player) error {
	// 验证是否是领袖
	if _, isLeader := player.(*models.Protagonist); !isLeader {
		return fmt.Errorf("只有领袖可以使用好感度能力")
	}

	// 获取角色和能力
	character := state.Character(models.CharacterName(cmd.CharacterName))
	if character == nil {
		return fmt.Errorf("角色 %s 不存在", cmd.CharacterName)
	}

	abilityIdx, err := strconv.Atoi(cmd.AbilityID)
	if err != nil || abilityIdx < 0 || abilityIdx >= len(character.GoodwillAbilityList) {
		return fmt.Errorf("无效的能力ID: %s", cmd.AbilityID)
	}

	ability := character.GoodwillAbilityList[abilityIdx]

	// 根据能力类型处理
	switch ability.Name {
	case "减少同位置学生1点不安":
		// 筛选同位置的学生角色
		filter := func(target models.TargetType) bool {
			if char, ok := target.(*models.Character); ok {
				return char.CurrentLocation == character.CurrentLocation &&
					char.ExistsTag("Student") &&
					char.Paranoia() > 0
			}
			return false
		}

		targets, err := cli.selectMultipleTargets(state, "选择要减少不安的学生", filter)
		if err != nil {
			return err
		}

		for _, target := range targets {
			if char, ok := target.(*models.Character); ok {
				char.SetParanoia(char.Paranoia() - 1)
				cli.logging.Info(fmt.Sprintf("减少了学生 %s 的不安", char.Name))
			}
		}

	default:
		target, err := cli.selectTarget(state)
		if err != nil {
			return err
		}

		// 创建执行上下文
		ctx := commands.CommandContext{
			GameState:     state,
			CurrentPlayer: player,
			IsLeader:      true,
		}

		// 执行能力效果
		if err := ability.Effect(ctx.GameState); err != nil {
			return err
		}
	}

	return nil
}

func (cli *CLI) handleInput(gameState *models.GameState) {
	for {
		fmt.Printf("[命令输入] > ") // 添加输入提示符
		input, err := cli.inputReader.ReadString('\n')
		if err != nil {
			cli.logging.Error("Failed to read input", zap.Error(err))
			continue
		}

		cmd, err := cli.cmdParser.Parse(strings.TrimSpace(input))
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			continue
		}

		// 执行命令
		ctx := commands.CommandContext{
			GameState: gameState,
		}

		if err := cmd.Execute(ctx); err != nil {
			fmt.Printf("Error executing command: %v\n", err)
			continue
		}
	}
}

// 处理能力命令

// findCard 根据卡牌ID查找玩家手牌中的卡牌
func findCard(player models.Player, cardID string) models.Card {
	for _, card := range player.GetHandCards() {
		if card.Id() == cardID {
			return card
		}
	}
	return nil
}

// findTarget 根据目标名称查找游戏中的目标（角色或位置）
func findTarget(gameState *models.GameState, targetInput string) models.TargetType {
	// 使用前缀区分目标类型
	parts := strings.SplitN(targetInput, "_", 2)
	if len(parts) != 2 {
		return nil
	}

	switch parts[0] {
	case "char":
		// 查找角色
		return findCharacter(gameState, parts[1])
	case "loc":
		// 查找位置
		return findLocation(gameState, parts[1])
	default:
		return nil
	}
}

// findCharacter 查找角色目标
func findCharacter(gameState *models.GameState, characterName string) models.TargetType {
	for _, c := range gameState.Script.Characters {
		if c.Name == models.CharacterName(characterName) {
			return c
		}
	}
	return nil
}

// findLocation 查找位置目标
func findLocation(gameState *models.GameState, locationID string) models.TargetType {
	return gameState.Board.GetLocation(models.LocationType(locationID))
}

// StartPlacementPhase 启动交互式卡牌放置阶段
func (cli *CLI) StartPlacementPhase(player models.Player, state *models.GameState) error {
	// 根据玩家类型设置最大选择次数
	var maxSelections int
	switch player.(type) {
	case *models.Mastermind:
		maxSelections = 3
	case *models.Protagonist:
		maxSelections = 1
	default:
		return fmt.Errorf("未知的玩家类型")
	}

	handCards := player.GetHandCards()

	// 缓存当前可用卡牌（实时更新）
	cli.cachedCards = make([]string, 0, len(handCards))
	for _, card := range handCards {
		cli.cachedCards = append(cli.cachedCards, card.Id())
	}

	// 追踪已选次数
	selectionsMade := 0
	for selectionsMade < maxSelections {
		if len(cli.cachedCards) == 0 {
			cli.logging.Error("无法继续放置: 没有可用卡牌",
				zap.String("阶段", "放置阶段"),
				zap.String("玩家类型", fmt.Sprintf("%T", player)),
				zap.Int("已选择次数", selectionsMade),
				zap.Int("最大次数", maxSelections))
			return fmt.Errorf("没有可用卡牌了")
		}

		// 显示剩余选择次数
		cli.logging.Info(fmt.Sprintf("剩余需要放置的卡牌数：%d/%d", maxSelections-selectionsMade, maxSelections))

		// 选择卡牌
		selectedIdx, err := cli.ui.Select("选择要放置的卡牌", cli.cachedCards)
		if err != nil {
			cli.logging.Error("选择卡牌失败", zap.Error(err))
			continue
		}
		selectedCardID := cli.cachedCards[selectedIdx]

		// 选择目标
		target, err := cli.selectTarget(state)
		if err != nil {
			cli.logging.Error("选择目标失败", zap.Error(err))
			continue
		}

		// 构造和验证命令
		cmdStr := fmt.Sprintf("place %s %s", selectedCardID, cli.formatTarget(target))
		cmd, err := cli.cmdParser.Parse(cmdStr)
		if err != nil {
			cli.logging.Error("解析命令失败", zap.String("cmdStr", cmdStr), zap.Error(err))
			continue
		}

		// 执行放置命令
		switch cmd := cmd.(type) {
		case *commands.GoodwillCommand:
			if err := cli.handleGoodwillCommand(cmd, state, player); err != nil {
				cli.logging.Error("执行好感度命令失败", zap.Error(err))
			}
		case *commands.PlaceCardCommand:
			// 创建包含当前玩家的命令上下文
			ctx := commands.CommandContext{
				GameState:     state,
				CurrentPlayer: player,
			}

			if err := cmd.Execute(ctx); err != nil {
				cli.logging.Error("执行放置命令失败", zap.Error(err))
				continue
			}

			// 成功后操作
			selectionsMade++
			// 更新可用卡牌列表
			cli.cachedCards = append(cli.cachedCards[:selectedIdx], cli.cachedCards[selectedIdx+1:]...)
			cli.logging.Info("放置卡牌成功", zap.String("卡牌ID", selectedCardID))
		} else {
			return fmt.Errorf("非法的放置命令类型 %T", cmd)
		}
	}

	// 最终验证放置数量
	switch p := player.(type) {
	case *models.Mastermind:
		if placed := len(state.Board.GetMastermindCards()); placed != maxSelections {
			return fmt.Errorf("需要精确放置3张卡牌，当前放置了%d张", placed)
		}
	case *models.Protagonist:
		if placed := len(state.Board.GetProtagonistCards(p)); placed != maxSelections {
			return fmt.Errorf("需要精确放置1张卡牌，当前放置了%d张", placed)
		}
	}

	return nil
}

func (cli *CLI) shouldAdvancePhase() bool {
	// 根据当前阶段和游戏状态判断是否应该进入下一阶段
	// 例如:所有玩家都已经放置了卡牌
	return false
}

func (cli *CLI) advanceToNextPhase() {

}
