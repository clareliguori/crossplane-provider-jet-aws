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
// +groupName=dax.aws.tf.crossplane.io
// +versionName=v1alpha1

package v1alpha1

import (
	xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	apis "github.com/crossplane-contrib/provider-tf-aws/apis"
	v1alpha1 "github.com/crossplane-contrib/provider-tf-aws/apis/dax/v1alpha1"
)

type DaxClusterObservation struct {
	Arn string `json:"arn" tf:"arn"`

	ClusterAddress string `json:"clusterAddress" tf:"cluster_address"`

	ConfigurationEndpoint string `json:"configurationEndpoint" tf:"configuration_endpoint"`

	Nodes []NodesObservation `json:"nodes" tf:"nodes"`

	Port int64 `json:"port" tf:"port"`
}

type DaxClusterParameters struct {
	AvailabilityZones []string `json:"availabilityZones,omitempty" tf:"availability_zones"`

	ClusterName string `json:"clusterName" tf:"cluster_name"`

	Description *string `json:"description,omitempty" tf:"description"`

	IamRoleArn string `json:"iamRoleArn" tf:"iam_role_arn"`

	MaintenanceWindow *string `json:"maintenanceWindow,omitempty" tf:"maintenance_window"`

	NodeType string `json:"nodeType" tf:"node_type"`

	NotificationTopicArn *string `json:"notificationTopicArn,omitempty" tf:"notification_topic_arn"`

	ParameterGroupName *string `json:"parameterGroupName,omitempty" tf:"parameter_group_name"`

	ReplicationFactor int64 `json:"replicationFactor" tf:"replication_factor"`

	SecurityGroupIds []string `json:"securityGroupIds,omitempty" tf:"security_group_ids"`

	ServerSideEncryption []ServerSideEncryptionParameters `json:"serverSideEncryption,omitempty" tf:"server_side_encryption"`

	SubnetGroupName *string `json:"subnetGroupName,omitempty" tf:"subnet_group_name"`

	Tags map[string]string `json:"tags,omitempty" tf:"tags"`

	TagsAll map[string]string `json:"tagsAll,omitempty" tf:"tags_all"`
}

type NodesObservation struct {
	Address string `json:"address" tf:"address"`

	AvailabilityZone string `json:"availabilityZone" tf:"availability_zone"`

	Id string `json:"id" tf:"id"`

	Port int64 `json:"port" tf:"port"`
}

type NodesParameters struct {
}

type ServerSideEncryptionObservation struct {
}

type ServerSideEncryptionParameters struct {
	Enabled *bool `json:"enabled,omitempty" tf:"enabled"`
}

// DaxClusterSpec defines the desired state of DaxCluster
type DaxClusterSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       DaxClusterParameters `json:"forProvider"`
}

// DaxClusterStatus defines the observed state of DaxCluster.
type DaxClusterStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          DaxClusterObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// DaxCluster is the Schema for the DaxClusters API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type DaxCluster struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DaxClusterSpec   `json:"spec"`
	Status            DaxClusterStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// DaxClusterList contains a list of DaxClusters
type DaxClusterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DaxCluster `json:"items"`
}

// Repository type metadata.
var (
	DaxClusterKind             = "DaxCluster"
	DaxClusterGroupKind        = schema.GroupKind{Group: v1alpha1.Group, Kind: DaxClusterKind}.String()
	DaxClusterKindAPIVersion   = DaxClusterKind + "." + v1alpha1.GroupVersion.String()
	DaxClusterGroupVersionKind = v1alpha1.GroupVersion.WithKind(DaxClusterKind)
)

func init() {
	apis.SchemaMap[v1alpha1.GroupVersion] = append(apis.SchemaMap[v1alpha1.GroupVersion], &DaxCluster{}, &DaxClusterList{})
}
