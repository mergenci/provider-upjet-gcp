// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

// Code generated by upjet. DO NOT EDIT.

package v1beta2

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type AuthorizedExternalNetworksInitParameters struct {

	// CIDR range for one authorized network of the instance.
	CidrRange *string `json:"cidrRange,omitempty" tf:"cidr_range,omitempty"`
}

type AuthorizedExternalNetworksObservation struct {

	// CIDR range for one authorized network of the instance.
	CidrRange *string `json:"cidrRange,omitempty" tf:"cidr_range,omitempty"`
}

type AuthorizedExternalNetworksParameters struct {

	// CIDR range for one authorized network of the instance.
	// +kubebuilder:validation:Optional
	CidrRange *string `json:"cidrRange,omitempty" tf:"cidr_range,omitempty"`
}

type ClientConnectionConfigInitParameters struct {

	// Configuration to enforce connectors only (ex: AuthProxy) connections to the database.
	RequireConnectors *bool `json:"requireConnectors,omitempty" tf:"require_connectors,omitempty"`

	// SSL config option for this instance.
	// Structure is documented below.
	SSLConfig *SSLConfigInitParameters `json:"sslConfig,omitempty" tf:"ssl_config,omitempty"`
}

type ClientConnectionConfigObservation struct {

	// Configuration to enforce connectors only (ex: AuthProxy) connections to the database.
	RequireConnectors *bool `json:"requireConnectors,omitempty" tf:"require_connectors,omitempty"`

	// SSL config option for this instance.
	// Structure is documented below.
	SSLConfig *SSLConfigObservation `json:"sslConfig,omitempty" tf:"ssl_config,omitempty"`
}

type ClientConnectionConfigParameters struct {

	// Configuration to enforce connectors only (ex: AuthProxy) connections to the database.
	// +kubebuilder:validation:Optional
	RequireConnectors *bool `json:"requireConnectors,omitempty" tf:"require_connectors,omitempty"`

	// SSL config option for this instance.
	// Structure is documented below.
	// +kubebuilder:validation:Optional
	SSLConfig *SSLConfigParameters `json:"sslConfig,omitempty" tf:"ssl_config,omitempty"`
}

