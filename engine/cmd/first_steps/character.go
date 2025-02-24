package first_steps

import "tragedy-looper/engine/internal/models"

// 1. 男学生(Boy Student) - 已经正确实现，作为参考
func NewBoyStudent(role *models.Role) *models.Character {
	abilities := []*models.CharacterAbilityData{
		{
			Name:         "减少同位置学生1点不安",
			Cost:         2,
			CanBeRefused: true,
			Effect: func(state *models.GameState) error {
				return nil
			},
		},
	}

	return models.NewCharacter(&models.CharacterData{
		Name:                "Boy Student",
		StartLocation:       models.LocationSchool,
		GoodwillLimit:       4,
		ParanoiaLimit:       3,
		GoodwillAbilityList: abilities,
	}, role)
}

// 2. 女学生(Girl Student)
func NewGirlStudent(role *models.Role) *models.Character {
	abilities := []*models.CharacterAbilityData{
		{
			Name:         "减少同位置学生1点不安",
			Cost:         2,
			CanBeRefused: true,
			Effect: func(state *models.GameState) error {
				return nil
			},
		},
	}

	return models.NewCharacter(&models.CharacterData{
		Name:                "Girl Student",
		StartLocation:       models.LocationSchool,
		GoodwillLimit:       4,
		ParanoiaLimit:       3,
		GoodwillAbilityList: abilities,
	}, role)
}

// 3. 大小姐(Rich Man's Daughter)
func NewRichMansDaughter(role *models.Role) *models.Character {
	abilities := []*models.CharacterAbilityData{
		{
			Name:         "增加同位置角色1点好感度",
			Cost:         3,
			CanBeRefused: true,
			Effect: func(state *models.GameState) error {
				return nil
			},
		},
	}

	return models.NewCharacter(&models.CharacterData{
		Name:                "Rich Man's Daughter",
		StartLocation:       models.LocationCity,
		GoodwillLimit:       4,
		ParanoiaLimit:       3,
		GoodwillAbilityList: abilities,
	}, role)
}

// 4. 班长(Class Rep)
func NewClassRep(role *models.Role) *models.Character {
	abilities := []*models.CharacterAbilityData{
		{
			Name:         "让领导者取回一张一次性卡牌",
			Cost:         2,
			CanBeRefused: true,
			Effect: func(state *models.GameState) error {
				return nil
			},
		},
	}

	return models.NewCharacter(&models.CharacterData{
		Name:                "Class Rep",
		StartLocation:       models.LocationSchool,
		GoodwillLimit:       4,
		ParanoiaLimit:       3,
		GoodwillAbilityList: abilities,
	}, role)
}

// 5. 神秘少年(Mystery Boy)
func NewMysteryBoy(role *models.Role) *models.Character {
	abilities := []*models.CharacterAbilityData{
		{
			Name:         "展示自己的角色身份",
			Cost:         3,
			CanBeRefused: false,
			Effect: func(state *models.GameState) error {
				return nil
			},
		},
	}

	return models.NewCharacter(&models.CharacterData{
		Name:                "Mystery Boy",
		StartLocation:       models.LocationCity,
		GoodwillLimit:       4,
		ParanoiaLimit:       3,
		GoodwillAbilityList: abilities,
	}, role)
}

// 6. 巫女(Shrine Maiden)
func NewShrineMaiden(role *models.Role) *models.Character {
	abilities := []*models.CharacterAbilityData{
		{
			Name:         "减少神社1点阴谋",
			Cost:         3,
			CanBeRefused: true,
			Effect: func(state *models.GameState) error {
				return nil
			},
		},
		{
			Name:         "展示同位置角色的身份",
			Cost:         5,
			CanBeRefused: true,
			Effect: func(state *models.GameState) error {
				return nil
			},
		},
	}

	return models.NewCharacter(&models.CharacterData{
		Name:                "Shrine Maiden",
		StartLocation:       models.LocationShrine,
		GoodwillLimit:       6,
		ParanoiaLimit:       3,
		GoodwillAbilityList: abilities,
	}, role)
}

