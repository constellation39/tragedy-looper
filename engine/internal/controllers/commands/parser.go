package commands

import (
	"fmt"
	"strings"
)

type CommandParser struct {
	// 存储命令序列，用于组合命令
	commandSequence []Command
}

func NewCommandParser() *CommandParser {
	return &CommandParser{
		commandSequence: make([]Command, 0),
	}
}

// Parse 解析命令字符串为命令对象
func (p *CommandParser) Parse(input string) (Command, error) {
	parts := strings.Fields(input)
	if len(parts) == 0 {
		return nil, fmt.Errorf("empty command")
	}

	cmdType := CommandType(parts[0])
	args := parts[1:]

	var cmd Command
	var err error

	switch cmdType {
	case CmdStartGame:
		cmd = NewStartGameCommand()
	case CmdPlaceCard:
		if len(args) < 1 {
			return nil, fmt.Errorf("usage: place <cardID>")
		}
		cmd = NewPlaceCardCommand(args[0])

	case CmdSelectChar:
		if len(args) < 1 {
			return nil, fmt.Errorf("usage: selectChar <characterName>")
		}
		cmd = NewSelectCharCommand(args[0])

	case CmdSelectLocation:
		if len(args) < 1 {
			return nil, fmt.Errorf("usage: selectLoc <locationType>")
		}
		cmd = NewSelectLocationCommand(args[0])

	case CmdPassAction:
		cmd = NewPassActionCommand()

	case CmdShowCards:
		cmd = NewShowCardsCommand()

	case CmdShowBoard:
		cmd = NewShowBoardCommand()

	case CmdStatus:
		if len(args) < 1 {
			return nil, fmt.Errorf("usage: status <target>")
		}
		cmd = NewStatusCommand(args[0])

	case CmdViewRules:
		cmd = NewViewRulesCommand()

	case CmdViewIncidents:
		cmd = NewViewIncidentsCommand()

	case CmdViewHistory:
		cmd = NewViewHistoryCommand()

	case CmdUseGoodwill:
		if len(args) < 2 {
			return nil, fmt.Errorf("usage: goodwill <characterName> <abilityID>")
		}
		cmd = NewGoodwillCommand(args[0], args[1])

	case CmdFinalGuess:
		cmd = NewFinalGuessCommand()

	case CmdNextPhase:
		cmd = NewNextPhaseCommand()

	case CmdEndTurn:
		cmd = NewEndTurnCommand()

	case CmdMakeNote:
		if len(args) < 1 {
			return nil, fmt.Errorf("usage: note <content>")
		}
		cmd = NewMakeNoteCommand(strings.Join(args, " "))

	case CmdHelp:
		cmd = NewHelpCommand()

	case CmdQuit:
		cmd = NewQuitCommand()

	default:
		return nil, fmt.Errorf("unknown command: %s", cmdType)
	}

	// 验证命令参数
	if err = cmd.Validate(); err != nil {
		return nil, fmt.Errorf("命令参数无效: %v", err)
	}

	return cmd, nil
}

// GetRequiredInputs 获取命令所需的输入描述
func (p *CommandParser) GetRequiredInputs(cmd Command) []string {
	return cmd.RequiredInputs()
}

// AddToSequence 将命令添加到序列中
func (p *CommandParser) AddToSequence(cmd Command) {
	p.commandSequence = append(p.commandSequence, cmd)
}

// GetCommandSequence 获取当前命令序列
func (p *CommandParser) GetCommandSequence() []Command {
	return p.commandSequence
}

// ClearSequence 清空命令序列
func (p *CommandParser) ClearSequence() {
	p.commandSequence = make([]Command, 0)
}

// IsSequenceComplete 检查命令序列是否完整可执行
func (p *CommandParser) IsSequenceComplete() bool {
	if len(p.commandSequence) == 0 {
		return false
	}

	// 检查是否是有效的命令组合
	firstCmd := p.commandSequence[0]
	switch firstCmd.Type() {
	case CmdPlaceCard:
		// place命令需要跟随一个选择目标的命令
		if len(p.commandSequence) < 2 {
			return false
		}
		secondCmd := p.commandSequence[1]
		return secondCmd.Type() == CmdSelectChar || secondCmd.Type() == CmdSelectLocation
	case CmdUseGoodwill:
		// goodwill命令需要跟随一个选择目标的命令
		if len(p.commandSequence) < 2 {
			return false
		}
		secondCmd := p.commandSequence[1]
		return secondCmd.Type() == CmdSelectChar || secondCmd.Type() == CmdSelectLocation
	default:
		// 其他命令可以单独执行
		return true
	}
}

// GetMissingInputs 获取命令序列中缺少的输入
func (p *CommandParser) GetMissingInputs() []string {
	if len(p.commandSequence) == 0 {
		return []string{"请输入一个命令"}
	}

	firstCmd := p.commandSequence[0]
	switch firstCmd.Type() {
	case CmdPlaceCard:
		if len(p.commandSequence) < 2 {
			return []string{"需要使用selectChar或selectLoc命令选择目标"}
		}
	case CmdUseGoodwill:
		if len(p.commandSequence) < 2 {
			return []string{"需要使用selectChar或selectLoc命令选择目标"}
		}
	}

	return nil
}

// ExecuteSequence 执行当前命令序列
func (p *CommandParser) ExecuteSequence(ctx CommandContext) error {
	if !p.IsSequenceComplete() {
		missingInputs := p.GetMissingInputs()
		if len(missingInputs) > 0 {
			return fmt.Errorf("命令序列不完整: %s", strings.Join(missingInputs, ", "))
		}
		return fmt.Errorf("命令序列不完整")
	}

	// 根据命令类型执行不同的组合逻辑
	firstCmd := p.commandSequence[0]
	switch firstCmd.Type() {
	case CmdPlaceCard:
		placeCmd := firstCmd.(*PlaceCardCommand)
		targetCmd := p.commandSequence[1]

		var target string
		if targetCmd.Type() == CmdSelectChar {
			target = "char_" + targetCmd.(*SelectCharCommand).CharacterName
		} else {
			target = "loc_" + targetCmd.(*SelectLocationCommand).LocationType
		}

		// 设置目标并执行
		placeCmd.SetTarget(target)
		return placeCmd.Execute(ctx)

	case CmdUseGoodwill:
		goodwillCmd := firstCmd.(*GoodwillCommand)
		targetCmd := p.commandSequence[1]

		var target string
		if targetCmd.Type() == CmdSelectChar {
			target = "char_" + targetCmd.(*SelectCharCommand).CharacterName
		} else {
			target = "loc_" + targetCmd.(*SelectLocationCommand).LocationType
		}

		// 设置目标并执行
		goodwillCmd.Target = target
		return goodwillCmd.Execute(ctx)

	default:
		// 单命令直接执行
		return firstCmd.Execute(ctx)
	}
}
