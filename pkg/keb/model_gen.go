// Package keb provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.8.2 DO NOT EDIT.
package keb

import (
	"time"
)

// Defines values for Status.
const (
	StatusDeleteError Status = "delete_error"

	StatusDeletePending Status = "delete_pending"

	StatusDeleted Status = "deleted"

	StatusDeleting Status = "deleting"

	StatusError Status = "error"

	StatusReady Status = "ready"

	StatusReconcileDisabled Status = "reconcile_disabled"

	StatusReconcilePending Status = "reconcile_pending"

	StatusReconcileSkipped Status = "reconcile_skipped"

	StatusReconciling Status = "reconciling"
)

// HTTPClusterResponse defines model for HTTPClusterResponse.
type HTTPClusterResponse struct {
	Cluster              string     `json:"cluster"`
	ClusterVersion       int64      `json:"clusterVersion"`
	ConfigurationVersion int64      `json:"configurationVersion"`
	Failures             *[]Failure `json:"failures,omitempty"`
	Status               Status     `json:"status"`
	StatusURL            string     `json:"statusURL"`
}

// HTTPClusterStatusResponse defines model for HTTPClusterStatusResponse.
type HTTPClusterStatusResponse struct {
	StatusChanges []StatusChange `json:"statusChanges"`
}

// HTTPErrorResponse defines model for HTTPErrorResponse.
type HTTPErrorResponse struct {
	Error string `json:"error"`
}

// HTTPReconcilerStatus defines model for HTTPReconcilerStatus.
type HTTPReconcilerStatus []Reconciliation

// HTTPReconciliationOperations defines model for HTTPReconciliationOperations.
type HTTPReconciliationOperations struct {
	Cluster    *Cluster     `json:"cluster,omitempty"`
	Operations *[]Operation `json:"operations,omitempty"`
}

// Cluster defines model for cluster.
type Cluster struct {
	// valid kubeconfig to cluster
	Kubeconfig   string       `json:"kubeconfig"`
	KymaConfig   KymaConfig   `json:"kymaConfig"`
	Metadata     Metadata     `json:"metadata"`
	RuntimeID    string       `json:"runtimeID"`
	RuntimeInput RuntimeInput `json:"runtimeInput"`
}

// Component defines model for component.
type Component struct {
	URL           string          `json:"URL"`
	Component     string          `json:"component"`
	Configuration []Configuration `json:"configuration"`
	Namespace     string          `json:"namespace"`
	Version       string          `json:"version"`
}

// Configuration defines model for configuration.
type Configuration struct {
	Key    string      `json:"key"`
	Secret bool        `json:"secret"`
	Value  interface{} `json:"value"`
}

// Failure defines model for failure.
type Failure struct {
	Component string `json:"component"`
	Reason    string `json:"reason"`
}

// KymaConfig defines model for kymaConfig.
type KymaConfig struct {
	Administrators []string    `json:"administrators"`
	Components     []Component `json:"components"`
	Profile        string      `json:"profile"`
	Version        string      `json:"version"`
}

// Metadata defines model for metadata.
type Metadata struct {
	GlobalAccountID string `json:"globalAccountID"`
	InstanceID      string `json:"instanceID"`
	Region          string `json:"region"`
	ServiceID       string `json:"serviceID"`
	ServicePlanID   string `json:"servicePlanID"`
	ServicePlanName string `json:"servicePlanName"`
	ShootName       string `json:"shootName"`
	SubAccountID    string `json:"subAccountID"`
}

// Operation defines model for operation.
type Operation struct {
	ClusterMetadata *Cluster  `json:"clusterMetadata,omitempty"`
	Component       string    `json:"component"`
	CorrelationID   string    `json:"correlationID"`
	Created         time.Time `json:"created"`
	Priority        int64     `json:"priority"`
	Reason          string    `json:"reason"`
	SchedulingID    string    `json:"schedulingID"`
	State           string    `json:"state"`
	Updated         time.Time `json:"updated"`
}

// OperationStop defines model for operationStop.
type OperationStop struct {
	Reason string `json:"reason"`
}

// ReconcilerStatus defines model for reconcilerStatus.
type ReconcilerStatus struct {
	Cluster  string    `json:"cluster"`
	Created  time.Time `json:"created"`
	Metadata Metadata  `json:"metadata"`
	Status   string    `json:"status"`
}

// Reconciliation defines model for reconciliation.
type Reconciliation struct {
	Created      time.Time `json:"created"`
	Lock         string    `json:"lock"`
	RuntimeID    string    `json:"runtimeID"`
	SchedulingID string    `json:"schedulingID"`
	Status       Status    `json:"status"`
	Updated      time.Time `json:"updated"`
}

// RuntimeInput defines model for runtimeInput.
type RuntimeInput struct {
	Description string `json:"description"`
	Name        string `json:"name"`
}

// Status defines model for status.
type Status string

// StatusChange defines model for statusChange.
type StatusChange struct {
	Duration int64     `json:"duration"`
	Started  time.Time `json:"started"`
	Status   Status    `json:"status"`
}

// StatusUpdate defines model for statusUpdate.
type StatusUpdate struct {
	Status Status `json:"status"`
}

// BadRequest defines model for BadRequest.
type BadRequest HTTPErrorResponse

// ClusterNotFound defines model for ClusterNotFound.
type ClusterNotFound HTTPErrorResponse

// InternalError defines model for InternalError.
type InternalError HTTPErrorResponse

// Ok defines model for Ok.
type Ok HTTPClusterResponse

// ReconcilationOperationsOKResponse defines model for ReconcilationOperationsOKResponse.
type ReconcilationOperationsOKResponse HTTPReconciliationOperations

// ReconcilationsOKResponse defines model for ReconcilationsOKResponse.
type ReconcilationsOKResponse HTTPReconcilerStatus

// PostClustersJSONBody defines parameters for PostClusters.
type PostClustersJSONBody Cluster

// PutClustersJSONBody defines parameters for PutClusters.
type PutClustersJSONBody Cluster

// PutClustersRuntimeIDStatusJSONBody defines parameters for PutClustersRuntimeIDStatus.
type PutClustersRuntimeIDStatusJSONBody StatusUpdate

// PostOperationsSchedulingIDCorrelationIDStopJSONBody defines parameters for PostOperationsSchedulingIDCorrelationIDStop.
type PostOperationsSchedulingIDCorrelationIDStopJSONBody OperationStop

// GetReconciliationsParams defines parameters for GetReconciliations.
type GetReconciliationsParams struct {
	RuntimeID *[]string `json:"runtimeID,omitempty"`
	Status    *[]Status `json:"status,omitempty"`
}

// PostClustersJSONRequestBody defines body for PostClusters for application/json ContentType.
type PostClustersJSONRequestBody PostClustersJSONBody

// PutClustersJSONRequestBody defines body for PutClusters for application/json ContentType.
type PutClustersJSONRequestBody PutClustersJSONBody

// PutClustersRuntimeIDStatusJSONRequestBody defines body for PutClustersRuntimeIDStatus for application/json ContentType.
type PutClustersRuntimeIDStatusJSONRequestBody PutClustersRuntimeIDStatusJSONBody

// PostOperationsSchedulingIDCorrelationIDStopJSONRequestBody defines body for PostOperationsSchedulingIDCorrelationIDStop for application/json ContentType.
type PostOperationsSchedulingIDCorrelationIDStopJSONRequestBody PostOperationsSchedulingIDCorrelationIDStopJSONBody
