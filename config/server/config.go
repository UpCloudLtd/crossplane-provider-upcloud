package server

import (
	"github.com/UpCloudLtd/crossplane-provider-upcloud/config/groupversion"

	"github.com/crossplane/upjet/v2/pkg/config"
)

// SDKResources is a list of all supported server resources implemented with Terraform legacy SDKv2.
var SDKResources = []string{}

// PluginFrameworkResources is a list of all supported server resources implemented with Terraform Plugin Framework.
var PluginFrameworkResources = []string{
	"upcloud_server",
	"upcloud_firewall_rules",
	"upcloud_server_group",
}

// AllResources is a list of all supported server resources.
var AllResources = append(SDKResources, PluginFrameworkResources...)

// Configure configures the server resources.
func Configure(p *config.Provider) {
	groupversion.Configure(AllResources, p, "server", "v1alpha1")

	p.AddResourceConfigurator("upcloud_server", func(r *config.Resource) {
		r.ExternalName = config.IdentifierFromProvider
		r.UseAsync = true

		// This is due to conflict between the cpu/mem and plan fields
		// Remove once the state setting is fixed
		r.LateInitializer = config.LateInitializer{
			IgnoredFields: []string{"cpu", "mem", "plan"},
		}

		r.References["network_interface.network"] = config.Reference{
			TerraformName: "upcloud_network",
		}

		r.References["storage_devices.storage"] = config.Reference{
			TerraformName: "upcloud_storage",
		}

		// TODO use r.SchemaElementOptions.SetEmbeddedObject("template") to flatten the yaml manifest from list to single embedded object
		// see https://github.com/crossplane/upjet/v2/blob/v1.2.4/pkg/config/resource.go#L550 for more info
	})

	p.AddResourceConfigurator("upcloud_server_group", func(r *config.Resource) {
		r.ExternalName = config.IdentifierFromProvider
		r.Kind = "ServerGroup"

		r.References["members"] = config.Reference{
			TerraformName: "upcloud_server",
		}
	})

	p.AddResourceConfigurator("upcloud_firewall_rules", func(r *config.Resource) {
		r.ExternalName = config.IdentifierFromProvider
		r.Kind = "FirewallRules"

		r.References["server_id"] = config.Reference{
			TerraformName: "upcloud_server",
		}
	})
}
