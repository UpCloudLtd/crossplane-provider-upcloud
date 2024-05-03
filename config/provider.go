/*
Copyright 2021 Upbound Inc.
*/

package config

import (
	// Note(turkenh): we are importing this to embed provider schema document
	_ "embed"
	"slices"

	ujconfig "github.com/crossplane/upjet/pkg/config"

	"github.com/UpCloudLtd/provider-upcloud/config/database"
	"github.com/UpCloudLtd/provider-upcloud/config/kubernetes"
	"github.com/UpCloudLtd/provider-upcloud/config/network"
	"github.com/UpCloudLtd/provider-upcloud/config/objectstorage"
	"github.com/UpCloudLtd/provider-upcloud/config/server"
	"github.com/UpCloudLtd/provider-upcloud/config/storage"
)

const (
	resourcePrefix = "upcloud"
	modulePath     = "github.com/UpCloudLtd/provider-upcloud"
)

//go:embed schema.json
var providerSchema string

//go:embed provider-metadata.yaml
var providerMetadata string

// GetProvider returns provider configuration
func GetProvider() *ujconfig.Provider {
	pc := ujconfig.NewProvider([]byte(providerSchema), resourcePrefix, modulePath, []byte(providerMetadata),
		ujconfig.WithRootGroup("upcloud.io"),
		ujconfig.WithIncludeList(ResourcesList()),
		ujconfig.WithFeaturesPackage("internal/features"),
	)

	for _, configure := range []func(provider *ujconfig.Provider){
		// add custom config functions
		server.Configure,
		network.Configure,
		storage.Configure,
		objectstorage.Configure,
		database.Configure,
		kubernetes.Configure,
	} {
		configure(pc)
	}

	pc.ConfigureResources()
	return pc
}

// ResourcesList returns the list of all resources that should be configured and generated
func ResourcesList() []string {
	allResources := slices.Concat(
		server.Resources,
		network.Resources,
		storage.Resources,
		objectstorage.Resources,
		database.Resources,
		kubernetes.Resources,
	)

	for i, name := range allResources {
		// $ is added to match the exact string since the format is regex.
		allResources[i] = name + "$"
	}

	return allResources
}
