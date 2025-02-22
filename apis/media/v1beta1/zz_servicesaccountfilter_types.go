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

type ServicesAccountFilterInitParameters struct {

	// The first quality bitrate. Sets the first video track to appear in the Live Streaming playlist to allow HLS native players to start downloading from this quality level at the beginning.
	FirstQualityBitrate *float64 `json:"firstQualityBitrate,omitempty" tf:"first_quality_bitrate,omitempty"`

	// A presentation_time_range block as defined below.
	PresentationTimeRange []ServicesAccountFilterPresentationTimeRangeInitParameters `json:"presentationTimeRange,omitempty" tf:"presentation_time_range,omitempty"`

	// One or more track_selection blocks as defined below.
	TrackSelection []ServicesAccountFilterTrackSelectionInitParameters `json:"trackSelection,omitempty" tf:"track_selection,omitempty"`
}

type ServicesAccountFilterObservation struct {

	// The first quality bitrate. Sets the first video track to appear in the Live Streaming playlist to allow HLS native players to start downloading from this quality level at the beginning.
	FirstQualityBitrate *float64 `json:"firstQualityBitrate,omitempty" tf:"first_quality_bitrate,omitempty"`

	// The ID of the Account Filter.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The Media Services account name. Changing this forces a new Account Filter to be created.
	MediaServicesAccountName *string `json:"mediaServicesAccountName,omitempty" tf:"media_services_account_name,omitempty"`

	// A presentation_time_range block as defined below.
	PresentationTimeRange []ServicesAccountFilterPresentationTimeRangeObservation `json:"presentationTimeRange,omitempty" tf:"presentation_time_range,omitempty"`

	// The name of the Resource Group where the Account Filter should exist. Changing this forces a new Account Filter to be created.
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// One or more track_selection blocks as defined below.
	TrackSelection []ServicesAccountFilterTrackSelectionObservation `json:"trackSelection,omitempty" tf:"track_selection,omitempty"`
}

type ServicesAccountFilterParameters struct {

	// The first quality bitrate. Sets the first video track to appear in the Live Streaming playlist to allow HLS native players to start downloading from this quality level at the beginning.
	// +kubebuilder:validation:Optional
	FirstQualityBitrate *float64 `json:"firstQualityBitrate,omitempty" tf:"first_quality_bitrate,omitempty"`

	// The Media Services account name. Changing this forces a new Account Filter to be created.
	// +crossplane:generate:reference:type=ServicesAccount
	// +kubebuilder:validation:Optional
	MediaServicesAccountName *string `json:"mediaServicesAccountName,omitempty" tf:"media_services_account_name,omitempty"`

	// Reference to a ServicesAccount to populate mediaServicesAccountName.
	// +kubebuilder:validation:Optional
	MediaServicesAccountNameRef *v1.Reference `json:"mediaServicesAccountNameRef,omitempty" tf:"-"`

	// Selector for a ServicesAccount to populate mediaServicesAccountName.
	// +kubebuilder:validation:Optional
	MediaServicesAccountNameSelector *v1.Selector `json:"mediaServicesAccountNameSelector,omitempty" tf:"-"`

	// A presentation_time_range block as defined below.
	// +kubebuilder:validation:Optional
	PresentationTimeRange []ServicesAccountFilterPresentationTimeRangeParameters `json:"presentationTimeRange,omitempty" tf:"presentation_time_range,omitempty"`

	// The name of the Resource Group where the Account Filter should exist. Changing this forces a new Account Filter to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/azure/v1beta1.ResourceGroup
	// +kubebuilder:validation:Optional
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// Reference to a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameRef *v1.Reference `json:"resourceGroupNameRef,omitempty" tf:"-"`

	// Selector for a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameSelector *v1.Selector `json:"resourceGroupNameSelector,omitempty" tf:"-"`

	// One or more track_selection blocks as defined below.
	// +kubebuilder:validation:Optional
	TrackSelection []ServicesAccountFilterTrackSelectionParameters `json:"trackSelection,omitempty" tf:"track_selection,omitempty"`
}

