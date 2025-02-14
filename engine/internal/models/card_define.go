package models

func InitMastermindCard(mastermind *Mastermind) (mastermindCard []Card) {
	// 初始化移动卡
	mastermindCard = append(mastermindCard, NewMovementCard(mastermind, HorizontalMovement, false))
	mastermindCard = append(mastermindCard, NewMovementCard(mastermind, VerticalMovement, false))
	mastermindCard = append(mastermindCard, NewMovementCard(mastermind, DiagonalMovement, true))

	// 初始化不安卡
	mastermindCard = append(mastermindCard, NewParanoiaCard(mastermind, +1, false))
	mastermindCard = append(mastermindCard, NewParanoiaCard(mastermind, +1, false))
	mastermindCard = append(mastermindCard, NewParanoiaCard(mastermind, -1, false))

	// 初始化阴谋卡
	mastermindCard = append(mastermindCard, NewIntrigueCard(mastermind, +1, false))
	mastermindCard = append(mastermindCard, NewIntrigueCard(mastermind, +2, true))

	//禁止友好
	mastermindCard = append(mastermindCard, NewForbidGoodwillCard(mastermind, false))
	//禁止不安
	mastermindCard = append(mastermindCard, NewForbidParanoiaCard(mastermind, false))
	return
}

func InitProtagonistCard(protagonist *Protagonist) (protagonistCard []Card) {
	// 初始化移动卡
	protagonistCard = append(protagonistCard, NewMovementCard(protagonist, HorizontalMovement, false))
	protagonistCard = append(protagonistCard, NewMovementCard(protagonist, VerticalMovement, false))

	// 初始化不安卡
	protagonistCard = append(protagonistCard, NewParanoiaCard(protagonist, +1, false))
	protagonistCard = append(protagonistCard, NewParanoiaCard(protagonist, -1, false))

	// 初始化友好卡
	protagonistCard = append(protagonistCard, NewGoodwillCard(protagonist, +1, false))
	protagonistCard = append(protagonistCard, NewGoodwillCard(protagonist, +2, true))

	// 禁止移动
	protagonistCard = append(protagonistCard, NewForbidMovementCard(protagonist, false))
	return
}