type InstanceInitParameters struct {

	// Annotations to allow client tools to store small amount of arbitrary data. This is distinct from labels.
	// Note: This field is non-authoritative, and will only manage the annotations present in your configuration.
	// Please refer to the field effective_annotations for all of the annotations present on the resource.
	// +mapType=granular
	Annotations map[string]*string `json:"annotations,omitempty" tf:"annotations,omitempty"`

	// 'Availability type of an Instance. Defaults to REGIONAL for both primary and read instances.
	// Note that primary and read instances can have different availability types.
	// Only READ_POOL instance supports ZONAL type. Users can't specify the zone for READ_POOL instance.
	// Zone is automatically chosen from the list of zones in the region specified.
	// Read pool of size 1 can only have zonal availability. Read pools with node count of 2 or more
	// can have regional availability (nodes are present in 2 or more zones in a region).'
	// Possible values are: AVAILABILITY_TYPE_UNSPECIFIED, ZONAL, REGIONAL.
	AvailabilityType *string `json:"availabilityType,omitempty" tf:"availability_type,omitempty"`

	// Client connection specific configurations.
	// Structure is documented below.
	ClientConnectionConfig *ClientConnectionConfigInitParameters `json:"clientConnectionConfig,omitempty" tf:"client_connection_config,omitempty"`

	// Database flags. Set at instance level. * They are copied from primary instance on read instance creation. * Read instances can set new or override existing flags that are relevant for reads, e.g. for enabling columnar cache on a read instance. Flags set on read instance may or may not be present on primary.
	// +mapType=granular
	DatabaseFlags map[string]*string `json:"databaseFlags,omitempty" tf:"database_flags,omitempty"`

	// User-settable and human-readable display name for the Instance.
	DisplayName *string `json:"displayName,omitempty" tf:"display_name,omitempty"`

	// The Compute Engine zone that the instance should serve from, per https://cloud.google.com/compute/docs/regions-zones This can ONLY be specified for ZONAL instances. If present for a REGIONAL instance, an error will be thrown. If this is absent for a ZONAL instance, instance is created in a random zone with available capacity.
	GceZone *string `json:"gceZone,omitempty" tf:"gce_zone,omitempty"`

	// The type of the instance.
	// If the instance type is READ_POOL, provide the associated PRIMARY/SECONDARY instance in the depends_on meta-data attribute.
	// If the instance type is SECONDARY, point to the cluster_type of the associated secondary cluster instead of mentioning SECONDARY.
	// Example: {instance_type = google_alloydb_cluster.<secondary_cluster_name>.
	// Use deletion_policy = "FORCE" in the associated secondary cluster and delete the cluster forcefully to delete the secondary cluster as well its associated secondary instance.
	// Possible values are: PRIMARY, READ_POOL, SECONDARY.
	// +crossplane:generate:reference:type=github.com/upbound/provider-gcp/apis/alloydb/v1beta2.Cluster
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractParamPath("cluster_type",false)
	InstanceType *string `json:"instanceType,omitempty" tf:"instance_type,omitempty"`

	// Reference to a Cluster in alloydb to populate instanceType.
	// +kubebuilder:validation:Optional
	InstanceTypeRef *v1.Reference `json:"instanceTypeRef,omitempty" tf:"-"`

	// Selector for a Cluster in alloydb to populate instanceType.
	// +kubebuilder:validation:Optional
	InstanceTypeSelector *v1.Selector `json:"instanceTypeSelector,omitempty" tf:"-"`

	// User-defined labels for the alloydb instance.
	// Note: This field is non-authoritative, and will only manage the labels present in your configuration.
	// Please refer to the field effective_labels for all of the labels present on the resource.
	// +mapType=granular
	Labels map[string]*string `json:"labels,omitempty" tf:"labels,omitempty"`

	// Configurations for the machines that host the underlying database engine.
	// Structure is documented below.
	MachineConfig *MachineConfigInitParameters `json:"machineConfig,omitempty" tf:"machine_config,omitempty"`

	// Instance level network configuration.
	// Structure is documented below.
	NetworkConfig *InstanceNetworkConfigInitParameters `json:"networkConfig,omitempty" tf:"network_config,omitempty"`

	// Configuration for query insights.
	// Structure is documented below.
	QueryInsightsConfig *QueryInsightsConfigInitParameters `json:"queryInsightsConfig,omitempty" tf:"query_insights_config,omitempty"`

	// Read pool specific config. If the instance type is READ_POOL, this configuration must be provided.
	// Structure is documented below.
	ReadPoolConfig *ReadPoolConfigInitParameters `json:"readPoolConfig,omitempty" tf:"read_pool_config,omitempty"`
}

type InstanceNetworkConfigInitParameters struct {

	// A list of external networks authorized to access this instance. This
	// field is only allowed to be set when enable_public_ip is set to
	// true.
	// Structure is documented below.
	AuthorizedExternalNetworks []AuthorizedExternalNetworksInitParameters `json:"authorizedExternalNetworks,omitempty" tf:"authorized_external_networks,omitempty"`

	// Enabling public ip for the instance. If a user wishes to disable this,
	// please also clear the list of the authorized external networks set on
	// the same instance.
	EnablePublicIP *bool `json:"enablePublicIp,omitempty" tf:"enable_public_ip,omitempty"`
}

type InstanceNetworkConfigObservation struct {

	// A list of external networks authorized to access this instance. This
	// field is only allowed to be set when enable_public_ip is set to
	// true.
	// Structure is documented below.
	AuthorizedExternalNetworks []AuthorizedExternalNetworksObservation `json:"authorizedExternalNetworks,omitempty" tf:"authorized_external_networks,omitempty"`

	// Enabling public ip for the instance. If a user wishes to disable this,
	// please also clear the list of the authorized external networks set on
	// the same instance.
	EnablePublicIP *bool `json:"enablePublicIp,omitempty" tf:"enable_public_ip,omitempty"`
}

