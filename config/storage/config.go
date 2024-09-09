package storage

import (
	"github.com/UpCloudLtd/provider-upcloud/config/groupversion"

	"github.com/crossplane/upjet/pkg/config"
)

// SDKResources is a list of all supported storage resources implemented with Terraform legacy SDKv2.
var SDKResources = []string{}

// PluginFrameworkResources is a list of all supported storage resources implemented with Terraform Plugin Framework.
var PluginFrameworkResources = []string{
	"upcloud_storage",
}

// AllResources is a list of all supported storage resources.
var AllResources = append(SDKResources, PluginFrameworkResources...)

// Configure configures the storage resources.
func Configure(p *config.Provider) {
	groupversion.Configure(AllResources, p, "storage", "v1alpha1")

	p.AddResourceConfigurator("upcloud_storage", func(r *config.Resource) {
		r.ExternalName = config.IdentifierFromProvider
		r.UseAsync = true
	})
}
