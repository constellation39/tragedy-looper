package commands

import (
	"fmt"
	"strings"
)

type CommandParser struct {
}

func NewCommandParser() *CommandParser {
	return &CommandParser{}
}

func (p *CommandParser) Parse(input string) (Command, error) {
	parts := strings.Fields(input)
	if len(parts) == 0 {
		return nil, fmt.Errorf("empty command")
	}

	cmdType := CommandType(parts[0])
	args := parts[1:]

	switch cmdType {
	case CmdPlaceCard:
		if len(args) != 2 {
			return nil, fmt.Errorf("usage: place <cardID> <target>")
		}
		return NewPlaceCardCommand(args[0], args[1]), nil

	case CmdUseGoodwill:
		if len(args) != 3 {
			return nil, fmt.Errorf("usage: goodwill <character> <abilityID> <target>")
		}
		return &GoodwillCommand{
			CharacterName: args[0],
			AbilityID:     args[1],
			Target:        args[2],
		}, nil

	default:
		return nil, fmt.Errorf("unknown command: %s", cmdType)
	}
}
