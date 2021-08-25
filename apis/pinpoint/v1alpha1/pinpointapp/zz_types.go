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
// +groupName=pinpoint.aws.tf.crossplane.io
// +versionName=v1alpha1

package v1alpha1

import (
	xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	apis "github.com/crossplane-contrib/provider-tf-aws/apis"
	v1alpha1 "github.com/crossplane-contrib/provider-tf-aws/apis/pinpoint/v1alpha1"
)

type CampaignHookObservation struct {
}

type CampaignHookParameters struct {
	LambdaFunctionName *string `json:"lambdaFunctionName,omitempty" tf:"lambda_function_name"`

	Mode *string `json:"mode,omitempty" tf:"mode"`

	WebUrl *string `json:"webUrl,omitempty" tf:"web_url"`
}

type LimitsObservation struct {
}

type LimitsParameters struct {
	Daily *int64 `json:"daily,omitempty" tf:"daily"`

	MaximumDuration *int64 `json:"maximumDuration,omitempty" tf:"maximum_duration"`

	MessagesPerSecond *int64 `json:"messagesPerSecond,omitempty" tf:"messages_per_second"`

	Total *int64 `json:"total,omitempty" tf:"total"`
}

type PinpointAppObservation struct {
	ApplicationId string `json:"applicationId" tf:"application_id"`

	Arn string `json:"arn" tf:"arn"`
}

type PinpointAppParameters struct {
	CampaignHook []CampaignHookParameters `json:"campaignHook,omitempty" tf:"campaign_hook"`

	Limits []LimitsParameters `json:"limits,omitempty" tf:"limits"`

	Name *string `json:"name,omitempty" tf:"name"`

	NamePrefix *string `json:"namePrefix,omitempty" tf:"name_prefix"`

	QuietTime []QuietTimeParameters `json:"quietTime,omitempty" tf:"quiet_time"`

	Tags map[string]string `json:"tags,omitempty" tf:"tags"`

	TagsAll map[string]string `json:"tagsAll,omitempty" tf:"tags_all"`
}

type QuietTimeObservation struct {
}

type QuietTimeParameters struct {
	End *string `json:"end,omitempty" tf:"end"`

	Start *string `json:"start,omitempty" tf:"start"`
}

// PinpointAppSpec defines the desired state of PinpointApp
type PinpointAppSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       PinpointAppParameters `json:"forProvider"`
}

// PinpointAppStatus defines the observed state of PinpointApp.
type PinpointAppStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          PinpointAppObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// PinpointApp is the Schema for the PinpointApps API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type PinpointApp struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              PinpointAppSpec   `json:"spec"`
	Status            PinpointAppStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// PinpointAppList contains a list of PinpointApps
type PinpointAppList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []PinpointApp `json:"items"`
}

// Repository type metadata.
var (
	PinpointAppKind             = "PinpointApp"
	PinpointAppGroupKind        = schema.GroupKind{Group: v1alpha1.Group, Kind: PinpointAppKind}.String()
	PinpointAppKindAPIVersion   = PinpointAppKind + "." + v1alpha1.GroupVersion.String()
	PinpointAppGroupVersionKind = v1alpha1.GroupVersion.WithKind(PinpointAppKind)
)

func init() {
	apis.SchemaMap[v1alpha1.GroupVersion] = append(apis.SchemaMap[v1alpha1.GroupVersion], &PinpointApp{}, &PinpointAppList{})
}
