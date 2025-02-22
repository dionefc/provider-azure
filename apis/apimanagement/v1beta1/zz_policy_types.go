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

type PolicyInitParameters_2 struct {

	// The XML Content for this Policy as a string.
	XMLContent *string `json:"xmlContent,omitempty" tf:"xml_content,omitempty"`

	// A link to a Policy XML Document, which must be publicly available.
	XMLLink *string `json:"xmlLink,omitempty" tf:"xml_link,omitempty"`
}

type PolicyObservation_2 struct {

	// The ID of the API Management service. Changing this forces a new API Management service Policy to be created.
	APIManagementID *string `json:"apiManagementId,omitempty" tf:"api_management_id,omitempty"`

	// The ID of the API Management service Policy.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The XML Content for this Policy as a string.
	XMLContent *string `json:"xmlContent,omitempty" tf:"xml_content,omitempty"`

	// A link to a Policy XML Document, which must be publicly available.
	XMLLink *string `json:"xmlLink,omitempty" tf:"xml_link,omitempty"`
}

type PolicyParameters_2 struct {

	// The ID of the API Management service. Changing this forces a new API Management service Policy to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/apimanagement/v1beta1.Management
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	APIManagementID *string `json:"apiManagementId,omitempty" tf:"api_management_id,omitempty"`

	// Reference to a Management in apimanagement to populate apiManagementId.
	// +kubebuilder:validation:Optional
	APIManagementIDRef *v1.Reference `json:"apiManagementIdRef,omitempty" tf:"-"`

	// Selector for a Management in apimanagement to populate apiManagementId.
	// +kubebuilder:validation:Optional
	APIManagementIDSelector *v1.Selector `json:"apiManagementIdSelector,omitempty" tf:"-"`

	// The XML Content for this Policy as a string.
	// +kubebuilder:validation:Optional
	XMLContent *string `json:"xmlContent,omitempty" tf:"xml_content,omitempty"`

	// A link to a Policy XML Document, which must be publicly available.
	// +kubebuilder:validation:Optional
	XMLLink *string `json:"xmlLink,omitempty" tf:"xml_link,omitempty"`
}

// PolicySpec defines the desired state of Policy
type PolicySpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     PolicyParameters_2 `json:"forProvider"`
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
	InitProvider PolicyInitParameters_2 `json:"initProvider,omitempty"`
}

// PolicyStatus defines the observed state of Policy.
type PolicyStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        PolicyObservation_2 `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Policy is the Schema for the Policys API. Manages a API Management service Policy.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type Policy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              PolicySpec   `json:"spec"`
	Status            PolicyStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// PolicyList contains a list of Policys
type PolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Policy `json:"items"`
}

// Repository type metadata.
var (
	Policy_Kind             = "Policy"
	Policy_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Policy_Kind}.String()
	Policy_KindAPIVersion   = Policy_Kind + "." + CRDGroupVersion.String()
	Policy_GroupVersionKind = CRDGroupVersion.WithKind(Policy_Kind)
)

func init() {
	SchemeBuilder.Register(&Policy{}, &PolicyList{})
}
