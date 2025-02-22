/*
Copyright 2022 Upbound Inc.
*/
// Code generated by angryjet. DO NOT EDIT.

package v1beta1

import xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"

// GetCondition of this IOTHubDeviceUpdateAccount.
func (mg *IOTHubDeviceUpdateAccount) GetCondition(ct xpv1.ConditionType) xpv1.Condition {
	return mg.Status.GetCondition(ct)
}

// GetDeletionPolicy of this IOTHubDeviceUpdateAccount.
func (mg *IOTHubDeviceUpdateAccount) GetDeletionPolicy() xpv1.DeletionPolicy {
	return mg.Spec.DeletionPolicy
}

// GetManagementPolicies of this IOTHubDeviceUpdateAccount.
func (mg *IOTHubDeviceUpdateAccount) GetManagementPolicies() xpv1.ManagementPolicies {
	return mg.Spec.ManagementPolicies
}

// GetProviderConfigReference of this IOTHubDeviceUpdateAccount.
func (mg *IOTHubDeviceUpdateAccount) GetProviderConfigReference() *xpv1.Reference {
	return mg.Spec.ProviderConfigReference
}

/*
GetProviderReference of this IOTHubDeviceUpdateAccount.
Deprecated: Use GetProviderConfigReference.
*/
func (mg *IOTHubDeviceUpdateAccount) GetProviderReference() *xpv1.Reference {
	return mg.Spec.ProviderReference
}

// GetPublishConnectionDetailsTo of this IOTHubDeviceUpdateAccount.
func (mg *IOTHubDeviceUpdateAccount) GetPublishConnectionDetailsTo() *xpv1.PublishConnectionDetailsTo {
	return mg.Spec.PublishConnectionDetailsTo
}

// GetWriteConnectionSecretToReference of this IOTHubDeviceUpdateAccount.
func (mg *IOTHubDeviceUpdateAccount) GetWriteConnectionSecretToReference() *xpv1.SecretReference {
	return mg.Spec.WriteConnectionSecretToReference
}

// SetConditions of this IOTHubDeviceUpdateAccount.
func (mg *IOTHubDeviceUpdateAccount) SetConditions(c ...xpv1.Condition) {
	mg.Status.SetConditions(c...)
}

// SetDeletionPolicy of this IOTHubDeviceUpdateAccount.
func (mg *IOTHubDeviceUpdateAccount) SetDeletionPolicy(r xpv1.DeletionPolicy) {
	mg.Spec.DeletionPolicy = r
}

// SetManagementPolicies of this IOTHubDeviceUpdateAccount.
func (mg *IOTHubDeviceUpdateAccount) SetManagementPolicies(r xpv1.ManagementPolicies) {
	mg.Spec.ManagementPolicies = r
}

// SetProviderConfigReference of this IOTHubDeviceUpdateAccount.
func (mg *IOTHubDeviceUpdateAccount) SetProviderConfigReference(r *xpv1.Reference) {
	mg.Spec.ProviderConfigReference = r
}

/*
SetProviderReference of this IOTHubDeviceUpdateAccount.
Deprecated: Use SetProviderConfigReference.
*/
func (mg *IOTHubDeviceUpdateAccount) SetProviderReference(r *xpv1.Reference) {
	mg.Spec.ProviderReference = r
}

// SetPublishConnectionDetailsTo of this IOTHubDeviceUpdateAccount.
func (mg *IOTHubDeviceUpdateAccount) SetPublishConnectionDetailsTo(r *xpv1.PublishConnectionDetailsTo) {
	mg.Spec.PublishConnectionDetailsTo = r
}

// SetWriteConnectionSecretToReference of this IOTHubDeviceUpdateAccount.
func (mg *IOTHubDeviceUpdateAccount) SetWriteConnectionSecretToReference(r *xpv1.SecretReference) {
	mg.Spec.WriteConnectionSecretToReference = r
}

// GetCondition of this IOTHubDeviceUpdateInstance.
func (mg *IOTHubDeviceUpdateInstance) GetCondition(ct xpv1.ConditionType) xpv1.Condition {
	return mg.Status.GetCondition(ct)
}

// GetDeletionPolicy of this IOTHubDeviceUpdateInstance.
func (mg *IOTHubDeviceUpdateInstance) GetDeletionPolicy() xpv1.DeletionPolicy {
	return mg.Spec.DeletionPolicy
}

// GetManagementPolicies of this IOTHubDeviceUpdateInstance.
func (mg *IOTHubDeviceUpdateInstance) GetManagementPolicies() xpv1.ManagementPolicies {
	return mg.Spec.ManagementPolicies
}

// GetProviderConfigReference of this IOTHubDeviceUpdateInstance.
func (mg *IOTHubDeviceUpdateInstance) GetProviderConfigReference() *xpv1.Reference {
	return mg.Spec.ProviderConfigReference
}

/*
GetProviderReference of this IOTHubDeviceUpdateInstance.
Deprecated: Use GetProviderConfigReference.
*/
func (mg *IOTHubDeviceUpdateInstance) GetProviderReference() *xpv1.Reference {
	return mg.Spec.ProviderReference
}

// GetPublishConnectionDetailsTo of this IOTHubDeviceUpdateInstance.
func (mg *IOTHubDeviceUpdateInstance) GetPublishConnectionDetailsTo() *xpv1.PublishConnectionDetailsTo {
	return mg.Spec.PublishConnectionDetailsTo
}

// GetWriteConnectionSecretToReference of this IOTHubDeviceUpdateInstance.
func (mg *IOTHubDeviceUpdateInstance) GetWriteConnectionSecretToReference() *xpv1.SecretReference {
	return mg.Spec.WriteConnectionSecretToReference
}

// SetConditions of this IOTHubDeviceUpdateInstance.
func (mg *IOTHubDeviceUpdateInstance) SetConditions(c ...xpv1.Condition) {
	mg.Status.SetConditions(c...)
}

// SetDeletionPolicy of this IOTHubDeviceUpdateInstance.
func (mg *IOTHubDeviceUpdateInstance) SetDeletionPolicy(r xpv1.DeletionPolicy) {
	mg.Spec.DeletionPolicy = r
}

// SetManagementPolicies of this IOTHubDeviceUpdateInstance.
func (mg *IOTHubDeviceUpdateInstance) SetManagementPolicies(r xpv1.ManagementPolicies) {
	mg.Spec.ManagementPolicies = r
}

// SetProviderConfigReference of this IOTHubDeviceUpdateInstance.
func (mg *IOTHubDeviceUpdateInstance) SetProviderConfigReference(r *xpv1.Reference) {
	mg.Spec.ProviderConfigReference = r
}

/*
SetProviderReference of this IOTHubDeviceUpdateInstance.
Deprecated: Use SetProviderConfigReference.
*/
func (mg *IOTHubDeviceUpdateInstance) SetProviderReference(r *xpv1.Reference) {
	mg.Spec.ProviderReference = r
}

// SetPublishConnectionDetailsTo of this IOTHubDeviceUpdateInstance.
func (mg *IOTHubDeviceUpdateInstance) SetPublishConnectionDetailsTo(r *xpv1.PublishConnectionDetailsTo) {
	mg.Spec.PublishConnectionDetailsTo = r
}

// SetWriteConnectionSecretToReference of this IOTHubDeviceUpdateInstance.
func (mg *IOTHubDeviceUpdateInstance) SetWriteConnectionSecretToReference(r *xpv1.SecretReference) {
	mg.Spec.WriteConnectionSecretToReference = r
}
