/*
Copyright 2022 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"

	manageddatabaselogicaldatabase "github.com/UpCloudLtd/provider-upcloud/internal/controller/database/manageddatabaselogicaldatabase"
	manageddatabasemysql "github.com/UpCloudLtd/provider-upcloud/internal/controller/database/manageddatabasemysql"
	manageddatabaseopensearch "github.com/UpCloudLtd/provider-upcloud/internal/controller/database/manageddatabaseopensearch"
	manageddatabasepostgresql "github.com/UpCloudLtd/provider-upcloud/internal/controller/database/manageddatabasepostgresql"
	manageddatabaseredis "github.com/UpCloudLtd/provider-upcloud/internal/controller/database/manageddatabaseredis"
	manageddatabaseuser "github.com/UpCloudLtd/provider-upcloud/internal/controller/database/manageddatabaseuser"
	network "github.com/UpCloudLtd/provider-upcloud/internal/controller/network/network"
	router "github.com/UpCloudLtd/provider-upcloud/internal/controller/network/router"
	managedobjectstorage "github.com/UpCloudLtd/provider-upcloud/internal/controller/objectstorage/managedobjectstorage"
	managedobjectstoragepolicy "github.com/UpCloudLtd/provider-upcloud/internal/controller/objectstorage/managedobjectstoragepolicy"
	managedobjectstorageuser "github.com/UpCloudLtd/provider-upcloud/internal/controller/objectstorage/managedobjectstorageuser"
	managedobjectstorageuseraccesskey "github.com/UpCloudLtd/provider-upcloud/internal/controller/objectstorage/managedobjectstorageuseraccesskey"
	managedobjectstorageuserpolicy "github.com/UpCloudLtd/provider-upcloud/internal/controller/objectstorage/managedobjectstorageuserpolicy"
	providerconfig "github.com/UpCloudLtd/provider-upcloud/internal/controller/providerconfig"
	firewallrules "github.com/UpCloudLtd/provider-upcloud/internal/controller/server/firewallrules"
	server "github.com/UpCloudLtd/provider-upcloud/internal/controller/server/server"
	servergroup "github.com/UpCloudLtd/provider-upcloud/internal/controller/server/servergroup"
	storage "github.com/UpCloudLtd/provider-upcloud/internal/controller/storage/storage"
	kubernetescluster "github.com/UpCloudLtd/provider-upcloud/internal/controller/uks/kubernetescluster"
	kubernetesnodegroup "github.com/UpCloudLtd/provider-upcloud/internal/controller/uks/kubernetesnodegroup"
)

// Setup creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		manageddatabaselogicaldatabase.Setup,
		manageddatabasemysql.Setup,
		manageddatabaseopensearch.Setup,
		manageddatabasepostgresql.Setup,
		manageddatabaseredis.Setup,
		manageddatabaseuser.Setup,
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
