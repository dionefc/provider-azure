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

type IdentityProviderAADInitParameters struct {

	// List of allowed AAD Tenants.
	AllowedTenants []*string `json:"allowedTenants,omitempty" tf:"allowed_tenants,omitempty"`

	// Client Id of the Application in the AAD Identity Provider.
	ClientID *string `json:"clientId,omitempty" tf:"client_id,omitempty"`

	// The AAD Tenant to use instead of Common when logging into Active Directory
	SigninTenant *string `json:"signinTenant,omitempty" tf:"signin_tenant,omitempty"`
}

type IdentityProviderAADObservation struct {

	// The Name of the API Management Service where this AAD Identity Provider should be created. Changing this forces a new resource to be created.
	APIManagementName *string `json:"apiManagementName,omitempty" tf:"api_management_name,omitempty"`

	// List of allowed AAD Tenants.
	AllowedTenants []*string `json:"allowedTenants,omitempty" tf:"allowed_tenants,omitempty"`

	// Client Id of the Application in the AAD Identity Provider.
	ClientID *string `json:"clientId,omitempty" tf:"client_id,omitempty"`

	// The ID of the API Management AAD Identity Provider.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The Name of the Resource Group where the API Management Service exists. Changing this forces a new resource to be created.
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// The AAD Tenant to use instead of Common when logging into Active Directory
	SigninTenant *string `json:"signinTenant,omitempty" tf:"signin_tenant,omitempty"`
}

type IdentityProviderAADParameters struct {

	// The Name of the API Management Service where this AAD Identity Provider should be created. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/apimanagement/v1beta1.Management
	// +kubebuilder:validation:Optional
	APIManagementName *string `json:"apiManagementName,omitempty" tf:"api_management_name,omitempty"`

	// Reference to a Management in apimanagement to populate apiManagementName.
	// +kubebuilder:validation:Optional
	APIManagementNameRef *v1.Reference `json:"apiManagementNameRef,omitempty" tf:"-"`

	// Selector for a Management in apimanagement to populate apiManagementName.
	// +kubebuilder:validation:Optional
	APIManagementNameSelector *v1.Selector `json:"apiManagementNameSelector,omitempty" tf:"-"`

	// List of allowed AAD Tenants.
	// +kubebuilder:validation:Optional
	AllowedTenants []*string `json:"allowedTenants,omitempty" tf:"allowed_tenants,omitempty"`

	// Client Id of the Application in the AAD Identity Provider.
	// +kubebuilder:validation:Optional
	ClientID *string `json:"clientId,omitempty" tf:"client_id,omitempty"`

	// Client secret of the Application in the AAD Identity Provider.
	// +kubebuilder:validation:Optional
	ClientSecretSecretRef v1.SecretKeySelector `json:"clientSecretSecretRef" tf:"-"`

	// The Name of the Resource Group where the API Management Service exists. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/azure/v1beta1.ResourceGroup
	// +kubebuilder:validation:Optional
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// Reference to a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameRef *v1.Reference `json:"resourceGroupNameRef,omitempty" tf:"-"`

	// Selector for a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameSelector *v1.Selector `json:"resourceGroupNameSelector,omitempty" tf:"-"`

	// The AAD Tenant to use instead of Common when logging into Active Directory
	// +kubebuilder:validation:Optional
	SigninTenant *string `json:"signinTenant,omitempty" tf:"signin_tenant,omitempty"`
}

// IdentityProviderAADSpec defines the desired state of IdentityProviderAAD
type IdentityProviderAADSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     IdentityProviderAADParameters `json:"forProvider"`
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
	InitProvider IdentityProviderAADInitParameters `json:"initProvider,omitempty"`
}

// IdentityProviderAADStatus defines the observed state of IdentityProviderAAD.
type IdentityProviderAADStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        IdentityProviderAADObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// IdentityProviderAAD is the Schema for the IdentityProviderAADs API. Manages an API Management AAD Identity Provider.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type IdentityProviderAAD struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.allowedTenants) || has(self.initProvider.allowedTenants)",message="allowedTenants is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.clientId) || has(self.initProvider.clientId)",message="clientId is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.clientSecretSecretRef)",message="clientSecretSecretRef is a required parameter"
	Spec   IdentityProviderAADSpec   `json:"spec"`
	Status IdentityProviderAADStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// IdentityProviderAADList contains a list of IdentityProviderAADs
type IdentityProviderAADList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []IdentityProviderAAD `json:"items"`
}

// Repository type metadata.
var (
	IdentityProviderAAD_Kind             = "IdentityProviderAAD"
	IdentityProviderAAD_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: IdentityProviderAAD_Kind}.String()
	IdentityProviderAAD_KindAPIVersion   = IdentityProviderAAD_Kind + "." + CRDGroupVersion.String()
	IdentityProviderAAD_GroupVersionKind = CRDGroupVersion.WithKind(IdentityProviderAAD_Kind)
)

func init() {
	SchemeBuilder.Register(&IdentityProviderAAD{}, &IdentityProviderAADList{})
}
