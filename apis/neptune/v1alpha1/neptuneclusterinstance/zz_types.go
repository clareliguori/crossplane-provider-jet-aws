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
// +groupName=neptune.aws.tf.crossplane.io
// +versionName=v1alpha1

package v1alpha1

import (
	xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	apis "github.com/crossplane-contrib/provider-tf-aws/apis"
	v1alpha1 "github.com/crossplane-contrib/provider-tf-aws/apis/neptune/v1alpha1"
)

type NeptuneClusterInstanceObservation struct {
	Address string `json:"address" tf:"address"`

	Arn string `json:"arn" tf:"arn"`

	DbiResourceId string `json:"dbiResourceId" tf:"dbi_resource_id"`

	Endpoint string `json:"endpoint" tf:"endpoint"`

	KmsKeyArn string `json:"kmsKeyArn" tf:"kms_key_arn"`

	StorageEncrypted bool `json:"storageEncrypted" tf:"storage_encrypted"`

	Writer bool `json:"writer" tf:"writer"`
}

type NeptuneClusterInstanceParameters struct {
	ApplyImmediately *bool `json:"applyImmediately,omitempty" tf:"apply_immediately"`

	AutoMinorVersionUpgrade *bool `json:"autoMinorVersionUpgrade,omitempty" tf:"auto_minor_version_upgrade"`

	AvailabilityZone *string `json:"availabilityZone,omitempty" tf:"availability_zone"`

	ClusterIdentifier string `json:"clusterIdentifier" tf:"cluster_identifier"`

	Engine *string `json:"engine,omitempty" tf:"engine"`

	EngineVersion *string `json:"engineVersion,omitempty" tf:"engine_version"`

	Identifier *string `json:"identifier,omitempty" tf:"identifier"`

	IdentifierPrefix *string `json:"identifierPrefix,omitempty" tf:"identifier_prefix"`

	InstanceClass string `json:"instanceClass" tf:"instance_class"`

	NeptuneParameterGroupName *string `json:"neptuneParameterGroupName,omitempty" tf:"neptune_parameter_group_name"`

	NeptuneSubnetGroupName *string `json:"neptuneSubnetGroupName,omitempty" tf:"neptune_subnet_group_name"`

	Port *int64 `json:"port,omitempty" tf:"port"`

	PreferredBackupWindow *string `json:"preferredBackupWindow,omitempty" tf:"preferred_backup_window"`

	PreferredMaintenanceWindow *string `json:"preferredMaintenanceWindow,omitempty" tf:"preferred_maintenance_window"`

	PromotionTier *int64 `json:"promotionTier,omitempty" tf:"promotion_tier"`

	PubliclyAccessible *bool `json:"publiclyAccessible,omitempty" tf:"publicly_accessible"`

	Tags map[string]string `json:"tags,omitempty" tf:"tags"`

	TagsAll map[string]string `json:"tagsAll,omitempty" tf:"tags_all"`
}

// NeptuneClusterInstanceSpec defines the desired state of NeptuneClusterInstance
type NeptuneClusterInstanceSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       NeptuneClusterInstanceParameters `json:"forProvider"`
}

// NeptuneClusterInstanceStatus defines the observed state of NeptuneClusterInstance.
type NeptuneClusterInstanceStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          NeptuneClusterInstanceObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// NeptuneClusterInstance is the Schema for the NeptuneClusterInstances API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type NeptuneClusterInstance struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              NeptuneClusterInstanceSpec   `json:"spec"`
	Status            NeptuneClusterInstanceStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// NeptuneClusterInstanceList contains a list of NeptuneClusterInstances
type NeptuneClusterInstanceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []NeptuneClusterInstance `json:"items"`
}

// Repository type metadata.
var (
	NeptuneClusterInstanceKind             = "NeptuneClusterInstance"
	NeptuneClusterInstanceGroupKind        = schema.GroupKind{Group: v1alpha1.Group, Kind: NeptuneClusterInstanceKind}.String()
	NeptuneClusterInstanceKindAPIVersion   = NeptuneClusterInstanceKind + "." + v1alpha1.GroupVersion.String()
	NeptuneClusterInstanceGroupVersionKind = v1alpha1.GroupVersion.WithKind(NeptuneClusterInstanceKind)
)

func init() {
	apis.SchemaMap[v1alpha1.GroupVersion] = append(apis.SchemaMap[v1alpha1.GroupVersion], &NeptuneClusterInstance{}, &NeptuneClusterInstanceList{})
}
