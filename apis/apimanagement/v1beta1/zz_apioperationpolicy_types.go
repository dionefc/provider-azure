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

type APIOperationPolicyInitParameters struct {

	// The XML Content for this Policy.
	XMLContent *string `json:"xmlContent,omitempty" tf:"xml_content,omitempty"`

	// A link to a Policy XML Document, which must be publicly available.
	XMLLink *string `json:"xmlLink,omitempty" tf:"xml_link,omitempty"`
}

type APIOperationPolicyObservation struct {

	// The name of the API Management Service. Changing this forces a new resource to be created.
	APIManagementName *string `json:"apiManagementName,omitempty" tf:"api_management_name,omitempty"`

	// The name of the API within the API Management Service where the Operation exists. Changing this forces a new resource to be created.
	APIName *string `json:"apiName,omitempty" tf:"api_name,omitempty"`

	// The ID of the API Management API Operation Policy.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The operation identifier within an API. Must be unique in the current API Management service instance. Changing this forces a new resource to be created.
	OperationID *string `json:"operationId,omitempty" tf:"operation_id,omitempty"`

	// The name of the Resource Group in which the API Management Service exists. Changing this forces a new resource to be created.
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// The XML Content for this Policy.
	XMLContent *string `json:"xmlContent,omitempty" tf:"xml_content,omitempty"`

	// A link to a Policy XML Document, which must be publicly available.
	XMLLink *string `json:"xmlLink,omitempty" tf:"xml_link,omitempty"`
}

type APIOperationPolicyParameters struct {

	// The name of the API Management Service. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/apimanagement/v1beta1.APIOperation
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractParamPath("api_management_name",false)
	// +kubebuilder:validation:Optional
	APIManagementName *string `json:"apiManagementName,omitempty" tf:"api_management_name,omitempty"`

	// Reference to a APIOperation in apimanagement to populate apiManagementName.
	// +kubebuilder:validation:Optional
	APIManagementNameRef *v1.Reference `json:"apiManagementNameRef,omitempty" tf:"-"`

	// Selector for a APIOperation in apimanagement to populate apiManagementName.
	// +kubebuilder:validation:Optional
	APIManagementNameSelector *v1.Selector `json:"apiManagementNameSelector,omitempty" tf:"-"`

	// The name of the API within the API Management Service where the Operation exists. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/apimanagement/v1beta1.APIOperation
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractParamPath("api_name",false)
	// +kubebuilder:validation:Optional
	APIName *string `json:"apiName,omitempty" tf:"api_name,omitempty"`

	// Reference to a APIOperation in apimanagement to populate apiName.
	// +kubebuilder:validation:Optional
	APINameRef *v1.Reference `json:"apiNameRef,omitempty" tf:"-"`

	// Selector for a APIOperation in apimanagement to populate apiName.
	// +kubebuilder:validation:Optional
	APINameSelector *v1.Selector `json:"apiNameSelector,omitempty" tf:"-"`

	// The operation identifier within an API. Must be unique in the current API Management service instance. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/apimanagement/v1beta1.APIOperation
	// +kubebuilder:validation:Optional
	OperationID *string `json:"operationId,omitempty" tf:"operation_id,omitempty"`

	// Reference to a APIOperation in apimanagement to populate operationId.
	// +kubebuilder:validation:Optional
	OperationIDRef *v1.Reference `json:"operationIdRef,omitempty" tf:"-"`

	// Selector for a APIOperation in apimanagement to populate operationId.
	// +kubebuilder:validation:Optional
	OperationIDSelector *v1.Selector `json:"operationIdSelector,omitempty" tf:"-"`

	// The name of the Resource Group in which the API Management Service exists. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/azure/v1beta1.ResourceGroup
	// +kubebuilder:validation:Optional
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// Reference to a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameRef *v1.Reference `json:"resourceGroupNameRef,omitempty" tf:"-"`

	// Selector for a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameSelector *v1.Selector `json:"resourceGroupNameSelector,omitempty" tf:"-"`

	// The XML Content for this Policy.
	// +kubebuilder:validation:Optional
	XMLContent *string `json:"xmlContent,omitempty" tf:"xml_content,omitempty"`

	// A link to a Policy XML Document, which must be publicly available.
	// +kubebuilder:validation:Optional
	XMLLink *string `json:"xmlLink,omitempty" tf:"xml_link,omitempty"`
}

// APIOperationPolicySpec defines the desired state of APIOperationPolicy
type APIOperationPolicySpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     APIOperationPolicyParameters `json:"forProvider"`
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
	InitProvider APIOperationPolicyInitParameters `json:"initProvider,omitempty"`
}

// APIOperationPolicyStatus defines the observed state of APIOperationPolicy.
type APIOperationPolicyStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        APIOperationPolicyObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// APIOperationPolicy is the Schema for the APIOperationPolicys API. Manages an API Management API Operation Policy
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type APIOperationPolicy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              APIOperationPolicySpec   `json:"spec"`
	Status            APIOperationPolicyStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// APIOperationPolicyList contains a list of APIOperationPolicys
type APIOperationPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []APIOperationPolicy `json:"items"`
}

// Repository type metadata.
var (
	APIOperationPolicy_Kind             = "APIOperationPolicy"
	APIOperationPolicy_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: APIOperationPolicy_Kind}.String()
	APIOperationPolicy_KindAPIVersion   = APIOperationPolicy_Kind + "." + CRDGroupVersion.String()
	APIOperationPolicy_GroupVersionKind = CRDGroupVersion.WithKind(APIOperationPolicy_Kind)
)

func init() {
	SchemeBuilder.Register(&APIOperationPolicy{}, &APIOperationPolicyList{})
}
