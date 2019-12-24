package operator

import (
	"github.com/3scale/3scale-operator/pkg/3scale/amp/component"
	"github.com/3scale/3scale-operator/pkg/3scale/amp/product"
	appsv1alpha1 "github.com/3scale/3scale-operator/pkg/apis/apps/v1alpha1"
)

type AmpImagesOptionsProvider struct {
	apimanagerSpec   *appsv1alpha1.APIManagerSpec
	ampImagesOptions *component.AmpImagesOptions
}

func NewAmpImagesOptionsProvider(apimanagerSpec *appsv1alpha1.APIManagerSpec) *AmpImagesOptionsProvider {
	return &AmpImagesOptionsProvider{
		apimanagerSpec:   apimanagerSpec,
		ampImagesOptions: component.NewAmpImagesOptions(),
	}
}

func (a *AmpImagesOptionsProvider) GetAmpImagesOptions() (*component.AmpImagesOptions, error) {
	a.ampImagesOptions.AppLabel = *a.apimanagerSpec.AppLabel
	a.ampImagesOptions.AmpRelease = product.ThreescaleRelease
	a.ampImagesOptions.InsecureImportPolicy = *a.apimanagerSpec.ImageStreamTagImportInsecure

	a.ampImagesOptions.ApicastImage = ApicastImageURL()
	if a.apimanagerSpec.Apicast != nil && a.apimanagerSpec.Apicast.Image != nil {
		a.ampImagesOptions.ApicastImage = *a.apimanagerSpec.Apicast.Image
	}

	a.ampImagesOptions.BackendImage = BackendImageURL()
	if a.apimanagerSpec.Backend != nil && a.apimanagerSpec.Backend.Image != nil {
		a.ampImagesOptions.BackendImage = *a.apimanagerSpec.Backend.Image
	}

	a.ampImagesOptions.SystemImage = SystemImageURL()
	if a.apimanagerSpec.System != nil && a.apimanagerSpec.System.Image != nil {
		a.ampImagesOptions.SystemImage = *a.apimanagerSpec.System.Image
	}

	a.ampImagesOptions.ZyncImage = ZyncImageURL()
	if a.apimanagerSpec.Zync != nil && a.apimanagerSpec.Zync.Image != nil {
		a.ampImagesOptions.ZyncImage = *a.apimanagerSpec.Zync.Image
	}

	a.ampImagesOptions.ZyncDatabasePostgreSQLImage = component.ZyncPostgreSQLImageURL()
	if a.apimanagerSpec.Zync != nil && a.apimanagerSpec.Zync.PostgreSQLImage != nil {
		a.ampImagesOptions.ZyncDatabasePostgreSQLImage = *a.apimanagerSpec.Zync.PostgreSQLImage
	}

	a.ampImagesOptions.SystemMemcachedImage = SystemMemcachedImageURL()
	if a.apimanagerSpec.System != nil && a.apimanagerSpec.System.MemcachedImage != nil {
		a.ampImagesOptions.SystemMemcachedImage = *a.apimanagerSpec.System.MemcachedImage
	}

	err := a.ampImagesOptions.Validate()
	return a.ampImagesOptions, err
}
