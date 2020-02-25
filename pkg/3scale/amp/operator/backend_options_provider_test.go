package operator

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/3scale/3scale-operator/pkg/3scale/amp/component"
	appsv1alpha1 "github.com/3scale/3scale-operator/pkg/apis/apps/v1alpha1"
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/controller-runtime/pkg/client/fake"
)

const (
	wildcardDomain             = "test.3scale.net"
	appLabel                   = "someLabel"
	apimanagerName             = "example-apimanager"
	namespace                  = "someNS"
	tenantName                 = "someTenant"
	trueValue                  = true
	listenerReplicaCount int64 = 3
	workerReplicaCount   int64 = 4
	cronReplicaCount     int64 = 5
)

func getInternalSecret(namespace string) *v1.Secret {
	return &v1.Secret{
		TypeMeta: metav1.TypeMeta{APIVersion: "v1", Kind: "Secret"},
		ObjectMeta: metav1.ObjectMeta{
			Name:      component.BackendSecretInternalApiSecretName,
			Namespace: namespace,
		},
		StringData: map[string]string{
			component.BackendSecretInternalApiUsernameFieldName: "someUserName",
			component.BackendSecretInternalApiPasswordFieldName: "somePasswd",
		},
		Type: v1.SecretTypeOpaque,
	}
}

func getListenerSecret(namespace string) *v1.Secret {
	return &v1.Secret{
		TypeMeta: metav1.TypeMeta{APIVersion: "v1", Kind: "Secret"},
		ObjectMeta: metav1.ObjectMeta{
			Name:      component.BackendSecretBackendListenerSecretName,
			Namespace: namespace,
		},
		StringData: map[string]string{
			component.BackendSecretBackendListenerServiceEndpointFieldName: "serviceValue",
			component.BackendSecretBackendListenerRouteEndpointFieldName:   "routeValue",
		},
		Type: v1.SecretTypeOpaque,
	}
}

func getRedisSecret(namespace string) *v1.Secret {
	return &v1.Secret{
		TypeMeta: metav1.TypeMeta{APIVersion: "v1", Kind: "Secret"},
		ObjectMeta: metav1.ObjectMeta{
			Name:      component.BackendSecretBackendRedisSecretName,
			Namespace: namespace,
		},
		StringData: map[string]string{
			component.BackendSecretBackendRedisStorageURLFieldName:           "storageURLValue",
			component.BackendSecretBackendRedisQueuesURLFieldName:            "queueURLValue",
			component.BackendSecretBackendRedisStorageSentinelHostsFieldName: "storageSentinelHostsValue",
			component.BackendSecretBackendRedisStorageSentinelRoleFieldName:  "storageSentinelRoleValue",
			component.BackendSecretBackendRedisQueuesSentinelHostsFieldName:  "queueSentinelHostsValue",
			component.BackendSecretBackendRedisQueuesSentinelRoleFieldName:   "queueSentinelRoleValue",
		},
		Type: v1.SecretTypeOpaque,
	}
}

func defaultBackendOoptions(opts *component.BackendOptions) *component.BackendOptions {
	return &component.BackendOptions{
		ServiceEndpoint:              component.DefaultBackendServiceEndpoint(),
		RouteEndpoint:                fmt.Sprintf("https://backend-%s.%s", tenantName, wildcardDomain),
		StorageURL:                   component.DefaultBackendRedisStorageURL(),
		QueuesURL:                    component.DefaultBackendRedisQueuesURL(),
		StorageSentinelHosts:         component.DefaultBackendStorageSentinelHosts(),
		StorageSentinelRole:          component.DefaultBackendStorageSentinelRole(),
		QueuesSentinelHosts:          component.DefaultBackendQueuesSentinelHosts(),
		QueuesSentinelRole:           component.DefaultBackendQueuesSentinelRole(),
		ListenerResourceRequirements: component.DefaultBackendListenerResourceRequirements(),
		WorkerResourceRequirements:   component.DefaultBackendWorkerResourceRequirements(),
		CronResourceRequirements:     component.DefaultCronResourceRequirements(),
		ListenerReplicas:             int32(listenerReplicaCount),
		WorkerReplicas:               int32(workerReplicaCount),
		CronReplicas:                 int32(cronReplicaCount),
		AppLabel:                     appLabel,
		SystemBackendUsername:        component.DefaultSystemBackendUsername(),
		SystemBackendPassword:        opts.SystemBackendPassword,
		TenantName:                   tenantName,
		WildcardDomain:               wildcardDomain,
	}
}