type InstanceNetworkConfigParameters struct {

	// A list of external networks authorized to access this instance. This
	// field is only allowed to be set when enable_public_ip is set to
	// true.
	// Structure is documented below.
	// +kubebuilder:validation:Optional
	AuthorizedExternalNetworks []AuthorizedExternalNetworksParameters `json:"authorizedExternalNetworks,omitempty" tf:"authorized_external_networks,omitempty"`

	// Enabling public ip for the instance. If a user wishes to disable this,
	// please also clear the list of the authorized external networks set on
	// the same instance.
	// +kubebuilder:validation:Optional
	EnablePublicIP *bool `json:"enablePublicIp,omitempty" tf:"enable_public_ip,omitempty"`
}

type InstanceObservation struct {

	// Annotations to allow client tools to store small amount of arbitrary data. This is distinct from labels.
	// Note: This field is non-authoritative, and will only manage the annotations present in your configuration.
	// Please refer to the field effective_annotations for all of the annotations present on the resource.
	// +mapType=granular
	Annotations map[string]*string `json:"annotations,omitempty" tf:"annotations,omitempty"`

	// 'Availability type of an Instance. Defaults to REGIONAL for both primary and read instances.
	// Note that primary and read instances can have different availability types.
	// Only READ_POOL instance supports ZONAL type. Users can't specify the zone for READ_POOL instance.
	// Zone is automatically chosen from the list of zones in the region specified.
	// Read pool of size 1 can only have zonal availability. Read pools with node count of 2 or more
	// can have regional availability (nodes are present in 2 or more zones in a region).'
	// Possible values are: AVAILABILITY_TYPE_UNSPECIFIED, ZONAL, REGIONAL.
	AvailabilityType *string `json:"availabilityType,omitempty" tf:"availability_type,omitempty"`

	// Client connection specific configurations.
	// Structure is documented below.
	ClientConnectionConfig *ClientConnectionConfigObservation `json:"clientConnectionConfig,omitempty" tf:"client_connection_config,omitempty"`

	// Identifies the alloydb cluster. Must be in the format
	// 'projects/{project}/locations/{location}/clusters/{cluster_id}'
	Cluster *string `json:"cluster,omitempty" tf:"cluster,omitempty"`

	// Time the Instance was created in UTC.
	CreateTime *string `json:"createTime,omitempty" tf:"create_time,omitempty"`

	// Database flags. Set at instance level. * They are copied from primary instance on read instance creation. * Read instances can set new or override existing flags that are relevant for reads, e.g. for enabling columnar cache on a read instance. Flags set on read instance may or may not be present on primary.
	// +mapType=granular
	DatabaseFlags map[string]*string `json:"databaseFlags,omitempty" tf:"database_flags,omitempty"`

	// User-settable and human-readable display name for the Instance.
	DisplayName *string `json:"displayName,omitempty" tf:"display_name,omitempty"`

	// +mapType=granular
	EffectiveAnnotations map[string]*string `json:"effectiveAnnotations,omitempty" tf:"effective_annotations,omitempty"`

	// +mapType=granular
	EffectiveLabels map[string]*string `json:"effectiveLabels,omitempty" tf:"effective_labels,omitempty"`

	// The Compute Engine zone that the instance should serve from, per https://cloud.google.com/compute/docs/regions-zones This can ONLY be specified for ZONAL instances. If present for a REGIONAL instance, an error will be thrown. If this is absent for a ZONAL instance, instance is created in a random zone with available capacity.
	GceZone *string `json:"gceZone,omitempty" tf:"gce_zone,omitempty"`

	// an identifier for the resource with format {{cluster}}/instances/{{instance_id}}
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The IP address for the Instance. This is the connection endpoint for an end-user application.
	IPAddress *string `json:"ipAddress,omitempty" tf:"ip_address,omitempty"`

	// The type of the instance.
	// If the instance type is READ_POOL, provide the associated PRIMARY/SECONDARY instance in the depends_on meta-data attribute.
	// If the instance type is SECONDARY, point to the cluster_type of the associated secondary cluster instead of mentioning SECONDARY.
	// Example: {instance_type = google_alloydb_cluster.<secondary_cluster_name>.
	// Use deletion_policy = "FORCE" in the associated secondary cluster and delete the cluster forcefully to delete the secondary cluster as well its associated secondary instance.
	// Possible values are: PRIMARY, READ_POOL, SECONDARY.
	InstanceType *string `json:"instanceType,omitempty" tf:"instance_type,omitempty"`

	// User-defined labels for the alloydb instance.
	// Note: This field is non-authoritative, and will only manage the labels present in your configuration.
	// Please refer to the field effective_labels for all of the labels present on the resource.
	// +mapType=granular
	Labels map[string]*string `json:"labels,omitempty" tf:"labels,omitempty"`

	// Configurations for the machines that host the underlying database engine.
	// Structure is documented below.
	MachineConfig *MachineConfigObservation `json:"machineConfig,omitempty" tf:"machine_config,omitempty"`

	// The name of the instance resource.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Instance level network configuration.
	// Structure is documented below.
	NetworkConfig *InstanceNetworkConfigObservation `json:"networkConfig,omitempty" tf:"network_config,omitempty"`

	// The public IP addresses for the Instance. This is available ONLY when
	// networkConfig.enablePublicIp is set to true. This is the connection
	// endpoint for an end-user application.
	PublicIPAddress *string `json:"publicIpAddress,omitempty" tf:"public_ip_address,omitempty"`

	// Configuration for query insights.
	// Structure is documented below.
	QueryInsightsConfig *QueryInsightsConfigObservation `json:"queryInsightsConfig,omitempty" tf:"query_insights_config,omitempty"`

	// Read pool specific config. If the instance type is READ_POOL, this configuration must be provided.
	// Structure is documented below.
	ReadPoolConfig *ReadPoolConfigObservation `json:"readPoolConfig,omitempty" tf:"read_pool_config,omitempty"`

	// Set to true if the current state of Instance does not match the user's intended state, and the service is actively updating the resource to reconcile them. This can happen due to user-triggered updates or system actions like failover or maintenance.
	Reconciling *bool `json:"reconciling,omitempty" tf:"reconciling,omitempty"`

	// The current state of the alloydb instance.
	State *string `json:"state,omitempty" tf:"state,omitempty"`

	// The combination of labels configured directly on the resource
	// and default labels configured on the provider.
	// +mapType=granular
	TerraformLabels map[string]*string `json:"terraformLabels,omitempty" tf:"terraform_labels,omitempty"`

	// The system-generated UID of the resource.
	UID *string `json:"uid,omitempty" tf:"uid,omitempty"`

	// Time the Instance was updated in UTC.
	UpdateTime *string `json:"updateTime,omitempty" tf:"update_time,omitempty"`
}

