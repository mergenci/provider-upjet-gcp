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

type BigqueryDateShardedSpecInitParameters struct {
}

type BigqueryDateShardedSpecObservation struct {

	// (Output)
	// The Data Catalog resource name of the dataset entry the current table belongs to, for example,
	// projects/{project_id}/locations/{location}/entrygroups/{entryGroupId}/entries/{entryId}
	Dataset *string `json:"dataset,omitempty" tf:"dataset,omitempty"`

	// (Output)
	// Total number of shards.
	ShardCount *float64 `json:"shardCount,omitempty" tf:"shard_count,omitempty"`

	// (Output)
	// The table name prefix of the shards. The name of any given shard is [tablePrefix]YYYYMMDD,
	// for example, for shard MyTable20180101, the tablePrefix is MyTable.
	TablePrefix *string `json:"tablePrefix,omitempty" tf:"table_prefix,omitempty"`
}

type BigqueryDateShardedSpecParameters struct {
}

type BigqueryTableSpecInitParameters struct {
}

type BigqueryTableSpecObservation struct {

	// (Output)
	// The table source type.
	TableSourceType *string `json:"tableSourceType,omitempty" tf:"table_source_type,omitempty"`

	// (Output)
	// Spec of a BigQuery table. This field should only be populated if tableSourceType is BIGQUERY_TABLE.
	// Structure is documented below.
	TableSpec []TableSpecObservation `json:"tableSpec,omitempty" tf:"table_spec,omitempty"`

	// (Output)
	// Table view specification. This field should only be populated if tableSourceType is BIGQUERY_VIEW.
	// Structure is documented below.
	ViewSpec []ViewSpecObservation `json:"viewSpec,omitempty" tf:"view_spec,omitempty"`
}

type BigqueryTableSpecParameters struct {
}

type EntryInitParameters struct {

	// Entry description, which can consist of several sentences or paragraphs that describe entry contents.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Display information such as title and description. A short name to identify the entry,
	// for example, "Analytics Data - Jan 2011".
	DisplayName *string `json:"displayName,omitempty" tf:"display_name,omitempty"`

	// The id of the entry to create.
	EntryID *string `json:"entryId,omitempty" tf:"entry_id,omitempty"`

	// Specification that applies to a Cloud Storage fileset. This is only valid on entries of type FILESET.
	// Structure is documented below.
	GcsFilesetSpec []GcsFilesetSpecInitParameters `json:"gcsFilesetSpec,omitempty" tf:"gcs_fileset_spec,omitempty"`

	// The resource this metadata entry refers to.
	// For Google Cloud Platform resources, linkedResource is the full name of the resource.
	// For example, the linkedResource for a table resource from BigQuery is:
	// //bigquery.googleapis.com/projects/projectId/datasets/datasetId/tables/tableId
	// Output only when Entry is of type in the EntryType enum. For entries with userSpecifiedType,
	// this field is optional and defaults to an empty string.
	LinkedResource *string `json:"linkedResource,omitempty" tf:"linked_resource,omitempty"`

	// Schema of the entry (e.g. BigQuery, GoogleSQL, Avro schema), as a json string. An entry might not have any schema
	// attached to it. See
	// https://cloud.google.com/data-catalog/docs/reference/rest/v1/projects.locations.entryGroups.entries#schema
	// for what fields this schema can contain.
	Schema *string `json:"schema,omitempty" tf:"schema,omitempty"`

	// The type of the entry. Only used for Entries with types in the EntryType enum.
	// Currently, only FILESET enum value is allowed. All other entries created through Data Catalog must use userSpecifiedType.
	// Possible values are: FILESET.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`

	// This field indicates the entry's source system that Data Catalog does not integrate with.
	// userSpecifiedSystem strings must begin with a letter or underscore and can only contain letters, numbers,
	// and underscores; are case insensitive; must be at least 1 character and at most 64 characters long.
	UserSpecifiedSystem *string `json:"userSpecifiedSystem,omitempty" tf:"user_specified_system,omitempty"`

	// Entry type if it does not fit any of the input-allowed values listed in EntryType enum above.
	// When creating an entry, users should check the enum values first, if nothing matches the entry
	// to be created, then provide a custom value, for example "my_special_type".
	// userSpecifiedType strings must begin with a letter or underscore and can only contain letters,
	// numbers, and underscores; are case insensitive; must be at least 1 character and at most 64 characters long.
	UserSpecifiedType *string `json:"userSpecifiedType,omitempty" tf:"user_specified_type,omitempty"`
}

