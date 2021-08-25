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
// +groupName=db.aws.tf.crossplane.io
// +versionName=v1alpha1

package v1alpha1

import (
	xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	apis "github.com/crossplane-contrib/provider-tf-aws/apis"
	v1alpha1 "github.com/crossplane-contrib/provider-tf-aws/apis/db/v1alpha1"
)

type DbProxyEndpointObservation struct {
	Arn string `json:"arn" tf:"arn"`

	Endpoint string `json:"endpoint" tf:"endpoint"`

	IsDefault bool `json:"isDefault" tf:"is_default"`

	VpcId string `json:"vpcId" tf:"vpc_id"`
}

type DbProxyEndpointParameters struct {
	DbProxyEndpointName string `json:"dbProxyEndpointName" tf:"db_proxy_endpoint_name"`

	DbProxyName string `json:"dbProxyName" tf:"db_proxy_name"`

	Tags map[string]string `json:"tags,omitempty" tf:"tags"`

	TagsAll map[string]string `json:"tagsAll,omitempty" tf:"tags_all"`

	TargetRole *string `json:"targetRole,omitempty" tf:"target_role"`

	VpcSecurityGroupIds []string `json:"vpcSecurityGroupIds,omitempty" tf:"vpc_security_group_ids"`

	VpcSubnetIds []string `json:"vpcSubnetIds" tf:"vpc_subnet_ids"`
}

// DbProxyEndpointSpec defines the desired state of DbProxyEndpoint
type DbProxyEndpointSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       DbProxyEndpointParameters `json:"forProvider"`
}

// DbProxyEndpointStatus defines the observed state of DbProxyEndpoint.
type DbProxyEndpointStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          DbProxyEndpointObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// DbProxyEndpoint is the Schema for the DbProxyEndpoints API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type DbProxyEndpoint struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DbProxyEndpointSpec   `json:"spec"`
	Status            DbProxyEndpointStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// DbProxyEndpointList contains a list of DbProxyEndpoints
type DbProxyEndpointList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DbProxyEndpoint `json:"items"`
}

// Repository type metadata.
var (
	DbProxyEndpointKind             = "DbProxyEndpoint"
	DbProxyEndpointGroupKind        = schema.GroupKind{Group: v1alpha1.Group, Kind: DbProxyEndpointKind}.String()
	DbProxyEndpointKindAPIVersion   = DbProxyEndpointKind + "." + v1alpha1.GroupVersion.String()
	DbProxyEndpointGroupVersionKind = v1alpha1.GroupVersion.WithKind(DbProxyEndpointKind)
)

func init() {
	apis.SchemaMap[v1alpha1.GroupVersion] = append(apis.SchemaMap[v1alpha1.GroupVersion], &DbProxyEndpoint{}, &DbProxyEndpointList{})
}
