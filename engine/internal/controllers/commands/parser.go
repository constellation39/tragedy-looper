package commands

import (
	"fmt"
	"strings"
)

type CommandParser struct {
	// 已移除commandSequence字段
}

func NewCommandParser() *CommandParser {
	return &CommandParser{}
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
	case CmdSelectScript:
		if len(args) < 1 {
			return nil, fmt.Errorf("usage: selectScript <剧本名称>")
		}
		cmd = NewSelectScriptCommand(args[0])
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
		cmd = NewQuitGameCommand()

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
