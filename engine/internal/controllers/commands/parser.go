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

	switch cmdType {
	case CmdPlaceCard:
		if len(args) < 1 {
			return nil, fmt.Errorf("usage: place <cardID>")
		}
		return NewPlaceCardCommand(args[0]), nil

	case CmdSelectChar:
		if len(args) < 1 {
			return nil, fmt.Errorf("usage: selectChar <characterName>")
		}
		return NewSelectCharCommand(args[0]), nil

	case CmdSelectLocation:
		if len(args) < 1 {
			return nil, fmt.Errorf("usage: selectLoc <locationType>")
		}
		return NewSelectLocationCommand(args[0]), nil

	case CmdPassAction:
		return NewPassActionCommand(), nil

	case CmdShowCards:
		return NewShowCardsCommand(), nil

	case CmdShowBoard:
		return NewShowBoardCommand(), nil

	case CmdStatus:
		if len(args) < 1 {
			return nil, fmt.Errorf("usage: status <target>")
		}
		return NewStatusCommand(args[0]), nil

	case CmdViewRules:
		return NewViewRulesCommand(), nil

	case CmdViewIncidents:
		return NewViewIncidentsCommand(), nil

	case CmdViewHistory:
		return NewViewHistoryCommand(), nil

	case CmdUseGoodwill:
		if len(args) < 2 {
			return nil, fmt.Errorf("usage: goodwill <characterName> <abilityID>")
		}
		return NewGoodwillCommand(args[0], args[1]), nil

	case CmdFinalGuess:
		return NewFinalGuessCommand(), nil

	case CmdNextPhase:
		return NewNextPhaseCommand(), nil

	case CmdEndTurn:
		return NewEndTurnCommand(), nil

	case CmdMakeNote:
		if len(args) < 1 {
			return nil, fmt.Errorf("usage: note <content>")
		}
		return NewMakeNoteCommand(strings.Join(args, " ")), nil

	case CmdHelp:
		return NewHelpCommand(), nil

	case CmdQuit:
		return NewQuitCommand(), nil

	default:
		return nil, fmt.Errorf("unknown command: %s", cmdType)
	}
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

// ExecuteSequence 执行当前命令序列
func (p *CommandParser) ExecuteSequence(ctx CommandContext) error {
	if !p.IsSequenceComplete() {
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
