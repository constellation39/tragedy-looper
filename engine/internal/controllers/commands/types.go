package commands

type CommandType string

const (
	// CmdSelectScript - Select a script to play
	// Syntax: selectScript <ScriptID>
	// Example: selectScript "first_steps_1"
	CmdSelectScript CommandType = "selectScript"

	// CmdStartGame - Start the game/loop with current setup
	// Syntax: start
	CmdStartGame CommandType = "start"

	// CmdPlaceCard - Place an action card on a character or location
	// Syntax: place <CardID> <LocationType|CharacterName>
	// Example: place "goodwill+1" "Student"
	// Example: place "intrigue+1" "Hospital"
	CmdPlaceCard CommandType = "place"

	// CmdPassAction - Skip current action/do nothing this turn
	// Syntax: pass
	CmdPassAction CommandType = "pass"

	// CmdShowCards - View your hand of action cards
	// Syntax: cards
	CmdShowCards CommandType = "cards"

	// CmdShowBoard - View the current board state with all characters and locations
	// Syntax: board
	CmdShowBoard CommandType = "board"

	// CmdStatus - Check status of a character or location
	// Syntax: status <CharacterName|LocationType>
	// Example: status "Student"
	// Example: status "Hospital"
	CmdStatus CommandType = "status"

	// CmdViewRules - View current script rules or role information
	// Syntax: rules [RoleName|PlotName]
	// Example: rules "Key Person"
	// Example: rules "Murder Plan"
	CmdViewRules CommandType = "rules"

	// CmdViewIncidents - View incidents that have occurred
	// Syntax: incidents [Day]
	// Example: incidents 2
	CmdViewIncidents CommandType = "incidents"

	// CmdViewHistory - View history log of events
	// Syntax: history [Day]
	// Example: history 1
	CmdViewHistory CommandType = "history"

	// CmdUseGoodwill - Use a character's goodwill ability
	// Syntax: goodwill <CharacterName> [TargetName]
	// Example: goodwill "Shrine Maiden" "Student"
	CmdUseGoodwill CommandType = "goodwill"

	// CmdSelectChar - Select a character as target for an action or ability
	// Syntax: selectChar <CharacterName>
	// Example: selectChar "Office Worker"
	CmdSelectChar CommandType = "selectChar"

	// CmdSelectLocation - Select a location as target
	// Syntax: selectLoc <LocationType>
	// Example: selectLoc "Shrine"
	CmdSelectLocation CommandType = "selectLoc"

	// CmdFinalGuess - Make final guess about character roles (Protagonist only)
	// Syntax: guess <CharacterName> <RoleName> [CharacterName RoleName...]
	// Example: guess "Student" "Key Person" "Office Worker" "Killer"
	CmdFinalGuess CommandType = "guess"

	// CmdNextPhase - Proceed to next game phase
	// Syntax: next
	CmdNextPhase CommandType = "next"

	// CmdEndTurn - End current turn
	// Syntax: end
	CmdEndTurn CommandType = "end"

	// CmdMakeNote - Add personal note/marking
	// Syntax: note <Text>
	// Example: note "I suspect the Student is the Key Person"
	CmdMakeNote CommandType = "note"

	// CmdHelp - Display help information
	// Syntax: help [CommandName]
	// Example: help place
	CmdHelp CommandType = "help"

	// CmdQuit - Exit the game
	// Syntax: quit
	CmdQuit CommandType = "quit"

	// CmdMoveCharacter - Move a character to a different location
	// Syntax: move <CharacterName> <LocationType>
	// Example: move "Student" "Hospital"
	CmdMoveCharacter CommandType = "move"

	// Additional commands based on game mechanics from your knowledge base

	// CmdCheckParanoia - Check characters at or above paranoia limit
	// Syntax: paranoia
	CmdCheckParanoia CommandType = "paranoia"

	// CmdCheckIntrigue - Check intrigue counters on characters/locations
	// Syntax: intrigue
	CmdCheckIntrigue CommandType = "intrigue"

	// CmdTriggerIncident - Trigger a scheduled incident (Mastermind only)
	// Syntax: triggerIncident <IncidentName> [TargetName]
	// Example: triggerIncident "Murder" "Student"
	CmdTriggerIncident CommandType = "triggerIncident"

	// CmdUseMastermindAbility - Use a character's mastermind ability (Mastermind only)
	// Syntax: mastermindAbility <CharacterName> [TargetName|LocationName]
	// Example: mastermindAbility "Brain" "Student"
	CmdUseMastermindAbility CommandType = "mastermindAbility"

	// CmdDayStart - Perform day start actions (Mastermind only)
	// Syntax: dayStart
	CmdDayStart CommandType = "dayStart"

	// CmdDayEnd - Perform day end actions (Mastermind only)
	// Syntax: dayEnd
	CmdDayEnd CommandType = "dayEnd"

	// CmdLoopStart - Perform loop start actions (Mastermind only)
	// Syntax: loopStart
	CmdLoopStart CommandType = "loopStart"

	// CmdLoopEnd - Perform loop end actions and check loss conditions (Mastermind only)
	// Syntax: loopEnd
	CmdLoopEnd CommandType = "loopEnd"

	// CmdRevealRole - Reveal a character's role (Mastermind only)
	// Syntax: revealRole <CharacterName>
	// Example: revealRole "Student"
	CmdRevealRole CommandType = "revealRole"

	// CmdKillCharacter - Kill a character (Mastermind only)
	// Syntax: kill <CharacterName>
	// Example: kill "Student"
	CmdKillCharacter CommandType = "kill"

	// CmdKillProtagonists - Kill the protagonists (Mastermind only)
	// Syntax: killProtagonists
	CmdKillProtagonists CommandType = "killProtagonists"

	// CmdCheckLossCondition - Check if a specific loss condition is met (Mastermind only)
	// Syntax: checkLoss <PlotName>
	// Example: checkLoss "Murder Plan"
	CmdCheckLossCondition CommandType = "checkLoss"

	// CmdSetupTimeSpiral - Set up a time spiral discussion between loops (Protagonists only)
	// Syntax: timeSpiral
	CmdSetupTimeSpiral CommandType = "timeSpiral"

	// CmdChangeLeader - Change the current leader among protagonists
	// Syntax: changeLeader <PlayerName>
	// Example: changeLeader "Player2"
	CmdChangeLeader CommandType = "changeLeader"
)
