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

// +kubebuilder:object:generate=true
// +groupName=ec2.aws.tf.crossplane.io
// +versionName=v1alpha1

package v1alpha1

import (
	xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	apis "github.com/crossplane-contrib/provider-tf-aws/apis"
	v1alpha1 "github.com/crossplane-contrib/provider-tf-aws/apis/ec2/v1alpha1"
)

type Ec2TransitGatewayPeeringAttachmentObservation struct {
}

type Ec2TransitGatewayPeeringAttachmentParameters struct {
	PeerAccountId *string `json:"peerAccountId,omitempty" tf:"peer_account_id"`

	PeerRegion string `json:"peerRegion" tf:"peer_region"`

	PeerTransitGatewayId string `json:"peerTransitGatewayId" tf:"peer_transit_gateway_id"`

	Tags map[string]string `json:"tags,omitempty" tf:"tags"`

	TagsAll map[string]string `json:"tagsAll,omitempty" tf:"tags_all"`

	TransitGatewayId string `json:"transitGatewayId" tf:"transit_gateway_id"`
}

// Ec2TransitGatewayPeeringAttachmentSpec defines the desired state of Ec2TransitGatewayPeeringAttachment
type Ec2TransitGatewayPeeringAttachmentSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       Ec2TransitGatewayPeeringAttachmentParameters `json:"forProvider"`
}

// Ec2TransitGatewayPeeringAttachmentStatus defines the observed state of Ec2TransitGatewayPeeringAttachment.
type Ec2TransitGatewayPeeringAttachmentStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          Ec2TransitGatewayPeeringAttachmentObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Ec2TransitGatewayPeeringAttachment is the Schema for the Ec2TransitGatewayPeeringAttachments API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type Ec2TransitGatewayPeeringAttachment struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              Ec2TransitGatewayPeeringAttachmentSpec   `json:"spec"`
	Status            Ec2TransitGatewayPeeringAttachmentStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// Ec2TransitGatewayPeeringAttachmentList contains a list of Ec2TransitGatewayPeeringAttachments
type Ec2TransitGatewayPeeringAttachmentList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Ec2TransitGatewayPeeringAttachment `json:"items"`
}

// Repository type metadata.
var (
	Ec2TransitGatewayPeeringAttachmentKind             = "Ec2TransitGatewayPeeringAttachment"
	Ec2TransitGatewayPeeringAttachmentGroupKind        = schema.GroupKind{Group: v1alpha1.Group, Kind: Ec2TransitGatewayPeeringAttachmentKind}.String()
	Ec2TransitGatewayPeeringAttachmentKindAPIVersion   = Ec2TransitGatewayPeeringAttachmentKind + "." + v1alpha1.GroupVersion.String()
	Ec2TransitGatewayPeeringAttachmentGroupVersionKind = v1alpha1.GroupVersion.WithKind(Ec2TransitGatewayPeeringAttachmentKind)
)

func init() {
	apis.SchemaMap[v1alpha1.GroupVersion] = append(apis.SchemaMap[v1alpha1.GroupVersion], &Ec2TransitGatewayPeeringAttachment{}, &Ec2TransitGatewayPeeringAttachmentList{})
}
