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
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type PeeringConnectionAccepterAccepterObservation struct {
}

type PeeringConnectionAccepterAccepterParameters struct {

	// +kubebuilder:validation:Optional
	AllowClassicLinkToRemoteVpc *bool `json:"allowClassicLinkToRemoteVpc,omitempty" tf:"allow_classic_link_to_remote_vpc,omitempty"`

	// +kubebuilder:validation:Optional
	AllowRemoteVpcDNSResolution *bool `json:"allowRemoteVpcDnsResolution,omitempty" tf:"allow_remote_vpc_dns_resolution,omitempty"`

	// +kubebuilder:validation:Optional
	AllowVpcToRemoteClassicLink *bool `json:"allowVpcToRemoteClassicLink,omitempty" tf:"allow_vpc_to_remote_classic_link,omitempty"`
}

type PeeringConnectionAccepterObservation struct {
	AcceptStatus *string `json:"acceptStatus,omitempty" tf:"accept_status,omitempty"`

	PeerOwnerID *string `json:"peerOwnerId,omitempty" tf:"peer_owner_id,omitempty"`

	PeerRegion *string `json:"peerRegion,omitempty" tf:"peer_region,omitempty"`

	PeerVpcID *string `json:"peerVpcId,omitempty" tf:"peer_vpc_id,omitempty"`

	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`

	VpcID *string `json:"vpcId,omitempty" tf:"vpc_id,omitempty"`
}

type PeeringConnectionAccepterParameters struct {

	// +kubebuilder:validation:Optional
	Accepter []PeeringConnectionAccepterAccepterParameters `json:"accepter,omitempty" tf:"accepter,omitempty"`

	// +kubebuilder:validation:Optional
	AutoAccept *bool `json:"autoAccept,omitempty" tf:"auto_accept,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +terrajet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// +kubebuilder:validation:Optional
	Requester []PeeringConnectionAccepterRequesterParameters `json:"requester,omitempty" tf:"requester,omitempty"`

	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// +kubebuilder:validation:Required
	VpcPeeringConnectionID *string `json:"vpcPeeringConnectionId" tf:"vpc_peering_connection_id,omitempty"`
}

type PeeringConnectionAccepterRequesterObservation struct {
}

type PeeringConnectionAccepterRequesterParameters struct {

	// +kubebuilder:validation:Optional
	AllowClassicLinkToRemoteVpc *bool `json:"allowClassicLinkToRemoteVpc,omitempty" tf:"allow_classic_link_to_remote_vpc,omitempty"`

	// +kubebuilder:validation:Optional
	AllowRemoteVpcDNSResolution *bool `json:"allowRemoteVpcDnsResolution,omitempty" tf:"allow_remote_vpc_dns_resolution,omitempty"`

	// +kubebuilder:validation:Optional
	AllowVpcToRemoteClassicLink *bool `json:"allowVpcToRemoteClassicLink,omitempty" tf:"allow_vpc_to_remote_classic_link,omitempty"`
}

// PeeringConnectionAccepterSpec defines the desired state of PeeringConnectionAccepter
type PeeringConnectionAccepterSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     PeeringConnectionAccepterParameters `json:"forProvider"`
}

// PeeringConnectionAccepterStatus defines the observed state of PeeringConnectionAccepter.
type PeeringConnectionAccepterStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        PeeringConnectionAccepterObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// PeeringConnectionAccepter is the Schema for the PeeringConnectionAccepters API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,awsjet}
type PeeringConnectionAccepter struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              PeeringConnectionAccepterSpec   `json:"spec"`
	Status            PeeringConnectionAccepterStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// PeeringConnectionAccepterList contains a list of PeeringConnectionAccepters
type PeeringConnectionAccepterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []PeeringConnectionAccepter `json:"items"`
}

// Repository type metadata.
var (
	PeeringConnectionAccepter_Kind             = "PeeringConnectionAccepter"
	PeeringConnectionAccepter_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: PeeringConnectionAccepter_Kind}.String()
	PeeringConnectionAccepter_KindAPIVersion   = PeeringConnectionAccepter_Kind + "." + CRDGroupVersion.String()
	PeeringConnectionAccepter_GroupVersionKind = CRDGroupVersion.WithKind(PeeringConnectionAccepter_Kind)
)

func init() {
	SchemeBuilder.Register(&PeeringConnectionAccepter{}, &PeeringConnectionAccepterList{})
}