// 7. 外星人(Alien)
func NewAlien(role *models.Role) *models.Character {
	abilities := []*models.CharacterAbilityData{
		{
			Name:         "杀死同位置的一个角色",
			Cost:         4,
			CanBeRefused: true,
			Effect: func(state *models.GameState) error {
				return nil
			},
		},
		{
			Name:         "复活同位置的一个尸体",
			Cost:         5,
			CanBeRefused: true,
			Effect: func(state *models.GameState) error {
				return nil
			},
		},
	}

	return models.NewCharacter(&models.CharacterData{
		Name:                "Alien",
		StartLocation:       models.LocationCity,
		GoodwillLimit:       6,
		ParanoiaLimit:       3,
		GoodwillAbilityList: abilities,
	}, role)
}

// 8. 神明(Godly)
func NewGodly(role *models.Role) *models.Character {
	abilities := []*models.CharacterAbilityData{
		{
			Name:         "揭示一个事件的凶手",
			Cost:         3,
			CanBeRefused: true,
			Effect: func(state *models.GameState) error {
				return nil
			},
		},
		{
			Name:         "减少同位置或角色1点阴谋",
			Cost:         5,
			CanBeRefused: true,
			Effect: func(state *models.GameState) error {
				return nil
			},
		},
	}

	return models.NewCharacter(&models.CharacterData{
		Name:                "Godly",
		StartLocation:       models.LocationShrine,
		GoodwillLimit:       6,
		ParanoiaLimit:       3,
		GoodwillAbilityList: abilities,
	}, role)
}

// 9. 警察(Police Officer)
func NewPoliceOfficer(role *models.Role) *models.Character {
	abilities := []*models.CharacterAbilityData{
		{
			Name:         "揭示前一个事件的凶手",
			Cost:         4,
			CanBeRefused: true,
			Effect: func(state *models.GameState) error {
				return nil
			},
		},
		{
			Name:         "阻止此处一次死亡",
			Cost:         5,
			CanBeRefused: true,
			Effect: func(state *models.GameState) error {
				return nil
			},
		},
	}

	return models.NewCharacter(&models.CharacterData{
		Name:                "Police Officer",
		StartLocation:       models.LocationCity,
		GoodwillLimit:       6,
		ParanoiaLimit:       3,
		GoodwillAbilityList: abilities,
	}, role)
}

// 10. 上班族(Office Worker)
func NewOfficeWorker(role *models.Role) *models.Character {
	abilities := []*models.CharacterAbilityData{
		{
			Name:         "展示自己的角色身份",
			Cost:         3,
			CanBeRefused: true,
			Effect: func(state *models.GameState) error {
				return nil
			},
		},
	}

	return models.NewCharacter(&models.CharacterData{
		Name:                "Office Worker",
		StartLocation:       models.LocationCity,
		GoodwillLimit:       4,
		ParanoiaLimit:       3,
		GoodwillAbilityList: abilities,
	}, role)
}

// 11. 告密者(Informer)
func NewInformer(role *models.Role) *models.Character {
	abilities := []*models.CharacterAbilityData{
		{
			Name:         "揭示副本A或B",
			Cost:         5,
			CanBeRefused: true,
			Effect: func(state *models.GameState) error {
				return nil
			},
		},
	}

	return models.NewCharacter(&models.CharacterData{
		Name:                "Informer",
		StartLocation:       models.LocationCity,
		GoodwillLimit:       6,
		ParanoiaLimit:       3,
		GoodwillAbilityList: abilities,
	}, role)
}

// 12. 偶像(Pop Idol)
func NewPopIdol(role *models.Role) *models.Character {
	abilities := []*models.CharacterAbilityData{
		{
			Name:         "减少同位置角色1点不安",
			Cost:         3,
			CanBeRefused: true,
			Effect: func(state *models.GameState) error {
				return nil
			},
		},
		{
			Name:         "增加同位置角色1点好感度",
			Cost:         4,
			CanBeRefused: true,
			Effect: func(state *models.GameState) error {
				return nil
			},
		},
	}

	return models.NewCharacter(&models.CharacterData{
		Name:                "Pop Idol",
		StartLocation:       models.LocationCity,
		GoodwillLimit:       5,
		ParanoiaLimit:       3,
		GoodwillAbilityList: abilities,
	}, role)
}

