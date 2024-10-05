package network

import (
	"github.com/UpCloudLtd/crossplane-provider-upcloud/config/groupversion"

	"github.com/crossplane/upjet/pkg/config"
)

// SDKResources is a list of all supported network resources implemented with Terraform legacy SDKv2.
var SDKResources = []string{}

// PluginFrameworkResources is a list of all supported network resources implemented with Terraform Plugin Framework.
var PluginFrameworkResources = []string{
	"upcloud_network",
	"upcloud_router",
}

// AllResources is a list of all supported network resources
var AllResources = append(SDKResources, PluginFrameworkResources...)

// Configure configures the network resources.
func Configure(p *config.Provider) {
	groupversion.Configure(AllResources, p, "network", "v1alpha1")

	p.AddResourceConfigurator("upcloud_network", func(r *config.Resource) {
		r.ExternalName = config.IdentifierFromProvider

		r.References["router"] = config.Reference{
			TerraformName: "upcloud_router",
		}
		if s, ok := r.TerraformResource.Schema["ip_networks"]; ok {
			s.Required = true
		}
	})

	p.AddResourceConfigurator("upcloud_router", func(r *config.Resource) {
		r.ExternalName = config.IdentifierFromProvider

		if s, ok := r.TerraformResource.Schema["labels"]; ok {
			s.Optional = false
			s.Computed = false
			s.Required = true
		}

		if s, ok := r.TerraformResource.Schema["static_route"]; ok {
			s.Optional = false
			s.Computed = false
			s.Required = true
		}
	})
}