type ServicesAccountFilterPresentationTimeRangeInitParameters struct {

	// The absolute end time boundary. Applies to Video on Demand (VoD).
	// For the Live Streaming presentation, it is silently ignored and applied when the presentation ends and the stream becomes VoD. This is a long value that represents an absolute end point of the presentation, rounded to the closest next GOP start. The unit is defined by unit_timescale_in_milliseconds, so an end_in_units of 180 would be for 3 minutes. Use start_in_units and end_in_units to trim the fragments that will be in the playlist (manifest). For example, start_in_units set to 20 and end_in_units set to 60 using unit_timescale_in_milliseconds in 1000 will generate a playlist that contains fragments from between 20 seconds and 60 seconds of the VoD presentation. If a fragment straddles the boundary, the entire fragment will be included in the manifest.
	EndInUnits *float64 `json:"endInUnits,omitempty" tf:"end_in_units,omitempty"`

	// Indicates whether the end_in_units property must be present. If true, end_in_units must be specified or a bad request code is returned. Applies to Live Streaming only. Allowed values: false, true.
	ForceEnd *bool `json:"forceEnd,omitempty" tf:"force_end,omitempty"`

	// The relative to end right edge. Applies to Live Streaming only.
	// This value defines the latest live position that a client can seek to. Using this property, you can delay live playback position and create a server-side buffer for players. The unit is defined by unit_timescale_in_milliseconds. The maximum live back off duration is 300 seconds. For example, a value of 20 means that the latest available content is 20 seconds delayed from the real live edge.
	LiveBackoffInUnits *float64 `json:"liveBackoffInUnits,omitempty" tf:"live_backoff_in_units,omitempty"`

	// The relative to end sliding window. Applies to Live Streaming only. Use presentation_window_in_units to apply a sliding window of fragments to include in a playlist. The unit is defined by unit_timescale_in_milliseconds. For example, set presentation_window_in_units to 120 to apply a two-minute sliding window. Media within 2 minutes of the live edge will be included in the playlist. If a fragment straddles the boundary, the entire fragment will be included in the playlist. The minimum presentation window duration is 60 seconds.
	PresentationWindowInUnits *float64 `json:"presentationWindowInUnits,omitempty" tf:"presentation_window_in_units,omitempty"`

	// The absolute start time boundary. Applies to Video on Demand (VoD) or Live Streaming. This is a long value that represents an absolute start point of the stream. The value gets rounded to the closest next GOP start. The unit is defined by unit_timescale_in_milliseconds, so a start_in_units of 15 would be for 15 seconds. Use start_in_units and end_in_units to trim the fragments that will be in the playlist (manifest). For example, start_in_units set to 20 and end_in_units set to 60 using unit_timescale_in_milliseconds in 1000 will generate a playlist that contains fragments from between 20 seconds and 60 seconds of the VoD presentation. If a fragment straddles the boundary, the entire fragment will be included in the manifest.
	StartInUnits *float64 `json:"startInUnits,omitempty" tf:"start_in_units,omitempty"`

	// Specified as the number of milliseconds in one unit timescale. For example, if you want to set a start_in_units at 30 seconds, you would use a value of 30 when using the unit_timescale_in_milliseconds in 1000. Or if you want to set start_in_units in 30 milliseconds, you would use a value of 30 when using the unit_timescale_in_milliseconds in 1. Applies timescale to start_in_units, start_timescale and presentation_window_in_timescale and live_backoff_in_timescale.
	UnitTimescaleInMilliseconds *float64 `json:"unitTimescaleInMilliseconds,omitempty" tf:"unit_timescale_in_milliseconds,omitempty"`
}

