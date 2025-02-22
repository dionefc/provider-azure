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

type ActionInitParameters struct {

	// A request_header block as defined below.
	RequestHeader []RequestHeaderInitParameters `json:"requestHeader,omitempty" tf:"request_header,omitempty"`

	// A response_header block as defined below.
	ResponseHeader []ResponseHeaderInitParameters `json:"responseHeader,omitempty" tf:"response_header,omitempty"`
}

type ActionObservation struct {

	// A request_header block as defined below.
	RequestHeader []RequestHeaderObservation `json:"requestHeader,omitempty" tf:"request_header,omitempty"`

	// A response_header block as defined below.
	ResponseHeader []ResponseHeaderObservation `json:"responseHeader,omitempty" tf:"response_header,omitempty"`
}

type ActionParameters struct {

	// A request_header block as defined below.
	// +kubebuilder:validation:Optional
	RequestHeader []RequestHeaderParameters `json:"requestHeader,omitempty" tf:"request_header,omitempty"`

	// A response_header block as defined below.
	// +kubebuilder:validation:Optional
	ResponseHeader []ResponseHeaderParameters `json:"responseHeader,omitempty" tf:"response_header,omitempty"`
}

type FrontdoorRulesEngineInitParameters struct {

	// Whether this Rules engine configuration is enabled? Defaults to true.
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// A rule block as defined below.
	Rule []FrontdoorRulesEngineRuleInitParameters `json:"rule,omitempty" tf:"rule,omitempty"`
}