func TestGetBackendOptions(t *testing.T) {
	tmpTrueValue := trueValue
	tmpAppLabel := appLabel
	tmpTenantName := tenantName
	tmpListenerReplicaCount := listenerReplicaCount
	tmpWorkerReplicaCount := workerReplicaCount
	tmpCronReplicaCount := cronReplicaCount

	cases := []struct {
		testName                    string
		resourceRequirementsEnabled bool
		internalSecret              *v1.Secret
		listenerSecret              *v1.Secret
		redisSecret                 *v1.Secret
		expectedOptionsFactory      func(*component.BackendOptions) *component.BackendOptions
	}{
		{"Default", true, nil, nil, nil,
			func(opts *component.BackendOptions) *component.BackendOptions {
				return defaultBackendOoptions(opts)
			},
		},
		{"WithoutResourceRequirements", false, nil, nil, nil,
			func(in *component.BackendOptions) *component.BackendOptions {
				opts := defaultBackendOoptions(in)
				opts.ListenerResourceRequirements = v1.ResourceRequirements{}
				opts.WorkerResourceRequirements = v1.ResourceRequirements{}
				opts.CronResourceRequirements = v1.ResourceRequirements{}
				return opts
			},
		},
		{"InternalSecret", true, getInternalSecret(namespace), nil, nil,
			func(in *component.BackendOptions) *component.BackendOptions {
				opts := defaultBackendOoptions(in)
				opts.SystemBackendUsername = "someUserName"
				opts.SystemBackendPassword = "somePasswd"
				return opts
			},
		},
		{"ListenerSecret", true, nil, getListenerSecret(namespace), nil,
			func(in *component.BackendOptions) *component.BackendOptions {
				opts := defaultBackendOoptions(in)
				opts.ServiceEndpoint = "serviceValue"
				opts.RouteEndpoint = "routeValue"
				return opts
			},
		},
		{"RedisSecret", true, nil, nil, getRedisSecret(namespace),
			func(in *component.BackendOptions) *component.BackendOptions {
				opts := defaultBackendOoptions(in)
				opts.StorageURL = "storageURLValue"
				opts.QueuesURL = "queueURLValue"
				opts.StorageSentinelHosts = "storageSentinelHostsValue"
				opts.StorageSentinelRole = "storageSentinelRoleValue"
				opts.QueuesSentinelHosts = "queueSentinelHostsValue"
				opts.QueuesSentinelRole = "queueSentinelRoleValue"
				return opts
			},
		},
	}

	for _, tc := range cases {
		t.Run(tc.testName, func(subT *testing.T) {
			apimanager := &appsv1alpha1.APIManager{
				ObjectMeta: metav1.ObjectMeta{
					Name:      apimanagerName,
					Namespace: namespace,
				},
				Spec: appsv1alpha1.APIManagerSpec{
					APIManagerCommonSpec: appsv1alpha1.APIManagerCommonSpec{
						AppLabel:                     &tmpAppLabel,
						ImageStreamTagImportInsecure: &tmpTrueValue,
						WildcardDomain:               wildcardDomain,
						TenantName:                   &tmpTenantName,
						ResourceRequirementsEnabled:  &tc.resourceRequirementsEnabled,
					},
					Backend: &appsv1alpha1.BackendSpec{
						ListenerSpec: &appsv1alpha1.BackendListenerSpec{Replicas: &tmpListenerReplicaCount},
						WorkerSpec:   &appsv1alpha1.BackendWorkerSpec{Replicas: &tmpWorkerReplicaCount},
						CronSpec:     &appsv1alpha1.BackendCronSpec{Replicas: &tmpCronReplicaCount},
					},
				},
			}
			objs := []runtime.Object{apimanager}
			if tc.internalSecret != nil {
				objs = append(objs, tc.internalSecret)
			}
			if tc.listenerSecret != nil {
				objs = append(objs, tc.listenerSecret)
			}
			if tc.redisSecret != nil {
				objs = append(objs, tc.redisSecret)
			}

			cl := fake.NewFakeClient(objs...)
			optsProvider := NewOperatorBackendOptionsProvider(&apimanager.Spec, namespace, cl)
			opts, err := optsProvider.GetBackendOptions()
			if err != nil {
				t.Error(err)
			}
			expectedOptions := tc.expectedOptionsFactory(opts)
			if !reflect.DeepEqual(expectedOptions, opts) {
				subT.Errorf("Resulting expected options differ: %s", cmp.Diff(expectedOptions, opts, cmpopts.IgnoreUnexported(resource.Quantity{})))
			}
		})
	}
}
