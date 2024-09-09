/*
Copyright 2021 Upbound Inc.
*/

package clients

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/crossplane/crossplane-runtime/pkg/resource"
	"github.com/crossplane/upjet/pkg/config"
	"github.com/crossplane/upjet/pkg/terraform"
	"github.com/pkg/errors"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"

	terraformframework "github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
	terraformsdk "github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	"github.com/UpCloudLtd/provider-upcloud/apis/v1beta1"
)

const (
	// error messages
	errNoProviderConfig     = "no providerConfigRef provided"
	errGetProviderConfig    = "cannot get referenced ProviderConfig"
	errTrackUsage           = "cannot track ProviderConfig usage"
	errExtractCredentials   = "cannot extract credentials"
	errUnmarshalCredentials = "cannot unmarshal upcloud credentials as JSON"
)

// TerraformSetupBuilder builds Terraform a terraform.SetupFn function which
// returns Terraform provider setup configuration
func TerraformSetupBuilder(version, providerSource, providerVersion string, provider *config.Provider) terraform.SetupFn {
	return func(ctx context.Context, client client.Client, mg resource.Managed) (terraform.Setup, error) {
		ps := terraform.Setup{
			Version: version,
			Requirement: terraform.ProviderRequirement{
				Source:  providerSource,
				Version: providerVersion,
			},
			Scheduler: terraform.NewNoOpProviderScheduler(),
		}

		configRef := mg.GetProviderConfigReference()
		if configRef == nil {
			return ps, errors.New(errNoProviderConfig)
		}
		pc := &v1beta1.ProviderConfig{}
		if err := client.Get(ctx, types.NamespacedName{Name: configRef.Name}, pc); err != nil {
			return ps, errors.Wrap(err, errGetProviderConfig)
		}

		t := resource.NewProviderConfigUsageTracker(client, &v1beta1.ProviderConfigUsage{})
		if err := t.Track(ctx, mg); err != nil {
			return ps, errors.Wrap(err, errTrackUsage)
		}

		data, err := resource.CommonCredentialExtractor(ctx, pc.Spec.Credentials.Source, client, pc.Spec.Credentials.CommonCredentialSelectors)
		if err != nil {
			return ps, errors.Wrap(err, errExtractCredentials)
		}
		creds := map[string]string{}
		if err := json.Unmarshal(data, &creds); err != nil {
			return ps, errors.Wrap(err, errUnmarshalCredentials)
		}

		// Set credentials in Terraform provider configuration.
		ps.Configuration = map[string]any{
			"username": creds["username"],
			"password": creds["password"],
		}

		// This configures the SDKv2 provider so that it knows how to get credentials and such
		diag := provider.TerraformProvider.Configure(ctx, &terraformsdk.ResourceConfig{
			Config: ps.Configuration,
		})

		if diag != nil && diag.HasError() {
			return ps, fmt.Errorf("failed to configure provider: %v", diag)
		}

		ps.Meta = provider.TerraformProvider.Meta()

		// This configures the Plugin Framework provider
		schemaResponse := &terraformframework.SchemaResponse{}
		provider.TerraformPluginFrameworkProvider.Schema(ctx, terraformframework.SchemaRequest{}, schemaResponse)

		provider.TerraformPluginFrameworkProvider.Configure(
			ctx,
			terraformframework.ConfigureRequest{
				TerraformVersion: version,
				Config: tfsdk.Config{
					Raw: tftypes.NewValue(
						tftypes.Map{ElementType: tftypes.String},
						map[string]tftypes.Value{
							"username": tftypes.NewValue(tftypes.String, creds["username"]),
							"password": tftypes.NewValue(tftypes.String, creds["password"]),
						}),
					Schema: schemaResponse.Schema,
				},
			},
			&terraformframework.ConfigureResponse{},
		)

		ps.FrameworkProvider = provider.TerraformPluginFrameworkProvider

		return ps, nil
	}
}
