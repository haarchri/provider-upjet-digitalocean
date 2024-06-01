/*
Copyright 2022 Upbound Inc.
*/

// Code generated by upjet. DO NOT EDIT.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type ProjectInitParameters struct {

	// the description of the project
	// the description of the project
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// the environment of the project's resources. The possible values are: Development, Staging, Production)
	// the environment of the project's resources
	Environment *string `json:"environment,omitempty" tf:"environment,omitempty"`

	// a boolean indicating whether or not the project is the default project. (Default: "false")
	// determine if the project is the default or not.
	IsDefault *bool `json:"isDefault,omitempty" tf:"is_default,omitempty"`

	// The name of the Project
	// the human-readable name for the project
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// the purpose of the project, (Default: "Web Application")
	// the purpose of the project
	Purpose *string `json:"purpose,omitempty" tf:"purpose,omitempty"`

	// a list of uniform resource names (URNs) for the resources associated with the project
	// the resources associated with the project
	// +listType=set
	Resources []*string `json:"resources,omitempty" tf:"resources,omitempty"`
}

type ProjectObservation struct {

	// the date and time when the project was created, (ISO8601)
	// the date and time when the project was created, (ISO8601)
	CreatedAt *string `json:"createdAt,omitempty" tf:"created_at,omitempty"`

	// the description of the project
	// the description of the project
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// the environment of the project's resources. The possible values are: Development, Staging, Production)
	// the environment of the project's resources
	Environment *string `json:"environment,omitempty" tf:"environment,omitempty"`

	// The id of the project
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// a boolean indicating whether or not the project is the default project. (Default: "false")
	// determine if the project is the default or not.
	IsDefault *bool `json:"isDefault,omitempty" tf:"is_default,omitempty"`

	// The name of the Project
	// the human-readable name for the project
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// the id of the project owner.
	// the id of the project owner.
	OwnerID *float64 `json:"ownerId,omitempty" tf:"owner_id,omitempty"`

	// the unique universal identifier of the project owner.
	// the unique universal identifier of the project owner.
	OwnerUUID *string `json:"ownerUuid,omitempty" tf:"owner_uuid,omitempty"`

	// the purpose of the project, (Default: "Web Application")
	// the purpose of the project
	Purpose *string `json:"purpose,omitempty" tf:"purpose,omitempty"`

	// a list of uniform resource names (URNs) for the resources associated with the project
	// the resources associated with the project
	// +listType=set
	Resources []*string `json:"resources,omitempty" tf:"resources,omitempty"`

	// the date and time when the project was last updated, (ISO8601)
	// the date and time when the project was last updated, (ISO8601)
	UpdatedAt *string `json:"updatedAt,omitempty" tf:"updated_at,omitempty"`
}

type ProjectParameters struct {

	// the description of the project
	// the description of the project
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// the environment of the project's resources. The possible values are: Development, Staging, Production)
	// the environment of the project's resources
	// +kubebuilder:validation:Optional
	Environment *string `json:"environment,omitempty" tf:"environment,omitempty"`

	// a boolean indicating whether or not the project is the default project. (Default: "false")
	// determine if the project is the default or not.
	// +kubebuilder:validation:Optional
	IsDefault *bool `json:"isDefault,omitempty" tf:"is_default,omitempty"`

	// The name of the Project
	// the human-readable name for the project
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// the purpose of the project, (Default: "Web Application")
	// the purpose of the project
	// +kubebuilder:validation:Optional
	Purpose *string `json:"purpose,omitempty" tf:"purpose,omitempty"`

	// a list of uniform resource names (URNs) for the resources associated with the project
	// the resources associated with the project
	// +kubebuilder:validation:Optional
	// +listType=set
	Resources []*string `json:"resources,omitempty" tf:"resources,omitempty"`
}

// ProjectSpec defines the desired state of Project
type ProjectSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ProjectParameters `json:"forProvider"`
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
	InitProvider ProjectInitParameters `json:"initProvider,omitempty"`
}

// ProjectStatus defines the observed state of Project.
type ProjectStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ProjectObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// Project is the Schema for the Projects API.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,do}
type Project struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || (has(self.initProvider) && has(self.initProvider.name))",message="spec.forProvider.name is a required parameter"
	Spec   ProjectSpec   `json:"spec"`
	Status ProjectStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ProjectList contains a list of Projects
type ProjectList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Project `json:"items"`
}

// Repository type metadata.
var (
	Project_Kind             = "Project"
	Project_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Project_Kind}.String()
	Project_KindAPIVersion   = Project_Kind + "." + CRDGroupVersion.String()
	Project_GroupVersionKind = CRDGroupVersion.WithKind(Project_Kind)
)

func init() {
	SchemeBuilder.Register(&Project{}, &ProjectList{})
}
