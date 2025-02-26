package first_steps

import (
	"fmt"
	"tragedy-looper/engine/internal/models"
)

// 1. 男学生(Boy Student) - 已经正确实现，作为参考
func NewBoyStudent(role *models.Role) *models.Character {
	character := models.NewCharacter(&models.CharacterData{
		Name:          "BoyStudent",
		StartLocation: models.LocationSchool,
		GoodwillLimit: 4,
		ParanoiaLimit: 3,
	}, role)

	character.GoodwillAbilityList = []*models.CharacterAbilityData{
		{
			Name:         "减少同位置学生1点不安",
			Cost:         2,
			CanBeRefused: true,
			Effect: func(state *models.GameState) error {
				location := state.Location(character.CurrentLocation)
				for _, char := range location.Characters {
					if !char.ExistsTag("Student") {
						continue
					}
					newParanoia := char.Paranoia() - 1
					if newParanoia < 0 {
						newParanoia = 0
					}
					char.SetParanoia(newParanoia)
				}
				return nil
			},
		},
	}

	return character
}

// 2. 女学生(Girl Student)
func NewGirlStudent(role *models.Role) *models.Character {
	character := models.NewCharacter(&models.CharacterData{
		Name:          "GirlStudent",
		StartLocation: models.LocationSchool,
		GoodwillLimit: 4,
		ParanoiaLimit: 3,
	}, role)
	
	character.GoodwillAbilityList = []*models.CharacterAbilityData{
		{
			Name:         "减少同位置学生1点不安",
			Cost:         2,
			CanBeRefused: true,
			Effect: func(state *models.GameState) error {
				location := state.Location(character.CurrentLocation)
				for _, char := range location.Characters {
					if !char.ExistsTag("Student") {
						continue
					}
					newParanoia := char.Paranoia() - 1
					if newParanoia < 0 {
						newParanoia = 0
					}
					char.SetParanoia(newParanoia)
				}
				return nil
			},
		},
	}

	return character
}

// 3. 大小姐(Rich Man's Daughter)
func NewRichMansDaughter(role *models.Role) *models.Character {
	character := models.NewCharacter(&models.CharacterData{
		Name:          "RichMansDaughter",
		StartLocation: models.LocationCity,
		GoodwillLimit: 4,
		ParanoiaLimit: 3,
	}, role)
	
	character.GoodwillAbilityList = []*models.CharacterAbilityData{
		{
			Name:         "增加同位置角色1点好感度",
			Cost:         3,
			CanBeRefused: true,
			Effect: func(state *models.GameState) error {
				// 需要玩家选择同位置的一个角色
				if ctx, ok := state.AbilityContext.(*models.AbilityContext); ok && ctx.Target != nil {
					if targetChar, ok := ctx.Target.(*models.Character); ok {
						// 验证目标角色是否在同一位置
						if targetChar.CurrentLocation == character.CurrentLocation {
							newGoodwill := targetChar.Goodwill() + 1
							if newGoodwill > targetChar.GoodwillLimit {
								newGoodwill = targetChar.GoodwillLimit
							}
							targetChar.SetGoodwill(newGoodwill)
						}
					}
				}
				return nil
			},
		},
	}

	return character
}

// 4. 班长(Class Rep)
func NewClassRep(role *models.Role) *models.Character {
	character := models.NewCharacter(&models.CharacterData{
		Name:          "ClassRep",
		StartLocation: models.LocationSchool,
		GoodwillLimit: 4,
		ParanoiaLimit: 3,
	}, role)
	
	character.GoodwillAbilityList = []*models.CharacterAbilityData{
		{
			Name:         "让领导者取回一张一次性卡牌",
			Cost:         2,
			CanBeRefused: true,
			Effect: func(state *models.GameState) error {
				// 找到领导者原型
				var leader *models.Protagonist
				for _, p := range state.Protagonists {
					if p.IsLeader {
						leader = p
						break
					}
				}
				
				if leader == nil || len(leader.OnceCards) == 0 {
					return nil // 没有领导者或没有一次性卡牌可取回
				}
				
				// 需要玩家选择一张领导者的一次性卡牌取回
				if ctx, ok := state.AbilityContext.(*models.AbilityContext); ok && ctx.Target != nil {
					if cardID, ok := ctx.Target.(string); ok {
						// 在一次性卡牌中查找该卡牌
						for i, card := range leader.OnceCards {
							if card.Id() == cardID {
								// 将卡牌从一次性列表移回手牌
								leader.HandCards = append(leader.HandCards, card)
								// 从一次性列表中移除
								leader.OnceCards = append(leader.OnceCards[:i], leader.OnceCards[i+1:]...)
								break
							}
						}
					}
				}
				return nil
			},
		},
	}

	return character
}

// 5. 神秘少年(Mystery Boy)
func NewMysteryBoy(role *models.Role) *models.Character {
	character := models.NewCharacter(&models.CharacterData{
		Name:          "MysteryBoy",
		StartLocation: models.LocationCity,
		GoodwillLimit: 4,
		ParanoiaLimit: 3,
	}, role)
	
	character.GoodwillAbilityList = []*models.CharacterAbilityData{
		{
			Name:         "展示自己的角色身份",
			Cost:         3,
			CanBeRefused: false,
			Effect: func(state *models.GameState) error {
				// 直接展示自己的角色身份
				if character.Role() != nil {
					state.LogEvent(fmt.Sprintf("神秘少年展示了身份：%s", character.Role().Type))
				} else {
					state.LogEvent("神秘少年没有特殊身份")
				}
				return nil
			},
		},
	}

	return character
}