type FrontdoorRulesEngineObservation struct {

	// Whether this Rules engine configuration is enabled? Defaults to true.
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// The name of the Front Door instance. Changing this forces a new resource to be created.
	FrontdoorName *string `json:"frontdoorName,omitempty" tf:"frontdoor_name,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// The name of the resource group. Changing this forces a new resource to be created.
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// A rule block as defined below.
	Rule []FrontdoorRulesEngineRuleObservation `json:"rule,omitempty" tf:"rule,omitempty"`
}

type FrontdoorRulesEngineParameters struct {

	// Whether this Rules engine configuration is enabled? Defaults to true.
	// +kubebuilder:validation:Optional
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// The name of the Front Door instance. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=FrontDoor
	// +kubebuilder:validation:Optional
	FrontdoorName *string `json:"frontdoorName,omitempty" tf:"frontdoor_name,omitempty"`

	// Reference to a FrontDoor to populate frontdoorName.
	// +kubebuilder:validation:Optional
	FrontdoorNameRef *v1.Reference `json:"frontdoorNameRef,omitempty" tf:"-"`

	// Selector for a FrontDoor to populate frontdoorName.
	// +kubebuilder:validation:Optional
	FrontdoorNameSelector *v1.Selector `json:"frontdoorNameSelector,omitempty" tf:"-"`

	// The name of the resource group. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/azure/v1beta1.ResourceGroup
	// +kubebuilder:validation:Optional
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// Reference to a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameRef *v1.Reference `json:"resourceGroupNameRef,omitempty" tf:"-"`

	// Selector for a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameSelector *v1.Selector `json:"resourceGroupNameSelector,omitempty" tf:"-"`

	// A rule block as defined below.
	// +kubebuilder:validation:Optional
	Rule []FrontdoorRulesEngineRuleParameters `json:"rule,omitempty" tf:"rule,omitempty"`
}

type FrontdoorRulesEngineRuleInitParameters struct {

	// An action block as defined below.
	Action []ActionInitParameters `json:"action,omitempty" tf:"action,omitempty"`

	// One or more match_condition block as defined below.
	MatchCondition []RuleMatchConditionInitParameters `json:"matchCondition,omitempty" tf:"match_condition,omitempty"`

	// The name of the rule.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Priority of the rule, must be unique per rules engine definition.
	Priority *float64 `json:"priority,omitempty" tf:"priority,omitempty"`
}

type FrontdoorRulesEngineRuleObservation struct {

	// An action block as defined below.
	Action []ActionObservation `json:"action,omitempty" tf:"action,omitempty"`

	// One or more match_condition block as defined below.
	MatchCondition []RuleMatchConditionObservation `json:"matchCondition,omitempty" tf:"match_condition,omitempty"`

	// The name of the rule.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Priority of the rule, must be unique per rules engine definition.
	Priority *float64 `json:"priority,omitempty" tf:"priority,omitempty"`
}

type FrontdoorRulesEngineRuleParameters struct {

	// An action block as defined below.
	// +kubebuilder:validation:Optional
	Action []ActionParameters `json:"action,omitempty" tf:"action,omitempty"`

	// One or more match_condition block as defined below.
	// +kubebuilder:validation:Optional
	MatchCondition []RuleMatchConditionParameters `json:"matchCondition,omitempty" tf:"match_condition,omitempty"`

	// The name of the rule.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Priority of the rule, must be unique per rules engine definition.
	// +kubebuilder:validation:Optional
	Priority *float64 `json:"priority,omitempty" tf:"priority,omitempty"`
}

type RequestHeaderInitParameters struct {

	// can be set to Overwrite, Append or Delete.
	HeaderActionType *string `json:"headerActionType,omitempty" tf:"header_action_type,omitempty"`

	// header name (string).
	HeaderName *string `json:"headerName,omitempty" tf:"header_name,omitempty"`

	// value name (string).
	Value *string `json:"value,omitempty" tf:"value,omitempty"`
}

type RequestHeaderObservation struct {

	// can be set to Overwrite, Append or Delete.
	HeaderActionType *string `json:"headerActionType,omitempty" tf:"header_action_type,omitempty"`

	// header name (string).
	HeaderName *string `json:"headerName,omitempty" tf:"header_name,omitempty"`

	// value name (string).
	Value *string `json:"value,omitempty" tf:"value,omitempty"`
}

type RequestHeaderParameters struct {

	// can be set to Overwrite, Append or Delete.
	// +kubebuilder:validation:Optional
	HeaderActionType *string `json:"headerActionType,omitempty" tf:"header_action_type,omitempty"`

	// header name (string).
	// +kubebuilder:validation:Optional
	HeaderName *string `json:"headerName,omitempty" tf:"header_name,omitempty"`

	// value name (string).
	// +kubebuilder:validation:Optional
	Value *string `json:"value,omitempty" tf:"value,omitempty"`
}

type ResponseHeaderInitParameters struct {

	// can be set to Overwrite, Append or Delete.
	HeaderActionType *string `json:"headerActionType,omitempty" tf:"header_action_type,omitempty"`

	// header name (string).
	HeaderName *string `json:"headerName,omitempty" tf:"header_name,omitempty"`

	// value name (string).
	Value *string `json:"value,omitempty" tf:"value,omitempty"`
}

type ResponseHeaderObservation struct {

	// can be set to Overwrite, Append or Delete.
	HeaderActionType *string `json:"headerActionType,omitempty" tf:"header_action_type,omitempty"`

	// header name (string).
	HeaderName *string `json:"headerName,omitempty" tf:"header_name,omitempty"`

	// value name (string).
	Value *string `json:"value,omitempty" tf:"value,omitempty"`
}

type ResponseHeaderParameters struct {

	// can be set to Overwrite, Append or Delete.
	// +kubebuilder:validation:Optional
	HeaderActionType *string `json:"headerActionType,omitempty" tf:"header_action_type,omitempty"`

	// header name (string).
	// +kubebuilder:validation:Optional
	HeaderName *string `json:"headerName,omitempty" tf:"header_name,omitempty"`

	// value name (string).
	// +kubebuilder:validation:Optional
	Value *string `json:"value,omitempty" tf:"value,omitempty"`
}

type RuleMatchConditionInitParameters struct {

	// can be set to true or false to negate the given condition. Defaults to true.
	NegateCondition *bool `json:"negateCondition,omitempty" tf:"negate_condition,omitempty"`

	// can be set to Any, IPMatch, GeoMatch, Equal, Contains, LessThan, GreaterThan, LessThanOrEqual, GreaterThanOrEqual, BeginsWith or EndsWith
	Operator *string `json:"operator,omitempty" tf:"operator,omitempty"`

	// match against a specific key when variable is set to PostArgs or RequestHeader. It cannot be used with QueryString and RequestMethod.
	Selector *string `json:"selector,omitempty" tf:"selector,omitempty"`

	// can be set to one or more values out of Lowercase, RemoveNulls, Trim, Uppercase, UrlDecode and UrlEncode
	Transform []*string `json:"transform,omitempty" tf:"transform,omitempty"`

	// value name (string).
	Value []*string `json:"value,omitempty" tf:"value,omitempty"`

	// can be set to IsMobile, RemoteAddr, RequestMethod, QueryString, PostArgs, RequestURI, RequestPath, RequestFilename, RequestFilenameExtension,RequestHeader,RequestBody or RequestScheme.
	Variable *string `json:"variable,omitempty" tf:"variable,omitempty"`
}

type RuleMatchConditionObservation struct {

	// can be set to true or false to negate the given condition. Defaults to true.
	NegateCondition *bool `json:"negateCondition,omitempty" tf:"negate_condition,omitempty"`

	// can be set to Any, IPMatch, GeoMatch, Equal, Contains, LessThan, GreaterThan, LessThanOrEqual, GreaterThanOrEqual, BeginsWith or EndsWith
	Operator *string `json:"operator,omitempty" tf:"operator,omitempty"`

	// match against a specific key when variable is set to PostArgs or RequestHeader. It cannot be used with QueryString and RequestMethod.
	Selector *string `json:"selector,omitempty" tf:"selector,omitempty"`

	// can be set to one or more values out of Lowercase, RemoveNulls, Trim, Uppercase, UrlDecode and UrlEncode
	Transform []*string `json:"transform,omitempty" tf:"transform,omitempty"`

	// value name (string).
	Value []*string `json:"value,omitempty" tf:"value,omitempty"`

	// can be set to IsMobile, RemoteAddr, RequestMethod, QueryString, PostArgs, RequestURI, RequestPath, RequestFilename, RequestFilenameExtension,RequestHeader,RequestBody or RequestScheme.
	Variable *string `json:"variable,omitempty" tf:"variable,omitempty"`
}

type RuleMatchConditionParameters struct {

	// can be set to true or false to negate the given condition. Defaults to true.
	// +kubebuilder:validation:Optional
	NegateCondition *bool `json:"negateCondition,omitempty" tf:"negate_condition,omitempty"`

	// can be set to Any, IPMatch, GeoMatch, Equal, Contains, LessThan, GreaterThan, LessThanOrEqual, GreaterThanOrEqual, BeginsWith or EndsWith
	// +kubebuilder:validation:Optional
	Operator *string `json:"operator,omitempty" tf:"operator,omitempty"`

	// match against a specific key when variable is set to PostArgs or RequestHeader. It cannot be used with QueryString and RequestMethod.
	// +kubebuilder:validation:Optional
	Selector *string `json:"selector,omitempty" tf:"selector,omitempty"`

	// can be set to one or more values out of Lowercase, RemoveNulls, Trim, Uppercase, UrlDecode and UrlEncode
	// +kubebuilder:validation:Optional
	Transform []*string `json:"transform,omitempty" tf:"transform,omitempty"`

	// value name (string).
	// +kubebuilder:validation:Optional
	Value []*string `json:"value,omitempty" tf:"value,omitempty"`

	// can be set to IsMobile, RemoteAddr, RequestMethod, QueryString, PostArgs, RequestURI, RequestPath, RequestFilename, RequestFilenameExtension,RequestHeader,RequestBody or RequestScheme.
	// +kubebuilder:validation:Optional
	Variable *string `json:"variable,omitempty" tf:"variable,omitempty"`
}

// FrontdoorRulesEngineSpec defines the desired state of FrontdoorRulesEngine
type FrontdoorRulesEngineSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     FrontdoorRulesEngineParameters `json:"forProvider"`
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
	InitProvider FrontdoorRulesEngineInitParameters `json:"initProvider,omitempty"`
}

// FrontdoorRulesEngineStatus defines the observed state of FrontdoorRulesEngine.
type FrontdoorRulesEngineStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        FrontdoorRulesEngineObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// FrontdoorRulesEngine is the Schema for the FrontdoorRulesEngines API. Manages an Azure Front Door (classic) Rules Engine configuration and rules.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type FrontdoorRulesEngine struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              FrontdoorRulesEngineSpec   `json:"spec"`
	Status            FrontdoorRulesEngineStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// FrontdoorRulesEngineList contains a list of FrontdoorRulesEngines
type FrontdoorRulesEngineList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []FrontdoorRulesEngine `json:"items"`
}

// Repository type metadata.
var (
	FrontdoorRulesEngine_Kind             = "FrontdoorRulesEngine"
	FrontdoorRulesEngine_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: FrontdoorRulesEngine_Kind}.String()
	FrontdoorRulesEngine_KindAPIVersion   = FrontdoorRulesEngine_Kind + "." + CRDGroupVersion.String()
	FrontdoorRulesEngine_GroupVersionKind = CRDGroupVersion.WithKind(FrontdoorRulesEngine_Kind)
)

func init() {
	SchemeBuilder.Register(&FrontdoorRulesEngine{}, &FrontdoorRulesEngineList{})
}