type ServicesAccountFilterPresentationTimeRangeObservation struct {

	// The absolute end time boundary. Applies to Video on Demand (VoD).
	// For the Live Streaming presentation, it is silently ignored and applied when the presentation ends and the stream becomes VoD. This is a long value that represents an absolute end point of the presentation, rounded to the closest next GOP start. The unit is defined by unit_timescale_in_milliseconds, so an end_in_units of 180 would be for 3 minutes. Use start_in_units and end_in_units to trim the fragments that will be in the playlist (manifest). For example, start_in_units set to 20 and end_in_units set to 60 using unit_timescale_in_milliseconds in 1000 will generate a playlist that contains fragments from between 20 seconds and 60 seconds of the VoD presentation. If a fragment straddles the boundary, the entire fragment will be included in the manifest.
	EndInUnits *float64 `json:"endInUnits,omitempty" tf:"end_in_units,omitempty"`

	// Indicates whether the end_in_units property must be present. If true, end_in_units must be specified or a bad request code is returned. Applies to Live Streaming only. Allowed values: false, true.
	ForceEnd *bool `json:"forceEnd,omitempty" tf:"force_end,omitempty"`

	// The relative to end right edge. Applies to Live Streaming only.
	// This value defines the latest live position that a client can seek to. Using this property, you can delay live playback position and create a server-side buffer for players. The unit is defined by unit_timescale_in_milliseconds. The maximum live back off duration is 300 seconds. For example, a value of 20 means that the latest available content is 20 seconds delayed from the real live edge.
	LiveBackoffInUnits *float64 `json:"liveBackoffInUnits,omitempty" tf:"live_backoff_in_units,omitempty"`

	// The relative to end sliding window. Applies to Live Streaming only. Use presentation_window_in_units to apply a sliding window of fragments to include in a playlist. The unit is defined by unit_timescale_in_milliseconds. For example, set presentation_window_in_units to 120 to apply a two-minute sliding window. Media within 2 minutes of the live edge will be included in the playlist. If a fragment straddles the boundary, the entire fragment will be included in the playlist. The minimum presentation window duration is 60 seconds.
	PresentationWindowInUnits *float64 `json:"presentationWindowInUnits,omitempty" tf:"presentation_window_in_units,omitempty"`

	// The absolute start time boundary. Applies to Video on Demand (VoD) or Live Streaming. This is a long value that represents an absolute start point of the stream. The value gets rounded to the closest next GOP start. The unit is defined by unit_timescale_in_milliseconds, so a start_in_units of 15 would be for 15 seconds. Use start_in_units and end_in_units to trim the fragments that will be in the playlist (manifest). For example, start_in_units set to 20 and end_in_units set to 60 using unit_timescale_in_milliseconds in 1000 will generate a playlist that contains fragments from between 20 seconds and 60 seconds of the VoD presentation. If a fragment straddles the boundary, the entire fragment will be included in the manifest.
	StartInUnits *float64 `json:"startInUnits,omitempty" tf:"start_in_units,omitempty"`

	// Specified as the number of milliseconds in one unit timescale. For example, if you want to set a start_in_units at 30 seconds, you would use a value of 30 when using the unit_timescale_in_milliseconds in 1000. Or if you want to set start_in_units in 30 milliseconds, you would use a value of 30 when using the unit_timescale_in_milliseconds in 1. Applies timescale to start_in_units, start_timescale and presentation_window_in_timescale and live_backoff_in_timescale.
	UnitTimescaleInMilliseconds *float64 `json:"unitTimescaleInMilliseconds,omitempty" tf:"unit_timescale_in_milliseconds,omitempty"`
}

type ServicesAccountFilterPresentationTimeRangeParameters struct {

	// The absolute end time boundary. Applies to Video on Demand (VoD).
	// For the Live Streaming presentation, it is silently ignored and applied when the presentation ends and the stream becomes VoD. This is a long value that represents an absolute end point of the presentation, rounded to the closest next GOP start. The unit is defined by unit_timescale_in_milliseconds, so an end_in_units of 180 would be for 3 minutes. Use start_in_units and end_in_units to trim the fragments that will be in the playlist (manifest). For example, start_in_units set to 20 and end_in_units set to 60 using unit_timescale_in_milliseconds in 1000 will generate a playlist that contains fragments from between 20 seconds and 60 seconds of the VoD presentation. If a fragment straddles the boundary, the entire fragment will be included in the manifest.
	// +kubebuilder:validation:Optional
	EndInUnits *float64 `json:"endInUnits,omitempty" tf:"end_in_units,omitempty"`

	// Indicates whether the end_in_units property must be present. If true, end_in_units must be specified or a bad request code is returned. Applies to Live Streaming only. Allowed values: false, true.
	// +kubebuilder:validation:Optional
	ForceEnd *bool `json:"forceEnd,omitempty" tf:"force_end,omitempty"`

	// The relative to end right edge. Applies to Live Streaming only.
	// This value defines the latest live position that a client can seek to. Using this property, you can delay live playback position and create a server-side buffer for players. The unit is defined by unit_timescale_in_milliseconds. The maximum live back off duration is 300 seconds. For example, a value of 20 means that the latest available content is 20 seconds delayed from the real live edge.
	// +kubebuilder:validation:Optional
	LiveBackoffInUnits *float64 `json:"liveBackoffInUnits,omitempty" tf:"live_backoff_in_units,omitempty"`

	// The relative to end sliding window. Applies to Live Streaming only. Use presentation_window_in_units to apply a sliding window of fragments to include in a playlist. The unit is defined by unit_timescale_in_milliseconds. For example, set presentation_window_in_units to 120 to apply a two-minute sliding window. Media within 2 minutes of the live edge will be included in the playlist. If a fragment straddles the boundary, the entire fragment will be included in the playlist. The minimum presentation window duration is 60 seconds.
	// +kubebuilder:validation:Optional
	PresentationWindowInUnits *float64 `json:"presentationWindowInUnits,omitempty" tf:"presentation_window_in_units,omitempty"`

	// The absolute start time boundary. Applies to Video on Demand (VoD) or Live Streaming. This is a long value that represents an absolute start point of the stream. The value gets rounded to the closest next GOP start. The unit is defined by unit_timescale_in_milliseconds, so a start_in_units of 15 would be for 15 seconds. Use start_in_units and end_in_units to trim the fragments that will be in the playlist (manifest). For example, start_in_units set to 20 and end_in_units set to 60 using unit_timescale_in_milliseconds in 1000 will generate a playlist that contains fragments from between 20 seconds and 60 seconds of the VoD presentation. If a fragment straddles the boundary, the entire fragment will be included in the manifest.
	// +kubebuilder:validation:Optional
	StartInUnits *float64 `json:"startInUnits,omitempty" tf:"start_in_units,omitempty"`

	// Specified as the number of milliseconds in one unit timescale. For example, if you want to set a start_in_units at 30 seconds, you would use a value of 30 when using the unit_timescale_in_milliseconds in 1000. Or if you want to set start_in_units in 30 milliseconds, you would use a value of 30 when using the unit_timescale_in_milliseconds in 1. Applies timescale to start_in_units, start_timescale and presentation_window_in_timescale and live_backoff_in_timescale.
	// +kubebuilder:validation:Optional
	UnitTimescaleInMilliseconds *float64 `json:"unitTimescaleInMilliseconds,omitempty" tf:"unit_timescale_in_milliseconds,omitempty"`
}

