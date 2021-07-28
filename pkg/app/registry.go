package app

import (
	"fmt"

	"github.com/kyma-incubator/reconciler/pkg/chart"
	"github.com/kyma-incubator/reconciler/pkg/cluster"
	"github.com/kyma-incubator/reconciler/pkg/db"
	"github.com/kyma-incubator/reconciler/pkg/kv"
	"github.com/kyma-incubator/reconciler/pkg/logger"
	"github.com/kyma-incubator/reconciler/pkg/metrics"
	"github.com/kyma-incubator/reconciler/pkg/workspace"
	"go.uber.org/zap"
)

type ApplicationRegistry struct {
	debug             bool
	logger            *zap.Logger
	connectionFactory db.ConnectionFactory
	inventory         cluster.Inventory
	kvRepository      *kv.Repository
	workspaceFactory  *workspace.Factory
	chartProvider     *chart.Provider
	initialized       bool
}

func NewApplicationRegistry(cf db.ConnectionFactory, debug bool) (*ApplicationRegistry, error) {
	or := &ApplicationRegistry{
		debug:             debug,
		connectionFactory: cf,
	}
	if err := or.init(); err != nil {
		return nil, err
	}
	return or, nil
}

func (or *ApplicationRegistry) init() error {
	if or.initialized {
		return nil
	}
	var err error
	if or.logger, err = logger.NewLogger(or.debug); err != nil {
		return err
	}
	if or.inventory, err = or.initInventory(); err != nil {
		return err
	}
	if or.kvRepository, err = or.initRepository(); err != nil {
		return err
	}
	or.workspaceFactory = &workspace.Factory{}
	if or.chartProvider, err = or.initChartProvider(); err != nil {
		return err
	}
	or.initialized = true
	return nil
}

func (or *ApplicationRegistry) Close() error {
	if !or.initialized {
		return nil
	}
	if err := or.kvRepository.Close(); err != nil {
		return err
	}
	return nil
}

func (or *ApplicationRegistry) Logger() *zap.Logger {
	return or.logger
}

func (or *ApplicationRegistry) Inventory() cluster.Inventory {
	return or.inventory
}

func (or *ApplicationRegistry) KVRepository() *kv.Repository {
	return or.kvRepository
}

func (or *ApplicationRegistry) WorkspaceFactory() *workspace.Factory {
	return or.workspaceFactory
}

func (or *ApplicationRegistry) ChartProvider() *chart.Provider {
	return or.chartProvider
}

func (or *ApplicationRegistry) initRepository() (*kv.Repository, error) {
	var err error

	var repository *kv.Repository
	if or.connectionFactory == nil {
		or.logger.Fatal("Failed to create configuration entry repository because connection factory is undefined")
	}
	repository, err = kv.NewRepository(or.connectionFactory, or.debug)
	if err != nil {
		or.logger.Error(fmt.Sprintf("Failed to create configuration entry repository: %s", err))
		return nil, err
	}

	return repository, nil
}

func (or *ApplicationRegistry) initInventory() (cluster.Inventory, error) {
	var err error

	if or.connectionFactory == nil {
		or.logger.Fatal("Failed to create cluster inventory because connection factory is undefined")
	}
	collector := metrics.NewReconciliationStatusCollector()
	or.inventory, err = cluster.NewInventory(or.connectionFactory, or.debug, collector)
	if err != nil {
		or.logger.Error(fmt.Sprintf("Failed to create cluster inventory: %s", err))
		return nil, err
	}

	return or.inventory, nil
}

func (or *ApplicationRegistry) initChartProvider() (*chart.Provider, error) {
	var err error

	or.chartProvider, err = chart.NewProvider(or.workspaceFactory, or.debug)
	if err != nil {
		return nil, err
	}

	return or.chartProvider, nil
}
