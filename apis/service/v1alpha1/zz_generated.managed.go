/*
Copyright 2021 The Crossplane Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
// Code generated by angryjet. DO NOT EDIT.

package v1alpha1

import xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"

// GetCondition of this ServiceDiscoveryHttpNamespace.
func (mg *ServiceDiscoveryHttpNamespace) GetCondition(ct xpv1.ConditionType) xpv1.Condition {
	return mg.Status.GetCondition(ct)
}

// GetDeletionPolicy of this ServiceDiscoveryHttpNamespace.
func (mg *ServiceDiscoveryHttpNamespace) GetDeletionPolicy() xpv1.DeletionPolicy {
	return mg.Spec.DeletionPolicy
}

// GetProviderConfigReference of this ServiceDiscoveryHttpNamespace.
func (mg *ServiceDiscoveryHttpNamespace) GetProviderConfigReference() *xpv1.Reference {
	return mg.Spec.ProviderConfigReference
}

/*
GetProviderReference of this ServiceDiscoveryHttpNamespace.
Deprecated: Use GetProviderConfigReference.
*/
func (mg *ServiceDiscoveryHttpNamespace) GetProviderReference() *xpv1.Reference {
	return mg.Spec.ProviderReference
}

// GetWriteConnectionSecretToReference of this ServiceDiscoveryHttpNamespace.
func (mg *ServiceDiscoveryHttpNamespace) GetWriteConnectionSecretToReference() *xpv1.SecretReference {
	return mg.Spec.WriteConnectionSecretToReference
}

// SetConditions of this ServiceDiscoveryHttpNamespace.
func (mg *ServiceDiscoveryHttpNamespace) SetConditions(c ...xpv1.Condition) {
	mg.Status.SetConditions(c...)
}

// SetDeletionPolicy of this ServiceDiscoveryHttpNamespace.
func (mg *ServiceDiscoveryHttpNamespace) SetDeletionPolicy(r xpv1.DeletionPolicy) {
	mg.Spec.DeletionPolicy = r
}

// SetProviderConfigReference of this ServiceDiscoveryHttpNamespace.
func (mg *ServiceDiscoveryHttpNamespace) SetProviderConfigReference(r *xpv1.Reference) {
	mg.Spec.ProviderConfigReference = r
}

/*
SetProviderReference of this ServiceDiscoveryHttpNamespace.
Deprecated: Use SetProviderConfigReference.
*/
func (mg *ServiceDiscoveryHttpNamespace) SetProviderReference(r *xpv1.Reference) {
	mg.Spec.ProviderReference = r
}

// SetWriteConnectionSecretToReference of this ServiceDiscoveryHttpNamespace.
func (mg *ServiceDiscoveryHttpNamespace) SetWriteConnectionSecretToReference(r *xpv1.SecretReference) {
	mg.Spec.WriteConnectionSecretToReference = r
}

// GetCondition of this ServiceDiscoveryPrivateDnsNamespace.
func (mg *ServiceDiscoveryPrivateDnsNamespace) GetCondition(ct xpv1.ConditionType) xpv1.Condition {
	return mg.Status.GetCondition(ct)
}

// GetDeletionPolicy of this ServiceDiscoveryPrivateDnsNamespace.
func (mg *ServiceDiscoveryPrivateDnsNamespace) GetDeletionPolicy() xpv1.DeletionPolicy {
	return mg.Spec.DeletionPolicy
}

// GetProviderConfigReference of this ServiceDiscoveryPrivateDnsNamespace.
func (mg *ServiceDiscoveryPrivateDnsNamespace) GetProviderConfigReference() *xpv1.Reference {
	return mg.Spec.ProviderConfigReference
}

/*
GetProviderReference of this ServiceDiscoveryPrivateDnsNamespace.
Deprecated: Use GetProviderConfigReference.
*/
func (mg *ServiceDiscoveryPrivateDnsNamespace) GetProviderReference() *xpv1.Reference {
	return mg.Spec.ProviderReference
}

// GetWriteConnectionSecretToReference of this ServiceDiscoveryPrivateDnsNamespace.
func (mg *ServiceDiscoveryPrivateDnsNamespace) GetWriteConnectionSecretToReference() *xpv1.SecretReference {
	return mg.Spec.WriteConnectionSecretToReference
}

// SetConditions of this ServiceDiscoveryPrivateDnsNamespace.
func (mg *ServiceDiscoveryPrivateDnsNamespace) SetConditions(c ...xpv1.Condition) {
	mg.Status.SetConditions(c...)
}

// SetDeletionPolicy of this ServiceDiscoveryPrivateDnsNamespace.
func (mg *ServiceDiscoveryPrivateDnsNamespace) SetDeletionPolicy(r xpv1.DeletionPolicy) {
	mg.Spec.DeletionPolicy = r
}

// SetProviderConfigReference of this ServiceDiscoveryPrivateDnsNamespace.
func (mg *ServiceDiscoveryPrivateDnsNamespace) SetProviderConfigReference(r *xpv1.Reference) {
	mg.Spec.ProviderConfigReference = r
}

/*
SetProviderReference of this ServiceDiscoveryPrivateDnsNamespace.
Deprecated: Use SetProviderConfigReference.
*/
func (mg *ServiceDiscoveryPrivateDnsNamespace) SetProviderReference(r *xpv1.Reference) {
	mg.Spec.ProviderReference = r
}