type ServicesAccountFilterTrackSelectionInitParameters struct {

	// One or more selection blocks as defined above.
	Condition []TrackSelectionConditionInitParameters `json:"condition,omitempty" tf:"condition,omitempty"`
}

type ServicesAccountFilterTrackSelectionObservation struct {

	// One or more selection blocks as defined above.
	Condition []TrackSelectionConditionObservation `json:"condition,omitempty" tf:"condition,omitempty"`
}

type ServicesAccountFilterTrackSelectionParameters struct {

	// One or more selection blocks as defined above.
	// +kubebuilder:validation:Optional
	Condition []TrackSelectionConditionParameters `json:"condition,omitempty" tf:"condition,omitempty"`
}

type TrackSelectionConditionInitParameters struct {

	// The condition operation to test a track property against. Supported values are Equal and NotEqual.
	Operation *string `json:"operation,omitempty" tf:"operation,omitempty"`

	// The track property to compare. Supported values are Bitrate, FourCC, Language, Name and Type. Check documentation for more details.
	Property *string `json:"property,omitempty" tf:"property,omitempty"`

	// The track property value to match or not match.
	Value *string `json:"value,omitempty" tf:"value,omitempty"`
}

type TrackSelectionConditionObservation struct {

	// The condition operation to test a track property against. Supported values are Equal and NotEqual.
	Operation *string `json:"operation,omitempty" tf:"operation,omitempty"`

	// The track property to compare. Supported values are Bitrate, FourCC, Language, Name and Type. Check documentation for more details.
	Property *string `json:"property,omitempty" tf:"property,omitempty"`

	// The track property value to match or not match.
	Value *string `json:"value,omitempty" tf:"value,omitempty"`
}

type TrackSelectionConditionParameters struct {

	// The condition operation to test a track property against. Supported values are Equal and NotEqual.
	// +kubebuilder:validation:Optional
	Operation *string `json:"operation,omitempty" tf:"operation,omitempty"`

	// The track property to compare. Supported values are Bitrate, FourCC, Language, Name and Type. Check documentation for more details.
	// +kubebuilder:validation:Optional
	Property *string `json:"property,omitempty" tf:"property,omitempty"`

	// The track property value to match or not match.
	// +kubebuilder:validation:Optional
	Value *string `json:"value,omitempty" tf:"value,omitempty"`
}

// ServicesAccountFilterSpec defines the desired state of ServicesAccountFilter
type ServicesAccountFilterSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ServicesAccountFilterParameters `json:"forProvider"`
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
	InitProvider ServicesAccountFilterInitParameters `json:"initProvider,omitempty"`
}

// ServicesAccountFilterStatus defines the observed state of ServicesAccountFilter.
type ServicesAccountFilterStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ServicesAccountFilterObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// ServicesAccountFilter is the Schema for the ServicesAccountFilters API. Manages a Media Services Account Filter.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type ServicesAccountFilter struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ServicesAccountFilterSpec   `json:"spec"`
	Status            ServicesAccountFilterStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ServicesAccountFilterList contains a list of ServicesAccountFilters
type ServicesAccountFilterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ServicesAccountFilter `json:"items"`
}

// Repository type metadata.
var (
	ServicesAccountFilter_Kind             = "ServicesAccountFilter"
	ServicesAccountFilter_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: ServicesAccountFilter_Kind}.String()
	ServicesAccountFilter_KindAPIVersion   = ServicesAccountFilter_Kind + "." + CRDGroupVersion.String()
	ServicesAccountFilter_GroupVersionKind = CRDGroupVersion.WithKind(ServicesAccountFilter_Kind)
)

func init() {
	SchemeBuilder.Register(&ServicesAccountFilter{}, &ServicesAccountFilterList{})
}
