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

type ConfigurationInitParameters struct {

	// The assignment type for the Guest Configuration Assignment. Possible values are Audit, ApplyAndAutoCorrect, ApplyAndMonitor and DeployAndAutoCorrect.
	AssignmentType *string `json:"assignmentType,omitempty" tf:"assignment_type,omitempty"`

	// The content hash for the Guest Configuration package.
	ContentHash *string `json:"contentHash,omitempty" tf:"content_hash,omitempty"`

	// The content URI where the Guest Configuration package is stored.
	ContentURI *string `json:"contentUri,omitempty" tf:"content_uri,omitempty"`

	// One or more parameter blocks as defined below which define what configuration parameters and values against.
	Parameter []ParameterInitParameters `json:"parameter,omitempty" tf:"parameter,omitempty"`

	// The version of the Guest Configuration that will be assigned in this Guest Configuration Assignment.
	Version *string `json:"version,omitempty" tf:"version,omitempty"`
}

type ConfigurationObservation struct {

	// The assignment type for the Guest Configuration Assignment. Possible values are Audit, ApplyAndAutoCorrect, ApplyAndMonitor and DeployAndAutoCorrect.
	AssignmentType *string `json:"assignmentType,omitempty" tf:"assignment_type,omitempty"`

	// The content hash for the Guest Configuration package.
	ContentHash *string `json:"contentHash,omitempty" tf:"content_hash,omitempty"`

	// The content URI where the Guest Configuration package is stored.
	ContentURI *string `json:"contentUri,omitempty" tf:"content_uri,omitempty"`

	// One or more parameter blocks as defined below which define what configuration parameters and values against.
	Parameter []ParameterObservation `json:"parameter,omitempty" tf:"parameter,omitempty"`

	// The version of the Guest Configuration that will be assigned in this Guest Configuration Assignment.
	Version *string `json:"version,omitempty" tf:"version,omitempty"`
}

type ConfigurationParameters struct {

	// The assignment type for the Guest Configuration Assignment. Possible values are Audit, ApplyAndAutoCorrect, ApplyAndMonitor and DeployAndAutoCorrect.
	// +kubebuilder:validation:Optional
	AssignmentType *string `json:"assignmentType,omitempty" tf:"assignment_type,omitempty"`

	// The content hash for the Guest Configuration package.
	// +kubebuilder:validation:Optional
	ContentHash *string `json:"contentHash,omitempty" tf:"content_hash,omitempty"`

	// The content URI where the Guest Configuration package is stored.
	// +kubebuilder:validation:Optional
	ContentURI *string `json:"contentUri,omitempty" tf:"content_uri,omitempty"`

	// One or more parameter blocks as defined below which define what configuration parameters and values against.
	// +kubebuilder:validation:Optional
	Parameter []ParameterParameters `json:"parameter,omitempty" tf:"parameter,omitempty"`

	// The version of the Guest Configuration that will be assigned in this Guest Configuration Assignment.
	// +kubebuilder:validation:Optional
	Version *string `json:"version,omitempty" tf:"version,omitempty"`
}

type ParameterInitParameters struct {

	// The name of the configuration parameter to check.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The value to check the configuration parameter with.
	Value *string `json:"value,omitempty" tf:"value,omitempty"`
}

type ParameterObservation struct {

	// The name of the configuration parameter to check.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The value to check the configuration parameter with.
	Value *string `json:"value,omitempty" tf:"value,omitempty"`
}

type ParameterParameters struct {

	// The name of the configuration parameter to check.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The value to check the configuration parameter with.
	// +kubebuilder:validation:Optional
	Value *string `json:"value,omitempty" tf:"value,omitempty"`
}

type PolicyVirtualMachineConfigurationAssignmentInitParameters struct {

	// A configuration block as defined below.
	Configuration []ConfigurationInitParameters `json:"configuration,omitempty" tf:"configuration,omitempty"`

	// The Azure location where the Policy Virtual Machine Configuration Assignment should exist. Changing this forces a new resource to be created.
	Location *string `json:"location,omitempty" tf:"location,omitempty"`
}

