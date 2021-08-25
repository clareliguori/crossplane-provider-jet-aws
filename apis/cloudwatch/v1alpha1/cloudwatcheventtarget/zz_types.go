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
// +groupName=cloudwatch.aws.tf.crossplane.io
// +versionName=v1alpha1

package v1alpha1

import (
	xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	apis "github.com/crossplane-contrib/provider-tf-aws/apis"
	v1alpha1 "github.com/crossplane-contrib/provider-tf-aws/apis/cloudwatch/v1alpha1"
)

type BatchTargetObservation struct {
}

type BatchTargetParameters struct {
	ArraySize *int64 `json:"arraySize,omitempty" tf:"array_size"`

	JobAttempts *int64 `json:"jobAttempts,omitempty" tf:"job_attempts"`

	JobDefinition string `json:"jobDefinition" tf:"job_definition"`

	JobName string `json:"jobName" tf:"job_name"`
}

type CloudwatchEventTargetObservation struct {
}

type CloudwatchEventTargetParameters struct {
	Arn string `json:"arn" tf:"arn"`

	BatchTarget []BatchTargetParameters `json:"batchTarget,omitempty" tf:"batch_target"`

	DeadLetterConfig []DeadLetterConfigParameters `json:"deadLetterConfig,omitempty" tf:"dead_letter_config"`

	EcsTarget []EcsTargetParameters `json:"ecsTarget,omitempty" tf:"ecs_target"`

	EventBusName *string `json:"eventBusName,omitempty" tf:"event_bus_name"`

	HttpTarget []HttpTargetParameters `json:"httpTarget,omitempty" tf:"http_target"`

	Input *string `json:"input,omitempty" tf:"input"`

	InputPath *string `json:"inputPath,omitempty" tf:"input_path"`

	InputTransformer []InputTransformerParameters `json:"inputTransformer,omitempty" tf:"input_transformer"`

	KinesisTarget []KinesisTargetParameters `json:"kinesisTarget,omitempty" tf:"kinesis_target"`

	RedshiftTarget []RedshiftTargetParameters `json:"redshiftTarget,omitempty" tf:"redshift_target"`

	RetryPolicy []RetryPolicyParameters `json:"retryPolicy,omitempty" tf:"retry_policy"`

	RoleArn *string `json:"roleArn,omitempty" tf:"role_arn"`

	Rule string `json:"rule" tf:"rule"`

	RunCommandTargets []RunCommandTargetsParameters `json:"runCommandTargets,omitempty" tf:"run_command_targets"`

	SqsTarget []SqsTargetParameters `json:"sqsTarget,omitempty" tf:"sqs_target"`

	TargetId *string `json:"targetId,omitempty" tf:"target_id"`
}

type DeadLetterConfigObservation struct {
}

type DeadLetterConfigParameters struct {
	Arn *string `json:"arn,omitempty" tf:"arn"`
}

type EcsTargetObservation struct {
}

type EcsTargetParameters struct {
	EnableEcsManagedTags *bool `json:"enableEcsManagedTags,omitempty" tf:"enable_ecs_managed_tags"`

	EnableExecuteCommand *bool `json:"enableExecuteCommand,omitempty" tf:"enable_execute_command"`

	Group *string `json:"group,omitempty" tf:"group"`

	LaunchType *string `json:"launchType,omitempty" tf:"launch_type"`

	NetworkConfiguration []NetworkConfigurationParameters `json:"networkConfiguration,omitempty" tf:"network_configuration"`

	PlacementConstraint []PlacementConstraintParameters `json:"placementConstraint,omitempty" tf:"placement_constraint"`

	PlatformVersion *string `json:"platformVersion,omitempty" tf:"platform_version"`

	PropagateTags *string `json:"propagateTags,omitempty" tf:"propagate_tags"`

	Tags map[string]string `json:"tags,omitempty" tf:"tags"`

	TaskCount *int64 `json:"taskCount,omitempty" tf:"task_count"`

	TaskDefinitionArn string `json:"taskDefinitionArn" tf:"task_definition_arn"`
}

type HttpTargetObservation struct {
}

