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

type RdsClusterObservation struct {
	Arn string `json:"arn" tf:"arn"`

	ClusterResourceID string `json:"clusterResourceId" tf:"cluster_resource_id"`

	Endpoint string `json:"endpoint" tf:"endpoint"`

	EngineVersionActual string `json:"engineVersionActual" tf:"engine_version_actual"`

	HostedZoneID string `json:"hostedZoneId" tf:"hosted_zone_id"`

	ReaderEndpoint string `json:"readerEndpoint" tf:"reader_endpoint"`
}

type RdsClusterParameters struct {

	// +kubebuilder:validation:Optional
	AllowMajorVersionUpgrade *bool `json:"allowMajorVersionUpgrade,omitempty" tf:"allow_major_version_upgrade"`

	// +kubebuilder:validation:Optional
	ApplyImmediately *bool `json:"applyImmediately,omitempty" tf:"apply_immediately"`

	// +kubebuilder:validation:Optional
	AvailabilityZones []string `json:"availabilityZones,omitempty" tf:"availability_zones"`

	// +kubebuilder:validation:Optional
	BacktrackWindow *int64 `json:"backtrackWindow,omitempty" tf:"backtrack_window"`

	// +kubebuilder:validation:Optional
	BackupRetentionPeriod *int64 `json:"backupRetentionPeriod,omitempty" tf:"backup_retention_period"`

	// +kubebuilder:validation:Optional
	ClusterIdentifier *string `json:"clusterIdentifier,omitempty" tf:"cluster_identifier"`

	// +kubebuilder:validation:Optional
	ClusterIdentifierPrefix *string `json:"clusterIdentifierPrefix,omitempty" tf:"cluster_identifier_prefix"`

	// +kubebuilder:validation:Optional
	ClusterMembers []string `json:"clusterMembers,omitempty" tf:"cluster_members"`

	// +kubebuilder:validation:Optional
	CopyTagsToSnapshot *bool `json:"copyTagsToSnapshot,omitempty" tf:"copy_tags_to_snapshot"`

	// +kubebuilder:validation:Optional
	DBClusterParameterGroupName *string `json:"dbClusterParameterGroupName,omitempty" tf:"db_cluster_parameter_group_name"`

	// +kubebuilder:validation:Optional
	DBSubnetGroupName *string `json:"dbSubnetGroupName,omitempty" tf:"db_subnet_group_name"`

	// +kubebuilder:validation:Optional
	DatabaseName *string `json:"databaseName,omitempty" tf:"database_name"`

	// +kubebuilder:validation:Optional
	DeletionProtection *bool `json:"deletionProtection,omitempty" tf:"deletion_protection"`

	// +kubebuilder:validation:Optional
	EnableHTTPEndpoint *bool `json:"enableHttpEndpoint,omitempty" tf:"enable_http_endpoint"`

	// +kubebuilder:validation:Optional
	EnabledCloudwatchLogsExports []string `json:"enabledCloudwatchLogsExports,omitempty" tf:"enabled_cloudwatch_logs_exports"`

	// +kubebuilder:validation:Optional
	Engine *string `json:"engine,omitempty" tf:"engine"`

	// +kubebuilder:validation:Optional
	EngineMode *string `json:"engineMode,omitempty" tf:"engine_mode"`

	// +kubebuilder:validation:Optional
	EngineVersion *string `json:"engineVersion,omitempty" tf:"engine_version"`

	// +kubebuilder:validation:Optional
	FinalSnapshotIdentifier *string `json:"finalSnapshotIdentifier,omitempty" tf:"final_snapshot_identifier"`

	// +kubebuilder:validation:Optional
	GlobalClusterIdentifier *string `json:"globalClusterIdentifier,omitempty" tf:"global_cluster_identifier"`

	// +kubebuilder:validation:Optional
	IamDatabaseAuthenticationEnabled *bool `json:"iamDatabaseAuthenticationEnabled,omitempty" tf:"iam_database_authentication_enabled"`

	// +kubebuilder:validation:Optional
	IamRoles []string `json:"iamRoles,omitempty" tf:"iam_roles"`

	// +kubebuilder:validation:Optional
	KmsKeyID *string `json:"kmsKeyId,omitempty" tf:"kms_key_id"`

	// +kubebuilder:validation:Optional
	MasterPassword *string `json:"masterPassword,omitempty" tf:"master_password"`

	// +kubebuilder:validation:Optional
	MasterUsername *string `json:"masterUsername,omitempty" tf:"master_username"`

	// +kubebuilder:validation:Optional
	Port *int64 `json:"port,omitempty" tf:"port"`

	// +kubebuilder:validation:Optional
	PreferredBackupWindow *string `json:"preferredBackupWindow,omitempty" tf:"preferred_backup_window"`

	// +kubebuilder:validation:Optional
	PreferredMaintenanceWindow *string `json:"preferredMaintenanceWindow,omitempty" tf:"preferred_maintenance_window"`

	// Region is the region you'd like your resource to be created in.
	// +terrajet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region string `json:"region" tf:"-"`

	// +kubebuilder:validation:Optional
	ReplicationSourceIdentifier *string `json:"replicationSourceIdentifier,omitempty" tf:"replication_source_identifier"`

	// +kubebuilder:validation:Optional
	RestoreToPointInTime []RestoreToPointInTimeParameters `json:"restoreToPointInTime,omitempty" tf:"restore_to_point_in_time"`

	// +kubebuilder:validation:Optional
	S3Import []S3ImportParameters `json:"s3Import,omitempty" tf:"s3_import"`

	// +kubebuilder:validation:Optional
	ScalingConfiguration []ScalingConfigurationParameters `json:"scalingConfiguration,omitempty" tf:"scaling_configuration"`

	// +kubebuilder:validation:Optional
	SkipFinalSnapshot *bool `json:"skipFinalSnapshot,omitempty" tf:"skip_final_snapshot"`

	// +kubebuilder:validation:Optional
	SnapshotIdentifier *string `json:"snapshotIdentifier,omitempty" tf:"snapshot_identifier"`

	// +kubebuilder:validation:Optional
	SourceRegion *string `json:"sourceRegion,omitempty" tf:"source_region"`

	// +kubebuilder:validation:Optional
	StorageEncrypted *bool `json:"storageEncrypted,omitempty" tf:"storage_encrypted"`

	// +kubebuilder:validation:Optional
	Tags map[string]string `json:"tags,omitempty" tf:"tags"`

	// +kubebuilder:validation:Optional
	TagsAll map[string]string `json:"tagsAll,omitempty" tf:"tags_all"`

	// +kubebuilder:validation:Optional
	VpcSecurityGroupIds []string `json:"vpcSecurityGroupIds,omitempty" tf:"vpc_security_group_ids"`
}

