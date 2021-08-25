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
// +groupName=s3.aws.tf.crossplane.io
// +versionName=v1alpha1

package v1alpha1

import (
	xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	apis "github.com/crossplane-contrib/provider-tf-aws/apis"
	v1alpha1 "github.com/crossplane-contrib/provider-tf-aws/apis/s3/v1alpha1"
)

type AccessControlTranslationObservation struct {
}

type AccessControlTranslationParameters struct {
	Owner string `json:"owner" tf:"owner"`
}

type ApplyServerSideEncryptionByDefaultObservation struct {
}

type ApplyServerSideEncryptionByDefaultParameters struct {
	KmsMasterKeyId *string `json:"kmsMasterKeyId,omitempty" tf:"kms_master_key_id"`

	SseAlgorithm string `json:"sseAlgorithm" tf:"sse_algorithm"`
}

type CorsRuleObservation struct {
}

type CorsRuleParameters struct {
	AllowedHeaders []string `json:"allowedHeaders,omitempty" tf:"allowed_headers"`

	AllowedMethods []string `json:"allowedMethods" tf:"allowed_methods"`

	AllowedOrigins []string `json:"allowedOrigins" tf:"allowed_origins"`

	ExposeHeaders []string `json:"exposeHeaders,omitempty" tf:"expose_headers"`

	MaxAgeSeconds *int64 `json:"maxAgeSeconds,omitempty" tf:"max_age_seconds"`
}

type DefaultRetentionObservation struct {
}

type DefaultRetentionParameters struct {
	Days *int64 `json:"days,omitempty" tf:"days"`

	Mode string `json:"mode" tf:"mode"`

	Years *int64 `json:"years,omitempty" tf:"years"`
}

type DestinationObservation struct {
}

type DestinationParameters struct {
	AccessControlTranslation []AccessControlTranslationParameters `json:"accessControlTranslation,omitempty" tf:"access_control_translation"`

	AccountId *string `json:"accountId,omitempty" tf:"account_id"`

	Bucket string `json:"bucket" tf:"bucket"`

	ReplicaKmsKeyId *string `json:"replicaKmsKeyId,omitempty" tf:"replica_kms_key_id"`

	StorageClass *string `json:"storageClass,omitempty" tf:"storage_class"`
}

type ExpirationObservation struct {
}

type ExpirationParameters struct {
	Date *string `json:"date,omitempty" tf:"date"`

	Days *int64 `json:"days,omitempty" tf:"days"`

	ExpiredObjectDeleteMarker *bool `json:"expiredObjectDeleteMarker,omitempty" tf:"expired_object_delete_marker"`
}

type FilterObservation struct {
}

type FilterParameters struct {
	Prefix *string `json:"prefix,omitempty" tf:"prefix"`

	Tags map[string]string `json:"tags,omitempty" tf:"tags"`
}

type GrantObservation struct {
}

type GrantParameters struct {
	Id *string `json:"id,omitempty" tf:"id"`

	Permissions []string `json:"permissions" tf:"permissions"`

	Type string `json:"type" tf:"type"`

	Uri *string `json:"uri,omitempty" tf:"uri"`
}

type LifecycleRuleObservation struct {
}

type LifecycleRuleParameters struct {
	AbortIncompleteMultipartUploadDays *int64 `json:"abortIncompleteMultipartUploadDays,omitempty" tf:"abort_incomplete_multipart_upload_days"`

	Enabled bool `json:"enabled" tf:"enabled"`

	Expiration []ExpirationParameters `json:"expiration,omitempty" tf:"expiration"`

	Id *string `json:"id,omitempty" tf:"id"`

	NoncurrentVersionExpiration []NoncurrentVersionExpirationParameters `json:"noncurrentVersionExpiration,omitempty" tf:"noncurrent_version_expiration"`

	NoncurrentVersionTransition []NoncurrentVersionTransitionParameters `json:"noncurrentVersionTransition,omitempty" tf:"noncurrent_version_transition"`

	Prefix *string `json:"prefix,omitempty" tf:"prefix"`

	Tags map[string]string `json:"tags,omitempty" tf:"tags"`

	Transition []TransitionParameters `json:"transition,omitempty" tf:"transition"`
}