// 6. 巫女(Shrine Maiden)
func NewShrineMaiden(role *models.Role) *models.Character {
	character := models.NewCharacter(&models.CharacterData{
		Name:          "ShrineMaiden",
		StartLocation: models.LocationShrine,
		GoodwillLimit: 6,
		ParanoiaLimit: 3,
	}, role)
	
	character.GoodwillAbilityList = []*models.CharacterAbilityData{
		{
			Name:         "减少神社1点阴谋",
			Cost:         3,
			CanBeRefused: true,
			Effect: func(state *models.GameState) error {
				shrine := state.Location(models.LocationShrine)
				if shrine != nil {
					newIntrigue := shrine.Intrigue() - 1
					if newIntrigue < 0 {
						newIntrigue = 0
					}
					shrine.SetIntrigue(newIntrigue)
				}
				return nil
			},
		},
		{
			Name:         "展示同位置角色的身份",
			Cost:         5,
			CanBeRefused: true,
			Effect: func(state *models.GameState) error {
				if ctx, ok := state.AbilityContext.(*models.AbilityContext); ok && ctx.Target != nil {
					if targetChar, ok := ctx.Target.(*models.Character); ok {
						// 验证目标角色是否在同一位置
						if targetChar.CurrentLocation == character.CurrentLocation {
							if targetChar.Role() != nil {
								state.LogEvent(fmt.Sprintf("巫女展示了%s的身份：%s", targetChar.Name, targetChar.Role().Type))
							} else {
								state.LogEvent(fmt.Sprintf("%s没有特殊身份", targetChar.Name))
							}
						}
					}
				}
				return nil
			},
		},
	}

	return character
}

// 7. 外星人(Alien)
func NewAlien(role *models.Role) *models.Character {
	character := models.NewCharacter(&models.CharacterData{
		Name:          "Alien",
		StartLocation: models.LocationCity,
		GoodwillLimit: 6,
		ParanoiaLimit: 3,
	}, role)
	
	character.GoodwillAbilityList = []*models.CharacterAbilityData{
		{
			Name:         "杀死同位置的一个角色",
			Cost:         4,
			CanBeRefused: true,
			Effect: func(state *models.GameState) error {
				if ctx, ok := state.AbilityContext.(*models.AbilityContext); ok && ctx.Target != nil {
					if targetChar, ok := ctx.Target.(*models.Character); ok {
						// 验证目标角色是否在同一位置
						if targetChar.CurrentLocation == character.CurrentLocation && targetChar.IsAlive() {
							targetChar.Kill()
							state.LogEvent(fmt.Sprintf("外星人杀死了%s", targetChar.Name))
						}
					}
				}
				return nil
			},
		},
		{
			Name:         "复活同位置的一个尸体",
			Cost:         5,
			CanBeRefused: true,
			Effect: func(state *models.GameState) error {
				if ctx, ok := state.AbilityContext.(*models.AbilityContext); ok && ctx.Target != nil {
					if targetChar, ok := ctx.Target.(*models.Character); ok {
						// 验证目标角色是否在同一位置且已死亡
						if targetChar.CurrentLocation == character.CurrentLocation && !targetChar.IsAlive() {
							targetChar.Revive()
							state.LogEvent(fmt.Sprintf("外星人复活了%s", targetChar.Name))
						}
					}
				}
				return nil
			},
		},
	}

	return character
}

// 8. 神明(Godly)
func NewGodly(role *models.Role) *models.Character {
	character := models.NewCharacter(&models.CharacterData{
		Name:          "Godly",
		StartLocation: models.LocationShrine,
		GoodwillLimit: 6,
		ParanoiaLimit: 3,
	}, role)
	
	character.GoodwillAbilityList = []*models.CharacterAbilityData{
		{
			Name:         "揭示一个事件的凶手",
			Cost:         3,
			CanBeRefused: true,
			Effect: func(state *models.GameState) error {
				if ctx, ok := state.AbilityContext.(*models.AbilityContext); ok && ctx.Target != nil {
					if incidentID, ok := ctx.Target.(string); ok {
						// 查找该事件
						for _, incident := range state.Incidents {
							if incident.ID == incidentID && incident.Culprit != nil {
								state.LogEvent(fmt.Sprintf("神明揭示事件 %s 的凶手是 %s", incident.Name, incident.Culprit.Name))
								break
							}
						}
					}
				}
				return nil
			},
		},
		{
			Name:         "减少同位置或角色1点阴谋",
			Cost:         5,
			CanBeRefused: true,
			Effect: func(state *models.GameState) error {
				if ctx, ok := state.AbilityContext.(*models.AbilityContext); ok && ctx.Target != nil {
					// 可以针对位置或角色
					switch target := ctx.Target.(type) {
					case *models.Location:
						newIntrigue := target.Intrigue() - 1
						if newIntrigue < 0 {
							newIntrigue = 0
						}
						target.SetIntrigue(newIntrigue)
						state.LogEvent(fmt.Sprintf("神明减少了%s的1点阴谋", target.LocationType))
					case *models.Character:
						newIntrigue := target.