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
// +groupName=subnet.aws.tf.crossplane.io
// +versionName=v1alpha1

package v1alpha1

import (
	xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	apis "github.com/crossplane-contrib/provider-tf-aws/apis"
	v1alpha1 "github.com/crossplane-contrib/provider-tf-aws/apis/subnet/v1alpha1"
)

type SubnetObservation struct {
	Arn string `json:"arn" tf:"arn"`

	Ipv6CidrBlockAssociationId string `json:"ipv6CidrBlockAssociationId" tf:"ipv6_cidr_block_association_id"`

	OwnerId string `json:"ownerId" tf:"owner_id"`
}

type SubnetParameters struct {
	AssignIpv6AddressOnCreation *bool `json:"assignIpv6AddressOnCreation,omitempty" tf:"assign_ipv6_address_on_creation"`

	AvailabilityZone *string `json:"availabilityZone,omitempty" tf:"availability_zone"`

	AvailabilityZoneId *string `json:"availabilityZoneId,omitempty" tf:"availability_zone_id"`

	CidrBlock string `json:"cidrBlock" tf:"cidr_block"`

	CustomerOwnedIpv4Pool *string `json:"customerOwnedIpv4Pool,omitempty" tf:"customer_owned_ipv4_pool"`

	Ipv6CidrBlock *string `json:"ipv6CidrBlock,omitempty" tf:"ipv6_cidr_block"`

	MapCustomerOwnedIpOnLaunch *bool `json:"mapCustomerOwnedIpOnLaunch,omitempty" tf:"map_customer_owned_ip_on_launch"`

	MapPublicIpOnLaunch *bool `json:"mapPublicIpOnLaunch,omitempty" tf:"map_public_ip_on_launch"`

	OutpostArn *string `json:"outpostArn,omitempty" tf:"outpost_arn"`

	Tags map[string]string `json:"tags,omitempty" tf:"tags"`

	TagsAll map[string]string `json:"tagsAll,omitempty" tf:"tags_all"`

	VpcId string `json:"vpcId" tf:"vpc_id"`
}

// SubnetSpec defines the desired state of Subnet
type SubnetSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       SubnetParameters `json:"forProvider"`
}

// SubnetStatus defines the observed state of Subnet.
type SubnetStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          SubnetObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Subnet is the Schema for the Subnets API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type Subnet struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              SubnetSpec   `json:"spec"`
	Status            SubnetStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SubnetList contains a list of Subnets
type SubnetList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Subnet `json:"items"`
}

// Repository type metadata.
var (
	SubnetKind             = "Subnet"
	SubnetGroupKind        = schema.GroupKind{Group: v1alpha1.Group, Kind: SubnetKind}.String()
	SubnetKindAPIVersion   = SubnetKind + "." + v1alpha1.GroupVersion.String()
	SubnetGroupVersionKind = v1alpha1.GroupVersion.WithKind(SubnetKind)
)

func init() {
	apis.SchemaMap[v1alpha1.GroupVersion] = append(apis.SchemaMap[v1alpha1.GroupVersion], &Subnet{}, &SubnetList{})
}
