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

type DbInstanceObservation struct {
	Address string `json:"address" tf:"address"`

	Arn string `json:"arn" tf:"arn"`

	Endpoint string `json:"endpoint" tf:"endpoint"`

	HostedZoneId string `json:"hostedZoneId" tf:"hosted_zone_id"`

	LatestRestorableTime string `json:"latestRestorableTime" tf:"latest_restorable_time"`

	Replicas []string `json:"replicas" tf:"replicas"`

	ResourceId string `json:"resourceId" tf:"resource_id"`

	Status string `json:"status" tf:"status"`
}

type DbInstanceParameters struct {
	AllocatedStorage *int64 `json:"allocatedStorage,omitempty" tf:"allocated_storage"`

	AllowMajorVersionUpgrade *bool `json:"allowMajorVersionUpgrade,omitempty" tf:"allow_major_version_upgrade"`

	ApplyImmediately *bool `json:"applyImmediately,omitempty" tf:"apply_immediately"`

	AutoMinorVersionUpgrade *bool `json:"autoMinorVersionUpgrade,omitempty" tf:"auto_minor_version_upgrade"`

	AvailabilityZone *string `json:"availabilityZone,omitempty" tf:"availability_zone"`

	BackupRetentionPeriod *int64 `json:"backupRetentionPeriod,omitempty" tf:"backup_retention_period"`

	BackupWindow *string `json:"backupWindow,omitempty" tf:"backup_window"`

	CaCertIdentifier *string `json:"caCertIdentifier,omitempty" tf:"ca_cert_identifier"`

	CharacterSetName *string `json:"characterSetName,omitempty" tf:"character_set_name"`

	CopyTagsToSnapshot *bool `json:"copyTagsToSnapshot,omitempty" tf:"copy_tags_to_snapshot"`

	CustomerOwnedIpEnabled *bool `json:"customerOwnedIpEnabled,omitempty" tf:"customer_owned_ip_enabled"`

	DbSubnetGroupName *string `json:"dbSubnetGroupName,omitempty" tf:"db_subnet_group_name"`

	DeleteAutomatedBackups *bool `json:"deleteAutomatedBackups,omitempty" tf:"delete_automated_backups"`

	DeletionProtection *bool `json:"deletionProtection,omitempty" tf:"deletion_protection"`

	Domain *string `json:"domain,omitempty" tf:"domain"`

	DomainIamRoleName *string `json:"domainIamRoleName,omitempty" tf:"domain_iam_role_name"`

	EnabledCloudwatchLogsExports []string `json:"enabledCloudwatchLogsExports,omitempty" tf:"enabled_cloudwatch_logs_exports"`

	Engine *string `json:"engine,omitempty" tf:"engine"`

	EngineVersion *string `json:"engineVersion,omitempty" tf:"engine_version"`

	FinalSnapshotIdentifier *string `json:"finalSnapshotIdentifier,omitempty" tf:"final_snapshot_identifier"`

	IamDatabaseAuthenticationEnabled *bool `json:"iamDatabaseAuthenticationEnabled,omitempty" tf:"iam_database_authentication_enabled"`

	Identifier *string `json:"identifier,omitempty" tf:"identifier"`

	IdentifierPrefix *string `json:"identifierPrefix,omitempty" tf:"identifier_prefix"`

	InstanceClass string `json:"instanceClass" tf:"instance_class"`

	Iops *int64 `json:"iops,omitempty" tf:"iops"`

	KmsKeyId *string `json:"kmsKeyId,omitempty" tf:"kms_key_id"`

	LicenseModel *string `json:"licenseModel,omitempty" tf:"license_model"`

	MaintenanceWindow *string `json:"maintenanceWindow,omitempty" tf:"maintenance_window"`

	MaxAllocatedStorage *int64 `json:"maxAllocatedStorage,omitempty" tf:"max_allocated_storage"`

	MonitoringInterval *int64 `json:"monitoringInterval,omitempty" tf:"monitoring_interval"`

	MonitoringRoleArn *string `json:"monitoringRoleArn,omitempty" tf:"monitoring_role_arn"`

	MultiAz *bool `json:"multiAz,omitempty" tf:"multi_az"`

	Name *string `json:"name,omitempty" tf:"name"`

	NcharCharacterSetName *string `json:"ncharCharacterSetName,omitempty" tf:"nchar_character_set_name"`

	OptionGroupName *string `json:"optionGroupName,omitempty" tf:"option_group_name"`

	ParameterGroupName *string `json:"parameterGroupName,omitempty" tf:"parameter_group_name"`

	Password *string `json:"password,omitempty" tf:"password"`

	PerformanceInsightsEnabled *bool `json:"performanceInsightsEnabled,omitempty" tf:"performance_insights_enabled"`

	PerformanceInsightsKmsKeyId *string `json:"performanceInsightsKmsKeyId,omitempty" tf:"performance_insights_kms_key_id"`

	PerformanceInsightsRetentionPeriod *int64 `json:"performanceInsightsRetentionPeriod,omitempty" tf:"performance_insights_retention_period"`

	Port *int64 `json:"port,omitempty" tf:"port"`

	PubliclyAccessible *bool `json:"publiclyAccessible,omitempty" tf:"publicly_accessible"`

	ReplicateSourceDb *string `json:"replicateSourceDb,omitempty" tf:"replicate_source_db"`

	RestoreToPointInTime []RestoreToPointInTimeParameters `json:"restoreToPointInTime,omitempty" tf:"restore_to_point_in_time"`

	S3Import []S3ImportParameters `json:"s3Import,omitempty" tf:"s3_import"`

	SecurityGroupNames []string `json:"securityGroupNames,omitempty" tf:"security_group_names"`

	SkipFinalSnapshot *bool `json:"skipFinalSnapshot,omitempty" tf:"skip_final_snapshot"`

	SnapshotIdentifier *string `json:"snapshotIdentifier,omitempty" tf:"snapshot_identifier"`

	StorageEncrypted *bool `json:"storageEncrypted,omitempty" tf:"storage_encrypted"`

	StorageType *string `json:"storageType,omitempty" tf:"storage_type"`

	Tags map[string]string `json:"tags,omitempty" tf:"tags"`

	TagsAll map[string]string `json:"tagsAll,omitempty" tf:"tags_all"`

	Timezone *string `json:"timezone,omitempty" tf:"timezone"`

	Username *string `json:"username,omitempty" tf:"username"`

	VpcSecurityGroupIds []string `json:"vpcSecurityGroupIds,omitempty" tf:"vpc_security_group_ids"`
}