// 13. 记者(Journalist)
func NewJournalist(role *models.Role) *models.Character {
	abilities := []*models.CharacterAbilityData{
		{
			Name:         "减少任意角色1点不安",
			Cost:         2,
			CanBeRefused: true,
			Effect: func(state *models.GameState) error {
				return nil
			},
		},
		{
			Name:         "增加同位置或角色1点阴谋",
			Cost:         2,
			CanBeRefused: true,
			Effect: func(state *models.GameState) error {
				return nil
			},
		},
	}

	return models.NewCharacter(&models.CharacterData{
		Name:                "Journalist",
		StartLocation:       models.LocationCity,
		GoodwillLimit:       4,
		ParanoiaLimit:       3,
		GoodwillAbilityList: abilities,
	}, role)
}

// 14. 老板(Boss)
func NewBoss(role *models.Role) *models.Character {
	abilities := []*models.CharacterAbilityData{
		{
			Name:         "揭示势力范围内角色的身份",
			Cost:         5,
			CanBeRefused: true,
			Effect: func(state *models.GameState) error {
				return nil
			},
		},
	}

	return models.NewCharacter(&models.CharacterData{
		Name:                "Boss",
		StartLocation:       models.LocationCity,
		GoodwillLimit:       6,
		ParanoiaLimit:       3,
		GoodwillAbilityList: abilities,
	}, role)
}

// 15. 医生(Doctor)
func NewDoctor(role *models.Role) *models.Character {
	abilities := []*models.CharacterAbilityData{
		{
			Name:         "增减同位置角色1点不安",
			Cost:         2,
			CanBeRefused: true,
			Effect: func(state *models.GameState) error {
				return nil
			},
		},
		{
			Name:         "解除病人的位置限制",
			Cost:         3,
			CanBeRefused: true,
			Effect: func(state *models.GameState) error {
				return nil
			},
		},
	}

	return models.NewCharacter(&models.CharacterData{
		Name:                "Doctor",
		StartLocation:       models.LocationHospital,
		GoodwillLimit:       4,
		ParanoiaLimit:       3,
		GoodwillAbilityList: abilities,
	}, role)
}

// 16. 病人(Patient)
func NewPatient(role *models.Role) *models.Character {
	return models.NewCharacter(&models.CharacterData{
		Name:                "Patient",
		StartLocation:       models.LocationHospital,
		GoodwillLimit:       4,
		ParanoiaLimit:       2,                                // 特殊的不安上限
		GoodwillAbilityList: []*models.CharacterAbilityData{}, // 病人没有特殊能力
	}, role)
}

// 17. 护士(Nurse)
func NewNurse(role *models.Role) *models.Character {
	abilities := []*models.CharacterAbilityData{
		{
			Name:         "减少同位置恐慌角色1点不安",
			Cost:         2,
			CanBeRefused: false,
			Effect: func(state *models.GameState) error {
				return nil
			},
		},
	}

	return models.NewCharacter(&models.CharacterData{
		Name:                "Nurse",
		StartLocation:       models.LocationHospital,
		GoodwillLimit:       4,
		ParanoiaLimit:       3,
		GoodwillAbilityList: abilities,
	}, role)
}

// 18. 手下(Henchman)
func NewHenchman(role *models.Role) *models.Character {
	abilities := []*models.CharacterAbilityData{
		{
			Name:         "不触发事件",
			Cost:         3,
			CanBeRefused: true,
			Effect: func(state *models.GameState) error {
				return nil
			},
		},
	}

	return models.NewCharacter(&models.CharacterData{
		Name:                "Henchman",
		StartLocation:       models.LocationCity,
		GoodwillLimit:       4,
		ParanoiaLimit:       3,
		GoodwillAbilityList: abilities,
	}, role)
}

// 19. 局外人(Outsider)
func NewOutsider(role *models.Role) *models.Character {
	abilities := []*models.CharacterAbilityData{
		{
			Name:         "展示自己的身份",
			Cost:         3,
			CanBeRefused: true,
			Effect: func(state *models.GameState) error {
				return nil
			},
		},
	}

	return models.NewCharacter(&models.CharacterData{
		Name:                "Outsider",
		StartLocation:       models.LocationCity,
		GoodwillLimit:       4,
		ParanoiaLimit:       3,
		GoodwillAbilityList: abilities,
	}, role)
}
