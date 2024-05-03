package storage

import (
	"github.com/UpCloudLtd/provider-upcloud/config/groupversion"
	"github.com/crossplane/upjet/pkg/config"
)

// Resources is a list of all supported storage resources.
var Resources = []string{
	"upcloud_storage",
}

// Configure configures the storage resources.
func Configure(p *config.Provider) {
	groupversion.Configure(Resources, p, "storage", "v1alpha1")

	p.AddResourceConfigurator("upcloud_storage", func(r *config.Resource) {
		r.ExternalName = config.IdentifierFromProvider
		r.UseAsync = true
	})
}
