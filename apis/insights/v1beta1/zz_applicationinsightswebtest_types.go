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

type ApplicationInsightsWebTestInitParameters struct {

	// An XML configuration specification for a WebTest (see here for more information).
	Configuration *string `json:"configuration,omitempty" tf:"configuration,omitempty"`

	// Purpose/user defined descriptive test for this WebTest.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Is the test actively being monitored.
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// Interval in seconds between test runs for this WebTest. Valid options are 300, 600 and 900. Defaults to 300.
	Frequency *float64 `json:"frequency,omitempty" tf:"frequency,omitempty"`

	// A list of where to physically run the tests from to give global coverage for accessibility of your application.
	GeoLocations []*string `json:"geoLocations,omitempty" tf:"geo_locations,omitempty"`

	// The kind of web test that this web test watches. Choices are ping and multistep. Changing this forces a new resource to be created.
	Kind *string `json:"kind,omitempty" tf:"kind,omitempty"`

	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created. It needs to correlate with location of parent resource (azurerm_application_insights).
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// Specifies the name of the Application Insights WebTest. Changing this forces a new resource to be created.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Allow for retries should this WebTest fail.
	RetryEnabled *bool `json:"retryEnabled,omitempty" tf:"retry_enabled,omitempty"`

	// A mapping of tags to assign to the resource.
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// Seconds until this WebTest will timeout and fail. Default is 30.
	Timeout *float64 `json:"timeout,omitempty" tf:"timeout,omitempty"`
}

type ApplicationInsightsWebTestObservation struct {

	// The ID of the Application Insights component on which the WebTest operates. Changing this forces a new resource to be created.
	ApplicationInsightsID *string `json:"applicationInsightsId,omitempty" tf:"application_insights_id,omitempty"`

	// An XML configuration specification for a WebTest (see here for more information).
	Configuration *string `json:"configuration,omitempty" tf:"configuration,omitempty"`

	// Purpose/user defined descriptive test for this WebTest.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Is the test actively being monitored.
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// Interval in seconds between test runs for this WebTest. Valid options are 300, 600 and 900. Defaults to 300.
	Frequency *float64 `json:"frequency,omitempty" tf:"frequency,omitempty"`

	// A list of where to physically run the tests from to give global coverage for accessibility of your application.
	GeoLocations []*string `json:"geoLocations,omitempty" tf:"geo_locations,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The kind of web test that this web test watches. Choices are ping and multistep. Changing this forces a new resource to be created.
	Kind *string `json:"kind,omitempty" tf:"kind,omitempty"`

	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created. It needs to correlate with location of parent resource (azurerm_application_insights).
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// Specifies the name of the Application Insights WebTest. Changing this forces a new resource to be created.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The name of the resource group in which to create the Application Insights WebTest. Changing this forces a new resource
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// Allow for retries should this WebTest fail.
	RetryEnabled *bool `json:"retryEnabled,omitempty" tf:"retry_enabled,omitempty"`

	SyntheticMonitorID *string `json:"syntheticMonitorId,omitempty" tf:"synthetic_monitor_id,omitempty"`

	// A mapping of tags to assign to the resource.
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// Seconds until this WebTest will timeout and fail. Default is 30.
	Timeout *float64 `json:"timeout,omitempty" tf:"timeout,omitempty"`
}

type ApplicationInsightsWebTestParameters struct {

	// The ID of the Application Insights component on which the WebTest operates. Changing this forces a new resource to be created.
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

	// An XML configuration specification for a WebTest (see here for more information).
	// +kubebuilder:validation:Optional
	Configuration *string `json:"configuration,omitempty" tf:"configuration,omitempty"`

	// Purpose/user defined descriptive test for this WebTest.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Is the test actively being monitored.
	// +kubebuilder:validation:Optional
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// Interval in seconds between test runs for this WebTest. Valid options are 300, 600 and 900. Defaults to 300.
	// +kubebuilder:validation:Optional
	Frequency *float64 `json:"frequency,omitempty" tf:"frequency,omitempty"`

	// A list of where to physically run the tests from to give global coverage for accessibility of your application.
	// +kubebuilder:validation:Optional
	GeoLocations []*string `json:"geoLocations,omitempty" tf:"geo_locations,omitempty"`

	// The kind of web test that this web test watches. Choices are ping and multistep. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	Kind *string `json:"kind,omitempty" tf:"kind,omitempty"`

	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created. It needs to correlate with location of parent resource (azurerm_application_insights).
	// +kubebuilder:validation:Optional
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// Specifies the name of the Application Insights WebTest. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The name of the resource group in which to create the Application Insights WebTest. Changing this forces a new resource
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/azure/v1beta1.ResourceGroup
	// +kubebuilder:validation:Optional
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// Reference to a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameRef *v1.Reference `json:"resourceGroupNameRef,omitempty" tf:"-"`

	// Selector for a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameSelector *v1.Selector `json:"resourceGroupNameSelector,omitempty" tf:"-"`

	// Allow for retries should this WebTest fail.
	// +kubebuilder:validation:Optional
	RetryEnabled *bool `json:"retryEnabled,omitempty" tf:"retry_enabled,omitempty"`

	// A mapping of tags to assign to the resource.
	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// Seconds until this WebTest will timeout and fail. Default is 30.
	// +kubebuilder:validation:Optional
	Timeout *float64 `json:"timeout,omitempty" tf:"timeout,omitempty"`
}

// ApplicationInsightsWebTestSpec defines the desired state of ApplicationInsightsWebTest
type ApplicationInsightsWebTestSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ApplicationInsightsWebTestParameters `json:"forProvider"`
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
	InitProvider ApplicationInsightsWebTestInitParameters `json:"initProvider,omitempty"`
}

// ApplicationInsightsWebTestStatus defines the observed state of ApplicationInsightsWebTest.
type ApplicationInsightsWebTestStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ApplicationInsightsWebTestObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// ApplicationInsightsWebTest is the Schema for the ApplicationInsightsWebTests API. Manages an Application Insights WebTest.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type ApplicationInsightsWebTest struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.configuration) || has(self.initProvider.configuration)",message="configuration is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.geoLocations) || has(self.initProvider.geoLocations)",message="geoLocations is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.kind) || has(self.initProvider.kind)",message="kind is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.location) || has(self.initProvider.location)",message="location is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || has(self.initProvider.name)",message="name is a required parameter"
	Spec   ApplicationInsightsWebTestSpec   `json:"spec"`
	Status ApplicationInsightsWebTestStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ApplicationInsightsWebTestList contains a list of ApplicationInsightsWebTests
type ApplicationInsightsWebTestList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ApplicationInsightsWebTest `json:"items"`
}

// Repository type metadata.
var (
	ApplicationInsightsWebTest_Kind             = "ApplicationInsightsWebTest"
	ApplicationInsightsWebTest_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: ApplicationInsightsWebTest_Kind}.String()
	ApplicationInsightsWebTest_KindAPIVersion   = ApplicationInsightsWebTest_Kind + "." + CRDGroupVersion.String()
	ApplicationInsightsWebTest_GroupVersionKind = CRDGroupVersion.WithKind(ApplicationInsightsWebTest_Kind)
)

func init() {
	SchemeBuilder.Register(&ApplicationInsightsWebTest{}, &ApplicationInsightsWebTestList{})
}