type EntryObservation struct {

	// Specification for a group of BigQuery tables with name pattern [prefix]YYYYMMDD.
	// Context: https://cloud.google.com/bigquery/docs/partitioned-tables#partitioning_versus_sharding.
	// Structure is documented below.
	BigqueryDateShardedSpec []BigqueryDateShardedSpecObservation `json:"bigqueryDateShardedSpec,omitempty" tf:"bigquery_date_sharded_spec,omitempty"`

	// Specification that applies to a BigQuery table. This is only valid on entries of type TABLE.
	// Structure is documented below.
	BigqueryTableSpec []BigqueryTableSpecObservation `json:"bigqueryTableSpec,omitempty" tf:"bigquery_table_spec,omitempty"`

	// Entry description, which can consist of several sentences or paragraphs that describe entry contents.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Display information such as title and description. A short name to identify the entry,
	// for example, "Analytics Data - Jan 2011".
	DisplayName *string `json:"displayName,omitempty" tf:"display_name,omitempty"`

	// The name of the entry group this entry is in.
	EntryGroup *string `json:"entryGroup,omitempty" tf:"entry_group,omitempty"`

	// The id of the entry to create.
	EntryID *string `json:"entryId,omitempty" tf:"entry_id,omitempty"`

	// Specification that applies to a Cloud Storage fileset. This is only valid on entries of type FILESET.
	// Structure is documented below.
	GcsFilesetSpec []GcsFilesetSpecObservation `json:"gcsFilesetSpec,omitempty" tf:"gcs_fileset_spec,omitempty"`

	// an identifier for the resource with format {{name}}
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// This field indicates the entry's source system that Data Catalog integrates with, such as BigQuery or Pub/Sub.
	IntegratedSystem *string `json:"integratedSystem,omitempty" tf:"integrated_system,omitempty"`

	// The resource this metadata entry refers to.
	// For Google Cloud Platform resources, linkedResource is the full name of the resource.
	// For example, the linkedResource for a table resource from BigQuery is:
	// //bigquery.googleapis.com/projects/projectId/datasets/datasetId/tables/tableId
	// Output only when Entry is of type in the EntryType enum. For entries with userSpecifiedType,
	// this field is optional and defaults to an empty string.
	LinkedResource *string `json:"linkedResource,omitempty" tf:"linked_resource,omitempty"`

	// The Data Catalog resource name of the entry in URL format.
	// Example: projects/{project_id}/locations/{location}/entryGroups/{entryGroupId}/entries/{entryId}.
	// Note that this Entry and its child resources may not actually be stored in the location in this name.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Schema of the entry (e.g. BigQuery, GoogleSQL, Avro schema), as a json string. An entry might not have any schema
	// attached to it. See
	// https://cloud.google.com/data-catalog/docs/reference/rest/v1/projects.locations.entryGroups.entries#schema
	// for what fields this schema can contain.
	Schema *string `json:"schema,omitempty" tf:"schema,omitempty"`

	// The type of the entry. Only used for Entries with types in the EntryType enum.
	// Currently, only FILESET enum value is allowed. All other entries created through Data Catalog must use userSpecifiedType.
	// Possible values are: FILESET.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`

	// This field indicates the entry's source system that Data Catalog does not integrate with.
	// userSpecifiedSystem strings must begin with a letter or underscore and can only contain letters, numbers,
	// and underscores; are case insensitive; must be at least 1 character and at most 64 characters long.
	UserSpecifiedSystem *string `json:"userSpecifiedSystem,omitempty" tf:"user_specified_system,omitempty"`

	// Entry type if it does not fit any of the input-allowed values listed in EntryType enum above.
	// When creating an entry, users should check the enum values first, if nothing matches the entry
	// to be created, then provide a custom value, for example "my_special_type".
	// userSpecifiedType strings must begin with a letter or underscore and can only contain letters,
	// numbers, and underscores; are case insensitive; must be at least 1 character and at most 64 characters long.
	UserSpecifiedType *string `json:"userSpecifiedType,omitempty" tf:"user_specified_type,omitempty"`
}

