/*
Copyright 2021.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package controllers

import (
	"context"

	apierrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/tools/record"
	"sigs.k8s.io/cluster-api/util"
	"sigs.k8s.io/cluster-api/util/annotations"
	"sigs.k8s.io/cluster-api/util/predicates"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"

	"sigs.k8s.io/controller-runtime/pkg/controller/controllerutil"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"

	"github.com/pkg/errors"

	"github.com/giantswarm/cluster-api-provider-kvm/api/v1alpha4"
	"github.com/giantswarm/cluster-api-provider-kvm/pkg/kvm/metrics"
	"github.com/giantswarm/cluster-api-provider-kvm/pkg/kvm/scope"
	"github.com/giantswarm/cluster-api-provider-kvm/pkg/kvm/services/namespace"
)

const (
	KVMClusterController = "kvmcluster"
)

// KVMClusterReconciler reconciles a KVMCluster object
type KVMClusterReconciler struct {
	client.Client
	Recorder         record.EventRecorder
	Scheme           *runtime.Scheme
	WatchFilterValue string
}

//+kubebuilder:rbac:groups=infrastructure.cluster.x-k8s.io.infrastructure.cluster.x-k8s.io,resources=kvmclusters,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=infrastructure.cluster.x-k8s.io.infrastructure.cluster.x-k8s.io,resources=kvmclusters/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=infrastructure.cluster.x-k8s.io.infrastructure.cluster.x-k8s.io,resources=kvmclusters/finalizers,verbs=update

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
// TODO(user): Modify the Reconcile function to compare the state specified by
// the KVMCluster object against the actual cluster state, and then
// perform operations to make the cluster state reflect the state specified by
// the user.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.8.3/pkg/reconcile
func (r *KVMClusterReconciler) Reconcile(ctx context.Context, req ctrl.Request) (_ ctrl.Result, reconcileError error) {
	log := ctrl.LoggerFrom(ctx)

	// Fetch the KVMCluster instance
	kvmCluster := &v1alpha4.KVMCluster{}
	err := r.Get(ctx, req.NamespacedName, kvmCluster)
	if err != nil {
		if apierrors.IsNotFound(err) {
			return reconcile.Result{}, nil
		}
		return reconcile.Result{}, err
	}

	defer func() {
		if reconcileError != nil {
			r.Recorder.Event(kvmCluster, "Warning", "failed-to-reconcile", reconcileError.Error())
			return
		}

		metrics.CaptureLastReconciled(KVMClusterController)
	}()

	// Fetch the Cluster.
	cluster, err := util.GetOwnerCluster(ctx, r.Client, kvmCluster.ObjectMeta)
	if err != nil {
		return reconcile.Result{}, err
	}

	if cluster == nil {
		log.Info("Cluster Controller has not yet set OwnerRef")
		return reconcile.Result{}, nil
	}

	if annotations.IsPaused(cluster, kvmCluster) {
		log.Info("KVMCluster or linked Cluster is marked as paused. Won't reconcile")
		return reconcile.Result{}, nil
	}

	// log = log.WithValues("cluster", cluster.Name)

	// Create the scope.
	clusterScope, err := scope.NewClusterScope(scope.ClusterScopeParams{
		Client:         r.Client,
		Cluster:        cluster,
		KVMCluster:     kvmCluster,
		ControllerName: KVMClusterController,
	})
	if err != nil {
		return reconcile.Result{}, errors.Errorf("failed to create scope: %+v", err)
	}

	// Always close the scope when exiting this function so we can persist any KVMCluster changes.
	defer func() {
		if err := clusterScope.Close(); err != nil && reconcileError == nil {
			reconcileError = err
		}
	}()

	// Handle deleted clusters
	if !kvmCluster.DeletionTimestamp.IsZero() {
		return r.reconcileDelete(ctx, clusterScope)
	}

	// Handle non-deleted clusters
	return r.reconcileNormal(ctx, clusterScope)
}

// SetupWithManager sets up the controller with the Manager.
func (r *KVMClusterReconciler) SetupWithManager(ctx context.Context, mgr ctrl.Manager) error {
	log := ctrl.LoggerFrom(ctx)

	return ctrl.NewControllerManagedBy(mgr).
		For(&v1alpha4.KVMCluster{}).
		WithEventFilter(predicates.ResourceNotPausedAndHasFilterLabel(log, r.WatchFilterValue)).
		Complete(r)
}

func (r *KVMClusterReconciler) reconcileDelete(ctx context.Context, scope *scope.ClusterScope) (_ reconcile.Result, reconcileError error) {

	nsService := namespace.NewService()

	if err := scope.PatchObject(); err != nil {
		return reconcile.Result{}, err
	}
	err := nsService.DeleteNamespace()
	if err != nil {
		return reconcile.Result{}, err
	}

	// Cluster is deleted so remove the finalizer.
	controllerutil.RemoveFinalizer(scope.KVMCluster, v1alpha4.ClusterFinalizer)

	return reconcile.Result{}, nil
}

func (r *KVMClusterReconciler) reconcileNormal(ctx context.Context, scope *scope.ClusterScope) (_ reconcile.Result, reconcileError error) {

	// If the KVMCluster doesn't have our finalizer, add it.
	controllerutil.AddFinalizer(scope.KVMCluster, v1alpha4.ClusterFinalizer)

	// Register the finalizer immediately to avoid orphaning resources on delete
	if err := scope.PatchObject(); err != nil {
		return reconcile.Result{}, err
	}

	nsService := namespace.NewService()
	err := nsService.ReconcileNamespace()
	if err != nil {
		return reconcile.Result{}, err
	}
	return reconcile.Result{}, nil
}
