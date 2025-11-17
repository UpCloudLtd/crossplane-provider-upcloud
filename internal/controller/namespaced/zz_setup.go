/*
Copyright 2022 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/v2/pkg/controller"

	manageddatabaselogicaldatabase "github.com/UpCloudLtd/crossplane-provider-upcloud/internal/controller/namespaced/database/manageddatabaselogicaldatabase"
	manageddatabasemysql "github.com/UpCloudLtd/crossplane-provider-upcloud/internal/controller/namespaced/database/manageddatabasemysql"
	manageddatabaseopensearch "github.com/UpCloudLtd/crossplane-provider-upcloud/internal/controller/namespaced/database/manageddatabaseopensearch"
	manageddatabasepostgresql "github.com/UpCloudLtd/crossplane-provider-upcloud/internal/controller/namespaced/database/manageddatabasepostgresql"
	manageddatabaseuser "github.com/UpCloudLtd/crossplane-provider-upcloud/internal/controller/namespaced/database/manageddatabaseuser"
	manageddatabasevalkey "github.com/UpCloudLtd/crossplane-provider-upcloud/internal/controller/namespaced/database/manageddatabasevalkey"
	network "github.com/UpCloudLtd/crossplane-provider-upcloud/internal/controller/namespaced/network/network"
	router "github.com/UpCloudLtd/crossplane-provider-upcloud/internal/controller/namespaced/network/router"
	managedobjectstorage "github.com/UpCloudLtd/crossplane-provider-upcloud/internal/controller/namespaced/objectstorage/managedobjectstorage"
	managedobjectstoragepolicy "github.com/UpCloudLtd/crossplane-provider-upcloud/internal/controller/namespaced/objectstorage/managedobjectstoragepolicy"
	managedobjectstorageuser "github.com/UpCloudLtd/crossplane-provider-upcloud/internal/controller/namespaced/objectstorage/managedobjectstorageuser"
	managedobjectstorageuseraccesskey "github.com/UpCloudLtd/crossplane-provider-upcloud/internal/controller/namespaced/objectstorage/managedobjectstorageuseraccesskey"
	managedobjectstorageuserpolicy "github.com/UpCloudLtd/crossplane-provider-upcloud/internal/controller/namespaced/objectstorage/managedobjectstorageuserpolicy"
	providerconfig "github.com/UpCloudLtd/crossplane-provider-upcloud/internal/controller/namespaced/providerconfig"
	firewallrules "github.com/UpCloudLtd/crossplane-provider-upcloud/internal/controller/namespaced/server/firewallrules"
	server "github.com/UpCloudLtd/crossplane-provider-upcloud/internal/controller/namespaced/server/server"
	servergroup "github.com/UpCloudLtd/crossplane-provider-upcloud/internal/controller/namespaced/server/servergroup"
	storage "github.com/UpCloudLtd/crossplane-provider-upcloud/internal/controller/namespaced/storage/storage"
	kubernetescluster "github.com/UpCloudLtd/crossplane-provider-upcloud/internal/controller/namespaced/uks/kubernetescluster"
	kubernetesnodegroup "github.com/UpCloudLtd/crossplane-provider-upcloud/internal/controller/namespaced/uks/kubernetesnodegroup"
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
