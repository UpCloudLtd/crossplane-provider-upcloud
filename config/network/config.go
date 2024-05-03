package network

import (
	"github.com/UpCloudLtd/provider-upcloud/config/groupversion"
	"github.com/crossplane/upjet/pkg/config"
)

// Resources is a list of all supported network resources.
var Resources = []string{
	"upcloud_network",
	"upcloud_router",
}

// Configure configures the network resources.
func Configure(p *config.Provider) {
	groupversion.Configure(Resources, p, "network", "v1alpha1")

	p.AddResourceConfigurator("upcloud_network", func(r *config.Resource) {
		r.ExternalName = config.IdentifierFromProvider

		r.References["router"] = config.Reference{
			Type: "Router",
		}
	})

	p.AddResourceConfigurator("upcloud_router", func(r *config.Resource) {
		r.ExternalName = config.IdentifierFromProvider
	})
}
