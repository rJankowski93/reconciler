// Code generated by mockery 2.7.5. DO NOT EDIT.

package mock

import (
	actions "github.com/kyma-incubator/reconciler/pkg/reconciler/instances/istio/actions"
	kubernetes "github.com/kyma-incubator/reconciler/pkg/reconciler/kubernetes"

	mock "github.com/stretchr/testify/mock"

	workspace "github.com/kyma-incubator/reconciler/pkg/reconciler/workspace"

	zap "go.uber.org/zap"
)

// IstioPerformer is an autogenerated mock type for the IstioPerformer type
type IstioPerformer struct {
	mock.Mock
}

// Install provides a mock function with given fields: kubeConfig, manifest, logger
func (_m *IstioPerformer) Install(kubeConfig string, manifest string, logger *zap.SugaredLogger) error {
	ret := _m.Called(kubeConfig, manifest, logger)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string, *zap.SugaredLogger) error); ok {
		r0 = rf(kubeConfig, manifest, logger)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// PatchMutatingWebhook provides a mock function with given fields: kubeClient, logger
func (_m *IstioPerformer) PatchMutatingWebhook(kubeClient kubernetes.Client, logger *zap.SugaredLogger) error {
	ret := _m.Called(kubeClient, logger)

	var r0 error
	if rf, ok := ret.Get(0).(func(kubernetes.Client, *zap.SugaredLogger) error); ok {
		r0 = rf(kubeClient, logger)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ResetProxy provides a mock function with given fields: kubeConfig, version, logger
func (_m *IstioPerformer) ResetProxy(kubeConfig string, version actions.IstioVersion, logger *zap.SugaredLogger) error {
	ret := _m.Called(kubeConfig, version, logger)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, actions.IstioVersion, *zap.SugaredLogger) error); ok {
		r0 = rf(kubeConfig, version, logger)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Update provides a mock function with given fields: kubeConfig, manifest, logger
func (_m *IstioPerformer) Update(kubeConfig string, manifest string, logger *zap.SugaredLogger) error {
	ret := _m.Called(kubeConfig, manifest, logger)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string, *zap.SugaredLogger) error); ok {
		r0 = rf(kubeConfig, manifest, logger)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Version provides a mock function with given fields: _a0, branchVersion, istioChart, kubeConfig, logger
func (_m *IstioPerformer) Version(_a0 workspace.Factory, branchVersion string, istioChart string, kubeConfig string, logger *zap.SugaredLogger) (actions.IstioVersion, error) {
	ret := _m.Called(_a0, branchVersion, istioChart, kubeConfig, logger)

	var r0 actions.IstioVersion
	if rf, ok := ret.Get(0).(func(workspace.Factory, string, string, string, *zap.SugaredLogger) actions.IstioVersion); ok {
		r0 = rf(_a0, branchVersion, istioChart, kubeConfig, logger)
	} else {
		r0 = ret.Get(0).(actions.IstioVersion)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(workspace.Factory, string, string, string, *zap.SugaredLogger) error); ok {
		r1 = rf(_a0, branchVersion, istioChart, kubeConfig, logger)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