type InstanceParameters struct {

	// Annotations to allow client tools to store small amount of arbitrary data. This is distinct from labels.
	// Note: This field is non-authoritative, and will only manage the annotations present in your configuration.
	// Please refer to the field effective_annotations for all of the annotations present on the resource.
	// +kubebuilder:validation:Optional
	// +mapType=granular
	Annotations map[string]*string `json:"annotations,omitempty" tf:"annotations,omitempty"`

	// 'Availability type of an Instance. Defaults to REGIONAL for both primary and read instances.
	// Note that primary and read instances can have different availability types.
	// Only READ_POOL instance supports ZONAL type. Users can't specify the zone for READ_POOL instance.
	// Zone is automatically chosen from the list of zones in the region specified.
	// Read pool of size 1 can only have zonal availability. Read pools with node count of 2 or more
	// can have regional availability (nodes are present in 2 or more zones in a region).'
	// Possible values are: AVAILABILITY_TYPE_UNSPECIFIED, ZONAL, REGIONAL.
	// +kubebuilder:validation:Optional
	AvailabilityType *string `json:"availabilityType,omitempty" tf:"availability_type,omitempty"`

	// Client connection specific configurations.
	// Structure is documented below.
	// +kubebuilder:validation:Optional
	ClientConnectionConfig *ClientConnectionConfigParameters `json:"clientConnectionConfig,omitempty" tf:"client_connection_config,omitempty"`

	// Identifies the alloydb cluster. Must be in the format
	// 'projects/{project}/locations/{location}/clusters/{cluster_id}'
	// +crossplane:generate:reference:type=github.com/upbound/provider-gcp/apis/alloydb/v1beta2.Cluster
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractParamPath("name",true)
	// +kubebuilder:validation:Optional
	Cluster *string `json:"cluster,omitempty" tf:"cluster,omitempty"`

	// Reference to a Cluster in alloydb to populate cluster.
	// +kubebuilder:validation:Optional
	ClusterRef *v1.Reference `json:"clusterRef,omitempty" tf:"-"`

	// Selector for a Cluster in alloydb to populate cluster.
	// +kubebuilder:validation:Optional
	ClusterSelector *v1.Selector `json:"clusterSelector,omitempty" tf:"-"`

	// Database flags. Set at instance level. * They are copied from primary instance on read instance creation. * Read instances can set new or override existing flags that are relevant for reads, e.g. for enabling columnar cache on a read instance. Flags set on read instance may or may not be present on primary.
	// +kubebuilder:validation:Optional
	// +mapType=granular
	DatabaseFlags map[string]*string `json:"databaseFlags,omitempty" tf:"database_flags,omitempty"`

	// User-settable and human-readable display name for the Instance.
	// +kubebuilder:validation:Optional
	DisplayName *string `json:"displayName,omitempty" tf:"display_name,omitempty"`

	// The Compute Engine zone that the instance should serve from, per https://cloud.google.com/compute/docs/regions-zones This can ONLY be specified for ZONAL instances. If present for a REGIONAL instance, an error will be thrown. If this is absent for a ZONAL instance, instance is created in a random zone with available capacity.
	// +kubebuilder:validation:Optional
	GceZone *string `json:"gceZone,omitempty" tf:"gce_zone,omitempty"`

	// The type of the instance.
	// If the instance type is READ_POOL, provide the associated PRIMARY/SECONDARY instance in the depends_on meta-data attribute.
	// If the instance type is SECONDARY, point to the cluster_type of the associated secondary cluster instead of mentioning SECONDARY.
	// Example: {instance_type = google_alloydb_cluster.<secondary_cluster_name>.
	// Use deletion_policy = "FORCE" in the associated secondary cluster and delete the cluster forcefully to delete the secondary cluster as well its associated secondary instance.
	// Possible values are: PRIMARY, READ_POOL, SECONDARY.
	// +crossplane:generate:reference:type=github.com/upbound/provider-gcp/apis/alloydb/v1beta2.Cluster
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractParamPath("cluster_type",false)
	// +kubebuilder:validation:Optional
	InstanceType *string `json:"instanceType,omitempty" tf:"instance_type,omitempty"`

	// Reference to a Cluster in alloydb to populate instanceType.
	// +kubebuilder:validation:Optional
	InstanceTypeRef *v1.Reference `json:"instanceTypeRef,omitempty" tf:"-"`

	// Selector for a Cluster in alloydb to populate instanceType.
	// +kubebuilder:validation:Optional
	InstanceTypeSelector *v1.Selector `json:"instanceTypeSelector,omitempty" tf:"-"`

	// User-defined labels for the alloydb instance.
	// Note: This field is non-authoritative, and will only manage the labels present in your configuration.
	// Please refer to the field effective_labels for all of the labels present on the resource.
	// +kubebuilder:validation:Optional
	// +mapType=granular
	Labels map[string]*string `json:"labels,omitempty" tf:"labels,omitempty"`

	// Configurations for the machines that host the underlying database engine.
	// Structure is documented below.
	// +kubebuilder:validation:Optional
	MachineConfig *MachineConfigParameters `json:"machineConfig,omitempty" tf:"machine_config,omitempty"`

	// Instance level network configuration.
	// Structure is documented below.
	// +kubebuilder:validation:Optional
	NetworkConfig *InstanceNetworkConfigParameters `json:"networkConfig,omitempty" tf:"network_config,omitempty"`

	// Configuration for query insights.
	// Structure is documented below.
	// +kubebuilder:validation:Optional
	QueryInsightsConfig *QueryInsightsConfigParameters `json:"queryInsightsConfig,omitempty" tf:"query_insights_config,omitempty"`

	// Read pool specific config. If the instance type is READ_POOL, this configuration must be provided.
	// Structure is documented below.
	// +kubebuilder:validation:Optional
	ReadPoolConfig *ReadPoolConfigParameters `json:"readPoolConfig,omitempty" tf:"read_pool_config,omitempty"`
}

