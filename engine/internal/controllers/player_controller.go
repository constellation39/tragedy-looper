package controllers

var MastermindCLI *CLI
var ProtagonistCLI *CLI

func SetMastermindCLI(cli *CLI) {
	MastermindCLI = cli
}

func SetProtagonistCLI(cli *CLI) {
	ProtagonistCLI = cli
}