type RestoreToPointInTimeObservation struct {
}

type RestoreToPointInTimeParameters struct {

	// +kubebuilder:validation:Optional
	RestoreToTime *string `json:"restoreToTime,omitempty" tf:"restore_to_time"`

	// +kubebuilder:validation:Optional
	RestoreType *string `json:"restoreType,omitempty" tf:"restore_type"`

	// +kubebuilder:validation:Required
	SourceClusterIdentifier string `json:"sourceClusterIdentifier" tf:"source_cluster_identifier"`

	// +kubebuilder:validation:Optional
	UseLatestRestorableTime *bool `json:"useLatestRestorableTime,omitempty" tf:"use_latest_restorable_time"`
}

type S3ImportObservation struct {
}

type S3ImportParameters struct {

	// +kubebuilder:validation:Required
	BucketName string `json:"bucketName" tf:"bucket_name"`

	// +kubebuilder:validation:Optional
	BucketPrefix *string `json:"bucketPrefix,omitempty" tf:"bucket_prefix"`

	// +kubebuilder:validation:Required
	IngestionRole string `json:"ingestionRole" tf:"ingestion_role"`

	// +kubebuilder:validation:Required
	SourceEngine string `json:"sourceEngine" tf:"source_engine"`

	// +kubebuilder:validation:Required
	SourceEngineVersion string `json:"sourceEngineVersion" tf:"source_engine_version"`
}

type ScalingConfigurationObservation struct {
}

type ScalingConfigurationParameters struct {

	// +kubebuilder:validation:Optional
	AutoPause *bool `json:"autoPause,omitempty" tf:"auto_pause"`

	// +kubebuilder:validation:Optional
	MaxCapacity *int64 `json:"maxCapacity,omitempty" tf:"max_capacity"`

	// +kubebuilder:validation:Optional
	MinCapacity *int64 `json:"minCapacity,omitempty" tf:"min_capacity"`

	// +kubebuilder:validation:Optional
	SecondsUntilAutoPause *int64 `json:"secondsUntilAutoPause,omitempty" tf:"seconds_until_auto_pause"`

	// +kubebuilder:validation:Optional
	TimeoutAction *string `json:"timeoutAction,omitempty" tf:"timeout_action"`
}

// RdsClusterSpec defines the desired state of RdsCluster
type RdsClusterSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       RdsClusterParameters `json:"forProvider"`
}

// RdsClusterStatus defines the observed state of RdsCluster.
type RdsClusterStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          RdsClusterObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// RdsCluster is the Schema for the RdsClusters API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tfaws}
type RdsCluster struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              RdsClusterSpec   `json:"spec"`
	Status            RdsClusterStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// RdsClusterList contains a list of RdsClusters
type RdsClusterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []RdsCluster `json:"items"`
}

// Repository type metadata.
var (
	RdsClusterKind             = "RdsCluster"
	RdsClusterGroupKind        = schema.GroupKind{Group: Group, Kind: RdsClusterKind}.String()
	RdsClusterKindAPIVersion   = RdsClusterKind + "." + GroupVersion.String()
	RdsClusterGroupVersionKind = GroupVersion.WithKind(RdsClusterKind)
)

func init() {
	SchemeBuilder.Register(&RdsCluster{}, &RdsClusterList{})
}