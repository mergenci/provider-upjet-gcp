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

type DNSAuthorizationInitParameters struct {

	// A human-readable description of the resource.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// A domain which is being authorized. A DnsAuthorization resource covers a
	// single domain and its wildcard, e.g. authorization for "example.com" can
	// be used to issue certificates for "example.com" and "*.example.com".
	Domain *string `json:"domain,omitempty" tf:"domain,omitempty"`

	// Set of label tags associated with the DNS Authorization resource.
	// Note: This field is non-authoritative, and will only manage the labels present in your configuration.
	// Please refer to the field effective_labels for all of the labels present on the resource.
	// +mapType=granular
	Labels map[string]*string `json:"labels,omitempty" tf:"labels,omitempty"`

	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// type of DNS authorization. If unset during the resource creation, FIXED_RECORD will
	// be used for global resources, and PER_PROJECT_RECORD will be used for other locations.
	// FIXED_RECORD DNS authorization uses DNS-01 validation method
	// PER_PROJECT_RECORD DNS authorization allows for independent management
	// of Google-managed certificates with DNS authorization across multiple
	// projects.
	// Possible values are: FIXED_RECORD, PER_PROJECT_RECORD.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type DNSAuthorizationObservation struct {

	// The structure describing the DNS Resource Record that needs to be added
	// to DNS configuration for the authorization to be usable by
	// certificate.
	// Structure is documented below.
	DNSResourceRecord []DNSResourceRecordObservation `json:"dnsResourceRecord,omitempty" tf:"dns_resource_record,omitempty"`

	// A human-readable description of the resource.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// A domain which is being authorized. A DnsAuthorization resource covers a
	// single domain and its wildcard, e.g. authorization for "example.com" can
	// be used to issue certificates for "example.com" and "*.example.com".
	Domain *string `json:"domain,omitempty" tf:"domain,omitempty"`

	// +mapType=granular
	EffectiveLabels map[string]*string `json:"effectiveLabels,omitempty" tf:"effective_labels,omitempty"`

	// an identifier for the resource with format projects/{{project}}/locations/{{location}}/dnsAuthorizations/{{name}}
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Set of label tags associated with the DNS Authorization resource.
	// Note: This field is non-authoritative, and will only manage the labels present in your configuration.
	// Please refer to the field effective_labels for all of the labels present on the resource.
	// +mapType=granular
	Labels map[string]*string `json:"labels,omitempty" tf:"labels,omitempty"`

	// The Certificate Manager location. If not specified, "global" is used.
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// The combination of labels configured directly on the resource
	// and default labels configured on the provider.
	// +mapType=granular
	TerraformLabels map[string]*string `json:"terraformLabels,omitempty" tf:"terraform_labels,omitempty"`

	// type of DNS authorization. If unset during the resource creation, FIXED_RECORD will
	// be used for global resources, and PER_PROJECT_RECORD will be used for other locations.
	// FIXED_RECORD DNS authorization uses DNS-01 validation method
	// PER_PROJECT_RECORD DNS authorization allows for independent management
	// of Google-managed certificates with DNS authorization across multiple
	// projects.
	// Possible values are: FIXED_RECORD, PER_PROJECT_RECORD.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type DNSAuthorizationParameters struct {

	// A human-readable description of the resource.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// A domain which is being authorized. A DnsAuthorization resource covers a
	// single domain and its wildcard, e.g. authorization for "example.com" can
	// be used to issue certificates for "example.com" and "*.example.com".
	// +kubebuilder:validation:Optional
	Domain *string `json:"domain,omitempty" tf:"domain,omitempty"`

	// Set of label tags associated with the DNS Authorization resource.
	// Note: This field is non-authoritative, and will only manage the labels present in your configuration.
	// Please refer to the field effective_labels for all of the labels present on the resource.
	// +kubebuilder:validation:Optional
	// +mapType=granular
	Labels map[string]*string `json:"labels,omitempty" tf:"labels,omitempty"`

	// The Certificate Manager location. If not specified, "global" is used.
	// +kubebuilder:validation:Optional
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	// +kubebuilder:validation:Optional
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// type of DNS authorization. If unset during the resource creation, FIXED_RECORD will
	// be used for global resources, and PER_PROJECT_RECORD will be used for other locations.
	// FIXED_RECORD DNS authorization uses DNS-01 validation method
	// PER_PROJECT_RECORD DNS authorization allows for independent management
	// of Google-managed certificates with DNS authorization across multiple
	// projects.
	// Possible values are: FIXED_RECORD, PER_PROJECT_RECORD.
	// +kubebuilder:validation:Optional
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type DNSResourceRecordInitParameters struct {
}

type DNSResourceRecordObservation struct {

	// (Output)
	// Data of the DNS Resource Record.
	Data *string `json:"data,omitempty" tf:"data,omitempty"`

	// (Output)
	// Fully qualified name of the DNS Resource Record.
	// E.g. _acme-challenge.example.com.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// (Output)
	// Type of the DNS Resource Record.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type DNSResourceRecordParameters struct {
}

// DNSAuthorizationSpec defines the desired state of DNSAuthorization
type DNSAuthorizationSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     DNSAuthorizationParameters `json:"forProvider"`
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
	InitProvider DNSAuthorizationInitParameters `json:"initProvider,omitempty"`
}

// DNSAuthorizationStatus defines the observed state of DNSAuthorization.
type DNSAuthorizationStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        DNSAuthorizationObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// DNSAuthorization is the Schema for the DNSAuthorizations API. DnsAuthorization represents a HTTP-reachable backend for a DnsAuthorization.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,gcp}
type DNSAuthorization struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.domain) || (has(self.initProvider) && has(self.initProvider.domain))",message="spec.forProvider.domain is a required parameter"
	Spec   DNSAuthorizationSpec   `json:"spec"`
	Status DNSAuthorizationStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// DNSAuthorizationList contains a list of DNSAuthorizations
type DNSAuthorizationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DNSAuthorization `json:"items"`
}

// Repository type metadata.
var (
	DNSAuthorization_Kind             = "DNSAuthorization"
	DNSAuthorization_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: DNSAuthorization_Kind}.String()
	DNSAuthorization_KindAPIVersion   = DNSAuthorization_Kind + "." + CRDGroupVersion.String()
	DNSAuthorization_GroupVersionKind = CRDGroupVersion.WithKind(DNSAuthorization_Kind)
)

func init() {
	SchemeBuilder.Register(&DNSAuthorization{}, &DNSAuthorizationList{})
}