type MachineConfigInitParameters struct {

	// The number of CPU's in the VM instance.
	CPUCount *float64 `json:"cpuCount,omitempty" tf:"cpu_count,omitempty"`
}

type MachineConfigObservation struct {

	// The number of CPU's in the VM instance.
	CPUCount *float64 `json:"cpuCount,omitempty" tf:"cpu_count,omitempty"`
}

type MachineConfigParameters struct {

	// The number of CPU's in the VM instance.
	// +kubebuilder:validation:Optional
	CPUCount *float64 `json:"cpuCount,omitempty" tf:"cpu_count,omitempty"`
}

type QueryInsightsConfigInitParameters struct {

	// Number of query execution plans captured by Insights per minute for all queries combined. The default value is 5. Any integer between 0 and 20 is considered valid.
	QueryPlansPerMinute *float64 `json:"queryPlansPerMinute,omitempty" tf:"query_plans_per_minute,omitempty"`

	// Query string length. The default value is 1024. Any integer between 256 and 4500 is considered valid.
	QueryStringLength *float64 `json:"queryStringLength,omitempty" tf:"query_string_length,omitempty"`

	// Record application tags for an instance. This flag is turned "on" by default.
	RecordApplicationTags *bool `json:"recordApplicationTags,omitempty" tf:"record_application_tags,omitempty"`

	// Record client address for an instance. Client address is PII information. This flag is turned "on" by default.
	RecordClientAddress *bool `json:"recordClientAddress,omitempty" tf:"record_client_address,omitempty"`
}

