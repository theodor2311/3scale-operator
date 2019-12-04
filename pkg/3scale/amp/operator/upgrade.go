package operator

import (
	appsv1alpha1 "github.com/3scale/3scale-operator/pkg/apis/apps/v1alpha1"
	"github.com/go-logr/logr"
	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
)

type UpgradeApiManager struct {
	Cr              *appsv1alpha1.APIManager
	Client          client.Client
	Logger          logr.Logger
	ApiClientReader client.Reader
	Scheme          *runtime.Scheme
}

func (u *UpgradeApiManager) Upgrade() (reconcile.Result, error) {
	res, err := u.upgradeImages()
	if res.Requeue || err != nil {
		return res, err
	}

	return reconcile.Result{}, nil
}

func (u *UpgradeApiManager) upgradeImages() (reconcile.Result, error) {
	res, err := u.upgradeAMPImageStreams()
	if res.Requeue || err != nil {
		return res, err
	}

	if !u.Cr.IsExternalDatabaseEnabled() {
		res, err = u.upgradeBackendRedisImageStream()
		if res.Requeue || err != nil {
			return res, err
		}

		res, err = u.upgradeSystemRedisImageStream()
		if res.Requeue || err != nil {
			return res, err
		}

		res, err = u.upgradeSystemDatabaseImageStream()
		if res.Requeue || err != nil {
			return res, err
		}
	}

	// TODO should there be some kind of coordination/wait between the ImageStream
	// tags creation and the trigger update/swap ??
	res, err = u.upgradeDeploymentConfigs()
	if res.Requeue || err != nil {
		return res, err
	}

	// TODO because we are deprecating the use of the "latest" ImageStreamTag
	// that tag is not necessary anymore. At this moment the upgrade does NOT take
	// care of removing that ImageStreamTag although it is not used anymore.
	// Should we take care of doing that?. Keep in mind that this would probably
	// require some kind of coordination / wait to make sure we do not remove the
	// tag until the DeploymentConfigs are successfully using the new tag

	return reconcile.Result{}, nil
}

func (u *UpgradeApiManager) upgradeDeploymentConfigs() (reconcile.Result, error) {
	res, err := u.upgradeAPIcastDeploymentConfigs()
	if res.Requeue || err != nil {
		return res, err
	}

	return reconcile.Result{}, nil
}

func (u *UpgradeApiManager) upgradeAPIcastDeploymentConfigs() (reconcile.Result, error) {
	baseReconciler := NewBaseReconciler(u.Client, u.ApiClientReader, u.Scheme, u.Logger)
	baseLogicReconciler := NewBaseLogicReconciler(baseReconciler)
	reconciler := NewApicastReconciler(NewBaseAPIManagerLogicReconciler(baseLogicReconciler, u.Cr))
	// TODO we are only interested on the DCs reconciliation, specifically on the ImageChange trigger
	// so maybe is not a good idea to implement upgrade procedure by reconcile procedure??
	// the Reconcile method will try to reconcile everything that is reconciled in the Apicast reconciler.
	// This is, DeploymentConfigs(and not only the ImageChange trigger), ConfigMaps, Secrets, ...
	return reconciler.Reconcile()
}

func (u *UpgradeApiManager) upgradeAMPImageStreams() (reconcile.Result, error) {
	// implement upgrade procedure by reconcile procedure
	baseReconciler := NewBaseReconciler(u.Client, u.ApiClientReader, u.Scheme, u.Logger)
	baseLogicReconciler := NewBaseLogicReconciler(baseReconciler)
	reconciler := NewAMPImagesReconciler(NewBaseAPIManagerLogicReconciler(baseLogicReconciler, u.Cr))
	return reconciler.Reconcile()
}

func (u *UpgradeApiManager) upgradeBackendRedisImageStream() (reconcile.Result, error) {
	redis, err := Redis(u.Cr)
	if err != nil {
		return reconcile.Result{}, err
	}

	baseReconciler := NewBaseReconciler(u.Client, u.ApiClientReader, u.Scheme, u.Logger)
	baseLogicReconciler := NewBaseLogicReconciler(baseReconciler)
	reconciler := NewImageStreamBaseReconciler(NewBaseAPIManagerLogicReconciler(baseLogicReconciler, u.Cr), NewImageStreamGenericReconciler())
	return reconcile.Result{}, reconciler.Reconcile(redis.BackendImageStream())
}

func (u *UpgradeApiManager) upgradeSystemRedisImageStream() (reconcile.Result, error) {
	redis, err := Redis(u.Cr)
	if err != nil {
		return reconcile.Result{}, err
	}

	baseReconciler := NewBaseReconciler(u.Client, u.ApiClientReader, u.Scheme, u.Logger)
	baseLogicReconciler := NewBaseLogicReconciler(baseReconciler)
	reconciler := NewImageStreamBaseReconciler(NewBaseAPIManagerLogicReconciler(baseLogicReconciler, u.Cr), NewImageStreamGenericReconciler())
	return reconcile.Result{}, reconciler.Reconcile(redis.SystemImageStream())
}

func (u *UpgradeApiManager) upgradeSystemDatabaseImageStream() (reconcile.Result, error) {
	if u.Cr.Spec.System.DatabaseSpec != nil && u.Cr.Spec.System.DatabaseSpec.PostgreSQL != nil {
		return u.upgradeSystemPostgreSQLImageStream()
	}

	// default is MySQL
	return u.upgradeSystemMySQLImageStream()
}

func (u *UpgradeApiManager) upgradeSystemMySQLImageStream() (reconcile.Result, error) {
	baseReconciler := NewBaseReconciler(u.Client, u.ApiClientReader, u.Scheme, u.Logger)
	baseLogicReconciler := NewBaseLogicReconciler(baseReconciler)
	reconciler := NewSystemMySQLImageReconciler(NewBaseAPIManagerLogicReconciler(baseLogicReconciler, u.Cr))
	return reconciler.Reconcile()
}

func (u *UpgradeApiManager) upgradeSystemPostgreSQLImageStream() (reconcile.Result, error) {
	baseReconciler := NewBaseReconciler(u.Client, u.ApiClientReader, u.Scheme, u.Logger)
	baseLogicReconciler := NewBaseLogicReconciler(baseReconciler)
	reconciler := NewSystemPostgreSQLImageReconciler(NewBaseAPIManagerLogicReconciler(baseLogicReconciler, u.Cr))
	return reconciler.Reconcile()
}
