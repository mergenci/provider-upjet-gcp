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

type GuestAcceleratorsInitParameters struct {

	// The number of the guest accelerator cards exposed to
	// this instance.
	AcceleratorCount *float64 `json:"acceleratorCount,omitempty" tf:"accelerator_count,omitempty"`

	// The full or partial URL of the accelerator type to
	// attach to this instance. For example:
	// projects/my-project/zones/us-central1-c/acceleratorTypes/nvidia-tesla-p100
	// If you are creating an instance template, specify only the accelerator name.
	AcceleratorType *string `json:"acceleratorType,omitempty" tf:"accelerator_type,omitempty"`
}

type GuestAcceleratorsObservation struct {

	// The number of the guest accelerator cards exposed to
	// this instance.
	AcceleratorCount *float64 `json:"acceleratorCount,omitempty" tf:"accelerator_count,omitempty"`

	// The full or partial URL of the accelerator type to
	// attach to this instance. For example:
	// projects/my-project/zones/us-central1-c/acceleratorTypes/nvidia-tesla-p100
	// If you are creating an instance template, specify only the accelerator name.
	AcceleratorType *string `json:"acceleratorType,omitempty" tf:"accelerator_type,omitempty"`
}

type GuestAcceleratorsParameters struct {

	// The number of the guest accelerator cards exposed to
	// this instance.
	// +kubebuilder:validation:Optional
	AcceleratorCount *float64 `json:"acceleratorCount" tf:"accelerator_count,omitempty"`

	// The full or partial URL of the accelerator type to
	// attach to this instance. For example:
	// projects/my-project/zones/us-central1-c/acceleratorTypes/nvidia-tesla-p100
	// If you are creating an instance template, specify only the accelerator name.
	// +kubebuilder:validation:Optional
	AcceleratorType *string `json:"acceleratorType" tf:"accelerator_type,omitempty"`
}

type InstancePropertiesInitParameters struct {

	// Guest accelerator type and count.
	// Structure is documented below.
	GuestAccelerators []GuestAcceleratorsInitParameters `json:"guestAccelerators,omitempty" tf:"guest_accelerators,omitempty"`

	// The amount of local ssd to reserve with each instance. This
	// reserves disks of type local-ssd.
	// Structure is documented below.
	LocalSsds []LocalSsdsInitParameters `json:"localSsds,omitempty" tf:"local_ssds,omitempty"`

	// The name of the machine type to reserve.
	MachineType *string `json:"machineType,omitempty" tf:"machine_type,omitempty"`

	// The minimum CPU platform for the reservation. For example,
	// "Intel Skylake". See
	// the CPU platform availability reference](https://cloud.google.com/compute/docs/instances/specify-min-cpu-platform#availablezones)
	// for information on available CPU platforms.
	MinCPUPlatform *string `json:"minCpuPlatform,omitempty" tf:"min_cpu_platform,omitempty"`
}

type InstancePropertiesObservation struct {

	// Guest accelerator type and count.
	// Structure is documented below.
	GuestAccelerators []GuestAcceleratorsObservation `json:"guestAccelerators,omitempty" tf:"guest_accelerators,omitempty"`

	// The amount of local ssd to reserve with each instance. This
	// reserves disks of type local-ssd.
	// Structure is documented below.
	LocalSsds []LocalSsdsObservation `json:"localSsds,omitempty" tf:"local_ssds,omitempty"`

	// The name of the machine type to reserve.
	MachineType *string `json:"machineType,omitempty" tf:"machine_type,omitempty"`

	// The minimum CPU platform for the reservation. For example,
	// "Intel Skylake". See
	// the CPU platform availability reference](https://cloud.google.com/compute/docs/instances/specify-min-cpu-platform#availablezones)
	// for information on available CPU platforms.
	MinCPUPlatform *string `json:"minCpuPlatform,omitempty" tf:"min_cpu_platform,omitempty"`
}

type InstancePropertiesParameters struct {

	// Guest accelerator type and count.
	// Structure is documented below.
	// +kubebuilder:validation:Optional
	GuestAccelerators []GuestAcceleratorsParameters `json:"guestAccelerators,omitempty" tf:"guest_accelerators,omitempty"`

	// The amount of local ssd to reserve with each instance. This
	// reserves disks of type local-ssd.
	// Structure is documented below.
	// +kubebuilder:validation:Optional
	LocalSsds []LocalSsdsParameters `json:"localSsds,omitempty" tf:"local_ssds,omitempty"`

	// The name of the machine type to reserve.
	// +kubebuilder:validation:Optional
	MachineType *string `json:"machineType" tf:"machine_type,omitempty"`

	// The minimum CPU platform for the reservation. For example,
	// "Intel Skylake". See
	// the CPU platform availability reference](https://cloud.google.com/compute/docs/instances/specify-min-cpu-platform#availablezones)
	// for information on available CPU platforms.
	// +kubebuilder:validation:Optional
	MinCPUPlatform *string `json:"minCpuPlatform,omitempty" tf:"min_cpu_platform,omitempty"`
}