type HttpTargetParameters struct {
	HeaderParameters map[string]string `json:"headerParameters,omitempty" tf:"header_parameters"`

	PathParameterValues []string `json:"pathParameterValues,omitempty" tf:"path_parameter_values"`

	QueryStringParameters map[string]string `json:"queryStringParameters,omitempty" tf:"query_string_parameters"`
}

type InputTransformerObservation struct {
}

type InputTransformerParameters struct {
	InputPaths map[string]string `json:"inputPaths,omitempty" tf:"input_paths"`

	InputTemplate string `json:"inputTemplate" tf:"input_template"`
}

type KinesisTargetObservation struct {
}

type KinesisTargetParameters struct {
	PartitionKeyPath *string `json:"partitionKeyPath,omitempty" tf:"partition_key_path"`
}

type NetworkConfigurationObservation struct {
}

type NetworkConfigurationParameters struct {
	AssignPublicIp *bool `json:"assignPublicIp,omitempty" tf:"assign_public_ip"`

	SecurityGroups []string `json:"securityGroups,omitempty" tf:"security_groups"`

	Subnets []string `json:"subnets" tf:"subnets"`
}

type PlacementConstraintObservation struct {
}

type PlacementConstraintParameters struct {
	Expression *string `json:"expression,omitempty" tf:"expression"`

	Type string `json:"type" tf:"type"`
}

type RedshiftTargetObservation struct {
}

type RedshiftTargetParameters struct {
	Database string `json:"database" tf:"database"`

	DbUser *string `json:"dbUser,omitempty" tf:"db_user"`

	SecretsManagerArn *string `json:"secretsManagerArn,omitempty" tf:"secrets_manager_arn"`

	Sql *string `json:"sql,omitempty" tf:"sql"`

	StatementName *string `json:"statementName,omitempty" tf:"statement_name"`

	WithEvent *bool `json:"withEvent,omitempty" tf:"with_event"`
}

type RetryPolicyObservation struct {
}

type RetryPolicyParameters struct {
	MaximumEventAgeInSeconds *int64 `json:"maximumEventAgeInSeconds,omitempty" tf:"maximum_event_age_in_seconds"`

	MaximumRetryAttempts *int64 `json:"maximumRetryAttempts,omitempty" tf:"maximum_retry_attempts"`
}

type RunCommandTargetsObservation struct {
}

type RunCommandTargetsParameters struct {
	Key string `json:"key" tf:"key"`

	Values []string `json:"values" tf:"values"`
}

type SqsTargetObservation struct {
}

type SqsTargetParameters struct {
	MessageGroupId *string `json:"messageGroupId,omitempty" tf:"message_group_id"`
}

// CloudwatchEventTargetSpec defines the desired state of CloudwatchEventTarget
type CloudwatchEventTargetSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       CloudwatchEventTargetParameters `json:"forProvider"`
}

// CloudwatchEventTargetStatus defines the observed state of CloudwatchEventTarget.
type CloudwatchEventTargetStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          CloudwatchEventTargetObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// CloudwatchEventTarget is the Schema for the CloudwatchEventTargets API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type CloudwatchEventTarget struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              CloudwatchEventTargetSpec   `json:"spec"`
	Status            CloudwatchEventTargetStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// CloudwatchEventTargetList contains a list of CloudwatchEventTargets
type CloudwatchEventTargetList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []CloudwatchEventTarget `json:"items"`
}

// Repository type metadata.
var (
	CloudwatchEventTargetKind             = "CloudwatchEventTarget"
	CloudwatchEventTargetGroupKind        = schema.GroupKind{Group: v1alpha1.Group, Kind: CloudwatchEventTargetKind}.String()
	CloudwatchEventTargetKindAPIVersion   = CloudwatchEventTargetKind + "." + v1alpha1.GroupVersion.String()
	CloudwatchEventTargetGroupVersionKind = v1alpha1.GroupVersion.WithKind(CloudwatchEventTargetKind)
)

func init() {
	apis.SchemaMap[v1alpha1.GroupVersion] = append(apis.SchemaMap[v1alpha1.GroupVersion], &CloudwatchEventTarget{}, &CloudwatchEventTargetList{})
}
