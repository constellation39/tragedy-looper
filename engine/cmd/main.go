package main

import (
	"framework/logger"
	"go.uber.org/zap"
	"tragedy-looper/engine/cmd/first_steps"
	"tragedy-looper/engine/internal/controllers"
)

func main() {
	loggerOptions := logger.DefaultOptions()
	loggerOptions.Level = zap.DebugLevel
	loggerOptions.Filename = "logs/tragedy-looper.log"
	logging, err := logger.NewLogger(loggerOptions)
	if err != nil {
		panic(err)
	}
	defer func(logging *zap.Logger) {
		_ = logging.Sync()
	}(logging)

	logging.Info("Start Game.")

	// 1. 创建脚本
	firstSteps1 := first_steps.NewFirstSteps1()

	// 2. 使用 GameController 来控制游戏
	gameController := controllers.NewGameController(logging, firstSteps1)

	// 3. 启动游戏
	err = gameController.StartGame()
	if err != nil {
		logging.Error("Failed to start game.", zap.Error(err))
		return
	}

	logging.Info("Game Started.")
}
