package main

import (
	"framework/logger"
	"go.uber.org/zap"
	"tragedy-looper/engine/cmd/first_steps"
	"tragedy-looper/engine/internal/controllers"
)

func main() {
	// 配置日志选项
	loggerOptions := logger.DefaultOptions()
	loggerOptions.Level = zap.DebugLevel
	loggerOptions.Filename = "logs/tragedy-looper.log"
	logging, err := logger.NewLogger(loggerOptions)
	if err != nil {
		panic("Failed to initialize logger: " + err.Error())
	}
	defer func(logging *zap.Logger) {
		_ = logging.Sync()
	}(logging)

	// 使用结构化的日志
	logging.Debug("Tragedy Looper game initializing...",
		zap.String("log_level", loggerOptions.Level.String()),
		zap.String("log_file", loggerOptions.Filename),
	)

	firstSteps1 := first_steps.NewFirstSteps1()
	logging.Debug("First steps scenario loaded",
		zap.String("scenario", "FirstSteps1"),
	)

	CLI := controllers.NewCLI(logging)

	controllers.SetMastermindCLI(CLI)
	controllers.SetProtagonistCLI(CLI)

	gameController := controllers.NewGameController(logging, firstSteps1)
	logging.Debug("Game controller initialized")

	logging.Debug("Attempting to start game...")
	if err = gameController.StartGame(); err != nil {
		logging.Error("Game start game failed",
			zap.Error(err),
			zap.String("scenario", "FirstSteps1"),
		)
		return
	}

	logging.Debug("Game successfully started",
		zap.String("status", "running"),
		zap.String("scenario", "FirstSteps1"),
	)
}