type LoggingObservation struct {
}

type LoggingParameters struct {
	TargetBucket string `json:"targetBucket" tf:"target_bucket"`

	TargetPrefix *string `json:"targetPrefix,omitempty" tf:"target_prefix"`
}

type NoncurrentVersionExpirationObservation struct {
}

type NoncurrentVersionExpirationParameters struct {
	Days *int64 `json:"days,omitempty" tf:"days"`
}

type NoncurrentVersionTransitionObservation struct {
}

type NoncurrentVersionTransitionParameters struct {
	Days *int64 `json:"days,omitempty" tf:"days"`

	StorageClass string `json:"storageClass" tf:"storage_class"`
}

type ObjectLockConfigurationObservation struct {
}

type ObjectLockConfigurationParameters struct {
	ObjectLockEnabled string `json:"objectLockEnabled" tf:"object_lock_enabled"`

	Rule []RuleParameters `json:"rule,omitempty" tf:"rule"`
}

type ReplicationConfigurationObservation struct {
}

type ReplicationConfigurationParameters struct {
	Role string `json:"role" tf:"role"`

	Rules []RulesParameters `json:"rules" tf:"rules"`
}

type RuleObservation struct {
}

type RuleParameters struct {
	DefaultRetention []DefaultRetentionParameters `json:"defaultRetention" tf:"default_retention"`
}

type RulesObservation struct {
}

type RulesParameters struct {
	DeleteMarkerReplicationStatus *string `json:"deleteMarkerReplicationStatus,omitempty" tf:"delete_marker_replication_status"`

	Destination []DestinationParameters `json:"destination" tf:"destination"`

	Filter []FilterParameters `json:"filter,omitempty" tf:"filter"`

	Id *string `json:"id,omitempty" tf:"id"`

	Prefix *string `json:"prefix,omitempty" tf:"prefix"`

	Priority *int64 `json:"priority,omitempty" tf:"priority"`

	SourceSelectionCriteria []SourceSelectionCriteriaParameters `json:"sourceSelectionCriteria,omitempty" tf:"source_selection_criteria"`

	Status string `json:"status" tf:"status"`
}

type S3BucketObservation struct {
	BucketDomainName string `json:"bucketDomainName" tf:"bucket_domain_name"`

	BucketRegionalDomainName string `json:"bucketRegionalDomainName" tf:"bucket_regional_domain_name"`

	Region string `json:"region" tf:"region"`
}

type S3BucketParameters struct {
	AccelerationStatus *string `json:"accelerationStatus,omitempty" tf:"acceleration_status"`

	Acl *string `json:"acl,omitempty" tf:"acl"`

	Arn *string `json:"arn,omitempty" tf:"arn"`

	Bucket *string `json:"bucket,omitempty" tf:"bucket"`

	BucketPrefix *string `json:"bucketPrefix,omitempty" tf:"bucket_prefix"`

	CorsRule []CorsRuleParameters `json:"corsRule,omitempty" tf:"cors_rule"`

	ForceDestroy *bool `json:"forceDestroy,omitempty" tf:"force_destroy"`

	Grant []GrantParameters `json:"grant,omitempty" tf:"grant"`

	HostedZoneId *string `json:"hostedZoneId,omitempty" tf:"hosted_zone_id"`

	LifecycleRule []LifecycleRuleParameters `json:"lifecycleRule,omitempty" tf:"lifecycle_rule"`

	Logging []LoggingParameters `json:"logging,omitempty" tf:"logging"`

	ObjectLockConfiguration []ObjectLockConfigurationParameters `json:"objectLockConfiguration,omitempty" tf:"object_lock_configuration"`

	Policy *string `json:"policy,omitempty" tf:"policy"`

	ReplicationConfiguration []ReplicationConfigurationParameters `json:"replicationConfiguration,omitempty" tf:"replication_configuration"`

	RequestPayer *string `json:"requestPayer,omitempty" tf:"request_payer"`

	ServerSideEncryptionConfiguration []ServerSideEncryptionConfigurationParameters `json:"serverSideEncryptionConfiguration,omitempty" tf:"server_side_encryption_configuration"`

	Tags map[string]string `json:"tags,omitempty" tf:"tags"`

	TagsAll map[string]string `json:"tagsAll,omitempty" tf:"tags_all"`

	Versioning []VersioningParameters `json:"versioning,omitempty" tf:"versioning"`

	Website []WebsiteParameters `json:"website,omitempty" tf:"website"`

	WebsiteDomain *string `json:"websiteDomain,omitempty" tf:"website_domain"`

	WebsiteEndpoint *string `json:"websiteEndpoint,omitempty" tf:"website_endpoint"`
}

