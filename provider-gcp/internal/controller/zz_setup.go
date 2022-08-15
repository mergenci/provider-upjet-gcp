/*
Copyright 2021 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/upbound/upjet/pkg/controller"

	folder "github.com/upbound/official-providers/provider-gcp/internal/controller/cloudplatform/folder"
	project "github.com/upbound/official-providers/provider-gcp/internal/controller/cloudplatform/project"
	projectdefaultserviceaccounts "github.com/upbound/official-providers/provider-gcp/internal/controller/cloudplatform/projectdefaultserviceaccounts"
	projectiamauditconfig "github.com/upbound/official-providers/provider-gcp/internal/controller/cloudplatform/projectiamauditconfig"
	projectiambinding "github.com/upbound/official-providers/provider-gcp/internal/controller/cloudplatform/projectiambinding"
	projectiammember "github.com/upbound/official-providers/provider-gcp/internal/controller/cloudplatform/projectiammember"
	projectiampolicy "github.com/upbound/official-providers/provider-gcp/internal/controller/cloudplatform/projectiampolicy"
	projectservice "github.com/upbound/official-providers/provider-gcp/internal/controller/cloudplatform/projectservice"
	projectusageexportbucket "github.com/upbound/official-providers/provider-gcp/internal/controller/cloudplatform/projectusageexportbucket"
	serviceaccount "github.com/upbound/official-providers/provider-gcp/internal/controller/cloudplatform/serviceaccount"
	serviceaccountiambinding "github.com/upbound/official-providers/provider-gcp/internal/controller/cloudplatform/serviceaccountiambinding"
	serviceaccountiammember "github.com/upbound/official-providers/provider-gcp/internal/controller/cloudplatform/serviceaccountiammember"
	serviceaccountiampolicy "github.com/upbound/official-providers/provider-gcp/internal/controller/cloudplatform/serviceaccountiampolicy"
	serviceaccountkey "github.com/upbound/official-providers/provider-gcp/internal/controller/cloudplatform/serviceaccountkey"
	servicenetworkingpeereddnsdomain "github.com/upbound/official-providers/provider-gcp/internal/controller/cloudplatform/servicenetworkingpeereddnsdomain"
	domainmapping "github.com/upbound/official-providers/provider-gcp/internal/controller/cloudrun/domainmapping"
	service "github.com/upbound/official-providers/provider-gcp/internal/controller/cloudrun/service"
	serviceiambinding "github.com/upbound/official-providers/provider-gcp/internal/controller/cloudrun/serviceiambinding"
	serviceiammember "github.com/upbound/official-providers/provider-gcp/internal/controller/cloudrun/serviceiammember"
	serviceiampolicy "github.com/upbound/official-providers/provider-gcp/internal/controller/cloudrun/serviceiampolicy"
	address "github.com/upbound/official-providers/provider-gcp/internal/controller/compute/address"
	disk "github.com/upbound/official-providers/provider-gcp/internal/controller/compute/disk"
	diskiambinding "github.com/upbound/official-providers/provider-gcp/internal/controller/compute/diskiambinding"
	diskiammember "github.com/upbound/official-providers/provider-gcp/internal/controller/compute/diskiammember"
	diskiampolicy "github.com/upbound/official-providers/provider-gcp/internal/controller/compute/diskiampolicy"
	diskresourcepolicyattachment "github.com/upbound/official-providers/provider-gcp/internal/controller/compute/diskresourcepolicyattachment"
	externalvpngateway "github.com/upbound/official-providers/provider-gcp/internal/controller/compute/externalvpngateway"
	firewall "github.com/upbound/official-providers/provider-gcp/internal/controller/compute/firewall"
	globaladdress "github.com/upbound/official-providers/provider-gcp/internal/controller/compute/globaladdress"
	globalnetworkendpoint "github.com/upbound/official-providers/provider-gcp/internal/controller/compute/globalnetworkendpoint"
	globalnetworkendpointgroup "github.com/upbound/official-providers/provider-gcp/internal/controller/compute/globalnetworkendpointgroup"
	havpngateway "github.com/upbound/official-providers/provider-gcp/internal/controller/compute/havpngateway"
	healthcheck "github.com/upbound/official-providers/provider-gcp/internal/controller/compute/healthcheck"
	httphealthcheck "github.com/upbound/official-providers/provider-gcp/internal/controller/compute/httphealthcheck"
	httpshealthcheck "github.com/upbound/official-providers/provider-gcp/internal/controller/compute/httpshealthcheck"
	image "github.com/upbound/official-providers/provider-gcp/internal/controller/compute/image"
	imageiambinding "github.com/upbound/official-providers/provider-gcp/internal/controller/compute/imageiambinding"
	imageiammember "github.com/upbound/official-providers/provider-gcp/internal/controller/compute/imageiammember"
	imageiampolicy "github.com/upbound/official-providers/provider-gcp/internal/controller/compute/imageiampolicy"
	instance "github.com/upbound/official-providers/provider-gcp/internal/controller/compute/instance"
	instancefromtemplate "github.com/upbound/official-providers/provider-gcp/internal/controller/compute/instancefromtemplate"
	instancegroup "github.com/upbound/official-providers/provider-gcp/internal/controller/compute/instancegroup"
	instancegroupmanager "github.com/upbound/official-providers/provider-gcp/internal/controller/compute/instancegroupmanager"
	instanceiambinding "github.com/upbound/official-providers/provider-gcp/internal/controller/compute/instanceiambinding"
	instanceiammember "github.com/upbound/official-providers/provider-gcp/internal/controller/compute/instanceiammember"
	instanceiampolicy "github.com/upbound/official-providers/provider-gcp/internal/controller/compute/instanceiampolicy"
	instancetemplate "github.com/upbound/official-providers/provider-gcp/internal/controller/compute/instancetemplate"
	interconnectattachment "github.com/upbound/official-providers/provider-gcp/internal/controller/compute/interconnectattachment"
	managedsslcertificate "github.com/upbound/official-providers/provider-gcp/internal/controller/compute/managedsslcertificate"
	network "github.com/upbound/official-providers/provider-gcp/internal/controller/compute/network"
	networkendpoint "github.com/upbound/official-providers/provider-gcp/internal/controller/compute/networkendpoint"
	networkendpointgroup "github.com/upbound/official-providers/provider-gcp/internal/controller/compute/networkendpointgroup"
	resourcepolicy "github.com/upbound/official-providers/provider-gcp/internal/controller/compute/resourcepolicy"
	router "github.com/upbound/official-providers/provider-gcp/internal/controller/compute/router"
	routernat "github.com/upbound/official-providers/provider-gcp/internal/controller/compute/routernat"
	subnetwork "github.com/upbound/official-providers/provider-gcp/internal/controller/compute/subnetwork"
	targetpool "github.com/upbound/official-providers/provider-gcp/internal/controller/compute/targetpool"
	cluster "github.com/upbound/official-providers/provider-gcp/internal/controller/container/cluster"
	nodepool "github.com/upbound/official-providers/provider-gcp/internal/controller/container/nodepool"
	managedzone "github.com/upbound/official-providers/provider-gcp/internal/controller/dns/managedzone"
	policy "github.com/upbound/official-providers/provider-gcp/internal/controller/dns/policy"
	recordset "github.com/upbound/official-providers/provider-gcp/internal/controller/dns/recordset"
	alertpolicy "github.com/upbound/official-providers/provider-gcp/internal/controller/monitoring/alertpolicy"
	notificationchannel "github.com/upbound/official-providers/provider-gcp/internal/controller/monitoring/notificationchannel"
	uptimecheckconfig "github.com/upbound/official-providers/provider-gcp/internal/controller/monitoring/uptimecheckconfig"
	providerconfig "github.com/upbound/official-providers/provider-gcp/internal/controller/providerconfig"
	litereservation "github.com/upbound/official-providers/provider-gcp/internal/controller/pubsub/litereservation"
	litesubscription "github.com/upbound/official-providers/provider-gcp/internal/controller/pubsub/litesubscription"
	litetopic "github.com/upbound/official-providers/provider-gcp/internal/controller/pubsub/litetopic"
	schema "github.com/upbound/official-providers/provider-gcp/internal/controller/pubsub/schema"
	subscription "github.com/upbound/official-providers/provider-gcp/internal/controller/pubsub/subscription"
	subscriptioniambinding "github.com/upbound/official-providers/provider-gcp/internal/controller/pubsub/subscriptioniambinding"
	subscriptioniammember "github.com/upbound/official-providers/provider-gcp/internal/controller/pubsub/subscriptioniammember"
	subscriptioniampolicy "github.com/upbound/official-providers/provider-gcp/internal/controller/pubsub/subscriptioniampolicy"
	topic "github.com/upbound/official-providers/provider-gcp/internal/controller/pubsub/topic"
	topiciambinding "github.com/upbound/official-providers/provider-gcp/internal/controller/pubsub/topiciambinding"
	topiciammember "github.com/upbound/official-providers/provider-gcp/internal/controller/pubsub/topiciammember"
	topiciampolicy "github.com/upbound/official-providers/provider-gcp/internal/controller/pubsub/topiciampolicy"
	instanceredis "github.com/upbound/official-providers/provider-gcp/internal/controller/redis/instance"
	secret "github.com/upbound/official-providers/provider-gcp/internal/controller/secretmanager/secret"
	secretversion "github.com/upbound/official-providers/provider-gcp/internal/controller/secretmanager/secretversion"
	database "github.com/upbound/official-providers/provider-gcp/internal/controller/sql/database"
	databaseinstance "github.com/upbound/official-providers/provider-gcp/internal/controller/sql/databaseinstance"
	sourcerepresentationinstance "github.com/upbound/official-providers/provider-gcp/internal/controller/sql/sourcerepresentationinstance"
	sslcert "github.com/upbound/official-providers/provider-gcp/internal/controller/sql/sslcert"
	user "github.com/upbound/official-providers/provider-gcp/internal/controller/sql/user"
	bucket "github.com/upbound/official-providers/provider-gcp/internal/controller/storage/bucket"
	bucketaccesscontrol "github.com/upbound/official-providers/provider-gcp/internal/controller/storage/bucketaccesscontrol"
)

// Setup creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		folder.Setup,
		project.Setup,
		projectdefaultserviceaccounts.Setup,
		projectiamauditconfig.Setup,
		projectiambinding.Setup,
		projectiammember.Setup,
		projectiampolicy.Setup,
		projectservice.Setup,
		projectusageexportbucket.Setup,
		serviceaccount.Setup,
		serviceaccountiambinding.Setup,
		serviceaccountiammember.Setup,
		serviceaccountiampolicy.Setup,
		serviceaccountkey.Setup,
		domainmapping.Setup,
		service.Setup,
		serviceiambinding.Setup,
		serviceiammember.Setup,
		serviceiampolicy.Setup,
		servicenetworkingpeereddnsdomain.Setup,
		address.Setup,
		disk.Setup,
		diskiambinding.Setup,
		diskiammember.Setup,
		diskiampolicy.Setup,
		diskresourcepolicyattachment.Setup,
		externalvpngateway.Setup,
		firewall.Setup,
		globaladdress.Setup,
		globalnetworkendpoint.Setup,
		globalnetworkendpointgroup.Setup,
		havpngateway.Setup,
		healthcheck.Setup,
		httphealthcheck.Setup,
		httpshealthcheck.Setup,
		image.Setup,
		imageiambinding.Setup,
		imageiammember.Setup,
		imageiampolicy.Setup,
		instance.Setup,
		instancefromtemplate.Setup,
		instancegroup.Setup,
		instancegroupmanager.Setup,
		instanceiambinding.Setup,
		instanceiammember.Setup,
		instanceiampolicy.Setup,
		instancetemplate.Setup,
		interconnectattachment.Setup,
		managedsslcertificate.Setup,
		network.Setup,
		networkendpoint.Setup,
		networkendpointgroup.Setup,
		resourcepolicy.Setup,
		router.Setup,
		routernat.Setup,
		subnetwork.Setup,
		targetpool.Setup,
		cluster.Setup,
		nodepool.Setup,
		managedzone.Setup,
		policy.Setup,
		recordset.Setup,
		alertpolicy.Setup,
		notificationchannel.Setup,
		uptimecheckconfig.Setup,
		providerconfig.Setup,
		litereservation.Setup,
		litesubscription.Setup,
		litetopic.Setup,
		schema.Setup,
		subscription.Setup,
		subscriptioniambinding.Setup,
		subscriptioniammember.Setup,
		subscriptioniampolicy.Setup,
		topic.Setup,
		topiciambinding.Setup,
		topiciammember.Setup,
		topiciampolicy.Setup,
		instanceredis.Setup,
		secret.Setup,
		secretversion.Setup,
		database.Setup,
		databaseinstance.Setup,
		sourcerepresentationinstance.Setup,
		sslcert.Setup,
		user.Setup,
		bucket.Setup,
		bucketaccesscontrol.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