type EntryParameters struct {

	// Entry description, which can consist of several sentences or paragraphs that describe entry contents.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Display information such as title and description. A short name to identify the entry,
	// for example, "Analytics Data - Jan 2011".
	// +kubebuilder:validation:Optional
	DisplayName *string `json:"displayName,omitempty" tf:"display_name,omitempty"`

	// The name of the entry group this entry is in.
	// +crossplane:generate:reference:type=github.com/upbound/provider-gcp/apis/datacatalog/v1beta1.EntryGroup
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	EntryGroup *string `json:"entryGroup,omitempty" tf:"entry_group,omitempty"`

	// Reference to a EntryGroup in datacatalog to populate entryGroup.
	// +kubebuilder:validation:Optional
	EntryGroupRef *v1.Reference `json:"entryGroupRef,omitempty" tf:"-"`

	// Selector for a EntryGroup in datacatalog to populate entryGroup.
	// +kubebuilder:validation:Optional
	EntryGroupSelector *v1.Selector `json:"entryGroupSelector,omitempty" tf:"-"`

	// The id of the entry to create.
	// +kubebuilder:validation:Optional
	EntryID *string `json:"entryId,omitempty" tf:"entry_id,omitempty"`

	// Specification that applies to a Cloud Storage fileset. This is only valid on entries of type FILESET.
	// Structure is documented below.
	// +kubebuilder:validation:Optional
	GcsFilesetSpec []GcsFilesetSpecParameters `json:"gcsFilesetSpec,omitempty" tf:"gcs_fileset_spec,omitempty"`

	// The resource this metadata entry refers to.
	// For Google Cloud Platform resources, linkedResource is the full name of the resource.
	// For example, the linkedResource for a table resource from BigQuery is:
	// //bigquery.googleapis.com/projects/projectId/datasets/datasetId/tables/tableId
	// Output only when Entry is of type in the EntryType enum. For entries with userSpecifiedType,
	// this field is optional and defaults to an empty string.
	// +kubebuilder:validation:Optional
	LinkedResource *string `json:"linkedResource,omitempty" tf:"linked_resource,omitempty"`

	// Schema of the entry (e.g. BigQuery, GoogleSQL, Avro schema), as a json string. An entry might not have any schema
	// attached to it. See
	// https://cloud.google.com/data-catalog/docs/reference/rest/v1/projects.locations.entryGroups.entries#schema
	// for what fields this schema can contain.
	// +kubebuilder:validation:Optional
	Schema *string `json:"schema,omitempty" tf:"schema,omitempty"`

	// The type of the entry. Only used for Entries with types in the EntryType enum.
	// Currently, only FILESET enum value is allowed. All other entries created through Data Catalog must use userSpecifiedType.
	// Possible values are: FILESET.
	// +kubebuilder:validation:Optional
	Type *string `json:"type,omitempty" tf:"type,omitempty"`

	// This field indicates the entry's source system that Data Catalog does not integrate with.
	// userSpecifiedSystem strings must begin with a letter or underscore and can only contain letters, numbers,
	// and underscores; are case insensitive; must be at least 1 character and at most 64 characters long.
	// +kubebuilder:validation:Optional
	UserSpecifiedSystem *string `json:"userSpecifiedSystem,omitempty" tf:"user_specified_system,omitempty"`

	// Entry type if it does not fit any of the input-allowed values listed in EntryType enum above.
	// When creating an entry, users should check the enum values first, if nothing matches the entry
	// to be created, then provide a custom value, for example "my_special_type".
	// userSpecifiedType strings must begin with a letter or underscore and can only contain letters,
	// numbers, and underscores; are case insensitive; must be at least 1 character and at most 64 characters long.
	// +kubebuilder:validation:Optional
	UserSpecifiedType *string `json:"userSpecifiedType,omitempty" tf:"user_specified_type,omitempty"`
}

