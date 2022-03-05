package application

import (
	"github.com/hxzhouh/zrpc/pkg/logger"
	"github.com/hxzhouh/zrpc/pkg/server"
	"github.com/hxzhouh/zrpc/pkg/server/signals"
	"github.com/hxzhouh/zrpc/pkg/utils/xgo"
	"github.com/hxzhouh/zrpc/pkg/utils/zcycle"
	"go.uber.org/zap"
	"golang.org/x/sync/errgroup"
	"sync"
)

type Application struct {
	HideBanner bool // 是否关闭 banner
	smu        *sync.RWMutex
	logger     *zap.Logger
	initOnce   sync.Once
	stopped    chan struct{}
	servers    []server.Server //可以是grpc  服务，也可以是http服务。
	//server     grpc.Server
	// startupOnce  sync.Once
	stopOnce sync.Once
	cycle    *zcycle.Cycle
}

func (app *Application) initialize() {
	app.initOnce.Do(func() {
		app.logger = logger.DefaultLogger
		_ = app.printBanner()
		app.cycle = zcycle.NewCycle()
	})
}
func (app *Application) Serve(s ...server.Server) error {
	app.smu.Lock()
	defer app.smu.Unlock()
	app.servers = append(app.servers, s...)
	return nil
}

func (app *Application) Startup(fns ...func() error) error {
	app.initialize()
	return xgo.SerialUntilError(fns...)()
}

//printBanner init
func (app *Application) printBanner() error {
	if app.HideBanner {
		return nil
	}
	const banner = `Welcome to zrpc, starting application ...`
	app.logger.Info(banner)
	return nil
}

// Run run application
func (app *Application) Run(servers ...server.Server) error {
	app.smu.Lock()
	app.servers = append(app.servers, servers...)
	app.smu.Unlock()

	app.waitSignals() //start signal listen task in goroutine
	defer app.clean()
	//
	//// todo jobs not graceful
	//_ = app.startJobs()

	// start servers and govern server
	app.cycle.Run(app.startServers)
	// start workers
	//app.cycle.Run(app.startWorkers)

	//blocking and wait quit
	if err := <-app.cycle.Wait(); err != nil {
		app.logger.Error("zrpc shutdown with error", zap.Error(err))
		return err
	}
	app.logger.Info("shutdown zrpc, bye!")
	return nil
}

//clean after app quit
func (app *Application) clean() {
	_ = app.logger.Sync()
}

// 依次启动挂载再上面的服务。
func (app *Application) startServers() error {
	var eg errgroup.Group
	//var ctx, cancel = context.WithTimeout(context.Background(), 3*time.Second)
	go func() {
		<-app.stopped
		//cancel()
	}()
	// start multi servers
	for _, s := range app.servers {
		s := s
		eg.Go(func() (err error) {
			//defer func() {
			//	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
			//	defer cancel()
			//	_ = registry.DefaultRegisterer.UnregisterService(ctx, s.Info())
			//	//app.logger.Info("exit server", xlog.FieldMod(ecode.ModApp), xlog.FieldEvent("exit"), xlog.FieldName(s.Info().Name), xlog.FieldErr(err), xlog.FieldAddr(s.Info().Label()))
			//}()
			//
			//time.AfterFunc(time.Second, func() {
			//	_ = registry.DefaultRegisterer.RegisterService(ctx, s.Info())
			//	//app.logger.Info("start server", xlog.FieldMod(ecode.ModApp), xlog.FieldEvent("init"), xlog.FieldName(s.Info().Name), xlog.FieldAddr(s.Info().Label()), xlog.Any("scheme", s.Info().Scheme))
			//})
			err = s.Serve()
			return
		})
	}
	return eg.Wait()
}

func (app *Application) waitSignals() {
	app.logger.Info("init listen signal")
	signals.Shutdown(func(grace bool) { //when get shutdown signal

		//if grace {
		//	_ = app.GracefulStop(context.TODO())
		//} else {
		_ = app.Stop()
		//}
	})
}

// Stop application immediately after necessary cleanup
func (app *Application) Stop() (err error) {
	app.stopOnce.Do(func() {
		//	app.stopped <- struct{}{}
		//	app.runHooks(hooks.Stage_BeforeStop)
		//
		//	//stop servers
		//	app.smu.RLock()
		//	for _, s := range app.servers {
		//		func(s server.Server) {
		//			app.cycle.Run(s.Stop)
		//		}(s)
		//	}
		//	app.smu.RUnlock()
		//
		//	//stop workers
		//	for _, w := range app.workers {
		//		func(w worker.Worker) {
		//			app.cycle.Run(w.Stop)
		//		}(w)
		//	}
		//	<-app.cycle.Done()
		//	app.runHooks(hooks.Stage_AfterStop)
		//	app.cycle.Close()
	})
	return
}
