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

type APIReleaseInitParameters struct {

	// The Release Notes.
	Notes *string `json:"notes,omitempty" tf:"notes,omitempty"`
}

type APIReleaseObservation struct {

	// The ID of the API Management API. Changing this forces a new API Management API Release to be created.
	APIID *string `json:"apiId,omitempty" tf:"api_id,omitempty"`

	// The ID of the API Management API Release.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The Release Notes.
	Notes *string `json:"notes,omitempty" tf:"notes,omitempty"`
}

type APIReleaseParameters struct {

	// The ID of the API Management API. Changing this forces a new API Management API Release to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/apimanagement/v1beta1.API
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	APIID *string `json:"apiId,omitempty" tf:"api_id,omitempty"`

	// Reference to a API in apimanagement to populate apiId.
	// +kubebuilder:validation:Optional
	APIIDRef *v1.Reference `json:"apiIdRef,omitempty" tf:"-"`

	// Selector for a API in apimanagement to populate apiId.
	// +kubebuilder:validation:Optional
	APIIDSelector *v1.Selector `json:"apiIdSelector,omitempty" tf:"-"`

	// The Release Notes.
	// +kubebuilder:validation:Optional
	Notes *string `json:"notes,omitempty" tf:"notes,omitempty"`
}

// APIReleaseSpec defines the desired state of APIRelease
type APIReleaseSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     APIReleaseParameters `json:"forProvider"`
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
	InitProvider APIReleaseInitParameters `json:"initProvider,omitempty"`
}

// APIReleaseStatus defines the observed state of APIRelease.
type APIReleaseStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        APIReleaseObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// APIRelease is the Schema for the APIReleases API. Manages a API Management API Release.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type APIRelease struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              APIReleaseSpec   `json:"spec"`
	Status            APIReleaseStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// APIReleaseList contains a list of APIReleases
type APIReleaseList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []APIRelease `json:"items"`
}

// Repository type metadata.
var (
	APIRelease_Kind             = "APIRelease"
	APIRelease_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: APIRelease_Kind}.String()
	APIRelease_KindAPIVersion   = APIRelease_Kind + "." + CRDGroupVersion.String()
	APIRelease_GroupVersionKind = CRDGroupVersion.WithKind(APIRelease_Kind)
)

func init() {
	SchemeBuilder.Register(&APIRelease{}, &APIReleaseList{})
}
