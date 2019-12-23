package operator

import (
	"fmt"

	"github.com/3scale/3scale-operator/pkg/3scale/amp/component"
	appsv1alpha1 "github.com/3scale/3scale-operator/pkg/apis/apps/v1alpha1"
	"github.com/3scale/3scale-operator/pkg/helper"
	v1 "k8s.io/api/core/v1"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

type OperatorBackendOptionsProvider struct {
	apimanagerSpec *appsv1alpha1.APIManagerSpec
	namespace      string
	client         client.Client
	backendOptions *component.BackendOptions
	secretSource   *helper.SecretSource
}

func NewOperatorBackendOptionsProvider(apimanagerSpec *appsv1alpha1.APIManagerSpec, namespace string, client client.Client) *OperatorBackendOptionsProvider {
	return &OperatorBackendOptionsProvider{
		apimanagerSpec: apimanagerSpec,
		namespace:      namespace,
		client:         client,
		backendOptions: component.NewBackendOptions(),
		secretSource:   helper.NewSecretSource(client, namespace),
	}
}

func (o *OperatorBackendOptionsProvider) GetBackendOptions() (*component.BackendOptions, error) {
	o.backendOptions.AppLabel = *o.apimanagerSpec.AppLabel
	o.backendOptions.TenantName = *o.apimanagerSpec.TenantName
	o.backendOptions.WildcardDomain = o.apimanagerSpec.WildcardDomain

	err := o.setSecretBasedOptions()
	if err != nil {
		return nil, err
	}

	o.setResourceRequirementsOptions()
	o.setReplicas()

	err = o.backendOptions.Validate()
	return o.backendOptions, err
}

func (o *OperatorBackendOptionsProvider) setSecretBasedOptions() error {
	cases := []struct {
		field       *string
		secretName  string
		secretField string
		defValue    string
	}{
		{
			&o.backendOptions.SystemBackendUsername,
			component.BackendSecretInternalApiSecretName,
			component.BackendSecretInternalApiUsernameFieldName,
			component.DefaultSystemBackendUsername(),
		},
		{
			&o.backendOptions.SystemBackendPassword,
			component.BackendSecretInternalApiSecretName,
			component.BackendSecretInternalApiPasswordFieldName,
			component.DefaultSystemBackendPassword(),
		},
		{
			&o.backendOptions.ServiceEndpoint,
			component.BackendSecretBackendListenerSecretName,
			component.BackendSecretBackendListenerServiceEndpointFieldName,
			component.DefaultBackendServiceEndpoint(),
		},
		{
			&o.backendOptions.RouteEndpoint,
			component.BackendSecretBackendListenerSecretName,
			component.BackendSecretBackendListenerRouteEndpointFieldName,
			fmt.Sprintf("https://backend-%s.%s", *o.apimanagerSpec.TenantName, o.apimanagerSpec.WildcardDomain),
		},
		{
			&o.backendOptions.StorageURL,
			component.BackendSecretBackendRedisSecretName,
			component.BackendSecretBackendRedisStorageURLFieldName,
			component.DefaultBackendRedisStorageURL(),
		},
		{
			&o.backendOptions.QueuesURL,
			component.BackendSecretBackendRedisSecretName,
			component.BackendSecretBackendRedisQueuesURLFieldName,
			component.DefaultBackendRedisQueuesURL(),
		},
		{
			&o.backendOptions.StorageSentinelHosts,
			component.BackendSecretBackendRedisSecretName,
			component.BackendSecretBackendRedisStorageSentinelHostsFieldName,
			component.DefaultBackendStorageSentinelHosts(),
		},
		{
			&o.backendOptions.StorageSentinelRole,
			component.BackendSecretBackendRedisSecretName,
			component.BackendSecretBackendRedisStorageSentinelRoleFieldName,
			component.DefaultBackendStorageSentinelRole(),
		},
		{
			&o.backendOptions.QueuesSentinelHosts,
			component.BackendSecretBackendRedisSecretName,
			component.BackendSecretBackendRedisQueuesSentinelHostsFieldName,
			component.DefaultBackendQueuesSentinelHosts(),
		},
		{
			&o.backendOptions.QueuesSentinelRole,
			component.BackendSecretBackendRedisSecretName,
			component.BackendSecretBackendRedisQueuesSentinelRoleFieldName,
			component.DefaultBackendQueuesSentinelRole(),
		},
	}

	for _, option := range cases {
		val, err := o.secretSource.FieldValue(option.secretName, option.secretField, option.defValue)
		if err != nil {
			return err
		}
		// not nil value is ensured
		*option.field = *val
	}

	return nil
}

func (o *OperatorBackendOptionsProvider) setResourceRequirementsOptions() {
	if *o.apimanagerSpec.ResourceRequirementsEnabled {
		o.backendOptions.ListenerResourceRequirements = component.DefaultBackendListenerResourceRequirements()
		o.backendOptions.WorkerResourceRequirements = component.DefaultBackendWorkerResourceRequirements()
		o.backendOptions.CronResourceRequirements = component.DefaultCronResourceRequirements()
	} else {
		o.backendOptions.ListenerResourceRequirements = v1.ResourceRequirements{}
		o.backendOptions.WorkerResourceRequirements = v1.ResourceRequirements{}
		o.backendOptions.CronResourceRequirements = v1.ResourceRequirements{}
	}
}

func (o *OperatorBackendOptionsProvider) setReplicas() {
	o.backendOptions.ListenerReplicas = int32(*o.apimanagerSpec.Backend.ListenerSpec.Replicas)
	o.backendOptions.WorkerReplicas = int32(*o.apimanagerSpec.Backend.WorkerSpec.Replicas)
	o.backendOptions.CronReplicas = int32(*o.apimanagerSpec.Backend.CronSpec.Replicas)
}