type ServerSideEncryptionConfigurationObservation struct {
}

type ServerSideEncryptionConfigurationParameters struct {
	Rule []ServerSideEncryptionConfigurationRuleParameters `json:"rule" tf:"rule"`
}

type ServerSideEncryptionConfigurationRuleObservation struct {
}

type ServerSideEncryptionConfigurationRuleParameters struct {
	ApplyServerSideEncryptionByDefault []ApplyServerSideEncryptionByDefaultParameters `json:"applyServerSideEncryptionByDefault" tf:"apply_server_side_encryption_by_default"`

	BucketKeyEnabled *bool `json:"bucketKeyEnabled,omitempty" tf:"bucket_key_enabled"`
}

type SourceSelectionCriteriaObservation struct {
}

type SourceSelectionCriteriaParameters struct {
	SseKmsEncryptedObjects []SseKmsEncryptedObjectsParameters `json:"sseKmsEncryptedObjects,omitempty" tf:"sse_kms_encrypted_objects"`
}

type SseKmsEncryptedObjectsObservation struct {
}

type SseKmsEncryptedObjectsParameters struct {
	Enabled bool `json:"enabled" tf:"enabled"`
}

type TransitionObservation struct {
}

type TransitionParameters struct {
	Date *string `json:"date,omitempty" tf:"date"`

	Days *int64 `json:"days,omitempty" tf:"days"`

	StorageClass string `json:"storageClass" tf:"storage_class"`
}

type VersioningObservation struct {
}

type VersioningParameters struct {
	Enabled *bool `json:"enabled,omitempty" tf:"enabled"`

	MfaDelete *bool `json:"mfaDelete,omitempty" tf:"mfa_delete"`
}

type WebsiteObservation struct {
}

type WebsiteParameters struct {
	ErrorDocument *string `json:"errorDocument,omitempty" tf:"error_document"`

	IndexDocument *string `json:"indexDocument,omitempty" tf:"index_document"`

	RedirectAllRequestsTo *string `json:"redirectAllRequestsTo,omitempty" tf:"redirect_all_requests_to"`

	RoutingRules *string `json:"routingRules,omitempty" tf:"routing_rules"`
}

// S3BucketSpec defines the desired state of S3Bucket
type S3BucketSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       S3BucketParameters `json:"forProvider"`
}

// S3BucketStatus defines the observed state of S3Bucket.
type S3BucketStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          S3BucketObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// S3Bucket is the Schema for the S3Buckets API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type S3Bucket struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              S3BucketSpec   `json:"spec"`
	Status            S3BucketStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// S3BucketList contains a list of S3Buckets
type S3BucketList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []S3Bucket `json:"items"`
}

// Repository type metadata.
var (
	S3BucketKind             = "S3Bucket"
	S3BucketGroupKind        = schema.GroupKind{Group: v1alpha1.Group, Kind: S3BucketKind}.String()
	S3BucketKindAPIVersion   = S3BucketKind + "." + v1alpha1.GroupVersion.String()
	S3BucketGroupVersionKind = v1alpha1.GroupVersion.WithKind(S3BucketKind)
)

func init() {
	apis.SchemaMap[v1alpha1.GroupVersion] = append(apis.SchemaMap[v1alpha1.GroupVersion], &S3Bucket{}, &S3BucketList{})
}
