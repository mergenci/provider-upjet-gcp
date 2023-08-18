/*
Copyright 2021 The Crossplane Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by upjet. DO NOT EDIT.

package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type ConditionInitParameters struct {
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	Expression *string `json:"expression,omitempty" tf:"expression,omitempty"`

	Title *string `json:"title,omitempty" tf:"title,omitempty"`
}

type ConditionObservation struct {
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	Expression *string `json:"expression,omitempty" tf:"expression,omitempty"`

	Title *string `json:"title,omitempty" tf:"title,omitempty"`
}

type ConditionParameters struct {

	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// +kubebuilder:validation:Optional
	Expression *string `json:"expression" tf:"expression,omitempty"`

	// +kubebuilder:validation:Optional
	Title *string `json:"title" tf:"title,omitempty"`
}

type ServiceIAMMemberInitParameters struct {
	Condition []ConditionInitParameters `json:"condition,omitempty" tf:"condition,omitempty"`

	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	Member *string `json:"member,omitempty" tf:"member,omitempty"`

	Role *string `json:"role,omitempty" tf:"role,omitempty"`
}

type ServiceIAMMemberObservation struct {
	Condition []ConditionObservation `json:"condition,omitempty" tf:"condition,omitempty"`

	Etag *string `json:"etag,omitempty" tf:"etag,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	Member *string `json:"member,omitempty" tf:"member,omitempty"`

	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	Role *string `json:"role,omitempty" tf:"role,omitempty"`

	Service *string `json:"service,omitempty" tf:"service,omitempty"`
}

type ServiceIAMMemberParameters struct {

	// +kubebuilder:validation:Optional
	Condition []ConditionParameters `json:"condition,omitempty" tf:"condition,omitempty"`

	// +kubebuilder:validation:Optional
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// +kubebuilder:validation:Optional
	Member *string `json:"member,omitempty" tf:"member,omitempty"`

	// +crossplane:generate:reference:type=github.com/upbound/provider-gcp/apis/cloudplatform/v1beta1.Project
	// +kubebuilder:validation:Optional
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// Reference to a Project in cloudplatform to populate project.
	// +kubebuilder:validation:Optional
	ProjectRef *v1.Reference `json:"projectRef,omitempty" tf:"-"`

	// Selector for a Project in cloudplatform to populate project.
	// +kubebuilder:validation:Optional
	ProjectSelector *v1.Selector `json:"projectSelector,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	Role *string `json:"role,omitempty" tf:"role,omitempty"`

	// +crossplane:generate:reference:type=Service
	// +kubebuilder:validation:Optional
	Service *string `json:"service,omitempty" tf:"service,omitempty"`

	// Reference to a Service to populate service.
	// +kubebuilder:validation:Optional
	ServiceRef *v1.Reference `json:"serviceRef,omitempty" tf:"-"`

	// Selector for a Service to populate service.
	// +kubebuilder:validation:Optional
	ServiceSelector *v1.Selector `json:"serviceSelector,omitempty" tf:"-"`
}

// ServiceIAMMemberSpec defines the desired state of ServiceIAMMember
type ServiceIAMMemberSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ServiceIAMMemberParameters `json:"forProvider"`
	// THIS IS AN ALPHA FIELD. Do not use it in production. It is not honored
	// unless the relevant Crossplane feature flag is enabled, and may be
	// changed or removed without notice.
	// InitProvider holds the same fields as ForProvider, with the exception
	// of Identifier and other resource reference fields. The fields that are
	// in InitProvider are merged into ForProvider when the resource is created.
	// The same fields are also added to the terraform ignore_changes hook, to
	// avoid updating them after creation. This is useful for fields that are
	// required on creation, but we do not desire to update them after creation,
	// for example because of an external controller is managing them, like an
	// autoscaler.
	InitProvider ServiceIAMMemberInitParameters `json:"initProvider,omitempty"`
}

// ServiceIAMMemberStatus defines the observed state of ServiceIAMMember.
type ServiceIAMMemberStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ServiceIAMMemberObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// ServiceIAMMember is the Schema for the ServiceIAMMembers API. <no value>
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,gcp}
type ServiceIAMMember struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.location) || has(self.initProvider.location)",message="location is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.member) || has(self.initProvider.member)",message="member is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.role) || has(self.initProvider.role)",message="role is a required parameter"
	Spec   ServiceIAMMemberSpec   `json:"spec"`
	Status ServiceIAMMemberStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ServiceIAMMemberList contains a list of ServiceIAMMembers
type ServiceIAMMemberList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ServiceIAMMember `json:"items"`
}

// Repository type metadata.
var (
	ServiceIAMMember_Kind             = "ServiceIAMMember"
	ServiceIAMMember_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: ServiceIAMMember_Kind}.String()
	ServiceIAMMember_KindAPIVersion   = ServiceIAMMember_Kind + "." + CRDGroupVersion.String()
	ServiceIAMMember_GroupVersionKind = CRDGroupVersion.WithKind(ServiceIAMMember_Kind)
)

func init() {
	SchemeBuilder.Register(&ServiceIAMMember{}, &ServiceIAMMemberList{})
}
