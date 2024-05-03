package groupversion

import "github.com/crossplane/upjet/pkg/config"

// Configure sets the short group and version for each resource in the group
func Configure(resources []string, p *config.Provider, shortName string, version string) {
	for _, res := range resources {
		p.AddResourceConfigurator(res, func(r *config.Resource) {
			r.ShortGroup = shortName
			r.Version = version
		})
	}
}
