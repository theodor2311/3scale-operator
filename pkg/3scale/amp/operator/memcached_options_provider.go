package operator

import (
	"github.com/3scale/3scale-operator/pkg/3scale/amp/component"
	appsv1alpha1 "github.com/3scale/3scale-operator/pkg/apis/apps/v1alpha1"
	v1 "k8s.io/api/core/v1"
)

type MemcachedOptionsProvider struct {
	apimanagerSpec   *appsv1alpha1.APIManagerSpec
	memcachedOptions *component.MemcachedOptions
}

func NewMemcachedOptionsProvider(apimanagerSpec *appsv1alpha1.APIManagerSpec) *MemcachedOptionsProvider {
	return &MemcachedOptionsProvider{
		apimanagerSpec:   apimanagerSpec,
		memcachedOptions: component.NewMemcachedOptions(),
	}
}

func (m *MemcachedOptionsProvider) GetMemcachedOptions() (*component.MemcachedOptions, error) {
	m.memcachedOptions.AppLabel = *m.apimanagerSpec.AppLabel

	m.setResourceRequirementsOptions()

	err := m.memcachedOptions.Validate()
	return m.memcachedOptions, err
}

func (m *MemcachedOptionsProvider) setResourceRequirementsOptions() {
	if *m.apimanagerSpec.ResourceRequirementsEnabled {
		m.memcachedOptions.ResourceRequirements = component.DefaultMemcachedResourceRequirements()
	} else {
		m.memcachedOptions.ResourceRequirements = v1.ResourceRequirements{}
	}
}
