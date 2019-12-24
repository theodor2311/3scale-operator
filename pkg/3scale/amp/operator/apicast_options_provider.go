package operator

import (
	"strconv"

	"github.com/3scale/3scale-operator/pkg/3scale/amp/component"
	appsv1alpha1 "github.com/3scale/3scale-operator/pkg/apis/apps/v1alpha1"
	v1 "k8s.io/api/core/v1"
)

type ApicastOptionsProvider struct {
	apimanagerSpec *appsv1alpha1.APIManagerSpec
	apicastOptions *component.ApicastOptions
}

func NewApicastOptionsProvider(apimanagerSpec *appsv1alpha1.APIManagerSpec) *ApicastOptionsProvider {
	return &ApicastOptionsProvider{
		apimanagerSpec: apimanagerSpec,
		apicastOptions: component.NewApicastOptions(),
	}
}

func (a *ApicastOptionsProvider) GetApicastOptions() (*component.ApicastOptions, error) {
	a.apicastOptions.AppLabel = *a.apimanagerSpec.AppLabel
	a.apicastOptions.TenantName = *a.apimanagerSpec.TenantName
	a.apicastOptions.WildcardDomain = a.apimanagerSpec.WildcardDomain
	a.apicastOptions.ManagementAPI = *a.apimanagerSpec.Apicast.ApicastManagementAPI
	a.apicastOptions.OpenSSLVerify = strconv.FormatBool(*a.apimanagerSpec.Apicast.OpenSSLVerify)
	a.apicastOptions.ResponseCodes = strconv.FormatBool(*a.apimanagerSpec.Apicast.IncludeResponseCodes)

	a.setResourceRequirementsOptions()
	a.setReplicas()

	err := a.apicastOptions.Validate()
	return a.apicastOptions, err
}

func (a *ApicastOptionsProvider) setResourceRequirementsOptions() {
	if *a.apimanagerSpec.ResourceRequirementsEnabled {
		a.apicastOptions.ProductionResourceRequirements = component.DefaultProductionResourceRequirements()
		a.apicastOptions.StagingResourceRequirements = component.DefaultStagingResourceRequirements()
	} else {
		a.apicastOptions.ProductionResourceRequirements = v1.ResourceRequirements{}
		a.apicastOptions.StagingResourceRequirements = v1.ResourceRequirements{}
	}
}

func (a *ApicastOptionsProvider) setReplicas() {
	a.apicastOptions.ProductionReplicas = int32(*a.apimanagerSpec.Apicast.ProductionSpec.Replicas)
	a.apicastOptions.StagingReplicas = int32(*a.apimanagerSpec.Apicast.StagingSpec.Replicas)
}
