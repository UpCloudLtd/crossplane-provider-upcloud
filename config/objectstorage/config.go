package objectstorage

import (
	"context"
	"fmt"

	"github.com/UpCloudLtd/crossplane-provider-upcloud/config/groupversion"

	"github.com/crossplane/upjet/v2/pkg/config"
)

// SDKResources is a list of all supported object storage resources implemented with Terraform legacy SDKv2.
var SDKResources = []string{}

// PluginFrameworkResources is a list of all supported object storage resources implemented with Terraform Plugin Framework.
var PluginFrameworkResources = []string{
	"upcloud_managed_object_storage",
	"upcloud_managed_object_storage_policy",
	"upcloud_managed_object_storage_user",
	"upcloud_managed_object_storage_user_access_key",
	"upcloud_managed_object_storage_user_policy",
}

// AllResources is a list of all supported object storage resources.
var AllResources = append(SDKResources, PluginFrameworkResources...)

// Configure configures the object storage resources.
func Configure(p *config.Provider) {
	groupversion.Configure(AllResources, p, "objectstorage", "v1alpha1")

	p.AddResourceConfigurator("upcloud_managed_object_storage", func(r *config.Resource) {
		r.Kind = "ManagedObjectStorage"
		r.ExternalName = config.IdentifierFromProvider
		r.UseAsync = true
	})

	p.AddResourceConfigurator("upcloud_managed_object_storage_policy", func(r *config.Resource) {
		r.Kind = "ManagedObjectStoragePolicy"
		r.ExternalName = config.TemplatedStringAsIdentifier("name", "{{ .parameters.service_uuid }}/{{ .external_name }}")

		r.References["service_uuid"] = config.Reference{
			TerraformName: "upcloud_managed_object_storage",
		}
	})

	p.AddResourceConfigurator("upcloud_managed_object_storage_user", func(r *config.Resource) {
		r.Kind = "ManagedObjectStorageUser"
		r.ExternalName = config.TemplatedStringAsIdentifier("username", "{{ .parameters.service_uuid }}/{{ .external_name }}")

		r.References["service_uuid"] = config.Reference{
			TerraformName: "upcloud_managed_object_storage",
		}
	})

	p.AddResourceConfigurator("upcloud_managed_object_storage_user_access_key", func(r *config.Resource) {
		r.Kind = "ManagedObjectStorageUserAccessKey"
		r.ExternalName = config.TemplatedStringAsIdentifier("access_key_id", "{{ .parameters.service_uuid }}/{{ .parameters.username }}/{{ .external_name }}")

		// This is needed because otherwise upjet will try to set the k8s object name to `access_key_id` argument.
		// `access_key_id` is a field that comes directly from the API, so we don't want to override in any way
		r.ExternalName.SetIdentifierArgumentFn = func(_ map[string]any, _ string) {}
		r.ExternalName.OmittedFields = []string{}

		r.References["service_uuid"] = config.Reference{
			TerraformName: "upcloud_managed_object_storage",
		}

		r.References["username"] = config.Reference{
			TerraformName: "upcloud_managed_object_storage_user",
		}
	})

	p.AddResourceConfigurator("upcloud_managed_object_storage_user_policy", func(r *config.Resource) {
		r.Kind = "ManagedObjectStorageUserPolicy"

		// We set external name by hand because the only field that could be external name is at the same time a reference property
		r.ExternalName = config.ExternalName{
			SetIdentifierArgumentFn: func(_ map[string]any, _ string) {}, // no need to do anything here
			GetIDFn: func(_ context.Context, _ string, parameters map[string]any, _ map[string]any) (string, error) {
				serviceUUID, ok := parameters["service_uuid"].(string)
				if !ok {
					return "", fmt.Errorf("service_uuid is not a string or is empty")
				}

				username, ok := parameters["username"].(string)
				if !ok {
					return "", fmt.Errorf("username is not a string or is empty")
				}

				name, ok := parameters["name"].(string)
				if !ok {
					return "", fmt.Errorf("name is not a string or is empty")
				}

				return fmt.Sprintf("%s/%s/%s", serviceUUID, username, name), nil
			},
			GetExternalNameFn: func(tfstate map[string]any) (string, error) {
				name, ok := tfstate["name"].(string)
				if !ok {
					return "", fmt.Errorf("name is not a string or is empty")
				}

				return name, nil
			},
		}

		r.References["service_uuid"] = config.Reference{
			TerraformName: "upcloud_managed_object_storage",
		}

		r.References["username"] = config.Reference{
			TerraformName: "upcloud_managed_object_storage_user",
		}

		r.References["name"] = config.Reference{
			TerraformName: "upcloud_managed_object_storage_policy",
		}
	})
}