type QueryInsightsConfigObservation struct {

	// Number of query execution plans captured by Insights per minute for all queries combined. The default value is 5. Any integer between 0 and 20 is considered valid.
	QueryPlansPerMinute *float64 `json:"queryPlansPerMinute,omitempty" tf:"query_plans_per_minute,omitempty"`

	// Query string length. The default value is 1024. Any integer between 256 and 4500 is considered valid.
	QueryStringLength *float64 `json:"queryStringLength,omitempty" tf:"query_string_length,omitempty"`

	// Record application tags for an instance. This flag is turned "on" by default.
	RecordApplicationTags *bool `json:"recordApplicationTags,omitempty" tf:"record_application_tags,omitempty"`

	// Record client address for an instance. Client address is PII information. This flag is turned "on" by default.
	RecordClientAddress *bool `json:"recordClientAddress,omitempty" tf:"record_client_address,omitempty"`
}

type QueryInsightsConfigParameters struct {

	// Number of query execution plans captured by Insights per minute for all queries combined. The default value is 5. Any integer between 0 and 20 is considered valid.
	// +kubebuilder:validation:Optional
	QueryPlansPerMinute *float64 `json:"queryPlansPerMinute,omitempty" tf:"query_plans_per_minute,omitempty"`

	// Query string length. The default value is 1024. Any integer between 256 and 4500 is considered valid.
	// +kubebuilder:validation:Optional
	QueryStringLength *float64 `json:"queryStringLength,omitempty" tf:"query_string_length,omitempty"`

	// Record application tags for an instance. This flag is turned "on" by default.
	// +kubebuilder:validation:Optional
	RecordApplicationTags *bool `json:"recordApplicationTags,omitempty" tf:"record_application_tags,omitempty"`

	// Record client address for an instance. Client address is PII information. This flag is turned "on" by default.
	// +kubebuilder:validation:Optional
	RecordClientAddress *bool `json:"recordClientAddress,omitempty" tf:"record_client_address,omitempty"`
}

