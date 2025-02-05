// SPDX-FileCopyrightText: 2023 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

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

type AlertInitParameters struct {

	// The comparison operator used against the alert's threshold. Must be one of greater_than or less_than.
	// The comparison operator used against the alert's threshold. Enum: 'greater_than' 'less_than
	Comparison *string `json:"comparison,omitempty" tf:"comparison,omitempty"`

	// A human-friendly display name.
	// A human-friendly display name for the alert.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The notification settings for a trigger alert.
	// The notification settings for a trigger alert.
	Notifications []NotificationsInitParameters `json:"notifications,omitempty" tf:"notifications,omitempty"`

	// Period of time the threshold must be exceeded to trigger the alert. Must be one of 2m, 3m, 5m, 10m, 15m, 30m or 1h.
	// Period of time the threshold must be exceeded to trigger the alert. Enum '2m' '3m' '5m' '10m' '15m' '30m' '1h'
	Period *string `json:"period,omitempty" tf:"period,omitempty"`

	// The threshold at which the alert will enter a trigger state. The specific threshold is dependent on the alert type.
	// The threshold at which the alert will enter a trigger state. The specific threshold is dependent on the alert type.
	Threshold *float64 `json:"threshold,omitempty" tf:"threshold,omitempty"`

	// The type of health check to perform. Must be one of latency, down, down_global or ssl_expiry.
	// The type of health check to perform. Enum: 'latency' 'down' 'down_global' 'ssl_expiry'
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type AlertObservation struct {

	// A unique identifier for a check
	// A unique identifier for a check.
	CheckID *string `json:"checkId,omitempty" tf:"check_id,omitempty"`

	// The comparison operator used against the alert's threshold. Must be one of greater_than or less_than.
	// The comparison operator used against the alert's threshold. Enum: 'greater_than' 'less_than
	Comparison *string `json:"comparison,omitempty" tf:"comparison,omitempty"`

	// The id of the alert.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// A human-friendly display name.
	// A human-friendly display name for the alert.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The notification settings for a trigger alert.
	// The notification settings for a trigger alert.
	Notifications []NotificationsObservation `json:"notifications,omitempty" tf:"notifications,omitempty"`

	// Period of time the threshold must be exceeded to trigger the alert. Must be one of 2m, 3m, 5m, 10m, 15m, 30m or 1h.
	// Period of time the threshold must be exceeded to trigger the alert. Enum '2m' '3m' '5m' '10m' '15m' '30m' '1h'
	Period *string `json:"period,omitempty" tf:"period,omitempty"`

	// The threshold at which the alert will enter a trigger state. The specific threshold is dependent on the alert type.
	// The threshold at which the alert will enter a trigger state. The specific threshold is dependent on the alert type.
	Threshold *float64 `json:"threshold,omitempty" tf:"threshold,omitempty"`

	// The type of health check to perform. Must be one of latency, down, down_global or ssl_expiry.
	// The type of health check to perform. Enum: 'latency' 'down' 'down_global' 'ssl_expiry'
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type AlertParameters struct {

	// A unique identifier for a check
	// A unique identifier for a check.
	// +crossplane:generate:reference:type=Check
	// +kubebuilder:validation:Optional
	CheckID *string `json:"checkId,omitempty" tf:"check_id,omitempty"`

	// Reference to a Check to populate checkId.
	// +kubebuilder:validation:Optional
	CheckIDRef *v1.Reference `json:"checkIdRef,omitempty" tf:"-"`

	// Selector for a Check to populate checkId.
	// +kubebuilder:validation:Optional
	CheckIDSelector *v1.Selector `json:"checkIdSelector,omitempty" tf:"-"`

	// The comparison operator used against the alert's threshold. Must be one of greater_than or less_than.
	// The comparison operator used against the alert's threshold. Enum: 'greater_than' 'less_than
	// +kubebuilder:validation:Optional
	Comparison *string `json:"comparison,omitempty" tf:"comparison,omitempty"`

	// A human-friendly display name.
	// A human-friendly display name for the alert.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The notification settings for a trigger alert.
	// The notification settings for a trigger alert.
	// +kubebuilder:validation:Optional
	Notifications []NotificationsParameters `json:"notifications,omitempty" tf:"notifications,omitempty"`

	// Period of time the threshold must be exceeded to trigger the alert. Must be one of 2m, 3m, 5m, 10m, 15m, 30m or 1h.
	// Period of time the threshold must be exceeded to trigger the alert. Enum '2m' '3m' '5m' '10m' '15m' '30m' '1h'
	// +kubebuilder:validation:Optional
	Period *string `json:"period,omitempty" tf:"period,omitempty"`

	// The threshold at which the alert will enter a trigger state. The specific threshold is dependent on the alert type.
	// The threshold at which the alert will enter a trigger state. The specific threshold is dependent on the alert type.
	// +kubebuilder:validation:Optional
	Threshold *float64 `json:"threshold,omitempty" tf:"threshold,omitempty"`

	// The type of health check to perform. Must be one of latency, down, down_global or ssl_expiry.
	// The type of health check to perform. Enum: 'latency' 'down' 'down_global' 'ssl_expiry'
	// +kubebuilder:validation:Optional
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type NotificationsInitParameters struct {

	// List of email addresses to sent notifications to.
	// List of email addresses to sent notifications to
	Email []*string `json:"email,omitempty" tf:"email,omitempty"`

	Slack []SlackInitParameters `json:"slack,omitempty" tf:"slack,omitempty"`
}

type NotificationsObservation struct {

	// List of email addresses to sent notifications to.
	// List of email addresses to sent notifications to
	Email []*string `json:"email,omitempty" tf:"email,omitempty"`

	Slack []SlackObservation `json:"slack,omitempty" tf:"slack,omitempty"`
}

type NotificationsParameters struct {

	// List of email addresses to sent notifications to.
	// List of email addresses to sent notifications to
	// +kubebuilder:validation:Optional
	Email []*string `json:"email,omitempty" tf:"email,omitempty"`

	// +kubebuilder:validation:Optional
	Slack []SlackParameters `json:"slack,omitempty" tf:"slack,omitempty"`
}

type SlackInitParameters struct {

	// The Slack channel to send alerts to.
	// The Slack channel to send alerts to
	Channel *string `json:"channel,omitempty" tf:"channel,omitempty"`

	// The webhook URL for Slack.
	// The webhook URL for Slack
	URL *string `json:"url,omitempty" tf:"url,omitempty"`
}

type SlackObservation struct {

	// The Slack channel to send alerts to.
	// The Slack channel to send alerts to
	Channel *string `json:"channel,omitempty" tf:"channel,omitempty"`

	// The webhook URL for Slack.
	// The webhook URL for Slack
	URL *string `json:"url,omitempty" tf:"url,omitempty"`
}

type SlackParameters struct {

	// The Slack channel to send alerts to.
	// The Slack channel to send alerts to
	// +kubebuilder:validation:Optional
	Channel *string `json:"channel" tf:"channel,omitempty"`

	// The webhook URL for Slack.
	// The webhook URL for Slack
	// +kubebuilder:validation:Optional
	URL *string `json:"url" tf:"url,omitempty"`
}

// AlertSpec defines the desired state of Alert
type AlertSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     AlertParameters `json:"forProvider"`
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
	InitProvider AlertInitParameters `json:"initProvider,omitempty"`
}

// AlertStatus defines the observed state of Alert.
type AlertStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        AlertObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Alert is the Schema for the Alerts API.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,do}
type Alert struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || (has(self.initProvider) && has(self.initProvider.name))",message="spec.forProvider.name is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.notifications) || (has(self.initProvider) && has(self.initProvider.notifications))",message="spec.forProvider.notifications is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.type) || (has(self.initProvider) && has(self.initProvider.type))",message="spec.forProvider.type is a required parameter"
	Spec   AlertSpec   `json:"spec"`
	Status AlertStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// AlertList contains a list of Alerts
type AlertList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Alert `json:"items"`
}

// Repository type metadata.
var (
	Alert_Kind             = "Alert"
	Alert_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Alert_Kind}.String()
	Alert_KindAPIVersion   = Alert_Kind + "." + CRDGroupVersion.String()
	Alert_GroupVersionKind = CRDGroupVersion.WithKind(Alert_Kind)
)

func init() {
	SchemeBuilder.Register(&Alert{}, &AlertList{})
}
