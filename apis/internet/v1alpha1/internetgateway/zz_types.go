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
// +groupName=internet.aws.tf.crossplane.io
// +versionName=v1alpha1

package v1alpha1

import (
	xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	apis "github.com/crossplane-contrib/provider-tf-aws/apis"
	v1alpha1 "github.com/crossplane-contrib/provider-tf-aws/apis/internet/v1alpha1"
)

type InternetGatewayObservation struct {
	Arn string `json:"arn" tf:"arn"`

	OwnerId string `json:"ownerId" tf:"owner_id"`
}

type InternetGatewayParameters struct {
	Tags map[string]string `json:"tags,omitempty" tf:"tags"`

	TagsAll map[string]string `json:"tagsAll,omitempty" tf:"tags_all"`

	VpcId *string `json:"vpcId,omitempty" tf:"vpc_id"`
}

// InternetGatewaySpec defines the desired state of InternetGateway
type InternetGatewaySpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       InternetGatewayParameters `json:"forProvider"`
}

// InternetGatewayStatus defines the observed state of InternetGateway.
type InternetGatewayStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          InternetGatewayObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// InternetGateway is the Schema for the InternetGateways API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type InternetGateway struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              InternetGatewaySpec   `json:"spec"`
	Status            InternetGatewayStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// InternetGatewayList contains a list of InternetGateways
type InternetGatewayList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []InternetGateway `json:"items"`
}

// Repository type metadata.
var (
	InternetGatewayKind             = "InternetGateway"
	InternetGatewayGroupKind        = schema.GroupKind{Group: v1alpha1.Group, Kind: InternetGatewayKind}.String()
	InternetGatewayKindAPIVersion   = InternetGatewayKind + "." + v1alpha1.GroupVersion.String()
	InternetGatewayGroupVersionKind = v1alpha1.GroupVersion.WithKind(InternetGatewayKind)
)

func init() {
	apis.SchemaMap[v1alpha1.GroupVersion] = append(apis.SchemaMap[v1alpha1.GroupVersion], &InternetGateway{}, &InternetGatewayList{})
}
