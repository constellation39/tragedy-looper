package first_steps

import (
	"tragedy-looper/engine/internal/models"
)

func NewFirstSteps1() *models.Script {
	// FirstSteps1 初学者剧本
	// RuleY Murder Plan 谋杀计划
	// RuleX1 Shadow of the Ripper 开膛手魔影 Rule2
	// Character 男生	Person 人
	// Character 女学生	Key Person 关键人物
	// Character 巫女	Serial Killer 连环杀手
	// Character 警官	Conspiracy Theorist 阴谋论者
	// Character 上班族	Killer 杀手
	// Character 医生	Brain 黑幕

	firstSteps1 := &models.Script{
		MainPlot:   MurderPlan,
		SubPlots:   []*models.Plot{ShadowOfTheRipper},
		Characters: make([]*models.Character, 0),
		Incidents: []models.Incident{
			&MurderIncident{},
			&SuicideIncident{},
		},
		MaxLoops:    3,
		DaysPerLoop: 3,
	}

	// 创建角色及其对应的身份
	characters := []*models.Character{
		// 男生 - Person(普通人)
		NewBoyStudent(&models.Role{
			Type: models.RolePersonType,
			Name: "Person",
			Abilities: []models.RoleAbility{
				&models.RolePerson{},
			},
		}),

		// 女学生 - Key Person(关键人物)
		NewGirlStudent(&models.Role{
			Type: KeyPerson,
			Name: "Key Person",
			Abilities: []models.RoleAbility{
				&KeyPersonRoleAbility{},
			},
		}),

		// 巫女 - Serial Killer(连环杀手)
		NewShrineMaiden(&models.Role{
			Type: SerialKiller,
			Name: "Serial Killer",
			Abilities: []models.RoleAbility{
				&SerialKillerAbility{},
			},
		}),

		// 警察 - Conspiracy Theorist(阴谋论者)
		NewPoliceOfficer(&models.Role{
			Type: ConspiracyTheorist,
			Name: "Conspiracy Theorist",
			Abilities: []models.RoleAbility{
				&ConspiracyTheoristAbility{},
			},
		}),

		// 上班族 - Killer(杀手)
		NewOfficeWorker(&models.Role{
			Type: Killer,
			Name: "Killer",
			Abilities: []models.RoleAbility{
				&KillerAbility{},
			},
		}),

		// 医生 - Brain(黑幕)
		NewDoctor(&models.Role{
			Type: Brain,
			Name: "Brain",
			Abilities: []models.RoleAbility{
				&BrainAbility{},
			},
		}),
	}

	firstSteps1.Characters = characters

	return firstSteps1
}
