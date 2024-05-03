package kubernetes

import (
	"github.com/UpCloudLtd/provider-upcloud/config/groupversion"
	"github.com/crossplane/upjet/pkg/config"
)

// Resources is a list of all supported kubernetes resources.
var Resources = []string{
	"upcloud_kubernetes_cluster",
	"upcloud_kubernetes_node_group",
}

// Configure configures the kubernetes resources.
func Configure(p *config.Provider) {
	groupversion.Configure(Resources, p, "uks", "v1alpha1")

	p.AddResourceConfigurator("upcloud_kubernetes_cluster", func(r *config.Resource) {
		r.ExternalName = config.IdentifierFromProvider
		r.Kind = "KubernetesCluster"
		r.UseAsync = true

		r.References["network"] = config.Reference{
			Type: "github.com/UpCloudLtd/provider-upcloud/apis/network/v1alpha1.Network",
		}
	})

	p.AddResourceConfigurator("upcloud_kubernetes_node_group", func(r *config.Resource) {
		r.ExternalName = config.TemplatedStringAsIdentifier("name", "{{ .parameters.cluster }}/{{ .external_name }}")
		r.Kind = "KubernetesNodeGroup"

		r.References["cluster"] = config.Reference{
			Type: "KubernetesCluster",
		}
	})
}
