package cat

import (
	"os"
	"sync/atomic"
)

var isEnabled uint32 = 0

func Init(domain string) {
	if err := config.Init(domain); err != nil {
		logger.Warning("Cat initialize failed.")
		return
	}
	enable()

	go background(&router)
	go background(&monitor)
	go background(&sender)
	aggregator.Background()
}

func SetConfigFile(configFile string) {
	config.SetConfigFile(configFile)
}

func SetLogPath(logPaht string) {
	config.SetLogPath(logPaht)
}

func enable() {
	if atomic.SwapUint32(&isEnabled, 1) == 0 {
		logger.Info("Cat has been enabled.")
	}
}

func disable() {
	if atomic.SwapUint32(&isEnabled, 0) == 1 {
		logger.Info("Cat has been disabled.")
	}
}

func IsEnabled() bool {
	return atomic.LoadUint32(&isEnabled) > 0
}

func Shutdown() {
	scheduler.shutdown()
}

func DebugOn() {
	logger.logger.SetOutput(os.Stdout)
}
