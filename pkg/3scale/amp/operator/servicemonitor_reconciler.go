package operator

import (
	"context"
	"errors"
	"fmt"

	"github.com/3scale/3scale-operator/pkg/helper"
	monitoringv1 "github.com/coreos/prometheus-operator/pkg/apis/monitoring/v1"
	"github.com/operator-framework/operator-sdk/pkg/k8sutil"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/client-go/discovery"
)

// ErrServiceMonitorsNotPresent custom error type
var ErrServiceMonitorsNotPresent = errors.New("no ServiceMonitors registered with the API")

type ServiceMonitorReconciler interface {
	IsUpdateNeeded(desired, existing *monitoringv1.ServiceMonitor) bool
}

type ServiceMonitorBaseReconciler struct {
	BaseAPIManagerLogicReconciler
	reconciler ServiceMonitorReconciler
}

func NewServiceMonitorBaseReconciler(baseAPIManagerLogicReconciler BaseAPIManagerLogicReconciler, reconciler ServiceMonitorReconciler) *ServiceMonitorBaseReconciler {
	return &ServiceMonitorBaseReconciler{
		BaseAPIManagerLogicReconciler: baseAPIManagerLogicReconciler,
		reconciler:                    reconciler,
	}
}

func (r *ServiceMonitorBaseReconciler) Reconcile(desired *monitoringv1.ServiceMonitor) error {
	objectInfo := ObjectInfo(desired)

	exists, err := r.hasServiceMonitors()
	if err != nil {
		return err
	}
	if !exists {
		r.Logger().Info("Install grafana-operator in your cluster to create grafanadashboards objects", "Error creating object", objectInfo)
		return nil
	}

	existing := &monitoringv1.ServiceMonitor{}
	err = r.Client().Get(
		context.TODO(),
		types.NamespacedName{Name: desired.Name, Namespace: r.apiManager.GetNamespace()},
		existing)
	if err != nil {
		if apierrors.IsNotFound(err) {
			createErr := r.createResource(desired)
			if createErr != nil {
				r.Logger().Error(createErr, fmt.Sprintf("Error creating object %s. Requeuing request...", objectInfo))
				return createErr
			}
			return nil
		}
		return err
	}

	update, err := r.isUpdateNeeded(desired, existing)
	if err != nil {
		return err
	}

	if update {
		return r.updateResource(existing)
	}

	return nil
}

// hasServiceMonitors checks if ServiceMonitor is registered in the cluster.
func (r *ServiceMonitorBaseReconciler) hasServiceMonitors() (bool, error) {
	dc := discovery.NewDiscoveryClientForConfigOrDie(r.cfg)

	return k8sutil.ResourceExists(dc,
		monitoringv1.SchemeGroupVersion.String(),
		monitoringv1.ServiceMonitorsKind)
}

func (r *ServiceMonitorBaseReconciler) isUpdateNeeded(desired, existing *monitoringv1.ServiceMonitor) (bool, error) {
	updated := helper.EnsureObjectMeta(&existing.ObjectMeta, &desired.ObjectMeta)

	updatedTmp, err := r.ensureOwnerReference(existing)
	if err != nil {
		return false, nil
	}

	updated = updated || updatedTmp

	updatedTmp = r.reconciler.IsUpdateNeeded(desired, existing)
	updated = updated || updatedTmp

	return updated, nil
}

type CreateOnlyServiceMonitorReconciler struct {
}

func NewCreateOnlyServiceMonitorReconciler() *CreateOnlyServiceMonitorReconciler {
	return &CreateOnlyServiceMonitorReconciler{}
}

func (r *CreateOnlyServiceMonitorReconciler) IsUpdateNeeded(desired, existing *monitoringv1.ServiceMonitor) bool {
	return false
}
