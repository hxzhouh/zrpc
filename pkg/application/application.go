package application

import (
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"sync"
)

type Application struct {
	HideBanner bool // 是否关闭 banner
	logger     *zap.Logger
	initOnce   sync.Once
	server     grpc.Server
}

func (app *Application) initialize() {
	app.initOnce.Do(func() {
		//assign
		//app.cycle = xcycle.NewCycle()
		//app.smu = &sync.RWMutex{}
		//app.servers = make([]server.Server, 0)
		//app.workers = make([]worker.Worker, 0)
		//app.jobs = make(map[string]job.Runner)
		//app.logger = xlog.JupiterLogger
		//app.configParser = toml.Unmarshal
		//app.disableMap = make(map[Disable]bool)
		//app.stopped = make(chan struct{})
		//app.components = make([]component.Component, 0)
		////private method
		//
		//_ = app.parseFlags()
		_ = app.printBanner()
	})
}

func (app *Application) Startup(fns ...func() error) error {
	app.initialize()
	return nil
}

//printBanner init
func (app *Application) printBanner() error {
	if app.HideBanner {
		return nil
	}

	//if xdebug.IsTestingMode() {
	//	return nil
	//}

	const banner = `
 ____ ____ ____ ____ 
/_   /  __/  __/   _\
 /   |  \/|  \/|  /  
/   /|    |  __|  \_ 
\____\_/\_\_/  \____/
 Welcome to zrpc, starting application ...
`
	app.logger.Info(banner)
	return nil
}
