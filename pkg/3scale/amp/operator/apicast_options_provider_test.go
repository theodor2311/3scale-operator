package operator

import (
	"reflect"
	"testing"

	"github.com/3scale/3scale-operator/pkg/3scale/amp/component"
	appsv1alpha1 "github.com/3scale/3scale-operator/pkg/apis/apps/v1alpha1"
	v1 "k8s.io/api/core/v1"
)

func TestGetApicastOptions(t *testing.T) {
	wildcardDomain := "test.3scale.net"
	appLabel := "someLabel"
	tenantName := "someTenant"
	apicastManagementAPI := "disabled"
	trueValue := true
	var oneValue int64 = 1

	cases := []struct {
		name                        string
		resourceRequirementsEnabled bool
		testFunc                    func(*testing.T, *component.ApicastOptions)
	}{
		{"WithResourceRequirements", true,
			func(subT *testing.T, opts *component.ApicastOptions) {
				if !reflect.DeepEqual(opts.ProductionResourceRequirements, component.DefaultProductionResourceRequirements()) {
					subT.Error("production resource requirements do not match")
				}
				if !reflect.DeepEqual(opts.StagingResourceRequirements, component.DefaultStagingResourceRequirements()) {
					subT.Error("production resource requirements do not match")
				}
			},
		},
		{"WithoutResourceRequirements", false,
			func(subT *testing.T, opts *component.ApicastOptions) {
				if !reflect.DeepEqual(opts.ProductionResourceRequirements, v1.ResourceRequirements{}) {
					subT.Error("production resource requirements do not match")
				}
				if !reflect.DeepEqual(opts.StagingResourceRequirements, v1.ResourceRequirements{}) {
					subT.Error("production resource requirements do not match")
				}
			},
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(subT *testing.T) {
			resourceRequirementsEnabled := tc.resourceRequirementsEnabled
			apimanager := &appsv1alpha1.APIManagerSpec{
				APIManagerCommonSpec: appsv1alpha1.APIManagerCommonSpec{
					WildcardDomain:               wildcardDomain,
					AppLabel:                     &appLabel,
					ImageStreamTagImportInsecure: &trueValue,
					TenantName:                   &tenantName,
					ResourceRequirementsEnabled:  &resourceRequirementsEnabled,
				},
				Apicast: &appsv1alpha1.ApicastSpec{
					ApicastManagementAPI: &apicastManagementAPI,
					OpenSSLVerify:        &trueValue,
					IncludeResponseCodes: &trueValue,
					StagingSpec: &appsv1alpha1.ApicastStagingSpec{
						Replicas: &oneValue,
					},
					ProductionSpec: &appsv1alpha1.ApicastProductionSpec{
						Replicas: &oneValue,
					},
				},
			}
			optsProvider := NewApicastOptionsProvider(apimanager)
			opts, err := optsProvider.GetApicastOptions()
			if err != nil {
				subT.Error(err)
			}
			tc.testFunc(subT, opts)
			if opts.AppLabel != appLabel {
				subT.Errorf("got: %s, expected: %s", opts.AppLabel, appLabel)
			}
			if opts.TenantName != tenantName {
				subT.Errorf("got: %s, expected: %s", opts.TenantName, tenantName)
			}
			if opts.WildcardDomain != wildcardDomain {
				subT.Errorf("got: %s, expected: %s", opts.WildcardDomain, wildcardDomain)
			}
			if opts.ManagementAPI != apicastManagementAPI {
				subT.Errorf("got: %s, expected: %s", opts.ManagementAPI, apicastManagementAPI)
			}
			if opts.ProductionReplicas != 1 {
				subT.Errorf("got: %d, expected: 1", opts.ProductionReplicas)
			}
			if opts.StagingReplicas != 1 {
				subT.Errorf("got: %d, expected: 1", opts.StagingReplicas)
			}
		})
	}

}
