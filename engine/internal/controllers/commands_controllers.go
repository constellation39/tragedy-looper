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

type CLI struct {
	logging     *zap.Logger
	inputReader *bufio.Reader
	cmdParser   *commands.CommandParser
}

// NewCLI 创建新的CLI控制器
func NewCLI(logger *zap.Logger) *CLI {
	return &CLI{
		logging:     logger,
		inputReader: bufio.NewReader(os.Stdin),
		cmdParser:   commands.NewCommandParser(),
	}
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
	//fmt.Printf("[%s]> ", cli.currentPhase)
	for {
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

// PlaceCommand 实现卡牌放置命令
type PlaceCommand struct {
	CardID        string
	Target        string
	gameState     *models.GameState
	currentPlayer models.Player
}

func (c *PlaceCommand) Type() commands.CommandType {
	return commands.CmdPlaceCard
}

// Execute 执行卡牌放置命令
func (c *PlaceCommand) Execute(ctx commands.CommandContext) error {
	// 实现卡牌放置逻辑
	card := findCard(c.currentPlayer, c.CardID)
	if card == nil {
		return fmt.Errorf("无效的卡牌ID")
	}

	target := findTarget(c.gameState, c.Target)
	if target == nil {
		return fmt.Errorf("无效的目标，正确格式：char_角色ID 或 loc_位置ID")
	}

	if err := c.gameState.Board.SetCard(target, card); err != nil {
		return err
	}

	return c.currentPlayer.PlaceCards(card)
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
	for {
		cli.logging.Info(fmt.Sprintf("当前手牌：%v", player.GetHandCardIDs()))
		input, _ := cli.inputReader.ReadString('\n')
		cmd, err := cli.cmdParser.Parse(strings.TrimSpace(input))
		if err != nil {
			cli.logging.Error("解析命令失败", zap.Error(err))
			continue
		}

		if placeCmd, ok := cmd.(*PlaceCommand); ok {
			placeCmd.gameState = state
			placeCmd.currentPlayer = player
			if err := placeCmd.Execute(commands.CommandContext{GameState: state}); err != nil {
				cli.logging.Error("执行放置命令失败", zap.Error(err))
			} else {
				break
			}
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