type LocalSsdsInitParameters struct {

	// The size of the disk in base-2 GB.
	DiskSizeGb *float64 `json:"diskSizeGb,omitempty" tf:"disk_size_gb,omitempty"`

	// The disk interface to use for attaching this disk.
	// Default value is SCSI.
	// Possible values are: SCSI, NVME.
	Interface *string `json:"interface,omitempty" tf:"interface,omitempty"`
}

type LocalSsdsObservation struct {

	// The size of the disk in base-2 GB.
	DiskSizeGb *float64 `json:"diskSizeGb,omitempty" tf:"disk_size_gb,omitempty"`

	// The disk interface to use for attaching this disk.
	// Default value is SCSI.
	// Possible values are: SCSI, NVME.
	Interface *string `json:"interface,omitempty" tf:"interface,omitempty"`
}

type LocalSsdsParameters struct {

	// The size of the disk in base-2 GB.
	// +kubebuilder:validation:Optional
	DiskSizeGb *float64 `json:"diskSizeGb" tf:"disk_size_gb,omitempty"`

	// The disk interface to use for attaching this disk.
	// Default value is SCSI.
	// Possible values are: SCSI, NVME.
	// +kubebuilder:validation:Optional
	Interface *string `json:"interface,omitempty" tf:"interface,omitempty"`
}

type ReservationInitParameters struct {

	// An optional description of this resource.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// The share setting for reservations.
	// Structure is documented below.
	ShareSettings []ReservationShareSettingsInitParameters `json:"shareSettings,omitempty" tf:"share_settings,omitempty"`

	// Reservation for instances with specific machine shapes.
	// Structure is documented below.
	SpecificReservation []ReservationSpecificReservationInitParameters `json:"specificReservation,omitempty" tf:"specific_reservation,omitempty"`

	// When set to true, only VMs that target this reservation by name can
	// consume this reservation. Otherwise, it can be consumed by VMs with
	// affinity for any reservation. Defaults to false.
	SpecificReservationRequired *bool `json:"specificReservationRequired,omitempty" tf:"specific_reservation_required,omitempty"`
}

type ReservationObservation struct {

	// Full or partial URL to a parent commitment. This field displays for
	// reservations that are tied to a commitment.
	Commitment *string `json:"commitment,omitempty" tf:"commitment,omitempty"`

	// Creation timestamp in RFC3339 text format.
	CreationTimestamp *string `json:"creationTimestamp,omitempty" tf:"creation_timestamp,omitempty"`

	// An optional description of this resource.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// an identifier for the resource with format projects/{{project}}/zones/{{zone}}/reservations/{{name}}
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// The URI of the created resource.
	SelfLink *string `json:"selfLink,omitempty" tf:"self_link,omitempty"`

	// The share setting for reservations.
	// Structure is documented below.
	ShareSettings []ReservationShareSettingsObservation `json:"shareSettings,omitempty" tf:"share_settings,omitempty"`

	// Reservation for instances with specific machine shapes.
	// Structure is documented below.
	SpecificReservation []ReservationSpecificReservationObservation `json:"specificReservation,omitempty" tf:"specific_reservation,omitempty"`

	// When set to true, only VMs that target this reservation by name can
	// consume this reservation. Otherwise, it can be consumed by VMs with
	// affinity for any reservation. Defaults to false.
	SpecificReservationRequired *bool `json:"specificReservationRequired,omitempty" tf:"specific_reservation_required,omitempty"`

	// The status of the reservation.
	Status *string `json:"status,omitempty" tf:"status,omitempty"`

	// The zone where the reservation is made.
	Zone *string `json:"zone,omitempty" tf:"zone,omitempty"`
}

type ReservationParameters struct {

	// An optional description of this resource.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	// +kubebuilder:validation:Optional
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// The share setting for reservations.
	// Structure is documented below.
	// +kubebuilder:validation:Optional
	ShareSettings []ReservationShareSettingsParameters `json:"shareSettings,omitempty" tf:"share_settings,omitempty"`

	// Reservation for instances with specific machine shapes.
	// Structure is documented below.
	// +kubebuilder:validation:Optional
	SpecificReservation []ReservationSpecificReservationParameters `json:"specificReservation,omitempty" tf:"specific_reservation,omitempty"`

	// When set to true, only VMs that target this reservation by name can
	// consume this reservation. Otherwise, it can be consumed by VMs with
	// affinity for any reservation. Defaults to false.
	// +kubebuilder:validation:Optional
	SpecificReservationRequired *bool `json:"specificReservationRequired,omitempty" tf:"specific_reservation_required,omitempty"`

	// The zone where the reservation is made.
	// +kubebuilder:validation:Required
	Zone *string `json:"zone" tf:"zone,omitempty"`
}

