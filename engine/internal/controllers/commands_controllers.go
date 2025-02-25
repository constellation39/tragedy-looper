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

// UI 接口用于多种交互模式
type UI interface {
	Select(title string, options []string) (int, error)
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

// 格式化目标参数
func (cli *CLI) formatTarget(target models.TargetType) string {
	switch t := target.(type) {
	case *models.Character:
		return fmt.Sprintf("char_%s", t.Name)
	case *models.Location:
		return fmt.Sprintf("loc_%s", t)
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
		if placeCmd, ok := cmd.(*commands.PlaceCardCommand); ok {
			// 创建包含当前玩家的命令上下文
			ctx := commands.CommandContext{
				GameState:     state,
				CurrentPlayer: player,
			}
			
			if err := placeCmd.Execute(ctx); err != nil {
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
