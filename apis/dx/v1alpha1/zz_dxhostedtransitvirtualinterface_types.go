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

// Code generated by terrajet. DO NOT EDIT.

package v1alpha1

import (
	xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

type DxHostedTransitVirtualInterfaceObservation struct {
	AmazonSideAsn string `json:"amazonSideAsn" tf:"amazon_side_asn"`

	Arn string `json:"arn" tf:"arn"`

	AwsDevice string `json:"awsDevice" tf:"aws_device"`

	JumboFrameCapable bool `json:"jumboFrameCapable" tf:"jumbo_frame_capable"`
}

type DxHostedTransitVirtualInterfaceParameters struct {

	// +kubebuilder:validation:Required
	AddressFamily string `json:"addressFamily" tf:"address_family"`

	// +kubebuilder:validation:Optional
	AmazonAddress *string `json:"amazonAddress,omitempty" tf:"amazon_address"`

	// +kubebuilder:validation:Required
	BgpAsn int64 `json:"bgpAsn" tf:"bgp_asn"`

	// +kubebuilder:validation:Optional
	BgpAuthKey *string `json:"bgpAuthKey,omitempty" tf:"bgp_auth_key"`

	// +kubebuilder:validation:Required
	ConnectionID string `json:"connectionId" tf:"connection_id"`

	// +kubebuilder:validation:Optional
	CustomerAddress *string `json:"customerAddress,omitempty" tf:"customer_address"`

	// +kubebuilder:validation:Optional
	Mtu *int64 `json:"mtu,omitempty" tf:"mtu"`

	// +kubebuilder:validation:Required
	Name string `json:"name" tf:"name"`

	// +kubebuilder:validation:Required
	OwnerAccountID string `json:"ownerAccountId" tf:"owner_account_id"`

	// Region is the region you'd like your resource to be created in.
	// +terrajet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region string `json:"region" tf:"-"`

	// +kubebuilder:validation:Required
	Vlan int64 `json:"vlan" tf:"vlan"`
}

// DxHostedTransitVirtualInterfaceSpec defines the desired state of DxHostedTransitVirtualInterface
type DxHostedTransitVirtualInterfaceSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       DxHostedTransitVirtualInterfaceParameters `json:"forProvider"`
}

// DxHostedTransitVirtualInterfaceStatus defines the observed state of DxHostedTransitVirtualInterface.
type DxHostedTransitVirtualInterfaceStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          DxHostedTransitVirtualInterfaceObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// DxHostedTransitVirtualInterface is the Schema for the DxHostedTransitVirtualInterfaces API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tfaws}
type DxHostedTransitVirtualInterface struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DxHostedTransitVirtualInterfaceSpec   `json:"spec"`
	Status            DxHostedTransitVirtualInterfaceStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// DxHostedTransitVirtualInterfaceList contains a list of DxHostedTransitVirtualInterfaces
type DxHostedTransitVirtualInterfaceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DxHostedTransitVirtualInterface `json:"items"`
}

// Repository type metadata.
var (
	DxHostedTransitVirtualInterfaceKind             = "DxHostedTransitVirtualInterface"
	DxHostedTransitVirtualInterfaceGroupKind        = schema.GroupKind{Group: Group, Kind: DxHostedTransitVirtualInterfaceKind}.String()
	DxHostedTransitVirtualInterfaceKindAPIVersion   = DxHostedTransitVirtualInterfaceKind + "." + GroupVersion.String()
	DxHostedTransitVirtualInterfaceGroupVersionKind = GroupVersion.WithKind(DxHostedTransitVirtualInterfaceKind)
)

func init() {
	SchemeBuilder.Register(&DxHostedTransitVirtualInterface{}, &DxHostedTransitVirtualInterfaceList{})
}