type RestoreToPointInTimeObservation struct {
}

type RestoreToPointInTimeParameters struct {
	RestoreTime *string `json:"restoreTime,omitempty" tf:"restore_time"`

	SourceDbInstanceIdentifier *string `json:"sourceDbInstanceIdentifier,omitempty" tf:"source_db_instance_identifier"`

	SourceDbiResourceId *string `json:"sourceDbiResourceId,omitempty" tf:"source_dbi_resource_id"`

	UseLatestRestorableTime *bool `json:"useLatestRestorableTime,omitempty" tf:"use_latest_restorable_time"`
}

type S3ImportObservation struct {
}

type S3ImportParameters struct {
	BucketName string `json:"bucketName" tf:"bucket_name"`

	BucketPrefix *string `json:"bucketPrefix,omitempty" tf:"bucket_prefix"`

	IngestionRole string `json:"ingestionRole" tf:"ingestion_role"`

	SourceEngine string `json:"sourceEngine" tf:"source_engine"`

	SourceEngineVersion string `json:"sourceEngineVersion" tf:"source_engine_version"`
}

// DbInstanceSpec defines the desired state of DbInstance
type DbInstanceSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       DbInstanceParameters `json:"forProvider"`
}

// DbInstanceStatus defines the observed state of DbInstance.
type DbInstanceStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          DbInstanceObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// DbInstance is the Schema for the DbInstances API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type DbInstance struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DbInstanceSpec   `json:"spec"`
	Status            DbInstanceStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// DbInstanceList contains a list of DbInstances
type DbInstanceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DbInstance `json:"items"`
}

// Repository type metadata.
var (
	DbInstanceKind             = "DbInstance"
	DbInstanceGroupKind        = schema.GroupKind{Group: v1alpha1.Group, Kind: DbInstanceKind}.String()
	DbInstanceKindAPIVersion   = DbInstanceKind + "." + v1alpha1.GroupVersion.String()
	DbInstanceGroupVersionKind = v1alpha1.GroupVersion.WithKind(DbInstanceKind)
)

func init() {
	apis.SchemaMap[v1alpha1.GroupVersion] = append(apis.SchemaMap[v1alpha1.GroupVersion], &DbInstance{}, &DbInstanceList{})
}