type PolicyVirtualMachineConfigurationAssignmentObservation struct {

	// A configuration block as defined below.
	Configuration []ConfigurationObservation `json:"configuration,omitempty" tf:"configuration,omitempty"`

	// The ID of the Policy Virtual Machine Configuration Assignment.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The Azure location where the Policy Virtual Machine Configuration Assignment should exist. Changing this forces a new resource to be created.
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// The resource ID of the Policy Virtual Machine which this Guest Configuration Assignment should apply to. Changing this forces a new resource to be created.
	VirtualMachineID *string `json:"virtualMachineId,omitempty" tf:"virtual_machine_id,omitempty"`
}

type PolicyVirtualMachineConfigurationAssignmentParameters struct {

	// A configuration block as defined below.
	// +kubebuilder:validation:Optional
	Configuration []ConfigurationParameters `json:"configuration,omitempty" tf:"configuration,omitempty"`

	// The Azure location where the Policy Virtual Machine Configuration Assignment should exist. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// The resource ID of the Policy Virtual Machine which this Guest Configuration Assignment should apply to. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/compute/v1beta1.WindowsVirtualMachine
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	VirtualMachineID *string `json:"virtualMachineId,omitempty" tf:"virtual_machine_id,omitempty"`

	// Reference to a WindowsVirtualMachine in compute to populate virtualMachineId.
	// +kubebuilder:validation:Optional
	VirtualMachineIDRef *v1.Reference `json:"virtualMachineIdRef,omitempty" tf:"-"`

	// Selector for a WindowsVirtualMachine in compute to populate virtualMachineId.
	// +kubebuilder:validation:Optional
	VirtualMachineIDSelector *v1.Selector `json:"virtualMachineIdSelector,omitempty" tf:"-"`
}

// PolicyVirtualMachineConfigurationAssignmentSpec defines the desired state of PolicyVirtualMachineConfigurationAssignment
type PolicyVirtualMachineConfigurationAssignmentSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     PolicyVirtualMachineConfigurationAssignmentParameters `json:"forProvider"`
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
	InitProvider PolicyVirtualMachineConfigurationAssignmentInitParameters `json:"initProvider,omitempty"`
}

// PolicyVirtualMachineConfigurationAssignmentStatus defines the observed state of PolicyVirtualMachineConfigurationAssignment.
type PolicyVirtualMachineConfigurationAssignmentStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        PolicyVirtualMachineConfigurationAssignmentObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// PolicyVirtualMachineConfigurationAssignment is the Schema for the PolicyVirtualMachineConfigurationAssignments API. Applies a Guest Configuration Policy to a Virtual Machine.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type PolicyVirtualMachineConfigurationAssignment struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.configuration) || has(self.initProvider.configuration)",message="configuration is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.location) || has(self.initProvider.location)",message="location is a required parameter"
	Spec   PolicyVirtualMachineConfigurationAssignmentSpec   `json:"spec"`
	Status PolicyVirtualMachineConfigurationAssignmentStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// PolicyVirtualMachineConfigurationAssignmentList contains a list of PolicyVirtualMachineConfigurationAssignments
type PolicyVirtualMachineConfigurationAssignmentList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []PolicyVirtualMachineConfigurationAssignment `json:"items"`
}

// Repository type metadata.
var (
	PolicyVirtualMachineConfigurationAssignment_Kind             = "PolicyVirtualMachineConfigurationAssignment"
	PolicyVirtualMachineConfigurationAssignment_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: PolicyVirtualMachineConfigurationAssignment_Kind}.String()
	PolicyVirtualMachineConfigurationAssignment_KindAPIVersion   = PolicyVirtualMachineConfigurationAssignment_Kind + "." + CRDGroupVersion.String()
	PolicyVirtualMachineConfigurationAssignment_GroupVersionKind = CRDGroupVersion.WithKind(PolicyVirtualMachineConfigurationAssignment_Kind)
)

func init() {
	SchemeBuilder.Register(&PolicyVirtualMachineConfigurationAssignment{}, &PolicyVirtualMachineConfigurationAssignmentList{})
}
