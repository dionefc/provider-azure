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

type HealthcareMedtechServiceFHIRDestinationInitParameters struct {

	// Specifies the destination Fhir mappings of the Med Tech Service Fhir Destination.
	DestinationFHIRMappingJSON *string `json:"destinationFhirMappingJson,omitempty" tf:"destination_fhir_mapping_json,omitempty"`

	// Specifies the destination identity resolution type where the Healthcare Med Tech Service Fhir Destination should be created. Possible values are Create, Lookup.
	DestinationIdentityResolutionType *string `json:"destinationIdentityResolutionType,omitempty" tf:"destination_identity_resolution_type,omitempty"`

	// Specifies the Azure Region where the Healthcare Med Tech Service Fhir Destination should be created. Changing this forces a new Healthcare Med Tech Service Fhir Destination to be created.
	Location *string `json:"location,omitempty" tf:"location,omitempty"`
}

type HealthcareMedtechServiceFHIRDestinationObservation struct {

	// Specifies the destination Fhir mappings of the Med Tech Service Fhir Destination.
	DestinationFHIRMappingJSON *string `json:"destinationFhirMappingJson,omitempty" tf:"destination_fhir_mapping_json,omitempty"`

	// Specifies the destination fhir service id of the Med Tech Service Fhir Destination.
	DestinationFHIRServiceID *string `json:"destinationFhirServiceId,omitempty" tf:"destination_fhir_service_id,omitempty"`

	// Specifies the destination identity resolution type where the Healthcare Med Tech Service Fhir Destination should be created. Possible values are Create, Lookup.
	DestinationIdentityResolutionType *string `json:"destinationIdentityResolutionType,omitempty" tf:"destination_identity_resolution_type,omitempty"`

	// The ID of the Healthcare Med Tech Service Fhir Destination.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Specifies the Azure Region where the Healthcare Med Tech Service Fhir Destination should be created. Changing this forces a new Healthcare Med Tech Service Fhir Destination to be created.
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// Specifies the name of the Healthcare Med Tech Service where the Healthcare Med Tech Service Fhir Destination should exist. Changing this forces a new Healthcare Med Tech Service Fhir Destination to be created.
	MedtechServiceID *string `json:"medtechServiceId,omitempty" tf:"medtech_service_id,omitempty"`
}

type HealthcareMedtechServiceFHIRDestinationParameters struct {

	// Specifies the destination Fhir mappings of the Med Tech Service Fhir Destination.
	// +kubebuilder:validation:Optional
	DestinationFHIRMappingJSON *string `json:"destinationFhirMappingJson,omitempty" tf:"destination_fhir_mapping_json,omitempty"`

	// Specifies the destination fhir service id of the Med Tech Service Fhir Destination.
	// +crossplane:generate:reference:type=HealthcareFHIRService
	// +crossplane:generate:reference:extractor=github.com/upbound/provider-azure/apis/rconfig.ExtractResourceID()
	// +kubebuilder:validation:Optional
	DestinationFHIRServiceID *string `json:"destinationFhirServiceId,omitempty" tf:"destination_fhir_service_id,omitempty"`

	// Reference to a HealthcareFHIRService to populate destinationFhirServiceId.
	// +kubebuilder:validation:Optional
	DestinationFHIRServiceIDRef *v1.Reference `json:"destinationFhirServiceIdRef,omitempty" tf:"-"`

	// Selector for a HealthcareFHIRService to populate destinationFhirServiceId.
	// +kubebuilder:validation:Optional
	DestinationFHIRServiceIDSelector *v1.Selector `json:"destinationFhirServiceIdSelector,omitempty" tf:"-"`

	// Specifies the destination identity resolution type where the Healthcare Med Tech Service Fhir Destination should be created. Possible values are Create, Lookup.
	// +kubebuilder:validation:Optional
	DestinationIdentityResolutionType *string `json:"destinationIdentityResolutionType,omitempty" tf:"destination_identity_resolution_type,omitempty"`

	// Specifies the Azure Region where the Healthcare Med Tech Service Fhir Destination should be created. Changing this forces a new Healthcare Med Tech Service Fhir Destination to be created.
	// +kubebuilder:validation:Optional
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// Specifies the name of the Healthcare Med Tech Service where the Healthcare Med Tech Service Fhir Destination should exist. Changing this forces a new Healthcare Med Tech Service Fhir Destination to be created.
	// +crossplane:generate:reference:type=HealthcareMedtechService
	// +crossplane:generate:reference:extractor=github.com/upbound/provider-azure/apis/rconfig.ExtractResourceID()
	// +kubebuilder:validation:Optional
	MedtechServiceID *string `json:"medtechServiceId,omitempty" tf:"medtech_service_id,omitempty"`

	// Reference to a HealthcareMedtechService to populate medtechServiceId.
	// +kubebuilder:validation:Optional
	MedtechServiceIDRef *v1.Reference `json:"medtechServiceIdRef,omitempty" tf:"-"`

	// Selector for a HealthcareMedtechService to populate medtechServiceId.
	// +kubebuilder:validation:Optional
	MedtechServiceIDSelector *v1.Selector `json:"medtechServiceIdSelector,omitempty" tf:"-"`
}

// HealthcareMedtechServiceFHIRDestinationSpec defines the desired state of HealthcareMedtechServiceFHIRDestination
type HealthcareMedtechServiceFHIRDestinationSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     HealthcareMedtechServiceFHIRDestinationParameters `json:"forProvider"`
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
	InitProvider HealthcareMedtechServiceFHIRDestinationInitParameters `json:"initProvider,omitempty"`
}

// HealthcareMedtechServiceFHIRDestinationStatus defines the observed state of HealthcareMedtechServiceFHIRDestination.
type HealthcareMedtechServiceFHIRDestinationStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        HealthcareMedtechServiceFHIRDestinationObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// HealthcareMedtechServiceFHIRDestination is the Schema for the HealthcareMedtechServiceFHIRDestinations API. Manages a Healthcare Med Tech (Internet of Medical Things) Service Fhir Destination.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type HealthcareMedtechServiceFHIRDestination struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.destinationFhirMappingJson) || has(self.initProvider.destinationFhirMappingJson)",message="destinationFhirMappingJson is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.destinationIdentityResolutionType) || has(self.initProvider.destinationIdentityResolutionType)",message="destinationIdentityResolutionType is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.location) || has(self.initProvider.location)",message="location is a required parameter"
	Spec   HealthcareMedtechServiceFHIRDestinationSpec   `json:"spec"`
	Status HealthcareMedtechServiceFHIRDestinationStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// HealthcareMedtechServiceFHIRDestinationList contains a list of HealthcareMedtechServiceFHIRDestinations
type HealthcareMedtechServiceFHIRDestinationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []HealthcareMedtechServiceFHIRDestination `json:"items"`
}

// Repository type metadata.
var (
	HealthcareMedtechServiceFHIRDestination_Kind             = "HealthcareMedtechServiceFHIRDestination"
	HealthcareMedtechServiceFHIRDestination_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: HealthcareMedtechServiceFHIRDestination_Kind}.String()
	HealthcareMedtechServiceFHIRDestination_KindAPIVersion   = HealthcareMedtechServiceFHIRDestination_Kind + "." + CRDGroupVersion.String()
	HealthcareMedtechServiceFHIRDestination_GroupVersionKind = CRDGroupVersion.WithKind(HealthcareMedtechServiceFHIRDestination_Kind)
)

func init() {
	SchemeBuilder.Register(&HealthcareMedtechServiceFHIRDestination{}, &HealthcareMedtechServiceFHIRDestinationList{})
}
