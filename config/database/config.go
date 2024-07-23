package database

import (
	"github.com/UpCloudLtd/provider-upcloud/config/groupversion"

	"github.com/crossplane/upjet/pkg/config"
)

// SDKResources is a list of all supported database resources implemented with Terraform legacy SDKv2.
var SDKResources = []string{
	"upcloud_managed_database_postgresql",
	"upcloud_managed_database_mysql",
	"upcloud_managed_database_opensearch",
	"upcloud_managed_database_redis",
	"upcloud_managed_database_logical_database",
	"upcloud_managed_database_user",
}

// PluginFrameworkResources is a list of all supported database resources implemented with Terraform Plugin Framework.
var PluginFrameworkResources = []string{}

// AllResources is a list of all supported database resources.
var AllResources = append(SDKResources, PluginFrameworkResources...)

// Configure configures the database resources.
func Configure(p *config.Provider) {
	groupversion.Configure(AllResources, p, "database", "v1alpha1")

	p.AddResourceConfigurator("upcloud_managed_database_postgresql", func(r *config.Resource) {
		r.ExternalName = config.IdentifierFromProvider
		r.Kind = "ManagedDatabasePostgresql"
		r.UseAsync = true

		r.References["network.uuid"] = config.Reference{
			TerraformName: "upcloud_network",
		}

		if s, ok := r.TerraformResource.Schema["node_states"]; ok {
			s.Optional = true
			s.Computed = true
		}
	})

	p.AddResourceConfigurator("upcloud_managed_database_mysql", func(r *config.Resource) {
		r.ExternalName = config.IdentifierFromProvider
		r.Kind = "ManagedDatabaseMysql"
		r.UseAsync = true

		r.References["network.uuid"] = config.Reference{
			TerraformName: "upcloud_network",
		}

		if s, ok := r.TerraformResource.Schema["node_states"]; ok {
			s.Optional = true
			s.Computed = true
		}
	})

	p.AddResourceConfigurator("upcloud_managed_database_opensearch", func(r *config.Resource) {
		r.ExternalName = config.IdentifierFromProvider
		r.Kind = "ManagedDatabaseOpensearch"
		r.UseAsync = true

		r.References["network.uuid"] = config.Reference{
			TerraformName: "upcloud_network",
		}

		if s, ok := r.TerraformResource.Schema["node_states"]; ok {
			s.Optional = true
			s.Computed = true
		}
	})

	p.AddResourceConfigurator("upcloud_managed_database_redis", func(r *config.Resource) {
		r.ExternalName = config.IdentifierFromProvider
		r.Kind = "ManagedDatabaseRedis"
		r.UseAsync = true

		r.References["network.uuid"] = config.Reference{
			TerraformName: "upcloud_network",
		}

		if s, ok := r.TerraformResource.Schema["node_states"]; ok {
			s.Optional = true
			s.Computed = true
		}
	})

	p.AddResourceConfigurator("upcloud_managed_database_logical_database", func(r *config.Resource) {
		r.ExternalName = config.TemplatedStringAsIdentifier("name", "{{ .parameters.service }}/{{ .external_name }}")
		r.Kind = "ManagedDatabaseLogicalDatabase"

		// TODO: use generic reference type if that's ever implemented in upjet
		// https://github.com/crossplane/upjet/issues/95
		r.References["service"] = config.Reference{
			TerraformName: "upcloud_managed_database_postgresql",
		}

		if schema, ok := r.TerraformResource.Schema["service"]; ok {
			schema.Description = "The service to which the logical database belongs. Please note that reference fields (`serviceRef` and `serviceSelector`) only work for PostgreSQL databases. For other databases you need to leverage compositions and patches to pass database service ID to logical database `service` field. See https://docs.crossplane.io/latest/concepts/patch-and-transform/#patching-between-resources for more info."
		}
	})

	p.AddResourceConfigurator("upcloud_managed_database_user", func(r *config.Resource) {
		r.ExternalName = config.TemplatedStringAsIdentifier("username", "{{ .parameters.service }}/{{ .external_name }}")
		r.Kind = "ManagedDatabaseUser"

		r.References["service"] = config.Reference{
			TerraformName: "upcloud_managed_database_postgresql",
		}

		if schema, ok := r.TerraformResource.Schema["service"]; ok {
			schema.Description = "The service to which the logical database belongs. Please note that reference fields (`serviceRef` and `serviceSelector`) only work for PostgreSQL databases. For other databases you need to leverage compositions and patches to pass database service ID to database user `service` field. See https://docs.crossplane.io/latest/concepts/patch-and-transform/#patching-between-resources for more info."
		}
	})
}