// SetWriteConnectionSecretToReference of this ServiceDiscoveryPrivateDnsNamespace.
func (mg *ServiceDiscoveryPrivateDnsNamespace) SetWriteConnectionSecretToReference(r *xpv1.SecretReference) {
	mg.Spec.WriteConnectionSecretToReference = r
}

// GetCondition of this ServiceDiscoveryPublicDnsNamespace.
func (mg *ServiceDiscoveryPublicDnsNamespace) GetCondition(ct xpv1.ConditionType) xpv1.Condition {
	return mg.Status.GetCondition(ct)
}

// GetDeletionPolicy of this ServiceDiscoveryPublicDnsNamespace.
func (mg *ServiceDiscoveryPublicDnsNamespace) GetDeletionPolicy() xpv1.DeletionPolicy {
	return mg.Spec.DeletionPolicy
}

// GetProviderConfigReference of this ServiceDiscoveryPublicDnsNamespace.
func (mg *ServiceDiscoveryPublicDnsNamespace) GetProviderConfigReference() *xpv1.Reference {
	return mg.Spec.ProviderConfigReference
}

/*
GetProviderReference of this ServiceDiscoveryPublicDnsNamespace.
Deprecated: Use GetProviderConfigReference.
*/
func (mg *ServiceDiscoveryPublicDnsNamespace) GetProviderReference() *xpv1.Reference {
	return mg.Spec.ProviderReference
}

// GetWriteConnectionSecretToReference of this ServiceDiscoveryPublicDnsNamespace.
func (mg *ServiceDiscoveryPublicDnsNamespace) GetWriteConnectionSecretToReference() *xpv1.SecretReference {
	return mg.Spec.WriteConnectionSecretToReference
}

// SetConditions of this ServiceDiscoveryPublicDnsNamespace.
func (mg *ServiceDiscoveryPublicDnsNamespace) SetConditions(c ...xpv1.Condition) {
	mg.Status.SetConditions(c...)
}

// SetDeletionPolicy of this ServiceDiscoveryPublicDnsNamespace.
func (mg *ServiceDiscoveryPublicDnsNamespace) SetDeletionPolicy(r xpv1.DeletionPolicy) {
	mg.Spec.DeletionPolicy = r
}

// SetProviderConfigReference of this ServiceDiscoveryPublicDnsNamespace.
func (mg *ServiceDiscoveryPublicDnsNamespace) SetProviderConfigReference(r *xpv1.Reference) {
	mg.Spec.ProviderConfigReference = r
}

/*
SetProviderReference of this ServiceDiscoveryPublicDnsNamespace.
Deprecated: Use SetProviderConfigReference.
*/
func (mg *ServiceDiscoveryPublicDnsNamespace) SetProviderReference(r *xpv1.Reference) {
	mg.Spec.ProviderReference = r
}

// SetWriteConnectionSecretToReference of this ServiceDiscoveryPublicDnsNamespace.
func (mg *ServiceDiscoveryPublicDnsNamespace) SetWriteConnectionSecretToReference(r *xpv1.SecretReference) {
	mg.Spec.WriteConnectionSecretToReference = r
}

// GetCondition of this ServiceDiscoveryService.
func (mg *ServiceDiscoveryService) GetCondition(ct xpv1.ConditionType) xpv1.Condition {
	return mg.Status.GetCondition(ct)
}

// GetDeletionPolicy of this ServiceDiscoveryService.
func (mg *ServiceDiscoveryService) GetDeletionPolicy() xpv1.DeletionPolicy {
	return mg.Spec.DeletionPolicy
}

// GetProviderConfigReference of this ServiceDiscoveryService.
func (mg *ServiceDiscoveryService) GetProviderConfigReference() *xpv1.Reference {
	return mg.Spec.ProviderConfigReference
}

/*
GetProviderReference of this ServiceDiscoveryService.
Deprecated: Use GetProviderConfigReference.
*/
func (mg *ServiceDiscoveryService) GetProviderReference() *xpv1.Reference {
	return mg.Spec.ProviderReference
}

// GetWriteConnectionSecretToReference of this ServiceDiscoveryService.
func (mg *ServiceDiscoveryService) GetWriteConnectionSecretToReference() *xpv1.SecretReference {
	return mg.Spec.WriteConnectionSecretToReference
}

// SetConditions of this ServiceDiscoveryService.
func (mg *ServiceDiscoveryService) SetConditions(c ...xpv1.Condition) {
	mg.Status.SetConditions(c...)
}

// SetDeletionPolicy of this ServiceDiscoveryService.
func (mg *ServiceDiscoveryService) SetDeletionPolicy(r xpv1.DeletionPolicy) {
	mg.Spec.DeletionPolicy = r
}

// SetProviderConfigReference of this ServiceDiscoveryService.
func (mg *ServiceDiscoveryService) SetProviderConfigReference(r *xpv1.Reference) {
	mg.Spec.ProviderConfigReference = r
}

/*
SetProviderReference of this ServiceDiscoveryService.
Deprecated: Use SetProviderConfigReference.
*/
func (mg *ServiceDiscoveryService) SetProviderReference(r *xpv1.Reference) {
	mg.Spec.ProviderReference = r
}

// SetWriteConnectionSecretToReference of this ServiceDiscoveryService.
func (mg *ServiceDiscoveryService) SetWriteConnectionSecretToReference(r *xpv1.SecretReference) {
	mg.Spec.WriteConnectionSecretToReference = r
}