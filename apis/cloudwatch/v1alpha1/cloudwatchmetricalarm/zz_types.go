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

type CloudwatchMetricAlarmObservation struct {
	Arn string `json:"arn" tf:"arn"`
}

type CloudwatchMetricAlarmParameters struct {
	ActionsEnabled *bool `json:"actionsEnabled,omitempty" tf:"actions_enabled"`

	AlarmActions []string `json:"alarmActions,omitempty" tf:"alarm_actions"`

	AlarmDescription *string `json:"alarmDescription,omitempty" tf:"alarm_description"`

	AlarmName string `json:"alarmName" tf:"alarm_name"`

	ComparisonOperator string `json:"comparisonOperator" tf:"comparison_operator"`

	DatapointsToAlarm *int64 `json:"datapointsToAlarm,omitempty" tf:"datapoints_to_alarm"`

	Dimensions map[string]string `json:"dimensions,omitempty" tf:"dimensions"`

	EvaluateLowSampleCountPercentiles *string `json:"evaluateLowSampleCountPercentiles,omitempty" tf:"evaluate_low_sample_count_percentiles"`

	EvaluationPeriods int64 `json:"evaluationPeriods" tf:"evaluation_periods"`

	ExtendedStatistic *string `json:"extendedStatistic,omitempty" tf:"extended_statistic"`

	InsufficientDataActions []string `json:"insufficientDataActions,omitempty" tf:"insufficient_data_actions"`

	MetricName *string `json:"metricName,omitempty" tf:"metric_name"`

	MetricQuery []MetricQueryParameters `json:"metricQuery,omitempty" tf:"metric_query"`

	Namespace *string `json:"namespace,omitempty" tf:"namespace"`

	OkActions []string `json:"okActions,omitempty" tf:"ok_actions"`

	Period *int64 `json:"period,omitempty" tf:"period"`

	Statistic *string `json:"statistic,omitempty" tf:"statistic"`

	Tags map[string]string `json:"tags,omitempty" tf:"tags"`

	TagsAll map[string]string `json:"tagsAll,omitempty" tf:"tags_all"`

	Threshold *float64 `json:"threshold,omitempty" tf:"threshold"`

	ThresholdMetricId *string `json:"thresholdMetricId,omitempty" tf:"threshold_metric_id"`

	TreatMissingData *string `json:"treatMissingData,omitempty" tf:"treat_missing_data"`

	Unit *string `json:"unit,omitempty" tf:"unit"`
}

type MetricObservation struct {
}

type MetricParameters struct {
	Dimensions map[string]string `json:"dimensions,omitempty" tf:"dimensions"`

	MetricName string `json:"metricName" tf:"metric_name"`

	Namespace *string `json:"namespace,omitempty" tf:"namespace"`

	Period int64 `json:"period" tf:"period"`

	Stat string `json:"stat" tf:"stat"`

	Unit *string `json:"unit,omitempty" tf:"unit"`
}

type MetricQueryObservation struct {
}

type MetricQueryParameters struct {
	Expression *string `json:"expression,omitempty" tf:"expression"`

	Id string `json:"id" tf:"id"`

	Label *string `json:"label,omitempty" tf:"label"`

	Metric []MetricParameters `json:"metric,omitempty" tf:"metric"`

	ReturnData *bool `json:"returnData,omitempty" tf:"return_data"`
}

// CloudwatchMetricAlarmSpec defines the desired state of CloudwatchMetricAlarm
type CloudwatchMetricAlarmSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       CloudwatchMetricAlarmParameters `json:"forProvider"`
}

// CloudwatchMetricAlarmStatus defines the observed state of CloudwatchMetricAlarm.
type CloudwatchMetricAlarmStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          CloudwatchMetricAlarmObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// CloudwatchMetricAlarm is the Schema for the CloudwatchMetricAlarms API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type CloudwatchMetricAlarm struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              CloudwatchMetricAlarmSpec   `json:"spec"`
	Status            CloudwatchMetricAlarmStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// CloudwatchMetricAlarmList contains a list of CloudwatchMetricAlarms
type CloudwatchMetricAlarmList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []CloudwatchMetricAlarm `json:"items"`
}

// Repository type metadata.
var (
	CloudwatchMetricAlarmKind             = "CloudwatchMetricAlarm"
	CloudwatchMetricAlarmGroupKind        = schema.GroupKind{Group: v1alpha1.Group, Kind: CloudwatchMetricAlarmKind}.String()
	CloudwatchMetricAlarmKindAPIVersion   = CloudwatchMetricAlarmKind + "." + v1alpha1.GroupVersion.String()
	CloudwatchMetricAlarmGroupVersionKind = v1alpha1.GroupVersion.WithKind(CloudwatchMetricAlarmKind)
)

func init() {
	apis.SchemaMap[v1alpha1.GroupVersion] = append(apis.SchemaMap[v1alpha1.GroupVersion], &CloudwatchMetricAlarm{}, &CloudwatchMetricAlarmList{})
}
