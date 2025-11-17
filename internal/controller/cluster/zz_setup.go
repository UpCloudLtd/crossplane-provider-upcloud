/*
Copyright 2022 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/v2/pkg/controller"

	manageddatabaselogicaldatabase "github.com/UpCloudLtd/crossplane-provider-upcloud/internal/controller/cluster/database/manageddatabaselogicaldatabase"
	manageddatabasemysql "github.com/UpCloudLtd/crossplane-provider-upcloud/internal/controller/cluster/database/manageddatabasemysql"
	manageddatabaseopensearch "github.com/UpCloudLtd/crossplane-provider-upcloud/internal/controller/cluster/database/manageddatabaseopensearch"
	manageddatabasepostgresql "github.com/UpCloudLtd/crossplane-provider-upcloud/internal/controller/cluster/database/manageddatabasepostgresql"
	manageddatabaseuser "github.com/UpCloudLtd/crossplane-provider-upcloud/internal/controller/cluster/database/manageddatabaseuser"
	manageddatabasevalkey "github.com/UpCloudLtd/crossplane-provider-upcloud/internal/controller/cluster/database/manageddatabasevalkey"
	network "github.com/UpCloudLtd/crossplane-provider-upcloud/internal/controller/cluster/network/network"
	router "github.com/UpCloudLtd/crossplane-provider-upcloud/internal/controller/cluster/network/router"
	managedobjectstorage "github.com/UpCloudLtd/crossplane-provider-upcloud/internal/controller/cluster/objectstorage/managedobjectstorage"
	managedobjectstoragepolicy "github.com/UpCloudLtd/crossplane-provider-upcloud/internal/controller/cluster/objectstorage/managedobjectstoragepolicy"
	managedobjectstorageuser "github.com/UpCloudLtd/crossplane-provider-upcloud/internal/controller/cluster/objectstorage/managedobjectstorageuser"
	managedobjectstorageuseraccesskey "github.com/UpCloudLtd/crossplane-provider-upcloud/internal/controller/cluster/objectstorage/managedobjectstorageuseraccesskey"
	managedobjectstorageuserpolicy "github.com/UpCloudLtd/crossplane-provider-upcloud/internal/controller/cluster/objectstorage/managedobjectstorageuserpolicy"
	providerconfig "github.com/UpCloudLtd/crossplane-provider-upcloud/internal/controller/cluster/providerconfig"
	firewallrules "github.com/UpCloudLtd/crossplane-provider-upcloud/internal/controller/cluster/server/firewallrules"
	server "github.com/UpCloudLtd/crossplane-provider-upcloud/internal/controller/cluster/server/server"
	servergroup "github.com/UpCloudLtd/crossplane-provider-upcloud/internal/controller/cluster/server/servergroup"
	storage "github.com/UpCloudLtd/crossplane-provider-upcloud/internal/controller/cluster/storage/storage"
	kubernetescluster "github.com/UpCloudLtd/crossplane-provider-upcloud/internal/controller/cluster/uks/kubernetescluster"
	kubernetesnodegroup "github.com/UpCloudLtd/crossplane-provider-upcloud/internal/controller/cluster/uks/kubernetesnodegroup"
)

// Setup creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		manageddatabaselogicaldatabase.Setup,
		manageddatabasemysql.Setup,
		manageddatabaseopensearch.Setup,
		manageddatabasepostgresql.Setup,
		manageddatabaseuser.Setup,
		manageddatabasevalkey.Setup,
		network.Setup,
		router.Setup,
		managedobjectstorage.Setup,
		managedobjectstoragepolicy.Setup,
		managedobjectstorageuser.Setup,
		managedobjectstorageuseraccesskey.Setup,
		managedobjectstorageuserpolicy.Setup,
		providerconfig.Setup,
		firewallrules.Setup,
		server.Setup,
		servergroup.Setup,
		storage.Setup,
		kubernetescluster.Setup,
		kubernetesnodegroup.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

// SetupGated creates all controllers with the supplied logger and adds them to
// the supplied manager gated.
func SetupGated(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		manageddatabaselogicaldatabase.SetupGated,
		manageddatabasemysql.SetupGated,
		manageddatabaseopensearch.SetupGated,
		manageddatabasepostgresql.SetupGated,
		manageddatabaseuser.SetupGated,
		manageddatabasevalkey.SetupGated,
		network.SetupGated,
		router.SetupGated,
		managedobjectstorage.SetupGated,
		managedobjectstoragepolicy.SetupGated,
		managedobjectstorageuser.SetupGated,
		managedobjectstorageuseraccesskey.SetupGated,
		managedobjectstorageuserpolicy.SetupGated,
		providerconfig.SetupGated,
		firewallrules.SetupGated,
		server.SetupGated,
		servergroup.SetupGated,
		storage.SetupGated,
		kubernetescluster.SetupGated,
		kubernetesnodegroup.SetupGated,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
