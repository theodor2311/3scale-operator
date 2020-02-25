package operator

import (
	"testing"

	"github.com/3scale/3scale-operator/pkg/3scale/amp/component"
	appsv1alpha1 "github.com/3scale/3scale-operator/pkg/apis/apps/v1alpha1"
)

func TestGetAmpImagesOptions(t *testing.T) {
	appLabel := "someLabel"
	apicastImage := "quay.io/3scale/apicast:mytag"
	backendImage := "quay.io/3scale/backend:mytag"
	systemImage := "quay.io/3scale/backend:mytag"
	zyncImage := "quay.io/3scale/zync:mytag"
	zyncPostgresqlImage := "postgresql-10:mytag"
	systemMemcachedImage := "memcached:mytag"
	trueValue := true

	cases := []struct {
		name       string
		apimanager *appsv1alpha1.APIManagerSpec
		testFunc   func(*testing.T, *component.AmpImagesOptions)
	}{
		{
			"apicastImage", &appsv1alpha1.APIManagerSpec{
				APIManagerCommonSpec: appsv1alpha1.APIManagerCommonSpec{
					AppLabel:                     &appLabel,
					ImageStreamTagImportInsecure: &trueValue,
				},
				Apicast: &appsv1alpha1.ApicastSpec{Image: &apicastImage},
			},
			func(subT *testing.T, opts *component.AmpImagesOptions) {
				if opts.ApicastImage != apicastImage {
					subT.Errorf("got: %s, expected: %s", opts.ApicastImage, apicastImage)
				}
			},
		},
		{
			"backendImage", &appsv1alpha1.APIManagerSpec{
				APIManagerCommonSpec: appsv1alpha1.APIManagerCommonSpec{
					AppLabel:                     &appLabel,
					ImageStreamTagImportInsecure: &trueValue,
				},
				Backend: &appsv1alpha1.BackendSpec{Image: &backendImage},
			},
			func(subT *testing.T, opts *component.AmpImagesOptions) {
				if opts.BackendImage != backendImage {
					subT.Errorf("got: %s, expected: %s", opts.BackendImage, backendImage)
				}
			},
		},
		{
			"systemImage", &appsv1alpha1.APIManagerSpec{
				APIManagerCommonSpec: appsv1alpha1.APIManagerCommonSpec{
					AppLabel:                     &appLabel,
					ImageStreamTagImportInsecure: &trueValue,
				},
				System: &appsv1alpha1.SystemSpec{Image: &systemImage},
			},
			func(subT *testing.T, opts *component.AmpImagesOptions) {
				if opts.SystemImage != systemImage {
					subT.Errorf("got: %s, expected: %s", opts.SystemImage, systemImage)
				}
			},
		},
		{
			"zyncImage", &appsv1alpha1.APIManagerSpec{
				APIManagerCommonSpec: appsv1alpha1.APIManagerCommonSpec{
					AppLabel:                     &appLabel,
					ImageStreamTagImportInsecure: &trueValue,
				},
				Zync: &appsv1alpha1.ZyncSpec{Image: &zyncImage},
			},
			func(subT *testing.T, opts *component.AmpImagesOptions) {
				if opts.ZyncImage != zyncImage {
					subT.Errorf("got: %s, expected: %s", opts.ZyncImage, zyncImage)
				}
			},
		},
		{
			"zyncPostgresqlImage", &appsv1alpha1.APIManagerSpec{
				APIManagerCommonSpec: appsv1alpha1.APIManagerCommonSpec{
					AppLabel:                     &appLabel,
					ImageStreamTagImportInsecure: &trueValue,
				},
				Zync: &appsv1alpha1.ZyncSpec{PostgreSQLImage: &zyncPostgresqlImage},
			},
			func(subT *testing.T, opts *component.AmpImagesOptions) {
				if opts.ZyncDatabasePostgreSQLImage != zyncPostgresqlImage {
					subT.Errorf("got: %s, expected: %s", opts.ZyncDatabasePostgreSQLImage, zyncPostgresqlImage)
				}
			},
		},
		{
			"systemMemcachedImage", &appsv1alpha1.APIManagerSpec{
				APIManagerCommonSpec: appsv1alpha1.APIManagerCommonSpec{
					AppLabel:                     &appLabel,
					ImageStreamTagImportInsecure: &trueValue,
				},
				System: &appsv1alpha1.SystemSpec{MemcachedImage: &systemMemcachedImage},
			},
			func(subT *testing.T, opts *component.AmpImagesOptions) {
				if opts.SystemMemcachedImage != systemMemcachedImage {
					subT.Errorf("got: %s, expected: %s", opts.SystemMemcachedImage, systemMemcachedImage)
				}
			},
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(subT *testing.T) {
			optsProvider := NewAmpImagesOptionsProvider(tc.apimanager)
			opts, err := optsProvider.GetAmpImagesOptions()
			if err != nil {
				subT.Error(err)
			}
			tc.testFunc(subT, opts)
			if opts.AppLabel != appLabel {
				subT.Errorf("AppLabel does not match, got: %s, expected: %s", opts.AppLabel, appLabel)
			}
		})
	}
}