type ReadPoolConfigInitParameters struct {

	// Read capacity, i.e. number of nodes in a read pool instance.
	NodeCount *float64 `json:"nodeCount,omitempty" tf:"node_count,omitempty"`
}

type ReadPoolConfigObservation struct {

	// Read capacity, i.e. number of nodes in a read pool instance.
	NodeCount *float64 `json:"nodeCount,omitempty" tf:"node_count,omitempty"`
}

type ReadPoolConfigParameters struct {

	// Read capacity, i.e. number of nodes in a read pool instance.
	// +kubebuilder:validation:Optional
	NodeCount *float64 `json:"nodeCount,omitempty" tf:"node_count,omitempty"`
}

type SSLConfigInitParameters struct {

	// SSL mode. Specifies client-server SSL/TLS connection behavior.
	// Possible values are: ENCRYPTED_ONLY, ALLOW_UNENCRYPTED_AND_ENCRYPTED.
	SSLMode *string `json:"sslMode,omitempty" tf:"ssl_mode,omitempty"`
}

type SSLConfigObservation struct {

	// SSL mode. Specifies client-server SSL/TLS connection behavior.
	// Possible values are: ENCRYPTED_ONLY, ALLOW_UNENCRYPTED_AND_ENCRYPTED.
	SSLMode *string `json:"sslMode,omitempty" tf:"ssl_mode,omitempty"`
}

type SSLConfigParameters struct {

	// SSL mode. Specifies client-server SSL/TLS connection behavior.
	// Possible values are: ENCRYPTED_ONLY, ALLOW_UNENCRYPTED_AND_ENCRYPTED.
	// +kubebuilder:validation:Optional
	SSLMode *string `json:"sslMode,omitempty" tf:"ssl_mode,omitempty"`
}

// InstanceSpec defines the desired state of Instance
type InstanceSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     InstanceParameters `json:"forProvider"`
	// THIS IS A BETA FIELD. It will be honored
	// unless the Management Policies feature flag is disabled.
	// InitProvider holds the same fields as ForProvider, with the exception
	// of Identifier and other resource reference fields. The fields that are
	// in InitProvider are merged into ForProvider when the resource is created.
	// The same fields are also added to the terraform ignore_changes hook, to
	// avoid updating them after creation. This is useful for fields that are
	// required on creation, but we do not desire to update them after creation,
	// for example because of an external controller is managing them, like an
	// autoscaler.
	InitProvider InstanceInitParameters `json:"initProvider,omitempty"`
}

// InstanceStatus defines the observed state of Instance.
type InstanceStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        InstanceObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// Instance is the Schema for the Instances API. A managed alloydb cluster instance.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,gcp}
type Instance struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              InstanceSpec   `json:"spec"`
	Status            InstanceStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// InstanceList contains a list of Instances
type InstanceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Instance `json:"items"`
}

// Repository type metadata.
var (
	Instance_Kind             = "Instance"
	Instance_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Instance_Kind}.String()
	Instance_KindAPIVersion   = Instance_Kind + "." + CRDGroupVersion.String()
	Instance_GroupVersionKind = CRDGroupVersion.WithKind(Instance_Kind)
)

func init() {
	SchemeBuilder.Register(&Instance{}, &InstanceList{})
}
