package cmd

import (
	"context"
	"github.com/kyma-incubator/reconciler/pkg/logger"
	"github.com/kyma-incubator/reconciler/pkg/scheduler/config"
	"github.com/kyma-incubator/reconciler/pkg/scheduler/service"
	"github.com/kyma-incubator/reconciler/pkg/scheduler/worker"
	"github.com/spf13/viper"
	"time"
)

func startScheduler(ctx context.Context, o *Options, configFile string) error {
	schedulerCfg, err := parseSchedulerConfig(configFile)
	if err != nil {
		return err
	}

	runtimeBuilder := service.NewRuntimeBuilder(o.Registry.ReconciliationRepository(), logger.NewLogger(o.Verbose))

	runtimeBuilder.
		RunRemote(
			o.Registry.Connnection(),
			o.Registry.Inventory(),
			schedulerCfg).
		WithWorkerPoolConfig(&worker.Config{
			PoolSize: o.Workers,
			//check-interval should be greater than "max-retires * retry-delay" to avoid queuing
			//of workers in case that component-reconciler isn't reachable
			OperationCheckInterval: 30 * time.Second,
			InvokerMaxRetries:      2,
			InvokerRetryDelay:      10 * time.Second,
		}).
		WithSchedulerConfig(
			&service.SchedulerConfig{
				InventoryWatchInterval:   o.WatchInterval,
				ClusterReconcileInterval: o.ClusterReconcileInterval,
				ClusterQueueSize:         10,
			}).
		WithBookkeeperConfig(&service.BookkeeperConfig{
			OperationsWatchInterval: 30 * time.Second,
			OrphanOperationTimeout:  o.OrphanOperationTimeout,
		}).
		Run(ctx)

	return nil
}

func parseSchedulerConfig(configFile string) (*config.Config, error) {
	viper.SetConfigFile(configFile)
	if err := viper.ReadInConfig(); err != nil {
		return &config.Config{}, err
	}

	var cfg config.Config
	return &cfg, viper.UnmarshalKey("mothership", &cfg)
}
