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

type ManagementLockInitParameters struct {

	// Specifies the Level to be used for this Lock. Possible values are CanNotDelete and ReadOnly. Changing this forces a new resource to be created.
	LockLevel *string `json:"lockLevel,omitempty" tf:"lock_level,omitempty"`

	// Specifies the name of the Management Lock. Changing this forces a new resource to be created.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Specifies some notes about the lock. Maximum of 512 characters. Changing this forces a new resource to be created.
	Notes *string `json:"notes,omitempty" tf:"notes,omitempty"`
}

type ManagementLockObservation struct {

	// The ID of the Management Lock
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Specifies the Level to be used for this Lock. Possible values are CanNotDelete and ReadOnly. Changing this forces a new resource to be created.
	LockLevel *string `json:"lockLevel,omitempty" tf:"lock_level,omitempty"`

	// Specifies the name of the Management Lock. Changing this forces a new resource to be created.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Specifies some notes about the lock. Maximum of 512 characters. Changing this forces a new resource to be created.
	Notes *string `json:"notes,omitempty" tf:"notes,omitempty"`

	// Specifies the scope at which the Management Lock should be created. Changing this forces a new resource to be created.
	Scope *string `json:"scope,omitempty" tf:"scope,omitempty"`
}

type ManagementLockParameters struct {

	// Specifies the Level to be used for this Lock. Possible values are CanNotDelete and ReadOnly. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	LockLevel *string `json:"lockLevel,omitempty" tf:"lock_level,omitempty"`

	// Specifies the name of the Management Lock. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Specifies some notes about the lock. Maximum of 512 characters. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	Notes *string `json:"notes,omitempty" tf:"notes,omitempty"`

	// Specifies the scope at which the Management Lock should be created. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/azure/v1beta1.ResourceGroup
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	Scope *string `json:"scope,omitempty" tf:"scope,omitempty"`

	// Reference to a ResourceGroup in azure to populate scope.
	// +kubebuilder:validation:Optional
	ScopeRef *v1.Reference `json:"scopeRef,omitempty" tf:"-"`

	// Selector for a ResourceGroup in azure to populate scope.
	// +kubebuilder:validation:Optional
	ScopeSelector *v1.Selector `json:"scopeSelector,omitempty" tf:"-"`
}

// ManagementLockSpec defines the desired state of ManagementLock
type ManagementLockSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ManagementLockParameters `json:"forProvider"`
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
	InitProvider ManagementLockInitParameters `json:"initProvider,omitempty"`
}

// ManagementLockStatus defines the observed state of ManagementLock.
type ManagementLockStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ManagementLockObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// ManagementLock is the Schema for the ManagementLocks API. Manages a Management Lock which is scoped to a Subscription, Resource Group or Resource.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type ManagementLock struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.lockLevel) || has(self.initProvider.lockLevel)",message="lockLevel is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || has(self.initProvider.name)",message="name is a required parameter"
	Spec   ManagementLockSpec   `json:"spec"`
	Status ManagementLockStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ManagementLockList contains a list of ManagementLocks
type ManagementLockList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ManagementLock `json:"items"`
}

// Repository type metadata.
var (
	ManagementLock_Kind             = "ManagementLock"
	ManagementLock_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: ManagementLock_Kind}.String()
	ManagementLock_KindAPIVersion   = ManagementLock_Kind + "." + CRDGroupVersion.String()
	ManagementLock_GroupVersionKind = CRDGroupVersion.WithKind(ManagementLock_Kind)
)

func init() {
	SchemeBuilder.Register(&ManagementLock{}, &ManagementLockList{})
}
