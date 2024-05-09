// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

// Code generated by upjet. DO NOT EDIT.

package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type RegionTargetHTTPProxyInitParameters struct {

	// An optional description of this resource.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// A reference to the RegionUrlMap resource that defines the mapping from URL
	// to the BackendService.
	// +crossplane:generate:reference:type=github.com/upbound/provider-gcp/apis/compute/v1beta1.RegionURLMap
	// +crossplane:generate:reference:extractor=github.com/upbound/provider-gcp/config/common.SelfLinkExtractor()
	URLMap *string `json:"urlMap,omitempty" tf:"url_map,omitempty"`

	// Reference to a RegionURLMap in compute to populate urlMap.
	// +kubebuilder:validation:Optional
	URLMapRef *v1.Reference `json:"urlMapRef,omitempty" tf:"-"`

	// Selector for a RegionURLMap in compute to populate urlMap.
	// +kubebuilder:validation:Optional
	URLMapSelector *v1.Selector `json:"urlMapSelector,omitempty" tf:"-"`
}

type RegionTargetHTTPProxyObservation struct {

	// Creation timestamp in RFC3339 text format.
	CreationTimestamp *string `json:"creationTimestamp,omitempty" tf:"creation_timestamp,omitempty"`

	// An optional description of this resource.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// an identifier for the resource with format projects/{{project}}/regions/{{region}}/targetHttpProxies/{{name}}
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// The unique identifier for the resource.
	ProxyID *float64 `json:"proxyId,omitempty" tf:"proxy_id,omitempty"`

	// The Region in which the created target https proxy should reside.
	// If it is not provided, the provider region is used.
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// The URI of the created resource.
	SelfLink *string `json:"selfLink,omitempty" tf:"self_link,omitempty"`

	// A reference to the RegionUrlMap resource that defines the mapping from URL
	// to the BackendService.
	URLMap *string `json:"urlMap,omitempty" tf:"url_map,omitempty"`
}

type RegionTargetHTTPProxyParameters struct {

	// An optional description of this resource.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	// +kubebuilder:validation:Optional
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// The Region in which the created target https proxy should reside.
	// If it is not provided, the provider region is used.
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"region,omitempty"`

	// A reference to the RegionUrlMap resource that defines the mapping from URL
	// to the BackendService.
	// +crossplane:generate:reference:type=github.com/upbound/provider-gcp/apis/compute/v1beta1.RegionURLMap
	// +crossplane:generate:reference:extractor=github.com/upbound/provider-gcp/config/common.SelfLinkExtractor()
	// +kubebuilder:validation:Optional
	URLMap *string `json:"urlMap,omitempty" tf:"url_map,omitempty"`

	// Reference to a RegionURLMap in compute to populate urlMap.
	// +kubebuilder:validation:Optional
	URLMapRef *v1.Reference `json:"urlMapRef,omitempty" tf:"-"`

	// Selector for a RegionURLMap in compute to populate urlMap.
	// +kubebuilder:validation:Optional
	URLMapSelector *v1.Selector `json:"urlMapSelector,omitempty" tf:"-"`
}

// RegionTargetHTTPProxySpec defines the desired state of RegionTargetHTTPProxy
type RegionTargetHTTPProxySpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     RegionTargetHTTPProxyParameters `json:"forProvider"`
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
	InitProvider RegionTargetHTTPProxyInitParameters `json:"initProvider,omitempty"`
}

// RegionTargetHTTPProxyStatus defines the observed state of RegionTargetHTTPProxy.
type RegionTargetHTTPProxyStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        RegionTargetHTTPProxyObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// RegionTargetHTTPProxy is the Schema for the RegionTargetHTTPProxys API. Represents a RegionTargetHttpProxy resource, which is used by one or more forwarding rules to route incoming HTTP requests to a URL map.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,gcp}
type RegionTargetHTTPProxy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              RegionTargetHTTPProxySpec   `json:"spec"`
	Status            RegionTargetHTTPProxyStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// RegionTargetHTTPProxyList contains a list of RegionTargetHTTPProxys
type RegionTargetHTTPProxyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []RegionTargetHTTPProxy `json:"items"`
}

// Repository type metadata.
var (
	RegionTargetHTTPProxy_Kind             = "RegionTargetHTTPProxy"
	RegionTargetHTTPProxy_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: RegionTargetHTTPProxy_Kind}.String()
	RegionTargetHTTPProxy_KindAPIVersion   = RegionTargetHTTPProxy_Kind + "." + CRDGroupVersion.String()
	RegionTargetHTTPProxy_GroupVersionKind = CRDGroupVersion.WithKind(RegionTargetHTTPProxy_Kind)
)

func init() {
	SchemeBuilder.Register(&RegionTargetHTTPProxy{}, &RegionTargetHTTPProxyList{})
}
