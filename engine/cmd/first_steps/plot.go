package first_steps

import "tragedy-looper/engine/internal/models"

// Implement the specific plot rules

// Light of the Avenger failure rule
type LightOfTheAvengerFailureRule struct{}

func (r *LightOfTheAvengerFailureRule) CheckCondition(gameState *models.GameState) bool {
	// Find the Brain character
	//for _, character := range gameState.Characters {
	//	if character.Role() != nil && character.Role().Type == models.Brain {
	//		// Check if the Brain's starting location has at least 2 Intrigue counters
	//		location := character.data.StartingLocation
	//		intrigueCount := gameState.Locations[location].IntrigueCount
	//		return intrigueCount >= 2
	//	}
	//}
	return false
}

func (r *LightOfTheAvengerFailureRule) GetTiming() models.DayPhase {
	return models.PhaseDayEnd
}

func (r *LightOfTheAvengerFailureRule) GetRuleType() models.RuleType {
	return models.Failure
}

func (r *LightOfTheAvengerFailureRule) GetDescription() string {
	return "At loop end, if there are at least 2 Intrigue counters on the Brain's starting location, the Protagonists lose."
}

// A Place to Protect failure rule
type APlaceToProtectFailureRule struct{}

func (r *APlaceToProtectFailureRule) CheckCondition(gameState *models.GameState) bool {
	// Check if the School has at least 2 Intrigue counters
	//intrigueCount := gameState.Locations[models.LocationSchool].IntrigueCount
	//return intrigueCount >= 2
	return false
}

func (r *APlaceToProtectFailureRule) GetTiming() models.DayPhase {
	return models.PhaseDayEnd
}

func (r *APlaceToProtectFailureRule) GetRuleType() models.RuleType {
	return models.Failure
}

func (r *APlaceToProtectFailureRule) GetDescription() string {
	return "At loop end, if there are at least 2 Intrigue counters on the School, the Protagonists lose."
}

// An Unsettling Rumor optional mastermind ability
type AnUnsettlingRumorOptionalRule struct {
	UsedThisLoop bool
}

func (r *AnUnsettlingRumorOptionalRule) CheckCondition(gameState *models.GameState) bool {
	// This rule doesn't check a condition; it's an ability
	return false
}

func (r *AnUnsettlingRumorOptionalRule) Execute(gameState *models.GameState) {
	//if !r.UsedThisLoop {
	//	// The Mastermind chooses a location to add 1 Intrigue counter
	//	// Placeholder for the actual location selection logic
	//	location := models.LocationChosenByMastermind()
	//	gameState.Locations[location].IntrigueCount++
	//	r.UsedThisLoop = true
	//}
}

func (r *AnUnsettlingRumorOptionalRule) GetTiming() models.DayPhase {
	return models.PhaseMastermindAbilities
}

func (r *AnUnsettlingRumorOptionalRule) GetRuleType() models.RuleType {
	return models.Optional
}

func (r *AnUnsettlingRumorOptionalRule) GetDescription() string {
	return "Once per loop, the Mastermind may add an Intrigue counter to any location."
}

var MurderPlan,
	LightOfTheAvenger,
	APlaceToProtect,
	ShadowOfTheRipper,
	AnUnsettlingRumor,
	AHideousScript *models.Plot

// Define the plots using the structures and rules
func init() {
	// Murder Plan
	// Source: First Steps Main Plots in your knowledge base
	murderPlan := models.NewPlot("murder_plan", "Murder Plan", models.MainPlot, "Roles to add: Key Person, Brain, Killer.")
	murderPlan.AddRequiredRole("KeyPerson", 1)
	murderPlan.AddRequiredRole("Brain", 1)
	murderPlan.AddRequiredRole("Killer", 1)
	MurderPlan = murderPlan

	// Light of the Avenger
	// Source: First Steps Main Plots in your knowledge base
	lightOfTheAvenger := models.NewPlot("light_of_the_avenger", "Light of the Avenger", models.MainPlot, "At loop end, if there are at least 2 Intrigue counters on the Brain's starting location, the Protagonists lose.")
	lightOfTheAvenger.AddRequiredRole("Brain", 1)
	lightOfTheAvenger.AddRule(&LightOfTheAvengerFailureRule{})
	LightOfTheAvenger = lightOfTheAvenger

	// A Place to Protect
	// Source: First Steps Main Plots in your knowledge base
	aPlaceToProtect := models.NewPlot("a_place_to_protect", "A Place to Protect", models.MainPlot, "At loop end, if there are at least 2 Intrigue counters on the School, the Protagonists lose.")
	aPlaceToProtect.AddRequiredRole("KeyPerson", 1)
	aPlaceToProtect.AddRequiredRole("Cultist", 1)
	aPlaceToProtect.AddRule(&APlaceToProtectFailureRule{})
	APlaceToProtect = aPlaceToProtect

	// Shadow of the Ripper
	// Source: First Steps Subplots in your knowledge base
	shadowOfTheRipper := models.NewPlot("shadow_of_the_ripper", "Shadow of the Ripper", models.SubPlot, "Roles to add: Conspiracy Theorist, Serial Killer.")
	shadowOfTheRipper.AddRequiredRole("ConspiracyTheorist", 1)
	shadowOfTheRipper.AddRequiredRole("SerialKiller", 1)
	ShadowOfTheRipper = shadowOfTheRipper

	// An Unsettling Rumor
	// Source: First Steps Subplots in your knowledge base
	anUnsettlingRumor := models.NewPlot("an_unsettling_rumor", "An Unsettling Rumor", models.SubPlot, "Once per loop, the Mastermind may add an Intrigue counter to any location.")
	anUnsettlingRumor.AddRequiredRole("ConspiracyTheorist", 1)
	anUnsettlingRumor.AddRule(&AnUnsettlingRumorOptionalRule{})
	AnUnsettlingRumor = anUnsettlingRumor

	// A Hideous script
	// Source: First Steps Subplots in your knowledge base
	aHideousScript := models.NewPlot("a_hideous_script", "A Hideous script", models.SubPlot, "Roles to add: Conspiracy Theorist, Friend, 0-2 Curmudgeons.")
	aHideousScript.AddRequiredRole("ConspiracyTheorist", 1)
	aHideousScript.AddRequiredRole("Friend", 1)
	//aHideousScript.AddOptionalRole("Curmudgeon", 0, 2) // Allows adding 0 to 2 Curmudgeons
	AHideousScript = aHideousScript

}
