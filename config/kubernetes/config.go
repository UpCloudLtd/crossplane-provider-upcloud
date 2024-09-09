package kubernetes

import (
	"github.com/UpCloudLtd/provider-upcloud/config/groupversion"

	"github.com/crossplane/upjet/pkg/config"
)

// SDKResources is a list of all supported kubernetes resources implemented with Terraform legacy SDKv2.
var SDKResources = []string{}

// PluginFrameworkResources is a list of all supported network resources implemented with Terraform Plugin Framework.
var PluginFrameworkResources = []string{
	"upcloud_kubernetes_cluster",
	"upcloud_kubernetes_node_group",
}

// AllResources is a list of all supported network resources.
var AllResources = append(SDKResources, PluginFrameworkResources...)

// Configure configures the kubernetes resources.
func Configure(p *config.Provider) {
	groupversion.Configure(AllResources, p, "uks", "v1alpha1")

	p.AddResourceConfigurator("upcloud_kubernetes_cluster", func(r *config.Resource) {
		r.ExternalName = config.IdentifierFromProvider
		r.Kind = "KubernetesCluster"
		r.UseAsync = true

		r.References["network"] = config.Reference{
			TerraformName: "upcloud_network",
		}

		if s, ok := r.TerraformResource.Schema["labels"]; ok {
			s.Optional = false
			s.Computed = false
			s.Required = true
		}
	})

	p.AddResourceConfigurator("upcloud_kubernetes_node_group", func(r *config.Resource) {
		r.ExternalName = config.TemplatedStringAsIdentifier("name", "{{ .parameters.cluster }}/{{ .external_name }}")
		r.Kind = "KubernetesNodeGroup"

		r.References["cluster"] = config.Reference{
			TerraformName: "upcloud_kubernetes_cluster",
		}
	})
}
