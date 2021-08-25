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
// +groupName=ses.aws.tf.crossplane.io
// +versionName=v1alpha1

package v1alpha1

import (
	xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	apis "github.com/crossplane-contrib/provider-tf-aws/apis"
	v1alpha1 "github.com/crossplane-contrib/provider-tf-aws/apis/ses/v1alpha1"
)

type AddHeaderActionObservation struct {
}

type AddHeaderActionParameters struct {
	HeaderName string `json:"headerName" tf:"header_name"`

	HeaderValue string `json:"headerValue" tf:"header_value"`

	Position int64 `json:"position" tf:"position"`
}

type BounceActionObservation struct {
}

type BounceActionParameters struct {
	Message string `json:"message" tf:"message"`

	Position int64 `json:"position" tf:"position"`

	Sender string `json:"sender" tf:"sender"`

	SmtpReplyCode string `json:"smtpReplyCode" tf:"smtp_reply_code"`

	StatusCode *string `json:"statusCode,omitempty" tf:"status_code"`

	TopicArn *string `json:"topicArn,omitempty" tf:"topic_arn"`
}

type LambdaActionObservation struct {
}

type LambdaActionParameters struct {
	FunctionArn string `json:"functionArn" tf:"function_arn"`

	InvocationType *string `json:"invocationType,omitempty" tf:"invocation_type"`

	Position int64 `json:"position" tf:"position"`

	TopicArn *string `json:"topicArn,omitempty" tf:"topic_arn"`
}

type S3ActionObservation struct {
}

type S3ActionParameters struct {
	BucketName string `json:"bucketName" tf:"bucket_name"`

	KmsKeyArn *string `json:"kmsKeyArn,omitempty" tf:"kms_key_arn"`

	ObjectKeyPrefix *string `json:"objectKeyPrefix,omitempty" tf:"object_key_prefix"`

	Position int64 `json:"position" tf:"position"`

	TopicArn *string `json:"topicArn,omitempty" tf:"topic_arn"`
}

type SesReceiptRuleObservation struct {
	Arn string `json:"arn" tf:"arn"`
}

type SesReceiptRuleParameters struct {
	AddHeaderAction []AddHeaderActionParameters `json:"addHeaderAction,omitempty" tf:"add_header_action"`

	After *string `json:"after,omitempty" tf:"after"`

	BounceAction []BounceActionParameters `json:"bounceAction,omitempty" tf:"bounce_action"`

	Enabled *bool `json:"enabled,omitempty" tf:"enabled"`

	LambdaAction []LambdaActionParameters `json:"lambdaAction,omitempty" tf:"lambda_action"`

	Name string `json:"name" tf:"name"`

	Recipients []string `json:"recipients,omitempty" tf:"recipients"`

	RuleSetName string `json:"ruleSetName" tf:"rule_set_name"`

	S3Action []S3ActionParameters `json:"s3Action,omitempty" tf:"s3_action"`

	ScanEnabled *bool `json:"scanEnabled,omitempty" tf:"scan_enabled"`

	SnsAction []SnsActionParameters `json:"snsAction,omitempty" tf:"sns_action"`

	StopAction []StopActionParameters `json:"stopAction,omitempty" tf:"stop_action"`

	TlsPolicy *string `json:"tlsPolicy,omitempty" tf:"tls_policy"`

	WorkmailAction []WorkmailActionParameters `json:"workmailAction,omitempty" tf:"workmail_action"`
}

type SnsActionObservation struct {
}

type SnsActionParameters struct {
	Encoding *string `json:"encoding,omitempty" tf:"encoding"`

	Position int64 `json:"position" tf:"position"`

	TopicArn string `json:"topicArn" tf:"topic_arn"`
}

type StopActionObservation struct {
}

type StopActionParameters struct {
	Position int64 `json:"position" tf:"position"`

	Scope string `json:"scope" tf:"scope"`

	TopicArn *string `json:"topicArn,omitempty" tf:"topic_arn"`
}

type WorkmailActionObservation struct {
}

type WorkmailActionParameters struct {
	OrganizationArn string `json:"organizationArn" tf:"organization_arn"`

	Position int64 `json:"position" tf:"position"`

	TopicArn *string `json:"topicArn,omitempty" tf:"topic_arn"`
}

// SesReceiptRuleSpec defines the desired state of SesReceiptRule
type SesReceiptRuleSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       SesReceiptRuleParameters `json:"forProvider"`
}

// SesReceiptRuleStatus defines the observed state of SesReceiptRule.
type SesReceiptRuleStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          SesReceiptRuleObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// SesReceiptRule is the Schema for the SesReceiptRules API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type SesReceiptRule struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              SesReceiptRuleSpec   `json:"spec"`
	Status            SesReceiptRuleStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SesReceiptRuleList contains a list of SesReceiptRules
type SesReceiptRuleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SesReceiptRule `json:"items"`
}

// Repository type metadata.
var (
	SesReceiptRuleKind             = "SesReceiptRule"
	SesReceiptRuleGroupKind        = schema.GroupKind{Group: v1alpha1.Group, Kind: SesReceiptRuleKind}.String()
	SesReceiptRuleKindAPIVersion   = SesReceiptRuleKind + "." + v1alpha1.GroupVersion.String()
	SesReceiptRuleGroupVersionKind = v1alpha1.GroupVersion.WithKind(SesReceiptRuleKind)
)

func init() {
	apis.SchemaMap[v1alpha1.GroupVersion] = append(apis.SchemaMap[v1alpha1.GroupVersion], &SesReceiptRule{}, &SesReceiptRuleList{})
}
