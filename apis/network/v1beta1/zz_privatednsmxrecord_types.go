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

type PrivateDNSMXRecordInitParameters struct {

	// One or more record blocks as defined below.
	Record []PrivateDNSMXRecordRecordInitParameters `json:"record,omitempty" tf:"record,omitempty"`

	// The Time To Live (TTL) of the DNS record in seconds.
	TTL *float64 `json:"ttl,omitempty" tf:"ttl,omitempty"`

	// A mapping of tags to assign to the resource.
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type PrivateDNSMXRecordObservation struct {

	// The FQDN of the DNS MX Record.
	Fqdn *string `json:"fqdn,omitempty" tf:"fqdn,omitempty"`

	// The Private DNS MX Record ID.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// One or more record blocks as defined below.
	Record []PrivateDNSMXRecordRecordObservation `json:"record,omitempty" tf:"record,omitempty"`

	// Specifies the resource group where the resource exists. Changing this forces a new resource to be created.
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// The Time To Live (TTL) of the DNS record in seconds.
	TTL *float64 `json:"ttl,omitempty" tf:"ttl,omitempty"`

	// A mapping of tags to assign to the resource.
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// Specifies the Private DNS Zone where the resource exists. Changing this forces a new resource to be created.
	ZoneName *string `json:"zoneName,omitempty" tf:"zone_name,omitempty"`
}

type PrivateDNSMXRecordParameters struct {

	// One or more record blocks as defined below.
	// +kubebuilder:validation:Optional
	Record []PrivateDNSMXRecordRecordParameters `json:"record,omitempty" tf:"record,omitempty"`

	// Specifies the resource group where the resource exists. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/azure/v1beta1.ResourceGroup
	// +kubebuilder:validation:Optional
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// Reference to a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameRef *v1.Reference `json:"resourceGroupNameRef,omitempty" tf:"-"`

	// Selector for a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameSelector *v1.Selector `json:"resourceGroupNameSelector,omitempty" tf:"-"`

	// The Time To Live (TTL) of the DNS record in seconds.
	// +kubebuilder:validation:Optional
	TTL *float64 `json:"ttl,omitempty" tf:"ttl,omitempty"`

	// A mapping of tags to assign to the resource.
	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// Specifies the Private DNS Zone where the resource exists. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=PrivateDNSZone
	// +kubebuilder:validation:Optional
	ZoneName *string `json:"zoneName,omitempty" tf:"zone_name,omitempty"`

	// Reference to a PrivateDNSZone to populate zoneName.
	// +kubebuilder:validation:Optional
	ZoneNameRef *v1.Reference `json:"zoneNameRef,omitempty" tf:"-"`

	// Selector for a PrivateDNSZone to populate zoneName.
	// +kubebuilder:validation:Optional
	ZoneNameSelector *v1.Selector `json:"zoneNameSelector,omitempty" tf:"-"`
}

type PrivateDNSMXRecordRecordInitParameters struct {

	// The FQDN of the exchange to MX record points to.
	Exchange *string `json:"exchange,omitempty" tf:"exchange,omitempty"`

	// The preference of the MX record.
	Preference *float64 `json:"preference,omitempty" tf:"preference,omitempty"`
}

type PrivateDNSMXRecordRecordObservation struct {

	// The FQDN of the exchange to MX record points to.
	Exchange *string `json:"exchange,omitempty" tf:"exchange,omitempty"`

	// The preference of the MX record.
	Preference *float64 `json:"preference,omitempty" tf:"preference,omitempty"`
}

type PrivateDNSMXRecordRecordParameters struct {

	// The FQDN of the exchange to MX record points to.
	// +kubebuilder:validation:Optional
	Exchange *string `json:"exchange,omitempty" tf:"exchange,omitempty"`

	// The preference of the MX record.
	// +kubebuilder:validation:Optional
	Preference *float64 `json:"preference,omitempty" tf:"preference,omitempty"`
}

// PrivateDNSMXRecordSpec defines the desired state of PrivateDNSMXRecord
type PrivateDNSMXRecordSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     PrivateDNSMXRecordParameters `json:"forProvider"`
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
	InitProvider PrivateDNSMXRecordInitParameters `json:"initProvider,omitempty"`
}

// PrivateDNSMXRecordStatus defines the observed state of PrivateDNSMXRecord.
type PrivateDNSMXRecordStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        PrivateDNSMXRecordObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// PrivateDNSMXRecord is the Schema for the PrivateDNSMXRecords API. Manages a Private DNS MX Record.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type PrivateDNSMXRecord struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.record) || has(self.initProvider.record)",message="record is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.ttl) || has(self.initProvider.ttl)",message="ttl is a required parameter"
	Spec   PrivateDNSMXRecordSpec   `json:"spec"`
	Status PrivateDNSMXRecordStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// PrivateDNSMXRecordList contains a list of PrivateDNSMXRecords
type PrivateDNSMXRecordList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []PrivateDNSMXRecord `json:"items"`
}

// Repository type metadata.
var (
	PrivateDNSMXRecord_Kind             = "PrivateDNSMXRecord"
	PrivateDNSMXRecord_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: PrivateDNSMXRecord_Kind}.String()
	PrivateDNSMXRecord_KindAPIVersion   = PrivateDNSMXRecord_Kind + "." + CRDGroupVersion.String()
	PrivateDNSMXRecord_GroupVersionKind = CRDGroupVersion.WithKind(PrivateDNSMXRecord_Kind)
)

func init() {
	SchemeBuilder.Register(&PrivateDNSMXRecord{}, &PrivateDNSMXRecordList{})
}