type ReservationShareSettingsInitParameters struct {

	// A map of project number and project config. This is only valid when shareType's value is SPECIFIC_PROJECTS.
	// Structure is documented below.
	ProjectMap []ShareSettingsProjectMapInitParameters `json:"projectMap,omitempty" tf:"project_map,omitempty"`

	// Type of sharing for this shared-reservation
	// Possible values are: LOCAL, SPECIFIC_PROJECTS.
	ShareType *string `json:"shareType,omitempty" tf:"share_type,omitempty"`
}

type ReservationShareSettingsObservation struct {

	// A map of project number and project config. This is only valid when shareType's value is SPECIFIC_PROJECTS.
	// Structure is documented below.
	ProjectMap []ShareSettingsProjectMapObservation `json:"projectMap,omitempty" tf:"project_map,omitempty"`

	// Type of sharing for this shared-reservation
	// Possible values are: LOCAL, SPECIFIC_PROJECTS.
	ShareType *string `json:"shareType,omitempty" tf:"share_type,omitempty"`
}

type ReservationShareSettingsParameters struct {

	// A map of project number and project config. This is only valid when shareType's value is SPECIFIC_PROJECTS.
	// Structure is documented below.
	// +kubebuilder:validation:Optional
	ProjectMap []ShareSettingsProjectMapParameters `json:"projectMap,omitempty" tf:"project_map,omitempty"`

	// Type of sharing for this shared-reservation
	// Possible values are: LOCAL, SPECIFIC_PROJECTS.
	// +kubebuilder:validation:Optional
	ShareType *string `json:"shareType,omitempty" tf:"share_type,omitempty"`
}

type ReservationSpecificReservationInitParameters struct {

	// The number of resources that are allocated.
	Count *float64 `json:"count,omitempty" tf:"count,omitempty"`

	// The instance properties for the reservation.
	// Structure is documented below.
	InstanceProperties []InstancePropertiesInitParameters `json:"instanceProperties,omitempty" tf:"instance_properties,omitempty"`
}

type ReservationSpecificReservationObservation struct {

	// The number of resources that are allocated.
	Count *float64 `json:"count,omitempty" tf:"count,omitempty"`

	// (Output)
	// How many instances are in use.
	InUseCount *float64 `json:"inUseCount,omitempty" tf:"in_use_count,omitempty"`

	// The instance properties for the reservation.
	// Structure is documented below.
	InstanceProperties []InstancePropertiesObservation `json:"instanceProperties,omitempty" tf:"instance_properties,omitempty"`
}

type ReservationSpecificReservationParameters struct {

	// The number of resources that are allocated.
	// +kubebuilder:validation:Optional
	Count *float64 `json:"count" tf:"count,omitempty"`

	// The instance properties for the reservation.
	// Structure is documented below.
	// +kubebuilder:validation:Optional
	InstanceProperties []InstancePropertiesParameters `json:"instanceProperties" tf:"instance_properties,omitempty"`
}

type ShareSettingsProjectMapInitParameters struct {

	// The identifier for this object. Format specified above.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The project id/number, should be same as the key of this project config in the project map.
	ProjectID *string `json:"projectId,omitempty" tf:"project_id,omitempty"`
}

type ShareSettingsProjectMapObservation struct {

	// The identifier for this object. Format specified above.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The project id/number, should be same as the key of this project config in the project map.
	ProjectID *string `json:"projectId,omitempty" tf:"project_id,omitempty"`
}

type ShareSettingsProjectMapParameters struct {

	// The identifier for this object. Format specified above.
	// +kubebuilder:validation:Optional
	ID *string `json:"id" tf:"id,omitempty"`

	// The project id/number, should be same as the key of this project config in the project map.
	// +kubebuilder:validation:Optional
	ProjectID *string `json:"projectId,omitempty" tf:"project_id,omitempty"`
}

// ReservationSpec defines the desired state of Reservation
type ReservationSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ReservationParameters `json:"forProvider"`
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
	InitProvider ReservationInitParameters `json:"initProvider,omitempty"`
}

// ReservationStatus defines the observed state of Reservation.
type ReservationStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ReservationObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Reservation is the Schema for the Reservations API. Represents a reservation resource.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,gcp}
type Reservation struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.specificReservation) || has(self.initProvider.specificReservation)",message="specificReservation is a required parameter"
	Spec   ReservationSpec   `json:"spec"`
	Status ReservationStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ReservationList contains a list of Reservations
type ReservationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Reservation `json:"items"`
}

// Repository type metadata.
var (
	Reservation_Kind             = "Reservation"
	Reservation_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Reservation_Kind}.String()
	Reservation_KindAPIVersion   = Reservation_Kind + "." + CRDGroupVersion.String()
	Reservation_GroupVersionKind = CRDGroupVersion.WithKind(Reservation_Kind)
)

func init() {
	SchemeBuilder.Register(&Reservation{}, &ReservationList{})
}
