/*
Copyright 2022 Upbound Inc.
*/

// Code generated by upjet. DO NOT EDIT.

package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type ApplicationInsightsAnalyticsItemInitParameters struct {

	// The content for the Analytics Item, for example the query text if type is query.
	Content *string `json:"content,omitempty" tf:"content,omitempty"`

	// The alias to use for the function. Required when type is function.
	FunctionAlias *string `json:"functionAlias,omitempty" tf:"function_alias,omitempty"`

	// Specifies the name of the Application Insights Analytics Item. Changing this forces a new resource to be created.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The scope for the Analytics Item. Can be shared or user. Changing this forces a new resource to be created. Must be shared for functions.
	Scope *string `json:"scope,omitempty" tf:"scope,omitempty"`

	// The type of Analytics Item to create. Can be one of query, function, folder, recent. Changing this forces a new resource to be created.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type ApplicationInsightsAnalyticsItemObservation struct {

	// The ID of the Application Insights component on which the Analytics Item exists. Changing this forces a new resource to be created.
	ApplicationInsightsID *string `json:"applicationInsightsId,omitempty" tf:"application_insights_id,omitempty"`

	// The content for the Analytics Item, for example the query text if type is query.
	Content *string `json:"content,omitempty" tf:"content,omitempty"`

	// The alias to use for the function. Required when type is function.
	FunctionAlias *string `json:"functionAlias,omitempty" tf:"function_alias,omitempty"`

	// The ID of the Application Insights Analytics Item.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Specifies the name of the Application Insights Analytics Item. Changing this forces a new resource to be created.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The scope for the Analytics Item. Can be shared or user. Changing this forces a new resource to be created. Must be shared for functions.
	Scope *string `json:"scope,omitempty" tf:"scope,omitempty"`

	// A string containing the time the Analytics Item was created.
	TimeCreated *string `json:"timeCreated,omitempty" tf:"time_created,omitempty"`

	// A string containing the time the Analytics Item was last modified.
	TimeModified *string `json:"timeModified,omitempty" tf:"time_modified,omitempty"`

	// The type of Analytics Item to create. Can be one of query, function, folder, recent. Changing this forces a new resource to be created.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`

	// A string indicating the version of the query format
	Version *string `json:"version,omitempty" tf:"version,omitempty"`
}

type ApplicationInsightsAnalyticsItemParameters struct {

	// The ID of the Application Insights component on which the Analytics Item exists. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/insights/v1beta1.ApplicationInsights
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	ApplicationInsightsID *string `json:"applicationInsightsId,omitempty" tf:"application_insights_id,omitempty"`

	// Reference to a ApplicationInsights in insights to populate applicationInsightsId.
	// +kubebuilder:validation:Optional
	ApplicationInsightsIDRef *v1.Reference `json:"applicationInsightsIdRef,omitempty" tf:"-"`

	// Selector for a ApplicationInsights in insights to populate applicationInsightsId.
	// +kubebuilder:validation:Optional
	ApplicationInsightsIDSelector *v1.Selector `json:"applicationInsightsIdSelector,omitempty" tf:"-"`

	// The content for the Analytics Item, for example the query text if type is query.
	// +kubebuilder:validation:Optional
	Content *string `json:"content,omitempty" tf:"content,omitempty"`

	// The alias to use for the function. Required when type is function.
	// +kubebuilder:validation:Optional
	FunctionAlias *string `json:"functionAlias,omitempty" tf:"function_alias,omitempty"`

	// Specifies the name of the Application Insights Analytics Item. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The scope for the Analytics Item. Can be shared or user. Changing this forces a new resource to be created. Must be shared for functions.
	// +kubebuilder:validation:Optional
	Scope *string `json:"scope,omitempty" tf:"scope,omitempty"`

	// The type of Analytics Item to create. Can be one of query, function, folder, recent. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

// ApplicationInsightsAnalyticsItemSpec defines the desired state of ApplicationInsightsAnalyticsItem
type ApplicationInsightsAnalyticsItemSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ApplicationInsightsAnalyticsItemParameters `json:"forProvider"`
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
	InitProvider ApplicationInsightsAnalyticsItemInitParameters `json:"initProvider,omitempty"`
}

// ApplicationInsightsAnalyticsItemStatus defines the observed state of ApplicationInsightsAnalyticsItem.
type ApplicationInsightsAnalyticsItemStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ApplicationInsightsAnalyticsItemObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// ApplicationInsightsAnalyticsItem is the Schema for the ApplicationInsightsAnalyticsItems API. Manages an Application Insights Analytics Item component.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type ApplicationInsightsAnalyticsItem struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.content) || has(self.initProvider.content)",message="content is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || has(self.initProvider.name)",message="name is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.scope) || has(self.initProvider.scope)",message="scope is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.type) || has(self.initProvider.type)",message="type is a required parameter"
	Spec   ApplicationInsightsAnalyticsItemSpec   `json:"spec"`
	Status ApplicationInsightsAnalyticsItemStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ApplicationInsightsAnalyticsItemList contains a list of ApplicationInsightsAnalyticsItems
type ApplicationInsightsAnalyticsItemList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ApplicationInsightsAnalyticsItem `json:"items"`
}

// Repository type metadata.
var (
	ApplicationInsightsAnalyticsItem_Kind             = "ApplicationInsightsAnalyticsItem"
	ApplicationInsightsAnalyticsItem_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: ApplicationInsightsAnalyticsItem_Kind}.String()
	ApplicationInsightsAnalyticsItem_KindAPIVersion   = ApplicationInsightsAnalyticsItem_Kind + "." + CRDGroupVersion.String()
	ApplicationInsightsAnalyticsItem_GroupVersionKind = CRDGroupVersion.WithKind(ApplicationInsightsAnalyticsItem_Kind)
)

func init() {
	SchemeBuilder.Register(&ApplicationInsightsAnalyticsItem{}, &ApplicationInsightsAnalyticsItemList{})
}