type GcsFilesetSpecInitParameters struct {

	// Patterns to identify a set of files in Google Cloud Storage.
	// See Cloud Storage documentation
	// for more information. Note that bucket wildcards are currently not supported. Examples of valid filePatterns:
	FilePatterns []*string `json:"filePatterns,omitempty" tf:"file_patterns,omitempty"`
}

type GcsFilesetSpecObservation struct {

	// Patterns to identify a set of files in Google Cloud Storage.
	// See Cloud Storage documentation
	// for more information. Note that bucket wildcards are currently not supported. Examples of valid filePatterns:
	FilePatterns []*string `json:"filePatterns,omitempty" tf:"file_patterns,omitempty"`

	// (Output)
	// Sample files contained in this fileset, not all files contained in this fileset are represented here.
	// Structure is documented below.
	SampleGcsFileSpecs []SampleGcsFileSpecsObservation `json:"sampleGcsFileSpecs,omitempty" tf:"sample_gcs_file_specs,omitempty"`
}

type GcsFilesetSpecParameters struct {

	// Patterns to identify a set of files in Google Cloud Storage.
	// See Cloud Storage documentation
	// for more information. Note that bucket wildcards are currently not supported. Examples of valid filePatterns:
	// +kubebuilder:validation:Optional
	FilePatterns []*string `json:"filePatterns" tf:"file_patterns,omitempty"`
}

type SampleGcsFileSpecsInitParameters struct {
}

type SampleGcsFileSpecsObservation struct {

	// (Output)
	// The full file path
	FilePath *string `json:"filePath,omitempty" tf:"file_path,omitempty"`

	// (Output)
	// The size of the file, in bytes.
	SizeBytes *float64 `json:"sizeBytes,omitempty" tf:"size_bytes,omitempty"`
}

type SampleGcsFileSpecsParameters struct {
}

type TableSpecInitParameters struct {
}

type TableSpecObservation struct {

	// (Output)
	// If the table is a dated shard, i.e., with name pattern [prefix]YYYYMMDD, groupedEntry is the
	// Data Catalog resource name of the date sharded grouped entry, for example,
	// projects/{project_id}/locations/{location}/entrygroups/{entryGroupId}/entries/{entryId}.
	// Otherwise, groupedEntry is empty.
	GroupedEntry *string `json:"groupedEntry,omitempty" tf:"grouped_entry,omitempty"`
}

type TableSpecParameters struct {
}

type ViewSpecInitParameters struct {
}

type ViewSpecObservation struct {

	// (Output)
	// The query that defines the table view.
	ViewQuery *string `json:"viewQuery,omitempty" tf:"view_query,omitempty"`
}

type ViewSpecParameters struct {
}

// EntrySpec defines the desired state of Entry
type EntrySpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     EntryParameters `json:"forProvider"`
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
	InitProvider EntryInitParameters `json:"initProvider,omitempty"`
}

// EntryStatus defines the observed state of Entry.
type EntryStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        EntryObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Entry is the Schema for the Entrys API. Entry Metadata.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,gcp}
type Entry struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.entryId) || has(self.initProvider.entryId)",message="entryId is a required parameter"
	Spec   EntrySpec   `json:"spec"`
	Status EntryStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// EntryList contains a list of Entrys
type EntryList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Entry `json:"items"`
}

// Repository type metadata.
var (
	Entry_Kind             = "Entry"
	Entry_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Entry_Kind}.String()
	Entry_KindAPIVersion   = Entry_Kind + "." + CRDGroupVersion.String()
	Entry_GroupVersionKind = CRDGroupVersion.WithKind(Entry_Kind)
)

func init() {
	SchemeBuilder.Register(&Entry{}, &EntryList{})
}
