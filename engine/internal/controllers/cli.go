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
	cli.logging.Info("- place <卡牌ID> <目标角色/位置>  放置行动卡")
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

func (cli *CLI) shouldAdvancePhase() bool {
	// 根据当前阶段和游戏状态判断是否应该进入下一阶段
	// 例如:所有玩家都已经放置了卡牌
	return false
}

func (cli *CLI) advanceToNextPhase() {

}
