package operator

import (
	"context"
	"errors"
	"fmt"

	"github.com/3scale/3scale-operator/pkg/helper"
	integreatlyv1alpha1 "github.com/integr8ly/grafana-operator/pkg/apis/integreatly/v1alpha1"
	"github.com/operator-framework/operator-sdk/pkg/k8sutil"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/client-go/discovery"
)

// ErrGrafanaDashboardsNotPresent custom error type
var ErrGrafanaDashboardsNotPresent = errors.New("no GrafanaDashboard registered with the API")

type GrafanaDashboardReconciler interface {
	IsUpdateNeeded(desired, existing *integreatlyv1alpha1.GrafanaDashboard) bool
}

type GrafanaDashboardBaseReconciler struct {
	BaseAPIManagerLogicReconciler
	reconciler GrafanaDashboardReconciler
}

func NewGrafanaDashboardBaseReconciler(baseAPIManagerLogicReconciler BaseAPIManagerLogicReconciler, reconciler GrafanaDashboardReconciler) *GrafanaDashboardBaseReconciler {
	return &GrafanaDashboardBaseReconciler{
		BaseAPIManagerLogicReconciler: baseAPIManagerLogicReconciler,
		reconciler:                    reconciler,
	}
}

func (r *GrafanaDashboardBaseReconciler) Reconcile(desired *integreatlyv1alpha1.GrafanaDashboard) error {
	objectInfo := ObjectInfo(desired)

	exists, err := r.hasGrafanaDashboards()
	if err != nil {
		return err
	}
	if !exists {
		r.Logger().Info("Install grafana-operator in your cluster to create grafanadashboards objects", "Error creating object", objectInfo)
		return nil
	}

	existing := &integreatlyv1alpha1.GrafanaDashboard{}
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

// hasGrafanaDashboards checks if GrafanaDashboard is registered in the cluster.
func (r *GrafanaDashboardBaseReconciler) hasGrafanaDashboards() (bool, error) {
	dc := discovery.NewDiscoveryClientForConfigOrDie(r.cfg)

	return k8sutil.ResourceExists(dc,
		integreatlyv1alpha1.SchemeGroupVersion.String(),
		integreatlyv1alpha1.GrafanaDashboardKind)
}

func (r *GrafanaDashboardBaseReconciler) isUpdateNeeded(desired, existing *integreatlyv1alpha1.GrafanaDashboard) (bool, error) {
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

type CreateOnlyGrafanaDashboardReconciler struct {
}

func NewCreateOnlyGrafanaDashboardReconciler() *CreateOnlyGrafanaDashboardReconciler {
	return &CreateOnlyGrafanaDashboardReconciler{}
}

func (r *CreateOnlyGrafanaDashboardReconciler) IsUpdateNeeded(desired, existing *integreatlyv1alpha1.GrafanaDashboard) bool {
	return false
}
