package network

import (
	"slices"

	"github.com/UpCloudLtd/provider-upcloud/config/groupversion"
	"github.com/crossplane/upjet/pkg/config"
)

// SDKResources is a list of all supported network resources implemented with Terraform legacy SDKv2.
var SDKResources = []string{
	"upcloud_router",
}

// PluginFrameworkResources is a list of all supported network resources implemented with Terraform Plugin Framework.
var PluginFrameworkResources = []string{
	"upcloud_network",
}

// AllResources is a list of all supported network resources.
var AllResources = slices.Concat(SDKResources, PluginFrameworkResources)

// Configure configures the network resources.
func Configure(p *config.Provider) {
	groupversion.Configure(AllResources, p, "network", "v1alpha1")

	p.AddResourceConfigurator("upcloud_network", func(r *config.Resource) {
		r.ExternalName = config.IdentifierFromProvider

		// r.References["router"] = config.Reference{
		// 	Type: "Router",
		// }
	})

	p.AddResourceConfigurator("upcloud_router", func(r *config.Resource) {
		r.ExternalName = config.IdentifierFromProvider
	})
}
